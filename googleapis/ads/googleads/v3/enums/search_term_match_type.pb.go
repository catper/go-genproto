// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/search_term_match_type.proto

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

// Possible match types for a keyword triggering an ad, including variants.
type SearchTermMatchTypeEnum_SearchTermMatchType int32

const (
	// Not specified.
	SearchTermMatchTypeEnum_UNSPECIFIED SearchTermMatchTypeEnum_SearchTermMatchType = 0
	// Used for return value only. Represents value unknown in this version.
	SearchTermMatchTypeEnum_UNKNOWN SearchTermMatchTypeEnum_SearchTermMatchType = 1
	// Broad match.
	SearchTermMatchTypeEnum_BROAD SearchTermMatchTypeEnum_SearchTermMatchType = 2
	// Exact match.
	SearchTermMatchTypeEnum_EXACT SearchTermMatchTypeEnum_SearchTermMatchType = 3
	// Phrase match.
	SearchTermMatchTypeEnum_PHRASE SearchTermMatchTypeEnum_SearchTermMatchType = 4
	// Exact match (close variant).
	SearchTermMatchTypeEnum_NEAR_EXACT SearchTermMatchTypeEnum_SearchTermMatchType = 5
	// Phrase match (close variant).
	SearchTermMatchTypeEnum_NEAR_PHRASE SearchTermMatchTypeEnum_SearchTermMatchType = 6
)

var SearchTermMatchTypeEnum_SearchTermMatchType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BROAD",
	3: "EXACT",
	4: "PHRASE",
	5: "NEAR_EXACT",
	6: "NEAR_PHRASE",
}

var SearchTermMatchTypeEnum_SearchTermMatchType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"BROAD":       2,
	"EXACT":       3,
	"PHRASE":      4,
	"NEAR_EXACT":  5,
	"NEAR_PHRASE": 6,
}

func (x SearchTermMatchTypeEnum_SearchTermMatchType) String() string {
	return proto.EnumName(SearchTermMatchTypeEnum_SearchTermMatchType_name, int32(x))
}

func (SearchTermMatchTypeEnum_SearchTermMatchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_37984e6ef3c9b9dd, []int{0, 0}
}

// Container for enum describing match types for a keyword triggering an ad.
type SearchTermMatchTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchTermMatchTypeEnum) Reset()         { *m = SearchTermMatchTypeEnum{} }
func (m *SearchTermMatchTypeEnum) String() string { return proto.CompactTextString(m) }
func (*SearchTermMatchTypeEnum) ProtoMessage()    {}
func (*SearchTermMatchTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_37984e6ef3c9b9dd, []int{0}
}

func (m *SearchTermMatchTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Unmarshal(m, b)
}
func (m *SearchTermMatchTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Marshal(b, m, deterministic)
}
func (m *SearchTermMatchTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermMatchTypeEnum.Merge(m, src)
}
func (m *SearchTermMatchTypeEnum) XXX_Size() int {
	return xxx_messageInfo_SearchTermMatchTypeEnum.Size(m)
}
func (m *SearchTermMatchTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermMatchTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermMatchTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.SearchTermMatchTypeEnum_SearchTermMatchType", SearchTermMatchTypeEnum_SearchTermMatchType_name, SearchTermMatchTypeEnum_SearchTermMatchType_value)
	proto.RegisterType((*SearchTermMatchTypeEnum)(nil), "google.ads.googleads.v3.enums.SearchTermMatchTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/search_term_match_type.proto", fileDescriptor_37984e6ef3c9b9dd)
}

