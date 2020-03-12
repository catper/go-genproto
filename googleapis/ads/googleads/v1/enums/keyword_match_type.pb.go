// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/keyword_match_type.proto

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

// Possible Keyword match types.
type KeywordMatchTypeEnum_KeywordMatchType int32

const (
	// Not specified.
	KeywordMatchTypeEnum_UNSPECIFIED KeywordMatchTypeEnum_KeywordMatchType = 0
	// Used for return value only. Represents value unknown in this version.
	KeywordMatchTypeEnum_UNKNOWN KeywordMatchTypeEnum_KeywordMatchType = 1
	// Exact match.
	KeywordMatchTypeEnum_EXACT KeywordMatchTypeEnum_KeywordMatchType = 2
	// Phrase match.
	KeywordMatchTypeEnum_PHRASE KeywordMatchTypeEnum_KeywordMatchType = 3
	// Broad match.
	KeywordMatchTypeEnum_BROAD KeywordMatchTypeEnum_KeywordMatchType = 4
)

var KeywordMatchTypeEnum_KeywordMatchType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "EXACT",
	3: "PHRASE",
	4: "BROAD",
}

var KeywordMatchTypeEnum_KeywordMatchType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"EXACT":       2,
	"PHRASE":      3,
	"BROAD":       4,
}

func (x KeywordMatchTypeEnum_KeywordMatchType) String() string {
	return proto.EnumName(KeywordMatchTypeEnum_KeywordMatchType_name, int32(x))
}

func (KeywordMatchTypeEnum_KeywordMatchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb7e96826193c362, []int{0, 0}
}

// Message describing Keyword match types.
type KeywordMatchTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordMatchTypeEnum) Reset()         { *m = KeywordMatchTypeEnum{} }
func (m *KeywordMatchTypeEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordMatchTypeEnum) ProtoMessage()    {}
func (*KeywordMatchTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb7e96826193c362, []int{0}
}

func (m *KeywordMatchTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordMatchTypeEnum.Unmarshal(m, b)
}
func (m *KeywordMatchTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordMatchTypeEnum.Marshal(b, m, deterministic)
}
func (m *KeywordMatchTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordMatchTypeEnum.Merge(m, src)
}
func (m *KeywordMatchTypeEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordMatchTypeEnum.Size(m)
}
func (m *KeywordMatchTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordMatchTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordMatchTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.KeywordMatchTypeEnum_KeywordMatchType", KeywordMatchTypeEnum_KeywordMatchType_name, KeywordMatchTypeEnum_KeywordMatchType_value)
	proto.RegisterType((*KeywordMatchTypeEnum)(nil), "google.ads.googleads.v1.enums.KeywordMatchTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/keyword_match_type.proto", fileDescriptor_bb7e96826193c362)
}

var fileDescriptor_bb7e96826193c362 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4e, 0xf2, 0x30,
	0x18, 0xfe, 0x18, 0x9f, 0x18, 0xcb, 0x81, 0xcb, 0xa2, 0x27, 0x46, 0x0e, 0xe0, 0x02, 0xba, 0x2c,
	0x26, 0x1e, 0xd4, 0xa3, 0x0e, 0x26, 0x12, 0xe2, 0x58, 0xf8, 0xd3, 0x98, 0x25, 0xa4, 0xd2, 0xa6,
	0xa2, 0xac, 0x5d, 0xe8, 0xc0, 0x70, 0x3b, 0x1e, 0x7a, 0x29, 0x5e, 0x0a, 0x57, 0x61, 0xda, 0xca,
	0x0e, 0x48, 0xf4, 0xa4, 0x79, 0xf2, 0x3e, 0x3f, 0x7d, 0xde, 0x17, 0x5c, 0x73, 0x29, 0xf9, 0x92,
	0xf9, 0x84, 0x2a, 0xdf, 0x42, 0x8d, 0x36, 0x81, 0xcf, 0xc4, 0x3a, 0x53, 0xfe, 0x1b, 0xdb, 0xbe,
	0xcb, 0x15, 0x9d, 0x65, 0xa4, 0x98, 0xbf, 0xcc, 0x8a, 0x6d, 0xce, 0x60, 0xbe, 0x92, 0x85, 0xf4,
	0x1a, 0x56, 0x0c, 0x09, 0x55, 0xb0, 0xf4, 0xc1, 0x4d, 0x00, 0x8d, 0xef, 0xe2, 0x72, 0x1f, 0x9b,
	0x2f, 0x7c, 0x22, 0x84, 0x2c, 0x48, 0xb1, 0x90, 0x42, 0x59, 0x73, 0xeb, 0x15, 0x9c, 0xf5, 0x6d,
	0xf0, 0xbd, 0xce, 0x1d, 0x6f, 0x73, 0x16, 0x89, 0x75, 0xd6, 0x1a, 0x02, 0xf7, 0x70, 0xee, 0x9d,
	0x82, 0xfa, 0x24, 0x1e, 0x25, 0x51, 0xbb, 0x77, 0xdb, 0x8b, 0x3a, 0xee, 0x3f, 0xaf, 0x0e, 0x8e,
	0x27, 0x71, 0x3f, 0x1e, 0x3c, 0xc4, 0x6e, 0xc5, 0x3b, 0x01, 0x47, 0xd1, 0x23, 0x6e, 0x8f, 0x5d,
	0xc7, 0x03, 0xa0, 0x96, 0xdc, 0x0d, 0xf1, 0x28, 0x72, 0xab, 0x7a, 0x1c, 0x0e, 0x07, 0xb8, 0xe3,
	0xfe, 0x0f, 0x77, 0x15, 0xd0, 0x9c, 0xcb, 0x0c, 0xfe, 0xd9, 0x37, 0x3c, 0x3f, 0xfc, 0x37, 0xd1,
	0x45, 0x93, 0xca, 0x53, 0xf8, 0xe3, 0xe3, 0x72, 0x49, 0x04, 0x87, 0x72, 0xc5, 0x7d, 0xce, 0x84,
	0x59, 0x63, 0x7f, 0xaf, 0x7c, 0xa1, 0x7e, 0x39, 0xdf, 0x8d, 0x79, 0x3f, 0x9c, 0x6a, 0x17, 0xe3,
	0x4f, 0xa7, 0xd1, 0xb5, 0x51, 0x98, 0x2a, 0x68, 0xa1, 0x46, 0xd3, 0x00, 0xea, 0xdd, 0xd5, 0xd7,
	0x9e, 0x4f, 0x31, 0x55, 0x69, 0xc9, 0xa7, 0xd3, 0x20, 0x35, 0xfc, 0xce, 0x69, 0xda, 0x21, 0x42,
	0x98, 0x2a, 0x84, 0x4a, 0x05, 0x42, 0xd3, 0x00, 0x21, 0xa3, 0x79, 0xae, 0x99, 0x62, 0x57, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x57, 0x9e, 0x75, 0xf2, 0xd6, 0x01, 0x00, 0x00,
}
