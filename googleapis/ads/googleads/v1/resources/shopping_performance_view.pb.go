// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/shopping_performance_view.proto

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

// Shopping performance view.
// Provides Shopping campaign statistics aggregated at several product dimension
// levels. Product dimension values from Merchant Center such as brand,
// category, custom attributes, product condition and product type will reflect
// the state of each dimension as of the date and time when the corresponding
// event was recorded.
type ShoppingPerformanceView struct {
	// Immutable. The resource name of the Shopping performance view.
	// Shopping performance view resource names have the form:
	// `customers/{customer_id}/shoppingPerformanceView`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShoppingPerformanceView) Reset()         { *m = ShoppingPerformanceView{} }
func (m *ShoppingPerformanceView) String() string { return proto.CompactTextString(m) }
func (*ShoppingPerformanceView) ProtoMessage()    {}
func (*ShoppingPerformanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_37285869b217afa9, []int{0}
}

func (m *ShoppingPerformanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShoppingPerformanceView.Unmarshal(m, b)
}
func (m *ShoppingPerformanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShoppingPerformanceView.Marshal(b, m, deterministic)
}
func (m *ShoppingPerformanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShoppingPerformanceView.Merge(m, src)
}
func (m *ShoppingPerformanceView) XXX_Size() int {
	return xxx_messageInfo_ShoppingPerformanceView.Size(m)
}
func (m *ShoppingPerformanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_ShoppingPerformanceView.DiscardUnknown(m)
}

var xxx_messageInfo_ShoppingPerformanceView proto.InternalMessageInfo

func (m *ShoppingPerformanceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ShoppingPerformanceView)(nil), "google.ads.googleads.v1.resources.ShoppingPerformanceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/shopping_performance_view.proto", fileDescriptor_37285869b217afa9)
}

var fileDescriptor_37285869b217afa9 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0x2e, 0x57, 0x30, 0xe8, 0xa6, 0x1b, 0xb5, 0x14, 0xb4, 0x42, 0xc1, 0x85, 0xcc,
	0x18, 0xdd, 0xc8, 0xb8, 0x9a, 0x6e, 0x0a, 0x2e, 0xa4, 0x54, 0xc8, 0x42, 0x22, 0x71, 0x9a, 0x4c,
	0xd3, 0x81, 0x66, 0x4e, 0x98, 0x49, 0xd3, 0x85, 0xf8, 0x00, 0xbe, 0x86, 0x4b, 0x1f, 0xc5, 0x77,
	0x10, 0x5c, 0xf7, 0x11, 0x5c, 0x49, 0x9b, 0xcc, 0xb4, 0x08, 0x55, 0xdc, 0xfd, 0x70, 0xbe, 0xf9,
	0xce, 0xcf, 0x61, 0x3c, 0x9a, 0x02, 0xa4, 0x13, 0x8e, 0x59, 0xa2, 0x71, 0x15, 0x17, 0xa9, 0xf4,
	0xb1, 0xe2, 0x1a, 0xa6, 0x2a, 0xe6, 0x1a, 0xeb, 0x31, 0xe4, 0xb9, 0x90, 0x69, 0x94, 0x73, 0x35,
	0x02, 0x95, 0x31, 0x19, 0xf3, 0xa8, 0x14, 0x7c, 0x86, 0x72, 0x05, 0x05, 0x34, 0xda, 0xd5, 0x3b,
	0xc4, 0x12, 0x8d, 0xac, 0x02, 0x95, 0x3e, 0xb2, 0x8a, 0xe6, 0xa1, 0xd9, 0x92, 0x0b, 0x3c, 0x12,
	0x7c, 0x92, 0x44, 0x43, 0x3e, 0x66, 0xa5, 0x00, 0x55, 0x39, 0x9a, 0x07, 0x6b, 0x80, 0x79, 0x56,
	0x8f, 0x5a, 0x6b, 0x23, 0x26, 0x25, 0x14, 0xac, 0x10, 0x20, 0x75, 0x35, 0x3d, 0x7e, 0x77, 0xbc,
	0xbd, 0xdb, 0xba, 0x60, 0x7f, 0xd5, 0x2f, 0x10, 0x7c, 0xd6, 0xb8, 0xf7, 0x76, 0x8d, 0x2b, 0x92,
	0x2c, 0xe3, 0xfb, 0xce, 0x91, 0x73, 0xb2, 0xdd, 0xbd, 0xfc, 0xa0, 0xff, 0x3f, 0xe9, 0xb9, 0x77,
	0xb6, 0x2a, 0x5b, 0xa7, 0x5c, 0x68, 0x14, 0x43, 0x86, 0x37, 0x08, 0x07, 0x3b, 0x46, 0x77, 0xc3,
	0x32, 0x4e, 0xe2, 0x39, 0x7d, 0xf8, 0xbb, 0xa4, 0x71, 0x1a, 0x4f, 0x75, 0x01, 0x19, 0x57, 0x1a,
	0x3f, 0x9a, 0xf8, 0x64, 0x8f, 0xfc, 0x8d, 0xee, 0x3e, 0xbb, 0x5e, 0x27, 0x86, 0x0c, 0xfd, 0x7a,
	0xe3, 0x6e, 0x6b, 0xc3, 0xc2, 0xfe, 0xe2, 0x4e, 0x7d, 0xe7, 0xee, 0xba, 0x56, 0xa4, 0x30, 0x61,
	0x32, 0x45, 0xa0, 0x52, 0x9c, 0x72, 0xb9, 0xbc, 0x22, 0x5e, 0x55, 0xff, 0xe1, 0x23, 0x5c, 0xd9,
	0xf4, 0xe2, 0xfe, 0xeb, 0x51, 0xfa, 0xea, 0xb6, 0x7b, 0x95, 0x92, 0x26, 0x1a, 0x55, 0x71, 0x91,
	0x02, 0x1f, 0x0d, 0x0c, 0xf9, 0x66, 0x98, 0x90, 0x26, 0x3a, 0xb4, 0x4c, 0x18, 0xf8, 0xa1, 0x65,
	0xe6, 0x6e, 0xa7, 0x1a, 0x10, 0x42, 0x13, 0x4d, 0x88, 0xa5, 0x08, 0x09, 0x7c, 0x42, 0x2c, 0x37,
	0xdc, 0x5a, 0x96, 0xbd, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x26, 0xd6, 0x5e, 0x69, 0xb4, 0x02,
	0x00, 0x00,
}
