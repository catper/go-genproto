// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/day_of_week.proto

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

// Enumerates days of the week, e.g., "Monday".
type DayOfWeekEnum_DayOfWeek int32

const (
	// Not specified.
	DayOfWeekEnum_UNSPECIFIED DayOfWeekEnum_DayOfWeek = 0
	// The value is unknown in this version.
	DayOfWeekEnum_UNKNOWN DayOfWeekEnum_DayOfWeek = 1
	// Monday.
	DayOfWeekEnum_MONDAY DayOfWeekEnum_DayOfWeek = 2
	// Tuesday.
	DayOfWeekEnum_TUESDAY DayOfWeekEnum_DayOfWeek = 3
	// Wednesday.
	DayOfWeekEnum_WEDNESDAY DayOfWeekEnum_DayOfWeek = 4
	// Thursday.
	DayOfWeekEnum_THURSDAY DayOfWeekEnum_DayOfWeek = 5
	// Friday.
	DayOfWeekEnum_FRIDAY DayOfWeekEnum_DayOfWeek = 6
	// Saturday.
	DayOfWeekEnum_SATURDAY DayOfWeekEnum_DayOfWeek = 7
	// Sunday.
	DayOfWeekEnum_SUNDAY DayOfWeekEnum_DayOfWeek = 8
)

var DayOfWeekEnum_DayOfWeek_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MONDAY",
	3: "TUESDAY",
	4: "WEDNESDAY",
	5: "THURSDAY",
	6: "FRIDAY",
	7: "SATURDAY",
	8: "SUNDAY",
}

var DayOfWeekEnum_DayOfWeek_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MONDAY":      2,
	"TUESDAY":     3,
	"WEDNESDAY":   4,
	"THURSDAY":    5,
	"FRIDAY":      6,
	"SATURDAY":    7,
	"SUNDAY":      8,
}

func (x DayOfWeekEnum_DayOfWeek) String() string {
	return proto.EnumName(DayOfWeekEnum_DayOfWeek_name, int32(x))
}

func (DayOfWeekEnum_DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_26c427a973cf03ad, []int{0, 0}
}

// Container for enumeration of days of the week, e.g., "Monday".
type DayOfWeekEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DayOfWeekEnum) Reset()         { *m = DayOfWeekEnum{} }
func (m *DayOfWeekEnum) String() string { return proto.CompactTextString(m) }
func (*DayOfWeekEnum) ProtoMessage()    {}
func (*DayOfWeekEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_26c427a973cf03ad, []int{0}
}

func (m *DayOfWeekEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DayOfWeekEnum.Unmarshal(m, b)
}
func (m *DayOfWeekEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DayOfWeekEnum.Marshal(b, m, deterministic)
}
func (m *DayOfWeekEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DayOfWeekEnum.Merge(m, src)
}
func (m *DayOfWeekEnum) XXX_Size() int {
	return xxx_messageInfo_DayOfWeekEnum.Size(m)
}
func (m *DayOfWeekEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DayOfWeekEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DayOfWeekEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.DayOfWeekEnum_DayOfWeek", DayOfWeekEnum_DayOfWeek_name, DayOfWeekEnum_DayOfWeek_value)
	proto.RegisterType((*DayOfWeekEnum)(nil), "google.ads.googleads.v2.enums.DayOfWeekEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/day_of_week.proto", fileDescriptor_26c427a973cf03ad)
}

