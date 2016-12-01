package DatabaseConnector

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	mdl "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	_ "github.com/lib/pq"
)

type DatabaseConnector struct {
	Database          *sql.DB
	CheckDevEUISTMT   *sql.Stmt
	GetNodeHeaderSTMT *sql.Stmt
	InsertMessageSTMT *sql.Stmt
}

type WorkRequest struct {
	Query         string
	ResultChannel chan (WorkResult)
	Arguments     []interface{}
	F             func(w *WorkRequest)
}

type WorkResult struct {
	Result interface{}
	err    error
}

func newDBC(db *sql.DB) *DatabaseConnector {
	dbc := DatabaseConnector{}
	dbc.Database = db
	return &dbc
}

var instantiated *DatabaseConnector
var once sync.Once
var WorkQueue = make(chan WorkRequest, 100)

// Get the instantiated instance of the DatabaseConnector or create it.
func GetInstance() *DatabaseConnector {
	once.Do(func() {
		settings := mdl.GetConfiguration().Db
		dbConnectionInfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%v sslmode=disable", settings.User, settings.Password, settings.Name, settings.Port)
		println(dbConnectionInfo)
		actualDb, err := sql.Open("postgres", dbConnectionInfo)
		instantiated = newDBC(actualDb)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("not connected yet")
		StartDispatcher(settings.NumberOfWorkers)
	})
	return instantiated
}

func Connect() error {
	db := GetInstance()
	err := db.Database.Ping()
	if err != nil {
		return err
	}
	db.CheckDevEUISTMT, err = db.Database.Prepare("select exists(select 1 from nodes where deveui=$1)")
	if err != nil {
		return err
	}
	db.GetNodeHeaderSTMT, err = db.Database.Prepare("select sensors.id, number_of_values, lenght_of_values, header_order, conversion_expression, data_type " +
		"from sensors " +
		"join public.sensortypes on sensors.sensortype_id = sensortypes.id " +
		"where deveui =$1 " +
		"order by header_order asc;")
	if err != nil {
		return err
	}
	db.InsertMessageSTMT, err = db.Database.Prepare("INSERT INTO public.messages(" +
		"deveui, created_at, down) " +
		"VALUES ($1, NOW(), false) " +
		"RETURNING id;")
	return err
}

func Close() error {
	return GetInstance().Database.Close()
}

//Not used, but can be used as example
// Checks if the devEUI exists in the database.
// Uses a database worker to execute the query.
func CheckDevEUI(devEUI string) bool {
	//log.Println(devEUI)
	result := make(chan WorkResult)
	args := make([]interface{}, 1)
	args[0] = devEUI
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		rows, err := GetInstance().CheckDevEUISTMT.Query(w.Arguments...)
		defer rows.Close()
		checkErr(err)
		if err != nil {
			w.ResultChannel <- WorkResult{Result: false, err: err}
			return
		}
		rows.Next()
		var exisistsboolean = false
		err = rows.Scan(&exisistsboolean)
		w.ResultChannel <- WorkResult{Result: exisistsboolean, err: err}
	}}
	defer close(result)
	var workResult = <-result
	checkErr(workResult.err)
	exists := workResult.Result.(bool)
	log.Printf("Work done for Q:%s A:%t\n", devEUI, exists)
	return exists
}

//Store a message
func AddMessage(devEUI string) (mdl.MessageUplinkI, error) {
	//create response channel
	result := make(chan WorkResult)
	defer close(result)
	//create arguments
	args := make([]interface{}, 1)
	args[0] = devEUI
	//create and add new WorkRequest
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		var messageId int64
		rows, err := GetInstance().InsertMessageSTMT.Query(w.Arguments...)

		checkErr(err)
		if err != nil {
			log.Println("check errors1")
			w.ResultChannel <- WorkResult{Result: 0, err: err}
			return
		}
		defer rows.Close()

		rows.Next()
		err = rows.Scan(&messageId)
		if err != nil {
			w.ResultChannel <- WorkResult{Result: 0, err: err}
		}
		w.ResultChannel <- WorkResult{Result: messageId, err: nil}
	}}
	response := <-result
	if response.err != nil {
		log.Println("check errors3")
		return nil, response.err
	}
	message := mdl.NewMessageUplink(response.Result.(int64), devEUI)

	return message, response.err
}

//Get a message
// TODO define return type
func GetMessage(MessageId []int) {
	// TODO define return type
}

//Get messages from one node
// TODO define return type
func GetNodeMessages(NodeId, maxMessages int) {
	// TODO
}

//Get sensors that belong to one node
func GetNodeSensors(devEUI string) []mdl.Sensor {
	result := make(chan WorkResult)
	defer close(result)
	args := make([]interface{}, 1)
	args[0] = devEUI
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		//rows, err := GetInstance().Database.Query(w.Query, w.Arguments...)
		rows, err := GetInstance().GetNodeHeaderSTMT.Query(w.Arguments...)
		defer rows.Close()
		checkErr(err)
		if err != nil {
			w.ResultChannel <- WorkResult{Result: false, err: err}
			return
		}
		sensors := make([]mdl.Sensor, 0)
		var id int64
		var number_of_values, lenght_of_values, header_order, data_type int
		var conversion_expression string

		for rows.Next() {
			err = rows.Scan(&id, &number_of_values, &lenght_of_values, &header_order, &conversion_expression, &data_type)
			if err != nil {
				panic(err.Error())
			}
			s := mdl.NewHeaderSensor(
				id,
				number_of_values,
				lenght_of_values,
				header_order,
				data_type,
				conversion_expression)
			sensors = append(sensors, s)
		}
		w.ResultChannel <- WorkResult{Result: sensors, err: err}
	}}
	var workResult = <-result
	checkErr(workResult.err)
	sensors := workResult.Result.([]mdl.Sensor)
	return sensors
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func panicErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
