package DatabaseConnector

import (
	"errors"
	"testing"

	mdl "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	if !isOffline() {
		t.Errorf("Already initialized before connecting")
	}
	errPing := Connect()
	if errPing != nil {
		// t.Log(DbConnectionInfo)
		t.Errorf("Could not do Connect()\n%+v", errPing)
	}
}

func TestCheckDevEUI(t *testing.T) {
	errPing := Connect()
	if errPing != nil {
		t.Errorf("Could not do Connect()\n%+v", errPing)
	}
	result := CheckDevEUI("not existend node _?!#@")
	if result == true {
		t.Errorf("Found a nonexisting node?!\n%+v", result)
	}
	result = CheckDevEUI("A4C12BF")
	if result != true {
		t.Errorf("Could not insert node: %+v", result)
	}

}

func TestAddMessage(t *testing.T) {
	Connect()
	message, err := AddMessage("A4C12BF")
	if err != nil {
		t.Errorf("Could not insert message: %+v", err)
	}
	if message == nil {
		t.Errorf("Could not insert message: %+v", message)
	}
	message, err = AddMessage("not existend node _?!#@")
	if err == nil {
		t.Errorf("Could insert message, which is not possible: %+v", err)
	}
	if message != nil {
		t.Errorf("Could insert message, which is not possible: %+v", message)
	}

}

func TestGetNodeSensors(t *testing.T) {
	sensors := GetNodeSensors("")
	if len(sensors) != 0 {
		t.Errorf("Node \"\" has sensors!?: %+v", sensors)
	}
	sensors = GetNodeSensors("A4C12BF")
	if len(sensors) <= 0 {
		t.Errorf("Node \"\" has no sensors?: %+v", sensors)
	}
}

func TestStoreMessagePayloads(t *testing.T) {
	var err error
	var m mdl.MessageUplinkI
	err = StoreMessagePayloads(nil)
	if err.Error() != "nil given as message parameter" {
		t.Errorf("Did not catch nil exception")
	}
	m = new(mdl.MessageUplink)
	err = StoreMessagePayloads(m)
	if err.Error() != "Message has not been initalized/stored" {
		t.Errorf("Works without an message id")
	}
	m = mdl.NewMessageUplink(33, "A4C12BF")
	err = StoreMessagePayloads(m)
	if err.Error() != "Nothing to store!" {
		t.Errorf("Works without payloads")
	}

	sensors := GetNodeSensors("A4C12BF")
	m, err = AddMessage("A4C12BF")
	if err != nil {
		t.Errorf("Could not add message for node A4C12BF. error: %+v", err)
	}
	m.RemovePayloads()
	m.AddPayloadString("Howdee1", sensors[0])
	m.AddPayloadString("Howdee2", sensors[1])
	err = StoreMessagePayloads(m)
	if err != nil {
		t.Errorf("Could not insert payloads for node A4C12BF. error: %+v", err)
	}
}

func TestStoreDownlinkMessage(t *testing.T) {
	dm := new(mdl.MessageDownLink)

	dm.Id = 1
	dmerror := StoreDownlinkMessage(dm)
	if dmerror.Error() != "Message already has an id, can not insert it" {
		t.Errorf("Allowed an already proccesed message to be stored again; %+v", dmerror)
	}
	dm.Id = 0
	dmerror = StoreDownlinkMessage(dm)
	if dmerror.Error() != "Message has an empty payload" {
		t.Errorf("Allowed a message to be stored without a payload; %+v", dmerror)
	}
	dm.Payload = "DOWNLINKMESSAGE"
	dmerror = StoreDownlinkMessage(dm)
	if dmerror.Error() != "Message has no DevEUI set" {
		t.Errorf("Allowed a message to be stored without a DevEUI; %+v", dmerror)
	}
	dm.Deveui = "A4C12BF"

	dmerror = StoreDownlinkMessage(dm)
	if dmerror != nil {
		t.Errorf("Could not store downlink message: %+v", dmerror)
	}
	if dm.Id == 0 {
		t.Error("Downlink message ID should be setted")
	}
	if dm.Time.IsZero() {
		t.Error("Time should be setted on the message")
	}
}

func TestUpdateHeader(t *testing.T) {
	newHeader := make([]mdl.Sensor, 0)
	deveui := ""
	var err error
	err = UpdateHeader(deveui, newHeader)
	if err == nil || err.Error() != "No Sensors given" {
		t.Errorf("Should not implicitly delete the current header")
	}
	newHeader = append(newHeader, mdl.NewSensor(0, 1, 4, 1, 0, 1, 1, 1, 1, "description", "conversion_expression", false))
	err = UpdateHeader(deveui, newHeader)
	if err == nil || err.Error() != "Deveui must not be empty" {
		t.Errorf("Deveui must not be empty")
	}
	deveui = "TEST_DEVEUI_THAT_DOES_NOT_EXIST"
	err = UpdateHeader(deveui, newHeader)
	if err == nil || err.Error() != "Deveui does not exist" {
		t.Errorf("Deveui does not exist")
	}
	deveui = "A4C12BF"
	err = UpdateHeader(deveui, newHeader)
	if err != nil {
		t.Errorf("did not expect an error: %+v", err)
	}

}

func TestGetFullHeader(t *testing.T) {
	sensors, err := GetFullHeader("")
	if len(sensors) != 0 {
		t.Errorf("Node \"\" has sensors!?: %+v", sensors)
	}
	if err != nil {
		t.Errorf("Encountered no error: %+v", err)
	}
	sensors, err = GetFullHeader("A4C12BF")
	if err != nil {
		t.Errorf("Encountered an error: %+v", err)
	}
	if len(sensors) <= 0 {
		t.Errorf("Node \"\" has no sensors?: %+v", sensors)
	}
}

func TestChangeSensorActivationState(t *testing.T) {
	sens, err := GetFullHeader("A4C12BF")
	for i := range sens {
		sens[i].Soft_deleted = false
	}
	ChangeSensorActivationState(sens)
	sensorStateHelperTest(t, sens[0].Id, false, err)

}

func sensorStateHelperTest(t *testing.T, id int64, expectedstate bool, err error) {
	if err != nil {
		t.Error("query failed")
	}
	var res bool
	err = GetInstance().Database.QueryRow("select soft_deleted from sensors where id=$1", id).Scan(&res)
	if res != expectedstate {
		t.Errorf("Softdelete with id: %+v should be %+v but was %+v.\n %+v", id, expectedstate, res, err)
	}

}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// The following is the code under test
	panicErr(errors.New("Some error"))
}

func TestClose(t *testing.T) {
	errPing := Connect()
	if errPing != nil {
		t.Errorf("Could not do Connect()\n%+v", errPing)
	}
	errClose := Close()
	if errClose != nil {
		t.Errorf("Could not do Close()\n%+v", errClose)
	}
}
