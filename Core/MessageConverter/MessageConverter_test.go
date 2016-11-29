package MessageConverter

import "testing"

var fibTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

var invalidByteArray = []byte{0x00, 0x00, 0x00}
var validByteArrays = [][]byte{
	{0xFF},
	{0x00},
	{0x99},
	{0x99, 0x23},
	{0xA2, 0xCB},
	{0xA2, 0xCB, 0xF2, 0xC3},
	{0x11, 0x11, 0x11, 0x11},
	{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0},
	{0x3B, 0x17, 0x24, 0x92, 0xA8, 0xF2, 0x9D, 0x2E}}

func TestUintTypes(t *testing.T) {
	var testData = []struct {
		expectedResult string // expectedResult
		payload        []byte // input
	}{
		{"255", validByteArrays[0]},
		{"0", validByteArrays[1]},
		{"153", validByteArrays[2]},
		{"39203", validByteArrays[3]},
		{"41675", validByteArrays[4]},
		{"2731274947", validByteArrays[5]},
		{"286331153", validByteArrays[6]},
		{"1311768467463790320", validByteArrays[7]},
		{"4257912185020390702", validByteArrays[8]},
	}
	messageConverter := New()
	for _, item := range testData {
		result, _ := messageConverter.ConvertSingleValue(item.payload, 0)
		if result != item.expectedResult {
			t.Errorf("Expected %s, was %s with payload %v", item.expectedResult, result, item.payload)
		}
	}
	_, err := messageConverter.ConvertSingleValue(invalidByteArray, 0)
	if err == nil {
		t.Errorf("Uint message length should be invalid and produce an error %v", invalidByteArray)
	}
}

func TestIntTypes(t *testing.T) {
	var testData = []struct {
		expectedResult string // expectedResult
		payload        []byte // input
	}{
		{"-1", validByteArrays[0]},
		{"0", validByteArrays[1]},
		{"-103", validByteArrays[2]},
		{"-26333", validByteArrays[3]},
		{"-23861", validByteArrays[4]},
		{"-1563692349", validByteArrays[5]},
		{"286331153", validByteArrays[6]},
		{"1311768467463790320", validByteArrays[7]},
		{"4257912185020390702", validByteArrays[8]},
	}
	messageConverter := New()
	for _, item := range testData {
		result, _ := messageConverter.ConvertSingleValue(item.payload, 1)
		if result != item.expectedResult {
			t.Errorf("Expected %s, was %s with payload %v", item.expectedResult, result, item.payload)
		}
	}
	_, err := messageConverter.ConvertSingleValue(invalidByteArray, 1)
	if err == nil {
		t.Errorf("Int message length should be invalid and produce an error %v", invalidByteArray)
	}
}
func TestFloatTypes(t *testing.T) {
	var testData = []struct {
		expectedResult string // expectedResult
		payload        []byte // input
	}{
		{"-5.528029E-18", validByteArrays[5]},
		{"1.144374E-28", validByteArrays[6]},
		{"5.626349E-221", validByteArrays[7]},
		{"4.785832E-24", validByteArrays[8]},
	}
	messageConverter := New()
	for _, item := range testData {
		result, _ := messageConverter.ConvertSingleValue(item.payload, 2)
		if result != item.expectedResult {
			t.Errorf("Expected %s, was %s with payload %v", item.expectedResult, result, item.payload)
		}
	}
	_, err := messageConverter.ConvertSingleValue(invalidByteArray, 2)
	if err == nil {
		t.Errorf("Float message length should be invalid and produce an error %v", invalidByteArray)
	}
}

func TestStringConversion(t *testing.T) {
	var testData = []struct {
		expectedResult string // expectedResult
		payload        []byte // input
	}{
		{"Test", []byte{0x54, 0x65, 0x73, 0x74}},
		{"Bits", []byte{0x42, 0x69, 0x74, 0x73}},
		{"Bytes", []byte{0x42, 0x79, 0x74, 0x65, 0x73}},
		{"supercalifragilisticexpialidocious", []byte{0x73, 0x75, 0x70, 0x65, 0x72,
			0x63, 0x61, 0x6C, 0x69, 0x66, 0x72, 0x61, 0x67, 0x69, 0x6C, 0x69, 0x73,
			0x74, 0x69, 0x63, 0x65, 0x78, 0x70, 0x69, 0x61, 0x6C, 0x69, 0x64, 0x6F,
			0x63, 0x69, 0x6F, 0x75, 0x73}},
	}
	messageConverter := New()
	for _, item := range testData {
		result, _ := messageConverter.ConvertSingleValue(item.payload, 3)
		if result != item.expectedResult {
			t.Errorf("Expected %s, was %s with payload %v", item.expectedResult, result, item.payload)
		}
	}
}

func TestBooleanTypes(t *testing.T) {
	var testData = []struct {
		expectedResult string // expectedResult
		payload        []byte // input
	}{
		{"true", []byte{0xFF}},
		{"true", []byte{0x01}},
		{"false", []byte{0x00}},
	}
	messageConverter := New()
	for _, item := range testData {
		result, _ := messageConverter.ConvertSingleValue(item.payload, 5)
		if result != item.expectedResult {
			t.Errorf("Expected %s, was %s with payload %v", item.expectedResult, result, item.payload)
		}
	}
	_, err := messageConverter.ConvertSingleValue(invalidByteArray, 5)
	if err == nil {
		t.Errorf("Boolean message length should be invalid and produce an error %v", invalidByteArray)
	}
}

func TestHexStringConversion(t *testing.T) {
	var testData = []struct {
		expectedResult string // expectedResult
		payload        []byte // input
	}{
		{"FF", validByteArrays[0]},
		{"00", validByteArrays[1]},
		{"99", validByteArrays[2]},
		{"9923", validByteArrays[3]},
		{"A2CB", validByteArrays[4]},
		{"A2CBF2C3", validByteArrays[5]},
		{"11111111", validByteArrays[6]},
		{"123456789ABCDEF0", validByteArrays[7]},
		{"3B172492A8F29D2E", validByteArrays[8]},
		{"3B1724", []byte{0x3B, 0x17, 0x24}}, //testing a three length slice
	}
	messageConverter := New()
	for _, item := range testData {
		result, _ := messageConverter.ConvertSingleValue(item.payload, 4)
		if result != item.expectedResult {
			t.Errorf("Expected %s, was %s with payload %v", item.expectedResult, result, item.payload)
		}
	}
}

func TestInvalidType(t *testing.T) {
	messageConverter := New()
	_, err := messageConverter.ConvertSingleValue(invalidByteArray, 99)
	if err == nil {
		t.Errorf("Boolean message length should be invalid and produce an error %v", invalidByteArray)
	}
}
