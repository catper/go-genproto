// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_schedule_view.proto

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

// An ad schedule view summarizes the performance of campaigns by
// AdSchedule criteria.
type AdScheduleView struct {
	// Immutable. The resource name of the ad schedule view.
	// AdSchedule view resource names have the form:
	//
	// `customers/{customer_id}/adScheduleViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdScheduleView) Reset()         { *m = AdScheduleView{} }
func (m *AdScheduleView) String() string { return proto.CompactTextString(m) }
func (*AdScheduleView) ProtoMessage()    {}
func (*AdScheduleView) Descriptor() ([]byte, []int) {
	return fileDescriptor_88641aea73ff6f50, []int{0}
}

func (m *AdScheduleView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdScheduleView.Unmarshal(m, b)
}
func (m *AdScheduleView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdScheduleView.Marshal(b, m, deterministic)
}
func (m *AdScheduleView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdScheduleView.Merge(m, src)
}
func (m *AdScheduleView) XXX_Size() int {
	return xxx_messageInfo_AdScheduleView.Size(m)
}
func (m *AdScheduleView) XXX_DiscardUnknown() {
	xxx_messageInfo_AdScheduleView.DiscardUnknown(m)
}

var xxx_messageInfo_AdScheduleView proto.InternalMessageInfo

func (m *AdScheduleView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AdScheduleView)(nil), "google.ads.googleads.v1.resources.AdScheduleView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_schedule_view.proto", fileDescriptor_88641aea73ff6f50)
}

var fileDescriptor_88641aea73ff6f50 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x1c, 0xc6, 0x49, 0x5e, 0x5e, 0xc1, 0xa0, 0x0e, 0x75, 0xd1, 0x22, 0x68, 0x85, 0xa2, 0x2e, 0x77,
	0x04, 0x07, 0xe5, 0x9c, 0xae, 0x4b, 0xc1, 0x41, 0x4a, 0x95, 0x0c, 0x12, 0x08, 0xd7, 0xdc, 0xdf,
	0xf4, 0x20, 0xc9, 0x95, 0x5c, 0x9a, 0x0e, 0xa5, 0xe0, 0x67, 0x71, 0xf4, 0xa3, 0x08, 0x7e, 0x07,
	0xe7, 0x7e, 0x04, 0x27, 0x69, 0x2f, 0x77, 0x6d, 0x1d, 0xd4, 0xed, 0x81, 0xff, 0xef, 0x79, 0xee,
	0xe1, 0x39, 0xef, 0x3a, 0x91, 0x32, 0x49, 0x01, 0x33, 0xae, 0xb0, 0x96, 0x0b, 0x55, 0xf9, 0xb8,
	0x00, 0x25, 0xc7, 0x45, 0x0c, 0x0a, 0x33, 0x1e, 0xa9, 0x78, 0x08, 0x7c, 0x9c, 0x42, 0x54, 0x09,
	0x98, 0xa0, 0x51, 0x21, 0x4b, 0xd9, 0x68, 0x69, 0x1c, 0x31, 0xae, 0x90, 0x75, 0xa2, 0xca, 0x47,
	0xd6, 0xd9, 0x3c, 0x36, 0xe1, 0x23, 0x81, 0x9f, 0x04, 0xa4, 0x3c, 0x1a, 0xc0, 0x90, 0x55, 0x42,
	0x16, 0x3a, 0xa3, 0x79, 0xb8, 0x06, 0x18, 0x5b, 0x7d, 0x3a, 0x5a, 0x3b, 0xb1, 0x3c, 0x97, 0x25,
	0x2b, 0x85, 0xcc, 0x95, 0xbe, 0x9e, 0xbe, 0x3b, 0xde, 0x1e, 0xe5, 0xf7, 0x75, 0xad, 0x40, 0xc0,
	0xa4, 0xf1, 0xe0, 0xed, 0x9a, 0x88, 0x28, 0x67, 0x19, 0x1c, 0x38, 0x27, 0xce, 0xf9, 0x76, 0x07,
	0x7f, 0xd0, 0xff, 0x9f, 0xf4, 0xc2, 0x3b, 0x5b, 0x75, 0xac, 0xd5, 0x48, 0x28, 0x14, 0xcb, 0x0c,
	0x6f, 0xe6, 0xf4, 0x77, 0x4c, 0xca, 0x1d, 0xcb, 0x80, 0xc0, 0x9c, 0x0e, 0xfe, 0xec, 0x6d, 0x5c,
	0xc5, 0x63, 0x55, 0xca, 0x0c, 0x0a, 0x85, 0xa7, 0x46, 0xce, 0x30, 0xdb, 0x80, 0x14, 0x9e, 0x7e,
	0x5f, 0x74, 0xd6, 0x79, 0x76, 0xbd, 0x76, 0x2c, 0x33, 0xf4, 0xeb, 0xa6, 0x9d, 0xfd, 0xcd, 0x27,
	0x7b, 0x8b, 0x39, 0x7a, 0xce, 0xe3, 0x6d, 0xed, 0x4c, 0x64, 0xca, 0xf2, 0x04, 0xc9, 0x22, 0xc1,
	0x09, 0xe4, 0xcb, 0xb1, 0xf0, 0xaa, 0xf3, 0x0f, 0xdf, 0x7c, 0x63, 0xd5, 0x8b, 0xfb, 0xaf, 0x4b,
	0xe9, 0xab, 0xdb, 0xea, 0xea, 0x48, 0xca, 0x15, 0xd2, 0x72, 0xa1, 0x02, 0x1f, 0xf5, 0x0d, 0xf9,
	0x66, 0x98, 0x90, 0x72, 0x15, 0x5a, 0x26, 0x0c, 0xfc, 0xd0, 0x32, 0x73, 0xb7, 0xad, 0x0f, 0x84,
	0x50, 0xae, 0x08, 0xb1, 0x14, 0x21, 0x81, 0x4f, 0x88, 0xe5, 0x06, 0x5b, 0xcb, 0xb2, 0x97, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xac, 0xa0, 0xeb, 0xee, 0x92, 0x02, 0x00, 0x00,
}
