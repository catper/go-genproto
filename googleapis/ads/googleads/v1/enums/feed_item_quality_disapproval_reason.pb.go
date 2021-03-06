// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/feed_item_quality_disapproval_reason.proto

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

// The possible quality evaluation disapproval reasons of a feed item.
type FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason int32

const (
	// No value has been specified.
	FeedItemQualityDisapprovalReasonEnum_UNSPECIFIED FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemQualityDisapprovalReasonEnum_UNKNOWN FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 1
	// Price contains repetitive headers.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_REPETITIVE_HEADERS FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 2
	// Price contains repetitive description.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_REPETITIVE_DESCRIPTION FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 3
	// Price contains inconsistent items.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_INCONSISTENT_ROWS FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 4
	// Price contains qualifiers in description.
	FeedItemQualityDisapprovalReasonEnum_PRICE_DESCRIPTION_HAS_PRICE_QUALIFIERS FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 5
	// Price contains an unsupported language.
	FeedItemQualityDisapprovalReasonEnum_PRICE_UNSUPPORTED_LANGUAGE FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 6
	// Price item header is not relevant to the price type.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_ROW_HEADER_TABLE_TYPE_MISMATCH FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 7
	// Price item header has promotional text.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_ROW_HEADER_HAS_PROMOTIONAL_TEXT FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 8
	// Price item description is not relevant to the item header.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_ROW_DESCRIPTION_NOT_RELEVANT FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 9
	// Price item description contains promotional text.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_ROW_DESCRIPTION_HAS_PROMOTIONAL_TEXT FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 10
	// Price item header and description are repetitive.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_ROW_HEADER_DESCRIPTION_REPETITIVE FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 11
	// Price item is in a foreign language, nonsense, or can't be rated.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_ROW_UNRATEABLE FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 12
	// Price item price is invalid or inaccurate.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_ROW_PRICE_INVALID FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 13
	// Price item URL is invalid or irrelevant.
	FeedItemQualityDisapprovalReasonEnum_PRICE_TABLE_ROW_URL_INVALID FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 14
	// Price item header or description has price.
	FeedItemQualityDisapprovalReasonEnum_PRICE_HEADER_OR_DESCRIPTION_HAS_PRICE FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 15
	// Structured snippet values do not match the header.
	FeedItemQualityDisapprovalReasonEnum_STRUCTURED_SNIPPETS_HEADER_POLICY_VIOLATED FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 16
	// Structured snippet values are repeated.
	FeedItemQualityDisapprovalReasonEnum_STRUCTURED_SNIPPETS_REPEATED_VALUES FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 17
	// Structured snippet values violate editorial guidelines like punctuation.
	FeedItemQualityDisapprovalReasonEnum_STRUCTURED_SNIPPETS_EDITORIAL_GUIDELINES FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 18
	// Structured snippet contain promotional text.
	FeedItemQualityDisapprovalReasonEnum_STRUCTURED_SNIPPETS_HAS_PROMOTIONAL_TEXT FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason = 19
)

var FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "PRICE_TABLE_REPETITIVE_HEADERS",
	3:  "PRICE_TABLE_REPETITIVE_DESCRIPTION",
	4:  "PRICE_TABLE_INCONSISTENT_ROWS",
	5:  "PRICE_DESCRIPTION_HAS_PRICE_QUALIFIERS",
	6:  "PRICE_UNSUPPORTED_LANGUAGE",
	7:  "PRICE_TABLE_ROW_HEADER_TABLE_TYPE_MISMATCH",
	8:  "PRICE_TABLE_ROW_HEADER_HAS_PROMOTIONAL_TEXT",
	9:  "PRICE_TABLE_ROW_DESCRIPTION_NOT_RELEVANT",
	10: "PRICE_TABLE_ROW_DESCRIPTION_HAS_PROMOTIONAL_TEXT",
	11: "PRICE_TABLE_ROW_HEADER_DESCRIPTION_REPETITIVE",
	12: "PRICE_TABLE_ROW_UNRATEABLE",
	13: "PRICE_TABLE_ROW_PRICE_INVALID",
	14: "PRICE_TABLE_ROW_URL_INVALID",
	15: "PRICE_HEADER_OR_DESCRIPTION_HAS_PRICE",
	16: "STRUCTURED_SNIPPETS_HEADER_POLICY_VIOLATED",
	17: "STRUCTURED_SNIPPETS_REPEATED_VALUES",
	18: "STRUCTURED_SNIPPETS_EDITORIAL_GUIDELINES",
	19: "STRUCTURED_SNIPPETS_HAS_PROMOTIONAL_TEXT",
}

var FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason_value = map[string]int32{
	"UNSPECIFIED":                                      0,
	"UNKNOWN":                                          1,
	"PRICE_TABLE_REPETITIVE_HEADERS":                   2,
	"PRICE_TABLE_REPETITIVE_DESCRIPTION":               3,
	"PRICE_TABLE_INCONSISTENT_ROWS":                    4,
	"PRICE_DESCRIPTION_HAS_PRICE_QUALIFIERS":           5,
	"PRICE_UNSUPPORTED_LANGUAGE":                       6,
	"PRICE_TABLE_ROW_HEADER_TABLE_TYPE_MISMATCH":       7,
	"PRICE_TABLE_ROW_HEADER_HAS_PROMOTIONAL_TEXT":      8,
	"PRICE_TABLE_ROW_DESCRIPTION_NOT_RELEVANT":         9,
	"PRICE_TABLE_ROW_DESCRIPTION_HAS_PROMOTIONAL_TEXT": 10,
	"PRICE_TABLE_ROW_HEADER_DESCRIPTION_REPETITIVE":    11,
	"PRICE_TABLE_ROW_UNRATEABLE":                       12,
	"PRICE_TABLE_ROW_PRICE_INVALID":                    13,
	"PRICE_TABLE_ROW_URL_INVALID":                      14,
	"PRICE_HEADER_OR_DESCRIPTION_HAS_PRICE":            15,
	"STRUCTURED_SNIPPETS_HEADER_POLICY_VIOLATED":       16,
	"STRUCTURED_SNIPPETS_REPEATED_VALUES":              17,
	"STRUCTURED_SNIPPETS_EDITORIAL_GUIDELINES":         18,
	"STRUCTURED_SNIPPETS_HAS_PROMOTIONAL_TEXT":         19,
}

func (x FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason) String() string {
	return proto.EnumName(FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason_name, int32(x))
}

func (FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f17372a6460ec60d, []int{0, 0}
}

// Container for enum describing possible quality evaluation disapproval reasons
// of a feed item.
type FeedItemQualityDisapprovalReasonEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemQualityDisapprovalReasonEnum) Reset()         { *m = FeedItemQualityDisapprovalReasonEnum{} }
func (m *FeedItemQualityDisapprovalReasonEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemQualityDisapprovalReasonEnum) ProtoMessage()    {}
func (*FeedItemQualityDisapprovalReasonEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f17372a6460ec60d, []int{0}
}

func (m *FeedItemQualityDisapprovalReasonEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemQualityDisapprovalReasonEnum.Unmarshal(m, b)
}
func (m *FeedItemQualityDisapprovalReasonEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemQualityDisapprovalReasonEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemQualityDisapprovalReasonEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemQualityDisapprovalReasonEnum.Merge(m, src)
}
func (m *FeedItemQualityDisapprovalReasonEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemQualityDisapprovalReasonEnum.Size(m)
}
func (m *FeedItemQualityDisapprovalReasonEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemQualityDisapprovalReasonEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemQualityDisapprovalReasonEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason", FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason_name, FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason_value)
	proto.RegisterType((*FeedItemQualityDisapprovalReasonEnum)(nil), "google.ads.googleads.v1.enums.FeedItemQualityDisapprovalReasonEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/feed_item_quality_disapproval_reason.proto", fileDescriptor_f17372a6460ec60d)
}

