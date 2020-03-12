// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/feed_attribute_reference_error.proto

package errors

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

// Enum describing possible feed attribute reference errors.
type FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError int32

const (
	// Enum unspecified.
	FeedAttributeReferenceErrorEnum_UNSPECIFIED FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 0
	// The received error code is not known in this version.
	FeedAttributeReferenceErrorEnum_UNKNOWN FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 1
	// A feed referenced by ID has been removed.
	FeedAttributeReferenceErrorEnum_CANNOT_REFERENCE_REMOVED_FEED FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 2
	// There is no enabled feed with the given name.
	FeedAttributeReferenceErrorEnum_INVALID_FEED_NAME FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 3
	// There is no feed attribute in an enabled feed with the given name.
	FeedAttributeReferenceErrorEnum_INVALID_FEED_ATTRIBUTE_NAME FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 4
)

var FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CANNOT_REFERENCE_REMOVED_FEED",
	3: "INVALID_FEED_NAME",
	4: "INVALID_FEED_ATTRIBUTE_NAME",
}

var FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"CANNOT_REFERENCE_REMOVED_FEED": 2,
	"INVALID_FEED_NAME":             3,
	"INVALID_FEED_ATTRIBUTE_NAME":   4,
}

func (x FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError) String() string {
	return proto.EnumName(FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name, int32(x))
}

func (FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc8d6ce7179580a8, []int{0, 0}
}

// Container for enum describing possible feed attribute reference errors.
type FeedAttributeReferenceErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedAttributeReferenceErrorEnum) Reset()         { *m = FeedAttributeReferenceErrorEnum{} }
func (m *FeedAttributeReferenceErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedAttributeReferenceErrorEnum) ProtoMessage()    {}
func (*FeedAttributeReferenceErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8d6ce7179580a8, []int{0}
}

func (m *FeedAttributeReferenceErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Unmarshal(m, b)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Marshal(b, m, deterministic)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttributeReferenceErrorEnum.Merge(m, src)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Size(m)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttributeReferenceErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttributeReferenceErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError", FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name, FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_value)
	proto.RegisterType((*FeedAttributeReferenceErrorEnum)(nil), "google.ads.googleads.v3.errors.FeedAttributeReferenceErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/feed_attribute_reference_error.proto", fileDescriptor_fc8d6ce7179580a8)
}

var fileDescriptor_fc8d6ce7179580a8 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x6e, 0xa3, 0x30,
	0x14, 0x86, 0x07, 0x32, 0x9a, 0x91, 0x9c, 0xc5, 0x30, 0x48, 0xdd, 0x34, 0x6d, 0xd2, 0x72, 0x00,
	0xb3, 0x60, 0xe7, 0xae, 0x1c, 0x30, 0x11, 0x6a, 0xe3, 0x44, 0x94, 0x50, 0xa9, 0x42, 0x42, 0x24,
	0x38, 0x08, 0x29, 0xc1, 0x91, 0x4d, 0x72, 0x94, 0x1e, 0xa0, 0xcb, 0x1e, 0xa1, 0x47, 0xe8, 0x51,
	0x7a, 0x85, 0x6e, 0x2a, 0x70, 0x89, 0xd4, 0x45, 0xb3, 0xf2, 0xaf, 0xe7, 0xef, 0xfd, 0xff, 0xd3,
	0x7b, 0xc0, 0x2d, 0x38, 0x2f, 0x36, 0xcc, 0xce, 0x72, 0x69, 0x2b, 0xd9, 0xa8, 0x83, 0x63, 0x33,
	0x21, 0xb8, 0x90, 0xf6, 0x9a, 0xb1, 0x3c, 0xcd, 0xea, 0x5a, 0x94, 0xcb, 0x7d, 0xcd, 0x52, 0xc1,
	0xd6, 0x4c, 0xb0, 0x6a, 0xc5, 0xd2, 0xf6, 0x1f, 0xee, 0x04, 0xaf, 0xb9, 0x39, 0x54, 0x9d, 0x30,
	0xcb, 0x25, 0x3c, 0x9a, 0xc0, 0x83, 0x03, 0x95, 0xc9, 0xf9, 0x45, 0x17, 0xb2, 0x2b, 0xed, 0xac,
	0xaa, 0x78, 0x9d, 0xd5, 0x25, 0xaf, 0xa4, 0xea, 0xb6, 0x5e, 0x35, 0x30, 0xf2, 0x19, 0xcb, 0x71,
	0x97, 0x12, 0x76, 0x21, 0xa4, 0x69, 0x27, 0xd5, 0x7e, 0x6b, 0x3d, 0x69, 0x60, 0x70, 0x82, 0x31,
	0xff, 0x81, 0xfe, 0x82, 0xde, 0xcf, 0x89, 0x1b, 0xf8, 0x01, 0xf1, 0x8c, 0x5f, 0x66, 0x1f, 0xfc,
	0x5d, 0xd0, 0x5b, 0x3a, 0x7b, 0xa0, 0x86, 0x66, 0x5e, 0x83, 0x4b, 0x17, 0x53, 0x3a, 0x8b, 0xd2,
	0x90, 0xf8, 0x24, 0x24, 0xd4, 0x25, 0x69, 0x48, 0xa6, 0xb3, 0x98, 0x78, 0xa9, 0x4f, 0x88, 0x67,
	0xe8, 0xe6, 0x19, 0xf8, 0x1f, 0xd0, 0x18, 0xdf, 0x05, 0xaa, 0x92, 0x52, 0x3c, 0x25, 0x46, 0xcf,
	0x1c, 0x81, 0xc1, 0xb7, 0x32, 0x8e, 0xa2, 0x30, 0x18, 0x2f, 0x22, 0xa2, 0x80, 0xdf, 0xe3, 0x0f,
	0x0d, 0x58, 0x2b, 0xbe, 0x85, 0xa7, 0x37, 0x30, 0xbe, 0x3a, 0x31, 0xfc, 0xbc, 0xd9, 0xc2, 0x5c,
	0x7b, 0xf4, 0xbe, 0x3c, 0x0a, 0xbe, 0xc9, 0xaa, 0x02, 0x72, 0x51, 0xd8, 0x05, 0xab, 0xda, 0x1d,
	0x75, 0xa7, 0xd9, 0x95, 0xf2, 0xa7, 0x4b, 0xdd, 0xa8, 0xe7, 0x59, 0xef, 0x4d, 0x30, 0x7e, 0xd1,
	0x87, 0x13, 0x65, 0x86, 0x73, 0x09, 0x95, 0x6c, 0x54, 0xec, 0xc0, 0x36, 0x52, 0xbe, 0x75, 0x40,
	0x82, 0x73, 0x99, 0x1c, 0x81, 0x24, 0x76, 0x12, 0x05, 0xbc, 0xeb, 0x96, 0xaa, 0x22, 0x84, 0x73,
	0x89, 0xd0, 0x11, 0x41, 0x28, 0x76, 0x10, 0x52, 0xd0, 0xf2, 0x4f, 0x3b, 0x9d, 0xf3, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x21, 0x68, 0x14, 0x5f, 0x46, 0x02, 0x00, 0x00,
}
