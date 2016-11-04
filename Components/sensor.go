package components

type Sensor interface {
}

type sensor struct {
	dataType byte // or int or string, or anything that shows what type of data this sensor is returning
}
