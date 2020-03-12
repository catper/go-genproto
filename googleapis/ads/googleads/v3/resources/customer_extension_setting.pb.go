// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/customer_extension_setting.proto

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

// A customer extension setting.
type CustomerExtensionSetting struct {
	// The resource name of the customer extension setting.
	// CustomerExtensionSetting resource names have the form:
	//
	// `customers/{customer_id}/customerExtensionSettings/{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The extension type of the customer extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v3.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// The resource names of the extension feed items to serve under the customer.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []*wrappers.StringValue `protobuf:"bytes,3,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device               enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,4,opt,name=device,proto3,enum=google.ads.googleads.v3.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *CustomerExtensionSetting) Reset()         { *m = CustomerExtensionSetting{} }
func (m *CustomerExtensionSetting) String() string { return proto.CompactTextString(m) }
func (*CustomerExtensionSetting) ProtoMessage()    {}
func (*CustomerExtensionSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_aad619780d8a95dd, []int{0}
}

func (m *CustomerExtensionSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerExtensionSetting.Unmarshal(m, b)
}
func (m *CustomerExtensionSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerExtensionSetting.Marshal(b, m, deterministic)
}
func (m *CustomerExtensionSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerExtensionSetting.Merge(m, src)
}
func (m *CustomerExtensionSetting) XXX_Size() int {
	return xxx_messageInfo_CustomerExtensionSetting.Size(m)
}
func (m *CustomerExtensionSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerExtensionSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerExtensionSetting proto.InternalMessageInfo

func (m *CustomerExtensionSetting) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *CustomerExtensionSetting) GetExtensionFeedItems() []*wrappers.StringValue {
	if m != nil {
		return m.ExtensionFeedItems
	}
	return nil
}

func (m *CustomerExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if m != nil {
		return m.Device
	}
	return enums.ExtensionSettingDeviceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CustomerExtensionSetting)(nil), "google.ads.googleads.v3.resources.CustomerExtensionSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/customer_extension_setting.proto", fileDescriptor_aad619780d8a95dd)
}

