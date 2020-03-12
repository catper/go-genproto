// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/target_cpa_opt_in_recommendation_goal.proto

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

// Goal of TargetCpaOptIn recommendation.
type TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal int32

const (
	// Not specified.
	TargetCpaOptInRecommendationGoalEnum_UNSPECIFIED TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal = 0
	// Used for return value only. Represents value unknown in this version.
	TargetCpaOptInRecommendationGoalEnum_UNKNOWN TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal = 1
	// Recommendation to set Target CPA to maintain the same cost.
	TargetCpaOptInRecommendationGoalEnum_SAME_COST TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal = 2
	// Recommendation to set Target CPA to maintain the same conversions.
	TargetCpaOptInRecommendationGoalEnum_SAME_CONVERSIONS TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal = 3
	// Recommendation to set Target CPA to maintain the same CPA.
	TargetCpaOptInRecommendationGoalEnum_SAME_CPA TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal = 4
	// Recommendation to set Target CPA to a value that is as close as possible
	// to, yet lower than, the actual CPA (computed for past 28 days).
	TargetCpaOptInRecommendationGoalEnum_CLOSEST_CPA TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal = 5
)

var TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SAME_COST",
	3: "SAME_CONVERSIONS",
	4: "SAME_CPA",
	5: "CLOSEST_CPA",
}

var TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal_value = map[string]int32{
	"UNSPECIFIED":      0,
	"UNKNOWN":          1,
	"SAME_COST":        2,
	"SAME_CONVERSIONS": 3,
	"SAME_CPA":         4,
	"CLOSEST_CPA":      5,
}

func (x TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal) String() string {
	return proto.EnumName(TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal_name, int32(x))
}

func (TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6cda105fb4f0f31c, []int{0, 0}
}

// Container for enum describing goals for TargetCpaOptIn recommendation.
type TargetCpaOptInRecommendationGoalEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetCpaOptInRecommendationGoalEnum) Reset()         { *m = TargetCpaOptInRecommendationGoalEnum{} }
func (m *TargetCpaOptInRecommendationGoalEnum) String() string { return proto.CompactTextString(m) }
func (*TargetCpaOptInRecommendationGoalEnum) ProtoMessage()    {}
func (*TargetCpaOptInRecommendationGoalEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cda105fb4f0f31c, []int{0}
}

func (m *TargetCpaOptInRecommendationGoalEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpaOptInRecommendationGoalEnum.Unmarshal(m, b)
}
func (m *TargetCpaOptInRecommendationGoalEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpaOptInRecommendationGoalEnum.Marshal(b, m, deterministic)
}
func (m *TargetCpaOptInRecommendationGoalEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpaOptInRecommendationGoalEnum.Merge(m, src)
}
func (m *TargetCpaOptInRecommendationGoalEnum) XXX_Size() int {
	return xxx_messageInfo_TargetCpaOptInRecommendationGoalEnum.Size(m)
}
func (m *TargetCpaOptInRecommendationGoalEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpaOptInRecommendationGoalEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpaOptInRecommendationGoalEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal", TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal_name, TargetCpaOptInRecommendationGoalEnum_TargetCpaOptInRecommendationGoal_value)
	proto.RegisterType((*TargetCpaOptInRecommendationGoalEnum)(nil), "google.ads.googleads.v2.enums.TargetCpaOptInRecommendationGoalEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/target_cpa_opt_in_recommendation_goal.proto", fileDescriptor_6cda105fb4f0f31c)
}

