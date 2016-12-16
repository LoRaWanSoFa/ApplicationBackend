//This package handels the communication with the database.
//It only contains prepared querys that are needed for the backend.
//It has a maximum amount of querys that can happen simultaniously
package DatabaseConnector

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	mdl "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	_ "github.com/lib/pq"
)

type DatabaseConnector struct {
	Database                        *sql.DB
	checkDevEUISTMT                 *sql.Stmt
	getNodeHeaderSTMT               *sql.Stmt
	insertMessageSTMT               *sql.Stmt
	insertPayloadSTMT               *sql.Stmt
	insertDownlinkMessageSTMT       *sql.Stmt
	getFullHeaderSTMT               *sql.Stmt
	changeSensorActivationStateSTMT *sql.Stmt
	checkSensorTypeSTMT             *sql.Stmt
	addSensorTypeSTMT               *sql.Stmt
	addSensorSTMT                   *sql.Stmt
	updateSensorOrderSTMT           *sql.Stmt
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
		dbConnectionInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=disable", settings.User, settings.Password, settings.Name, settings.Network, settings.Port)
		println(dbConnectionInfo)
		actualDb, err := sql.Open("postgres", dbConnectionInfo)
		instantiated = newDBC(actualDb)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("not connected yet")
		startDispatcher(settings.NumberOfWorkers)
	})
	return instantiated
}

func Connect() error {
	//IDEA move sql.Open to here so we can reconnect.
	db := GetInstance()
	err := db.Database.Ping()
	if err != nil {
		return err
	}
	time.Sleep(1000 * time.Millisecond) //Wait for connections to finnish setting up. 1 sec is propably too long, can be finetuned later
	db.checkDevEUISTMT, err = db.Database.Prepare("select exists(select 1 from nodes where deveui=$1)")
	if err != nil {
		return err
	}
	db.getNodeHeaderSTMT, err = db.Database.Prepare("select sensors.id, number_of_values, lenght_of_values, header_order, conversion_expression, data_type " +
		"from sensors " +
		"join public.sensortypes on sensors.sensortype_id = sensortypes.id " +
		"where deveui =$1 and header_order >= 0" +
		"order by header_order asc;")
	if err != nil {
		return err
	}
	db.insertMessageSTMT, err = db.Database.Prepare("INSERT INTO public.messages(" +
		"deveui, created_at, down) " +
		"VALUES ($1, NOW(), false) " +
		"RETURNING id;")
	if err != nil {
		return err
	}
	db.insertDownlinkMessageSTMT, err = db.Database.Prepare("INSERT INTO public.messages(" +
		"deveui, created_at, down) " +
		"VALUES ($1, $2, true) " +
		"RETURNING id;")
	if err != nil {
		return err
	}
	db.getFullHeaderSTMT, err = db.Database.Prepare("select sensors.id, sensortypes.id, io_address, io_type, number_of_values, lenght_of_values, header_order, conversion_expression, description, data_type, sensor_type, soft_deleted " +
		"from sensors " +
		"join public.sensortypes on sensors.sensortype_id = sensortypes.id " +
		"where deveui =$1 " +
		"order by soft_deleted, header_order;")
	if err != nil {
		return err
	}
	db.checkSensorTypeSTMT, err = db.Database.Prepare("select id from sensortypes where sensor_type=$1 limit 1")
	if err != nil {
		return err
	}
	db.changeSensorActivationStateSTMT, err = db.Database.Prepare("update sensors set soft_deleted=$1 where id=$2")
	if err != nil {
		return err
	}
	db.insertPayloadSTMT, err = db.Database.Prepare("INSERT INTO public.message_payloads(" +
		"message_id, sensor_id, payload) " +
		"VALUES ($1, $2, $3);")
	return err
}

func Close() error {
	return GetInstance().Database.Close()
}

