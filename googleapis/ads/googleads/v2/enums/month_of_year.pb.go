// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/month_of_year.proto

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

// Enumerates months of the year, e.g., "January".
type MonthOfYearEnum_MonthOfYear int32

const (
	// Not specified.
	MonthOfYearEnum_UNSPECIFIED MonthOfYearEnum_MonthOfYear = 0
	// The value is unknown in this version.
	MonthOfYearEnum_UNKNOWN MonthOfYearEnum_MonthOfYear = 1
	// January.
	MonthOfYearEnum_JANUARY MonthOfYearEnum_MonthOfYear = 2
	// February.
	MonthOfYearEnum_FEBRUARY MonthOfYearEnum_MonthOfYear = 3
	// March.
	MonthOfYearEnum_MARCH MonthOfYearEnum_MonthOfYear = 4
	// April.
	MonthOfYearEnum_APRIL MonthOfYearEnum_MonthOfYear = 5
	// May.
	MonthOfYearEnum_MAY MonthOfYearEnum_MonthOfYear = 6
	// June.
	MonthOfYearEnum_JUNE MonthOfYearEnum_MonthOfYear = 7
	// July.
	MonthOfYearEnum_JULY MonthOfYearEnum_MonthOfYear = 8
	// August.
	MonthOfYearEnum_AUGUST MonthOfYearEnum_MonthOfYear = 9
	// September.
	MonthOfYearEnum_SEPTEMBER MonthOfYearEnum_MonthOfYear = 10
	// October.
	MonthOfYearEnum_OCTOBER MonthOfYearEnum_MonthOfYear = 11
	// November.
	MonthOfYearEnum_NOVEMBER MonthOfYearEnum_MonthOfYear = 12
	// December.
	MonthOfYearEnum_DECEMBER MonthOfYearEnum_MonthOfYear = 13
)

var MonthOfYearEnum_MonthOfYear_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "JANUARY",
	3:  "FEBRUARY",
	4:  "MARCH",
	5:  "APRIL",
	6:  "MAY",
	7:  "JUNE",
	8:  "JULY",
	9:  "AUGUST",
	10: "SEPTEMBER",
	11: "OCTOBER",
	12: "NOVEMBER",
	13: "DECEMBER",
}

var MonthOfYearEnum_MonthOfYear_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"JANUARY":     2,
	"FEBRUARY":    3,
	"MARCH":       4,
	"APRIL":       5,
	"MAY":         6,
	"JUNE":        7,
	"JULY":        8,
	"AUGUST":      9,
	"SEPTEMBER":   10,
	"OCTOBER":     11,
	"NOVEMBER":    12,
	"DECEMBER":    13,
}

func (x MonthOfYearEnum_MonthOfYear) String() string {
	return proto.EnumName(MonthOfYearEnum_MonthOfYear_name, int32(x))
}

func (MonthOfYearEnum_MonthOfYear) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e16e44aeadbe2fc9, []int{0, 0}
}

// Container for enumeration of months of the year, e.g., "January".
type MonthOfYearEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonthOfYearEnum) Reset()         { *m = MonthOfYearEnum{} }
func (m *MonthOfYearEnum) String() string { return proto.CompactTextString(m) }
func (*MonthOfYearEnum) ProtoMessage()    {}
func (*MonthOfYearEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16e44aeadbe2fc9, []int{0}
}

func (m *MonthOfYearEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonthOfYearEnum.Unmarshal(m, b)
}
func (m *MonthOfYearEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonthOfYearEnum.Marshal(b, m, deterministic)
}
func (m *MonthOfYearEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonthOfYearEnum.Merge(m, src)
}
func (m *MonthOfYearEnum) XXX_Size() int {
	return xxx_messageInfo_MonthOfYearEnum.Size(m)
}
func (m *MonthOfYearEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MonthOfYearEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MonthOfYearEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.MonthOfYearEnum_MonthOfYear", MonthOfYearEnum_MonthOfYear_name, MonthOfYearEnum_MonthOfYear_value)
	proto.RegisterType((*MonthOfYearEnum)(nil), "google.ads.googleads.v2.enums.MonthOfYearEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/month_of_year.proto", fileDescriptor_e16e44aeadbe2fc9)
}

