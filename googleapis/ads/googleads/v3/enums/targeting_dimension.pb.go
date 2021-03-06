// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/targeting_dimension.proto

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

// Enum describing possible targeting dimensions.
type TargetingDimensionEnum_TargetingDimension int32

const (
	// Not specified.
	TargetingDimensionEnum_UNSPECIFIED TargetingDimensionEnum_TargetingDimension = 0
	// Used for return value only. Represents value unknown in this version.
	TargetingDimensionEnum_UNKNOWN TargetingDimensionEnum_TargetingDimension = 1
	// Keyword criteria, e.g. 'mars cruise'. KEYWORD may be used as a custom bid
	// dimension. Keywords are always a targeting dimension, so may not be set
	// as a target "ALL" dimension with TargetRestriction.
	TargetingDimensionEnum_KEYWORD TargetingDimensionEnum_TargetingDimension = 2
	// Audience criteria, which include user list, user interest, custom
	// affinity,  and custom in market.
	TargetingDimensionEnum_AUDIENCE TargetingDimensionEnum_TargetingDimension = 3
	// Topic criteria for targeting categories of content, e.g.
	// 'category::Animals>Pets' Used for Display and Video targeting.
	TargetingDimensionEnum_TOPIC TargetingDimensionEnum_TargetingDimension = 4
	// Criteria for targeting gender.
	TargetingDimensionEnum_GENDER TargetingDimensionEnum_TargetingDimension = 5
	// Criteria for targeting age ranges.
	TargetingDimensionEnum_AGE_RANGE TargetingDimensionEnum_TargetingDimension = 6
	// Placement criteria, which include websites like 'www.flowers4sale.com',
	// as well as mobile applications, mobile app categories, YouTube videos,
	// and YouTube channels.
	TargetingDimensionEnum_PLACEMENT TargetingDimensionEnum_TargetingDimension = 7
	// Criteria for parental status targeting.
	TargetingDimensionEnum_PARENTAL_STATUS TargetingDimensionEnum_TargetingDimension = 8
	// Criteria for income range targeting.
	TargetingDimensionEnum_INCOME_RANGE TargetingDimensionEnum_TargetingDimension = 9
)

var TargetingDimensionEnum_TargetingDimension_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "KEYWORD",
	3: "AUDIENCE",
	4: "TOPIC",
	5: "GENDER",
	6: "AGE_RANGE",
	7: "PLACEMENT",
	8: "PARENTAL_STATUS",
	9: "INCOME_RANGE",
}

var TargetingDimensionEnum_TargetingDimension_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"KEYWORD":         2,
	"AUDIENCE":        3,
	"TOPIC":           4,
	"GENDER":          5,
	"AGE_RANGE":       6,
	"PLACEMENT":       7,
	"PARENTAL_STATUS": 8,
	"INCOME_RANGE":    9,
}

func (x TargetingDimensionEnum_TargetingDimension) String() string {
	return proto.EnumName(TargetingDimensionEnum_TargetingDimension_name, int32(x))
}

func (TargetingDimensionEnum_TargetingDimension) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e42413d270777352, []int{0, 0}
}

// The dimensions that can be targeted.
type TargetingDimensionEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetingDimensionEnum) Reset()         { *m = TargetingDimensionEnum{} }
func (m *TargetingDimensionEnum) String() string { return proto.CompactTextString(m) }
func (*TargetingDimensionEnum) ProtoMessage()    {}
func (*TargetingDimensionEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e42413d270777352, []int{0}
}

func (m *TargetingDimensionEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetingDimensionEnum.Unmarshal(m, b)
}
func (m *TargetingDimensionEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetingDimensionEnum.Marshal(b, m, deterministic)
}
func (m *TargetingDimensionEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetingDimensionEnum.Merge(m, src)
}
func (m *TargetingDimensionEnum) XXX_Size() int {
	return xxx_messageInfo_TargetingDimensionEnum.Size(m)
}
func (m *TargetingDimensionEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetingDimensionEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TargetingDimensionEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.TargetingDimensionEnum_TargetingDimension", TargetingDimensionEnum_TargetingDimension_name, TargetingDimensionEnum_TargetingDimension_value)
	proto.RegisterType((*TargetingDimensionEnum)(nil), "google.ads.googleads.v3.enums.TargetingDimensionEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/targeting_dimension.proto", fileDescriptor_e42413d270777352)
}

