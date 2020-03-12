// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/merchant_center_link.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// A data sharing connection, proposed or in use,
// between a Google Ads Customer and a Merchant Center account.
type MerchantCenterLink struct {
	// The resource name of the merchant center link.
	// Merchant center link resource names have the form:
	//
	// `customers/{customer_id}/merchantCenterLinks/{merchant_center_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the Merchant Center account.
	// This field is readonly.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Merchant Center account.
	// This field is readonly.
	MerchantCenterAccountName *wrappers.StringValue `protobuf:"bytes,4,opt,name=merchant_center_account_name,json=merchantCenterAccountName,proto3" json:"merchant_center_account_name,omitempty"`
	// The status of the link.
	Status               enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *MerchantCenterLink) Reset()         { *m = MerchantCenterLink{} }
func (m *MerchantCenterLink) String() string { return proto.CompactTextString(m) }
func (*MerchantCenterLink) ProtoMessage()    {}
func (*MerchantCenterLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_5922ba0a523027c6, []int{0}
}

func (m *MerchantCenterLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCenterLink.Unmarshal(m, b)
}
func (m *MerchantCenterLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCenterLink.Marshal(b, m, deterministic)
}
func (m *MerchantCenterLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCenterLink.Merge(m, src)
}
func (m *MerchantCenterLink) XXX_Size() int {
	return xxx_messageInfo_MerchantCenterLink.Size(m)
}
func (m *MerchantCenterLink) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCenterLink.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCenterLink proto.InternalMessageInfo

func (m *MerchantCenterLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MerchantCenterLink) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MerchantCenterLink) GetMerchantCenterAccountName() *wrappers.StringValue {
	if m != nil {
		return m.MerchantCenterAccountName
	}
	return nil
}

func (m *MerchantCenterLink) GetStatus() enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.MerchantCenterLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MerchantCenterLink)(nil), "google.ads.googleads.v3.resources.MerchantCenterLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/merchant_center_link.proto", fileDescriptor_5922ba0a523027c6)
}

