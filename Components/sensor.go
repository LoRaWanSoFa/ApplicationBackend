package components

type Sensor struct {
	Id                    int64  //sensor.id
	IoType                int64  //NotNeeded
	IoAddress             int    //NotNeeded
	NumberOfValues        int    //lenght of bytes that belong to this sensor (between 1 and 8)
	LenghtOfValues        int    //number of values that belong to this sensor (1,2,4 or 8)
	HeaderOrder           int    //index in which the messages are send that belong to this sensor
	Description           string //NotNeeded - user input
	Conversion_expression string //math expression that transforms the message
	DataType              int    // or int or string, or anything that shows what type of data this sensor is returning
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
	return Sensor{id, 0, 0, number_of_values, lenght_of_values, header_order, "", conversion_expression, data_type}
}

//Default constructor
func NewSensor(
	id,
	ioType int64,
	ioAddress,
	number_of_values,
	lenght_of_values,
	header_order,
	data_type int,
	description,
	conversion_expression string) Sensor {
	return Sensor{id, ioType, ioAddress, number_of_values, lenght_of_values, header_order, description, conversion_expression, data_type}
}
