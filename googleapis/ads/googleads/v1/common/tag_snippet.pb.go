// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/tag_snippet.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// The site tag and event snippet pair for a TrackingCodeType.
type TagSnippet struct {
	// The type of the generated tag snippets for tracking conversions.
	Type enums.TrackingCodeTypeEnum_TrackingCodeType `protobuf:"varint,1,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.TrackingCodeTypeEnum_TrackingCodeType" json:"type,omitempty"`
	// The format of the web page where the tracking tag and snippet will be
	// installed, e.g. HTML.
	PageFormat enums.TrackingCodePageFormatEnum_TrackingCodePageFormat `protobuf:"varint,2,opt,name=page_format,json=pageFormat,proto3,enum=google.ads.googleads.v1.enums.TrackingCodePageFormatEnum_TrackingCodePageFormat" json:"page_format,omitempty"`
	// The site tag that adds visitors to your basic remarketing lists and sets
	// new cookies on your domain.
	GlobalSiteTag *wrappers.StringValue `protobuf:"bytes,3,opt,name=global_site_tag,json=globalSiteTag,proto3" json:"global_site_tag,omitempty"`
	// The event snippet that works with the site tag to track actions that
	// should be counted as conversions.
	EventSnippet         *wrappers.StringValue `protobuf:"bytes,4,opt,name=event_snippet,json=eventSnippet,proto3" json:"event_snippet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TagSnippet) Reset()         { *m = TagSnippet{} }
func (m *TagSnippet) String() string { return proto.CompactTextString(m) }
func (*TagSnippet) ProtoMessage()    {}
func (*TagSnippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_24cd5966e9dfdab2, []int{0}
}

func (m *TagSnippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagSnippet.Unmarshal(m, b)
}
func (m *TagSnippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagSnippet.Marshal(b, m, deterministic)
}
func (m *TagSnippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagSnippet.Merge(m, src)
}
func (m *TagSnippet) XXX_Size() int {
	return xxx_messageInfo_TagSnippet.Size(m)
}
func (m *TagSnippet) XXX_DiscardUnknown() {
	xxx_messageInfo_TagSnippet.DiscardUnknown(m)
}

var xxx_messageInfo_TagSnippet proto.InternalMessageInfo

func (m *TagSnippet) GetType() enums.TrackingCodeTypeEnum_TrackingCodeType {
	if m != nil {
		return m.Type
	}
	return enums.TrackingCodeTypeEnum_UNSPECIFIED
}

func (m *TagSnippet) GetPageFormat() enums.TrackingCodePageFormatEnum_TrackingCodePageFormat {
	if m != nil {
		return m.PageFormat
	}
	return enums.TrackingCodePageFormatEnum_UNSPECIFIED
}

func (m *TagSnippet) GetGlobalSiteTag() *wrappers.StringValue {
	if m != nil {
		return m.GlobalSiteTag
	}
	return nil
}

func (m *TagSnippet) GetEventSnippet() *wrappers.StringValue {
	if m != nil {
		return m.EventSnippet
	}
	return nil
}

func init() {
	proto.RegisterType((*TagSnippet)(nil), "google.ads.googleads.v1.common.TagSnippet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/tag_snippet.proto", fileDescriptor_24cd5966e9dfdab2)
}

var fileDescriptor_24cd5966e9dfdab2 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0x6e, 0xd9, 0x41, 0x5d, 0x57, 0xf0, 0x29, 0x94, 0x52, 0x4a, 0x4e, 0x3d, 0x49, 0x73,
	0x07, 0x3b, 0x68, 0xec, 0xe0, 0x36, 0x5b, 0xaf, 0x21, 0x09, 0x61, 0x8c, 0x80, 0x79, 0x89, 0x5f,
	0x85, 0x98, 0x2d, 0x69, 0x96, 0x92, 0xd1, 0xaf, 0xb3, 0xe3, 0x3e, 0xca, 0x3e, 0xc6, 0x8e, 0xfd,
	0x14, 0xc3, 0x92, 0xec, 0x0e, 0x4a, 0x46, 0x7b, 0xf2, 0xcf, 0x7a, 0xbf, 0x3f, 0xef, 0x49, 0x8f,
	0xbc, 0x15, 0x5a, 0x8b, 0x1a, 0x19, 0x54, 0x96, 0x05, 0xd8, 0xa1, 0x5d, 0xce, 0x36, 0xba, 0x69,
	0xb4, 0x62, 0x0e, 0x44, 0x69, 0x95, 0x34, 0x06, 0x1d, 0x35, 0xad, 0x76, 0x3a, 0x3b, 0x0f, 0x34,
	0x0a, 0x95, 0xa5, 0x83, 0x82, 0xee, 0x72, 0x1a, 0x14, 0xa7, 0x1f, 0xf7, 0x39, 0xa2, 0xda, 0x36,
	0x96, 0xb9, 0x16, 0x36, 0xdf, 0xa4, 0x12, 0xe5, 0x46, 0x57, 0x58, 0x1a, 0x10, 0x58, 0xde, 0xe9,
	0xb6, 0x81, 0x68, 0x7f, 0xfa, 0xfe, 0x25, 0x72, 0x77, 0x6f, 0x30, 0xea, 0x62, 0x5b, 0xcc, 0xff,
	0xad, 0xb7, 0x77, 0xec, 0x47, 0x0b, 0xc6, 0x60, 0x6b, 0x63, 0xfd, 0xac, 0xf7, 0x35, 0x92, 0x81,
	0x52, 0xda, 0x81, 0x93, 0x5a, 0xc5, 0xea, 0xf8, 0x4f, 0x4a, 0xc8, 0x02, 0xc4, 0x3c, 0x4c, 0x9a,
	0x7d, 0x21, 0x87, 0x9d, 0xf5, 0x28, 0xb9, 0x48, 0x2e, 0xdf, 0x5c, 0x4d, 0xe8, 0xbe, 0x91, 0x7d,
	0x4f, 0x74, 0x11, 0x7b, 0xba, 0xd1, 0x15, 0x2e, 0xee, 0x0d, 0x7e, 0x52, 0xdb, 0xe6, 0xc9, 0xe1,
	0xcc, 0x3b, 0x66, 0xdf, 0xc9, 0xd1, 0x3f, 0x33, 0x8f, 0x52, 0x1f, 0x30, 0x7d, 0x41, 0xc0, 0x14,
	0x04, 0x7e, 0xf6, 0xe2, 0x27, 0x31, 0x8f, 0xa5, 0x19, 0x31, 0x03, 0xce, 0x26, 0xe4, 0x44, 0xd4,
	0x7a, 0x0d, 0x75, 0x69, 0xa5, 0xc3, 0xd2, 0x81, 0x18, 0x1d, 0x5c, 0x24, 0x97, 0x47, 0x57, 0x67,
	0x7d, 0x6c, 0x7f, 0x67, 0x74, 0xee, 0x5a, 0xa9, 0xc4, 0x12, 0xea, 0x2d, 0xce, 0x8e, 0x83, 0x68,
	0x2e, 0x1d, 0x2e, 0x40, 0x64, 0x05, 0x39, 0xc6, 0x1d, 0x2a, 0xd7, 0x6f, 0xc3, 0xe8, 0xf0, 0x19,
	0x1e, 0xaf, 0xbd, 0x24, 0xde, 0xea, 0xf5, 0x43, 0x42, 0xc6, 0x1b, 0xdd, 0xd0, 0xff, 0x2f, 0xd0,
	0xf5, 0xc9, 0xe3, 0x43, 0x4c, 0x3b, 0xd3, 0x69, 0xf2, 0x75, 0x12, 0x25, 0x42, 0xd7, 0xa0, 0x04,
	0xd5, 0xad, 0x60, 0x02, 0x95, 0x8f, 0xec, 0x97, 0xc4, 0x48, 0xbb, 0x6f, 0x89, 0x3f, 0x84, 0xcf,
	0xcf, 0xf4, 0xe0, 0xb6, 0x28, 0x7e, 0xa5, 0xe7, 0xb7, 0xc1, 0xac, 0xa8, 0x2c, 0x0d, 0xb0, 0x43,
	0xcb, 0x9c, 0xde, 0x78, 0xda, 0xef, 0x9e, 0xb0, 0x2a, 0x2a, 0xbb, 0x1a, 0x08, 0xab, 0x65, 0xbe,
	0x0a, 0x84, 0x87, 0x74, 0x1c, 0x4e, 0x39, 0x2f, 0x2a, 0xcb, 0xf9, 0x40, 0xe1, 0x7c, 0x99, 0x73,
	0x1e, 0x48, 0xeb, 0x57, 0xbe, 0xbb, 0x77, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x9e, 0x62,
	0x27, 0x61, 0x03, 0x00, 0x00,
}
