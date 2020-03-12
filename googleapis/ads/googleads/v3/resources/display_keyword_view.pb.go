// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/display_keyword_view.proto

package resources

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

// A display keyword view.
type DisplayKeywordView struct {
	// The resource name of the display keyword view.
	// Display Keyword view resource names have the form:
	//
	// `customers/{customer_id}/displayKeywordViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisplayKeywordView) Reset()         { *m = DisplayKeywordView{} }
func (m *DisplayKeywordView) String() string { return proto.CompactTextString(m) }
func (*DisplayKeywordView) ProtoMessage()    {}
func (*DisplayKeywordView) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b0a6220f800cacf, []int{0}
}

func (m *DisplayKeywordView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisplayKeywordView.Unmarshal(m, b)
}
func (m *DisplayKeywordView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisplayKeywordView.Marshal(b, m, deterministic)
}
func (m *DisplayKeywordView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisplayKeywordView.Merge(m, src)
}
func (m *DisplayKeywordView) XXX_Size() int {
	return xxx_messageInfo_DisplayKeywordView.Size(m)
}
func (m *DisplayKeywordView) XXX_DiscardUnknown() {
	xxx_messageInfo_DisplayKeywordView.DiscardUnknown(m)
}

var xxx_messageInfo_DisplayKeywordView proto.InternalMessageInfo

func (m *DisplayKeywordView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*DisplayKeywordView)(nil), "google.ads.googleads.v3.resources.DisplayKeywordView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/display_keyword_view.proto", fileDescriptor_5b0a6220f800cacf)
}

var fileDescriptor_5b0a6220f800cacf = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x04, 0xc1, 0xa0, 0x97, 0x5e, 0xd4, 0xe2, 0xc1, 0x2a, 0x05, 0x41, 0xd8, 0x3d,
	0xe4, 0xb6, 0x0a, 0x92, 0x22, 0x14, 0x14, 0xa4, 0xf4, 0x90, 0x83, 0x04, 0xca, 0x9a, 0x5d, 0xc2,
	0x62, 0xb3, 0x13, 0x77, 0xd2, 0x96, 0x52, 0x7a, 0xf3, 0x49, 0x3c, 0x8a, 0x4f, 0xe2, 0xa3, 0xf4,
	0x29, 0xa4, 0xdd, 0xee, 0xb6, 0x60, 0xd1, 0xdb, 0x4f, 0xe6, 0xfb, 0xff, 0x7f, 0x67, 0x12, 0xdd,
	0x16, 0x00, 0xc5, 0x50, 0x52, 0x2e, 0x90, 0x5a, 0xb9, 0x54, 0xe3, 0x98, 0x1a, 0x89, 0x30, 0x32,
	0xb9, 0x44, 0x2a, 0x14, 0x56, 0x43, 0x3e, 0x1d, 0xbc, 0xca, 0xe9, 0x04, 0x8c, 0x18, 0x8c, 0x95,
	0x9c, 0x90, 0xca, 0x40, 0x0d, 0x8d, 0x96, 0xb5, 0x10, 0x2e, 0x90, 0x78, 0x37, 0x19, 0xc7, 0xc4,
	0xbb, 0x9b, 0xa7, 0xae, 0xa0, 0x52, 0x3e, 0xd3, 0xba, 0x9b, 0x67, 0x5b, 0x23, 0xae, 0x35, 0xd4,
	0xbc, 0x56, 0xa0, 0xd1, 0x4e, 0x2f, 0xbe, 0x82, 0xa8, 0x71, 0x6f, 0xab, 0x1f, 0x6d, 0x73, 0xaa,
	0xe4, 0xa4, 0x71, 0x19, 0x1d, 0xb9, 0x98, 0x81, 0xe6, 0xa5, 0x3c, 0x09, 0xce, 0x83, 0xab, 0x83,
	0xfe, 0xa1, 0xfb, 0xf8, 0xc4, 0x4b, 0xc9, 0xde, 0x16, 0x89, 0x8e, 0xae, 0x37, 0x2f, 0x5a, 0xab,
	0x4a, 0x21, 0xc9, 0xa1, 0xa4, 0x3b, 0x62, 0xef, 0xf2, 0x11, 0xd6, 0x50, 0x4a, 0x83, 0x74, 0xe6,
	0xe4, 0xdc, 0xad, 0xbe, 0x05, 0x22, 0x9d, 0xed, 0xba, 0xc7, 0xbc, 0xf3, 0x1e, 0x46, 0xed, 0x1c,
	0x4a, 0xf2, 0xef, 0x45, 0x3a, 0xc7, 0xbf, 0xeb, 0x7b, 0xcb, 0x8d, 0x7b, 0xc1, 0xf3, 0xc3, 0xda,
	0x5d, 0xc0, 0x90, 0xeb, 0x82, 0x80, 0x29, 0x68, 0x21, 0xf5, 0xea, 0x1e, 0x74, 0xb3, 0xc3, 0x1f,
	0x3f, 0xeb, 0xc6, 0xab, 0x8f, 0x70, 0xaf, 0x9b, 0x24, 0x9f, 0x61, 0xab, 0x6b, 0x23, 0x13, 0x81,
	0xc4, 0xca, 0xa5, 0x4a, 0x63, 0xd2, 0x77, 0xe4, 0xb7, 0x63, 0xb2, 0x44, 0x60, 0xe6, 0x99, 0x2c,
	0x8d, 0x33, 0xcf, 0x2c, 0xc2, 0xb6, 0x1d, 0x30, 0x96, 0x08, 0x64, 0xcc, 0x53, 0x8c, 0xa5, 0x31,
	0x63, 0x9e, 0x7b, 0xd9, 0x5f, 0x3d, 0x36, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x51, 0xa4,
	0x8b, 0x58, 0x02, 0x00, 0x00,
}
