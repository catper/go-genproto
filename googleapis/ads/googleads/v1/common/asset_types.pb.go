// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/asset_types.proto

package common

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

// A YouTube asset.
type YoutubeVideoAsset struct {
	// YouTube video id. This is the 11 character string value used in the
	// YouTube video URL.
	YoutubeVideoId       *wrappers.StringValue `protobuf:"bytes,1,opt,name=youtube_video_id,json=youtubeVideoId,proto3" json:"youtube_video_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *YoutubeVideoAsset) Reset()         { *m = YoutubeVideoAsset{} }
func (m *YoutubeVideoAsset) String() string { return proto.CompactTextString(m) }
func (*YoutubeVideoAsset) ProtoMessage()    {}
func (*YoutubeVideoAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84b8dff1cc51362, []int{0}
}

func (m *YoutubeVideoAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YoutubeVideoAsset.Unmarshal(m, b)
}
func (m *YoutubeVideoAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YoutubeVideoAsset.Marshal(b, m, deterministic)
}
func (m *YoutubeVideoAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YoutubeVideoAsset.Merge(m, src)
}
func (m *YoutubeVideoAsset) XXX_Size() int {
	return xxx_messageInfo_YoutubeVideoAsset.Size(m)
}
func (m *YoutubeVideoAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_YoutubeVideoAsset.DiscardUnknown(m)
}

var xxx_messageInfo_YoutubeVideoAsset proto.InternalMessageInfo

func (m *YoutubeVideoAsset) GetYoutubeVideoId() *wrappers.StringValue {
	if m != nil {
		return m.YoutubeVideoId
	}
	return nil
}

// A MediaBundle asset.
type MediaBundleAsset struct {
	// Media bundle (ZIP file) asset data. The format of the uploaded ZIP file
	// depends on the ad field where it will be used. For more information on the
	// format, see the documentation of the ad field where you plan on using the
	// MediaBundleAsset. This field is mutate only.
	Data                 *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaBundleAsset) Reset()         { *m = MediaBundleAsset{} }
func (m *MediaBundleAsset) String() string { return proto.CompactTextString(m) }
func (*MediaBundleAsset) ProtoMessage()    {}
func (*MediaBundleAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84b8dff1cc51362, []int{1}
}

func (m *MediaBundleAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaBundleAsset.Unmarshal(m, b)
}
func (m *MediaBundleAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaBundleAsset.Marshal(b, m, deterministic)
}
func (m *MediaBundleAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaBundleAsset.Merge(m, src)
}
func (m *MediaBundleAsset) XXX_Size() int {
	return xxx_messageInfo_MediaBundleAsset.Size(m)
}
func (m *MediaBundleAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaBundleAsset.DiscardUnknown(m)
}

var xxx_messageInfo_MediaBundleAsset proto.InternalMessageInfo

func (m *MediaBundleAsset) GetData() *wrappers.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

// An Image asset.
type ImageAsset struct {
	// The raw bytes data of an image. This field is mutate only.
	Data *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// File size of the image asset in bytes.
	FileSize *wrappers.Int64Value `protobuf:"bytes,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// MIME type of the image asset.
	MimeType enums.MimeTypeEnum_MimeType `protobuf:"varint,3,opt,name=mime_type,json=mimeType,proto3,enum=google.ads.googleads.v1.enums.MimeTypeEnum_MimeType" json:"mime_type,omitempty"`
	// Metadata for this image at its original size.
	FullSize             *ImageDimension `protobuf:"bytes,4,opt,name=full_size,json=fullSize,proto3" json:"full_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ImageAsset) Reset()         { *m = ImageAsset{} }
func (m *ImageAsset) String() string { return proto.CompactTextString(m) }
func (*ImageAsset) ProtoMessage()    {}
func (*ImageAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84b8dff1cc51362, []int{2}
}

func (m *ImageAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageAsset.Unmarshal(m, b)
}
func (m *ImageAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageAsset.Marshal(b, m, deterministic)
}
func (m *ImageAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageAsset.Merge(m, src)
}
func (m *ImageAsset) XXX_Size() int {
	return xxx_messageInfo_ImageAsset.Size(m)
}
func (m *ImageAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageAsset.DiscardUnknown(m)
}

var xxx_messageInfo_ImageAsset proto.InternalMessageInfo

func (m *ImageAsset) GetData() *wrappers.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ImageAsset) GetFileSize() *wrappers.Int64Value {
	if m != nil {
		return m.FileSize
	}
	return nil
}

func (m *ImageAsset) GetMimeType() enums.MimeTypeEnum_MimeType {
	if m != nil {
		return m.MimeType
	}
	return enums.MimeTypeEnum_UNSPECIFIED
}

func (m *ImageAsset) GetFullSize() *ImageDimension {
	if m != nil {
		return m.FullSize
	}
	return nil
}

// Metadata for an image at a certain size, either original or resized.
type ImageDimension struct {
	// Height of the image.
	HeightPixels *wrappers.Int64Value `protobuf:"bytes,1,opt,name=height_pixels,json=heightPixels,proto3" json:"height_pixels,omitempty"`
	// Width of the image.
	WidthPixels *wrappers.Int64Value `protobuf:"bytes,2,opt,name=width_pixels,json=widthPixels,proto3" json:"width_pixels,omitempty"`
	// A URL that returns the image with this height and width.
	Url                  *wrappers.StringValue `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ImageDimension) Reset()         { *m = ImageDimension{} }
