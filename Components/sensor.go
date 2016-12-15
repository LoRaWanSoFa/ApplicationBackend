package components

type Sensor struct {
	Id                    int64  //sensor.id
	SensorTypeId          int64  //sensortype.id
	IoAddress             int    //
	IoType                int    //
	SensorType            int    //Not the ID but the actual sensortype number
	NumberOfValues        int    //lenght of bytes that belong to this sensor (between 1 and 8)
	LenghtOfValues        int    //number of values that belong to this sensor (1,2,4 or 8)
	HeaderOrder           int    //index in which the messages are send that belong to this sensor
	Description           string //NotNeeded - user input
	Conversion_expression string //math expression that transforms the message
	DataType              int    // or int or string, or anything that shows what type of data this sensor is returning
	Soft_deleted          bool   //true if the sensor is deleted
}

// Headersensor will return a sensor object with default values for fields that are not needed durring proccesing
func NewHeaderSensor(
	id int64,
	number_of_values,
	lenght_of_values,
	header_order,
	data_type int,
	conversion_expression string) Sensor {
	//return sensor(id:id, IoType:io_type)
	return Sensor{id, 0, 0, 0, 0, number_of_values, lenght_of_values, header_order, "", conversion_expression, data_type, false}
}

//Default constructor
func NewSensor(
	id,
	sTypeId int64,
	ioAddress,
	iotype,
	sensor_type,
	number_of_values,
	lenght_of_values,
	header_order,
	data_type int,
	description,
	conversion_expression string,
	soft_deleted bool) Sensor {
	return Sensor{
		Id:                    id,
		SensorTypeId:          sTypeId,
		IoAddress:             ioAddress,
		IoType:                iotype,
		SensorType:            sensor_type,
		NumberOfValues:        number_of_values,
		LenghtOfValues:        lenght_of_values,
		HeaderOrder:           header_order,
		Description:           description,
		Conversion_expression: conversion_expression,
		DataType:              data_type,
		Soft_deleted:          soft_deleted}
}

func (s *Sensor) SameSensor(otherSensor Sensor) bool {
	b := s.SensorType == otherSensor.SensorType
	b = b && (s.IoAddress == otherSensor.IoAddress)
	b = b && (s.IoType == otherSensor.IoType)
	b = b && (s.LenghtOfValues == otherSensor.LenghtOfValues)
	b = b && (s.NumberOfValues == otherSensor.NumberOfValues)
	return b
}
