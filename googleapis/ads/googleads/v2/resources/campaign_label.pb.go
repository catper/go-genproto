// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_label.proto

package resources

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

// Represents a relationship between a campaign and a label.
type CampaignLabel struct {
	// Immutable. Name of the resource.
	// Campaign label resource names have the form:
	// `customers/{customer_id}/campaignLabels/{campaign_id}~{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The campaign to which the label is attached.
	Campaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Immutable. The label assigned to the campaign.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CampaignLabel) Reset()         { *m = CampaignLabel{} }
func (m *CampaignLabel) String() string { return proto.CompactTextString(m) }
func (*CampaignLabel) ProtoMessage()    {}
func (*CampaignLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a6c77ba70ba8330, []int{0}
}

func (m *CampaignLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignLabel.Unmarshal(m, b)
}
func (m *CampaignLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignLabel.Marshal(b, m, deterministic)
}
func (m *CampaignLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignLabel.Merge(m, src)
}
func (m *CampaignLabel) XXX_Size() int {
	return xxx_messageInfo_CampaignLabel.Size(m)
}
func (m *CampaignLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignLabel.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignLabel proto.InternalMessageInfo

func (m *CampaignLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignLabel) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*CampaignLabel)(nil), "google.ads.googleads.v2.resources.CampaignLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_label.proto", fileDescriptor_1a6c77ba70ba8330)
}

var fileDescriptor_1a6c77ba70ba8330 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x8a, 0xd4, 0x30,
	0x18, 0xa7, 0x1d, 0x56, 0x34, 0xba, 0x97, 0x9e, 0xc6, 0x65, 0x59, 0x67, 0x94, 0x5d, 0xc6, 0x4b,
	0x02, 0x55, 0x3c, 0xc4, 0x53, 0xea, 0x61, 0x41, 0x44, 0x96, 0x59, 0xe8, 0x61, 0x29, 0x0e, 0x69,
	0x9b, 0x8d, 0x85, 0x36, 0x29, 0x49, 0x3b, 0x1e, 0x64, 0xef, 0x3e, 0x87, 0x47, 0x1f, 0xc5, 0xa7,
	0xf0, 0xbc, 0x8f, 0xe0, 0x45, 0x69, 0xf3, 0x67, 0x66, 0x0e, 0x3a, 0x73, 0xfb, 0x85, 0xdf, 0xbf,
	0xef, 0x0b, 0x1f, 0x78, 0xc3, 0xa5, 0xe4, 0x35, 0x43, 0xb4, 0xd4, 0xc8, 0xc0, 0x01, 0xad, 0x63,
	0xa4, 0x98, 0x96, 0xbd, 0x2a, 0x98, 0x46, 0x05, 0x6d, 0x5a, 0x5a, 0x71, 0xb1, 0xaa, 0x69, 0xce,
	0x6a, 0xd8, 0x2a, 0xd9, 0xc9, 0x68, 0x6e, 0xc4, 0x90, 0x96, 0x1a, 0x7a, 0x1f, 0x5c, 0xc7, 0xd0,
	0xfb, 0x4e, 0x9e, 0xb9, 0xe8, 0xb6, 0x42, 0xb7, 0x15, 0xab, 0xcb, 0x55, 0xce, 0x3e, 0xd3, 0x75,
	0x25, 0x95, 0xc9, 0x38, 0x79, 0xba, 0x25, 0x70, 0x36, 0x4b, 0x9d, 0x59, 0x6a, 0x7c, 0xe5, 0xfd,
	0x2d, 0xfa, 0xa2, 0x68, 0xdb, 0x32, 0xa5, 0x2d, 0x7f, 0xba, 0x65, 0xa5, 0x42, 0xc8, 0x8e, 0x76,
	0x95, 0x14, 0x96, 0x7d, 0xfe, 0x6d, 0x02, 0x8e, 0xdf, 0xd9, 0xa9, 0x3f, 0x0c, 0x43, 0x47, 0xd7,
	0xe0, 0xd8, 0x35, 0xac, 0x04, 0x6d, 0xd8, 0x34, 0x98, 0x05, 0x8b, 0x47, 0x09, 0xfc, 0x45, 0x8e,
	0x7e, 0x93, 0x05, 0xb8, 0xd8, 0xac, 0x60, 0x51, 0x5b, 0x69, 0x58, 0xc8, 0x06, 0xed, 0xc4, 0x2c,
	0x9f, 0xb8, 0x90, 0x8f, 0xb4, 0x61, 0x51, 0x01, 0x1e, 0xba, 0xbf, 0x99, 0x86, 0xb3, 0x60, 0xf1,
	0x38, 0x3e, 0xb5, 0x76, 0xe8, 0xe6, 0x86, 0xd7, 0x9d, 0xaa, 0x04, 0x4f, 0x69, 0xdd, 0xb3, 0xe4,
	0xe5, 0xd8, 0xf6, 0x02, 0xcc, 0xf7, 0xb6, 0x2d, 0x7d, 0x70, 0x74, 0x03, 0x8e, 0xc6, 0x7f, 0x9f,
	0x4e, 0x0e, 0x68, 0xb8, 0x18, 0x1b, 0x66, 0xe0, 0xec, 0x9f, 0x0d, 0x66, 0x0f, 0x13, 0x89, 0xe9,
	0x3d, 0xf9, 0x74, 0xe8, 0xee, 0xd1, 0xeb, 0xa2, 0xd7, 0x9d, 0x6c, 0x98, 0xd2, 0xe8, 0xab, 0x83,
	0x77, 0xfe, 0x38, 0x46, 0xcd, 0xc0, 0xec, 0x1c, 0xcb, 0x5d, 0xf2, 0x27, 0x00, 0xe7, 0x85, 0x6c,
	0xe0, 0xde, 0x73, 0x49, 0xa2, 0x9d, 0xba, 0xab, 0x61, 0xbd, 0xab, 0xe0, 0xe6, 0xbd, 0x35, 0x72,
	0x59, 0x53, 0xc1, 0xa1, 0x54, 0x1c, 0x71, 0x26, 0xc6, 0xe5, 0xd1, 0x66, 0xdc, 0xff, 0x9c, 0xef,
	0x5b, 0x8f, 0xbe, 0x87, 0x93, 0x4b, 0x42, 0x7e, 0x84, 0xf3, 0x4b, 0x13, 0x49, 0x4a, 0x0d, 0x0d,
	0x1c, 0x50, 0x1a, 0xc3, 0xa5, 0x53, 0xfe, 0x74, 0x9a, 0x8c, 0x94, 0x3a, 0xf3, 0x9a, 0x2c, 0x8d,
	0x33, 0xaf, 0xb9, 0x0f, 0xcf, 0x0d, 0x81, 0x31, 0x29, 0x35, 0xc6, 0x5e, 0x85, 0x71, 0x1a, 0x63,
	0xec, 0x75, 0xf9, 0x83, 0x71, 0xd8, 0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x34, 0x23, 0x88,
	0xa5, 0x6a, 0x03, 0x00, 0x00,
}
