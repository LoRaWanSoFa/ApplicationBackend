package components

import "testing"

func TestSameSensor(t *testing.T) {
	sensor := NewSensor(0, 0, 0, 0, 1, 2, 1, 0, "", "", false)
	otherSensor := NewSensor(2, 0, 0, 0, 1, 2, 1, 0, "added a description", "+3", true)
	differentSensor := NewSensor(0, 0, 0, 0, 2, 2, 1, 0, "", "", false)
	if !sensor.SameSensor(otherSensor) {
		t.Errorf("The sensors should be the same sensor.\n First Sensor: %+v \n Second Sensor %+v", sensor, otherSensor)
	}
	if sensor.SameSensor(differentSensor) {
		t.Errorf("The sensors should be different.\n First Sensor: %+v \n Second Sensor %+v", sensor, differentSensor)
	}
}
