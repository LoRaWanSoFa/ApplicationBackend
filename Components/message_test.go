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
func TestToJson(t *testing.T) {
	m := NewMessageUplink(23, "devEUI")
	firstSensor := NewHeaderSensor(77, 1, 2, 1, 1, "+1")
	secondSensor := NewHeaderSensor(78, 1, 3, 2, 1, "+2")
	m.AddPayloadString("test", firstSensor)
	m.AddPayloadString("test2", firstSensor)
	m.AddPayloadString("nose", secondSensor)
	m.AddPayloadString("cheese", secondSensor)

	json := m.ToJson()
	if len(json["77"]) != 2 {
		t.Errorf("Json should contain 2 entries for id 77, found %d", len(json))
	} else if json["77"][0] != "test" {
		t.Errorf("Json should contain \"test\" as first entry for id 77, found %s", json["77"][0])
	} else if json["77"][1] != "test2" {
		t.Errorf("Json should contain \"test2\" as second entry for id 77, found %s", json["77"][0])
	}
	if len(json["78"]) != 2 {
		t.Errorf("Json should contain 2 entries for id 77, found %d", len(json))
	} else if json["78"][0] != "nose" {
		t.Errorf("Json should contain \"nose\" as first entry for id 78, found %s", json["77"][0])
	} else if json["78"][1] != "cheese" {
		t.Errorf("Json should contain \"cheese\" as second entry for id 78, found %s", json["77"][0])
	}

}
