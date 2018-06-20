// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proio/model/mc.proto

/*
	Package mc is a generated protocol buffer package.

	It is generated from these files:
		proio/model/mc.proto

	It has these top-level messages:
		MCParameters
		PythiaParameters
*/
package mc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MCParameters struct {
	Number           *uint64  `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Processid        *int32   `protobuf:"varint,2,opt,name=processid" json:"processid,omitempty"`
	PtHat            *float64 `protobuf:"fixed64,3,opt,name=pt_hat,json=ptHat" json:"pt_hat,omitempty"`
	AlphaEm          *float64 `protobuf:"fixed64,4,opt,name=alpha_em,json=alphaEm" json:"alpha_em,omitempty"`
	AlphaS           *float64 `protobuf:"fixed64,5,opt,name=alpha_s,json=alphaS" json:"alpha_s,omitempty"`
	ScaleQFac        *float64 `protobuf:"fixed64,6,opt,name=scale_q_fac,json=scaleQFac" json:"scale_q_fac,omitempty"`
	Weight           *float64 `protobuf:"fixed64,7,opt,name=weight" json:"weight,omitempty"`
	X1               *float64 `protobuf:"fixed64,8,opt,name=x1" json:"x1,omitempty"`
	X2               *float64 `protobuf:"fixed64,9,opt,name=x2" json:"x2,omitempty"`
	Id1              *uint64  `protobuf:"varint,10,opt,name=id1" json:"id1,omitempty"`
	Id2              *uint64  `protobuf:"varint,11,opt,name=id2" json:"id2,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MCParameters) Reset()                    { *m = MCParameters{} }
func (m *MCParameters) String() string            { return proto.CompactTextString(m) }
func (*MCParameters) ProtoMessage()               {}
func (*MCParameters) Descriptor() ([]byte, []int) { return fileDescriptorMc, []int{0} }

func (m *MCParameters) GetNumber() uint64 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *MCParameters) GetProcessid() int32 {
	if m != nil && m.Processid != nil {
		return *m.Processid
	}
	return 0
}

func (m *MCParameters) GetPtHat() float64 {
	if m != nil && m.PtHat != nil {
		return *m.PtHat
	}
	return 0
}

func (m *MCParameters) GetAlphaEm() float64 {
	if m != nil && m.AlphaEm != nil {
		return *m.AlphaEm
	}
	return 0
}

func (m *MCParameters) GetAlphaS() float64 {
	if m != nil && m.AlphaS != nil {
		return *m.AlphaS
	}
	return 0
}

func (m *MCParameters) GetScaleQFac() float64 {
	if m != nil && m.ScaleQFac != nil {
		return *m.ScaleQFac
	}
	return 0
}

func (m *MCParameters) GetWeight() float64 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return 0
}

func (m *MCParameters) GetX1() float64 {
	if m != nil && m.X1 != nil {
		return *m.X1
	}
	return 0
}

func (m *MCParameters) GetX2() float64 {
	if m != nil && m.X2 != nil {
		return *m.X2
	}
	return 0
}

func (m *MCParameters) GetId1() uint64 {
	if m != nil && m.Id1 != nil {
		return *m.Id1
	}
	return 0
}

func (m *MCParameters) GetId2() uint64 {
	if m != nil && m.Id2 != nil {
		return *m.Id2
	}
	return 0
}

