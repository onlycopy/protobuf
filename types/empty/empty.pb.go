// Code generated by protoc-gen-go.
// source: empty.proto
// DO NOT EDIT!

/*
Package empty is a generated protocol buffer package.

It is generated from these files:
	empty.proto

It has these top-level messages:
	Empty
*/
package empty

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Empty ...
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Empty)(nil), "dkfbasel.protobuf.Empty")
}

func init() { proto.RegisterFile("empty.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0x2d, 0x28,
	0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xc9, 0x4e, 0x4b, 0x4a, 0x2c, 0x4e,
	0xcd, 0x81, 0xf0, 0x93, 0x4a, 0xd3, 0x94, 0xd8, 0xb9, 0x58, 0x5d, 0x41, 0x2a, 0x9c, 0xb4, 0xa2,
	0x34, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x61, 0x0a, 0xf5, 0x61,
	0x0a, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0xc1, 0xa6, 0x25, 0xb1, 0x81, 0x45, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0x91, 0x36, 0xfb, 0x5d, 0x00, 0x00, 0x00,
}
