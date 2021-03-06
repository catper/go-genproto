// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/minute_of_hour.proto

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

// Enumerates of quarter-hours. E.g. "FIFTEEN"
type MinuteOfHourEnum_MinuteOfHour int32

const (
	// Not specified.
	MinuteOfHourEnum_UNSPECIFIED MinuteOfHourEnum_MinuteOfHour = 0
	// The value is unknown in this version.
	MinuteOfHourEnum_UNKNOWN MinuteOfHourEnum_MinuteOfHour = 1
	// Zero minutes past the hour.
	MinuteOfHourEnum_ZERO MinuteOfHourEnum_MinuteOfHour = 2
	// Fifteen minutes past the hour.
	MinuteOfHourEnum_FIFTEEN MinuteOfHourEnum_MinuteOfHour = 3
	// Thirty minutes past the hour.
	MinuteOfHourEnum_THIRTY MinuteOfHourEnum_MinuteOfHour = 4
	// Forty-five minutes past the hour.
	MinuteOfHourEnum_FORTY_FIVE MinuteOfHourEnum_MinuteOfHour = 5
)

var MinuteOfHourEnum_MinuteOfHour_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ZERO",
	3: "FIFTEEN",
	4: "THIRTY",
	5: "FORTY_FIVE",
}

var MinuteOfHourEnum_MinuteOfHour_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ZERO":        2,
	"FIFTEEN":     3,
	"THIRTY":      4,
	"FORTY_FIVE":  5,
}

func (x MinuteOfHourEnum_MinuteOfHour) String() string {
	return proto.EnumName(MinuteOfHourEnum_MinuteOfHour_name, int32(x))
}

func (MinuteOfHourEnum_MinuteOfHour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d5c875df57002631, []int{0, 0}
}

// Container for enumeration of quarter-hours.
type MinuteOfHourEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MinuteOfHourEnum) Reset()         { *m = MinuteOfHourEnum{} }
func (m *MinuteOfHourEnum) String() string { return proto.CompactTextString(m) }
func (*MinuteOfHourEnum) ProtoMessage()    {}
func (*MinuteOfHourEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c875df57002631, []int{0}
}

func (m *MinuteOfHourEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MinuteOfHourEnum.Unmarshal(m, b)
}
func (m *MinuteOfHourEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MinuteOfHourEnum.Marshal(b, m, deterministic)
}
func (m *MinuteOfHourEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinuteOfHourEnum.Merge(m, src)
}
func (m *MinuteOfHourEnum) XXX_Size() int {
	return xxx_messageInfo_MinuteOfHourEnum.Size(m)
}
func (m *MinuteOfHourEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MinuteOfHourEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MinuteOfHourEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.MinuteOfHourEnum_MinuteOfHour", MinuteOfHourEnum_MinuteOfHour_name, MinuteOfHourEnum_MinuteOfHour_value)
	proto.RegisterType((*MinuteOfHourEnum)(nil), "google.ads.googleads.v2.enums.MinuteOfHourEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/minute_of_hour.proto", fileDescriptor_d5c875df57002631)
}

var fileDescriptor_d5c875df57002631 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x4e, 0xc2, 0x40,
	0x18, 0x95, 0x82, 0x68, 0x06, 0xa3, 0x63, 0x97, 0x46, 0x16, 0x70, 0x80, 0x69, 0x52, 0x77, 0xe3,
	0xaa, 0xe8, 0x14, 0x1a, 0x63, 0x4b, 0xb0, 0xd4, 0x40, 0x9a, 0x34, 0xd5, 0x96, 0xda, 0x84, 0xce,
	0x47, 0x3a, 0x2d, 0x07, 0x72, 0xe9, 0x51, 0x3c, 0x88, 0x0b, 0x4f, 0x61, 0x3a, 0x23, 0x84, 0x8d,
	0x6e, 0x26, 0x2f, 0xdf, 0xfb, 0xc9, 0x9b, 0x87, 0xcc, 0x0c, 0x20, 0x5b, 0xa7, 0x46, 0x9c, 0x08,
	0x43, 0xc1, 0x06, 0x6d, 0x4d, 0x23, 0xe5, 0x75, 0x21, 0x8c, 0x22, 0xe7, 0x75, 0x95, 0x46, 0xb0,
	0x8a, 0xde, 0xa0, 0x2e, 0xc9, 0xa6, 0x84, 0x0a, 0xf4, 0xbe, 0x12, 0x92, 0x38, 0x11, 0x64, 0xef,
	0x21, 0x5b, 0x93, 0x48, 0xcf, 0xd5, 0xf5, 0x2e, 0x72, 0x93, 0x1b, 0x31, 0xe7, 0x50, 0xc5, 0x55,
	0x0e, 0x5c, 0x28, 0xf3, 0x50, 0x20, 0xfc, 0x28, 0x43, 0xbd, 0xd5, 0x04, 0xea, 0x92, 0xf1, 0xba,
	0x18, 0x46, 0xe8, 0xec, 0xf0, 0xa6, 0x5f, 0xa0, 0xde, 0xdc, 0x7d, 0x9a, 0xb2, 0x3b, 0xc7, 0x76,
	0xd8, 0x3d, 0x3e, 0xd2, 0x7b, 0xe8, 0x64, 0xee, 0x3e, 0xb8, 0xde, 0xb3, 0x8b, 0x5b, 0xfa, 0x29,
	0xea, 0x2c, 0xd9, 0xcc, 0xc3, 0x5a, 0x73, 0xb6, 0x1d, 0xdb, 0x67, 0xcc, 0xc5, 0x6d, 0x1d, 0xa1,
	0xae, 0x3f, 0x71, 0x66, 0xfe, 0x02, 0x77, 0xf4, 0x73, 0x84, 0x6c, 0x6f, 0xe6, 0x2f, 0x22, 0xdb,
	0x09, 0x18, 0x3e, 0x1e, 0x7d, 0xb5, 0xd0, 0xe0, 0x15, 0x0a, 0xf2, 0x6f, 0xf1, 0xd1, 0xe5, 0x61,
	0x89, 0x69, 0xd3, 0x76, 0xda, 0x5a, 0x8e, 0x7e, 0x3d, 0x19, 0xac, 0x63, 0x9e, 0x11, 0x28, 0x33,
	0x23, 0x4b, 0xb9, 0xfc, 0xcb, 0x6e, 0xb0, 0x4d, 0x2e, 0xfe, 0xd8, 0xef, 0x56, 0xbe, 0xef, 0x5a,
	0x7b, 0x6c, 0x59, 0x1f, 0x5a, 0x7f, 0xac, 0xa2, 0xac, 0x44, 0x10, 0x05, 0x1b, 0x14, 0x98, 0xa4,
	0x19, 0x41, 0x7c, 0xee, 0xf8, 0xd0, 0x4a, 0x44, 0xb8, 0xe7, 0xc3, 0xc0, 0x0c, 0x25, 0xff, 0xad,
	0x0d, 0xd4, 0x91, 0x52, 0x2b, 0x11, 0x94, 0xee, 0x15, 0x94, 0x06, 0x26, 0xa5, 0x52, 0xf3, 0xd2,
	0x95, 0xc5, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0x0b, 0x2d, 0xe3, 0xd7, 0x01, 0x00,
	0x00,
}
