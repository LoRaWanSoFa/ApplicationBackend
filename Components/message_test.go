package components

import "testing"

func TestNewMessageUplink(t *testing.T) {
	m := NewMessageUplink(23, "devEUI")
	if m.GetId() != 23 {
		t.Errorf("Expected %d, was %d", 23, m.GetId())
	}
	if len(m.GetPayloads()) != 0 {
		t.Errorf("Expected %d, was %d", 0, len(m.GetPayloads()))
	}
	switch v := m.(type) {
	case MessageUplinkI:
	default:
		t.Errorf("Type unknown: %v", v)
	}
	if m.GetDevEUI() != "devEUI" {
		t.Errorf("Expected %s, was %s", "devEUI", m.GetDevEUI())
	}
}

func TestAddPayload(t *testing.T) {
	m := NewMessageUplink(23, "devEUI")
	if len(m.GetPayloads()) != 0 {
		t.Errorf("Expected %d, was %d", 0, len(m.GetPayloads()))
	}
	s := NewHeaderSensor(77, 1, 2, 1, 1, "+1")
	m.AddPayload([]byte("test"), s)
	if len(m.GetPayloads()) != 1 {
		t.Errorf("Expected %d, was %d", 0, len(m.GetPayloads()))
	}
	payloads := m.GetPayloads()
	if string(payloads[0].GetPayload().([]byte)) != "test" {
		t.Errorf("Expected %+v, was %+v", "test", string(payloads[0].GetPayload().([]byte)))
	}
}

func TestAddPayloadString(t *testing.T) {
	m := NewMessageUplink(23, "devEUI")
	s := NewHeaderSensor(77, 1, 2, 1, 1, "+1")
	m.AddPayload([]byte("test"), s)
	m.RemovePayloads()
	m.AddPayloadString("TestString", s)
	payloads := m.GetPayloads()
	if len(payloads) != 1 {
		t.Errorf("Expected %+v, was lenght %+v", "lenght 1", len(payloads))
	}
	if payloads[0].GetPayload() != "TestString" {
		t.Errorf("Expected %+v, was %+v", "test", string(payloads[0].GetPayload().([]byte)))
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
