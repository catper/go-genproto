// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/size_limit_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible size limit errors.
type SizeLimitErrorEnum_SizeLimitError int32

const (
	// Enum unspecified.
	SizeLimitErrorEnum_UNSPECIFIED SizeLimitErrorEnum_SizeLimitError = 0
	// The received error code is not known in this version.
	SizeLimitErrorEnum_UNKNOWN SizeLimitErrorEnum_SizeLimitError = 1
	// The number of entries in the request exceeds the system limit.
	SizeLimitErrorEnum_REQUEST_SIZE_LIMIT_EXCEEDED SizeLimitErrorEnum_SizeLimitError = 2
	// The number of entries in the response exceeds the system limit.
	SizeLimitErrorEnum_RESPONSE_SIZE_LIMIT_EXCEEDED SizeLimitErrorEnum_SizeLimitError = 3
)

var SizeLimitErrorEnum_SizeLimitError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "REQUEST_SIZE_LIMIT_EXCEEDED",
	3: "RESPONSE_SIZE_LIMIT_EXCEEDED",
}

var SizeLimitErrorEnum_SizeLimitError_value = map[string]int32{
	"UNSPECIFIED":                  0,
	"UNKNOWN":                      1,
	"REQUEST_SIZE_LIMIT_EXCEEDED":  2,
	"RESPONSE_SIZE_LIMIT_EXCEEDED": 3,
}

func (x SizeLimitErrorEnum_SizeLimitError) String() string {
	return proto.EnumName(SizeLimitErrorEnum_SizeLimitError_name, int32(x))
}

func (SizeLimitErrorEnum_SizeLimitError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bcc468de64921988, []int{0, 0}
}

// Container for enum describing possible size limit errors.
type SizeLimitErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizeLimitErrorEnum) Reset()         { *m = SizeLimitErrorEnum{} }
func (m *SizeLimitErrorEnum) String() string { return proto.CompactTextString(m) }
func (*SizeLimitErrorEnum) ProtoMessage()    {}
func (*SizeLimitErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcc468de64921988, []int{0}
}

func (m *SizeLimitErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SizeLimitErrorEnum.Unmarshal(m, b)
}
func (m *SizeLimitErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SizeLimitErrorEnum.Marshal(b, m, deterministic)
}
func (m *SizeLimitErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizeLimitErrorEnum.Merge(m, src)
}
func (m *SizeLimitErrorEnum) XXX_Size() int {
	return xxx_messageInfo_SizeLimitErrorEnum.Size(m)
}
func (m *SizeLimitErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SizeLimitErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SizeLimitErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.SizeLimitErrorEnum_SizeLimitError", SizeLimitErrorEnum_SizeLimitError_name, SizeLimitErrorEnum_SizeLimitError_value)
	proto.RegisterType((*SizeLimitErrorEnum)(nil), "google.ads.googleads.v1.errors.SizeLimitErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/size_limit_error.proto", fileDescriptor_bcc468de64921988)
}

var fileDescriptor_bcc468de64921988 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x4d, 0x0a, 0x0a, 0x53, 0xd0, 0x12, 0x77, 0x5a, 0xaa, 0xe4, 0x00, 0x13, 0x82, 0xb8,
	0x19, 0x57, 0x69, 0xf3, 0x2c, 0xc1, 0x9a, 0xc6, 0x4e, 0x5b, 0xa5, 0x04, 0x42, 0x34, 0x61, 0x18,
	0x68, 0x67, 0x6a, 0x26, 0x76, 0xd1, 0x0b, 0x78, 0x0f, 0x97, 0x1e, 0xc5, 0xa3, 0xb8, 0xf0, 0x0c,
	0x92, 0x8c, 0x2d, 0x14, 0xd4, 0xd5, 0xfc, 0x3c, 0xbe, 0xff, 0x9f, 0xf7, 0x3f, 0x74, 0xc9, 0xa4,
	0x64, 0xf3, 0xdc, 0x49, 0x33, 0xe5, 0x68, 0x59, 0xa9, 0x95, 0xeb, 0xe4, 0x45, 0x21, 0x0b, 0xe5,
	0x28, 0xbe, 0xce, 0x93, 0x39, 0x5f, 0xf0, 0x32, 0xa9, 0x27, 0x78, 0x59, 0xc8, 0x52, 0x5a, 0x1d,
	0xcd, 0xe2, 0x34, 0x53, 0x78, 0x6b, 0xc3, 0x2b, 0x17, 0x6b, 0xdb, 0x49, 0x7b, 0x13, 0xbb, 0xe4,
	0x4e, 0x2a, 0x84, 0x2c, 0xd3, 0x92, 0x4b, 0xa1, 0xb4, 0xdb, 0x7e, 0x35, 0x90, 0x45, 0xf9, 0x3a,
	0x1f, 0x54, 0xb9, 0x50, 0x39, 0x40, 0xbc, 0x2c, 0xec, 0x67, 0x74, 0xb8, 0x3b, 0xb5, 0x8e, 0x50,
	0x73, 0x12, 0xd2, 0x08, 0x7a, 0xc1, 0x75, 0x00, 0x7e, 0x6b, 0xcf, 0x6a, 0xa2, 0x83, 0x49, 0x78,
	0x13, 0x0e, 0xef, 0xc3, 0x96, 0x61, 0x9d, 0xa1, 0xd3, 0x11, 0xdc, 0x4d, 0x80, 0x8e, 0x13, 0x1a,
	0xcc, 0x20, 0x19, 0x04, 0xb7, 0xc1, 0x38, 0x81, 0x87, 0x1e, 0x80, 0x0f, 0x7e, 0xcb, 0xb4, 0xce,
	0x51, 0x7b, 0x04, 0x34, 0x1a, 0x86, 0x14, 0x7e, 0x25, 0x1a, 0xdd, 0x2f, 0x03, 0xd9, 0x4f, 0x72,
	0x81, 0xff, 0xaf, 0xd3, 0x3d, 0xde, 0xdd, 0x2b, 0xaa, 0x5a, 0x44, 0xc6, 0xcc, 0xff, 0xb1, 0x31,
	0x39, 0x4f, 0x05, 0xc3, 0xb2, 0x60, 0x0e, 0xcb, 0x45, 0xdd, 0x71, 0x73, 0xcc, 0x25, 0x57, 0x7f,
	0xdd, 0xf6, 0x4a, 0x3f, 0x6f, 0x66, 0xa3, 0xef, 0x79, 0xef, 0x66, 0xa7, 0xaf, 0xc3, 0xbc, 0x4c,
	0x61, 0x2d, 0x2b, 0x35, 0x75, 0x71, 0xfd, 0xa5, 0xfa, 0xd8, 0x00, 0xb1, 0x97, 0xa9, 0x78, 0x0b,
	0xc4, 0x53, 0x37, 0xd6, 0xc0, 0xa7, 0x69, 0xeb, 0x29, 0x21, 0x5e, 0xa6, 0x08, 0xd9, 0x22, 0x84,
	0x4c, 0x5d, 0x42, 0x34, 0xf4, 0xb8, 0x5f, 0x6f, 0x77, 0xf1, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x16,
	0xd6, 0x32, 0xa2, 0xf8, 0x01, 0x00, 0x00,
}
