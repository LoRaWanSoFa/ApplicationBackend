// Code generated by protoc-gen-gogo.
// source: core.proto
// DO NOT EDIT!

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Metadata struct {
	Region      string  `protobuf:"bytes,99,opt,name=Region,proto3" json:"Region,omitempty"`
	DutyRX1     uint32  `protobuf:"varint,1,opt,name=DutyRX1,proto3" json:"DutyRX1,omitempty"`
	DutyRX2     uint32  `protobuf:"varint,2,opt,name=DutyRX2,proto3" json:"DutyRX2,omitempty"`
	Frequency   float32 `protobuf:"fixed32,3,opt,name=Frequency,proto3" json:"Frequency,omitempty"`
	DataRate    string  `protobuf:"bytes,4,opt,name=DataRate,proto3" json:"DataRate,omitempty"`
	CodingRate  string  `protobuf:"bytes,5,opt,name=CodingRate,proto3" json:"CodingRate,omitempty"`
	Timestamp   uint32  `protobuf:"varint,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Rssi        int32   `protobuf:"varint,7,opt,name=Rssi,proto3" json:"Rssi,omitempty"`
	Lsnr        float32 `protobuf:"fixed32,8,opt,name=Lsnr,proto3" json:"Lsnr,omitempty"`
	PayloadSize uint32  `protobuf:"varint,9,opt,name=PayloadSize,proto3" json:"PayloadSize,omitempty"`
	Time        string  `protobuf:"bytes,10,opt,name=Time,proto3" json:"Time,omitempty"`
	RFChain     uint32  `protobuf:"varint,11,opt,name=RFChain,proto3" json:"RFChain,omitempty"`
	CRCStatus   int32   `protobuf:"varint,12,opt,name=CRCStatus,proto3" json:"CRCStatus,omitempty"`
	Modulation  string  `protobuf:"bytes,13,opt,name=Modulation,proto3" json:"Modulation,omitempty"`
	InvPolarity bool    `protobuf:"varint,14,opt,name=InvPolarity,proto3" json:"InvPolarity,omitempty"`
	Power       uint32  `protobuf:"varint,15,opt,name=Power,proto3" json:"Power,omitempty"`
	Channel     uint32  `protobuf:"varint,16,opt,name=Channel,proto3" json:"Channel,omitempty"`
	GatewayEUI  string  `protobuf:"bytes,20,opt,name=GatewayEUI,proto3" json:"GatewayEUI,omitempty"`
	Altitude    int32   `protobuf:"varint,21,opt,name=Altitude,proto3" json:"Altitude,omitempty"`
	Longitude   float32 `protobuf:"fixed32,22,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Latitude    float32 `protobuf:"fixed32,23,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	ServerTime  string  `protobuf:"bytes,31,opt,name=ServerTime,proto3" json:"ServerTime,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorCore, []int{0} }

type StatsMetadata struct {
	Altitude  int32   `protobuf:"varint,1,opt,name=Altitude,proto3" json:"Altitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Latitude  float32 `protobuf:"fixed32,3,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
}

func (m *StatsMetadata) Reset()                    { *m = StatsMetadata{} }
func (m *StatsMetadata) String() string            { return proto.CompactTextString(m) }
func (*StatsMetadata) ProtoMessage()               {}
func (*StatsMetadata) Descriptor() ([]byte, []int) { return fileDescriptorCore, []int{1} }

