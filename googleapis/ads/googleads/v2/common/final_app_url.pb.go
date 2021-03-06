// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/common/final_app_url.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A URL for deep linking into an app for the given operating system.
type FinalAppUrl struct {
	// The operating system targeted by this URL. Required.
	OsType enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType `protobuf:"varint,1,opt,name=os_type,json=osType,proto3,enum=google.ads.googleads.v2.enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType" json:"os_type,omitempty"`
	// The app deep link URL. Deep links specify a location in an app that
	// corresponds to the content you'd like to show, and should be of the form
	// {scheme}://{host_path}
	// The scheme identifies which app to open. For your app, you can use a custom
	// scheme that starts with the app's name. The host and path specify the
	// unique location in the app where your content exists.
	// Example: "exampleapp://productid_1234". Required.
	Url                  *wrappers.StringValue `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FinalAppUrl) Reset()         { *m = FinalAppUrl{} }
func (m *FinalAppUrl) String() string { return proto.CompactTextString(m) }
func (*FinalAppUrl) ProtoMessage()    {}
func (*FinalAppUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f9a04033e8150d6, []int{0}
}

func (m *FinalAppUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinalAppUrl.Unmarshal(m, b)
}
func (m *FinalAppUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinalAppUrl.Marshal(b, m, deterministic)
}
func (m *FinalAppUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinalAppUrl.Merge(m, src)
}
func (m *FinalAppUrl) XXX_Size() int {
	return xxx_messageInfo_FinalAppUrl.Size(m)
}
func (m *FinalAppUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_FinalAppUrl.DiscardUnknown(m)
}

var xxx_messageInfo_FinalAppUrl proto.InternalMessageInfo

func (m *FinalAppUrl) GetOsType() enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType {
	if m != nil {
		return m.OsType
	}
	return enums.AppUrlOperatingSystemTypeEnum_UNSPECIFIED
}

func (m *FinalAppUrl) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func init() {
	proto.RegisterType((*FinalAppUrl)(nil), "google.ads.googleads.v2.common.FinalAppUrl")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/common/final_app_url.proto", fileDescriptor_4f9a04033e8150d6)
}

var fileDescriptor_4f9a04033e8150d6 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x6a, 0xeb, 0x30,
	0x14, 0x86, 0xb1, 0x03, 0xb9, 0xe0, 0xc0, 0xe5, 0x92, 0x29, 0x84, 0x10, 0x42, 0xa6, 0x4c, 0x12,
	0xe8, 0x6e, 0xba, 0x93, 0x73, 0xdb, 0x66, 0x6c, 0x48, 0x5a, 0x0f, 0xc5, 0x60, 0x94, 0x58, 0x11,
	0x06, 0x5b, 0x47, 0x48, 0x72, 0x4a, 0x5e, 0xa7, 0x63, 0x87, 0x3e, 0x48, 0x1f, 0xa5, 0x7d, 0x89,
	0x62, 0xc9, 0x0e, 0x5d, 0xdc, 0x49, 0xc7, 0x3e, 0xdf, 0xf9, 0xff, 0xff, 0x48, 0x11, 0x11, 0x00,
	0xa2, 0xe4, 0x98, 0xe5, 0x06, 0xfb, 0xb2, 0xa9, 0xce, 0x04, 0x1f, 0xa1, 0xaa, 0x40, 0xe2, 0x53,
	0x21, 0x59, 0x99, 0x31, 0xa5, 0xb2, 0x5a, 0x97, 0x48, 0x69, 0xb0, 0x30, 0x9e, 0x7b, 0x10, 0xb1,
	0xdc, 0xa0, 0xeb, 0x0c, 0x3a, 0x13, 0xe4, 0x67, 0xa6, 0x71, 0x9f, 0x26, 0x97, 0x75, 0x65, 0x70,
	0x2b, 0x96, 0x81, 0xe2, 0x9a, 0xd9, 0x42, 0x8a, 0xcc, 0x5c, 0x8c, 0xe5, 0x55, 0x66, 0x2f, 0x8a,
	0x7b, 0x8b, 0x69, 0x6b, 0x81, 0xdd, 0xd7, 0xa1, 0x3e, 0xe1, 0x67, 0xcd, 0x94, 0xe2, 0xda, 0xb4,
	0xfd, 0x59, 0x67, 0xa1, 0x0a, 0xcc, 0xa4, 0x04, 0xcb, 0x6c, 0x01, 0xb2, 0xed, 0x2e, 0xdf, 0x82,
	0x68, 0x74, 0xd7, 0x04, 0x8f, 0x95, 0x7a, 0xd4, 0xe5, 0x18, 0xa2, 0x5f, 0x60, 0x9c, 0xfc, 0x24,
	0x58, 0x04, 0xab, 0xdf, 0x24, 0x41, 0x7d, 0x2b, 0xb8, 0x88, 0xc8, 0xcf, 0xdd, 0x77, 0x01, 0xf7,
	0x2e, 0xdf, 0xc3, 0x45, 0xf1, 0x5b, 0x59, 0x57, 0xfd, 0xdd, 0xdd, 0x10, 0x4c, 0x73, 0x8e, 0x51,
	0x34, 0xa8, 0x75, 0x39, 0x09, 0x17, 0xc1, 0x6a, 0x44, 0x66, 0x9d, 0x59, 0xb7, 0x0c, 0xda, 0x5b,
	0x5d, 0x48, 0x91, 0xb0, 0xb2, 0xe6, 0xbb, 0x06, 0x5c, 0x7f, 0x06, 0xd1, 0xf2, 0x08, 0x15, 0xfa,
	0xf9, 0x62, 0xd7, 0x7f, 0xbe, 0x2d, 0xb5, 0x6d, 0xc4, 0xb6, 0xc1, 0xd3, 0x4d, 0x3b, 0x23, 0xa0,
	0x64, 0x52, 0x20, 0xd0, 0x02, 0x0b, 0x2e, 0x9d, 0x55, 0x77, 0xf9, 0xaa, 0x30, 0x7d, 0xef, 0xfb,
	0xcf, 0x1f, 0x2f, 0xe1, 0x60, 0x13, 0xc7, 0xaf, 0xe1, 0x7c, 0xe3, 0xc5, 0xe2, 0xdc, 0x20, 0x5f,
	0x36, 0x55, 0x42, 0xd0, 0x7f, 0x87, 0xbd, 0x77, 0x40, 0x1a, 0xe7, 0x26, 0xbd, 0x02, 0x69, 0x42,
	0x52, 0x0f, 0x7c, 0x84, 0x4b, 0xff, 0x97, 0xd2, 0x38, 0x37, 0x94, 0x5e, 0x11, 0x4a, 0x13, 0x42,
	0xa9, 0x87, 0x0e, 0x43, 0x97, 0xee, 0xef, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xda, 0x7a,
	0xf3, 0x7c, 0x02, 0x00, 0x00,
}
