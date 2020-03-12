// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/detail_placement_view.proto

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

// A view with metrics aggregated by ad group and URL or YouTube video.
type DetailPlacementView struct {
	// The resource name of the detail placement view.
	// Detail placement view resource names have the form:
	//
	// `customers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The automatic placement string at detail level, e. g. website URL, mobile
	// application ID, or a YouTube video ID.
	Placement *wrappers.StringValue `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	// The display name is URL name for websites, YouTube video name for YouTube
	// videos, and translated mobile app name for mobile apps.
	DisplayName *wrappers.StringValue `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// URL of the group placement, e.g. domain, link to the mobile application in
	// app store, or a YouTube channel URL.
	GroupPlacementTargetUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=group_placement_target_url,json=groupPlacementTargetUrl,proto3" json:"group_placement_target_url,omitempty"`
	// URL of the placement, e.g. website, link to the mobile application in app
	// store, or a YouTube video URL.
	TargetUrl *wrappers.StringValue `protobuf:"bytes,5,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
	// Type of the placement, e.g. Website, YouTube Video, and Mobile Application.
	PlacementType        enums.PlacementTypeEnum_PlacementType `protobuf:"varint,6,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v3.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DetailPlacementView) Reset()         { *m = DetailPlacementView{} }
func (m *DetailPlacementView) String() string { return proto.CompactTextString(m) }
func (*DetailPlacementView) ProtoMessage()    {}
func (*DetailPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_55e3b8e503380d64, []int{0}
}

func (m *DetailPlacementView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailPlacementView.Unmarshal(m, b)
}
func (m *DetailPlacementView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailPlacementView.Marshal(b, m, deterministic)
}
func (m *DetailPlacementView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailPlacementView.Merge(m, src)
}
func (m *DetailPlacementView) XXX_Size() int {
	return xxx_messageInfo_DetailPlacementView.Size(m)
}
func (m *DetailPlacementView) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailPlacementView.DiscardUnknown(m)
}

var xxx_messageInfo_DetailPlacementView proto.InternalMessageInfo

func (m *DetailPlacementView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DetailPlacementView) GetPlacement() *wrappers.StringValue {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *DetailPlacementView) GetDisplayName() *wrappers.StringValue {
	if m != nil {
		return m.DisplayName
	}
	return nil
}

func (m *DetailPlacementView) GetGroupPlacementTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.GroupPlacementTargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.TargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetPlacementType() enums.PlacementTypeEnum_PlacementType {
	if m != nil {
		return m.PlacementType
	}
	return enums.PlacementTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*DetailPlacementView)(nil), "google.ads.googleads.v3.resources.DetailPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/detail_placement_view.proto", fileDescriptor_55e3b8e503380d64)
}

