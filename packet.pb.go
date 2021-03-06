// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packet.proto

/*
	Package grpcx is a generated protocol buffer package.

	It is generated from these files:
		packet.proto

	It has these top-level messages:
		PackHeader
		NetPack
*/
package grpcx

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type CompType int32

const (
	CompType_None CompType = 0
	CompType_GZip CompType = 1
)

var CompType_name = map[int32]string{
	0: "None",
	1: "GZip",
}
var CompType_value = map[string]int32{
	"None": 0,
	"GZip": 1,
}

func (x CompType) String() string {
	return proto.EnumName(CompType_name, int32(x))
}
func (CompType) EnumDescriptor() ([]byte, []int) { return fileDescriptorPacket, []int{0} }

type PackType int32

const (
	PackType_UNKNOW PackType = 0
	PackType_SINI   PackType = 1
	PackType_SREQ   PackType = 2
	PackType_SRSP   PackType = 3
	PackType_REQ    PackType = 4
	PackType_RSP    PackType = 5
	PackType_Notify PackType = 6
	PackType_PING   PackType = 7
	PackType_PONG   PackType = 8
	PackType_EOF    PackType = 9
	PackType_ERROR  PackType = 10
)

var PackType_name = map[int32]string{
	0:  "UNKNOW",
	1:  "SINI",
	2:  "SREQ",
	3:  "SRSP",
	4:  "REQ",
	5:  "RSP",
	6:  "Notify",
	7:  "PING",
	8:  "PONG",
	9:  "EOF",
	10: "ERROR",
}
var PackType_value = map[string]int32{
	"UNKNOW": 0,
	"SINI":   1,
	"SREQ":   2,
	"SRSP":   3,
	"REQ":    4,
	"RSP":    5,
	"Notify": 6,
	"PING":   7,
	"PONG":   8,
	"EOF":    9,
	"ERROR":  10,
}

func (x PackType) String() string {
	return proto.EnumName(PackType_name, int32(x))
}
func (PackType) EnumDescriptor() ([]byte, []int) { return fileDescriptorPacket, []int{1} }

type PackHeader struct {
	Ptype     PackType         `protobuf:"varint,1,opt,name=ptype,proto3,enum=grpcx.PackType" json:"ptype,omitempty"`
	Methord   string           `protobuf:"bytes,2,opt,name=methord,proto3" json:"methord,omitempty"`
	Sessionid int64            `protobuf:"varint,3,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
	Metadata  map[int32][]byte `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *PackHeader) Reset()                    { *m = PackHeader{} }
func (m *PackHeader) String() string            { return proto.CompactTextString(m) }
func (*PackHeader) ProtoMessage()               {}
func (*PackHeader) Descriptor() ([]byte, []int) { return fileDescriptorPacket, []int{0} }

func (m *PackHeader) GetPtype() PackType {
	if m != nil {
		return m.Ptype
	}
	return PackType_UNKNOW
}

func (m *PackHeader) GetMethord() string {
	if m != nil {
		return m.Methord
	}
	return ""
}

func (m *PackHeader) GetSessionid() int64 {
	if m != nil {
		return m.Sessionid
	}
	return 0
}

