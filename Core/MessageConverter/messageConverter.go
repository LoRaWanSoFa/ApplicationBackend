package messageConverter

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

type conversionType int

// conversion types that can be received from messages
const (
	conversionUint   conversionType = iota // 0
	conversionInt                          // 1
	conversionFloat                        // 2
	conversionString                       // 3
	conversionBool                         // 4
)

type MessageConverter interface {
	ConvertSingleValue(payload []byte, conversion int) string
}

type messageConverter struct {
}

// Creates a new MessageConverter
func New() MessageConverter {
	return new(messageConverter)
}

// Converts a single value that was received from a message to a string
// representation of the received value. The conversion int refers to the type of
// the value that has been received. The types that can be input are uint(0),
// int(1), float(2), string(3) or bool(4).
func (m *messageConverter) ConvertSingleValue(payload []byte, conversion int) string {
	bits := binary.LittleEndian.Uint64(payload)
	float := math.Float64frombits(bits)
	return fmt.Sprintf("%f64 %d", float, conversionFloat)
}

// Conversion method for the uint type, which type of uint(8, 16, 32, 64) will
// be determined by the length of the byte slice. Returns a string representation
// of the resulting uint.
func convertUint(payload []byte) string {
	var result string
	var err error
	switch len(payload) {

	case 1: //uint8
		return string(payload[0])
	case 2: //uint16

	case 4: //uint32

	case 8: //uint64

	default:
		err = errors.New("illegal length of payload for a uint type")
	}
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return result
}

// Conversion method for the int type, which type of int(8, 16, 32, 64) will
// be determined by the length of the byte slice. Returns a string representation
// of the resulting int.
func convertInt(payload []byte) string {
	var result string
	var err error
	switch len(payload) {
	case 1: //int8

	case 2: //int16

	case 4: //int32

	case 8: //int64

	default:
		err = errors.New("illegal length of payload for an int type")
	}
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return result
}

// Conversion method for the float type, which type of float(32, 64) will
// be determined by the length of the byte slice. Returns a string representation
// of the resulting float.
func convertFloat(payload []byte) string {
	var float interface{}
	var err error
	buf := bytes.NewReader(payload)
	switch len(payload) {
	case 4: //float32
		float = float.(float32)
	case 8: //float64
		float = float.(float64)
	default:
		err = errors.New("illegal length of payload for a float type")
	}
	err = binary.Read(buf, binary.LittleEndian, &float)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return fmt.Sprintf("%f", float)
}

// Conversion method for the string type.
func convertString(payload []byte) string {
	return string(payload)
}

// Conversion method for the bool type.
// It is expected that the length of payload is never greater than  one.
// Result will be either "true" or "false" as a string.
func convertBool(payload []byte) string {
	if payload[0] > 0 {
		return "true"
	} else {
		return "false"
	}
}
