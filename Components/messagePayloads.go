package components

type messagePayloadI interface {
	GetId() int64
	SetId(id int64)
	GetPayload() interface{}
	SetPayload(data interface{})
	GetSensor() Sensor
	SetSensor(s Sensor)
	Equals(mpi messagePayloadI) bool
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

func (mpb *messagePayloadByte) Equals(mpi messagePayloadI) bool {
	if mpb.Id != mpi.GetId() {
		return false
	}
	otherPayload, ok := mpi.GetPayload().([]byte)
	if ok {
		if len(mpb.Payload) == len(otherPayload) {
			for i := range mpb.Payload {
				if mpb.Payload[i] != otherPayload[i] {
					return false
				}
			}
		} else {
			return false
		}
	} else {
		return false
	}
	if mpb.Sensor != mpi.GetSensor() {
		return false
	}

	return true
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

func (mps *messagePayloadString) Equals(mpi messagePayloadI) bool {
	if mps.Id != mpi.GetId() {
		return false
	}
	otherPayload, ok := mpi.GetPayload().(string)
	if ok {
		if mps.Payload != otherPayload {
			return false
		}
	} else {
		return false
	}
	if mps.Sensor != mpi.GetSensor() {
		return false
	}
	return true
}
