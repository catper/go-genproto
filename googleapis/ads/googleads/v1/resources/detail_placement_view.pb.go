// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/detail_placement_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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
	// Immutable. The resource name of the detail placement view.
	// Detail placement view resource names have the form:
	//
	// `customers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The automatic placement string at detail level, e. g. website URL, mobile
	// application ID, or a YouTube video ID.
	Placement *wrappers.StringValue `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	// Output only. The display name is URL name for websites, YouTube video name for YouTube
	// videos, and translated mobile app name for mobile apps.
	DisplayName *wrappers.StringValue `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. URL of the group placement, e.g. domain, link to the mobile application in
	// app store, or a YouTube channel URL.
	GroupPlacementTargetUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=group_placement_target_url,json=groupPlacementTargetUrl,proto3" json:"group_placement_target_url,omitempty"`
	// Output only. URL of the placement, e.g. website, link to the mobile application in app
	// store, or a YouTube video URL.
	TargetUrl *wrappers.StringValue `protobuf:"bytes,5,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
	// Output only. Type of the placement, e.g. Website, YouTube Video, and Mobile Application.
	PlacementType        enums.PlacementTypeEnum_PlacementType `protobuf:"varint,6,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v1.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DetailPlacementView) Reset()         { *m = DetailPlacementView{} }
func (m *DetailPlacementView) String() string { return proto.CompactTextString(m) }
func (*DetailPlacementView) ProtoMessage()    {}
func (*DetailPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_9703f91bccf33f03, []int{0}
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
	proto.RegisterType((*DetailPlacementView)(nil), "google.ads.googleads.v1.resources.DetailPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/detail_placement_view.proto", fileDescriptor_9703f91bccf33f03)
}

