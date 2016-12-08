package components

import "testing"

func TestMessagPayloadByte(t *testing.T) {
	mp := new(messagePayloadByte)

	//Test if wrong datatypes get setted as the payload
	var x int
	x = 700
	mp.SetPayload(x)
	if p, ok := mp.GetPayload().([]byte); !ok || p != nil {
		t.Errorf("Payload was set while it was not a byte! payload: %+v", p)
	}

	//Test other geter and setters
	mp.SetPayload([]byte("testbyte"))
	mp.SetId(23)
	s := NewHeaderSensor(6, 1, 1, 1, 1, "+1")
	mp.SetSensor(s)
	if mp.GetId() != 23 {
		t.Errorf("Could not get correct Id: %+v", mp.GetId())
	}
	payload := mp.GetPayload().([]byte)
	if string(payload) != "testbyte" {
		t.Errorf("Could not get correct Payload: %+v", mp.GetPayload())
	}
	if mp.GetSensor() != s {
		t.Errorf("Could not get correct sensor: %+v", mp.GetSensor())
	}

}

func TestMessagPayloadString(t *testing.T) {
	mp := new(messagePayloadString)
	//Test if wrong datatypes get setted as the payload
	var x int
	x = 700
	mp.SetPayload(x)
	if mp.GetPayload() != "" {
		t.Errorf("Payload was set while it was not a string! payload: %+v", mp.GetPayload())
	}
	//Test other geter and setters
	mp.SetPayload("testbyte")
	mp.SetId(23)
	s := NewHeaderSensor(6, 1, 1, 1, 1, "+1")
	mp.SetSensor(s)
	if mp.GetId() != 23 {
		t.Errorf("Could not get correct Id: %+v", mp.GetId())
	}
	if mp.GetPayload() != "testbyte" {
		t.Errorf("Could not get correct Payload: %+v", mp.GetPayload())
	}
	if mp.GetSensor() != s {
		t.Errorf("Could not get correct sensor: %+v", mp.GetSensor())
	}
}
