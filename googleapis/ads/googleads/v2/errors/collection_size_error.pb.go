// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/collection_size_error.proto

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

// Enum describing possible collection size errors.
type CollectionSizeErrorEnum_CollectionSizeError int32

const (
	// Enum unspecified.
	CollectionSizeErrorEnum_UNSPECIFIED CollectionSizeErrorEnum_CollectionSizeError = 0
	// The received error code is not known in this version.
	CollectionSizeErrorEnum_UNKNOWN CollectionSizeErrorEnum_CollectionSizeError = 1
	// Too few.
	CollectionSizeErrorEnum_TOO_FEW CollectionSizeErrorEnum_CollectionSizeError = 2
	// Too many.
	CollectionSizeErrorEnum_TOO_MANY CollectionSizeErrorEnum_CollectionSizeError = 3
)

var CollectionSizeErrorEnum_CollectionSizeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "TOO_FEW",
	3: "TOO_MANY",
}

var CollectionSizeErrorEnum_CollectionSizeError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"TOO_FEW":     2,
	"TOO_MANY":    3,
}

func (x CollectionSizeErrorEnum_CollectionSizeError) String() string {
	return proto.EnumName(CollectionSizeErrorEnum_CollectionSizeError_name, int32(x))
}

func (CollectionSizeErrorEnum_CollectionSizeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8be19f13e849550f, []int{0, 0}
}

// Container for enum describing possible collection size errors.
type CollectionSizeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionSizeErrorEnum) Reset()         { *m = CollectionSizeErrorEnum{} }
func (m *CollectionSizeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CollectionSizeErrorEnum) ProtoMessage()    {}
func (*CollectionSizeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8be19f13e849550f, []int{0}
}

func (m *CollectionSizeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionSizeErrorEnum.Unmarshal(m, b)
}
func (m *CollectionSizeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionSizeErrorEnum.Marshal(b, m, deterministic)
}
func (m *CollectionSizeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionSizeErrorEnum.Merge(m, src)
}
func (m *CollectionSizeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CollectionSizeErrorEnum.Size(m)
}
func (m *CollectionSizeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionSizeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionSizeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.CollectionSizeErrorEnum_CollectionSizeError", CollectionSizeErrorEnum_CollectionSizeError_name, CollectionSizeErrorEnum_CollectionSizeError_value)
	proto.RegisterType((*CollectionSizeErrorEnum)(nil), "google.ads.googleads.v2.errors.CollectionSizeErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/collection_size_error.proto", fileDescriptor_8be19f13e849550f)
}

var fileDescriptor_8be19f13e849550f = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x23, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0xe4, 0xfc, 0x9c, 0x9c, 0xd4, 0xe4, 0x92, 0xcc, 0xfc, 0xbc, 0xf8,
	0xe2, 0xcc, 0xaa, 0xd4, 0x78, 0xb0, 0xb0, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x1c, 0x44,
	0x83, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xaf, 0x5e, 0x99, 0x91, 0x1e, 0x44, 0xaf, 0x94, 0x0c,
	0xcc, 0xec, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0x44, 0x90, 0x39, 0xc5, 0x10, 0xdd,
	0x4a, 0x99, 0x5c, 0xe2, 0xce, 0x70, 0xc3, 0x83, 0x33, 0xab, 0x52, 0x5d, 0x41, 0xba, 0x5c, 0xf3,
	0x4a, 0x73, 0x95, 0xfc, 0xb8, 0x84, 0xb1, 0x48, 0x09, 0xf1, 0x73, 0x71, 0x87, 0xfa, 0x05, 0x07,
	0xb8, 0x3a, 0x7b, 0xba, 0x79, 0xba, 0xba, 0x08, 0x30, 0x08, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79,
	0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x82, 0x38, 0x21, 0xfe, 0xfe, 0xf1, 0x6e, 0xae, 0xe1, 0x02,
	0x4c, 0x42, 0x3c, 0x5c, 0x1c, 0x20, 0x8e, 0xaf, 0xa3, 0x5f, 0xa4, 0x00, 0xb3, 0xd3, 0x67, 0x46,
	0x2e, 0xa5, 0xe4, 0xfc, 0x5c, 0x3d, 0xfc, 0xee, 0x75, 0x92, 0xc0, 0x62, 0x69, 0x00, 0xc8, 0xad,
	0x01, 0x8c, 0x51, 0x2e, 0x50, 0xbd, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9,
	0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x9f, 0xc0, 0xc2, 0xad, 0x20, 0xb3, 0x18, 0x57, 0x30, 0x5a, 0x43,
	0xa8, 0x45, 0x4c, 0xcc, 0xee, 0x8e, 0x8e, 0xab, 0x98, 0xe4, 0xdc, 0x21, 0x86, 0x39, 0xa6, 0x14,
	0xeb, 0x41, 0x98, 0x20, 0x56, 0x98, 0x91, 0x1e, 0xd8, 0xca, 0xe2, 0x53, 0x30, 0x05, 0x31, 0x8e,
	0x29, 0xc5, 0x31, 0x70, 0x05, 0x31, 0x61, 0x46, 0x31, 0x10, 0x05, 0xaf, 0x98, 0x94, 0x20, 0xa2,
	0x56, 0x56, 0x8e, 0x29, 0xc5, 0x56, 0x56, 0x70, 0x25, 0x56, 0x56, 0x61, 0x46, 0x56, 0x56, 0x10,
	0x45, 0x49, 0x6c, 0x60, 0xd7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x53, 0x0f, 0x9f,
	0xe3, 0x01, 0x00, 0x00,
}