func (m *PackHeader) GetMetadata() map[int32][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type NetPack struct {
	Header *PackHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Body   []byte      `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *NetPack) Reset()                    { *m = NetPack{} }
func (m *NetPack) String() string            { return proto.CompactTextString(m) }
func (*NetPack) ProtoMessage()               {}
func (*NetPack) Descriptor() ([]byte, []int) { return fileDescriptorPacket, []int{1} }

func (m *NetPack) GetHeader() *PackHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NetPack) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*PackHeader)(nil), "grpcx.PackHeader")
	proto.RegisterType((*NetPack)(nil), "grpcx.NetPack")
	proto.RegisterEnum("grpcx.CompType", CompType_name, CompType_value)
	proto.RegisterEnum("grpcx.PackType", PackType_name, PackType_value)
}
func (m *PackHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PackHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ptype != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPacket(dAtA, i, uint64(m.Ptype))
	}
	if len(m.Methord) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Methord)))
		i += copy(dAtA[i:], m.Methord)
	}
	if m.Sessionid != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPacket(dAtA, i, uint64(m.Sessionid))
	}
	if len(m.Metadata) > 0 {
		for k, _ := range m.Metadata {
			dAtA[i] = 0x22
			i++
			v := m.Metadata[k]
			byteSize := 0
			if len(v) > 0 {
				byteSize = 1 + len(v) + sovPacket(uint64(len(v)))
			}
			mapSize := 1 + sovPacket(uint64(k)) + byteSize
			i = encodeVarintPacket(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintPacket(dAtA, i, uint64(k))
			if len(v) > 0 {
				dAtA[i] = 0x12
				i++
				i = encodeVarintPacket(dAtA, i, uint64(len(v)))
				i += copy(dAtA[i:], v)
			}
		}
	}
	return i, nil
}

func (m *NetPack) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetPack) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPacket(dAtA, i, uint64(m.Header.Size()))
		n1, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	return i, nil
}

func encodeFixed64Packet(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Packet(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PackHeader) Size() (n int) {
	var l int
	_ = l
	if m.Ptype != 0 {
		n += 1 + sovPacket(uint64(m.Ptype))
	}
	l = len(m.Methord)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.Sessionid != 0 {
		n += 1 + sovPacket(uint64(m.Sessionid))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovPacket(uint64(len(v)))
			}
			mapEntrySize := 1 + sovPacket(uint64(k)) + l
			n += mapEntrySize + 1 + sovPacket(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *NetPack) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PackHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: PackHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PackHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ptype", wireType)
			}
			m.Ptype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ptype |= (PackType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Methord", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Methord = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sessionid", wireType)
			}
			m.Sessionid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sessionid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapkey int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				mapkey |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if m.Metadata == nil {
				m.Metadata = make(map[int32][]byte)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPacket
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapbyteLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPacket
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapbyteLen |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intMapbyteLen := int(mapbyteLen)
				if intMapbyteLen < 0 {
					return ErrInvalidLengthPacket
				}
				postbytesIndex := iNdEx + intMapbyteLen
				if postbytesIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := make([]byte, mapbyteLen)
				copy(mapvalue, dAtA[iNdEx:postbytesIndex])
				iNdEx = postbytesIndex
				m.Metadata[mapkey] = mapvalue
			} else {
				var mapvalue []byte
				m.Metadata[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPacket
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
func (m *NetPack) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: NetPack: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetPack: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &PackHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPacket
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
				next, err := skipPacket(dAtA[start:])
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
	ErrInvalidLengthPacket = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("packet.proto", fileDescriptorPacket) }

var fileDescriptorPacket = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0xcd, 0xf7, 0x69, 0xd5, 0x71, 0x70, 0x11, 0x44, 0x62, 0x28, 0x08, 0xb1, 0x8b,
	0x2c, 0xea, 0x46, 0xec, 0x4e, 0x89, 0x6d, 0x11, 0x27, 0x71, 0xaa, 0x08, 0xee, 0xd2, 0x66, 0xb4,
	0xa1, 0xb6, 0x33, 0xa4, 0xa3, 0x18, 0x9f, 0xc4, 0x47, 0x72, 0xe9, 0x23, 0x48, 0xc5, 0xf7, 0xb8,
	0xcc, 0xa4, 0xbd, 0x1f, 0xdc, 0xdd, 0x6f, 0xce, 0xf9, 0x9d, 0x9c, 0x3f, 0x27, 0x30, 0x94, 0xe5,
	0x7a, 0xcb, 0x55, 0x2a, 0x1b, 0xa1, 0x04, 0x71, 0xbe, 0x34, 0x72, 0xfd, 0x63, 0xf4, 0x1f, 0x01,
	0x14, 0xe5, 0x7a, 0x3b, 0xe7, 0x65, 0xc5, 0x1b, 0xf2, 0x04, 0x1c, 0xa9, 0x5a, 0xc9, 0x43, 0x14,
	0xa3, 0xe4, 0xee, 0xe4, 0x5e, 0x6a, 0xac, 0x54, 0x1b, 0xef, 0x5b, 0xc9, 0x59, 0xd7, 0x25, 0x21,
	0x78, 0x3b, 0xae, 0x36, 0xa2, 0xa9, 0xc2, 0x7e, 0x8c, 0x92, 0x80, 0x9d, 0x9f, 0xe4, 0x11, 0x04,
	0x07, 0x7e, 0x38, 0xd4, 0x62, 0x5f, 0x57, 0xa1, 0x15, 0xa3, 0xc4, 0x62, 0x57, 0x05, 0x32, 0x05,
	0x7f, 0xc7, 0x55, 0x59, 0x95, 0xaa, 0x0c, 0xed, 0xd8, 0x4a, 0x06, 0x93, 0xc7, 0xd7, 0x36, 0x74,
	0x19, 0xd2, 0xb7, 0x27, 0x23, 0xdb, 0xab, 0xa6, 0x65, 0x97, 0x03, 0x0f, 0xa7, 0x70, 0xe7, 0x46,
	0x8b, 0x60, 0xb0, 0xb6, 0xbc, 0x35, 0x51, 0x1d, 0xa6, 0x91, 0x3c, 0x00, 0xe7, 0x7b, 0xf9, 0xf5,
	0x1b, 0x37, 0xa9, 0x86, 0xac, 0x7b, 0xbc, 0xe8, 0x3f, 0x47, 0xa3, 0x39, 0x78, 0x94, 0x2b, 0xbd,
	0x85, 0x3c, 0x05, 0x77, 0x63, 0x36, 0x99, 0xc9, 0xc1, 0xe4, 0xfe, 0xad, 0x08, 0xec, 0x24, 0x10,
	0x02, 0xf6, 0x4a, 0x54, 0xed, 0xe9, 0x73, 0x86, 0xc7, 0x11, 0xf8, 0xaf, 0xc4, 0x4e, 0xea, 0x73,
	0x10, 0x1f, 0x6c, 0x2a, 0xf6, 0x1c, 0xf7, 0x34, 0xcd, 0x3e, 0xd5, 0x12, 0xa3, 0xf1, 0x4f, 0xf0,
	0xcf, 0xe7, 0x22, 0x00, 0xee, 0x07, 0xfa, 0x86, 0xe6, 0x1f, 0x3b, 0x63, 0xb9, 0xa0, 0x0b, 0x8c,
	0x0c, 0xb1, 0xec, 0x1d, 0xee, 0x77, 0xb4, 0x2c, 0xb0, 0x45, 0x3c, 0xb0, 0x74, 0xc9, 0x36, 0xb0,
	0x2c, 0xb0, 0xa3, 0x67, 0xa9, 0x50, 0xf5, 0xe7, 0x16, 0xbb, 0xda, 0x2b, 0x16, 0x74, 0x86, 0x3d,
	0x43, 0x39, 0x9d, 0x61, 0x5f, 0x8b, 0x59, 0xfe, 0x1a, 0x07, 0x24, 0x00, 0x27, 0x63, 0x2c, 0x67,
	0x18, 0x5e, 0xe2, 0xdf, 0xc7, 0x08, 0xfd, 0x39, 0x46, 0xe8, 0xef, 0x31, 0x42, 0xbf, 0xfe, 0x45,
	0xbd, 0x95, 0x6b, 0xfe, 0xf6, 0xb3, 0x8b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x98, 0x46, 0x9a,
	0xfd, 0x01, 0x00, 0x00,
}
