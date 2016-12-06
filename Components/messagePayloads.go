package components

type messagePayloadI interface {
	GetId() int64
	SetId(id int64)
	GetPayload() interface{}
	SetPayload(data interface{})
	GetSensor() Sensor
	SetSensor(s Sensor)
}

type messagePayloadByte struct {
	Id      int64
	Payload []byte
	Sensor  Sensor
}

type messagePayloadString struct {
	Id      int64
	Payload string
	Sensor  Sensor
}

//messagePayloadByte

func (mpb *messagePayloadByte) GetId() int64 {
	return mpb.Id
}

func (mpb *messagePayloadByte) SetId(id int64) {
	mpb.Id = id
}

func (mpb *messagePayloadByte) GetPayload() interface{} {
	return mpb.Payload
}

func (mpb *messagePayloadByte) SetPayload(data interface{}) {
	converted, ok := data.([]byte)
	if ok {
		mpb.Payload = converted
	} else {
		mpb.Payload = nil
	}
}

func (mpb *messagePayloadByte) GetSensor() Sensor {
	return mpb.Sensor
}

func (mpb *messagePayloadByte) SetSensor(s Sensor) {
	mpb.Sensor = s
}

//messagePayloadString

func (mps *messagePayloadString) GetId() int64 {
	return mps.Id
}

func (mps *messagePayloadString) SetId(id int64) {
	mps.Id = id
}

func (mps *messagePayloadString) GetPayload() interface{} {
	return mps.Payload
}

func (mps *messagePayloadString) SetPayload(data interface{}) {
	converted, ok := data.(string)
	if ok {
		mps.Payload = converted
	} else {
		mps.Payload = ""
	}
}

func (mps *messagePayloadString) GetSensor() Sensor {
	return mps.Sensor
}

func (mps *messagePayloadString) SetSensor(s Sensor) {
	mps.Sensor = s
}
