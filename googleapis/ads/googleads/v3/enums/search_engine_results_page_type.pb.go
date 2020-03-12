// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/search_engine_results_page_type.proto

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

// The type of the search engine results page.
type SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType int32

const (
	// Not specified.
	SearchEngineResultsPageTypeEnum_UNSPECIFIED SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 0
	// Used for return value only. Represents value unknown in this version.
	SearchEngineResultsPageTypeEnum_UNKNOWN SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 1
	// Only ads were contained in the search engine results page.
	SearchEngineResultsPageTypeEnum_ADS_ONLY SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 2
	// Only organic results were contained in the search engine results page.
	SearchEngineResultsPageTypeEnum_ORGANIC_ONLY SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 3
	// Both ads and organic results were contained in the search engine results
	// page.
	SearchEngineResultsPageTypeEnum_ADS_AND_ORGANIC SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType = 4
)

var SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADS_ONLY",
	3: "ORGANIC_ONLY",
	4: "ADS_AND_ORGANIC",
}

var SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"ADS_ONLY":        2,
	"ORGANIC_ONLY":    3,
	"ADS_AND_ORGANIC": 4,
}

func (x SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType) String() string {
	return proto.EnumName(SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_name, int32(x))
}

func (SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d0c9169c79e638f, []int{0, 0}
}

// The type of the search engine results page.
type SearchEngineResultsPageTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchEngineResultsPageTypeEnum) Reset()         { *m = SearchEngineResultsPageTypeEnum{} }
func (m *SearchEngineResultsPageTypeEnum) String() string { return proto.CompactTextString(m) }
func (*SearchEngineResultsPageTypeEnum) ProtoMessage()    {}
func (*SearchEngineResultsPageTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0c9169c79e638f, []int{0}
}

func (m *SearchEngineResultsPageTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchEngineResultsPageTypeEnum.Unmarshal(m, b)
}
func (m *SearchEngineResultsPageTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchEngineResultsPageTypeEnum.Marshal(b, m, deterministic)
}
func (m *SearchEngineResultsPageTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchEngineResultsPageTypeEnum.Merge(m, src)
}
func (m *SearchEngineResultsPageTypeEnum) XXX_Size() int {
	return xxx_messageInfo_SearchEngineResultsPageTypeEnum.Size(m)
}
func (m *SearchEngineResultsPageTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchEngineResultsPageTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SearchEngineResultsPageTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType", SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_name, SearchEngineResultsPageTypeEnum_SearchEngineResultsPageType_value)
	proto.RegisterType((*SearchEngineResultsPageTypeEnum)(nil), "google.ads.googleads.v3.enums.SearchEngineResultsPageTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/search_engine_results_page_type.proto", fileDescriptor_8d0c9169c79e638f)
}

var fileDescriptor_8d0c9169c79e638f = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xd1, 0x4e, 0xf2, 0x30,
	0x14, 0xfe, 0x37, 0xfe, 0xa8, 0x29, 0x24, 0x2c, 0xf3, 0x4e, 0x25, 0x0a, 0x0f, 0xd0, 0x5d, 0xec,
	0xae, 0x5e, 0x15, 0x98, 0x84, 0x68, 0xca, 0x02, 0x82, 0xd1, 0x2c, 0x59, 0x2a, 0x6b, 0xea, 0x12,
	0x68, 0x9b, 0x75, 0x90, 0xf0, 0x1c, 0xbe, 0x81, 0x97, 0x3e, 0x8a, 0x8f, 0xe2, 0xbd, 0xf7, 0x66,
	0x2d, 0x70, 0xe7, 0x6e, 0x9a, 0x2f, 0xe7, 0x3b, 0xe7, 0xfb, 0xce, 0x77, 0x0a, 0x06, 0x5c, 0x4a,
	0xbe, 0x62, 0x01, 0xcd, 0x74, 0x60, 0x61, 0x85, 0xb6, 0x61, 0xc0, 0xc4, 0x66, 0xad, 0x03, 0xcd,
	0x68, 0xb1, 0x7c, 0x4b, 0x99, 0xe0, 0xb9, 0x60, 0x69, 0xc1, 0xf4, 0x66, 0x55, 0xea, 0x54, 0x51,
	0xce, 0xd2, 0x72, 0xa7, 0x18, 0x54, 0x85, 0x2c, 0xa5, 0xdf, 0xb1, 0x93, 0x90, 0x66, 0x1a, 0x1e,
	0x45, 0xe0, 0x36, 0x84, 0x46, 0xe4, 0xe2, 0xea, 0xe0, 0xa1, 0xf2, 0x80, 0x0a, 0x21, 0x4b, 0x5a,
	0xe6, 0x52, 0x68, 0x3b, 0xdc, 0x7b, 0x77, 0xc0, 0xf5, 0xcc, 0xd8, 0x44, 0xc6, 0x65, 0x6a, 0x4d,
	0x62, 0xca, 0xd9, 0xe3, 0x4e, 0xb1, 0x48, 0x6c, 0xd6, 0x3d, 0x05, 0x2e, 0x6b, 0x5a, 0xfc, 0x36,
	0x68, 0xce, 0xc9, 0x2c, 0x8e, 0x06, 0xe3, 0xbb, 0x71, 0x34, 0xf4, 0xfe, 0xf9, 0x4d, 0x70, 0x3a,
	0x27, 0xf7, 0x64, 0xf2, 0x44, 0x3c, 0xc7, 0x6f, 0x81, 0x33, 0x3c, 0x9c, 0xa5, 0x13, 0xf2, 0xf0,
	0xec, 0xb9, 0xbe, 0x07, 0x5a, 0x93, 0xe9, 0x08, 0x93, 0xf1, 0xc0, 0x56, 0x1a, 0xfe, 0x39, 0x68,
	0x57, 0x3c, 0x26, 0xc3, 0x74, 0xcf, 0x78, 0xff, 0xfb, 0x3f, 0x0e, 0xe8, 0x2e, 0xe5, 0x1a, 0xd6,
	0x26, 0xeb, 0xdf, 0xd4, 0x6c, 0x15, 0x57, 0xe9, 0x62, 0xe7, 0xa5, 0xbf, 0x97, 0xe0, 0x72, 0x45,
	0x05, 0x87, 0xb2, 0xe0, 0x01, 0x67, 0xc2, 0x64, 0x3f, 0x5c, 0x5c, 0xe5, 0xfa, 0x8f, 0x0f, 0xb8,
	0x35, 0xef, 0x87, 0xdb, 0x18, 0x61, 0xfc, 0xe9, 0x76, 0x46, 0x56, 0x0a, 0x67, 0x1a, 0x5a, 0x58,
	0xa1, 0x45, 0x08, 0xab, 0x23, 0xe9, 0xaf, 0x03, 0x9f, 0xe0, 0x4c, 0x27, 0x47, 0x3e, 0x59, 0x84,
	0x89, 0xe1, 0xbf, 0xdd, 0xae, 0x2d, 0x22, 0x84, 0x33, 0x8d, 0xd0, 0xb1, 0x03, 0xa1, 0x45, 0x88,
	0x90, 0xe9, 0x79, 0x3d, 0x31, 0x8b, 0x85, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xf0, 0x2c,
	0x2c, 0x18, 0x02, 0x00, 0x00,
}