var fileDescriptor_55e3b8e503380d64 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0x26, 0x6e, 0x57, 0xa8, 0xfa, 0x73, 0xe1, 0x5d, 0xcc, 0x0b, 0x65, 0xa4, 0x1b, 0x85, 0x5c,
	0x0c, 0x09, 0xe2, 0x3b, 0x95, 0x6d, 0xb8, 0x6c, 0x14, 0x76, 0x31, 0x42, 0xb6, 0x19, 0x36, 0x02,
	0x41, 0xb5, 0xcf, 0x8c, 0xc1, 0x96, 0x34, 0x49, 0x4e, 0x08, 0xa5, 0x97, 0x7b, 0x91, 0x5d, 0x8e,
	0x3d, 0xc9, 0x1e, 0xa5, 0x4f, 0x31, 0x62, 0xcb, 0x72, 0x02, 0xe9, 0x96, 0xbb, 0x4f, 0x3a, 0xdf,
	0xf7, 0x9d, 0x73, 0x74, 0x8e, 0xd0, 0xab, 0x4c, 0x88, 0xac, 0x00, 0xc2, 0x52, 0x4d, 0x1a, 0xb8,
	0x42, 0xf3, 0x90, 0x28, 0xd0, 0xa2, 0x52, 0x09, 0x68, 0x92, 0x82, 0x61, 0x79, 0x31, 0x93, 0x05,
	0x4b, 0xa0, 0x04, 0x6e, 0x66, 0xf3, 0x1c, 0x16, 0x58, 0x2a, 0x61, 0x84, 0x7f, 0xde, 0x68, 0x30,
	0x4b, 0x35, 0x76, 0x72, 0x3c, 0x0f, 0xb1, 0x93, 0xf7, 0x47, 0x0f, 0x65, 0x00, 0x5e, 0x95, 0x9a,
	0x74, 0xb6, 0x66, 0x29, 0xa1, 0xb1, 0xed, 0x3f, 0x6d, 0x35, 0x32, 0x77, 0x85, 0xd8, 0xd0, 0x33,
	0x1b, 0xaa, 0x4f, 0x37, 0xd5, 0x37, 0xb2, 0x50, 0x4c, 0x4a, 0x50, 0xda, 0xc6, 0xcf, 0xd6, 0xa4,
	0x8c, 0x73, 0x61, 0x98, 0xc9, 0x05, 0xb7, 0xd1, 0xe7, 0xbf, 0xf7, 0xd1, 0xe3, 0xb7, 0x75, 0x3f,
	0xe3, 0x36, 0x6f, 0x9c, 0xc3, 0xc2, 0x7f, 0x81, 0x4e, 0xda, 0x3c, 0x33, 0xce, 0x4a, 0x08, 0x7a,
	0x83, 0xde, 0xf0, 0x70, 0x72, 0xdc, 0x5e, 0x7e, 0x60, 0x25, 0xf8, 0x14, 0x1d, 0xba, 0x6a, 0x03,
	0x6f, 0xd0, 0x1b, 0x1e, 0x8d, 0xce, 0x6c, 0xd7, 0xb8, 0x2d, 0x07, 0x7f, 0x34, 0x2a, 0xe7, 0x59,
	0xcc, 0x8a, 0x0a, 0x26, 0x1d, 0xdd, 0x7f, 0x83, 0x8e, 0xd3, 0x5c, 0xcb, 0x82, 0x2d, 0x1b, 0xff,
	0xbd, 0x1d, 0xe4, 0x47, 0x56, 0x51, 0x27, 0xff, 0x82, 0xfa, 0x99, 0x12, 0x95, 0x5c, 0x9b, 0x83,
	0x61, 0x2a, 0x03, 0x33, 0xab, 0x54, 0x11, 0xec, 0xef, 0x60, 0xf7, 0xa4, 0xd6, 0xbb, 0xbe, 0x3f,
	0xd5, 0xea, 0xcf, 0xaa, 0xf0, 0x2f, 0x11, 0x5a, 0xb3, 0x7a, 0xb4, 0x4b, 0x63, 0xc6, 0x89, 0x01,
	0x9d, 0x6e, 0x8e, 0x30, 0x38, 0x18, 0xf4, 0x86, 0xa7, 0xa3, 0xd7, 0xf8, 0xa1, 0xd5, 0xa8, 0xe7,
	0x8e, 0xbb, 0x3a, 0x96, 0x12, 0xde, 0xf1, 0xaa, 0xdc, 0xbc, 0x99, 0x9c, 0xc8, 0xf5, 0x23, 0x35,
	0xf7, 0xd1, 0x77, 0xf4, 0xb2, 0xf3, 0xb1, 0x48, 0xe6, 0x1a, 0x27, 0xa2, 0x24, 0xdb, 0x66, 0x1a,
	0x25, 0x95, 0x36, 0xa2, 0x04, 0xa5, 0xc9, 0x6d, 0x0b, 0xef, 0xec, 0x36, 0x6f, 0x30, 0x35, 0xb9,
	0xdd, 0xba, 0xe3, 0x77, 0x57, 0x3f, 0x3c, 0x74, 0x91, 0x88, 0x12, 0xff, 0x77, 0xcb, 0xaf, 0x82,
	0x2d, 0x15, 0x8c, 0x57, 0x4f, 0x37, 0xee, 0x7d, 0x7d, 0x6f, 0xe5, 0x99, 0x28, 0x18, 0xcf, 0xb0,
	0x50, 0x19, 0xc9, 0x80, 0xd7, 0x0f, 0x4b, 0xba, 0x3e, 0xfe, 0xf1, 0x05, 0x2f, 0x1d, 0xfa, 0xe9,
	0xed, 0x5d, 0x47, 0xd1, 0x2f, 0xef, 0xfc, 0xba, 0xb1, 0x8c, 0x52, 0x8d, 0x1b, 0xb8, 0x42, 0x71,
	0x88, 0x27, 0x2d, 0xf3, 0x4f, 0xcb, 0x99, 0x46, 0xa9, 0x9e, 0x3a, 0xce, 0x34, 0x0e, 0xa7, 0x8e,
	0x73, 0xef, 0x5d, 0x34, 0x01, 0x4a, 0xa3, 0x54, 0x53, 0xea, 0x58, 0x94, 0xc6, 0x21, 0xa5, 0x8e,
	0x77, 0x73, 0x50, 0x17, 0x1b, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x28, 0xbe, 0xdd, 0xc6, 0x2e,
	0x04, 0x00, 0x00,
}
