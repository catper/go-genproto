// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/positive_geo_target_type.proto

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

// The possible positive geo target types.
type PositiveGeoTargetTypeEnum_PositiveGeoTargetType int32

const (
	// Not specified.
	PositiveGeoTargetTypeEnum_UNSPECIFIED PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 0
	// The value is unknown in this version.
	PositiveGeoTargetTypeEnum_UNKNOWN PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 1
	// Specifies that an ad is triggered if the user is in,
	// or shows interest in, advertiser's targeted locations.
	PositiveGeoTargetTypeEnum_PRESENCE_OR_INTEREST PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 5
	// Specifies that an ad is triggered if the user
	// searches for advertiser's targeted locations.
	PositiveGeoTargetTypeEnum_SEARCH_INTEREST PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 6
	// Specifies that an ad is triggered if the user is in
	// or regularly in advertiser's targeted locations.
	PositiveGeoTargetTypeEnum_PRESENCE PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 7
)

var PositiveGeoTargetTypeEnum_PositiveGeoTargetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	5: "PRESENCE_OR_INTEREST",
	6: "SEARCH_INTEREST",
	7: "PRESENCE",
}

var PositiveGeoTargetTypeEnum_PositiveGeoTargetType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"PRESENCE_OR_INTEREST": 5,
	"SEARCH_INTEREST":      6,
	"PRESENCE":             7,
}

func (x PositiveGeoTargetTypeEnum_PositiveGeoTargetType) String() string {
	return proto.EnumName(PositiveGeoTargetTypeEnum_PositiveGeoTargetType_name, int32(x))
}

func (PositiveGeoTargetTypeEnum_PositiveGeoTargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_74e8d47d77af9abc, []int{0, 0}
}

// Container for enum describing possible positive geo target types.
type PositiveGeoTargetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositiveGeoTargetTypeEnum) Reset()         { *m = PositiveGeoTargetTypeEnum{} }
func (m *PositiveGeoTargetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*PositiveGeoTargetTypeEnum) ProtoMessage()    {}
func (*PositiveGeoTargetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_74e8d47d77af9abc, []int{0}
}

func (m *PositiveGeoTargetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositiveGeoTargetTypeEnum.Unmarshal(m, b)
}
func (m *PositiveGeoTargetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositiveGeoTargetTypeEnum.Marshal(b, m, deterministic)
}
func (m *PositiveGeoTargetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositiveGeoTargetTypeEnum.Merge(m, src)
}
func (m *PositiveGeoTargetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PositiveGeoTargetTypeEnum.Size(m)
}
func (m *PositiveGeoTargetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PositiveGeoTargetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PositiveGeoTargetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.PositiveGeoTargetTypeEnum_PositiveGeoTargetType", PositiveGeoTargetTypeEnum_PositiveGeoTargetType_name, PositiveGeoTargetTypeEnum_PositiveGeoTargetType_value)
	proto.RegisterType((*PositiveGeoTargetTypeEnum)(nil), "google.ads.googleads.v3.enums.PositiveGeoTargetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/positive_geo_target_type.proto", fileDescriptor_74e8d47d77af9abc)
}

var fileDescriptor_74e8d47d77af9abc = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xf3, 0x30,
	0x1c, 0xfd, 0xba, 0x0f, 0x37, 0xc9, 0x84, 0x95, 0xaa, 0xa0, 0xc3, 0x5d, 0x6c, 0x0f, 0x90, 0x5e,
	0xf4, 0x2e, 0x7a, 0xd3, 0xcd, 0x38, 0x87, 0xd0, 0x95, 0xb6, 0x9b, 0x20, 0x85, 0x52, 0x6d, 0x08,
	0x85, 0x2d, 0x29, 0x4d, 0x36, 0xd8, 0x53, 0xf8, 0x0e, 0x5e, 0xfa, 0x28, 0x3e, 0x8a, 0x37, 0xbe,
	0x82, 0x34, 0x59, 0xeb, 0xcd, 0xf4, 0x26, 0x1c, 0x72, 0xfe, 0xf0, 0x3b, 0x07, 0xdc, 0x50, 0xce,
	0xe9, 0x8a, 0xd8, 0x69, 0x26, 0x6c, 0x0d, 0x2b, 0xb4, 0x75, 0x6c, 0xc2, 0x36, 0x6b, 0x61, 0x17,
	0x5c, 0xe4, 0x32, 0xdf, 0x92, 0x84, 0x12, 0x9e, 0xc8, 0xb4, 0xa4, 0x44, 0x26, 0x72, 0x57, 0x10,
	0x58, 0x94, 0x5c, 0x72, 0x6b, 0xa0, 0x2d, 0x30, 0xcd, 0x04, 0x6c, 0xdc, 0x70, 0xeb, 0x40, 0xe5,
	0xee, 0x5f, 0xd5, 0xe1, 0x45, 0x6e, 0xa7, 0x8c, 0x71, 0x99, 0xca, 0x9c, 0x33, 0xa1, 0xcd, 0xa3,
	0x57, 0x03, 0x5c, 0xfa, 0xfb, 0xfc, 0x29, 0xe1, 0x91, 0x4a, 0x8f, 0x76, 0x05, 0xc1, 0x6c, 0xb3,
	0x1e, 0x95, 0xe0, 0xfc, 0x20, 0x69, 0xf5, 0x40, 0x77, 0xe1, 0x85, 0x3e, 0x9e, 0xcc, 0xee, 0x66,
	0xf8, 0xd6, 0xfc, 0x67, 0x75, 0x41, 0x67, 0xe1, 0x3d, 0x78, 0xf3, 0x47, 0xcf, 0x34, 0xac, 0x0b,
	0x70, 0xe6, 0x07, 0x38, 0xc4, 0xde, 0x04, 0x27, 0xf3, 0x20, 0x99, 0x79, 0x11, 0x0e, 0x70, 0x18,
	0x99, 0x47, 0xd6, 0x29, 0xe8, 0x85, 0xd8, 0x0d, 0x26, 0xf7, 0x3f, 0x9f, 0x6d, 0xeb, 0x04, 0x1c,
	0xd7, 0x72, 0xb3, 0x33, 0xfe, 0x32, 0xc0, 0xf0, 0x85, 0xaf, 0xe1, 0x9f, 0xad, 0xc6, 0xfd, 0x83,
	0x77, 0xf9, 0x55, 0x27, 0xdf, 0x78, 0x1a, 0xef, 0xcd, 0x94, 0xaf, 0x52, 0x46, 0x21, 0x2f, 0xa9,
	0x4d, 0x09, 0x53, 0x8d, 0xeb, 0x81, 0x8b, 0x5c, 0xfc, 0xb2, 0xf7, 0xb5, 0x7a, 0xdf, 0x5a, 0xff,
	0xa7, 0xae, 0xfb, 0xde, 0x1a, 0x4c, 0x75, 0x94, 0x9b, 0x09, 0xa8, 0x61, 0x85, 0x96, 0x0e, 0xac,
	0x06, 0x12, 0x1f, 0x35, 0x1f, 0xbb, 0x99, 0x88, 0x1b, 0x3e, 0x5e, 0x3a, 0xb1, 0xe2, 0x3f, 0x5b,
	0x43, 0xfd, 0x89, 0x90, 0x9b, 0x09, 0x84, 0x1a, 0x05, 0x42, 0x4b, 0x07, 0x21, 0xa5, 0x79, 0x6e,
	0xab, 0xc3, 0x9c, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x03, 0xaa, 0xe5, 0x07, 0x02, 0x00,
	0x00,
}
