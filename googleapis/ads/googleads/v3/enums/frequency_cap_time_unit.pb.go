// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/frequency_cap_time_unit.proto

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

// Unit of time the cap is defined at (e.g. day, week).
type FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit int32

const (
	// Not specified.
	FrequencyCapTimeUnitEnum_UNSPECIFIED FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit = 0
	// Used for return value only. Represents value unknown in this version.
	FrequencyCapTimeUnitEnum_UNKNOWN FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit = 1
	// The cap would define limit per one day.
	FrequencyCapTimeUnitEnum_DAY FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit = 2
	// The cap would define limit per one week.
	FrequencyCapTimeUnitEnum_WEEK FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit = 3
	// The cap would define limit per one month.
	FrequencyCapTimeUnitEnum_MONTH FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit = 4
)

var FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DAY",
	3: "WEEK",
	4: "MONTH",
}

var FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"DAY":         2,
	"WEEK":        3,
	"MONTH":       4,
}

func (x FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit) String() string {
	return proto.EnumName(FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit_name, int32(x))
}

func (FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0acbdb304da9e43b, []int{0, 0}
}

// Container for enum describing the unit of time the cap is defined at.
type FrequencyCapTimeUnitEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrequencyCapTimeUnitEnum) Reset()         { *m = FrequencyCapTimeUnitEnum{} }
func (m *FrequencyCapTimeUnitEnum) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapTimeUnitEnum) ProtoMessage()    {}
func (*FrequencyCapTimeUnitEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0acbdb304da9e43b, []int{0}
}

func (m *FrequencyCapTimeUnitEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapTimeUnitEnum.Unmarshal(m, b)
}
func (m *FrequencyCapTimeUnitEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapTimeUnitEnum.Marshal(b, m, deterministic)
}
func (m *FrequencyCapTimeUnitEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapTimeUnitEnum.Merge(m, src)
}
func (m *FrequencyCapTimeUnitEnum) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapTimeUnitEnum.Size(m)
}
func (m *FrequencyCapTimeUnitEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapTimeUnitEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapTimeUnitEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit", FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit_name, FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit_value)
	proto.RegisterType((*FrequencyCapTimeUnitEnum)(nil), "google.ads.googleads.v3.enums.FrequencyCapTimeUnitEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/frequency_cap_time_unit.proto", fileDescriptor_0acbdb304da9e43b)
}

var fileDescriptor_0acbdb304da9e43b = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4f, 0xc2, 0x30,
	0x18, 0xc6, 0x65, 0xa0, 0x68, 0x39, 0xb8, 0x2c, 0x1e, 0xd4, 0xc8, 0x01, 0x3e, 0x40, 0x77, 0xd8,
	0xad, 0x9c, 0x0a, 0x0c, 0x24, 0xc4, 0x41, 0x94, 0x3f, 0xd1, 0x2c, 0x21, 0x95, 0xd5, 0xa6, 0x09,
	0x6b, 0x27, 0xed, 0x48, 0xfc, 0x3a, 0x1e, 0xfd, 0x28, 0x7e, 0x14, 0x0f, 0x7e, 0x06, 0xb3, 0xd6,
	0x71, 0x42, 0x2f, 0xcb, 0x93, 0x3d, 0xef, 0xef, 0xe9, 0xf3, 0xbe, 0xa0, 0xc3, 0xa4, 0x64, 0x1b,
	0xea, 0x93, 0x44, 0xf9, 0x56, 0x16, 0x6a, 0x17, 0xf8, 0x54, 0xe4, 0xa9, 0xf2, 0x5f, 0xb6, 0xf4,
	0x35, 0xa7, 0x62, 0xfd, 0xb6, 0x5a, 0x93, 0x6c, 0xa5, 0x79, 0x4a, 0x57, 0xb9, 0xe0, 0x1a, 0x66,
	0x5b, 0xa9, 0xa5, 0xd7, 0xb4, 0x04, 0x24, 0x89, 0x82, 0x7b, 0x18, 0xee, 0x02, 0x68, 0xe0, 0xeb,
	0x9b, 0x32, 0x3b, 0xe3, 0x3e, 0x11, 0x42, 0x6a, 0xa2, 0xb9, 0x14, 0xca, 0xc2, 0x6d, 0x01, 0x2e,
	0x07, 0x65, 0x7a, 0x8f, 0x64, 0x33, 0x9e, 0xd2, 0xb9, 0xe0, 0x3a, 0x14, 0x79, 0xda, 0xbe, 0x07,
	0x17, 0x87, 0x3c, 0xef, 0x1c, 0x34, 0xe6, 0xd1, 0xc3, 0x34, 0xec, 0x8d, 0x06, 0xa3, 0xb0, 0xef,
	0x1e, 0x79, 0x0d, 0x50, 0x9f, 0x47, 0xe3, 0x68, 0xb2, 0x8c, 0xdc, 0x8a, 0x57, 0x07, 0xd5, 0x3e,
	0x7e, 0x74, 0x1d, 0xef, 0x14, 0xd4, 0x96, 0x61, 0x38, 0x76, 0xab, 0xde, 0x19, 0x38, 0xbe, 0x9b,
	0x44, 0xb3, 0x5b, 0xb7, 0xd6, 0xfd, 0xae, 0x80, 0xd6, 0x5a, 0xa6, 0xf0, 0xdf, 0xce, 0xdd, 0xab,
	0x43, 0xef, 0x4e, 0x8b, 0xc2, 0xd3, 0xca, 0x53, 0xf7, 0x97, 0x65, 0x72, 0x43, 0x04, 0x83, 0x72,
	0xcb, 0x7c, 0x46, 0x85, 0x59, 0xa7, 0x3c, 0x5e, 0xc6, 0xd5, 0x1f, 0xb7, 0xec, 0x98, 0xef, 0xbb,
	0x53, 0x1d, 0x62, 0xfc, 0xe1, 0x34, 0x87, 0x36, 0x0a, 0x27, 0x0a, 0x5a, 0x59, 0xa8, 0x45, 0x00,
	0x8b, 0xfd, 0xd5, 0x67, 0xe9, 0xc7, 0x38, 0x51, 0xf1, 0xde, 0x8f, 0x17, 0x41, 0x6c, 0xfc, 0x2f,
	0xa7, 0x65, 0x7f, 0x22, 0x84, 0x13, 0x85, 0xd0, 0x7e, 0x02, 0xa1, 0x45, 0x80, 0x90, 0x99, 0x79,
	0x3e, 0x31, 0xc5, 0x82, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x6b, 0xc4, 0x54, 0xe3, 0x01,
	0x00, 0x00,
}
