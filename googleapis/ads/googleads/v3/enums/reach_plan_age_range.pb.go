// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/reach_plan_age_range.proto

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

// Possible plannable age range values.
type ReachPlanAgeRangeEnum_ReachPlanAgeRange int32

const (
	// Not specified.
	ReachPlanAgeRangeEnum_UNSPECIFIED ReachPlanAgeRangeEnum_ReachPlanAgeRange = 0
	// The value is unknown in this version.
	ReachPlanAgeRangeEnum_UNKNOWN ReachPlanAgeRangeEnum_ReachPlanAgeRange = 1
	// Between 18 and 24 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_24 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503001
	// Between 18 and 34 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_34 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 2
	// Between 18 and 44 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_44 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 3
	// Between 18 and 49 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_49 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 4
	// Between 18 and 54 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_54 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 5
	// Between 18 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 6
	// Between 18 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_18_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 7
	// Between 21 and 34 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_21_34 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 8
	// Between 25 and 34 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_34 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503002
	// Between 25 and 44 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_44 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 9
	// Between 25 and 49 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_49 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 10
	// Between 25 and 54 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_54 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 11
	// Between 25 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 12
	// Between 25 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_25_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 13
	// Between 35 and 44 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_44 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503003
	// Between 35 and 49 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_49 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 14
	// Between 35 and 54 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_54 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 15
	// Between 35 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 16
	// Between 35 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_35_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 17
	// Between 45 and 54 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_45_54 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503004
	// Between 45 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_45_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 18
	// Between 45 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_45_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 19
	// Between 50 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_50_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 20
	// Between 55 and 64 years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_55_64 ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503005
	// Between 55 and 65+ years old.
	ReachPlanAgeRangeEnum_AGE_RANGE_55_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 21
	// 65 years old and beyond.
	ReachPlanAgeRangeEnum_AGE_RANGE_65_UP ReachPlanAgeRangeEnum_ReachPlanAgeRange = 503006
)

var ReachPlanAgeRangeEnum_ReachPlanAgeRange_name = map[int32]string{
	0:      "UNSPECIFIED",
	1:      "UNKNOWN",
	503001: "AGE_RANGE_18_24",
	2:      "AGE_RANGE_18_34",
	3:      "AGE_RANGE_18_44",
	4:      "AGE_RANGE_18_49",
	5:      "AGE_RANGE_18_54",
	6:      "AGE_RANGE_18_64",
	7:      "AGE_RANGE_18_65_UP",
	8:      "AGE_RANGE_21_34",
	503002: "AGE_RANGE_25_34",
	9:      "AGE_RANGE_25_44",
	10:     "AGE_RANGE_25_49",
	11:     "AGE_RANGE_25_54",
	12:     "AGE_RANGE_25_64",
	13:     "AGE_RANGE_25_65_UP",
	503003: "AGE_RANGE_35_44",
	14:     "AGE_RANGE_35_49",
	15:     "AGE_RANGE_35_54",
	16:     "AGE_RANGE_35_64",
	17:     "AGE_RANGE_35_65_UP",
	503004: "AGE_RANGE_45_54",
	18:     "AGE_RANGE_45_64",
	19:     "AGE_RANGE_45_65_UP",
	20:     "AGE_RANGE_50_65_UP",
	503005: "AGE_RANGE_55_64",
	21:     "AGE_RANGE_55_65_UP",
	503006: "AGE_RANGE_65_UP",
}

var ReachPlanAgeRangeEnum_ReachPlanAgeRange_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"AGE_RANGE_18_24":    503001,
	"AGE_RANGE_18_34":    2,
	"AGE_RANGE_18_44":    3,
	"AGE_RANGE_18_49":    4,
	"AGE_RANGE_18_54":    5,
	"AGE_RANGE_18_64":    6,
	"AGE_RANGE_18_65_UP": 7,
	"AGE_RANGE_21_34":    8,
	"AGE_RANGE_25_34":    503002,
	"AGE_RANGE_25_44":    9,
	"AGE_RANGE_25_49":    10,
	"AGE_RANGE_25_54":    11,
	"AGE_RANGE_25_64":    12,
	"AGE_RANGE_25_65_UP": 13,
	"AGE_RANGE_35_44":    503003,
	"AGE_RANGE_35_49":    14,
	"AGE_RANGE_35_54":    15,
	"AGE_RANGE_35_64":    16,
	"AGE_RANGE_35_65_UP": 17,
	"AGE_RANGE_45_54":    503004,
	"AGE_RANGE_45_64":    18,
	"AGE_RANGE_45_65_UP": 19,
	"AGE_RANGE_50_65_UP": 20,
	"AGE_RANGE_55_64":    503005,
	"AGE_RANGE_55_65_UP": 21,
	"AGE_RANGE_65_UP":    503006,
}

