// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dbtesterpb/flag_consul.proto

package dbtesterpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// See https://github.com/hashicorp/consul for more.
type Flag_Consul_V0_7_5 struct {
}

func (m *Flag_Consul_V0_7_5) Reset()                    { *m = Flag_Consul_V0_7_5{} }
func (m *Flag_Consul_V0_7_5) String() string            { return proto.CompactTextString(m) }
func (*Flag_Consul_V0_7_5) ProtoMessage()               {}
func (*Flag_Consul_V0_7_5) Descriptor() ([]byte, []int) { return fileDescriptorFlagConsul, []int{0} }

// See https://github.com/hashicorp/consul for more.
type Flag_Consul_V0_8_0 struct {
}

func (m *Flag_Consul_V0_8_0) Reset()                    { *m = Flag_Consul_V0_8_0{} }
func (m *Flag_Consul_V0_8_0) String() string            { return proto.CompactTextString(m) }
func (*Flag_Consul_V0_8_0) ProtoMessage()               {}
func (*Flag_Consul_V0_8_0) Descriptor() ([]byte, []int) { return fileDescriptorFlagConsul, []int{1} }

// See https://github.com/hashicorp/consul for more.
type Flag_Consul_V0_8_4 struct {
}

func (m *Flag_Consul_V0_8_4) Reset()                    { *m = Flag_Consul_V0_8_4{} }
func (m *Flag_Consul_V0_8_4) String() string            { return proto.CompactTextString(m) }
func (*Flag_Consul_V0_8_4) ProtoMessage()               {}
func (*Flag_Consul_V0_8_4) Descriptor() ([]byte, []int) { return fileDescriptorFlagConsul, []int{2} }

func init() {
	proto.RegisterType((*Flag_Consul_V0_7_5)(nil), "dbtesterpb.flag__consul__v0_7_5")
	proto.RegisterType((*Flag_Consul_V0_8_0)(nil), "dbtesterpb.flag__consul__v0_8_0")
	proto.RegisterType((*Flag_Consul_V0_8_4)(nil), "dbtesterpb.flag__consul__v0_8_4")
}
func (m *Flag_Consul_V0_7_5) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Flag_Consul_V0_7_5) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Flag_Consul_V0_8_0) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Flag_Consul_V0_8_0) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Flag_Consul_V0_8_4) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Flag_Consul_V0_8_4) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64FlagConsul(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32FlagConsul(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFlagConsul(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Flag_Consul_V0_7_5) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Flag_Consul_V0_8_0) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Flag_Consul_V0_8_4) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovFlagConsul(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFlagConsul(x uint64) (n int) {
	return sovFlagConsul(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Flag_Consul_V0_7_5) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlagConsul
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
			return fmt.Errorf("proto: flag__consul__v0_7_5: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: flag__consul__v0_7_5: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFlagConsul(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlagConsul
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
func (m *Flag_Consul_V0_8_0) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlagConsul
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
			return fmt.Errorf("proto: flag__consul__v0_8_0: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: flag__consul__v0_8_0: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFlagConsul(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlagConsul
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
func (m *Flag_Consul_V0_8_4) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlagConsul
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
			return fmt.Errorf("proto: flag__consul__v0_8_4: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: flag__consul__v0_8_4: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFlagConsul(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlagConsul
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
func skipFlagConsul(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlagConsul
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
					return 0, ErrIntOverflowFlagConsul
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
					return 0, ErrIntOverflowFlagConsul
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
				return 0, ErrInvalidLengthFlagConsul
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFlagConsul
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
				next, err := skipFlagConsul(dAtA[start:])
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
	ErrInvalidLengthFlagConsul = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlagConsul   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("dbtesterpb/flag_consul.proto", fileDescriptorFlagConsul) }

var fileDescriptorFlagConsul = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x49, 0x2a, 0x49,
	0x2d, 0x2e, 0x49, 0x2d, 0x2a, 0x48, 0xd2, 0x4f, 0xcb, 0x49, 0x4c, 0x8f, 0x4f, 0xce, 0xcf, 0x2b,
	0x2e, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0xc8, 0x4a, 0xe9, 0xa6, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x95,
	0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xaa, 0x24, 0xc6, 0x25, 0x02, 0x36,
	0x0f, 0x6a, 0x60, 0x7c, 0x7c, 0x99, 0x41, 0xbc, 0x79, 0xbc, 0x29, 0x56, 0x71, 0x8b, 0x78, 0x03,
	0x1c, 0xe2, 0x26, 0x4e, 0x22, 0x27, 0x1e, 0xca, 0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0x2d, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x31, 0xd1, 0x99, 0x57, 0xbf, 0x00, 0x00, 0x00,
}