type PythiaParameters struct {
	WeightSum        *float64 `protobuf:"fixed64,1,opt,name=weight_sum,json=weightSum" json:"weight_sum,omitempty"`
	MergingWeight    *float64 `protobuf:"fixed64,2,opt,name=merging_weight,json=mergingWeight" json:"merging_weight,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PythiaParameters) Reset()                    { *m = PythiaParameters{} }
func (m *PythiaParameters) String() string            { return proto.CompactTextString(m) }
func (*PythiaParameters) ProtoMessage()               {}
func (*PythiaParameters) Descriptor() ([]byte, []int) { return fileDescriptorMc, []int{1} }

func (m *PythiaParameters) GetWeightSum() float64 {
	if m != nil && m.WeightSum != nil {
		return *m.WeightSum
	}
	return 0
}

func (m *PythiaParameters) GetMergingWeight() float64 {
	if m != nil && m.MergingWeight != nil {
		return *m.MergingWeight
	}
	return 0
}

func init() {
	proto.RegisterType((*MCParameters)(nil), "proio.model.mc.MCParameters")
	proto.RegisterType((*PythiaParameters)(nil), "proio.model.mc.PythiaParameters")
}
func (m *MCParameters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MCParameters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Number != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMc(dAtA, i, uint64(*m.Number))
	}
	if m.Processid != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMc(dAtA, i, uint64(*m.Processid))
	}
	if m.PtHat != nil {
		dAtA[i] = 0x19
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.PtHat))))
		i += 8
	}
	if m.AlphaEm != nil {
		dAtA[i] = 0x21
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.AlphaEm))))
		i += 8
	}
	if m.AlphaS != nil {
		dAtA[i] = 0x29
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.AlphaS))))
		i += 8
	}
	if m.ScaleQFac != nil {
		dAtA[i] = 0x31
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.ScaleQFac))))
		i += 8
	}
	if m.Weight != nil {
		dAtA[i] = 0x39
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.Weight))))
		i += 8
	}
	if m.X1 != nil {
		dAtA[i] = 0x41
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.X1))))
		i += 8
	}
	if m.X2 != nil {
		dAtA[i] = 0x49
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.X2))))
		i += 8
	}
	if m.Id1 != nil {
		dAtA[i] = 0x50
		i++
		i = encodeVarintMc(dAtA, i, uint64(*m.Id1))
	}
	if m.Id2 != nil {
		dAtA[i] = 0x58
		i++
		i = encodeVarintMc(dAtA, i, uint64(*m.Id2))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PythiaParameters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PythiaParameters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WeightSum != nil {
		dAtA[i] = 0x9
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.WeightSum))))
		i += 8
	}
	if m.MergingWeight != nil {
		dAtA[i] = 0x11
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.MergingWeight))))
		i += 8
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MCParameters) Size() (n int) {
	var l int
	_ = l
	if m.Number != nil {
		n += 1 + sovMc(uint64(*m.Number))
	}
	if m.Processid != nil {
		n += 1 + sovMc(uint64(*m.Processid))
	}
	if m.PtHat != nil {
		n += 9
	}
	if m.AlphaEm != nil {
		n += 9
	}
	if m.AlphaS != nil {
		n += 9
	}
	if m.ScaleQFac != nil {
		n += 9
	}
	if m.Weight != nil {
		n += 9
	}
	if m.X1 != nil {
		n += 9
	}
	if m.X2 != nil {
		n += 9
	}
	if m.Id1 != nil {
		n += 1 + sovMc(uint64(*m.Id1))
	}
	if m.Id2 != nil {
		n += 1 + sovMc(uint64(*m.Id2))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PythiaParameters) Size() (n int) {
	var l int
	_ = l
	if m.WeightSum != nil {
		n += 9
	}
	if m.MergingWeight != nil {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMc(x uint64) (n int) {
	return sovMc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MCParameters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MCParameters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MCParameters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Number = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Processid", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Processid = &v
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field PtHat", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.PtHat = &v2
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlphaEm", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.AlphaEm = &v2
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlphaS", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.AlphaS = &v2
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScaleQFac", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.ScaleQFac = &v2
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.Weight = &v2
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field X1", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.X1 = &v2
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field X2", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.X2 = &v2
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id1", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Id1 = &v
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id2", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Id2 = &v
		default:
			iNdEx = preIndex
			skippy, err := skipMc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PythiaParameters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PythiaParameters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PythiaParameters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeightSum", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.WeightSum = &v2
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MergingWeight", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.MergingWeight = &v2
		default:
			iNdEx = preIndex
			skippy, err := skipMc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowMc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowMc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipMc(dAtA[start:])
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
	ErrInvalidLengthMc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proio/model/mc.proto", fileDescriptorMc) }

var fileDescriptorMc = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x6a, 0xe3, 0x30,
	0x14, 0x87, 0x47, 0x4e, 0xe2, 0xc4, 0x2f, 0x33, 0x21, 0x88, 0xf9, 0xa3, 0x81, 0x19, 0x13, 0x02,
	0x03, 0xd9, 0x8c, 0x3d, 0xf1, 0x11, 0x32, 0xb4, 0x74, 0x13, 0x48, 0x9d, 0x45, 0x4b, 0x37, 0x46,
	0x91, 0x55, 0x5b, 0x60, 0x45, 0xae, 0x2c, 0xd3, 0xf4, 0x0c, 0xbd, 0x40, 0x8f, 0xd4, 0x65, 0x8f,
	0x50, 0xd2, 0x8b, 0x94, 0xc8, 0x2e, 0x4d, 0x77, 0xef, 0xfb, 0x9e, 0x84, 0x9e, 0xde, 0x0f, 0xbe,
	0x96, 0x5a, 0x09, 0x15, 0x4a, 0x95, 0xf2, 0x22, 0x94, 0x2c, 0x28, 0xb5, 0x32, 0x0a, 0x8f, 0xac,
	0x0d, 0xac, 0x0d, 0x24, 0x9b, 0xde, 0x3b, 0xf0, 0x79, 0xf9, 0x7f, 0x45, 0x35, 0x95, 0xdc, 0x70,
	0x5d, 0xe1, 0xef, 0xe0, 0x6e, 0x6b, 0xb9, 0xe1, 0x9a, 0xa0, 0x09, 0x9a, 0x75, 0xe3, 0x96, 0xf0,
	0x2f, 0xf0, 0x4a, 0xad, 0x18, 0xaf, 0x2a, 0x91, 0x12, 0x67, 0x82, 0x66, 0xbd, 0xf8, 0x5d, 0xe0,
	0x6f, 0xe0, 0x96, 0x26, 0xc9, 0xa9, 0x21, 0x9d, 0x09, 0x9a, 0xa1, 0xb8, 0x57, 0x9a, 0x33, 0x6a,
	0xf0, 0x4f, 0x18, 0xd0, 0xa2, 0xcc, 0x69, 0xc2, 0x25, 0xe9, 0xda, 0x46, 0xdf, 0xf2, 0x89, 0xc4,
	0x3f, 0xa0, 0x29, 0x93, 0x8a, 0xf4, 0x6c, 0xc7, 0xb5, 0xb8, 0xc6, 0x3e, 0x0c, 0x2b, 0x46, 0x0b,
	0x9e, 0xdc, 0x24, 0xd7, 0x94, 0x11, 0xd7, 0x36, 0x3d, 0xab, 0xce, 0x4f, 0x29, 0x3b, 0x0c, 0x78,
	0xcb, 0x45, 0x96, 0x1b, 0xd2, 0x6f, 0xee, 0x35, 0x84, 0x47, 0xe0, 0xec, 0xe6, 0x64, 0x60, 0x9d,
	0xb3, 0x9b, 0x5b, 0x8e, 0x88, 0xd7, 0x72, 0x84, 0xc7, 0xd0, 0x11, 0xe9, 0x9c, 0x80, 0xfd, 0xd5,
	0xa1, 0x6c, 0x4c, 0x44, 0x86, 0x6f, 0x26, 0x9a, 0x5e, 0xc2, 0x78, 0x75, 0x67, 0x72, 0x41, 0x8f,
	0x16, 0xf2, 0x1b, 0xa0, 0x79, 0x21, 0xa9, 0x6a, 0x69, 0x97, 0x82, 0x62, 0xaf, 0x31, 0xeb, 0x5a,
	0xe2, 0x3f, 0x30, 0x92, 0x5c, 0x67, 0x62, 0x9b, 0x25, 0xed, 0x58, 0x8e, 0x3d, 0xf2, 0xa5, 0xb5,
	0x17, 0x56, 0x2e, 0xd6, 0x8f, 0x7b, 0x1f, 0x3d, 0xed, 0x7d, 0xf4, 0xbc, 0xf7, 0xd1, 0xc3, 0x8b,
	0xff, 0x09, 0x86, 0x47, 0x49, 0x2c, 0x9c, 0x25, 0xbb, 0xfa, 0x97, 0x09, 0x93, 0xd7, 0x9b, 0x80,
	0x29, 0x19, 0xa6, 0x9c, 0x89, 0x0d, 0x2f, 0x98, 0x52, 0x25, 0xd7, 0x61, 0x93, 0x64, 0xa6, 0xfe,
	0x7e, 0x8c, 0xf4, 0x35, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xe7, 0x67, 0x55, 0xe3, 0x01, 0x00, 0x00,
}