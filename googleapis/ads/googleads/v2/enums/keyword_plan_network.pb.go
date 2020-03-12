// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/keyword_plan_network.proto

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

// Enumerates keyword plan forecastable network types.
type KeywordPlanNetworkEnum_KeywordPlanNetwork int32

const (
	// Not specified.
	KeywordPlanNetworkEnum_UNSPECIFIED KeywordPlanNetworkEnum_KeywordPlanNetwork = 0
	// The value is unknown in this version.
	KeywordPlanNetworkEnum_UNKNOWN KeywordPlanNetworkEnum_KeywordPlanNetwork = 1
	// Google Search.
	KeywordPlanNetworkEnum_GOOGLE_SEARCH KeywordPlanNetworkEnum_KeywordPlanNetwork = 2
	// Google Search + Search partners.
	KeywordPlanNetworkEnum_GOOGLE_SEARCH_AND_PARTNERS KeywordPlanNetworkEnum_KeywordPlanNetwork = 3
)

var KeywordPlanNetworkEnum_KeywordPlanNetwork_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "GOOGLE_SEARCH",
	3: "GOOGLE_SEARCH_AND_PARTNERS",
}

var KeywordPlanNetworkEnum_KeywordPlanNetwork_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"GOOGLE_SEARCH":              2,
	"GOOGLE_SEARCH_AND_PARTNERS": 3,
}

func (x KeywordPlanNetworkEnum_KeywordPlanNetwork) String() string {
	return proto.EnumName(KeywordPlanNetworkEnum_KeywordPlanNetwork_name, int32(x))
}

func (KeywordPlanNetworkEnum_KeywordPlanNetwork) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2481e5f0090e92fb, []int{0, 0}
}

// Container for enumeration of keyword plan forecastable network types.
type KeywordPlanNetworkEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanNetworkEnum) Reset()         { *m = KeywordPlanNetworkEnum{} }
func (m *KeywordPlanNetworkEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanNetworkEnum) ProtoMessage()    {}
func (*KeywordPlanNetworkEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_2481e5f0090e92fb, []int{0}
}

func (m *KeywordPlanNetworkEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Unmarshal(m, b)
}
func (m *KeywordPlanNetworkEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanNetworkEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanNetworkEnum.Merge(m, src)
}
func (m *KeywordPlanNetworkEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Size(m)
}
func (m *KeywordPlanNetworkEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanNetworkEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanNetworkEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork", KeywordPlanNetworkEnum_KeywordPlanNetwork_name, KeywordPlanNetworkEnum_KeywordPlanNetwork_value)
	proto.RegisterType((*KeywordPlanNetworkEnum)(nil), "google.ads.googleads.v2.enums.KeywordPlanNetworkEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/keyword_plan_network.proto", fileDescriptor_2481e5f0090e92fb)
}

var fileDescriptor_2481e5f0090e92fb = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xf3, 0x30,
	0x1c, 0xfd, 0xd6, 0xc1, 0x27, 0x64, 0x88, 0xb5, 0x17, 0x0a, 0xc3, 0x09, 0xdb, 0x03, 0xa4, 0x50,
	0x6f, 0x24, 0x5e, 0x65, 0x5b, 0xad, 0x63, 0x92, 0x95, 0xcd, 0x4d, 0x90, 0x42, 0x89, 0x26, 0x84,
	0xb1, 0x2e, 0x29, 0x4d, 0xb7, 0xe1, 0x95, 0xef, 0xe2, 0xa5, 0x8f, 0xe2, 0xa3, 0x88, 0x0f, 0x21,
	0x4d, 0xdc, 0x40, 0x86, 0xde, 0x84, 0x43, 0xce, 0x1f, 0xce, 0xef, 0x80, 0x4b, 0xa1, 0x94, 0xc8,
	0xb8, 0x4f, 0x99, 0xf6, 0x2d, 0xac, 0xd0, 0x3a, 0xf0, 0xb9, 0x5c, 0x2d, 0xb5, 0xbf, 0xe0, 0xcf,
	0x1b, 0x55, 0xb0, 0x34, 0xcf, 0xa8, 0x4c, 0x25, 0x2f, 0x37, 0xaa, 0x58, 0xc0, 0xbc, 0x50, 0xa5,
	0xf2, 0x5a, 0x56, 0x0e, 0x29, 0xd3, 0x70, 0xe7, 0x84, 0xeb, 0x00, 0x1a, 0x67, 0xf3, 0x6c, 0x1b,
	0x9c, 0xcf, 0x7d, 0x2a, 0xa5, 0x2a, 0x69, 0x39, 0x57, 0x52, 0x5b, 0x73, 0xe7, 0x05, 0x9c, 0x0c,
	0x6d, 0x74, 0x9c, 0x51, 0x49, 0x6c, 0x70, 0x28, 0x57, 0xcb, 0x0e, 0x07, 0xde, 0x3e, 0xe3, 0x1d,
	0x81, 0xc6, 0x94, 0x4c, 0xe2, 0xb0, 0x37, 0xb8, 0x1e, 0x84, 0x7d, 0xf7, 0x9f, 0xd7, 0x00, 0x07,
	0x53, 0x32, 0x24, 0xa3, 0x7b, 0xe2, 0xd6, 0xbc, 0x63, 0x70, 0x18, 0x8d, 0x46, 0xd1, 0x6d, 0x98,
	0x4e, 0x42, 0x3c, 0xee, 0xdd, 0xb8, 0x8e, 0x77, 0x0e, 0x9a, 0x3f, 0xbe, 0x52, 0x4c, 0xfa, 0x69,
	0x8c, 0xc7, 0x77, 0x24, 0x1c, 0x4f, 0xdc, 0x7a, 0xf7, 0xb3, 0x06, 0xda, 0x4f, 0x6a, 0x09, 0xff,
	0x3c, 0xa2, 0x7b, 0xba, 0x5f, 0x25, 0xae, 0xfa, 0xc7, 0xb5, 0x87, 0xee, 0xb7, 0x53, 0xa8, 0x8c,
	0x4a, 0x01, 0x55, 0x21, 0x7c, 0xc1, 0xa5, 0xb9, 0x6e, 0x3b, 0x64, 0x3e, 0xd7, 0xbf, 0xec, 0x7a,
	0x65, 0xde, 0x57, 0xa7, 0x1e, 0x61, 0xfc, 0xe6, 0xb4, 0x22, 0x1b, 0x85, 0x99, 0x86, 0x16, 0x56,
	0x68, 0x16, 0xc0, 0x6a, 0x10, 0xfd, 0xbe, 0xe5, 0x13, 0xcc, 0x74, 0xb2, 0xe3, 0x93, 0x59, 0x90,
	0x18, 0xfe, 0xc3, 0x69, 0xdb, 0x4f, 0x84, 0x30, 0xd3, 0x08, 0xed, 0x14, 0x08, 0xcd, 0x02, 0x84,
	0x8c, 0xe6, 0xf1, 0xbf, 0x29, 0x76, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0x64, 0xe6, 0x50, 0x4e,
	0xef, 0x01, 0x00, 0x00,
}
