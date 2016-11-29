package components

type SensorI interface {
}

type Sensor struct {
	id                    int64
	ioType                int64 //NotNeeded
	ioAddress             int   //NotNeeded
	numberOfValues        int
	lenghtOfValues        int
	headerOrder           int
	description           string //NotNeeded
	conversion_expression string
	dataType              int // or int or string, or anything that shows what type of data this sensor is returning
}

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
	//return sensor(id:id, IoType:io_type)
	return Sensor{id, ioType, ioAddress, number_of_values, lenght_of_values, header_order, description, conversion_expression, data_type}
}
