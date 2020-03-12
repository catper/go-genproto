// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/resource_access_denied_error.proto

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

// Enum describing possible resource access denied errors.
type ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError int32

const (
	// Enum unspecified.
	ResourceAccessDeniedErrorEnum_UNSPECIFIED ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 0
	// The received error code is not known in this version.
	ResourceAccessDeniedErrorEnum_UNKNOWN ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 1
	// User did not have write access.
	ResourceAccessDeniedErrorEnum_WRITE_ACCESS_DENIED ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 3
)

var ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "WRITE_ACCESS_DENIED",
}

var ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"WRITE_ACCESS_DENIED": 3,
}

func (x ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError) String() string {
	return proto.EnumName(ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name, int32(x))
}

func (ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ebcb9332861fe0ce, []int{0, 0}
}

// Container for enum describing possible resource access denied errors.
type ResourceAccessDeniedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceAccessDeniedErrorEnum) Reset()         { *m = ResourceAccessDeniedErrorEnum{} }
func (m *ResourceAccessDeniedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ResourceAccessDeniedErrorEnum) ProtoMessage()    {}
func (*ResourceAccessDeniedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebcb9332861fe0ce, []int{0}
}

func (m *ResourceAccessDeniedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Unmarshal(m, b)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Marshal(b, m, deterministic)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceAccessDeniedErrorEnum.Merge(m, src)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Size(m)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceAccessDeniedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceAccessDeniedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError", ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name, ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_value)
	proto.RegisterType((*ResourceAccessDeniedErrorEnum)(nil), "google.ads.googleads.v1.errors.ResourceAccessDeniedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/resource_access_denied_error.proto", fileDescriptor_ebcb9332861fe0ce)
}

var fileDescriptor_ebcb9332861fe0ce = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x6a, 0x32, 0x31,
	0x14, 0x85, 0x7f, 0x15, 0xfe, 0x42, 0x5c, 0x54, 0xa6, 0x8b, 0xd2, 0xd2, 0xba, 0x98, 0x07, 0x48,
	0x18, 0xba, 0x4b, 0x57, 0xd1, 0x49, 0x45, 0x0a, 0x53, 0xd1, 0xaa, 0x50, 0x06, 0x86, 0x74, 0x12,
	0xc2, 0x80, 0x26, 0x92, 0xab, 0x3e, 0x50, 0x97, 0x7d, 0x94, 0x3e, 0x4a, 0xdf, 0xa0, 0xbb, 0x32,
	0x89, 0xce, 0x6e, 0xba, 0xca, 0x21, 0xf9, 0xee, 0x39, 0x27, 0x17, 0x31, 0x6d, 0xad, 0xde, 0x28,
	0x22, 0x24, 0x90, 0x20, 0x6b, 0x75, 0x4c, 0x88, 0x72, 0xce, 0x3a, 0x20, 0x4e, 0x81, 0x3d, 0xb8,
	0x52, 0x15, 0xa2, 0x2c, 0x15, 0x40, 0x21, 0x95, 0xa9, 0x94, 0x2c, 0xfc, 0x2b, 0xde, 0x39, 0xbb,
	0xb7, 0xd1, 0x30, 0xcc, 0x61, 0x21, 0x01, 0x37, 0x16, 0xf8, 0x98, 0xe0, 0x60, 0x71, 0x7b, 0x77,
	0x8e, 0xd8, 0x55, 0x44, 0x18, 0x63, 0xf7, 0x62, 0x5f, 0x59, 0x03, 0x61, 0x3a, 0x06, 0x74, 0x3f,
	0x3f, 0x65, 0x30, 0x1f, 0x91, 0xfa, 0x04, 0x5e, 0xcf, 0x72, 0x73, 0xd8, 0xc6, 0x73, 0x74, 0xd3,
	0x0a, 0x44, 0x97, 0xa8, 0xbf, 0xcc, 0x16, 0x33, 0x3e, 0x9e, 0x3e, 0x4d, 0x79, 0x3a, 0xf8, 0x17,
	0xf5, 0xd1, 0xc5, 0x32, 0x7b, 0xce, 0x5e, 0xd6, 0xd9, 0xa0, 0x13, 0x5d, 0xa3, 0xab, 0xf5, 0x7c,
	0xfa, 0xca, 0x0b, 0x36, 0x1e, 0xf3, 0xc5, 0xa2, 0x48, 0x79, 0x56, 0x53, 0xbd, 0xd1, 0x4f, 0x07,
	0xc5, 0xa5, 0xdd, 0xe2, 0xbf, 0x9b, 0x8f, 0x86, 0xad, 0xc1, 0xb3, 0xba, 0xfb, 0xac, 0xf3, 0x96,
	0x9e, 0x1c, 0xb4, 0xdd, 0x08, 0xa3, 0xb1, 0x75, 0x9a, 0x68, 0x65, 0xfc, 0xcf, 0xce, 0xeb, 0xdc,
	0x55, 0xd0, 0xb6, 0xdd, 0xc7, 0x70, 0x7c, 0x74, 0x7b, 0x13, 0xc6, 0x3e, 0xbb, 0xc3, 0x49, 0x30,
	0x63, 0x12, 0x70, 0x90, 0xb5, 0x5a, 0x25, 0xd8, 0x47, 0xc2, 0xd7, 0x19, 0xc8, 0x99, 0x84, 0xbc,
	0x01, 0xf2, 0x55, 0x92, 0x07, 0xe0, 0xbb, 0x1b, 0x87, 0x5b, 0x4a, 0x99, 0x04, 0x4a, 0x1b, 0x84,
	0xd2, 0x55, 0x42, 0x69, 0x80, 0xde, 0xff, 0xfb, 0x76, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0xe6, 0xc2, 0x05, 0xfa, 0x01, 0x00, 0x00,
}
