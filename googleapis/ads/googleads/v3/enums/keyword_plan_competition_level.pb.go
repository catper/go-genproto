// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/keyword_plan_competition_level.proto

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

// Competition level of a keyword.
type KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel int32

const (
	// Not specified.
	KeywordPlanCompetitionLevelEnum_UNSPECIFIED KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 0
	// The value is unknown in this version.
	KeywordPlanCompetitionLevelEnum_UNKNOWN KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 1
	// Low competition.
	KeywordPlanCompetitionLevelEnum_LOW KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 2
	// Medium competition.
	KeywordPlanCompetitionLevelEnum_MEDIUM KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 3
	// High competition.
	KeywordPlanCompetitionLevelEnum_HIGH KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 4
)

var KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "LOW",
	3: "MEDIUM",
	4: "HIGH",
}

var KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"LOW":         2,
	"MEDIUM":      3,
	"HIGH":        4,
}

func (x KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel) String() string {
	return proto.EnumName(KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_name, int32(x))
}

func (KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aaf960a010cf6f64, []int{0, 0}
}

// Container for enumeration of keyword competition levels. The competition
// level indicates how competitive ad placement is for a keyword and
// is determined by the number of advertisers bidding on that keyword relative
// to all keywords across Google. The competition level can depend on the
// location and Search Network targeting options you've selected.
type KeywordPlanCompetitionLevelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanCompetitionLevelEnum) Reset()         { *m = KeywordPlanCompetitionLevelEnum{} }
func (m *KeywordPlanCompetitionLevelEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCompetitionLevelEnum) ProtoMessage()    {}
func (*KeywordPlanCompetitionLevelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf960a010cf6f64, []int{0}
}

func (m *KeywordPlanCompetitionLevelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCompetitionLevelEnum.Unmarshal(m, b)
}
func (m *KeywordPlanCompetitionLevelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCompetitionLevelEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanCompetitionLevelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCompetitionLevelEnum.Merge(m, src)
}
func (m *KeywordPlanCompetitionLevelEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCompetitionLevelEnum.Size(m)
}
func (m *KeywordPlanCompetitionLevelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCompetitionLevelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCompetitionLevelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel", KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_name, KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_value)
	proto.RegisterType((*KeywordPlanCompetitionLevelEnum)(nil), "google.ads.googleads.v3.enums.KeywordPlanCompetitionLevelEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/keyword_plan_competition_level.proto", fileDescriptor_aaf960a010cf6f64)
}

var fileDescriptor_aaf960a010cf6f64 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xb1, 0x6e, 0xf2, 0x30,
	0x18, 0xfc, 0x09, 0x08, 0x7e, 0x99, 0xa1, 0x51, 0xc6, 0xb6, 0xa8, 0x85, 0x07, 0x70, 0x86, 0x6c,
	0xee, 0x14, 0x20, 0x85, 0x08, 0x08, 0x91, 0x2a, 0x40, 0x42, 0x91, 0x90, 0x8b, 0x2d, 0x2b, 0xaa,
	0x63, 0x47, 0x38, 0x50, 0x75, 0xe8, 0xcb, 0x74, 0xec, 0xa3, 0xf4, 0x51, 0xba, 0x77, 0xaf, 0x6c,
	0x17, 0x3a, 0x95, 0x25, 0x3a, 0xe5, 0xee, 0xbb, 0xf3, 0x1d, 0xe8, 0x33, 0x29, 0x19, 0xa7, 0x3e,
	0x26, 0xca, 0xb7, 0x50, 0xa3, 0x43, 0xe0, 0x53, 0xb1, 0x2f, 0x94, 0xff, 0x44, 0x5f, 0x9e, 0xe5,
	0x8e, 0x6c, 0x4a, 0x8e, 0xc5, 0x66, 0x2b, 0x8b, 0x92, 0x56, 0x79, 0x95, 0x4b, 0xb1, 0xe1, 0xf4,
	0x40, 0x39, 0x2c, 0x77, 0xb2, 0x92, 0x5e, 0xc7, 0x1e, 0x42, 0x4c, 0x14, 0x3c, 0x79, 0xc0, 0x43,
	0x00, 0x8d, 0xc7, 0xe5, 0xf5, 0x31, 0xa2, 0xcc, 0x7d, 0x2c, 0x84, 0xac, 0xb0, 0x76, 0x50, 0xf6,
	0xb8, 0xf7, 0x0a, 0x6e, 0x26, 0x36, 0x24, 0xe5, 0x58, 0x0c, 0x7e, 0x23, 0xa6, 0x3a, 0x21, 0x12,
	0xfb, 0xa2, 0xb7, 0x06, 0x57, 0x67, 0x24, 0xde, 0x05, 0x68, 0x2f, 0x92, 0x87, 0x34, 0x1a, 0xc4,
	0xf7, 0x71, 0x34, 0x74, 0xff, 0x79, 0x6d, 0xd0, 0x5a, 0x24, 0x93, 0x64, 0xbe, 0x4a, 0xdc, 0x9a,
	0xd7, 0x02, 0xf5, 0xe9, 0x7c, 0xe5, 0x3a, 0x1e, 0x00, 0xcd, 0x59, 0x34, 0x8c, 0x17, 0x33, 0xb7,
	0xee, 0xfd, 0x07, 0x8d, 0x71, 0x3c, 0x1a, 0xbb, 0x8d, 0xfe, 0x57, 0x0d, 0x74, 0xb7, 0xb2, 0x80,
	0x67, 0x2b, 0xf4, 0x6f, 0xcf, 0xe4, 0xa7, 0xba, 0x46, 0x5a, 0x5b, 0xff, 0x2c, 0x09, 0x99, 0xe4,
	0x58, 0x30, 0x28, 0x77, 0xcc, 0x67, 0x54, 0x98, 0x92, 0xc7, 0x65, 0xcb, 0x5c, 0xfd, 0x31, 0xf4,
	0x9d, 0xf9, 0xbe, 0x39, 0xf5, 0x51, 0x18, 0xbe, 0x3b, 0x9d, 0x91, 0xb5, 0x0a, 0x89, 0x82, 0x16,
	0x6a, 0xb4, 0x0c, 0xa0, 0x9e, 0x43, 0x7d, 0x1c, 0xf9, 0x2c, 0x24, 0x2a, 0x3b, 0xf1, 0xd9, 0x32,
	0xc8, 0x0c, 0xff, 0xe9, 0x74, 0xed, 0x4f, 0x84, 0x42, 0xa2, 0x10, 0x3a, 0x29, 0x10, 0x5a, 0x06,
	0x08, 0x19, 0xcd, 0x63, 0xd3, 0x3c, 0x2c, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x87, 0x6a,
	0xf6, 0x00, 0x02, 0x00, 0x00,
}
