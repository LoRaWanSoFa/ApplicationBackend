package DatabaseConnector

import (
	errors "errors"
	"testing"

	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	if !isOffline() {
		t.Errorf("Already initialized before connecting")
	}
	errPing := Connect()
	if errPing != nil {
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