var fileDescriptor_6cda105fb4f0f31c = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xdd, 0x4a, 0xf3, 0x30,
	0x18, 0xc7, 0xdf, 0x76, 0xef, 0x97, 0x99, 0x62, 0x29, 0x1e, 0x89, 0x03, 0x37, 0xf4, 0x34, 0x85,
	0x7a, 0x16, 0x8f, 0xba, 0x5a, 0x47, 0x51, 0xdb, 0xb2, 0x6e, 0x15, 0xa4, 0x50, 0xe2, 0x5a, 0x42,
	0xa1, 0x4d, 0x4a, 0x93, 0xed, 0x0a, 0xbc, 0x12, 0x0f, 0x05, 0x6f, 0xc4, 0x4b, 0xf1, 0x12, 0x3c,
	0x92, 0x26, 0xdb, 0xc0, 0x03, 0xdd, 0x49, 0xf8, 0xe7, 0xf9, 0xf8, 0x3d, 0x5f, 0xc0, 0x27, 0x8c,
	0x91, 0xaa, 0xb0, 0x70, 0xce, 0x2d, 0x25, 0x3b, 0xb5, 0xb2, 0xad, 0x82, 0x2e, 0x6b, 0x6e, 0x09,
	0xdc, 0x92, 0x42, 0x64, 0x8b, 0x06, 0x67, 0xac, 0x11, 0x59, 0x49, 0xb3, 0xb6, 0x58, 0xb0, 0xba,
	0x2e, 0x68, 0x8e, 0x45, 0xc9, 0x68, 0x46, 0x18, 0xae, 0x60, 0xd3, 0x32, 0xc1, 0xcc, 0x81, 0xca,
	0x87, 0x38, 0xe7, 0x70, 0x8b, 0x82, 0x2b, 0x1b, 0x4a, 0xd4, 0xf1, 0xc9, 0xa6, 0x52, 0x53, 0x5a,
	0x98, 0x52, 0x26, 0x24, 0x80, 0xab, 0xe4, 0xd1, 0xab, 0x06, 0xce, 0x66, 0xb2, 0x98, 0xdb, 0xe0,
	0xb0, 0x11, 0x3e, 0x9d, 0x7e, 0x29, 0x34, 0x61, 0xb8, 0xf2, 0xe8, 0xb2, 0x1e, 0x3d, 0x69, 0xe0,
	0x74, 0x57, 0xa0, 0x79, 0x08, 0xfa, 0xf3, 0x20, 0x8e, 0x3c, 0xd7, 0xbf, 0xf6, 0xbd, 0x2b, 0xe3,
	0x97, 0xd9, 0x07, 0xff, 0xe6, 0xc1, 0x4d, 0x10, 0xde, 0x07, 0x86, 0x66, 0x1e, 0x80, 0xbd, 0xd8,
	0xb9, 0xf3, 0x32, 0x37, 0x8c, 0x67, 0x86, 0x6e, 0x1e, 0x01, 0x63, 0xfd, 0x0d, 0x12, 0x6f, 0x1a,
	0xfb, 0x61, 0x10, 0x1b, 0x3d, 0x73, 0x1f, 0xfc, 0x57, 0xd6, 0xc8, 0x31, 0x7e, 0x77, 0x40, 0xf7,
	0x36, 0x8c, 0xbd, 0x78, 0x26, 0x0d, 0x7f, 0xc6, 0x1f, 0x1a, 0x18, 0x2e, 0x58, 0x0d, 0x7f, 0x9c,
	0x79, 0x7c, 0xbe, 0xab, 0xd3, 0xa8, 0x1b, 0x3e, 0xd2, 0x1e, 0xc6, 0x6b, 0x0e, 0x61, 0x15, 0xa6,
	0x04, 0xb2, 0x96, 0x58, 0xa4, 0xa0, 0x72, 0x35, 0x9b, 0xb3, 0x34, 0x25, 0xff, 0xe6, 0x4a, 0x97,
	0xf2, 0x7d, 0xd6, 0x7b, 0x13, 0xc7, 0x79, 0xd1, 0x07, 0x13, 0x85, 0x72, 0x72, 0x0e, 0x95, 0xec,
	0x54, 0x62, 0xc3, 0x6e, 0x7d, 0xfc, 0x6d, 0xe3, 0x4f, 0x9d, 0x9c, 0xa7, 0x5b, 0x7f, 0x9a, 0xd8,
	0xa9, 0xf4, 0xbf, 0xeb, 0x43, 0x65, 0x44, 0xc8, 0xc9, 0x39, 0x42, 0xdb, 0x08, 0x84, 0x12, 0x1b,
	0x21, 0x19, 0xf3, 0xf8, 0x57, 0x36, 0x76, 0xf1, 0x19, 0x00, 0x00, 0xff, 0xff, 0x74, 0x39, 0x28,
	0x1a, 0x3d, 0x02, 0x00, 0x00,
}