var fileDescriptor_f17372a6460ec60d = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcd, 0x72, 0xd3, 0x30,
	0x10, 0xc7, 0x69, 0x0a, 0x29, 0xa8, 0x40, 0x85, 0xb8, 0x15, 0x5a, 0x68, 0xa0, 0x05, 0x0a, 0xd8,
	0x64, 0xe0, 0x64, 0x4e, 0x8a, 0xad, 0x26, 0x1a, 0x5c, 0xd9, 0x95, 0x64, 0x97, 0x32, 0x99, 0xd1,
	0x18, 0x6c, 0x32, 0x99, 0x49, 0xec, 0x10, 0xa7, 0x9d, 0xe1, 0x75, 0x38, 0xf2, 0x02, 0xbc, 0x03,
	0x4f, 0xc1, 0x99, 0x47, 0xe0, 0xc4, 0xd8, 0xca, 0x17, 0x25, 0x69, 0x2f, 0x9e, 0xd5, 0xee, 0x6f,
	0x77, 0xff, 0x5e, 0xcd, 0x0a, 0xb4, 0x3a, 0x59, 0xd6, 0xe9, 0x25, 0x66, 0x14, 0xe7, 0xa6, 0x36,
	0x0b, 0xeb, 0xac, 0x6e, 0x26, 0xe9, 0x69, 0x3f, 0x37, 0x3f, 0x27, 0x49, 0xac, 0xba, 0xa3, 0xa4,
	0xaf, 0xbe, 0x9c, 0x46, 0xbd, 0xee, 0xe8, 0xab, 0x8a, 0xbb, 0x79, 0x34, 0x18, 0x0c, 0xb3, 0xb3,
	0xa8, 0xa7, 0x86, 0x49, 0x94, 0x67, 0xa9, 0x31, 0x18, 0x66, 0xa3, 0x0c, 0x6d, 0xe9, 0x74, 0x23,
	0x8a, 0x73, 0x63, 0x5a, 0xc9, 0x38, 0xab, 0x1b, 0x65, 0xa5, 0xcd, 0xfb, 0x93, 0x46, 0x83, 0xae,
	0x19, 0xa5, 0x69, 0x36, 0x8a, 0x46, 0xdd, 0x2c, 0xcd, 0x75, 0x72, 0xed, 0x57, 0x15, 0x3c, 0x3e,
	0x48, 0x92, 0x98, 0x8e, 0x92, 0xfe, 0x91, 0xee, 0xe4, 0xcc, 0x1a, 0xf1, 0xb2, 0x0f, 0x49, 0x4f,
	0xfb, 0xb5, 0x1f, 0x55, 0xf0, 0xf0, 0x32, 0x10, 0x6d, 0x80, 0xf5, 0x80, 0x09, 0x9f, 0xd8, 0xf4,
	0x80, 0x12, 0x07, 0x5e, 0x41, 0xeb, 0x60, 0x2d, 0x60, 0xef, 0x98, 0x77, 0xcc, 0xe0, 0x0a, 0xaa,
	0x81, 0x6d, 0x9f, 0x53, 0x9b, 0x28, 0x89, 0x1b, 0x2e, 0x51, 0x9c, 0xf8, 0x44, 0x52, 0x49, 0x43,
	0xa2, 0x5a, 0x04, 0x3b, 0x84, 0x0b, 0x58, 0x41, 0x7b, 0xa0, 0xb6, 0x84, 0x71, 0x88, 0xb0, 0x39,
	0xf5, 0x25, 0xf5, 0x18, 0x5c, 0x45, 0x3b, 0x60, 0x6b, 0x9e, 0xa3, 0xcc, 0xf6, 0x98, 0xa0, 0x42,
	0x12, 0x26, 0x15, 0xf7, 0x8e, 0x05, 0xbc, 0x8a, 0xf6, 0xc1, 0x9e, 0x46, 0xe6, 0x32, 0x55, 0x0b,
	0x0b, 0xa5, 0xbd, 0x47, 0x01, 0x76, 0x0b, 0x99, 0x5c, 0xc0, 0x6b, 0x68, 0x1b, 0x6c, 0x6a, 0x6f,
	0xc0, 0x44, 0xe0, 0xfb, 0x1e, 0x97, 0xc4, 0x51, 0x2e, 0x66, 0xcd, 0x00, 0x37, 0x09, 0xac, 0x22,
	0x03, 0xec, 0xff, 0x23, 0xcb, 0x3b, 0x1e, 0x6b, 0x1e, 0x3b, 0xe4, 0x89, 0x4f, 0xd4, 0x21, 0x15,
	0x87, 0x58, 0xda, 0x2d, 0xb8, 0x86, 0x4c, 0xf0, 0x7c, 0x09, 0xaf, 0x05, 0x78, 0x87, 0x5e, 0x21,
	0x07, 0xbb, 0x4a, 0x92, 0xf7, 0x12, 0x5e, 0x47, 0x2f, 0xc0, 0xd3, 0xf3, 0x09, 0xf3, 0xb2, 0x99,
	0x27, 0x15, 0x27, 0x2e, 0x09, 0x31, 0x93, 0xf0, 0x06, 0x7a, 0x03, 0x5e, 0x5d, 0x44, 0x2f, 0xec,
	0x01, 0x50, 0x1d, 0xbc, 0x5c, 0x22, 0x6a, 0x3e, 0x79, 0x36, 0x72, 0xb8, 0x3e, 0x9b, 0xcb, 0x2c,
	0x25, 0x60, 0x1c, 0x4b, 0x52, 0x1c, 0xe1, 0xcd, 0xf3, 0xd7, 0x50, 0xc4, 0xf5, 0x99, 0xb2, 0x10,
	0xbb, 0xd4, 0x81, 0xb7, 0xd0, 0x03, 0x70, 0xef, 0xbf, 0x12, 0xdc, 0x9d, 0x02, 0xb7, 0xd1, 0x33,
	0xb0, 0xab, 0x81, 0xb1, 0x18, 0x8f, 0x2f, 0xbe, 0x31, 0xb8, 0x51, 0x5c, 0x83, 0x90, 0x3c, 0xb0,
	0x65, 0xc0, 0x89, 0xa3, 0x04, 0xa3, 0xbe, 0x4f, 0xa4, 0x98, 0x24, 0xfa, 0x9e, 0x4b, 0xed, 0x13,
	0x15, 0x52, 0xcf, 0xc5, 0x92, 0x38, 0x10, 0xa2, 0x27, 0xe0, 0xd1, 0x22, 0xbe, 0xf8, 0xc5, 0x82,
	0x50, 0x21, 0x76, 0x03, 0x22, 0xe0, 0x9d, 0x62, 0xfc, 0x8b, 0x40, 0xe2, 0x50, 0xe9, 0x71, 0x8a,
	0x5d, 0xd5, 0x0c, 0xa8, 0x43, 0x5c, 0xca, 0x88, 0x80, 0x68, 0x19, 0xbd, 0x70, 0xec, 0x77, 0x1b,
	0x7f, 0x56, 0xc0, 0xce, 0xa7, 0xac, 0x6f, 0x5c, 0xb8, 0xa6, 0x8d, 0xdd, 0xcb, 0x96, 0xcb, 0x2f,
	0xf6, 0xd5, 0x5f, 0xf9, 0xd0, 0x18, 0xd7, 0xe9, 0x64, 0xbd, 0x28, 0xed, 0x18, 0xd9, 0xb0, 0x63,
	0x76, 0x92, 0xb4, 0xdc, 0xe6, 0xc9, 0x43, 0x32, 0xe8, 0xe6, 0x4b, 0xde, 0x95, 0xb7, 0xe5, 0xf7,
	0x5b, 0x65, 0xb5, 0x89, 0xf1, 0xf7, 0xca, 0x56, 0x53, 0x97, 0xc2, 0x71, 0x6e, 0x68, 0xb3, 0xb0,
	0xc2, 0xba, 0x51, 0x6c, 0x7c, 0xfe, 0x73, 0x12, 0x6f, 0xe3, 0x38, 0x6f, 0x4f, 0xe3, 0xed, 0xb0,
	0xde, 0x2e, 0xe3, 0xbf, 0x2b, 0x3b, 0xda, 0x69, 0x59, 0x38, 0xce, 0x2d, 0x6b, 0x4a, 0x58, 0x56,
	0x58, 0xb7, 0xac, 0x92, 0xf9, 0x58, 0x2d, 0x85, 0xbd, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0xcf, 0x49, 0x08, 0xef, 0x04, 0x00, 0x00,
}
