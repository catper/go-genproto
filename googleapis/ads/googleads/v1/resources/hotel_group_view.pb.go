// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/hotel_group_view.proto

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

// A hotel group view.
type HotelGroupView struct {
	// Immutable. The resource name of the hotel group view.
	// Hotel Group view resource names have the form:
	//
	// `customers/{customer_id}/hotelGroupViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelGroupView) Reset()         { *m = HotelGroupView{} }
func (m *HotelGroupView) String() string { return proto.CompactTextString(m) }
func (*HotelGroupView) ProtoMessage()    {}
func (*HotelGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa913b710ae8e895, []int{0}
}

func (m *HotelGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelGroupView.Unmarshal(m, b)
}
func (m *HotelGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelGroupView.Marshal(b, m, deterministic)
}
func (m *HotelGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelGroupView.Merge(m, src)
}
func (m *HotelGroupView) XXX_Size() int {
	return xxx_messageInfo_HotelGroupView.Size(m)
}
func (m *HotelGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_HotelGroupView proto.InternalMessageInfo

func (m *HotelGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*HotelGroupView)(nil), "google.ads.googleads.v1.resources.HotelGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/hotel_group_view.proto", fileDescriptor_fa913b710ae8e895)
}

var fileDescriptor_fa913b710ae8e895 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0x49, 0x2e, 0xf7, 0xc2, 0x0d, 0xf7, 0xba, 0xa8, 0x1b, 0x2d, 0x82, 0x56, 0x28, 0xea,
	0x66, 0x86, 0xe0, 0x42, 0x19, 0x57, 0xd3, 0x4d, 0xc5, 0x85, 0x94, 0x22, 0x59, 0x48, 0x20, 0x4c,
	0x93, 0x31, 0x1d, 0x48, 0xf2, 0x87, 0x99, 0x34, 0x5d, 0x94, 0x82, 0xcf, 0xe2, 0xd2, 0x47, 0x11,
	0x7c, 0x07, 0xd7, 0x7d, 0x04, 0x57, 0x92, 0x4e, 0x67, 0xda, 0xba, 0x50, 0x77, 0x07, 0xfe, 0xef,
	0x9c, 0x39, 0x9c, 0xf1, 0x2e, 0x53, 0x80, 0x34, 0xe3, 0x98, 0x25, 0x0a, 0x6b, 0xd9, 0xa8, 0xda,
	0xc7, 0x92, 0x2b, 0x98, 0xc8, 0x98, 0x2b, 0x3c, 0x86, 0x8a, 0x67, 0x51, 0x2a, 0x61, 0x52, 0x46,
	0xb5, 0xe0, 0x53, 0x54, 0x4a, 0xa8, 0xa0, 0xd5, 0xd1, 0x38, 0x62, 0x89, 0x42, 0xd6, 0x89, 0x6a,
	0x1f, 0x59, 0x67, 0xfb, 0xd0, 0x84, 0x97, 0x02, 0x3f, 0x08, 0x9e, 0x25, 0xd1, 0x88, 0x8f, 0x59,
	0x2d, 0x40, 0xea, 0x8c, 0xf6, 0xfe, 0x06, 0x60, 0x6c, 0xab, 0xd3, 0xc1, 0xc6, 0x89, 0x15, 0x05,
	0x54, 0xac, 0x12, 0x50, 0x28, 0x7d, 0x3d, 0x7e, 0x75, 0xbc, 0x9d, 0xeb, 0xa6, 0x57, 0xbf, 0xa9,
	0x15, 0x08, 0x3e, 0x6d, 0xdd, 0x79, 0xff, 0x4d, 0x44, 0x54, 0xb0, 0x9c, 0xef, 0x39, 0x47, 0xce,
	0xe9, 0xdf, 0x1e, 0x7e, 0xa3, 0xbf, 0xdf, 0xe9, 0x99, 0x77, 0xb2, 0xee, 0xb8, 0x52, 0xa5, 0x50,
	0x28, 0x86, 0x1c, 0x6f, 0xe7, 0x0c, 0xff, 0x99, 0x94, 0x5b, 0x96, 0x73, 0xc2, 0x17, 0x74, 0xf4,
	0x63, 0x6f, 0xeb, 0x22, 0x9e, 0xa8, 0x0a, 0x72, 0x2e, 0x15, 0x9e, 0x19, 0x39, 0xd7, 0x03, 0x5a,
	0x48, 0xe1, 0xd9, 0xe7, 0x45, 0xe7, 0xbd, 0x47, 0xd7, 0xeb, 0xc6, 0x90, 0xa3, 0x6f, 0x37, 0xed,
	0xed, 0x6e, 0x3f, 0x39, 0x68, 0xe6, 0x18, 0x38, 0xf7, 0x37, 0x2b, 0x67, 0x0a, 0x19, 0x2b, 0x52,
	0x04, 0x32, 0xc5, 0x29, 0x2f, 0x96, 0x63, 0xe1, 0x75, 0xe7, 0x2f, 0xbe, 0xf9, 0xca, 0xaa, 0x27,
	0xf7, 0x57, 0x9f, 0xd2, 0x67, 0xb7, 0xd3, 0xd7, 0x91, 0x34, 0x51, 0x48, 0xcb, 0x46, 0x05, 0x3e,
	0x1a, 0x1a, 0xf2, 0xc5, 0x30, 0x21, 0x4d, 0x54, 0x68, 0x99, 0x30, 0xf0, 0x43, 0xcb, 0x2c, 0xdc,
	0xae, 0x3e, 0x10, 0x42, 0x13, 0x45, 0x88, 0xa5, 0x08, 0x09, 0x7c, 0x42, 0x2c, 0x37, 0xfa, 0xb3,
	0x2c, 0x7b, 0xfe, 0x11, 0x00, 0x00, 0xff, 0xff, 0x59, 0x09, 0xd8, 0x46, 0x92, 0x02, 0x00, 0x00,
}