func init() {
	proto.RegisterType((*Metadata)(nil), "core.Metadata")
	proto.RegisterType((*StatsMetadata)(nil), "core.StatsMetadata")
}
func (m *Metadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Metadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DutyRX1 != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCore(data, i, uint64(m.DutyRX1))
	}
	if m.DutyRX2 != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCore(data, i, uint64(m.DutyRX2))
	}
	if m.Frequency != 0 {
		data[i] = 0x1d
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Frequency))))
	}
	if len(m.DataRate) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintCore(data, i, uint64(len(m.DataRate)))
		i += copy(data[i:], m.DataRate)
	}
	if len(m.CodingRate) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintCore(data, i, uint64(len(m.CodingRate)))
		i += copy(data[i:], m.CodingRate)
	}
	if m.Timestamp != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintCore(data, i, uint64(m.Timestamp))
	}
	if m.Rssi != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintCore(data, i, uint64(m.Rssi))
	}
	if m.Lsnr != 0 {
		data[i] = 0x45
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Lsnr))))
	}
	if m.PayloadSize != 0 {
		data[i] = 0x48
		i++
		i = encodeVarintCore(data, i, uint64(m.PayloadSize))
	}
	if len(m.Time) > 0 {
		data[i] = 0x52
		i++
		i = encodeVarintCore(data, i, uint64(len(m.Time)))
		i += copy(data[i:], m.Time)
	}
	if m.RFChain != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintCore(data, i, uint64(m.RFChain))
	}
	if m.CRCStatus != 0 {
		data[i] = 0x60
		i++
		i = encodeVarintCore(data, i, uint64(m.CRCStatus))
	}
	if len(m.Modulation) > 0 {
		data[i] = 0x6a
		i++
		i = encodeVarintCore(data, i, uint64(len(m.Modulation)))
		i += copy(data[i:], m.Modulation)
	}
	if m.InvPolarity {
		data[i] = 0x70
		i++
		if m.InvPolarity {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Power != 0 {
		data[i] = 0x78
		i++
		i = encodeVarintCore(data, i, uint64(m.Power))
	}
	if m.Channel != 0 {
		data[i] = 0x80
		i++
		data[i] = 0x1
		i++
		i = encodeVarintCore(data, i, uint64(m.Channel))
	}
	if len(m.GatewayEUI) > 0 {
		data[i] = 0xa2
		i++
		data[i] = 0x1
		i++
		i = encodeVarintCore(data, i, uint64(len(m.GatewayEUI)))
		i += copy(data[i:], m.GatewayEUI)
	}
	if m.Altitude != 0 {
		data[i] = 0xa8
		i++
		data[i] = 0x1
		i++
		i = encodeVarintCore(data, i, uint64(m.Altitude))
	}
	if m.Longitude != 0 {
		data[i] = 0xb5
		i++
		data[i] = 0x1
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Longitude))))
	}
	if m.Latitude != 0 {
		data[i] = 0xbd
		i++
		data[i] = 0x1
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Latitude))))
	}
	if len(m.ServerTime) > 0 {
		data[i] = 0xfa
		i++
		data[i] = 0x1
		i++
		i = encodeVarintCore(data, i, uint64(len(m.ServerTime)))
		i += copy(data[i:], m.ServerTime)
	}
	if len(m.Region) > 0 {
		data[i] = 0x9a
		i++
		data[i] = 0x6
		i++
		i = encodeVarintCore(data, i, uint64(len(m.Region)))
		i += copy(data[i:], m.Region)
	}
	return i, nil
}

func (m *StatsMetadata) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatsMetadata) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Altitude != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCore(data, i, uint64(m.Altitude))
	}
	if m.Longitude != 0 {
		data[i] = 0x15
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Longitude))))
	}
	if m.Latitude != 0 {
		data[i] = 0x1d
		i++
		i = encodeFixed32Core(data, i, uint32(math.Float32bits(float32(m.Latitude))))
	}
	return i, nil
}

