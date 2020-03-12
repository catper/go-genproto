// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/system_managed_entity_source.proto

package enums

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

// Enum listing the possible system managed entity sources.
type SystemManagedResourceSourceEnum_SystemManagedResourceSource int32

const (
	// Not specified.
	SystemManagedResourceSourceEnum_UNSPECIFIED SystemManagedResourceSourceEnum_SystemManagedResourceSource = 0
	// Used for return value only. Represents value unknown in this version.
	SystemManagedResourceSourceEnum_UNKNOWN SystemManagedResourceSourceEnum_SystemManagedResourceSource = 1
	// Generated ad variations experiment ad.
	SystemManagedResourceSourceEnum_AD_VARIATIONS SystemManagedResourceSourceEnum_SystemManagedResourceSource = 2
)

var SystemManagedResourceSourceEnum_SystemManagedResourceSource_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_VARIATIONS",
}

var SystemManagedResourceSourceEnum_SystemManagedResourceSource_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"AD_VARIATIONS": 2,
}

func (x SystemManagedResourceSourceEnum_SystemManagedResourceSource) String() string {
	return proto.EnumName(SystemManagedResourceSourceEnum_SystemManagedResourceSource_name, int32(x))
}

func (SystemManagedResourceSourceEnum_SystemManagedResourceSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5fe06dec8282bd9, []int{0, 0}
}

// Container for enum describing possible system managed entity sources.
type SystemManagedResourceSourceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemManagedResourceSourceEnum) Reset()         { *m = SystemManagedResourceSourceEnum{} }
func (m *SystemManagedResourceSourceEnum) String() string { return proto.CompactTextString(m) }
func (*SystemManagedResourceSourceEnum) ProtoMessage()    {}
func (*SystemManagedResourceSourceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5fe06dec8282bd9, []int{0}
}

func (m *SystemManagedResourceSourceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemManagedResourceSourceEnum.Unmarshal(m, b)
}
func (m *SystemManagedResourceSourceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemManagedResourceSourceEnum.Marshal(b, m, deterministic)
}
func (m *SystemManagedResourceSourceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemManagedResourceSourceEnum.Merge(m, src)
}
func (m *SystemManagedResourceSourceEnum) XXX_Size() int {
	return xxx_messageInfo_SystemManagedResourceSourceEnum.Size(m)
}
func (m *SystemManagedResourceSourceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemManagedResourceSourceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SystemManagedResourceSourceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource", SystemManagedResourceSourceEnum_SystemManagedResourceSource_name, SystemManagedResourceSourceEnum_SystemManagedResourceSource_value)
	proto.RegisterType((*SystemManagedResourceSourceEnum)(nil), "google.ads.googleads.v3.enums.SystemManagedResourceSourceEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/system_managed_entity_source.proto", fileDescriptor_a5fe06dec8282bd9)
}

var fileDescriptor_a5fe06dec8282bd9 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x75, 0x15, 0x14, 0x32, 0xc4, 0xd9, 0x47, 0x75, 0xca, 0xf6, 0x01, 0xe9, 0x43, 0xdf, 0xe2,
	0x8b, 0x99, 0xab, 0xa3, 0x88, 0xdd, 0x58, 0x5d, 0x05, 0x29, 0x94, 0xb8, 0x84, 0x50, 0x58, 0x93,
	0xb9, 0x9b, 0x0d, 0xf6, 0x3b, 0x3e, 0xfa, 0x29, 0x7e, 0x8a, 0xaf, 0xfe, 0x80, 0x34, 0x71, 0x05,
	0x1f, 0xdc, 0x4b, 0x38, 0xe4, 0x9e, 0x7b, 0xce, 0xb9, 0x07, 0xdd, 0x4a, 0xad, 0xe5, 0x42, 0x04,
	0x8c, 0x43, 0xe0, 0x60, 0x8d, 0x36, 0x61, 0x20, 0xd4, 0xba, 0x82, 0x00, 0xb6, 0x60, 0x44, 0x55,
	0x54, 0x4c, 0x31, 0x29, 0x78, 0x21, 0x94, 0x29, 0xcd, 0xb6, 0x00, 0xbd, 0x5e, 0xcd, 0x05, 0x5e,
	0xae, 0xb4, 0xd1, 0x7e, 0xd7, 0xad, 0x61, 0xc6, 0x01, 0x37, 0x0a, 0x78, 0x13, 0x62, 0xab, 0x70,
	0x7e, 0xb9, 0x33, 0x58, 0x96, 0x01, 0x53, 0x4a, 0x1b, 0x66, 0x4a, 0xad, 0xc0, 0x2d, 0xf7, 0xdf,
	0xd0, 0x75, 0x6a, 0x2d, 0x1e, 0x9d, 0xc3, 0x54, 0x38, 0xed, 0xd4, 0xbe, 0x91, 0x5a, 0x57, 0xfd,
	0x04, 0x5d, 0xec, 0xa1, 0xf8, 0xa7, 0xa8, 0x3d, 0x4b, 0xd2, 0x49, 0x74, 0x17, 0xdf, 0xc7, 0xd1,
	0xb0, 0x73, 0xe0, 0xb7, 0xd1, 0xf1, 0x2c, 0x79, 0x48, 0xc6, 0xcf, 0x49, 0xa7, 0xe5, 0x9f, 0xa1,
	0x13, 0x3a, 0x2c, 0x32, 0x3a, 0x8d, 0xe9, 0x53, 0x3c, 0x4e, 0xd2, 0x8e, 0x37, 0xf8, 0x6e, 0xa1,
	0xde, 0x5c, 0x57, 0x78, 0x6f, 0xec, 0xc1, 0xd5, 0x1f, 0xcf, 0xc8, 0xde, 0xed, 0x1c, 0x27, 0x75,
	0xf0, 0x49, 0xeb, 0x65, 0xf0, 0x2b, 0x20, 0xf5, 0x82, 0x29, 0x89, 0xf5, 0x4a, 0x06, 0x52, 0x28,
	0x7b, 0xd6, 0xae, 0xc9, 0x65, 0x09, 0xff, 0x14, 0x7b, 0x63, 0xdf, 0x77, 0xef, 0x70, 0x44, 0xe9,
	0x87, 0xd7, 0x1d, 0x39, 0x29, 0xca, 0x01, 0x3b, 0x58, 0xa3, 0x2c, 0xc4, 0x75, 0x01, 0xf0, 0xb9,
	0x9b, 0xe7, 0x94, 0x43, 0xde, 0xcc, 0xf3, 0x2c, 0xcc, 0xed, 0xfc, 0xcb, 0xeb, 0xb9, 0x4f, 0x42,
	0x28, 0x07, 0x42, 0x1a, 0x06, 0x21, 0x59, 0x48, 0x88, 0xe5, 0xbc, 0x1e, 0xd9, 0x60, 0xe1, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x46, 0x13, 0x02, 0xf0, 0x01, 0x00, 0x00,
}
