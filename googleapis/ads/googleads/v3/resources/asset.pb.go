// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/asset.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
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

// Asset is a part of an ad which can be shared across multiple ads.
// It can be an image (ImageAsset), a video (YoutubeVideoAsset), etc.
type Asset struct {
	// The resource name of the asset.
	// Asset resource names have the form:
	//
	// `customers/{customer_id}/assets/{asset_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the asset.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Optional name of the asset.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the asset.
	Type enums.AssetTypeEnum_AssetType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.AssetTypeEnum_AssetType" json:"type,omitempty"`
	// The specific type of the asset.
	//
	// Types that are valid to be assigned to AssetData:
	//	*Asset_YoutubeVideoAsset
	//	*Asset_MediaBundleAsset
	//	*Asset_ImageAsset
	//	*Asset_TextAsset
	AssetData            isAsset_AssetData `protobuf_oneof:"asset_data"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cc8f6a72ea26cf4, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Asset) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Asset) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Asset) GetType() enums.AssetTypeEnum_AssetType {
	if m != nil {
		return m.Type
	}
	return enums.AssetTypeEnum_UNSPECIFIED
}

type isAsset_AssetData interface {
	isAsset_AssetData()
}

type Asset_YoutubeVideoAsset struct {
	YoutubeVideoAsset *common.YoutubeVideoAsset `protobuf:"bytes,5,opt,name=youtube_video_asset,json=youtubeVideoAsset,proto3,oneof"`
}

type Asset_MediaBundleAsset struct {
	MediaBundleAsset *common.MediaBundleAsset `protobuf:"bytes,6,opt,name=media_bundle_asset,json=mediaBundleAsset,proto3,oneof"`
}

type Asset_ImageAsset struct {
	ImageAsset *common.ImageAsset `protobuf:"bytes,7,opt,name=image_asset,json=imageAsset,proto3,oneof"`
}

type Asset_TextAsset struct {
	TextAsset *common.TextAsset `protobuf:"bytes,8,opt,name=text_asset,json=textAsset,proto3,oneof"`
}

func (*Asset_YoutubeVideoAsset) isAsset_AssetData() {}

func (*Asset_MediaBundleAsset) isAsset_AssetData() {}

func (*Asset_ImageAsset) isAsset_AssetData() {}

func (*Asset_TextAsset) isAsset_AssetData() {}

func (m *Asset) GetAssetData() isAsset_AssetData {
	if m != nil {
		return m.AssetData
	}
	return nil
}

func (m *Asset) GetYoutubeVideoAsset() *common.YoutubeVideoAsset {
	if x, ok := m.GetAssetData().(*Asset_YoutubeVideoAsset); ok {
		return x.YoutubeVideoAsset
	}
	return nil
}

func (m *Asset) GetMediaBundleAsset() *common.MediaBundleAsset {
	if x, ok := m.GetAssetData().(*Asset_MediaBundleAsset); ok {
		return x.MediaBundleAsset
	}
	return nil
}

func (m *Asset) GetImageAsset() *common.ImageAsset {
	if x, ok := m.GetAssetData().(*Asset_ImageAsset); ok {
		return x.ImageAsset
	}
	return nil
}

func (m *Asset) GetTextAsset() *common.TextAsset {
	if x, ok := m.GetAssetData().(*Asset_TextAsset); ok {
		return x.TextAsset
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Asset) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Asset_YoutubeVideoAsset)(nil),
		(*Asset_MediaBundleAsset)(nil),
		(*Asset_ImageAsset)(nil),
		(*Asset_TextAsset)(nil),
	}
}

func init() {
	proto.RegisterType((*Asset)(nil), "google.ads.googleads.v3.resources.Asset")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/asset.proto", fileDescriptor_8cc8f6a72ea26cf4)
}

var fileDescriptor_8cc8f6a72ea26cf4 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xd6, 0x0e, 0xe6, 0x0d, 0x04, 0xe1, 0xa6, 0x8c, 0x69, 0xea, 0x40, 0x93, 0x0a,
	0x08, 0xa7, 0x2c, 0x68, 0x17, 0xe1, 0x2a, 0x95, 0xa6, 0x6d, 0x95, 0x86, 0xa6, 0x32, 0x55, 0x02,
	0x55, 0x2a, 0x6e, 0x6c, 0x22, 0x4b, 0xb5, 0x1d, 0xc5, 0x4e, 0x59, 0x35, 0xed, 0x21, 0x78, 0x05,
	0x2e, 0x79, 0x14, 0x1e, 0x65, 0x2f, 0xc0, 0x2d, 0x8a, 0xff, 0xa4, 0xd3, 0x50, 0xe9, 0x55, 0x8f,
	0xed, 0xef, 0xfb, 0x7d, 0x27, 0xee, 0x31, 0x78, 0x9b, 0x09, 0x91, 0x4d, 0x49, 0x88, 0xb0, 0x0c,
	0x4d, 0x59, 0x55, 0xb3, 0x28, 0x2c, 0x88, 0x14, 0x65, 0x91, 0x12, 0x19, 0x22, 0x29, 0x89, 0x82,
	0x79, 0x21, 0x94, 0x08, 0xf6, 0x8c, 0x06, 0x22, 0x2c, 0x61, 0x2d, 0x87, 0xb3, 0x08, 0xd6, 0xf2,
	0xed, 0xee, 0x32, 0x62, 0x2a, 0x18, 0x13, 0xdc, 0xe0, 0xc6, 0x6a, 0x9e, 0x13, 0x69, 0xa0, 0xdb,
	0x70, 0x99, 0x83, 0xf0, 0x92, 0xc9, 0x5b, 0x06, 0xab, 0x7f, 0xe6, 0xf4, 0x39, 0xad, 0xdb, 0xb4,
	0x47, 0xbb, 0xf6, 0x48, 0xaf, 0x26, 0xe5, 0xb7, 0xf0, 0x7b, 0x81, 0xf2, 0x9c, 0x14, 0x2e, 0x6a,
	0xe7, 0x96, 0x15, 0x71, 0x2e, 0x14, 0x52, 0x54, 0x70, 0x7b, 0xfa, 0xe2, 0x47, 0x13, 0x34, 0x93,
	0x2a, 0x2d, 0x78, 0x09, 0x1e, 0x3a, 0xf2, 0x98, 0x23, 0x46, 0x5a, 0x5e, 0xdb, 0xeb, 0x6c, 0x0c,
	0xb6, 0xdc, 0xe6, 0x47, 0xc4, 0x48, 0xf0, 0x06, 0xf8, 0x14, 0xb7, 0xfc, 0xb6, 0xd7, 0xd9, 0x3c,
	0x78, 0x6e, 0x3f, 0x02, 0xba, 0x64, 0x78, 0xca, 0xd5, 0xe1, 0xfb, 0x21, 0x9a, 0x96, 0x64, 0xe0,
	0x53, 0x1c, 0x74, 0x41, 0x43, 0x83, 0xd6, 0xb4, 0x7c, 0xe7, 0x1f, 0xf9, 0x27, 0x55, 0x50, 0x9e,
	0x19, 0xbd, 0x56, 0x06, 0x7d, 0xd0, 0xa8, 0x3e, 0xba, 0xd5, 0x68, 0x7b, 0x9d, 0x47, 0x07, 0x87,
	0x70, 0xd9, 0xd5, 0xeb, 0x5b, 0x82, 0xba, 0xef, 0x8b, 0x79, 0x4e, 0x8e, 0x78, 0xc9, 0x16, 0xab,
	0x81, 0x66, 0x04, 0x29, 0x78, 0x3a, 0x17, 0xa5, 0x2a, 0x27, 0x64, 0x3c, 0xa3, 0x98, 0x88, 0xb1,
	0xbe, 0xd4, 0x56, 0x53, 0x37, 0xf3, 0x6e, 0x29, 0xda, 0xfc, 0x65, 0xf0, 0xb3, 0xb1, 0x0e, 0x2b,
	0xa7, 0x26, 0x9f, 0xdc, 0x1b, 0x3c, 0x99, 0xdf, 0xdd, 0x0c, 0xbe, 0x82, 0x80, 0x11, 0x4c, 0xd1,
	0x78, 0x52, 0x72, 0x3c, 0x25, 0x36, 0x63, 0x5d, 0x67, 0x74, 0x57, 0x65, 0x9c, 0x55, 0xce, 0x9e,
	0x36, 0xba, 0x88, 0xc7, 0xec, 0xce, 0x5e, 0x70, 0x06, 0x36, 0x29, 0x43, 0x99, 0x43, 0xdf, 0xd7,
	0xe8, 0xd7, 0xab, 0xd0, 0xa7, 0x95, 0xc5, 0x41, 0x01, 0xad, 0x57, 0x41, 0x1f, 0x00, 0x45, 0x2e,
	0x95, 0xa5, 0x3d, 0xd0, 0xb4, 0x57, 0xab, 0x68, 0x17, 0xe4, 0x52, 0x39, 0xd8, 0x86, 0x72, 0x8b,
	0xf8, 0xe4, 0x26, 0x39, 0x02, 0xbb, 0x0b, 0x83, 0xad, 0x72, 0x2a, 0x2b, 0x63, 0x68, 0xc7, 0x2a,
	0x2d, 0xa5, 0x12, 0x8c, 0x14, 0x32, 0xbc, 0x72, 0xe5, 0xb5, 0x19, 0x70, 0x19, 0x5e, 0xe9, 0xdf,
	0xeb, 0xde, 0x16, 0x00, 0x66, 0xe4, 0x31, 0x52, 0xa8, 0xf7, 0xc7, 0x03, 0xfb, 0xa9, 0x60, 0x70,
	0xe5, 0xc3, 0xeb, 0x01, 0x9d, 0x71, 0x5e, 0x0d, 0xd4, 0xb9, 0xf7, 0xa5, 0x6f, 0x0d, 0x99, 0x98,
	0x22, 0x9e, 0x41, 0x51, 0x64, 0x61, 0x46, 0xb8, 0x1e, 0xb7, 0x70, 0xd1, 0xdb, 0x7f, 0xde, 0xfd,
	0x87, 0xba, 0xfa, 0xe9, 0xaf, 0x1d, 0x27, 0xc9, 0x2f, 0x7f, 0xef, 0xd8, 0x20, 0x13, 0x2c, 0xa1,
	0x29, 0xab, 0x6a, 0x18, 0xc1, 0x81, 0x53, 0xfe, 0x76, 0x9a, 0x51, 0x82, 0xe5, 0xa8, 0xd6, 0x8c,
	0x86, 0xd1, 0xa8, 0xd6, 0xdc, 0xf8, 0xfb, 0xe6, 0x20, 0x8e, 0x13, 0x2c, 0xe3, 0xb8, 0x56, 0xc5,
	0xf1, 0x30, 0x8a, 0xe3, 0x5a, 0x37, 0x59, 0xd7, 0xcd, 0x46, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0xe2, 0xeb, 0xa7, 0xa3, 0x04, 0x00, 0x00,
}