// Checks if the devEUI exists in the database.
// Uses a database worker to execute the query.
func CheckDevEUI(devEUI string) bool {
	//log.Println(devEUI)
	result := make(chan WorkResult)
	args := make([]interface{}, 1)
	args[0] = devEUI
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		rows, err := GetInstance().checkDevEUISTMT.Query(w.Arguments...)
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
		rows, err := GetInstance().insertMessageSTMT.Query(w.Arguments...)

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
func StoreMessagePayloads(message mdl.MessageUplinkI) error {
	if message == nil {
		return errors.New("nil given as message parameter")
	}
	if message.GetId() == 0 {
		return errors.New("Message has not been initalized/stored")
	}
	payloads := message.GetPayloads()
	if len(payloads) == 0 {
		return errors.New("Nothing to store!")
	}
	var parameters []interface{}
	var err error
	for _, payload := range payloads {
		parameters = make([]interface{}, 0)
		parameters = append(parameters, message.GetId())               //message id
		parameters = append(parameters, payload.GetSensor().Id)        //sensor id
		parameters = append(parameters, payload.GetPayload().(string)) //payload
		//log.Printf("parameters: %+v", parameters)
		err = insertPayload(parameters)
	}
	if err != nil {
		return nil
	}
	return nil
}

func insertPayload(parameters []interface{}) error {
	result := make(chan WorkResult)
	defer close(result)
	WorkQueue <- WorkRequest{Query: "", Arguments: parameters, ResultChannel: result, F: func(w *WorkRequest) {
		rows, err := GetInstance().insertPayloadSTMT.Query(w.Arguments...)

		checkErr(err)
		if err != nil {
			log.Println("Query could not be executed!")
			w.ResultChannel <- WorkResult{Result: nil, err: err}
			return
		}
		defer rows.Close()

		rows.Next()
		if err != nil {
			w.ResultChannel <- WorkResult{Result: nil, err: err}
		}
		w.ResultChannel <- WorkResult{Result: true, err: nil}
	}}
	response := <-result
	if response.err != nil {
		log.Println("The worker could not execute the work properly")
		return response.err
	}
	return nil
}

//Stores a DownlinkMessage which has an id,payload and deveui set.
//if no time is set NOW() will be used
func StoreDownlinkMessage(message *mdl.MessageDownLink) error {
	if message.Id != 0 {
		return errors.New("Message already has an id, can not insert it")
	}
	if message.Payload == "" {
		return errors.New("Message has an empty payload")
	}
	if message.Deveui == "" {
		return errors.New("Message has no DevEUI set")
	}
	if message.Time.IsZero() {
		message.Time = time.Now()
	}
	err := addDownlinkMessage(message)
	if err != nil {
		return err
	}
	var parameters []interface{}
	parameters = append(parameters, message.Id)
	parameters = append(parameters, nil)
	parameters = append(parameters, message.Payload)
	err = insertPayload(parameters)
	return err
}

// Executes the query to insert the message
// Formats the time to UTC and rounds it to second percision.
// Sets a new id for the message
func addDownlinkMessage(message *mdl.MessageDownLink) error {
	//create response channel
	result := make(chan WorkResult)
	defer close(result)
	//create arguments
	args := make([]interface{}, 2)
	args[0] = message.Deveui
	args[1] = message.Time.UTC().Round(time.Second).Format(time.RFC3339)
	//create and add new WorkRequest
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		var messageId int64
		rows, err := GetInstance().insertDownlinkMessageSTMT.Query(w.Arguments...)

		checkErr(err)
		if err != nil {
			log.Println("Query could not be executed")
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
		log.Println("Worker could finnish its work properly")
		return response.err
	}
	message.Id = response.Result.(int64)

	return response.err
}

//Get sensors that belong to one node
func GetNodeSensors(devEUI string) []mdl.Sensor {
	result := make(chan WorkResult)
	defer close(result)
	args := make([]interface{}, 1)
	args[0] = devEUI
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		//rows, err := GetInstance().Database.Query(w.Query, w.Arguments...)
		rows, err := GetInstance().getNodeHeaderSTMT.Query(w.Arguments...)
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
	if workResult.err != nil {
		log.Printf("A problem occured when getting the sensorheaders: %+v", workResult.err)
		return make([]mdl.Sensor, 0)
	}
	sensors := workResult.Result.([]mdl.Sensor)
	return sensors
}

//GetFullHeader returns all sensors connected to a node.
func GetFullHeader(devEUI string) ([]mdl.Sensor, error) {
	result := make(chan WorkResult)
	defer close(result)
	args := make([]interface{}, 1)
	args[0] = devEUI

	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		rows, err := GetInstance().getFullHeaderSTMT.Query(w.Arguments...)
		defer rows.Close()
		checkErr(err)
		if err != nil {
			w.ResultChannel <- WorkResult{Result: false, err: err}
			return
		}
		sensors := make([]mdl.Sensor, 0)
		var sid, stid int64
		var io_address, io_type, number_of_values, lenght_of_values, header_order, data_type, sensor_type int
		var conversion_expression, description string
		var soft_deleted bool

		for rows.Next() {
			err = rows.Scan(&sid, &stid, &io_address, &io_type, &number_of_values, &lenght_of_values, &header_order, &conversion_expression, &description, &data_type, &sensor_type, &soft_deleted)
			if err != nil {
				panic(err.Error())
			}
			s := mdl.NewSensor(
				sid,
				stid,
				io_address,
				io_type,
				sensor_type,
				number_of_values,
				lenght_of_values,
				header_order,
				data_type,
				description,
				conversion_expression,
				soft_deleted)
			sensors = append(sensors, s)
		}
		w.ResultChannel <- WorkResult{Result: sensors, err: err}
	}}

	var workResult = <-result
	if workResult.err != nil {
		log.Printf("A problem occured when getting the sensorheaders: %+v", workResult.err)
		return make([]mdl.Sensor, 0), workResult.err
	}
	sensors := workResult.Result.([]mdl.Sensor)
	return sensors, nil
}