var fileDescriptor_5922ba0a523027c6 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x27, 0x59, 0x2d, 0x18, 0xff, 0x1c, 0x72, 0x71, 0x5b, 0x17, 0xd9, 0x2a, 0x85, 0x85, 0xc2,
	0x0c, 0x34, 0xe2, 0x61, 0x14, 0x4a, 0x2a, 0x52, 0x14, 0x95, 0xb2, 0x95, 0x3d, 0xc8, 0xca, 0x32,
	0x4d, 0xc6, 0x18, 0x9a, 0x79, 0x13, 0x67, 0x26, 0xf5, 0x50, 0x7a, 0xf3, 0x93, 0x78, 0xf0, 0xe0,
	0x47, 0xf1, 0xa3, 0xf4, 0x53, 0xc8, 0xce, 0xbf, 0xaa, 0xd9, 0xb5, 0xb7, 0x5f, 0xf2, 0x7e, 0x7f,
	0xde, 0x7b, 0x79, 0x49, 0x9e, 0x57, 0x42, 0x54, 0x0d, 0xc3, 0xb4, 0x54, 0xd8, 0xc2, 0x25, 0x3a,
	0xcb, 0xb0, 0x64, 0x4a, 0x74, 0xb2, 0x60, 0x0a, 0x73, 0x26, 0x8b, 0xcf, 0x14, 0xf4, 0xa2, 0x60,
	0xa0, 0x99, 0x5c, 0x34, 0x35, 0x9c, 0xa2, 0x56, 0x0a, 0x2d, 0xd2, 0x6d, 0x2b, 0x41, 0xb4, 0x54,
	0x28, 0xa8, 0xd1, 0x59, 0x86, 0x82, 0x7a, 0x6b, 0x7f, 0x5d, 0x00, 0x83, 0x8e, 0xaf, 0x36, 0x5f,
	0x28, 0x4d, 0x75, 0xa7, 0x6c, 0xc6, 0xd6, 0xa6, 0x37, 0x68, 0xeb, 0xd0, 0x94, 0x2b, 0x3d, 0x74,
	0x25, 0xf3, 0x74, 0xd2, 0x7d, 0xc2, 0x5f, 0x25, 0x6d, 0x5b, 0x26, 0xbd, 0x74, 0xf4, 0x87, 0x94,
	0x02, 0x08, 0x4d, 0x75, 0x2d, 0xc0, 0x55, 0x1f, 0xfd, 0x18, 0x24, 0xe9, 0x5b, 0x17, 0xff, 0xc2,
	0xa4, 0xbf, 0xa9, 0xe1, 0x34, 0x7d, 0x9c, 0xdc, 0xf5, 0x31, 0x0b, 0xa0, 0x9c, 0x0d, 0xa3, 0x71,
	0x34, 0xb9, 0x35, 0xbd, 0xe3, 0x5f, 0xbe, 0xa3, 0x9c, 0xa5, 0xbb, 0x49, 0x5c, 0x97, 0xc3, 0xc1,
	0x38, 0x9a, 0xdc, 0xde, 0x7b, 0xe0, 0x46, 0x47, 0xbe, 0x0d, 0xf4, 0x0a, 0xf4, 0xd3, 0x27, 0x33,
	0xda, 0x74, 0x6c, 0x1a, 0xd7, 0x65, 0xfa, 0x31, 0x19, 0xfd, 0x3b, 0x26, 0x2d, 0x0a, 0xd1, 0x81,
	0xb6, 0x01, 0x37, 0x8c, 0xcd, 0xa8, 0x67, 0x73, 0xac, 0x65, 0x0d, 0x95, 0xf5, 0xd9, 0xe4, 0x7f,
	0x75, 0x9a, 0x5b, 0xbd, 0xe9, 0xa5, 0x49, 0x36, 0xec, 0xc2, 0x86, 0x37, 0xc7, 0xd1, 0xe4, 0xde,
	0xde, 0x7b, 0xb4, 0xee, 0xab, 0x98, 0x95, 0xa3, 0xfe, 0xcc, 0xc7, 0x46, 0xfe, 0x12, 0x3a, 0xbe,
	0xb6, 0x38, 0x75, 0x19, 0xe4, 0xcb, 0x65, 0x0e, 0xc9, 0xee, 0x95, 0xad, 0x43, 0x6d, 0xad, 0x50,
	0x21, 0x38, 0x5e, 0xb1, 0xd0, 0xfd, 0xa2, 0x53, 0x5a, 0x70, 0x26, 0x15, 0x3e, 0xf7, 0xf0, 0x02,
	0xf3, 0x1e, 0x51, 0xe1, 0xf3, 0x55, 0xd7, 0x70, 0x71, 0xf0, 0x2d, 0x4e, 0x76, 0x0a, 0xc1, 0xd1,
	0xb5, 0xc7, 0x76, 0x70, 0xbf, 0x1f, 0x7f, 0xb4, 0xdc, 0xe6, 0x51, 0xf4, 0xe1, 0xb5, 0x53, 0x57,
	0xa2, 0xa1, 0x50, 0x21, 0x21, 0x2b, 0x5c, 0x31, 0x30, 0xbb, 0xc6, 0x57, 0x33, 0xfc, 0xe7, 0x3f,
	0x78, 0x16, 0xd0, 0xf7, 0x78, 0x70, 0x98, 0xe7, 0x3f, 0xe3, 0xed, 0x43, 0x6b, 0x99, 0x97, 0x0a,
	0x59, 0xb8, 0x44, 0xb3, 0x0c, 0x4d, 0x3d, 0xf3, 0x97, 0xe7, 0xcc, 0xf3, 0x52, 0xcd, 0x03, 0x67,
	0x3e, 0xcb, 0xe6, 0x81, 0x73, 0x19, 0xef, 0xd8, 0x02, 0x21, 0x79, 0xa9, 0x08, 0x09, 0x2c, 0x42,
	0x66, 0x19, 0x21, 0x81, 0x77, 0xb2, 0x61, 0x9a, 0xcd, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x91,
	0x97, 0x6e, 0x3d, 0xb3, 0x03, 0x00, 0x00,
}
