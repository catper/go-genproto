// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/mobile_device_constant.proto

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

// A mobile device constant.
type MobileDeviceConstant struct {
	// The resource name of the mobile device constant.
	// Mobile device constant resource names have the form:
	//
	// `mobileDeviceConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the mobile device constant.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the mobile device.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The manufacturer of the mobile device.
	ManufacturerName *wrappers.StringValue `protobuf:"bytes,4,opt,name=manufacturer_name,json=manufacturerName,proto3" json:"manufacturer_name,omitempty"`
	// The operating system of the mobile device.
	OperatingSystemName *wrappers.StringValue `protobuf:"bytes,5,opt,name=operating_system_name,json=operatingSystemName,proto3" json:"operating_system_name,omitempty"`
	// The type of mobile device.
	Type                 enums.MobileDeviceTypeEnum_MobileDeviceType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.MobileDeviceTypeEnum_MobileDeviceType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *MobileDeviceConstant) Reset()         { *m = MobileDeviceConstant{} }
func (m *MobileDeviceConstant) String() string { return proto.CompactTextString(m) }
func (*MobileDeviceConstant) ProtoMessage()    {}
func (*MobileDeviceConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5b5d0a29b45d355, []int{0}
}

func (m *MobileDeviceConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileDeviceConstant.Unmarshal(m, b)
}
func (m *MobileDeviceConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileDeviceConstant.Marshal(b, m, deterministic)
}
func (m *MobileDeviceConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileDeviceConstant.Merge(m, src)
}
func (m *MobileDeviceConstant) XXX_Size() int {
	return xxx_messageInfo_MobileDeviceConstant.Size(m)
}
func (m *MobileDeviceConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileDeviceConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileDeviceConstant proto.InternalMessageInfo

func (m *MobileDeviceConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileDeviceConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileDeviceConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MobileDeviceConstant) GetManufacturerName() *wrappers.StringValue {
	if m != nil {
		return m.ManufacturerName
	}
	return nil
}

func (m *MobileDeviceConstant) GetOperatingSystemName() *wrappers.StringValue {
	if m != nil {
		return m.OperatingSystemName
	}
	return nil
}

func (m *MobileDeviceConstant) GetType() enums.MobileDeviceTypeEnum_MobileDeviceType {
	if m != nil {
		return m.Type
	}
	return enums.MobileDeviceTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MobileDeviceConstant)(nil), "google.ads.googleads.v3.resources.MobileDeviceConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/mobile_device_constant.proto", fileDescriptor_c5b5d0a29b45d355)
}

