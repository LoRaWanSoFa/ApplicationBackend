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

func TestMessagePayloadEquals(t *testing.T) {
	var testData = []struct {
		firstInt  int // expectedResult
		secondInt int // input
	}{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 6},
		{0, 7},
		{3, 0},
		{3, 4},
		{3, 5},
		{3, 8},
	}

	m := NewMessageUplink(23, "devEUI")
	firstSensor := NewHeaderSensor(77, 1, 2, 1, 1, "+1")
	secondSensor := NewHeaderSensor(78, 1, 3, 2, 1, "+2")
	m.AddPayload([]byte("test"), firstSensor)
	m.AddPayload([]byte("test2"), firstSensor)
	m.AddPayload([]byte("test"), secondSensor)
	m.AddPayloadString("test", firstSensor)
	m.AddPayloadString("test2", firstSensor)
	m.AddPayloadString("test", secondSensor)
	m.AddPayload([]byte("nose"), firstSensor)
	m.AddPayload([]byte("test"), firstSensor)
	m.AddPayloadString("test", firstSensor)
	m.GetPayloads()[7].SetId(1)
	m.GetPayloads()[8].SetId(1)

	p := m.GetPayloads()
	for i := range testData {
		if p[testData[i].firstInt].Equals(p[testData[i].secondInt]) {
			t.Errorf("The first and second message payloads should be different. First message: %+v Second Message: %+v", p[testData[i].firstInt], p[testData[i].secondInt])
		}
	}
	if !p[1].Equals(p[1]) {
		t.Errorf("The first and second message payloads should be equal. First message: %+v Second Message: %+v", p[1], p[1])
	}
	if !p[4].Equals(p[4]) {
		t.Errorf("The first and second message payloads should be equal. First message: %+v Second Message: %+v", p[4], p[4])
	}
}