func (m *ImageDimension) String() string { return proto.CompactTextString(m) }
func (*ImageDimension) ProtoMessage()    {}
func (*ImageDimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84b8dff1cc51362, []int{3}
}

func (m *ImageDimension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageDimension.Unmarshal(m, b)
}
func (m *ImageDimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageDimension.Marshal(b, m, deterministic)
}
func (m *ImageDimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageDimension.Merge(m, src)
}
func (m *ImageDimension) XXX_Size() int {
	return xxx_messageInfo_ImageDimension.Size(m)
}
func (m *ImageDimension) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageDimension.DiscardUnknown(m)
}

var xxx_messageInfo_ImageDimension proto.InternalMessageInfo

func (m *ImageDimension) GetHeightPixels() *wrappers.Int64Value {
	if m != nil {
		return m.HeightPixels
	}
	return nil
}

func (m *ImageDimension) GetWidthPixels() *wrappers.Int64Value {
	if m != nil {
		return m.WidthPixels
	}
	return nil
}

func (m *ImageDimension) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

// A Text asset.
type TextAsset struct {
	// Text content of the text asset.
	Text                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TextAsset) Reset()         { *m = TextAsset{} }
func (m *TextAsset) String() string { return proto.CompactTextString(m) }
func (*TextAsset) ProtoMessage()    {}
func (*TextAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84b8dff1cc51362, []int{4}
}

func (m *TextAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextAsset.Unmarshal(m, b)
}
func (m *TextAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextAsset.Marshal(b, m, deterministic)
}
func (m *TextAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextAsset.Merge(m, src)
}
func (m *TextAsset) XXX_Size() int {
	return xxx_messageInfo_TextAsset.Size(m)
}
func (m *TextAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_TextAsset.DiscardUnknown(m)
}

var xxx_messageInfo_TextAsset proto.InternalMessageInfo

func (m *TextAsset) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func init() {
	proto.RegisterType((*YoutubeVideoAsset)(nil), "google.ads.googleads.v1.common.YoutubeVideoAsset")
	proto.RegisterType((*MediaBundleAsset)(nil), "google.ads.googleads.v1.common.MediaBundleAsset")
	proto.RegisterType((*ImageAsset)(nil), "google.ads.googleads.v1.common.ImageAsset")
	proto.RegisterType((*ImageDimension)(nil), "google.ads.googleads.v1.common.ImageDimension")
	proto.RegisterType((*TextAsset)(nil), "google.ads.googleads.v1.common.TextAsset")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/asset_types.proto", fileDescriptor_f84b8dff1cc51362)
}