var fileDescriptor_26c427a973cf03ad = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0xbf, 0x19, 0x3e, 0xf9, 0x53, 0x44, 0xc9, 0x2c, 0x8d, 0x2c, 0xe0, 0x01, 0x3a, 0xc9,
	0xb8, 0xab, 0xab, 0xe2, 0x0c, 0x48, 0x8c, 0x03, 0x61, 0x18, 0x88, 0x66, 0x12, 0x52, 0x9d, 0xd2,
	0x10, 0xa0, 0x25, 0x14, 0x30, 0x3c, 0x80, 0xcf, 0xe0, 0xde, 0xa5, 0x8f, 0xe2, 0x6b, 0xb8, 0xf3,
	0x29, 0x4c, 0x5b, 0x66, 0x76, 0xba, 0x69, 0xce, 0xe9, 0xef, 0x9e, 0x9b, 0x7b, 0x2f, 0x70, 0x99,
	0x10, 0x6c, 0x49, 0x5d, 0x92, 0xca, 0xa3, 0x54, 0x6a, 0xef, 0xb9, 0x94, 0xef, 0x56, 0xd2, 0x4d,
	0xc9, 0x61, 0x2a, 0x66, 0xd3, 0x17, 0x4a, 0x17, 0x70, 0xbd, 0x11, 0x5b, 0xe1, 0x34, 0x4c, 0x15,
	0x24, 0xa9, 0x84, 0x79, 0x00, 0xee, 0x3d, 0xa8, 0x03, 0x17, 0x97, 0x59, 0xbf, 0xf5, 0xdc, 0x25,
	0x9c, 0x8b, 0x2d, 0xd9, 0xce, 0x05, 0x97, 0x26, 0xdc, 0x7a, 0xb3, 0x40, 0xcd, 0x27, 0x87, 0xfe,
	0x6c, 0x42, 0xe9, 0x22, 0xe0, 0xbb, 0x55, 0xeb, 0xd5, 0x02, 0x95, 0xfc, 0xc7, 0x39, 0x07, 0xd5,
	0x38, 0x8c, 0x06, 0xc1, 0x4d, 0xaf, 0xd3, 0x0b, 0xfc, 0xfa, 0x3f, 0xa7, 0x0a, 0x4a, 0x71, 0x78,
	0x17, 0xf6, 0x27, 0x61, 0xdd, 0x72, 0x00, 0x28, 0xde, 0xf7, 0x43, 0x1f, 0x3f, 0xd4, 0x6d, 0x05,
	0x46, 0x71, 0x10, 0x29, 0x53, 0x70, 0x6a, 0xa0, 0x32, 0x09, 0xfc, 0xd0, 0xd8, 0xff, 0xce, 0x29,
	0x28, 0x8f, 0x6e, 0xe3, 0xa1, 0x76, 0x27, 0x2a, 0xd5, 0x19, 0xf6, 0x94, 0x2e, 0x2a, 0x12, 0xe1,
	0x51, 0x3c, 0x54, 0xae, 0xa4, 0x48, 0x14, 0xeb, 0x7e, 0xe5, 0xf6, 0x97, 0x05, 0x9a, 0xcf, 0x62,
	0x05, 0xff, 0xdc, 0xae, 0x7d, 0x96, 0x8f, 0x3a, 0x50, 0xfb, 0x0c, 0xac, 0xc7, 0xf6, 0x31, 0xc0,
	0xc4, 0x92, 0x70, 0x06, 0xc5, 0x86, 0xb9, 0x8c, 0x72, 0xbd, 0x6d, 0x76, 0xcf, 0xf5, 0x5c, 0xfe,
	0x72, 0xde, 0x6b, 0xfd, 0xbe, 0xdb, 0x85, 0x2e, 0xc6, 0x1f, 0x76, 0xa3, 0x6b, 0x5a, 0xe1, 0x54,
	0x42, 0x23, 0x95, 0x1a, 0x7b, 0x50, 0x1d, 0x4a, 0x7e, 0x66, 0x3c, 0xc1, 0xa9, 0x4c, 0x72, 0x9e,
	0x8c, 0xbd, 0x44, 0xf3, 0x6f, 0xbb, 0x69, 0x3e, 0x11, 0xc2, 0xa9, 0x44, 0x28, 0xaf, 0x40, 0x68,
	0xec, 0x21, 0xa4, 0x6b, 0x9e, 0x8a, 0x7a, 0xb0, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f,
	0xaf, 0x68, 0x22, 0xf6, 0x01, 0x00, 0x00,
}