func (x ReachPlanAgeRangeEnum_ReachPlanAgeRange) String() string {
	return proto.EnumName(ReachPlanAgeRangeEnum_ReachPlanAgeRange_name, int32(x))
}

func (ReachPlanAgeRangeEnum_ReachPlanAgeRange) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aa9ffb3fe54db3d3, []int{0, 0}
}

// Message describing plannable age ranges.
type ReachPlanAgeRangeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReachPlanAgeRangeEnum) Reset()         { *m = ReachPlanAgeRangeEnum{} }
func (m *ReachPlanAgeRangeEnum) String() string { return proto.CompactTextString(m) }
func (*ReachPlanAgeRangeEnum) ProtoMessage()    {}
func (*ReachPlanAgeRangeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa9ffb3fe54db3d3, []int{0}
}

func (m *ReachPlanAgeRangeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReachPlanAgeRangeEnum.Unmarshal(m, b)
}
func (m *ReachPlanAgeRangeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReachPlanAgeRangeEnum.Marshal(b, m, deterministic)
}
func (m *ReachPlanAgeRangeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReachPlanAgeRangeEnum.Merge(m, src)
}
func (m *ReachPlanAgeRangeEnum) XXX_Size() int {
	return xxx_messageInfo_ReachPlanAgeRangeEnum.Size(m)
}
func (m *ReachPlanAgeRangeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ReachPlanAgeRangeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ReachPlanAgeRangeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ReachPlanAgeRangeEnum_ReachPlanAgeRange", ReachPlanAgeRangeEnum_ReachPlanAgeRange_name, ReachPlanAgeRangeEnum_ReachPlanAgeRange_value)
	proto.RegisterType((*ReachPlanAgeRangeEnum)(nil), "google.ads.googleads.v3.enums.ReachPlanAgeRangeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/reach_plan_age_range.proto", fileDescriptor_aa9ffb3fe54db3d3)
}