var fileDescriptor_aad619780d8a95dd = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x25, 0x59, 0x29, 0x18, 0x6d, 0x1f, 0x82, 0x0f, 0xb1, 0x54, 0xd9, 0x2a, 0x85, 0x7d, 0x9a,
	0xc1, 0xcd, 0xdb, 0x28, 0x42, 0x56, 0x6b, 0x51, 0xa1, 0x2c, 0xa9, 0xec, 0x83, 0x2c, 0x84, 0x69,
	0x72, 0x1b, 0x03, 0x9b, 0x99, 0x30, 0x33, 0x59, 0x2d, 0x52, 0xf0, 0xc1, 0x27, 0x7f, 0x86, 0x8f,
	0xfe, 0x14, 0xff, 0x88, 0xd0, 0x5f, 0x21, 0x9b, 0xf9, 0xd8, 0x5d, 0xcb, 0xda, 0x7d, 0xbb, 0x33,
	0xf7, 0x9c, 0x7b, 0xce, 0x9c, 0xdc, 0x04, 0xa3, 0x92, 0xf3, 0x72, 0x06, 0x98, 0x16, 0x12, 0xeb,
	0x72, 0x51, 0xcd, 0x63, 0x2c, 0x40, 0xf2, 0x56, 0xe4, 0x20, 0x71, 0xde, 0x4a, 0xc5, 0x6b, 0x10,
	0x19, 0x7c, 0x51, 0xc0, 0x64, 0xc5, 0x59, 0x26, 0x41, 0xa9, 0x8a, 0x95, 0xa8, 0x11, 0x5c, 0xf1,
	0xf0, 0x50, 0x13, 0x11, 0x2d, 0x24, 0x72, 0x33, 0xd0, 0x3c, 0x46, 0x6e, 0xc6, 0xfe, 0x8b, 0x4d,
	0x32, 0xc0, 0xda, 0x5a, 0xe2, 0x1b, 0x93, 0xb3, 0x02, 0xe6, 0x55, 0x0e, 0x5a, 0x60, 0x7f, 0xb8,
	0x2d, 0x5b, 0x5d, 0x36, 0x96, 0xf3, 0xd0, 0x72, 0x9a, 0xca, 0xbd, 0xc5, 0xb4, 0x1e, 0x9b, 0x56,
	0x77, 0x3a, 0x6f, 0x2f, 0xf0, 0x67, 0x41, 0x9b, 0x06, 0x84, 0x34, 0xfd, 0x83, 0x15, 0x2a, 0x65,
	0x8c, 0x2b, 0xaa, 0x2a, 0xce, 0x4c, 0xf7, 0xc9, 0x9f, 0x5e, 0x10, 0xbd, 0x32, 0x91, 0x1c, 0x5b,
	0xe5, 0x33, 0x6d, 0x3b, 0x7c, 0x1a, 0xec, 0x5a, 0xb1, 0x8c, 0xd1, 0x1a, 0x22, 0xaf, 0xef, 0x0d,
	0xee, 0xa6, 0xf7, 0xed, 0xe5, 0x29, 0xad, 0x21, 0x84, 0x60, 0x6f, 0xdd, 0x72, 0xe4, 0xf7, 0xbd,
	0xc1, 0xde, 0xf0, 0x25, 0xda, 0x14, 0x64, 0xf7, 0x4e, 0xe4, 0xd4, 0x3e, 0x5c, 0x36, 0x70, 0xcc,
	0xda, 0x7a, 0xfd, 0x26, 0xdd, 0x85, 0xd5, 0x63, 0x78, 0x1a, 0x3c, 0x58, 0xca, 0x5c, 0x00, 0x14,
	0x59, 0xa5, 0xa0, 0x96, 0x51, 0xaf, 0xdf, 0x1b, 0xdc, 0x1b, 0x1e, 0x58, 0x31, 0x9b, 0x02, 0x3a,
	0x53, 0xa2, 0x62, 0xe5, 0x84, 0xce, 0x5a, 0x48, 0x43, 0xc7, 0x7c, 0x03, 0x50, 0xbc, 0x5d, 0xf0,
	0xc2, 0x4f, 0xc1, 0x8e, 0xfe, 0x2a, 0xd1, 0x9d, 0xce, 0xee, 0x78, 0x5b, 0xbb, 0x26, 0x9c, 0xd7,
	0x1d, 0x79, 0xdd, 0xf7, 0x5a, 0x2b, 0x35, 0xf3, 0xc9, 0x77, 0xef, 0x3a, 0xf9, 0xe6, 0x05, 0xcf,
	0x96, 0x43, 0x4d, 0xd5, 0x54, 0x12, 0xe5, 0xbc, 0xc6, 0x1b, 0xf3, 0x7f, 0x6f, 0x97, 0x55, 0xe2,
	0xaf, 0xb6, 0xbc, 0x72, 0x1b, 0xfc, 0x2f, 0x7c, 0x05, 0x74, 0x73, 0xb9, 0xaf, 0x46, 0x3f, 0xfc,
	0xe0, 0x28, 0xe7, 0x35, 0xba, 0x75, 0xbd, 0x47, 0x8f, 0x36, 0x19, 0x1a, 0x2f, 0xc2, 0x1d, 0x7b,
	0x1f, 0xdf, 0x99, 0x19, 0x25, 0x9f, 0x51, 0x56, 0x22, 0x2e, 0x4a, 0x5c, 0x02, 0xeb, 0xa2, 0xc7,
	0xcb, 0xb7, 0xfd, 0xe7, 0x2f, 0x7c, 0xee, 0xaa, 0x9f, 0x7e, 0xef, 0x24, 0x49, 0x7e, 0xf9, 0x87,
	0x27, 0x7a, 0x64, 0x52, 0x48, 0xa4, 0xcb, 0x45, 0x35, 0x89, 0x51, 0x6a, 0x91, 0xbf, 0x2d, 0x66,
	0x9a, 0x14, 0x72, 0xea, 0x30, 0xd3, 0x49, 0x3c, 0x75, 0x98, 0x6b, 0xff, 0x48, 0x37, 0x08, 0x49,
	0x0a, 0x49, 0x88, 0x43, 0x11, 0x32, 0x89, 0x09, 0x71, 0xb8, 0xf3, 0x9d, 0xce, 0x6c, 0xfc, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xd3, 0x84, 0xd4, 0x16, 0x31, 0x04, 0x00, 0x00,
}
