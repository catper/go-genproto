// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/product_group_view.proto

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

// A product group view.
type ProductGroupView struct {
	// Immutable. The resource name of the product group view.
	// Product group view resource names have the form:
	//
	// `customers/{customer_id}/productGroupViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductGroupView) Reset()         { *m = ProductGroupView{} }
func (m *ProductGroupView) String() string { return proto.CompactTextString(m) }
func (*ProductGroupView) ProtoMessage()    {}
func (*ProductGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4fb0e8e8872a613, []int{0}
}

func (m *ProductGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductGroupView.Unmarshal(m, b)
}
func (m *ProductGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductGroupView.Marshal(b, m, deterministic)
}
func (m *ProductGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductGroupView.Merge(m, src)
}
func (m *ProductGroupView) XXX_Size() int {
	return xxx_messageInfo_ProductGroupView.Size(m)
}
func (m *ProductGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_ProductGroupView proto.InternalMessageInfo

func (m *ProductGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductGroupView)(nil), "google.ads.googleads.v2.resources.ProductGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/product_group_view.proto", fileDescriptor_c4fb0e8e8872a613)
}

var fileDescriptor_c4fb0e8e8872a613 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0x49, 0x2e, 0xf7, 0xc2, 0x0d, 0x0a, 0x52, 0x10, 0xb4, 0x08, 0x5a, 0xa1, 0xa0, 0x08,
	0x33, 0x18, 0x77, 0xd3, 0xd5, 0x74, 0x53, 0x70, 0x21, 0xa5, 0x8b, 0x2c, 0x24, 0x10, 0xa6, 0xc9,
	0x98, 0x0e, 0x36, 0xf9, 0xc3, 0x4c, 0x92, 0x2e, 0x4a, 0x37, 0x3e, 0x8a, 0x4b, 0x1f, 0xc5, 0x17,
	0x70, 0xeb, 0xba, 0x8f, 0xe0, 0x4a, 0xd2, 0xe9, 0x4c, 0x4b, 0x05, 0xc5, 0xdd, 0x81, 0xff, 0x3b,
	0x67, 0x0e, 0x67, 0x3c, 0x92, 0x02, 0xa4, 0x53, 0x8e, 0x59, 0xa2, 0xb0, 0x96, 0x8d, 0xaa, 0x7d,
	0x2c, 0xb9, 0x82, 0x4a, 0xc6, 0x5c, 0xe1, 0x42, 0x42, 0x52, 0xc5, 0x65, 0x94, 0x4a, 0xa8, 0x8a,
	0xa8, 0x16, 0x7c, 0x86, 0x0a, 0x09, 0x25, 0xb4, 0x3a, 0xda, 0x80, 0x58, 0xa2, 0x90, 0xf5, 0xa2,
	0xda, 0x47, 0xd6, 0xdb, 0x3e, 0x35, 0xf1, 0x85, 0xc0, 0x0f, 0x82, 0x4f, 0x93, 0x68, 0xcc, 0x27,
	0xac, 0x16, 0x20, 0x75, 0x46, 0xfb, 0x78, 0x0b, 0x30, 0xb6, 0xf5, 0xe9, 0x64, 0xeb, 0xc4, 0xf2,
	0x1c, 0x4a, 0x56, 0x0a, 0xc8, 0x95, 0xbe, 0x9e, 0xbf, 0x39, 0xde, 0xc1, 0x50, 0x37, 0x1b, 0x34,
	0xc5, 0x02, 0xc1, 0x67, 0xad, 0xc0, 0xdb, 0x37, 0x21, 0x51, 0xce, 0x32, 0x7e, 0xe4, 0x9c, 0x39,
	0x17, 0xff, 0xfb, 0xd7, 0xef, 0xf4, 0xef, 0x07, 0xbd, 0xf2, 0x2e, 0x37, 0x2d, 0xd7, 0xaa, 0x10,
	0x0a, 0xc5, 0x90, 0xe1, 0xdd, 0xa4, 0xd1, 0x9e, 0xc9, 0xb9, 0x63, 0x19, 0x27, 0x8f, 0x4b, 0x3a,
	0xf9, 0x85, 0xbb, 0xd5, 0x8b, 0x2b, 0x55, 0x42, 0xc6, 0xa5, 0xc2, 0x73, 0x23, 0x17, 0x66, 0x48,
	0x8b, 0x29, 0x3c, 0xff, 0xba, 0xed, 0xa2, 0xff, 0xe4, 0x7a, 0xdd, 0x18, 0x32, 0xf4, 0xe3, 0xba,
	0xfd, 0xc3, 0xdd, 0x87, 0x87, 0xcd, 0x34, 0x43, 0xe7, 0xfe, 0x76, 0xed, 0x4d, 0x61, 0xca, 0xf2,
	0x14, 0x81, 0x4c, 0x71, 0xca, 0xf3, 0xd5, 0x70, 0x78, 0xd3, 0xfd, 0x9b, 0x4f, 0xef, 0x59, 0xf5,
	0xec, 0xfe, 0x19, 0x50, 0xfa, 0xe2, 0x76, 0x06, 0x3a, 0x92, 0x26, 0x0a, 0x69, 0xd9, 0xa8, 0xc0,
	0x47, 0x23, 0x43, 0xbe, 0x1a, 0x26, 0xa4, 0x89, 0x0a, 0x2d, 0x13, 0x06, 0x7e, 0x68, 0x99, 0xa5,
	0xdb, 0xd5, 0x07, 0x42, 0x68, 0xa2, 0x08, 0xb1, 0x14, 0x21, 0x81, 0x4f, 0x88, 0xe5, 0xc6, 0xff,
	0x56, 0x65, 0x6f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x55, 0xbd, 0x32, 0xa0, 0x02, 0x00,
	0x00,
}