var fileDescriptor_aa9ffb3fe54db3d3 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x69, 0xf7, 0x0f, 0x5c, 0xa0, 0x5e, 0x46, 0x77, 0x81, 0x58, 0xa5, 0xed, 0x01, 0x1c,
	0x56, 0x3b, 0xd5, 0x16, 0xae, 0x5c, 0x08, 0xd5, 0x84, 0x14, 0xa2, 0xa2, 0x16, 0x09, 0x55, 0x8a,
	0xcc, 0x12, 0x99, 0x4a, 0xad, 0x5d, 0xd5, 0xdd, 0x5e, 0x82, 0xb7, 0xe0, 0x92, 0x47, 0xe1, 0x0d,
	0xc2, 0xbf, 0x07, 0x80, 0x2b, 0xde, 0x00, 0xd9, 0x69, 0x2a, 0x39, 0x16, 0xdc, 0x58, 0x47, 0xbf,
	0x73, 0xce, 0xf7, 0x7d, 0x17, 0xc7, 0xe0, 0x82, 0x4b, 0xc9, 0xe7, 0xb9, 0xcf, 0x32, 0xe5, 0x97,
	0xa5, 0xae, 0x6e, 0xb1, 0x9f, 0x8b, 0x9b, 0x85, 0xf2, 0x57, 0x39, 0xbb, 0xfe, 0x90, 0x2e, 0xe7,
	0x4c, 0xa4, 0x8c, 0xe7, 0xe9, 0x8a, 0x09, 0x9e, 0xa3, 0xe5, 0x4a, 0xae, 0xa5, 0x77, 0x52, 0x8e,
	0x23, 0x96, 0x29, 0xb4, 0xdd, 0x44, 0xb7, 0x18, 0x99, 0xcd, 0xc7, 0x4f, 0x2a, 0xe1, 0xe5, 0xcc,
	0x67, 0x42, 0xc8, 0x35, 0x5b, 0xcf, 0xa4, 0x50, 0xe5, 0xf2, 0xd9, 0xc7, 0x3d, 0xd0, 0x19, 0x69,
	0xed, 0x64, 0xce, 0x04, 0xe5, 0xf9, 0x48, 0x0b, 0x47, 0xe2, 0x66, 0x71, 0xf6, 0x67, 0x17, 0x1c,
	0x3a, 0x1d, 0xaf, 0x0d, 0x5a, 0xe3, 0xf8, 0x4d, 0x12, 0x3d, 0xbf, 0x7a, 0x79, 0x15, 0xbd, 0x80,
	0x77, 0xbc, 0x16, 0x38, 0x18, 0xc7, 0xaf, 0xe2, 0xd7, 0x6f, 0x63, 0xd8, 0xf0, 0x3a, 0xa0, 0x4d,
	0x87, 0x51, 0x3a, 0xa2, 0xf1, 0x30, 0x4a, 0xcf, 0x2f, 0xd2, 0x1e, 0x81, 0x45, 0xd1, 0xf5, 0x8e,
	0x6a, 0x18, 0x13, 0xd8, 0x74, 0x20, 0x21, 0x70, 0xc7, 0x85, 0x97, 0x70, 0xd7, 0x81, 0x01, 0x81,
	0x7b, 0x0e, 0xec, 0x13, 0xb8, 0xef, 0x1d, 0x03, 0xcf, 0x86, 0x41, 0x3a, 0x4e, 0xe0, 0x81, 0x3d,
	0xdc, 0x3b, 0xd7, 0x01, 0xee, 0xda, 0x61, 0x7b, 0x81, 0x86, 0x5f, 0xeb, 0x61, 0x7b, 0x81, 0xce,
	0x75, 0xcf, 0x85, 0x97, 0x10, 0x38, 0x30, 0x20, 0xb0, 0xe5, 0xc0, 0x3e, 0x81, 0xf7, 0xed, 0x5c,
	0x1a, 0x9a, 0x5c, 0x0f, 0xec, 0x08, 0xd8, 0x78, 0x7d, 0xab, 0x47, 0xc0, 0xc6, 0xed, 0xa1, 0x03,
	0x03, 0x02, 0xdb, 0x0e, 0xec, 0x13, 0x08, 0x6d, 0x37, 0x5c, 0xb9, 0x1d, 0xda, 0x6e, 0xc4, 0x28,
	0x7c, 0xaf, 0xbb, 0x11, 0xa3, 0xe1, 0xd9, 0x1a, 0xa4, 0xd2, 0x38, 0xb2, 0x79, 0xf0, 0x74, 0xc3,
	0x1f, 0xd9, 0xda, 0x81, 0x11, 0xf9, 0x51, 0x74, 0x6b, 0xe3, 0x95, 0x4c, 0xc7, 0x1e, 0x2f, 0xe1,
	0xcf, 0xa2, 0x3b, 0xf8, 0xdd, 0x00, 0xa7, 0xd7, 0x72, 0x81, 0xfe, 0x7b, 0xd1, 0x83, 0x63, 0xe7,
	0x2c, 0x13, 0x7d, 0xcb, 0x49, 0xe3, 0xdd, 0x60, 0xb3, 0xc8, 0xe5, 0x9c, 0x09, 0x8e, 0xe4, 0x8a,
	0xfb, 0x3c, 0x17, 0xe6, 0xd2, 0xab, 0x4f, 0xb5, 0x9c, 0xa9, 0x7f, 0xfc, 0xb1, 0x67, 0xe6, 0xfd,
	0xd4, 0xdc, 0x19, 0x52, 0xfa, 0xb9, 0x79, 0x32, 0x2c, 0xa5, 0x68, 0xa6, 0x50, 0x59, 0xea, 0x6a,
	0x82, 0x91, 0xfe, 0x1c, 0xea, 0x4b, 0xd5, 0x9f, 0xd2, 0x4c, 0x4d, 0xb7, 0xfd, 0xe9, 0x04, 0x4f,
	0x4d, 0xff, 0x57, 0xf3, 0xb4, 0x84, 0x61, 0x48, 0x33, 0x15, 0x86, 0xdb, 0x89, 0x30, 0x9c, 0xe0,
	0x30, 0x34, 0x33, 0xef, 0xf7, 0x4d, 0x30, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0x17, 0x4f, 0x5d,
	0x4f, 0xfb, 0x03, 0x00, 0x00,
}
