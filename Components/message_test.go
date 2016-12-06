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
