package mqttUplink

import (
	"encoding/binary"
	"errors"
	"fmt"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

type HeaderHandler interface {
	CreateNewHeader(payload []byte, devEUI string) ([]components.Sensor, error)
	StoreHeader(header []components.Sensor, devEUI string) error
}

type headerHandler struct {
}

func NewHeaderHandler() HeaderHandler {
	h := new(headerHandler)
	return h
}

func (h *headerHandler) CreateNewHeader(payload []byte, devEUI string) ([]components.Sensor, error) {
	var sensors []components.Sensor
	if h.checkLength(payload) {
		for i := 1; i < len(payload); i = i + 3 {
			sensor := h.createSensor(payload[i : i+3])
			sensor.HeaderOrder = (i + 2) / 3
			sensors = append(sensors, sensor)
		}
	} else {
		err := errors.New("Header of unkown length was send.")
		return nil, err
	}
	return sensors, nil
}

func (h *headerHandler) StoreHeader(sensor []components.Sensor, devEUI string) error {
	oldHeader := DatabaseConnector.GetNodeSensors(devEUI) //TODO: DatabaseConnector.GetFullHeader()
	for i := range oldHeader {
		fmt.Println(i)
	}
	return nil
}

func (h *headerHandler) createSensor(payload []byte) components.Sensor {
	var sensor components.Sensor

	ioType := payload[0]
	ioType = ioType >> 6
	sensor.IoType = int64(ioType)

	ioAddress := payload[0]
	ioAddress = ioAddress << 2 >> 4
	sensor.IoAddress = int(ioAddress)

	sensorType := payload[:2]
	sensorType[0] = sensorType[0] << 6 >> 6
	sensor.SensorType = int(binary.BigEndian.Uint16(sensorType))

	lenghtOfValues := payload[2]
	lenghtOfValues = lenghtOfValues >> 5
	if lenghtOfValues == 0 {
		lenghtOfValues = 1
	}
	sensor.LenghtOfValues = int(lenghtOfValues)

	numberOfValues := payload[2]
	numberOfValues = numberOfValues << 2 >> 5
	sensor.NumberOfValues = int(numberOfValues)

	sensor.Soft_deleted = false

	return sensor
}

func (hc *headerHandler) checkLength(payload []byte) bool {
	return len(payload)%3 == 1
}
