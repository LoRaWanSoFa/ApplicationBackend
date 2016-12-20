package components

//Sensor is a model that contains data from 2 tables in the database: sensors and sensortype.
type Sensor struct {
	ID                   int64  //sensor.id
	SensorTypeID         int64  //sensortype.id
	IoAddress            int    //
	IoType               int    //
	SensorType           int    //Not the ID but the actual sensortype number
	NumberOfValues       int    //lenght of bytes that belong to this sensor (between 1 and 8)
	LenghtOfValues       int    //number of values that belong to this sensor (1,2,4 or 8)
	HeaderOrder          int    //index in which the messages are send that belong to this sensor
	Description          string //NotNeeded - user input
	ConversionExpression string //math expression that transforms the message
	DataType             int    // or int or string, or anything that shows what type of data this sensor is returning
	SoftDeleted          bool   //true if the sensor is deleted
}

//NewHeaderSensor will return a sensor object with default values for fields that are not needed durring proccesing
func NewHeaderSensor(
	id int64,
	NumberOfValues,
	LenghtOfValues,
	HeaderOrder,
	DataType int,
	ConversionExpression string) Sensor {
	//return sensor(id:id, IoType:io_type)
	return Sensor{id, 0, 0, 0, 0, NumberOfValues, LenghtOfValues, HeaderOrder, "", ConversionExpression, DataType, false}
}

//NewSensor is the Default constructor
func NewSensor(
	id,
	sTypeID int64,
	ioAddress,
	iotype,
	SensorType,
	NumberOfValues,
	LenghtOfValues,
	HeaderOrder,
	DataType int,
	description,
	ConversionExpression string,
	SoftDeleted bool) Sensor {
	return Sensor{
		ID:                   id,
		SensorTypeID:         sTypeID,
		IoAddress:            ioAddress,
		IoType:               iotype,
		SensorType:           SensorType,
		NumberOfValues:       NumberOfValues,
		LenghtOfValues:       LenghtOfValues,
		HeaderOrder:          HeaderOrder,
		Description:          description,
		ConversionExpression: ConversionExpression,
		DataType:             DataType,
		SoftDeleted:          SoftDeleted}
}

//SameSensor compares this sensor to the given sensor.
func (s *Sensor) SameSensor(otherSensor Sensor) bool {
	b := s.SensorType == otherSensor.SensorType
	b = b && (s.IoAddress == otherSensor.IoAddress)
	b = b && (s.IoType == otherSensor.IoType)
	b = b && (s.LenghtOfValues == otherSensor.LenghtOfValues)
	b = b && (s.NumberOfValues == otherSensor.NumberOfValues)
	return b
}
