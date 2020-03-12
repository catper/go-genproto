// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/feed_mapping_criterion_type.proto

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

// Possible placeholder types for a feed mapping.
type FeedMappingCriterionTypeEnum_FeedMappingCriterionType int32

const (
	// Not specified.
	FeedMappingCriterionTypeEnum_UNSPECIFIED FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 0
	// Used for return value only. Represents value unknown in this version.
	FeedMappingCriterionTypeEnum_UNKNOWN FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 1
	// Allows campaign targeting at locations within a location feed.
	FeedMappingCriterionTypeEnum_LOCATION_EXTENSION_TARGETING FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 4
	// Allows url targeting for your dynamic search ads within a page feed.
	FeedMappingCriterionTypeEnum_DSA_PAGE_FEED FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 3
)

var FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	4: "LOCATION_EXTENSION_TARGETING",
	3: "DSA_PAGE_FEED",
}

var FeedMappingCriterionTypeEnum_FeedMappingCriterionType_value = map[string]int32{
	"UNSPECIFIED":                  0,
	"UNKNOWN":                      1,
	"LOCATION_EXTENSION_TARGETING": 4,
	"DSA_PAGE_FEED":                3,
}

func (x FeedMappingCriterionTypeEnum_FeedMappingCriterionType) String() string {
	return proto.EnumName(FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name, int32(x))
}

func (FeedMappingCriterionTypeEnum_FeedMappingCriterionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0fb3d458969bb53, []int{0, 0}
}

// Container for enum describing possible criterion types for a feed mapping.
type FeedMappingCriterionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedMappingCriterionTypeEnum) Reset()         { *m = FeedMappingCriterionTypeEnum{} }
func (m *FeedMappingCriterionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FeedMappingCriterionTypeEnum) ProtoMessage()    {}
func (*FeedMappingCriterionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0fb3d458969bb53, []int{0}
}

func (m *FeedMappingCriterionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Unmarshal(m, b)
}
func (m *FeedMappingCriterionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Marshal(b, m, deterministic)
}
func (m *FeedMappingCriterionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingCriterionTypeEnum.Merge(m, src)
}
func (m *FeedMappingCriterionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Size(m)
}
func (m *FeedMappingCriterionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingCriterionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingCriterionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.FeedMappingCriterionTypeEnum_FeedMappingCriterionType", FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name, FeedMappingCriterionTypeEnum_FeedMappingCriterionType_value)
	proto.RegisterType((*FeedMappingCriterionTypeEnum)(nil), "google.ads.googleads.v2.enums.FeedMappingCriterionTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/feed_mapping_criterion_type.proto", fileDescriptor_e0fb3d458969bb53)
}

var fileDescriptor_e0fb3d458969bb53 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x6a, 0xe3, 0x30,
	0x14, 0x1d, 0x27, 0xc3, 0x0c, 0x28, 0x0c, 0xe3, 0x7a, 0x55, 0x4a, 0x02, 0x4d, 0x0e, 0x20, 0x83,
	0xbb, 0x53, 0x17, 0x45, 0x49, 0x14, 0x63, 0xda, 0x3a, 0xa6, 0x71, 0xd2, 0x52, 0x0c, 0xc6, 0x8d,
	0x54, 0x61, 0x88, 0x25, 0x61, 0x39, 0x81, 0x1c, 0xa2, 0x97, 0xe8, 0xb2, 0x47, 0xe9, 0x51, 0xba,
	0xec, 0x09, 0x8a, 0xad, 0x3a, 0xbb, 0x74, 0x23, 0x1e, 0x7a, 0xff, 0xbd, 0xff, 0xdf, 0x03, 0x57,
	0x5c, 0x4a, 0xbe, 0x61, 0x6e, 0x46, 0xb5, 0x6b, 0x60, 0x8d, 0x76, 0x9e, 0xcb, 0xc4, 0xb6, 0xd0,
	0xee, 0x33, 0x63, 0x34, 0x2d, 0x32, 0xa5, 0x72, 0xc1, 0xd3, 0x75, 0x99, 0x57, 0xac, 0xcc, 0xa5,
	0x48, 0xab, 0xbd, 0x62, 0x50, 0x95, 0xb2, 0x92, 0xce, 0xc0, 0xa8, 0x60, 0x46, 0x35, 0x3c, 0x18,
	0xc0, 0x9d, 0x07, 0x1b, 0x83, 0xb3, 0x7e, 0xeb, 0xaf, 0x72, 0x37, 0x13, 0x42, 0x56, 0x59, 0x95,
	0x4b, 0xa1, 0x8d, 0x78, 0xf4, 0x62, 0x81, 0xfe, 0x8c, 0x31, 0x7a, 0x6b, 0x36, 0x4c, 0xda, 0x05,
	0xf1, 0x5e, 0x31, 0x22, 0xb6, 0xc5, 0xa8, 0x00, 0xa7, 0xc7, 0x78, 0xe7, 0x3f, 0xe8, 0x2d, 0xc3,
	0x45, 0x44, 0x26, 0xc1, 0x2c, 0x20, 0x53, 0xfb, 0x97, 0xd3, 0x03, 0x7f, 0x97, 0xe1, 0x75, 0x38,
	0xbf, 0x0f, 0x6d, 0xcb, 0x39, 0x07, 0xfd, 0x9b, 0xf9, 0x04, 0xc7, 0xc1, 0x3c, 0x4c, 0xc9, 0x43,
	0x4c, 0xc2, 0x45, 0x8d, 0x62, 0x7c, 0xe7, 0x93, 0x38, 0x08, 0x7d, 0xfb, 0xb7, 0x73, 0x02, 0xfe,
	0x4d, 0x17, 0x38, 0x8d, 0xb0, 0x4f, 0xd2, 0x19, 0x21, 0x53, 0xbb, 0x3b, 0xfe, 0xb4, 0xc0, 0x70,
	0x2d, 0x0b, 0xf8, 0x63, 0xa6, 0xf1, 0xe0, 0xd8, 0x49, 0x51, 0x1d, 0x2a, 0xb2, 0x1e, 0xc7, 0xdf,
	0x7a, 0x2e, 0x37, 0x99, 0xe0, 0x50, 0x96, 0xdc, 0xe5, 0x4c, 0x34, 0x91, 0xdb, 0x92, 0x55, 0xae,
	0x8f, 0x74, 0x7e, 0xd9, 0xbc, 0xaf, 0x9d, 0xae, 0x8f, 0xf1, 0x5b, 0x67, 0xe0, 0x1b, 0x2b, 0x4c,
	0x35, 0x34, 0xb0, 0x46, 0x2b, 0x0f, 0xd6, 0xf5, 0xe8, 0xf7, 0x96, 0x4f, 0x30, 0xd5, 0xc9, 0x81,
	0x4f, 0x56, 0x5e, 0xd2, 0xf0, 0x1f, 0x9d, 0xa1, 0xf9, 0x44, 0x08, 0x53, 0x8d, 0xd0, 0x61, 0x02,
	0xa1, 0x95, 0x87, 0x50, 0x33, 0xf3, 0xf4, 0xa7, 0x39, 0xec, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x0b, 0x44, 0xd2, 0x4d, 0x0b, 0x02, 0x00, 0x00,
}