func AddSensor(sensor mdl.Sensor) error {
	//does s.type exist?
	//no -> insert new type
	//inset s
	return nil
}

func getSensorTypeId(sensorType int) (int64, error) {
	result := make(chan WorkResult)
	defer close(result)
	args := make([]interface{}, 1)
	args[0] = sensorType
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: result, F: func(w *WorkRequest) {
		row := GetInstance().checkDevEUISTMT.QueryRow(w.Arguments...)
		var id int64
		err := row.Scan(&id)
		if err != nil {
			w.ResultChannel <- WorkResult{Result: 0, err: err}
			return
		}
		w.ResultChannel <- WorkResult{Result: id, err: err}
	}}
	var workResult = <-result
	checkErr(workResult.err)
	return workResult.Result.(int64), workResult.err
}

func ChangeSensorActivationState(sensors []mdl.Sensor) {
	for _, sensor := range sensors {
		ChangeSingleSensorActivationState(sensor)
	}
}

func UpdateHeader(devEUI string, newheader []mdl.Sensor) error {
	//check first
	if len(newheader) == 0 {
		return errors.New("No Sensors given")
	}
	if devEUI == "" {
		return errors.New("Deveui must not be empty")
	}
	if !CheckDevEUI(devEUI) {
		return errors.New("Deveui does not exist")
	}
	oldheader := GetNodeSensors(devEUI)
	newSensortyps := make(map[string]mdl.Sensor)
	if len(oldheader) == 0 {
		//old header does not exist, insert new ones

		for _, newS := range newheader {
			newSensortyps[fmt.Sprintf("%s %d", newS.Conversion_expression, newS.DataType)] = newS
		}
		for _, sensor := range newSensortyps {
			//Query: insert new sensortypes if not exist
			log.Printf("insert into sensor_typs where !%+v and !%+v on dubplicate ignore", sensor.Conversion_expression, sensor.DataType)
		}
		//Query: insert new sensors with Deveui
		return nil
	}
	sensortyps := make(map[string]bool)
	for _, s := range oldheader {
		sensortyps[fmt.Sprintf("%s %d", s.Conversion_expression, s.DataType)] = true
	}
	log.Printf("sensortyps map: %+v", sensortyps)
	for _, newS := range newheader {
		//check which newS is not in sensortyps; add those to newSensortyps
		newS.Conversion_expression = ""
	}
	//Query: softdelete where Deveui
	//Query: un-delete where Deveui and sensr iotype/ioaddress/sensorstype
	//Query: insert new sensortypes
	//Query: insert new sensors with Deveui

	return nil
}

func ChangeSingleSensorActivationState(sensor mdl.Sensor) {
	args := make([]interface{}, 2)
	log.Printf("deleted: %+v", sensor.Soft_deleted)
	args[0] = sensor.Soft_deleted
	args[1] = sensor.Id
	WorkQueue <- WorkRequest{Query: "", Arguments: args, ResultChannel: nil, F: func(w *WorkRequest) {
		_, err := GetInstance().changeSensorActivationStateSTMT.Exec(args...)
		if err != nil {
			log.Printf("Could not change the sensor state!\n %+v", err)
		}
	}}
}

func isOffline() bool {
	return instantiated == nil
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