var fileDescriptor_f84b8dff1cc51362 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0x71, 0x12, 0x46, 0xa2, 0x76, 0x59, 0xe7, 0x53, 0xc8, 0x4a, 0x29, 0x3e, 0xf5, 0x32,
	0xb9, 0xe9, 0xca, 0x18, 0x1e, 0x1b, 0x73, 0xda, 0xad, 0x84, 0x51, 0xc8, 0xd2, 0x12, 0xd8, 0x16,
	0x08, 0x4a, 0xf5, 0xe2, 0x08, 0x2c, 0xc9, 0x58, 0x72, 0x9a, 0xf4, 0x8b, 0xec, 0xbe, 0xe3, 0x3e,
	0xc6, 0x8e, 0xfb, 0x28, 0xfd, 0x14, 0x43, 0x92, 0x13, 0x3a, 0x46, 0xd6, 0xb2, 0x93, 0x9f, 0xa5,
	0xff, 0xff, 0xa7, 0xa7, 0xf7, 0x9e, 0xd0, 0x61, 0x22, 0x65, 0x92, 0x42, 0x48, 0xa8, 0x0a, 0x5d,
	0x68, 0xa2, 0x79, 0x27, 0xbc, 0x92, 0x9c, 0x4b, 0x11, 0x12, 0xa5, 0x40, 0x8f, 0xf5, 0x32, 0x03,
	0x85, 0xb3, 0x5c, 0x6a, 0xe9, 0xef, 0x39, 0x19, 0x26, 0x54, 0xe1, 0xb5, 0x03, 0xcf, 0x3b, 0xd8,
	0x39, 0xda, 0xcf, 0x37, 0x11, 0x41, 0x14, 0x5c, 0x85, 0x9c, 0x71, 0xb0, 0x3c, 0x87, 0x6b, 0x97,
	0xb8, 0xd0, 0xfe, 0x4d, 0x8a, 0x69, 0x78, 0x9d, 0x93, 0x2c, 0x83, 0xbc, 0x3c, 0xae, 0xbd, 0xbb,
	0xc2, 0x65, 0x2c, 0x24, 0x42, 0x48, 0x4d, 0x34, 0x93, 0xa2, 0xdc, 0x0d, 0xbe, 0xa2, 0xa7, 0x9f,
	0x65, 0xa1, 0x8b, 0x09, 0x0c, 0x19, 0x05, 0x19, 0x9b, 0x6c, 0xfd, 0x0f, 0x68, 0x67, 0xe9, 0x16,
	0xc7, 0x73, 0xb3, 0x3a, 0x66, 0xb4, 0xe5, 0xed, 0x7b, 0x07, 0x5b, 0x47, 0xbb, 0x65, 0xc6, 0x78,
	0x75, 0x1a, 0xbe, 0xd0, 0x39, 0x13, 0xc9, 0x90, 0xa4, 0x05, 0x0c, 0x9a, 0xcb, 0x3b, 0xa8, 0x1e,
	0x0d, 0x4e, 0xd0, 0xce, 0x39, 0x50, 0x46, 0xba, 0x85, 0xa0, 0x29, 0x38, 0x76, 0x88, 0x6a, 0x94,
	0x68, 0x52, 0xf2, 0x9e, 0xfd, 0xc5, 0xeb, 0x2e, 0x35, 0x28, 0x87, 0xb3, 0xc2, 0xe0, 0x5b, 0x05,
	0xa1, 0x1e, 0x27, 0xc9, 0x7f, 0xfa, 0xfd, 0x57, 0xa8, 0x31, 0x65, 0x29, 0x8c, 0x15, 0xbb, 0x81,
	0x56, 0x65, 0x83, 0xab, 0x27, 0xf4, 0xcb, 0x63, 0xe7, 0xaa, 0x1b, 0xf5, 0x05, 0xbb, 0x01, 0xff,
	0x13, 0x6a, 0xac, 0x8b, 0xdd, 0xaa, 0xee, 0x7b, 0x07, 0xcd, 0xa3, 0x63, 0xbc, 0xa9, 0x79, 0xb6,
	0x39, 0xf8, 0x9c, 0x71, 0xb8, 0x5c, 0x66, 0xf0, 0x5e, 0x14, 0x7c, 0xfd, 0x33, 0xa8, 0xf3, 0x32,
	0xf2, 0x3f, 0xa2, 0xc6, 0xb4, 0x48, 0x53, 0x97, 0x4c, 0xcd, 0x26, 0x83, 0xf1, 0xbf, 0xe7, 0x01,
	0xdb, 0xcb, 0x9f, 0x32, 0x0e, 0x42, 0x31, 0x29, 0x06, 0x75, 0x03, 0x30, 0xf9, 0x05, 0x3f, 0x3d,
	0xd4, 0xfc, 0x73, 0xd3, 0x7f, 0x87, 0x1e, 0xcf, 0x80, 0x25, 0x33, 0x3d, 0xce, 0xd8, 0x02, 0x52,
	0xb5, 0xb1, 0x4c, 0x77, 0x2e, 0xbc, 0xed, 0x1c, 0x7d, 0x6b, 0xf0, 0xdf, 0xa2, 0xed, 0x6b, 0x46,
	0xf5, 0x6c, 0x05, 0x78, 0x40, 0xc5, 0xb6, 0xac, 0xa1, 0xf4, 0x63, 0x54, 0x2d, 0xf2, 0xd4, 0x96,
	0xeb, 0xbe, 0x71, 0x31, 0xc2, 0xe0, 0x0d, 0x6a, 0x5c, 0xc2, 0x42, 0xbb, 0xe6, 0x1e, 0xa2, 0x9a,
	0x86, 0x85, 0x7e, 0xd0, 0xb0, 0x59, 0x65, 0xf7, 0xd6, 0x43, 0xc1, 0x95, 0xe4, 0xf7, 0xd4, 0xb0,
	0xfb, 0xc4, 0xf2, 0x4d, 0x0b, 0x54, 0xdf, 0xc0, 0xfa, 0xde, 0x97, 0xd3, 0xd2, 0x92, 0xc8, 0x94,
	0x88, 0x04, 0xcb, 0x3c, 0x09, 0x13, 0x10, 0xf6, 0xa8, 0xd5, 0xb3, 0xcb, 0x98, 0xda, 0xf4, 0xae,
	0x5f, 0xbb, 0xcf, 0xf7, 0x4a, 0xf5, 0x2c, 0x8e, 0x7f, 0x54, 0xf6, 0xce, 0x1c, 0x2c, 0xa6, 0x0a,
	0xbb, 0xd0, 0x44, 0xc3, 0x0e, 0x3e, 0xb1, 0xb2, 0x5f, 0x2b, 0xc1, 0x28, 0xa6, 0x6a, 0xb4, 0x16,
	0x8c, 0x86, 0x9d, 0x91, 0x13, 0xdc, 0x56, 0x02, 0xb7, 0x1a, 0x45, 0x31, 0x55, 0x51, 0xb4, 0x96,
	0x44, 0xd1, 0xb0, 0x13, 0x45, 0x4e, 0x34, 0x79, 0x64, 0xb3, 0x7b, 0xf1, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0x95, 0x59, 0x49, 0x74, 0x04, 0x00, 0x00,
}
