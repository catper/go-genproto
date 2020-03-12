// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/quota_error.proto

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

// Enum describing possible quota errors.
type QuotaErrorEnum_QuotaError int32

const (
	// Enum unspecified.
	QuotaErrorEnum_UNSPECIFIED QuotaErrorEnum_QuotaError = 0
	// The received error code is not known in this version.
	QuotaErrorEnum_UNKNOWN QuotaErrorEnum_QuotaError = 1
	// Too many requests.
	QuotaErrorEnum_RESOURCE_EXHAUSTED QuotaErrorEnum_QuotaError = 2
	// Access is prohibited.
	QuotaErrorEnum_ACCESS_PROHIBITED QuotaErrorEnum_QuotaError = 3
	// Too many requests in a short amount of time.
	QuotaErrorEnum_RESOURCE_TEMPORARILY_EXHAUSTED QuotaErrorEnum_QuotaError = 4
)

var QuotaErrorEnum_QuotaError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "RESOURCE_EXHAUSTED",
	3: "ACCESS_PROHIBITED",
	4: "RESOURCE_TEMPORARILY_EXHAUSTED",
}

var QuotaErrorEnum_QuotaError_value = map[string]int32{
	"UNSPECIFIED":                    0,
	"UNKNOWN":                        1,
	"RESOURCE_EXHAUSTED":             2,
	"ACCESS_PROHIBITED":              3,
	"RESOURCE_TEMPORARILY_EXHAUSTED": 4,
}

func (x QuotaErrorEnum_QuotaError) String() string {
	return proto.EnumName(QuotaErrorEnum_QuotaError_name, int32(x))
}

func (QuotaErrorEnum_QuotaError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_74f0d03d4b2122e8, []int{0, 0}
}

// Container for enum describing possible quota errors.
type QuotaErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaErrorEnum) Reset()         { *m = QuotaErrorEnum{} }
func (m *QuotaErrorEnum) String() string { return proto.CompactTextString(m) }
func (*QuotaErrorEnum) ProtoMessage()    {}
func (*QuotaErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f0d03d4b2122e8, []int{0}
}

func (m *QuotaErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaErrorEnum.Unmarshal(m, b)
}
func (m *QuotaErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaErrorEnum.Marshal(b, m, deterministic)
}
func (m *QuotaErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaErrorEnum.Merge(m, src)
}
func (m *QuotaErrorEnum) XXX_Size() int {
	return xxx_messageInfo_QuotaErrorEnum.Size(m)
}
func (m *QuotaErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.QuotaErrorEnum_QuotaError", QuotaErrorEnum_QuotaError_name, QuotaErrorEnum_QuotaError_value)
	proto.RegisterType((*QuotaErrorEnum)(nil), "google.ads.googleads.v3.errors.QuotaErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/quota_error.proto", fileDescriptor_74f0d03d4b2122e8)
}

var fileDescriptor_74f0d03d4b2122e8 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4e, 0x83, 0x40,
	0x18, 0xc5, 0x85, 0x1a, 0x4d, 0xa6, 0x89, 0xad, 0x93, 0xe8, 0xc2, 0x98, 0x2e, 0x38, 0xc0, 0x60,
	0xc2, 0x6e, 0x5c, 0x4d, 0xe9, 0xd8, 0x12, 0x15, 0x10, 0x4a, 0xfd, 0x13, 0x92, 0x06, 0xa5, 0x99,
	0x34, 0x69, 0x67, 0x2a, 0x43, 0xbb, 0xf3, 0x0c, 0xde, 0xc1, 0xa5, 0x47, 0xf1, 0x28, 0x3d, 0x85,
	0x19, 0x46, 0xa8, 0x1b, 0x5d, 0xf1, 0xf8, 0xf2, 0x7b, 0x6f, 0xbe, 0xf7, 0x81, 0x0b, 0x26, 0x04,
	0x5b, 0xcc, 0xec, 0x2c, 0x97, 0xb6, 0x96, 0x4a, 0x6d, 0x1c, 0x7b, 0x56, 0x14, 0xa2, 0x90, 0xf6,
	0xeb, 0x5a, 0x94, 0xd9, 0xb4, 0xfa, 0x41, 0xab, 0x42, 0x94, 0x02, 0xf6, 0x34, 0x86, 0xb2, 0x5c,
	0xa2, 0xc6, 0x81, 0x36, 0x0e, 0xd2, 0x8e, 0xb3, 0xf3, 0x3a, 0x71, 0x35, 0xb7, 0x33, 0xce, 0x45,
	0x99, 0x95, 0x73, 0xc1, 0xa5, 0x76, 0x5b, 0xef, 0x06, 0x38, 0xba, 0x53, 0x99, 0x54, 0xd1, 0x94,
	0xaf, 0x97, 0xd6, 0x1b, 0x00, 0xbb, 0x09, 0xec, 0x80, 0x76, 0xe2, 0xc7, 0x21, 0x75, 0xbd, 0x2b,
	0x8f, 0x0e, 0xba, 0x7b, 0xb0, 0x0d, 0x0e, 0x13, 0xff, 0xda, 0x0f, 0xee, 0xfd, 0xae, 0x01, 0x4f,
	0x01, 0x8c, 0x68, 0x1c, 0x24, 0x91, 0x4b, 0xa7, 0xf4, 0x61, 0x44, 0x92, 0x78, 0x4c, 0x07, 0x5d,
	0x13, 0x9e, 0x80, 0x63, 0xe2, 0xba, 0x34, 0x8e, 0xa7, 0x61, 0x14, 0x8c, 0xbc, 0xbe, 0xa7, 0xc6,
	0x2d, 0x68, 0x81, 0x5e, 0x83, 0x8f, 0xe9, 0x6d, 0x18, 0x44, 0x24, 0xf2, 0x6e, 0x1e, 0x7f, 0x59,
	0xf7, 0xfb, 0x5b, 0x03, 0x58, 0x2f, 0x62, 0x89, 0xfe, 0xaf, 0xd5, 0xef, 0xec, 0x76, 0x0c, 0x55,
	0x93, 0xd0, 0x78, 0x1a, 0xfc, 0x58, 0x98, 0x58, 0x64, 0x9c, 0x21, 0x51, 0x30, 0x9b, 0xcd, 0x78,
	0xd5, 0xb3, 0xbe, 0xe5, 0x6a, 0x2e, 0xff, 0x3a, 0xed, 0xa5, 0xfe, 0x7c, 0x98, 0xad, 0x21, 0x21,
	0x9f, 0x66, 0x6f, 0xa8, 0xc3, 0x48, 0x2e, 0x91, 0x96, 0x4a, 0x4d, 0x1c, 0x54, 0x3d, 0x29, 0xbf,
	0x6a, 0x20, 0x25, 0xb9, 0x4c, 0x1b, 0x20, 0x9d, 0x38, 0xa9, 0x06, 0xb6, 0xa6, 0xa5, 0xa7, 0x18,
	0x93, 0x5c, 0x62, 0xdc, 0x20, 0x18, 0x4f, 0x1c, 0x8c, 0x35, 0xf4, 0x7c, 0x50, 0x6d, 0xe7, 0x7c,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x06, 0xb6, 0xd1, 0xf7, 0x01, 0x00, 0x00,
}