var fileDescriptor_9703f91bccf33f03 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x55, 0x92, 0xb6, 0x52, 0xdd, 0xc7, 0x62, 0x58, 0x10, 0xaa, 0x0a, 0x52, 0xa4, 0x4a, 0x59,
	0x20, 0x5b, 0x09, 0xac, 0x06, 0xf1, 0x70, 0x54, 0x54, 0x89, 0x05, 0x8a, 0x02, 0x44, 0x02, 0x45,
	0x44, 0x4e, 0xe6, 0x76, 0x30, 0x9a, 0xb1, 0x8d, 0xed, 0x49, 0x14, 0x55, 0x5d, 0xf2, 0x23, 0x2c,
	0xf9, 0x14, 0xbe, 0xa2, 0xeb, 0x7e, 0x02, 0x6c, 0x50, 0x66, 0x3c, 0x4e, 0x22, 0xa5, 0x10, 0x76,
	0x67, 0xe6, 0x9e, 0x73, 0xee, 0x3d, 0xb6, 0x2f, 0x7a, 0x16, 0x4b, 0x19, 0x27, 0x40, 0x58, 0x64,
	0x48, 0x01, 0xe7, 0x68, 0xd2, 0x22, 0x1a, 0x8c, 0xcc, 0xf4, 0x18, 0x0c, 0x89, 0xc0, 0x32, 0x9e,
	0x0c, 0x55, 0xc2, 0xc6, 0x90, 0x82, 0xb0, 0xc3, 0x09, 0x87, 0x29, 0x56, 0x5a, 0x5a, 0x19, 0x9c,
	0x14, 0x1a, 0xcc, 0x22, 0x83, 0xbd, 0x1c, 0x4f, 0x5a, 0xd8, 0xcb, 0x8f, 0xda, 0xb7, 0x75, 0x00,
	0x91, 0xa5, 0x86, 0x2c, 0x6c, 0xed, 0x4c, 0x41, 0x61, 0x7b, 0xf4, 0xa0, 0xd4, 0x28, 0x4e, 0x2e,
	0x38, 0x24, 0xd1, 0x70, 0x04, 0x9f, 0xd9, 0x84, 0x4b, 0xed, 0x08, 0xf7, 0x96, 0x08, 0x65, 0x2b,
	0x57, 0xba, 0xef, 0x4a, 0xf9, 0xd7, 0x28, 0xbb, 0x20, 0x53, 0xcd, 0x94, 0x02, 0x6d, 0x5c, 0xfd,
	0x78, 0x49, 0xca, 0x84, 0x90, 0x96, 0x59, 0x2e, 0x85, 0xab, 0x3e, 0xfc, 0xbd, 0x85, 0xee, 0x9c,
	0xe5, 0x81, 0xbb, 0xe5, 0x60, 0x7d, 0x0e, 0xd3, 0xe0, 0x03, 0x3a, 0x28, 0xfb, 0x0c, 0x05, 0x4b,
	0xa1, 0x5e, 0x69, 0x54, 0x9a, 0xbb, 0x9d, 0x27, 0xd7, 0x74, 0xfb, 0x17, 0xc5, 0xe8, 0xd1, 0x22,
	0xbc, 0x43, 0x8a, 0x1b, 0x3c, 0x96, 0x29, 0x59, 0x63, 0xd6, 0xdb, 0x2f, 0xad, 0xde, 0xb0, 0x14,
	0x82, 0x17, 0x68, 0xd7, 0x1f, 0x42, 0xbd, 0xda, 0xa8, 0x34, 0xf7, 0xda, 0xc7, 0xce, 0x05, 0x97,
	0x21, 0xf0, 0x5b, 0xab, 0xb9, 0x88, 0xfb, 0x2c, 0xc9, 0xa0, 0x53, 0xbb, 0xa6, 0xb5, 0xde, 0x42,
	0x13, 0x9c, 0xa1, 0xfd, 0x88, 0x1b, 0x95, 0xb0, 0x59, 0x31, 0x5a, 0x6d, 0x53, 0x8f, 0x3d, 0x27,
	0xcb, 0xc7, 0xf8, 0x84, 0x8e, 0x62, 0x2d, 0x33, 0xb5, 0x74, 0xd1, 0x96, 0xe9, 0x18, 0xec, 0x30,
	0xd3, 0x49, 0x7d, 0x6b, 0x53, 0xcf, 0xbb, 0xb9, 0x89, 0xcf, 0xfb, 0x2e, 0xb7, 0x78, 0xaf, 0x93,
	0xe0, 0x25, 0x42, 0x4b, 0x7e, 0xdb, 0x1b, 0xe7, 0xb4, 0xde, 0xe1, 0x0b, 0x3a, 0x5c, 0x7d, 0x2d,
	0xf5, 0x9d, 0x46, 0xa5, 0x79, 0xd8, 0x7e, 0x8e, 0x6f, 0x7b, 0x85, 0xf9, 0x13, 0xc3, 0x8b, 0x61,
	0x66, 0x0a, 0x5e, 0x89, 0x2c, 0x5d, 0xfd, 0x53, 0xf4, 0x39, 0x50, 0xcb, 0xff, 0x42, 0x7b, 0x43,
	0xbf, 0xfe, 0xdf, 0xad, 0x06, 0x74, 0x9c, 0x19, 0x2b, 0x53, 0xd0, 0x86, 0x5c, 0x96, 0xf0, 0xca,
	0x6d, 0xcf, 0x0a, 0xd3, 0x90, 0xcb, 0xb5, 0x3b, 0x75, 0xd5, 0xf9, 0x56, 0x45, 0xa7, 0x63, 0x99,
	0xe2, 0x7f, 0x6e, 0x55, 0xa7, 0xbe, 0x66, 0x82, 0xee, 0xfc, 0x10, 0xbb, 0x95, 0x8f, 0xaf, 0x9d,
	0x3c, 0x96, 0x09, 0x13, 0x31, 0x96, 0x3a, 0x26, 0x31, 0x88, 0xfc, 0x88, 0xc9, 0x22, 0xc7, 0x5f,
	0x56, 0xfe, 0xa9, 0x47, 0xdf, 0xab, 0xb5, 0x73, 0x4a, 0x7f, 0x54, 0x4f, 0xce, 0x0b, 0x4b, 0x1a,
	0x19, 0x5c, 0xc0, 0x39, 0xea, 0xb7, 0x70, 0xaf, 0x64, 0xfe, 0x2c, 0x39, 0x03, 0x1a, 0x99, 0x81,
	0xe7, 0x0c, 0xfa, 0xad, 0x81, 0xe7, 0xdc, 0x54, 0x4f, 0x8b, 0x42, 0x18, 0xd2, 0xc8, 0x84, 0xa1,
	0x67, 0x85, 0x61, 0xbf, 0x15, 0x86, 0x9e, 0x37, 0xda, 0xc9, 0x87, 0x7d, 0xfc, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xa9, 0xf6, 0x40, 0x2d, 0x9e, 0x04, 0x00, 0x00,
}
