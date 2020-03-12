// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/keyword_plan_competition_level.proto

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
	return fileDescriptor_c20ed2e277b87d6e, []int{0, 0}
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
	return fileDescriptor_c20ed2e277b87d6e, []int{0}
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
	proto.RegisterEnum("google.ads.googleads.v1.enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel", KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_name, KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_value)
	proto.RegisterType((*KeywordPlanCompetitionLevelEnum)(nil), "google.ads.googleads.v1.enums.KeywordPlanCompetitionLevelEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/keyword_plan_competition_level.proto", fileDescriptor_c20ed2e277b87d6e)
}

var fileDescriptor_c20ed2e277b87d6e = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4f, 0xc2, 0x30,
	0x18, 0x95, 0x41, 0xc0, 0x94, 0x83, 0xcb, 0x8e, 0x2a, 0x51, 0xf8, 0x01, 0x5d, 0x16, 0x6f, 0xf5,
	0x34, 0x60, 0xc2, 0x02, 0x8c, 0x25, 0x06, 0x48, 0xc8, 0x12, 0x52, 0x59, 0xd3, 0x2c, 0x76, 0xed,
	0x42, 0x07, 0xc6, 0x83, 0x7f, 0xc6, 0xa3, 0x3f, 0xc5, 0x9f, 0xe2, 0xdd, 0xbb, 0x69, 0x2b, 0x78,
	0x72, 0x97, 0xe5, 0x65, 0xef, 0x7d, 0xef, 0xf5, 0x3d, 0xd0, 0xa7, 0x42, 0x50, 0x46, 0x5c, 0x9c,
	0x4a, 0xd7, 0x40, 0x85, 0x0e, 0x9e, 0x4b, 0xf8, 0x3e, 0x97, 0xee, 0x33, 0x79, 0x7d, 0x11, 0xbb,
	0x74, 0x53, 0x30, 0xcc, 0x37, 0x5b, 0x91, 0x17, 0xa4, 0xcc, 0xca, 0x4c, 0xf0, 0x0d, 0x23, 0x07,
	0xc2, 0x60, 0xb1, 0x13, 0xa5, 0x70, 0x3a, 0xe6, 0x10, 0xe2, 0x54, 0xc2, 0x93, 0x07, 0x3c, 0x78,
	0x50, 0x7b, 0x5c, 0x5e, 0x1f, 0x23, 0x8a, 0xcc, 0xc5, 0x9c, 0x8b, 0x12, 0x2b, 0x07, 0x69, 0x8e,
	0x7b, 0x6f, 0xe0, 0x66, 0x62, 0x42, 0x62, 0x86, 0xf9, 0xe0, 0x2f, 0x62, 0xaa, 0x12, 0x02, 0xbe,
	0xcf, 0x7b, 0x6b, 0x70, 0x55, 0x21, 0x71, 0x2e, 0x40, 0x7b, 0x11, 0x3d, 0xc6, 0xc1, 0x20, 0x7c,
	0x08, 0x83, 0xa1, 0x7d, 0xe6, 0xb4, 0x41, 0x6b, 0x11, 0x4d, 0xa2, 0xf9, 0x2a, 0xb2, 0x6b, 0x4e,
	0x0b, 0xd4, 0xa7, 0xf3, 0x95, 0x6d, 0x39, 0x00, 0x34, 0x67, 0xc1, 0x30, 0x5c, 0xcc, 0xec, 0xba,
	0x73, 0x0e, 0x1a, 0xe3, 0x70, 0x34, 0xb6, 0x1b, 0xfd, 0xef, 0x1a, 0xe8, 0x6e, 0x45, 0x0e, 0x2b,
	0x2b, 0xf4, 0x6f, 0x2b, 0xf2, 0x63, 0x55, 0x23, 0xae, 0xad, 0x7f, 0x97, 0x84, 0x54, 0x30, 0xcc,
	0x29, 0x14, 0x3b, 0xea, 0x52, 0xc2, 0x75, 0xc9, 0xe3, 0xb2, 0x45, 0x26, 0xff, 0x19, 0xfa, 0x5e,
	0x7f, 0xdf, 0xad, 0xfa, 0xc8, 0xf7, 0x3f, 0xac, 0xce, 0xc8, 0x58, 0xf9, 0xa9, 0x84, 0x06, 0x2a,
	0xb4, 0xf4, 0xa0, 0x9a, 0x43, 0x7e, 0x1e, 0xf9, 0xc4, 0x4f, 0x65, 0x72, 0xe2, 0x93, 0xa5, 0x97,
	0x68, 0xfe, 0xcb, 0xea, 0x9a, 0x9f, 0x08, 0xf9, 0xa9, 0x44, 0xe8, 0xa4, 0x40, 0x68, 0xe9, 0x21,
	0xa4, 0x35, 0x4f, 0x4d, 0xfd, 0xb0, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x77, 0x8e,
	0xaa, 0x00, 0x02, 0x00, 0x00,
}
