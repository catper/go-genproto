// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/minute_of_hour.proto

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
	return fileDescriptor_8fd8f371ee9dca04, []int{0, 0}
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
	return fileDescriptor_8fd8f371ee9dca04, []int{0}
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
	proto.RegisterEnum("google.ads.googleads.v1.enums.MinuteOfHourEnum_MinuteOfHour", MinuteOfHourEnum_MinuteOfHour_name, MinuteOfHourEnum_MinuteOfHour_value)
	proto.RegisterType((*MinuteOfHourEnum)(nil), "google.ads.googleads.v1.enums.MinuteOfHourEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/minute_of_hour.proto", fileDescriptor_8fd8f371ee9dca04)
}

var fileDescriptor_8fd8f371ee9dca04 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4e, 0xf2, 0x40,
	0x1c, 0xfd, 0x28, 0x7c, 0x68, 0x06, 0xa3, 0x63, 0x97, 0x46, 0x16, 0x70, 0x80, 0x69, 0xaa, 0xbb,
	0x71, 0x55, 0x74, 0x0a, 0x8d, 0xb1, 0x25, 0x58, 0x6a, 0x20, 0x4d, 0x9a, 0x6a, 0x4b, 0x6d, 0x42,
	0xe7, 0x47, 0x3a, 0x2d, 0x07, 0x72, 0xe9, 0x51, 0x3c, 0x88, 0x0b, 0x4f, 0x61, 0x3a, 0x23, 0x84,
	0x8d, 0x6e, 0x26, 0x2f, 0xbf, 0xf7, 0x27, 0x6f, 0x1e, 0xba, 0xca, 0x00, 0xb2, 0x75, 0x6a, 0xc4,
	0x89, 0x30, 0x14, 0x6c, 0xd0, 0xd6, 0x34, 0x52, 0x5e, 0x17, 0xc2, 0x28, 0x72, 0x5e, 0x57, 0x69,
	0x04, 0xab, 0xe8, 0x15, 0xea, 0x92, 0x6c, 0x4a, 0xa8, 0x40, 0xef, 0x2b, 0x21, 0x89, 0x13, 0x41,
	0xf6, 0x1e, 0xb2, 0x35, 0x89, 0xf4, 0x5c, 0x5c, 0xee, 0x22, 0x37, 0xb9, 0x11, 0x73, 0x0e, 0x55,
	0x5c, 0xe5, 0xc0, 0x85, 0x32, 0x0f, 0x05, 0xc2, 0x0f, 0x32, 0xd4, 0x5b, 0x4d, 0xa0, 0x2e, 0x19,
	0xaf, 0x8b, 0x61, 0x84, 0x4e, 0x0e, 0x6f, 0xfa, 0x19, 0xea, 0xcd, 0xdd, 0xc7, 0x29, 0xbb, 0x75,
	0x6c, 0x87, 0xdd, 0xe1, 0x7f, 0x7a, 0x0f, 0x1d, 0xcd, 0xdd, 0x7b, 0xd7, 0x7b, 0x72, 0x71, 0x4b,
	0x3f, 0x46, 0x9d, 0x25, 0x9b, 0x79, 0x58, 0x6b, 0xce, 0xb6, 0x63, 0xfb, 0x8c, 0xb9, 0xb8, 0xad,
	0x23, 0xd4, 0xf5, 0x27, 0xce, 0xcc, 0x5f, 0xe0, 0x8e, 0x7e, 0x8a, 0x90, 0xed, 0xcd, 0xfc, 0x45,
	0x64, 0x3b, 0x01, 0xc3, 0xff, 0x47, 0x9f, 0x2d, 0x34, 0x78, 0x81, 0x82, 0xfc, 0x59, 0x7c, 0x74,
	0x7e, 0x58, 0x62, 0xda, 0xb4, 0x9d, 0xb6, 0x96, 0xa3, 0x1f, 0x4f, 0x06, 0xeb, 0x98, 0x67, 0x04,
	0xca, 0xcc, 0xc8, 0x52, 0x2e, 0xff, 0xb2, 0x1b, 0x6c, 0x93, 0x8b, 0x5f, 0xf6, 0xbb, 0x91, 0xef,
	0x9b, 0xd6, 0x1e, 0x5b, 0xd6, 0xbb, 0xd6, 0x1f, 0xab, 0x28, 0x2b, 0x11, 0x44, 0xc1, 0x06, 0x05,
	0x26, 0x69, 0x46, 0x10, 0x1f, 0x3b, 0x3e, 0xb4, 0x12, 0x11, 0xee, 0xf9, 0x30, 0x30, 0x43, 0xc9,
	0x7f, 0x69, 0x03, 0x75, 0xa4, 0xd4, 0x4a, 0x04, 0xa5, 0x7b, 0x05, 0xa5, 0x81, 0x49, 0xa9, 0xd4,
	0x3c, 0x77, 0x65, 0xb1, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xe0, 0x63, 0x00, 0xd7,
	0x01, 0x00, 0x00,
}
