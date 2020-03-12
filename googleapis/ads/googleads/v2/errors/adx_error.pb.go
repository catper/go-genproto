// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/adx_error.proto

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

// Enum describing possible adx errors.
type AdxErrorEnum_AdxError int32

const (
	// Enum unspecified.
	AdxErrorEnum_UNSPECIFIED AdxErrorEnum_AdxError = 0
	// The received error code is not known in this version.
	AdxErrorEnum_UNKNOWN AdxErrorEnum_AdxError = 1
	// Attempt to use non-AdX feature by AdX customer.
	AdxErrorEnum_UNSUPPORTED_FEATURE AdxErrorEnum_AdxError = 2
)

var AdxErrorEnum_AdxError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UNSUPPORTED_FEATURE",
}

var AdxErrorEnum_AdxError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"UNSUPPORTED_FEATURE": 2,
}

func (x AdxErrorEnum_AdxError) String() string {
	return proto.EnumName(AdxErrorEnum_AdxError_name, int32(x))
}

func (AdxErrorEnum_AdxError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_64584b858be7dbce, []int{0, 0}
}

// Container for enum describing possible adx errors.
type AdxErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdxErrorEnum) Reset()         { *m = AdxErrorEnum{} }
func (m *AdxErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdxErrorEnum) ProtoMessage()    {}
func (*AdxErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_64584b858be7dbce, []int{0}
}

func (m *AdxErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdxErrorEnum.Unmarshal(m, b)
}
func (m *AdxErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdxErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdxErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdxErrorEnum.Merge(m, src)
}
func (m *AdxErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdxErrorEnum.Size(m)
}
func (m *AdxErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdxErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdxErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.AdxErrorEnum_AdxError", AdxErrorEnum_AdxError_name, AdxErrorEnum_AdxError_value)
	proto.RegisterType((*AdxErrorEnum)(nil), "google.ads.googleads.v2.errors.AdxErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/adx_error.proto", fileDescriptor_64584b858be7dbce)
}

var fileDescriptor_64584b858be7dbce = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xbf, 0xf6, 0x03, 0x95, 0x4c, 0xb1, 0xd4, 0x0b, 0x41, 0x64, 0x17, 0x7d, 0x80, 0x04,
	0xea, 0x5d, 0xbc, 0xca, 0x6c, 0x36, 0x86, 0xd0, 0xd5, 0x6d, 0xad, 0x20, 0x85, 0x11, 0x4d, 0x09,
	0x85, 0x2d, 0x29, 0x4d, 0x1d, 0x7b, 0x1e, 0x2f, 0x7d, 0x14, 0x9f, 0x44, 0x7c, 0x0a, 0x69, 0x8f,
	0xed, 0x9d, 0x5e, 0xe5, 0x97, 0xc3, 0xef, 0x9c, 0xf3, 0xe7, 0x20, 0xac, 0x8c, 0x51, 0xdb, 0x82,
	0x08, 0x69, 0x09, 0x60, 0x4b, 0xfb, 0x90, 0x14, 0x75, 0x6d, 0x6a, 0x4b, 0x84, 0x3c, 0x6c, 0x3a,
	0xc4, 0x55, 0x6d, 0x1a, 0xe3, 0x8f, 0x41, 0xc2, 0x42, 0x5a, 0x3c, 0xf8, 0x78, 0x1f, 0x62, 0xf0,
	0xaf, 0xae, 0xfb, 0x79, 0x55, 0x49, 0x84, 0xd6, 0xa6, 0x11, 0x4d, 0x69, 0xb4, 0x85, 0xee, 0xe0,
	0x01, 0x9d, 0x32, 0x79, 0xe0, 0xad, 0xca, 0xf5, 0xeb, 0x2e, 0x60, 0xe8, 0xa4, 0xff, 0xfb, 0xe7,
	0x68, 0x94, 0xc6, 0xab, 0x84, 0xdf, 0xcd, 0xa7, 0x73, 0x1e, 0x79, 0xff, 0xfc, 0x11, 0x3a, 0x4e,
	0xe3, 0xfb, 0x78, 0xf1, 0x18, 0x7b, 0x8e, 0x7f, 0x89, 0x2e, 0xd2, 0x78, 0x95, 0x26, 0xc9, 0x62,
	0xb9, 0xe6, 0xd1, 0x66, 0xca, 0xd9, 0x3a, 0x5d, 0x72, 0xcf, 0x9d, 0x7c, 0x3a, 0x28, 0x78, 0x31,
	0x3b, 0xfc, 0x77, 0xae, 0xc9, 0x59, 0xbf, 0x27, 0x69, 0x83, 0x24, 0xce, 0x53, 0xf4, 0xd3, 0xa0,
	0xcc, 0x56, 0x68, 0x85, 0x4d, 0xad, 0x88, 0x2a, 0x74, 0x17, 0xb3, 0x3f, 0x44, 0x55, 0xda, 0xdf,
	0xee, 0x72, 0x0b, 0xcf, 0x9b, 0xfb, 0x7f, 0xc6, 0xd8, 0xbb, 0x3b, 0x9e, 0xc1, 0x30, 0x26, 0x2d,
	0x06, 0x6c, 0x29, 0x0b, 0x71, 0xb7, 0xd2, 0x7e, 0xf4, 0x42, 0xce, 0xa4, 0xcd, 0x07, 0x21, 0xcf,
	0xc2, 0x1c, 0x84, 0x2f, 0x37, 0x80, 0x2a, 0xa5, 0x4c, 0x5a, 0x4a, 0x07, 0x85, 0xd2, 0x2c, 0xa4,
	0x14, 0xa4, 0xe7, 0xa3, 0x2e, 0xdd, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0xee, 0x4e,
	0x83, 0xb4, 0x01, 0x00, 0x00,
}
