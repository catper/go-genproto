// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_audience_view.proto

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

// An ad group audience view.
// Includes performance data from interests and remarketing lists for Display
// Network and YouTube Network ads, and remarketing lists for search ads (RLSA),
// aggregated at the audience level.
type AdGroupAudienceView struct {
	// Immutable. The resource name of the ad group audience view.
	// Ad group audience view resource names have the form:
	//
	// `customers/{customer_id}/adGroupAudienceViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAudienceView) Reset()         { *m = AdGroupAudienceView{} }
func (m *AdGroupAudienceView) String() string { return proto.CompactTextString(m) }
func (*AdGroupAudienceView) ProtoMessage()    {}
func (*AdGroupAudienceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8817e198a682ff, []int{0}
}

func (m *AdGroupAudienceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAudienceView.Unmarshal(m, b)
}
func (m *AdGroupAudienceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAudienceView.Marshal(b, m, deterministic)
}
func (m *AdGroupAudienceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAudienceView.Merge(m, src)
}
func (m *AdGroupAudienceView) XXX_Size() int {
	return xxx_messageInfo_AdGroupAudienceView.Size(m)
}
func (m *AdGroupAudienceView) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAudienceView.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAudienceView proto.InternalMessageInfo

func (m *AdGroupAudienceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AdGroupAudienceView)(nil), "google.ads.googleads.v1.resources.AdGroupAudienceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_audience_view.proto", fileDescriptor_bc8817e198a682ff)
}

var fileDescriptor_bc8817e198a682ff = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xc3, 0x40,
	0x1c, 0xc5, 0x49, 0x44, 0xc1, 0xa0, 0x4b, 0x5d, 0x6a, 0x11, 0xb4, 0x42, 0xc1, 0x41, 0xee, 0x08,
	0x3a, 0x9d, 0x20, 0x5c, 0x96, 0x82, 0x83, 0x94, 0x0e, 0x01, 0x25, 0x10, 0xae, 0xb9, 0x33, 0x1e,
	0x34, 0xf9, 0x87, 0xbb, 0x24, 0x1d, 0x4a, 0x47, 0xbf, 0x88, 0xa3, 0x1f, 0xc5, 0x4f, 0xe1, 0x26,
	0xf4, 0x23, 0x38, 0x49, 0x9a, 0xe4, 0x5a, 0xb0, 0x28, 0x6e, 0x0f, 0xfe, 0xbf, 0xf7, 0xf2, 0xf2,
	0x38, 0xe7, 0x36, 0x06, 0x88, 0xa7, 0x02, 0x33, 0xae, 0x71, 0x2d, 0x2b, 0x55, 0xba, 0x58, 0x09,
	0x0d, 0x85, 0x8a, 0x84, 0xc6, 0x8c, 0x87, 0xb1, 0x82, 0x22, 0x0b, 0x59, 0xc1, 0xa5, 0x48, 0x23,
	0x11, 0x96, 0x52, 0xcc, 0x50, 0xa6, 0x20, 0x87, 0x4e, 0xbf, 0x36, 0x21, 0xc6, 0x35, 0x32, 0x7e,
	0x54, 0xba, 0xc8, 0xf8, 0x7b, 0xa7, 0xed, 0x27, 0x32, 0x89, 0x9f, 0xa4, 0x98, 0xf2, 0x70, 0x22,
	0x9e, 0x59, 0x29, 0x41, 0xd5, 0x19, 0xbd, 0xe3, 0x0d, 0xa0, 0xb5, 0x35, 0xa7, 0x93, 0x8d, 0x13,
	0x4b, 0x53, 0xc8, 0x59, 0x2e, 0x21, 0xd5, 0xf5, 0xf5, 0xfc, 0xd3, 0x72, 0x8e, 0x28, 0x1f, 0x56,
	0xe5, 0x68, 0xd3, 0xcd, 0x97, 0x62, 0xd6, 0x79, 0x70, 0x0e, 0xdb, 0x9c, 0x30, 0x65, 0x89, 0xe8,
	0x5a, 0x67, 0xd6, 0xc5, 0xbe, 0x77, 0xfd, 0x41, 0x77, 0xbf, 0x28, 0x72, 0x2e, 0xd7, 0x45, 0x1b,
	0x95, 0x49, 0x8d, 0x22, 0x48, 0xf0, 0x96, 0xb0, 0xf1, 0x41, 0x1b, 0x75, 0xcf, 0x12, 0x41, 0x8a,
	0x25, 0x55, 0xff, 0x0b, 0xe8, 0x78, 0x51, 0xa1, 0x73, 0x48, 0x84, 0xd2, 0x78, 0xde, 0xca, 0x05,
	0x66, 0x3f, 0x49, 0x8d, 0xe7, 0xdb, 0xb7, 0x5e, 0x78, 0x2f, 0xb6, 0x33, 0x88, 0x20, 0x41, 0x7f,
	0xae, 0xed, 0x75, 0xb7, 0x54, 0x18, 0x55, 0x6b, 0x8d, 0xac, 0xc7, 0xbb, 0xc6, 0x1e, 0xc3, 0x94,
	0xa5, 0x31, 0x02, 0x15, 0xe3, 0x58, 0xa4, 0xab, 0x2d, 0xf1, 0xfa, 0x47, 0x7e, 0x79, 0x0b, 0x37,
	0x46, 0xbd, 0xda, 0x3b, 0x43, 0x4a, 0xdf, 0xec, 0xfe, 0xb0, 0x8e, 0xa4, 0x5c, 0xa3, 0x5a, 0x56,
	0xca, 0x77, 0xd1, 0xb8, 0x25, 0xdf, 0x5b, 0x26, 0xa0, 0x5c, 0x07, 0x86, 0x09, 0x7c, 0x37, 0x30,
	0xcc, 0xd2, 0x1e, 0xd4, 0x07, 0x42, 0x28, 0xd7, 0x84, 0x18, 0x8a, 0x10, 0xdf, 0x25, 0xc4, 0x70,
	0x93, 0xbd, 0x55, 0xd9, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xb2, 0x93, 0xc5, 0xb7,
	0x02, 0x00, 0x00,
}