func encodeFixed64Core(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Core(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCore(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Metadata) Size() (n int) {
	var l int
	_ = l
	if m.DutyRX1 != 0 {
		n += 1 + sovCore(uint64(m.DutyRX1))
	}
	if m.DutyRX2 != 0 {
		n += 1 + sovCore(uint64(m.DutyRX2))
	}
	if m.Frequency != 0 {
		n += 5
	}
	l = len(m.DataRate)
	if l > 0 {
		n += 1 + l + sovCore(uint64(l))
	}
	l = len(m.CodingRate)
	if l > 0 {
		n += 1 + l + sovCore(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovCore(uint64(m.Timestamp))
	}
	if m.Rssi != 0 {
		n += 1 + sovCore(uint64(m.Rssi))
	}
	if m.Lsnr != 0 {
		n += 5
	}
	if m.PayloadSize != 0 {
		n += 1 + sovCore(uint64(m.PayloadSize))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovCore(uint64(l))
	}
	if m.RFChain != 0 {
		n += 1 + sovCore(uint64(m.RFChain))
	}
	if m.CRCStatus != 0 {
		n += 1 + sovCore(uint64(m.CRCStatus))
	}
	l = len(m.Modulation)
	if l > 0 {
		n += 1 + l + sovCore(uint64(l))
	}
	if m.InvPolarity {
		n += 2
	}
	if m.Power != 0 {
		n += 1 + sovCore(uint64(m.Power))
	}
	if m.Channel != 0 {
		n += 2 + sovCore(uint64(m.Channel))
	}
	l = len(m.GatewayEUI)
	if l > 0 {
		n += 2 + l + sovCore(uint64(l))
	}
	if m.Altitude != 0 {
		n += 2 + sovCore(uint64(m.Altitude))
	}
	if m.Longitude != 0 {
		n += 6
	}
	if m.Latitude != 0 {
		n += 6
	}
	l = len(m.ServerTime)
	if l > 0 {
		n += 2 + l + sovCore(uint64(l))
	}
	l = len(m.Region)
	if l > 0 {
		n += 2 + l + sovCore(uint64(l))
	}
	return n
}

func (m *StatsMetadata) Size() (n int) {
	var l int
	_ = l
	if m.Altitude != 0 {
		n += 1 + sovCore(uint64(m.Altitude))
	}
	if m.Longitude != 0 {
		n += 5
	}
	if m.Latitude != 0 {
		n += 5
	}
	return n
}

func sovCore(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCore(x uint64) (n int) {
	return sovCore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutyRX1", wireType)
			}
			m.DutyRX1 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DutyRX1 |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DutyRX2", wireType)
			}
			m.DutyRX2 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DutyRX2 |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frequency", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Frequency = float32(math.Float32frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodingRate = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Timestamp |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rssi", wireType)
			}
			m.Rssi = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Rssi |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lsnr", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Lsnr = float32(math.Float32frombits(v))
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadSize", wireType)
			}
			m.PayloadSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.PayloadSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RFChain", wireType)
			}
			m.RFChain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RFChain |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CRCStatus", wireType)
			}
			m.CRCStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.CRCStatus |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modulation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Modulation = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvPolarity", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InvPolarity = bool(v != 0)
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Power |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			m.Channel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Channel |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayEUI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayEUI = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Altitude", wireType)
			}
			m.Altitude = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Altitude |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Longitude = float32(math.Float32frombits(v))
		case 23:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Latitude = float32(math.Float32frombits(v))
		case 31:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerTime = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 99:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Region", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCore
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Region = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatsMetadata) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatsMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatsMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Altitude", wireType)
			}
			m.Altitude = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Altitude |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Longitude = float32(math.Float32frombits(v))
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Latitude = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipCore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCore
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCore(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCore
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCore
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCore
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCore(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCore = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCore   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("core.proto", fileDescriptorCore) }

var fileDescriptorCore = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x65, 0xd3, 0x24, 0x4d, 0xb6, 0x04, 0xaa, 0x55, 0x29, 0x23, 0x84, 0x8c, 0xd5, 0x93, 0x4f,
	0x48, 0x94, 0x2f, 0x00, 0x97, 0xa2, 0x4a, 0xa9, 0x14, 0x6d, 0x40, 0xe2, 0x3a, 0xc4, 0xab, 0xd4,
	0x92, 0xbb, 0x5b, 0xec, 0x75, 0x2b, 0xf3, 0x25, 0x7c, 0x12, 0x47, 0x3e, 0x01, 0x85, 0x2b, 0x1f,
	0x81, 0x66, 0xd6, 0xd8, 0xee, 0xa5, 0xb7, 0x79, 0xef, 0x69, 0xe6, 0xbd, 0x99, 0x5d, 0x29, 0x37,
	0xae, 0x34, 0xaf, 0x6f, 0x4a, 0xe7, 0x9d, 0x1a, 0x53, 0x7d, 0xf2, 0x77, 0x2c, 0x67, 0x97, 0xc6,
	0x63, 0x86, 0x1e, 0x15, 0xc8, 0xfd, 0xb3, 0xda, 0x37, 0xfa, 0xcb, 0x1b, 0x10, 0xb1, 0x48, 0x16,
	0xfa, 0x3f, 0xec, 0x95, 0x53, 0x18, 0x0d, 0x95, 0x53, 0xf5, 0x52, 0xce, 0xcf, 0x4b, 0xf3, 0xad,
	0x36, 0x76, 0xd3, 0xc0, 0x5e, 0x2c, 0x92, 0x91, 0xee, 0x09, 0xf5, 0x42, 0xce, 0xce, 0xd0, 0xa3,
	0x46, 0x6f, 0x60, 0x1c, 0x8b, 0x64, 0xae, 0x3b, 0xac, 0x22, 0x29, 0x53, 0x97, 0xe5, 0x76, 0xcb,
	0xea, 0x84, 0xd5, 0x01, 0x43, 0x93, 0x3f, 0xe5, 0xd7, 0xa6, 0xf2, 0x78, 0x7d, 0x03, 0x53, 0x76,
	0xed, 0x09, 0xa5, 0xe4, 0x58, 0x57, 0x55, 0x0e, 0xfb, 0xb1, 0x48, 0x26, 0x9a, 0x6b, 0xe2, 0x96,
	0x95, 0x2d, 0x61, 0xc6, 0x31, 0xb8, 0x56, 0xb1, 0x3c, 0x58, 0x61, 0x53, 0x38, 0xcc, 0xd6, 0xf9,
	0x77, 0x03, 0x73, 0x9e, 0x33, 0xa4, 0xa8, 0x8b, 0xc6, 0x82, 0xe4, 0x04, 0x5c, 0xd3, 0xbe, 0xfa,
	0x3c, 0xbd, 0xc2, 0xdc, 0xc2, 0x41, 0xd8, 0xb7, 0x85, 0x94, 0x2a, 0xd5, 0xe9, 0xda, 0xa3, 0xaf,
	0x2b, 0x78, 0xcc, 0xe6, 0x3d, 0x41, 0x3b, 0x5d, 0xba, 0xac, 0x2e, 0xd0, 0xe7, 0xce, 0xc2, 0x22,
	0xec, 0xd4, 0x33, 0x94, 0xe6, 0xc2, 0xde, 0xae, 0x5c, 0x81, 0x65, 0xee, 0x1b, 0x78, 0x12, 0x8b,
	0x64, 0xa6, 0x87, 0x94, 0x3a, 0x92, 0x93, 0x95, 0xbb, 0x33, 0x25, 0x3c, 0x65, 0xdf, 0x00, 0x28,
	0x4f, 0x7a, 0x85, 0xd6, 0x9a, 0x02, 0x0e, 0x43, 0x9e, 0x16, 0x92, 0xe3, 0x47, 0xf4, 0xe6, 0x0e,
	0x9b, 0x0f, 0x9f, 0x2f, 0xe0, 0x28, 0x38, 0xf6, 0x0c, 0xbd, 0xc0, 0xbb, 0xc2, 0xe7, 0xbe, 0xce,
	0x0c, 0x3c, 0xe3, 0xb8, 0x1d, 0xa6, 0x5d, 0x96, 0xce, 0x6e, 0x83, 0x78, 0x1c, 0xde, 0xae, 0x23,
	0xa8, 0x73, 0x89, 0x6d, 0xe7, 0x73, 0x16, 0x3b, 0x4c, 0xae, 0x6b, 0x53, 0xde, 0x9a, 0x92, 0x2f,
	0xf7, 0x2a, 0xb8, 0xf6, 0x8c, 0x3a, 0x96, 0x53, 0x6d, 0xb6, 0x74, 0x83, 0x0d, 0x6b, 0x2d, 0x3a,
	0x31, 0x72, 0x41, 0x97, 0xaa, 0xba, 0x2f, 0x37, 0x8c, 0x27, 0x1e, 0x8a, 0x37, 0x7a, 0x28, 0xde,
	0xde, 0xfd, 0x78, 0xef, 0x0f, 0x7f, 0xee, 0x22, 0xf1, 0x6b, 0x17, 0x89, 0xdf, 0xbb, 0x48, 0xfc,
	0xf8, 0x13, 0x3d, 0xfa, 0x3a, 0xe5, 0x4f, 0xff, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0xd5, 0x19, 0x61, 0x02, 0x03, 0x00, 0x00,
}