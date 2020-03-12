// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/text_label.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
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

// A type of label displaying text on a colored background.
type TextLabel struct {
	// Background color of the label in RGB format. This string must match the
	// regular expression '^\#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$'.
	// Note: The background color may not be visible for manager accounts.
	BackgroundColor *wrappers.StringValue `protobuf:"bytes,1,opt,name=background_color,json=backgroundColor,proto3" json:"background_color,omitempty"`
	// A short description of the label. The length must be no more than 200
	// characters.
	Description          *wrappers.StringValue `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TextLabel) Reset()         { *m = TextLabel{} }
func (m *TextLabel) String() string { return proto.CompactTextString(m) }
func (*TextLabel) ProtoMessage()    {}
func (*TextLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f97c60eafd5c770, []int{0}
}

func (m *TextLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextLabel.Unmarshal(m, b)
}
func (m *TextLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextLabel.Marshal(b, m, deterministic)
}
func (m *TextLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextLabel.Merge(m, src)
}
func (m *TextLabel) XXX_Size() int {
	return xxx_messageInfo_TextLabel.Size(m)
}
func (m *TextLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_TextLabel.DiscardUnknown(m)
}

var xxx_messageInfo_TextLabel proto.InternalMessageInfo

func (m *TextLabel) GetBackgroundColor() *wrappers.StringValue {
	if m != nil {
		return m.BackgroundColor
	}
	return nil
}

func (m *TextLabel) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func init() {
	proto.RegisterType((*TextLabel)(nil), "google.ads.googleads.v1.common.TextLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/text_label.proto", fileDescriptor_0f97c60eafd5c770)
}

var fileDescriptor_0f97c60eafd5c770 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xfb, 0x30,
	0x1c, 0xc7, 0x69, 0xff, 0xf0, 0x07, 0x3b, 0x50, 0xd9, 0x69, 0x8c, 0x31, 0x64, 0x27, 0x4f, 0x09,
	0xd5, 0x5b, 0x04, 0xa1, 0x9b, 0xb0, 0x8b, 0x87, 0xa1, 0xd2, 0x83, 0x14, 0x46, 0xda, 0xc4, 0x50,
	0xcc, 0xf2, 0x0b, 0x49, 0x36, 0xf7, 0x20, 0x3e, 0x81, 0x47, 0x1f, 0xc5, 0x37, 0xd1, 0xa7, 0x90,
	0x34, 0x6d, 0xf5, 0xa2, 0x78, 0xea, 0x97, 0xe6, 0xf3, 0xf9, 0x7d, 0x93, 0x5f, 0x82, 0x05, 0x80,
	0x90, 0x1c, 0x53, 0x66, 0xdb, 0xe8, 0xd3, 0x2e, 0xc5, 0x15, 0x6c, 0x36, 0xa0, 0xb0, 0xe3, 0x7b,
	0xb7, 0x96, 0xb4, 0xe4, 0x12, 0x69, 0x03, 0x0e, 0x86, 0xd3, 0x40, 0x21, 0xca, 0x2c, 0xea, 0x05,
	0xb4, 0x4b, 0x51, 0x10, 0xc6, 0xed, 0x39, 0x6e, 0xe8, 0x72, 0xfb, 0x80, 0x9f, 0x0c, 0xd5, 0x9a,
	0x1b, 0x1b, 0xfc, 0xf1, 0xa4, 0x2b, 0xd4, 0x35, 0xa6, 0x4a, 0x81, 0xa3, 0xae, 0x06, 0xd5, 0x9e,
	0xce, 0x9e, 0xa3, 0xe4, 0xe0, 0x8e, 0xef, 0xdd, 0xb5, 0x6f, 0x1c, 0x2e, 0x93, 0xe3, 0x92, 0x56,
	0x8f, 0xc2, 0xc0, 0x56, 0xb1, 0x75, 0x05, 0x12, 0xcc, 0x28, 0x3a, 0x89, 0x4e, 0x07, 0x67, 0x93,
	0xb6, 0x1b, 0x75, 0x35, 0xe8, 0xd6, 0x99, 0x5a, 0x89, 0x9c, 0xca, 0x2d, 0xbf, 0x39, 0xfa, 0xb2,
	0x16, 0x5e, 0x1a, 0x5e, 0x26, 0x03, 0xc6, 0x6d, 0x65, 0x6a, 0xed, 0xcb, 0x46, 0xf1, 0x1f, 0x66,
	0x7c, 0x17, 0xe6, 0xef, 0x51, 0x32, 0xab, 0x60, 0x83, 0x7e, 0x7f, 0xfb, 0xfc, 0xb0, 0xbf, 0xfa,
	0xca, 0x8f, 0x5c, 0x45, 0xf7, 0x57, 0xad, 0x21, 0x40, 0x52, 0x25, 0x10, 0x18, 0x81, 0x05, 0x57,
	0x4d, 0x61, 0xb7, 0x6e, 0x5d, 0xdb, 0x9f, 0xb6, 0x7f, 0x11, 0x3e, 0x2f, 0xf1, 0xbf, 0x65, 0x96,
	0xbd, 0xc6, 0xd3, 0x65, 0x18, 0x96, 0x31, 0x8b, 0x42, 0xf4, 0x29, 0x4f, 0xd1, 0xa2, 0xc1, 0xde,
	0x3a, 0xa0, 0xc8, 0x98, 0x2d, 0x7a, 0xa0, 0xc8, 0xd3, 0x22, 0x00, 0x1f, 0xf1, 0x2c, 0xfc, 0x25,
	0x24, 0x63, 0x96, 0x90, 0x1e, 0x21, 0x24, 0x4f, 0x09, 0x09, 0x50, 0xf9, 0xbf, 0xb9, 0xdd, 0xf9,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x01, 0x9e, 0x77, 0x1a, 0x02, 0x00, 0x00,
}
