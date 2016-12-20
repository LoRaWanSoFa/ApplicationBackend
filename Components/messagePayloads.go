package components

type messagePayloadI interface {
	GetID() int64
	SetID(id int64)
	GetPayload() interface{}
	SetPayload(data interface{})
	GetSensor() Sensor
	SetSensor(s Sensor)
	Equals(mpi messagePayloadI) bool
}

type messagePayloadByte struct {
	ID      int64
	Payload []byte
	Sensor  Sensor
}

type messagePayloadString struct {
	ID      int64
	Payload string
	Sensor  Sensor
}

//messagePayloadByte

func (mpb *messagePayloadByte) GetID() int64 {
	return mpb.ID
}

func (mpb *messagePayloadByte) SetID(id int64) {
	mpb.ID = id
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
	if mpb.ID != mpi.GetID() {
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

func (mps *messagePayloadString) GetID() int64 {
	return mps.ID
}

func (mps *messagePayloadString) SetID(id int64) {
	mps.ID = id
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
	if mps.ID != mpi.GetID() {
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