var fileDescriptor_e42413d270777352 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4f, 0xcb, 0xd3, 0x30,
	0x1c, 0xb6, 0x7d, 0x7d, 0xf7, 0x6e, 0xd9, 0x64, 0x21, 0x82, 0x82, 0xb8, 0xc3, 0xf6, 0x01, 0xd2,
	0x43, 0x0f, 0x42, 0x3c, 0x65, 0x6d, 0x2c, 0x65, 0x5b, 0x5a, 0xba, 0x76, 0x43, 0x29, 0x8c, 0x6a,
	0x4b, 0x28, 0xac, 0xc9, 0x58, 0xba, 0x7d, 0x20, 0x6f, 0xfa, 0x39, 0x3c, 0xf9, 0x51, 0xc4, 0x0f,
	0x21, 0x6d, 0xed, 0x2e, 0x43, 0x2f, 0xe1, 0xf9, 0xe5, 0xf9, 0x43, 0xf2, 0xfc, 0xc0, 0x3b, 0xa1,
	0x94, 0x38, 0x16, 0x56, 0x96, 0x6b, 0xab, 0x83, 0x0d, 0xba, 0xda, 0x56, 0x21, 0x2f, 0x95, 0xb6,
	0xea, 0xec, 0x2c, 0x8a, 0xba, 0x94, 0xe2, 0x90, 0x97, 0x55, 0x21, 0x75, 0xa9, 0x24, 0x3e, 0x9d,
	0x55, 0xad, 0xd0, 0xac, 0x53, 0xe3, 0x2c, 0xd7, 0xf8, 0x66, 0xc4, 0x57, 0x1b, 0xb7, 0xc6, 0x37,
	0x6f, 0xfb, 0xdc, 0x53, 0x69, 0x65, 0x52, 0xaa, 0x3a, 0xab, 0x4b, 0x25, 0x75, 0x67, 0x5e, 0xfc,
	0x30, 0xc0, 0xab, 0xb8, 0x8f, 0x76, 0xfb, 0x64, 0x26, 0x2f, 0xd5, 0xe2, 0x9b, 0x01, 0xd0, 0x3d,
	0x85, 0xa6, 0x60, 0x9c, 0xf0, 0x6d, 0xc8, 0x1c, 0xff, 0x83, 0xcf, 0x5c, 0xf8, 0x0c, 0x8d, 0xc1,
	0x53, 0xc2, 0x57, 0x3c, 0xd8, 0x73, 0x68, 0x34, 0xc3, 0x8a, 0x7d, 0xdc, 0x07, 0x91, 0x0b, 0x4d,
	0x34, 0x01, 0x43, 0x9a, 0xb8, 0x3e, 0xe3, 0x0e, 0x83, 0x0f, 0x68, 0x04, 0x1e, 0xe3, 0x20, 0xf4,
	0x1d, 0xf8, 0x1c, 0x01, 0x30, 0xf0, 0x18, 0x77, 0x59, 0x04, 0x1f, 0xd1, 0x0b, 0x30, 0xa2, 0x1e,
	0x3b, 0x44, 0x94, 0x7b, 0x0c, 0x0e, 0x9a, 0x31, 0x5c, 0x53, 0x87, 0x6d, 0x18, 0x8f, 0xe1, 0x13,
	0x7a, 0x09, 0xa6, 0x21, 0x8d, 0x18, 0x8f, 0xe9, 0xfa, 0xb0, 0x8d, 0x69, 0x9c, 0x6c, 0xe1, 0x10,
	0x41, 0x30, 0xf1, 0xb9, 0x13, 0x6c, 0x7a, 0xd7, 0x68, 0xf9, 0xdb, 0x00, 0xf3, 0x2f, 0xaa, 0xc2,
	0xff, 0xad, 0x62, 0xf9, 0xfa, 0xfe, 0x3b, 0x61, 0xd3, 0x42, 0x68, 0x7c, 0x5a, 0xfe, 0x75, 0x0a,
	0x75, 0xcc, 0xa4, 0xc0, 0xea, 0x2c, 0x2c, 0x51, 0xc8, 0xb6, 0xa3, 0x7e, 0x1b, 0xa7, 0x52, 0xff,
	0x63, 0x39, 0xef, 0xdb, 0xf3, 0xab, 0xf9, 0xe0, 0x51, 0xfa, 0xdd, 0x9c, 0x79, 0x5d, 0x14, 0xcd,
	0x35, 0xee, 0x60, 0x83, 0x76, 0x36, 0x6e, 0x5a, 0xd5, 0x3f, 0x7b, 0x3e, 0xa5, 0xb9, 0x4e, 0x6f,
	0x7c, 0xba, 0xb3, 0xd3, 0x96, 0xff, 0x65, 0xce, 0xbb, 0x4b, 0x42, 0x68, 0xae, 0x09, 0xb9, 0x29,
	0x08, 0xd9, 0xd9, 0x84, 0xb4, 0x9a, 0xcf, 0x83, 0xf6, 0x61, 0xf6, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x45, 0xef, 0xff, 0xbe, 0x34, 0x02, 0x00, 0x00,
}
