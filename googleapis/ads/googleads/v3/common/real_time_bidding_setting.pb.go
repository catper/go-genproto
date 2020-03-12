// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/common/real_time_bidding_setting.proto

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

// Settings for Real-Time Bidding, a feature only available for campaigns
// targeting the Ad Exchange network.
type RealTimeBiddingSetting struct {
	// Whether the campaign is opted in to real-time bidding.
	OptIn                *wrappers.BoolValue `protobuf:"bytes,1,opt,name=opt_in,json=optIn,proto3" json:"opt_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RealTimeBiddingSetting) Reset()         { *m = RealTimeBiddingSetting{} }
func (m *RealTimeBiddingSetting) String() string { return proto.CompactTextString(m) }
func (*RealTimeBiddingSetting) ProtoMessage()    {}
func (*RealTimeBiddingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_575472eed6650645, []int{0}
}

func (m *RealTimeBiddingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealTimeBiddingSetting.Unmarshal(m, b)
}
func (m *RealTimeBiddingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealTimeBiddingSetting.Marshal(b, m, deterministic)
}
func (m *RealTimeBiddingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealTimeBiddingSetting.Merge(m, src)
}
func (m *RealTimeBiddingSetting) XXX_Size() int {
	return xxx_messageInfo_RealTimeBiddingSetting.Size(m)
}
func (m *RealTimeBiddingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_RealTimeBiddingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_RealTimeBiddingSetting proto.InternalMessageInfo

func (m *RealTimeBiddingSetting) GetOptIn() *wrappers.BoolValue {
	if m != nil {
		return m.OptIn
	}
	return nil
}

func init() {
	proto.RegisterType((*RealTimeBiddingSetting)(nil), "google.ads.googleads.v3.common.RealTimeBiddingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/common/real_time_bidding_setting.proto", fileDescriptor_575472eed6650645)
}

var fileDescriptor_575472eed6650645 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xec, 0x30,
	0x14, 0x86, 0xe9, 0x5c, 0xee, 0x2c, 0xea, 0x6e, 0x16, 0x22, 0xa3, 0x0c, 0x32, 0x2b, 0x57, 0x09,
	0x9a, 0x5d, 0x04, 0xa1, 0x55, 0x18, 0xc4, 0xcd, 0x30, 0x4a, 0x17, 0x52, 0x28, 0x99, 0x49, 0x0c,
	0x81, 0x34, 0x27, 0x34, 0x99, 0xf1, 0x7d, 0x5c, 0xfa, 0x28, 0x3e, 0x8a, 0x0f, 0xe0, 0x5a, 0xda,
	0xd3, 0x76, 0xa5, 0xae, 0xfa, 0xd3, 0x7c, 0xff, 0xf9, 0x0e, 0x27, 0xbd, 0xd1, 0x00, 0xda, 0x2a,
	0x2a, 0x64, 0xa0, 0x18, 0xdb, 0x74, 0x60, 0x74, 0x07, 0x75, 0x0d, 0x8e, 0x36, 0x4a, 0xd8, 0x2a,
	0x9a, 0x5a, 0x55, 0x5b, 0x23, 0xa5, 0x71, 0xba, 0x0a, 0x2a, 0x46, 0xe3, 0x34, 0xf1, 0x0d, 0x44,
	0x98, 0x2d, 0xb0, 0x44, 0x84, 0x0c, 0x64, 0xec, 0x93, 0x03, 0x23, 0xd8, 0x9f, 0xf7, 0xef, 0xb4,
	0xa3, 0xb7, 0xfb, 0x17, 0xfa, 0xda, 0x08, 0xef, 0x55, 0x13, 0xb0, 0x3f, 0x3f, 0x1b, 0xfc, 0xde,
	0x50, 0xe1, 0x1c, 0x44, 0x11, 0x0d, 0xb8, 0xfe, 0x75, 0xf9, 0x90, 0x1e, 0x6f, 0x94, 0xb0, 0x4f,
	0xa6, 0x56, 0x39, 0xea, 0x1f, 0xd1, 0x3e, 0xbb, 0x4c, 0xa7, 0xe0, 0x63, 0x65, 0xdc, 0x49, 0x72,
	0x9e, 0x5c, 0x1c, 0x5d, 0xcd, 0x7b, 0x3b, 0x19, 0x44, 0x24, 0x07, 0xb0, 0x85, 0xb0, 0x7b, 0xb5,
	0xf9, 0x0f, 0x3e, 0xde, 0xbb, 0xfc, 0x2b, 0x49, 0x97, 0x3b, 0xa8, 0xc9, 0xdf, 0x1b, 0xe7, 0xa7,
	0x3f, 0x1b, 0xd7, 0xed, 0xdc, 0x75, 0xf2, 0x7c, 0xd7, 0xd7, 0x35, 0x58, 0xe1, 0x34, 0x81, 0x46,
	0x53, 0xad, 0x5c, 0x67, 0x1d, 0x0e, 0xe8, 0x4d, 0xf8, 0xed, 0x9e, 0xd7, 0xf8, 0x79, 0x9b, 0xfc,
	0x5b, 0x65, 0xd9, 0xfb, 0x64, 0xb1, 0xc2, 0x61, 0x99, 0x0c, 0x04, 0x63, 0x9b, 0x0a, 0x46, 0x6e,
	0x3b, 0xec, 0x63, 0x00, 0xca, 0x4c, 0x86, 0x72, 0x04, 0xca, 0x82, 0x95, 0x08, 0x7c, 0x4e, 0x96,
	0xf8, 0x97, 0xf3, 0x4c, 0x06, 0xce, 0x47, 0x84, 0xf3, 0x82, 0x71, 0x8e, 0xd0, 0x76, 0xda, 0x6d,
	0xc7, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x5c, 0xb0, 0x50, 0xec, 0x01, 0x00, 0x00,
}
