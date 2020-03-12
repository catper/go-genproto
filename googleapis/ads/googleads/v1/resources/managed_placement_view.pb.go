// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/managed_placement_view.proto

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

// A managed placement view.
type ManagedPlacementView struct {
	// Immutable. The resource name of the Managed Placement view.
	// Managed placement view resource names have the form:
	//
	// `customers/{customer_id}/managedPlacementViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagedPlacementView) Reset()         { *m = ManagedPlacementView{} }
func (m *ManagedPlacementView) String() string { return proto.CompactTextString(m) }
func (*ManagedPlacementView) ProtoMessage()    {}
func (*ManagedPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_6953f60267bddfaf, []int{0}
}

func (m *ManagedPlacementView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagedPlacementView.Unmarshal(m, b)
}
func (m *ManagedPlacementView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagedPlacementView.Marshal(b, m, deterministic)
}
func (m *ManagedPlacementView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagedPlacementView.Merge(m, src)
}
func (m *ManagedPlacementView) XXX_Size() int {
	return xxx_messageInfo_ManagedPlacementView.Size(m)
}
func (m *ManagedPlacementView) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagedPlacementView.DiscardUnknown(m)
}

var xxx_messageInfo_ManagedPlacementView proto.InternalMessageInfo

func (m *ManagedPlacementView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ManagedPlacementView)(nil), "google.ads.googleads.v1.resources.ManagedPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/managed_placement_view.proto", fileDescriptor_6953f60267bddfaf)
}

var fileDescriptor_6953f60267bddfaf = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xc3, 0x40,
	0x1c, 0xc5, 0x49, 0x44, 0xc1, 0xa0, 0x4b, 0x71, 0xb0, 0x45, 0xd0, 0x0a, 0x05, 0x17, 0xef, 0x08,
	0xe2, 0x72, 0x82, 0x70, 0x75, 0x28, 0x08, 0x4a, 0xe9, 0x90, 0xa1, 0x04, 0xc2, 0x35, 0xf9, 0x1b,
	0x03, 0xb9, 0xbb, 0x90, 0x4b, 0xd3, 0xa1, 0x74, 0xf5, 0x83, 0x38, 0xfa, 0x51, 0xfc, 0x14, 0xce,
	0xf5, 0x1b, 0x38, 0x49, 0x7b, 0xbd, 0x6b, 0x85, 0xa2, 0xb8, 0x3d, 0xf8, 0xff, 0xde, 0xcb, 0xcb,
	0xe3, 0xbc, 0xdb, 0x54, 0xca, 0x34, 0x07, 0xcc, 0x12, 0x85, 0xb5, 0x5c, 0xa8, 0xda, 0xc7, 0x25,
	0x28, 0x39, 0x2e, 0x63, 0x50, 0x98, 0x33, 0xc1, 0x52, 0x48, 0xa2, 0x22, 0x67, 0x31, 0x70, 0x10,
	0x55, 0x54, 0x67, 0x30, 0x41, 0x45, 0x29, 0x2b, 0xd9, 0x68, 0x6b, 0x13, 0x62, 0x89, 0x42, 0xd6,
	0x8f, 0x6a, 0x1f, 0x59, 0x7f, 0xeb, 0xd4, 0x7c, 0xa2, 0xc8, 0xf0, 0x53, 0x06, 0x79, 0x12, 0x8d,
	0xe0, 0x99, 0xd5, 0x99, 0x2c, 0x75, 0x46, 0xab, 0xb9, 0x01, 0x18, 0xdb, 0xea, 0x74, 0xb2, 0x71,
	0x62, 0x42, 0xc8, 0x8a, 0x55, 0x99, 0x14, 0x4a, 0x5f, 0xcf, 0x3f, 0x1d, 0xef, 0xe8, 0x41, 0xb7,
	0xeb, 0x9b, 0x72, 0x41, 0x06, 0x93, 0xc6, 0xd0, 0x3b, 0x34, 0x41, 0x91, 0x60, 0x1c, 0x8e, 0x9d,
	0x33, 0xe7, 0x62, 0xbf, 0x7b, 0xfd, 0x41, 0x77, 0xbf, 0x28, 0xf6, 0x2e, 0xd7, 0x4d, 0x57, 0xaa,
	0xc8, 0x14, 0x8a, 0x25, 0xc7, 0xdb, 0xd2, 0x06, 0x07, 0x26, 0xeb, 0x91, 0x71, 0x20, 0x93, 0x39,
	0xad, 0xfe, 0x99, 0xd0, 0xb8, 0x8b, 0xc7, 0xaa, 0x92, 0x1c, 0x4a, 0x85, 0xa7, 0x46, 0xce, 0xcc,
	0xb0, 0x3f, 0x50, 0x85, 0xa7, 0xdb, 0xf7, 0x9e, 0x75, 0x5f, 0x5c, 0xaf, 0x13, 0x4b, 0x8e, 0xfe,
	0x5c, 0xbc, 0xdb, 0xdc, 0x56, 0xa2, 0xbf, 0x98, 0xac, 0xef, 0x0c, 0xef, 0x57, 0xfe, 0x54, 0xe6,
	0x4c, 0xa4, 0x48, 0x96, 0x29, 0x4e, 0x41, 0x2c, 0x07, 0xc5, 0xeb, 0x7f, 0xf9, 0xe5, 0x41, 0xdc,
	0x58, 0xf5, 0xea, 0xee, 0xf4, 0x28, 0x7d, 0x73, 0xdb, 0x3d, 0x1d, 0x49, 0x13, 0x85, 0xb4, 0x5c,
	0xa8, 0xc0, 0x47, 0x03, 0x43, 0xbe, 0x1b, 0x26, 0xa4, 0x89, 0x0a, 0x2d, 0x13, 0x06, 0x7e, 0x68,
	0x99, 0xb9, 0xdb, 0xd1, 0x07, 0x42, 0x68, 0xa2, 0x08, 0xb1, 0x14, 0x21, 0x81, 0x4f, 0x88, 0xe5,
	0x46, 0x7b, 0xcb, 0xb2, 0x57, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0xb8, 0xd9, 0x83, 0xbc,
	0x02, 0x00, 0x00,
}
