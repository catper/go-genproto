// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/vanity_pharma_display_url_mode.proto

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

// Enum describing possible display modes for vanity pharma URLs.
type VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode int32

const (
	// Not specified.
	VanityPharmaDisplayUrlModeEnum_UNSPECIFIED VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode = 0
	// Used for return value only. Represents value unknown in this version.
	VanityPharmaDisplayUrlModeEnum_UNKNOWN VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode = 1
	// Replace vanity pharma URL with manufacturer website url.
	VanityPharmaDisplayUrlModeEnum_MANUFACTURER_WEBSITE_URL VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode = 2
	// Replace vanity pharma URL with description of the website.
	VanityPharmaDisplayUrlModeEnum_WEBSITE_DESCRIPTION VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode = 3
)

var VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MANUFACTURER_WEBSITE_URL",
	3: "WEBSITE_DESCRIPTION",
}

var VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"MANUFACTURER_WEBSITE_URL": 2,
	"WEBSITE_DESCRIPTION":      3,
}

func (x VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode) String() string {
	return proto.EnumName(VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_name, int32(x))
}

func (VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_da95e5df129a03ce, []int{0, 0}
}

// The display mode for vanity pharma URLs.
type VanityPharmaDisplayUrlModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VanityPharmaDisplayUrlModeEnum) Reset()         { *m = VanityPharmaDisplayUrlModeEnum{} }
func (m *VanityPharmaDisplayUrlModeEnum) String() string { return proto.CompactTextString(m) }
func (*VanityPharmaDisplayUrlModeEnum) ProtoMessage()    {}
func (*VanityPharmaDisplayUrlModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_da95e5df129a03ce, []int{0}
}

func (m *VanityPharmaDisplayUrlModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.Unmarshal(m, b)
}
func (m *VanityPharmaDisplayUrlModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.Marshal(b, m, deterministic)
}
func (m *VanityPharmaDisplayUrlModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.Merge(m, src)
}
func (m *VanityPharmaDisplayUrlModeEnum) XXX_Size() int {
	return xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.Size(m)
}
func (m *VanityPharmaDisplayUrlModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_VanityPharmaDisplayUrlModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_VanityPharmaDisplayUrlModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode", VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_name, VanityPharmaDisplayUrlModeEnum_VanityPharmaDisplayUrlMode_value)
	proto.RegisterType((*VanityPharmaDisplayUrlModeEnum)(nil), "google.ads.googleads.v2.enums.VanityPharmaDisplayUrlModeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/vanity_pharma_display_url_mode.proto", fileDescriptor_da95e5df129a03ce)
}

var fileDescriptor_da95e5df129a03ce = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x6a, 0xb3, 0x40,
	0x18, 0xfd, 0x35, 0xf0, 0x17, 0x26, 0x8b, 0x8a, 0x5d, 0xb4, 0x84, 0xa4, 0x25, 0x39, 0xc0, 0x08,
	0x76, 0x37, 0x5d, 0x69, 0x62, 0x82, 0xb4, 0x31, 0x62, 0xa2, 0x81, 0x22, 0xc8, 0xb4, 0x23, 0x56,
	0xd0, 0x19, 0xeb, 0x98, 0x40, 0xce, 0xd1, 0x1b, 0x74, 0xd9, 0xa3, 0xf4, 0x28, 0x5d, 0xf7, 0x00,
	0xc5, 0x99, 0x9a, 0x5d, 0xba, 0x19, 0x1e, 0xf3, 0xbd, 0xef, 0xbd, 0xef, 0x3d, 0x60, 0x67, 0x8c,
	0x65, 0x45, 0x6a, 0x60, 0xc2, 0x0d, 0x09, 0x5b, 0xb4, 0x37, 0x8d, 0x94, 0xee, 0x4a, 0x6e, 0xec,
	0x31, 0xcd, 0x9b, 0x43, 0x52, 0xbd, 0xe0, 0xba, 0xc4, 0x09, 0xc9, 0x79, 0x55, 0xe0, 0x43, 0xb2,
	0xab, 0x8b, 0xa4, 0x64, 0x24, 0x85, 0x55, 0xcd, 0x1a, 0xa6, 0x8f, 0xe4, 0x22, 0xc4, 0x84, 0xc3,
	0xa3, 0x06, 0xdc, 0x9b, 0x50, 0x68, 0x0c, 0x86, 0x9d, 0x45, 0x95, 0x1b, 0x98, 0x52, 0xd6, 0xe0,
	0x26, 0x67, 0x94, 0xcb, 0xe5, 0xc9, 0x9b, 0x02, 0xae, 0x23, 0xe1, 0xe2, 0x0b, 0x93, 0x99, 0xf4,
	0x08, 0xeb, 0x62, 0xc9, 0x48, 0xea, 0xd0, 0x5d, 0x39, 0x79, 0x05, 0x83, 0xd3, 0x0c, 0xfd, 0x1c,
	0xf4, 0x43, 0x6f, 0xed, 0x3b, 0x53, 0x77, 0xee, 0x3a, 0x33, 0xed, 0x9f, 0xde, 0x07, 0x67, 0xa1,
	0x77, 0xef, 0xad, 0xb6, 0x9e, 0xa6, 0xe8, 0x43, 0x70, 0xb5, 0xb4, 0xbc, 0x70, 0x6e, 0x4d, 0x37,
	0x61, 0xe0, 0x04, 0xc9, 0xd6, 0xb1, 0xd7, 0xee, 0xc6, 0x49, 0xc2, 0xe0, 0x41, 0x53, 0xf5, 0x4b,
	0x70, 0xd1, 0x7d, 0xcc, 0x9c, 0xf5, 0x34, 0x70, 0xfd, 0x8d, 0xbb, 0xf2, 0xb4, 0x9e, 0xfd, 0xad,
	0x80, 0xf1, 0x33, 0x2b, 0xe1, 0x9f, 0xc9, 0xec, 0x9b, 0xd3, 0x67, 0xf9, 0x6d, 0x38, 0x5f, 0x79,
	0xfc, 0xed, 0x17, 0x66, 0xac, 0xc0, 0x34, 0x83, 0xac, 0xce, 0x8c, 0x2c, 0xa5, 0x22, 0x7a, 0xd7,
	0x77, 0x95, 0xf3, 0x13, 0xf5, 0xdf, 0x89, 0xf7, 0x5d, 0xed, 0x2d, 0x2c, 0xeb, 0x43, 0x1d, 0x2d,
	0xa4, 0x94, 0x45, 0x38, 0x94, 0xb0, 0x45, 0x91, 0x09, 0xdb, 0x92, 0xf8, 0x67, 0x37, 0x8f, 0x2d,
	0xc2, 0xe3, 0xe3, 0x3c, 0x8e, 0xcc, 0x58, 0xcc, 0xbf, 0xd4, 0xb1, 0xfc, 0x44, 0xc8, 0x22, 0x1c,
	0xa1, 0x23, 0x03, 0xa1, 0xc8, 0x44, 0x48, 0x70, 0x9e, 0xfe, 0x8b, 0xc3, 0x6e, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x34, 0x91, 0x2a, 0x2d, 0x16, 0x02, 0x00, 0x00,
}