var fileDescriptor_c5b5d0a29b45d355 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6b, 0x14, 0x31,
	0x14, 0x65, 0x66, 0xd7, 0x82, 0xf1, 0x03, 0x1d, 0x15, 0xb6, 0xb5, 0xc8, 0x56, 0x29, 0x2c, 0x88,
	0x19, 0xe9, 0x48, 0x1f, 0x22, 0x08, 0x53, 0x2b, 0xa5, 0x82, 0xb2, 0x6c, 0x65, 0x11, 0x59, 0x18,
	0xb2, 0x33, 0xb7, 0x43, 0x60, 0x93, 0x0c, 0x49, 0x66, 0x65, 0x11, 0x5f, 0xfd, 0x21, 0x3e, 0xfa,
	0x53, 0xfc, 0x29, 0x7d, 0xf7, 0x5d, 0x36, 0x99, 0x49, 0x17, 0xbb, 0xb5, 0x7d, 0x3b, 0x93, 0x7b,
	0xce, 0xb9, 0x27, 0xf7, 0x66, 0xd0, 0x9b, 0x52, 0xca, 0x72, 0x06, 0x31, 0x2d, 0x74, 0xec, 0xe0,
	0x12, 0xcd, 0x93, 0x58, 0x81, 0x96, 0xb5, 0xca, 0x41, 0xc7, 0x5c, 0x4e, 0xd9, 0x0c, 0xb2, 0x02,
	0xe6, 0x2c, 0x87, 0x2c, 0x97, 0x42, 0x1b, 0x2a, 0x0c, 0xae, 0x94, 0x34, 0x32, 0xda, 0x71, 0x22,
	0x4c, 0x0b, 0x8d, 0xbd, 0x1e, 0xcf, 0x13, 0xec, 0xf5, 0x5b, 0xfb, 0x97, 0xb5, 0x00, 0x51, 0xf3,
	0x7f, 0xed, 0xcd, 0xa2, 0x02, 0x67, 0xbd, 0xb5, 0xd9, 0xea, 0x2a, 0xe6, 0xd3, 0x34, 0xa5, 0x27,
	0x4d, 0xc9, 0x7e, 0x4d, 0xeb, 0xd3, 0xf8, 0xab, 0xa2, 0x55, 0x05, 0x4a, 0x37, 0xf5, 0xed, 0x15,
	0x29, 0x15, 0x42, 0x1a, 0x6a, 0x98, 0x14, 0x4d, 0xf5, 0xe9, 0x9f, 0x0e, 0x7a, 0xf8, 0xc1, 0x76,
	0x3d, 0xb4, 0x4d, 0xdf, 0x36, 0x57, 0x8a, 0x9e, 0xa1, 0x3b, 0x6d, 0xa3, 0x4c, 0x50, 0x0e, 0xbd,
	0xa0, 0x1f, 0x0c, 0x6e, 0x8e, 0x6e, 0xb7, 0x87, 0x1f, 0x29, 0x87, 0xe8, 0x39, 0x0a, 0x59, 0xd1,
	0x0b, 0xfb, 0xc1, 0xe0, 0xd6, 0xde, 0xe3, 0xe6, 0xce, 0xb8, 0x0d, 0x82, 0x8f, 0x85, 0xd9, 0x7f,
	0x35, 0xa6, 0xb3, 0x1a, 0x46, 0x21, 0x2b, 0xa2, 0x97, 0xa8, 0x6b, 0x8d, 0x3a, 0x96, 0xbe, 0x7d,
	0x81, 0x7e, 0x62, 0x14, 0x13, 0xa5, 0xe3, 0x5b, 0x66, 0x74, 0x8c, 0xee, 0x73, 0x2a, 0xea, 0x53,
	0x9a, 0x9b, 0x5a, 0x81, 0x72, 0x39, 0xba, 0xd7, 0x90, 0xdf, 0x5b, 0x95, 0xd9, 0xa4, 0x43, 0xf4,
	0x48, 0x56, 0xa0, 0xa8, 0x61, 0xa2, 0xcc, 0xf4, 0x42, 0x1b, 0xe0, 0xce, 0xee, 0xc6, 0x35, 0xec,
	0x1e, 0x78, 0xe9, 0x89, 0x55, 0x5a, 0xc7, 0xcf, 0xa8, 0xbb, 0x5c, 0x50, 0x6f, 0xa3, 0x1f, 0x0c,
	0xee, 0xee, 0x1d, 0xe2, 0xcb, 0x96, 0x6f, 0x37, 0x8b, 0x57, 0x67, 0xfc, 0x69, 0x51, 0xc1, 0x3b,
	0x51, 0xf3, 0x0b, 0x87, 0x23, 0xeb, 0x48, 0xa6, 0x67, 0x69, 0x86, 0x5e, 0x9c, 0x9b, 0x34, 0xa8,
	0x62, 0x1a, 0xe7, 0x92, 0xc7, 0x6b, 0xd7, 0x85, 0xf9, 0x9a, 0x53, 0x1d, 0x7f, 0x5b, 0xff, 0x60,
	0xbf, 0x1f, 0xfc, 0x08, 0xd1, 0x6e, 0x2e, 0x39, 0xbe, 0xf2, 0xc9, 0x1e, 0x6c, 0xae, 0xeb, 0x37,
	0x5c, 0x8e, 0x69, 0x18, 0x7c, 0x79, 0xdf, 0xe8, 0x4b, 0x39, 0xa3, 0xa2, 0xc4, 0x52, 0x95, 0x71,
	0x09, 0xc2, 0x0e, 0x31, 0x3e, 0x8f, 0xfd, 0x9f, 0x3f, 0xea, 0xb5, 0x47, 0x3f, 0xc3, 0xce, 0x51,
	0x9a, 0xfe, 0x0a, 0x77, 0x8e, 0x9c, 0x65, 0x5a, 0x68, 0xec, 0xe0, 0x12, 0x8d, 0x13, 0x3c, 0x6a,
	0x99, 0xbf, 0x5b, 0xce, 0x24, 0x2d, 0xf4, 0xc4, 0x73, 0x26, 0xe3, 0x64, 0xe2, 0x39, 0x67, 0xe1,
	0xae, 0x2b, 0x10, 0x92, 0x16, 0x9a, 0x10, 0xcf, 0x22, 0x64, 0x9c, 0x10, 0xe2, 0x79, 0xd3, 0x0d,
	0x1b, 0x36, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xba, 0x32, 0x30, 0xd1, 0xfd, 0x03, 0x00, 0x00,
}