var fileDescriptor_e16e44aeadbe2fc9 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x25, 0xe9, 0x4c, 0x1f, 0xce, 0x8c, 0xc6, 0xf2, 0x12, 0x31, 0x8b, 0x99, 0x0f, 0x70, 0x44,
	0xd8, 0x99, 0x95, 0x93, 0xf1, 0x94, 0x0e, 0xcd, 0x43, 0x69, 0x13, 0x14, 0x14, 0xa9, 0x0a, 0x24,
	0x0d, 0x95, 0x1a, 0xbb, 0x8a, 0xd3, 0x4a, 0xfc, 0x0e, 0x4b, 0xfe, 0x81, 0x1f, 0x60, 0xc7, 0x47,
	0xb0, 0xe1, 0x2b, 0x90, 0x63, 0x5a, 0x75, 0xc3, 0x6c, 0xa2, 0x73, 0xee, 0xb9, 0xe7, 0xde, 0xf8,
	0x5c, 0xf0, 0xba, 0x16, 0xa2, 0xde, 0x56, 0x76, 0x51, 0x4a, 0x5b, 0x43, 0x85, 0x0e, 0x8e, 0x5d,
	0xf1, 0x7d, 0x23, 0xed, 0x46, 0xf0, 0xee, 0xcb, 0x4a, 0xac, 0x57, 0x5f, 0xab, 0xa2, 0xc5, 0xbb,
	0x56, 0x74, 0x02, 0xdd, 0xea, 0x3e, 0x5c, 0x94, 0x12, 0x9f, 0x2c, 0xf8, 0xe0, 0xe0, 0xde, 0xf2,
	0xf2, 0xd5, 0x71, 0xe2, 0x6e, 0x63, 0x17, 0x9c, 0x8b, 0xae, 0xe8, 0x36, 0x82, 0x4b, 0x6d, 0xbe,
	0xff, 0x65, 0x80, 0x1b, 0x5f, 0x0d, 0x0d, 0xd7, 0x59, 0x55, 0xb4, 0x8c, 0xef, 0x9b, 0xfb, 0x1f,
	0x06, 0xb0, 0xce, 0x6a, 0xe8, 0x06, 0x58, 0x49, 0xb0, 0x88, 0x98, 0x37, 0x7b, 0x9c, 0xb1, 0x07,
	0xf8, 0x02, 0x59, 0x60, 0x94, 0x04, 0xef, 0x83, 0xf0, 0x43, 0x00, 0x0d, 0x45, 0x9e, 0x68, 0x90,
	0xd0, 0x38, 0x83, 0x26, 0xba, 0x02, 0xe3, 0x47, 0xe6, 0xc6, 0x3d, 0x1b, 0xa0, 0x09, 0xb8, 0xf4,
	0x69, 0xec, 0xbd, 0x83, 0x17, 0x0a, 0xd2, 0x28, 0x9e, 0xcd, 0xe1, 0x25, 0x1a, 0x81, 0x81, 0x4f,
	0x33, 0x38, 0x44, 0x63, 0x70, 0xf1, 0x94, 0x04, 0x0c, 0x8e, 0x34, 0x9a, 0x67, 0x70, 0x8c, 0x00,
	0x18, 0xd2, 0x64, 0x9a, 0x2c, 0x96, 0x70, 0x82, 0xae, 0xc1, 0x64, 0xc1, 0xa2, 0x25, 0xf3, 0x5d,
	0x16, 0x43, 0xa0, 0x16, 0x85, 0xde, 0x32, 0x54, 0xc4, 0x52, 0x8b, 0x82, 0x30, 0xd5, 0xd2, 0x95,
	0x62, 0x0f, 0xcc, 0xd3, 0xec, 0xda, 0xfd, 0x6d, 0x80, 0xbb, 0xcf, 0xa2, 0xc1, 0xcf, 0xe6, 0xe2,
	0xc2, 0xb3, 0x27, 0x46, 0x2a, 0x8b, 0xc8, 0xf8, 0xe8, 0xfe, 0xb3, 0xd4, 0x62, 0x5b, 0xf0, 0x1a,
	0x8b, 0xb6, 0xb6, 0xeb, 0x8a, 0xf7, 0x49, 0x1d, 0xaf, 0xb1, 0xdb, 0xc8, 0xff, 0x1c, 0xe7, 0x6d,
	0xff, 0xfd, 0x66, 0x0e, 0xa6, 0x94, 0x7e, 0x37, 0x6f, 0xa7, 0x7a, 0x14, 0x2d, 0x25, 0xd6, 0x50,
	0xa1, 0xd4, 0xc1, 0x2a, 0x62, 0xf9, 0xf3, 0xa8, 0xe7, 0xb4, 0x94, 0xf9, 0x49, 0xcf, 0x53, 0x27,
	0xef, 0xf5, 0x3f, 0xe6, 0x9d, 0x2e, 0x12, 0x42, 0x4b, 0x49, 0xc8, 0xa9, 0x83, 0x90, 0xd4, 0x21,
	0xa4, 0xef, 0xf9, 0x34, 0xec, 0x7f, 0xec, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x04,
	0xe3, 0xe0, 0x34, 0x02, 0x00, 0x00,
}
