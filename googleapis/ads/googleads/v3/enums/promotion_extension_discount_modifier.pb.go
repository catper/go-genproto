// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/promotion_extension_discount_modifier.proto

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

// A promotion extension discount modifier.
type PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier int32

const (
	// Not specified.
	PromotionExtensionDiscountModifierEnum_UNSPECIFIED PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier = 0
	// Used for return value only. Represents value unknown in this version.
	PromotionExtensionDiscountModifierEnum_UNKNOWN PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier = 1
	// 'Up to'.
	PromotionExtensionDiscountModifierEnum_UP_TO PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier = 2
)

var PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UP_TO",
}

var PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"UP_TO":       2,
}

func (x PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier) String() string {
	return proto.EnumName(PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_name, int32(x))
}

func (PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b6cacf196203e5e, []int{0, 0}
}

// Container for enum describing possible a promotion extension
// discount modifier.
type PromotionExtensionDiscountModifierEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromotionExtensionDiscountModifierEnum) Reset() {
	*m = PromotionExtensionDiscountModifierEnum{}
}
func (m *PromotionExtensionDiscountModifierEnum) String() string { return proto.CompactTextString(m) }
func (*PromotionExtensionDiscountModifierEnum) ProtoMessage()    {}
func (*PromotionExtensionDiscountModifierEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b6cacf196203e5e, []int{0}
}

func (m *PromotionExtensionDiscountModifierEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromotionExtensionDiscountModifierEnum.Unmarshal(m, b)
}
func (m *PromotionExtensionDiscountModifierEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromotionExtensionDiscountModifierEnum.Marshal(b, m, deterministic)
}
func (m *PromotionExtensionDiscountModifierEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromotionExtensionDiscountModifierEnum.Merge(m, src)
}
func (m *PromotionExtensionDiscountModifierEnum) XXX_Size() int {
	return xxx_messageInfo_PromotionExtensionDiscountModifierEnum.Size(m)
}
func (m *PromotionExtensionDiscountModifierEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PromotionExtensionDiscountModifierEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PromotionExtensionDiscountModifierEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier", PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_name, PromotionExtensionDiscountModifierEnum_PromotionExtensionDiscountModifier_value)
	proto.RegisterType((*PromotionExtensionDiscountModifierEnum)(nil), "google.ads.googleads.v3.enums.PromotionExtensionDiscountModifierEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/promotion_extension_discount_modifier.proto", fileDescriptor_7b6cacf196203e5e)
}

var fileDescriptor_7b6cacf196203e5e = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4e, 0xfb, 0x30,
	0x10, 0xc6, 0xff, 0xcd, 0x5f, 0x80, 0x70, 0x07, 0xaa, 0x8c, 0x88, 0x0e, 0xed, 0x00, 0x9b, 0x3d,
	0x78, 0x33, 0x53, 0x4a, 0x43, 0x55, 0xa1, 0xa6, 0x91, 0xa0, 0x45, 0x42, 0x91, 0xaa, 0x50, 0x1b,
	0xcb, 0x52, 0xe3, 0x8b, 0xe2, 0xb4, 0xf0, 0x3c, 0x8c, 0x3c, 0x0a, 0x8f, 0xc2, 0x33, 0x30, 0x20,
	0xdb, 0x4d, 0x36, 0xe8, 0x62, 0x7d, 0xf2, 0xdd, 0x7d, 0xbf, 0xbb, 0x0f, 0x4d, 0x25, 0x80, 0xdc,
	0x08, 0x92, 0x73, 0x43, 0xbc, 0xb4, 0x6a, 0x47, 0x89, 0xd0, 0xdb, 0xc2, 0x90, 0xb2, 0x82, 0x02,
	0x6a, 0x05, 0x7a, 0x25, 0xde, 0x6a, 0xa1, 0x8d, 0x55, 0x5c, 0x99, 0x35, 0x6c, 0x75, 0xbd, 0x2a,
	0x80, 0xab, 0x17, 0x25, 0x2a, 0x5c, 0x56, 0x50, 0x43, 0xd8, 0xf7, 0xf3, 0x38, 0xe7, 0x06, 0xb7,
	0x56, 0x78, 0x47, 0xb1, 0xb3, 0x3a, 0xbf, 0x68, 0x48, 0xa5, 0x22, 0xb9, 0xd6, 0x50, 0xe7, 0xd6,
	0xd7, 0xf8, 0xe1, 0xe1, 0x2b, 0xba, 0x4c, 0x1b, 0x56, 0xdc, 0xa0, 0xc6, 0x7b, 0xd2, 0x6c, 0x0f,
	0x8a, 0xf5, 0xb6, 0x18, 0xce, 0xd0, 0xf0, 0x70, 0x67, 0x78, 0x86, 0xba, 0x8b, 0xe4, 0x3e, 0x8d,
	0x6f, 0xa6, 0xb7, 0xd3, 0x78, 0xdc, 0xfb, 0x17, 0x76, 0xd1, 0xc9, 0x22, 0xb9, 0x4b, 0xe6, 0x8f,
	0x49, 0xaf, 0x13, 0x9e, 0xa2, 0xa3, 0x45, 0xba, 0x7a, 0x98, 0xf7, 0x82, 0xd1, 0x77, 0x07, 0x0d,
	0xd6, 0x50, 0xe0, 0x3f, 0x97, 0x1f, 0x5d, 0x1d, 0x46, 0xa6, 0xf6, 0x8e, 0xb4, 0xf3, 0x34, 0xda,
	0x3b, 0x49, 0xd8, 0xe4, 0x5a, 0x62, 0xa8, 0x24, 0x91, 0x42, 0xbb, 0x2b, 0x9b, 0x84, 0x4b, 0x65,
	0x7e, 0x09, 0xfc, 0xda, 0xbd, 0xef, 0xc1, 0xff, 0x49, 0x14, 0x7d, 0x04, 0xfd, 0x89, 0xb7, 0x8a,
	0xb8, 0xc1, 0x5e, 0x5a, 0xb5, 0xa4, 0xd8, 0x06, 0x61, 0x3e, 0x9b, 0x7a, 0x16, 0x71, 0x93, 0xb5,
	0xf5, 0x6c, 0x49, 0x33, 0x57, 0xff, 0x0a, 0x06, 0xfe, 0x93, 0xb1, 0x88, 0x1b, 0xc6, 0xda, 0x0e,
	0xc6, 0x96, 0x94, 0x31, 0xd7, 0xf3, 0x7c, 0xec, 0x16, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x83, 0x52, 0x77, 0xd7, 0x08, 0x02, 0x00, 0x00,
}
