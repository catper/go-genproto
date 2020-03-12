// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/feed_item_target_type.proto

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

// Possible type of a feed item target.
type FeedItemTargetTypeEnum_FeedItemTargetType int32

const (
	// Not specified.
	FeedItemTargetTypeEnum_UNSPECIFIED FeedItemTargetTypeEnum_FeedItemTargetType = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemTargetTypeEnum_UNKNOWN FeedItemTargetTypeEnum_FeedItemTargetType = 1
	// Feed item targets a campaign.
	FeedItemTargetTypeEnum_CAMPAIGN FeedItemTargetTypeEnum_FeedItemTargetType = 2
	// Feed item targets an ad group.
	FeedItemTargetTypeEnum_AD_GROUP FeedItemTargetTypeEnum_FeedItemTargetType = 3
	// Feed item targets a criterion.
	FeedItemTargetTypeEnum_CRITERION FeedItemTargetTypeEnum_FeedItemTargetType = 4
)

var FeedItemTargetTypeEnum_FeedItemTargetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CAMPAIGN",
	3: "AD_GROUP",
	4: "CRITERION",
}

var FeedItemTargetTypeEnum_FeedItemTargetType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CAMPAIGN":    2,
	"AD_GROUP":    3,
	"CRITERION":   4,
}

func (x FeedItemTargetTypeEnum_FeedItemTargetType) String() string {
	return proto.EnumName(FeedItemTargetTypeEnum_FeedItemTargetType_name, int32(x))
}

func (FeedItemTargetTypeEnum_FeedItemTargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d72b8d5b7777593c, []int{0, 0}
}

// Container for enum describing possible types of a feed item target.
type FeedItemTargetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemTargetTypeEnum) Reset()         { *m = FeedItemTargetTypeEnum{} }
func (m *FeedItemTargetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemTargetTypeEnum) ProtoMessage()    {}
func (*FeedItemTargetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d72b8d5b7777593c, []int{0}
}

func (m *FeedItemTargetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemTargetTypeEnum.Unmarshal(m, b)
}
func (m *FeedItemTargetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemTargetTypeEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemTargetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemTargetTypeEnum.Merge(m, src)
}
func (m *FeedItemTargetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemTargetTypeEnum.Size(m)
}
func (m *FeedItemTargetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemTargetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemTargetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FeedItemTargetTypeEnum_FeedItemTargetType", FeedItemTargetTypeEnum_FeedItemTargetType_name, FeedItemTargetTypeEnum_FeedItemTargetType_value)
	proto.RegisterType((*FeedItemTargetTypeEnum)(nil), "google.ads.googleads.v3.enums.FeedItemTargetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/feed_item_target_type.proto", fileDescriptor_d72b8d5b7777593c)
}

var fileDescriptor_d72b8d5b7777593c = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0x76, 0x9d, 0xf8, 0x27, 0x53, 0x2c, 0x3d, 0x28, 0x88, 0x3b, 0x6c, 0x0f, 0x90, 0x1e, 0x7a,
	0x32, 0x9e, 0xb2, 0xad, 0x2b, 0x45, 0xec, 0xca, 0xdc, 0x26, 0x48, 0xa5, 0x54, 0xf3, 0x33, 0x14,
	0xd6, 0xa4, 0x2c, 0xd9, 0x64, 0xaf, 0xe3, 0xd1, 0x47, 0xf1, 0x51, 0xc4, 0x87, 0x90, 0xa6, 0x6e,
	0x97, 0xa1, 0x97, 0xf0, 0x25, 0xdf, 0x1f, 0xbe, 0x7c, 0xe8, 0x9a, 0x4b, 0xc9, 0xe7, 0xe0, 0x66,
	0x4c, 0xb9, 0x35, 0xac, 0xd0, 0xca, 0x73, 0x41, 0x2c, 0x0b, 0xe5, 0xbe, 0x02, 0xb0, 0x34, 0xd7,
	0x50, 0xa4, 0x3a, 0x5b, 0x70, 0xd0, 0xa9, 0x5e, 0x97, 0x80, 0xcb, 0x85, 0xd4, 0xd2, 0x69, 0xd7,
	0x7a, 0x9c, 0x31, 0x85, 0xb7, 0x56, 0xbc, 0xf2, 0xb0, 0xb1, 0x5e, 0x5e, 0x6d, 0x92, 0xcb, 0xdc,
	0xcd, 0x84, 0x90, 0x3a, 0xd3, 0xb9, 0x14, 0xaa, 0x36, 0x77, 0xdf, 0xd0, 0xf9, 0x10, 0x80, 0x85,
	0x1a, 0x8a, 0x89, 0x49, 0x9e, 0xac, 0x4b, 0xf0, 0xc5, 0xb2, 0xe8, 0x3e, 0x21, 0x67, 0x97, 0x71,
	0xce, 0x50, 0x6b, 0x1a, 0xdd, 0xc7, 0x7e, 0x3f, 0x1c, 0x86, 0xfe, 0xc0, 0xde, 0x73, 0x5a, 0xe8,
	0x70, 0x1a, 0xdd, 0x46, 0xa3, 0x87, 0xc8, 0x6e, 0x38, 0x27, 0xe8, 0xa8, 0x4f, 0xef, 0x62, 0x1a,
	0x06, 0x91, 0x6d, 0x55, 0x37, 0x3a, 0x48, 0x83, 0xf1, 0x68, 0x1a, 0xdb, 0x4d, 0xe7, 0x14, 0x1d,
	0xf7, 0xc7, 0xe1, 0xc4, 0x1f, 0x87, 0xa3, 0xc8, 0xde, 0xef, 0x7d, 0x37, 0x50, 0xe7, 0x45, 0x16,
	0xf8, 0xdf, 0xf2, 0xbd, 0x8b, 0xdd, 0x0a, 0x71, 0xd5, 0x3b, 0x6e, 0x3c, 0xf6, 0x7e, 0x9d, 0x5c,
	0xce, 0x33, 0xc1, 0xb1, 0x5c, 0x70, 0x97, 0x83, 0x30, 0xbf, 0xda, 0x2c, 0x58, 0xe6, 0xea, 0x8f,
	0x41, 0x6f, 0xcc, 0xf9, 0x6e, 0x35, 0x03, 0x4a, 0x3f, 0xac, 0x76, 0x50, 0x47, 0x51, 0xa6, 0x70,
	0x0d, 0x2b, 0x34, 0xf3, 0x70, 0x35, 0x84, 0xfa, 0xdc, 0xf0, 0x09, 0x65, 0x2a, 0xd9, 0xf2, 0xc9,
	0xcc, 0x4b, 0x0c, 0xff, 0x65, 0x75, 0xea, 0x47, 0x42, 0x28, 0x53, 0x84, 0x6c, 0x15, 0x84, 0xcc,
	0x3c, 0x42, 0x8c, 0xe6, 0xf9, 0xc0, 0x14, 0xf3, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x7a,
	0x11, 0xc7, 0xe8, 0x01, 0x00, 0x00,
}