var fileDescriptor_37984e6ef3c9b9dd = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x18, 0xc5, 0x6d, 0xe7, 0x26, 0x66, 0xa0, 0xa5, 0x5e, 0x28, 0xe2, 0x2e, 0xb6, 0x07, 0x48, 0x2f,
	0x7a, 0x17, 0xaf, 0xd2, 0xad, 0xce, 0x21, 0x76, 0x65, 0xff, 0x14, 0x29, 0x8c, 0xb8, 0x86, 0x3a,
	0x58, 0x92, 0xd2, 0x74, 0x83, 0x3d, 0x86, 0xaf, 0xe0, 0xa5, 0x8f, 0xe2, 0xa3, 0x08, 0xbe, 0x83,
	0x24, 0xd9, 0x7a, 0x35, 0xbd, 0x29, 0xbf, 0xe6, 0x7c, 0xe7, 0xf0, 0x7d, 0x07, 0xa0, 0x4c, 0x88,
	0x6c, 0x45, 0x3d, 0x92, 0x4a, 0xcf, 0xa0, 0xa2, 0x8d, 0xef, 0x51, 0xbe, 0x66, 0xd2, 0x93, 0x94,
	0x14, 0x8b, 0xb7, 0x79, 0x49, 0x0b, 0x36, 0x67, 0xa4, 0x54, 0xb8, 0xcd, 0x29, 0xcc, 0x0b, 0x51,
	0x0a, 0xb7, 0x65, 0x0c, 0x90, 0xa4, 0x12, 0x56, 0x5e, 0xb8, 0xf1, 0xa1, 0xf6, 0x5e, 0xdf, 0xec,
	0xa3, 0xf3, 0xa5, 0x47, 0x38, 0x17, 0x25, 0x29, 0x97, 0x82, 0x4b, 0x63, 0xee, 0xbc, 0x5b, 0xe0,
	0x72, 0xac, 0xd3, 0x27, 0xb4, 0x60, 0x8f, 0x2a, 0x7b, 0xb2, 0xcd, 0x69, 0xc8, 0xd7, 0xac, 0xb3,
	0x01, 0x17, 0x07, 0x24, 0xf7, 0x1c, 0x34, 0xa7, 0xd1, 0x38, 0x0e, 0xbb, 0x83, 0xbb, 0x41, 0xd8,
	0x73, 0x8e, 0xdc, 0x26, 0x38, 0x99, 0x46, 0x0f, 0xd1, 0xf0, 0x29, 0x72, 0x2c, 0xf7, 0x14, 0xd4,
	0x83, 0xd1, 0x10, 0xf7, 0x1c, 0x5b, 0x61, 0xf8, 0x8c, 0xbb, 0x13, 0xa7, 0xe6, 0x02, 0xd0, 0x88,
	0xef, 0x47, 0x78, 0x1c, 0x3a, 0xc7, 0xee, 0x19, 0x00, 0x51, 0x88, 0x47, 0x73, 0xa3, 0xd5, 0x55,
	0x9e, 0xfe, 0xdf, 0x0d, 0x34, 0x82, 0x1f, 0x0b, 0xb4, 0x17, 0x82, 0xc1, 0x7f, 0xef, 0x0a, 0xae,
	0x0e, 0xec, 0x16, 0xab, 0x9b, 0x62, 0xeb, 0x25, 0xd8, 0x59, 0x33, 0xb1, 0x22, 0x3c, 0x83, 0xa2,
	0xc8, 0xbc, 0x8c, 0x72, 0x7d, 0xf1, 0xbe, 0xde, 0x7c, 0x29, 0xff, 0x68, 0xfb, 0x56, 0x7f, 0x3f,
	0xec, 0x5a, 0x1f, 0xe3, 0x4f, 0xbb, 0xd5, 0x37, 0x51, 0x38, 0x95, 0xd0, 0xa0, 0xa2, 0x99, 0x0f,
	0x55, 0x45, 0xf2, 0x6b, 0xaf, 0x27, 0x38, 0x95, 0x49, 0xa5, 0x27, 0x33, 0x3f, 0xd1, 0xfa, 0xb7,
	0xdd, 0x36, 0x8f, 0x08, 0xe1, 0x54, 0x22, 0x54, 0x4d, 0x20, 0x34, 0xf3, 0x11, 0xd2, 0x33, 0xaf,
	0x0d, 0xbd, 0x98, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x65, 0x16, 0xc4, 0x05, 0x02, 0x00,
	0x00,
}
