// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_label.proto

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
	return fileDescriptor_a1b047812890bf04, []int{0}
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
	proto.RegisterType((*CampaignLabel)(nil), "google.ads.googleads.v1.resources.CampaignLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_label.proto", fileDescriptor_a1b047812890bf04)
}

var fileDescriptor_a1b047812890bf04 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbf, 0x6e, 0xd4, 0x30,
	0x18, 0x57, 0x72, 0x2a, 0x02, 0x43, 0x97, 0x4c, 0x47, 0x55, 0x95, 0x3b, 0x50, 0xab, 0x63, 0xb1,
	0x15, 0x40, 0x0c, 0x66, 0x72, 0x18, 0x2a, 0x21, 0x84, 0xaa, 0xab, 0x94, 0xa1, 0x8a, 0x38, 0x39,
	0x89, 0x6b, 0x22, 0x25, 0x76, 0x64, 0x27, 0xc7, 0x80, 0xba, 0xf3, 0x1c, 0x8c, 0x3c, 0x0a, 0x4f,
	0xc1, 0xdc, 0x47, 0x60, 0x01, 0x25, 0xfe, 0xd3, 0xbb, 0x81, 0xde, 0x6d, 0x3f, 0xeb, 0xf7, 0xef,
	0xfb, 0xac, 0x0f, 0xbc, 0xe5, 0x52, 0xf2, 0x9a, 0x21, 0x5a, 0x6a, 0x64, 0xe0, 0x80, 0xd6, 0x31,
	0x52, 0x4c, 0xcb, 0x5e, 0x15, 0x4c, 0xa3, 0x82, 0x36, 0x2d, 0xad, 0xb8, 0x58, 0xd5, 0x34, 0x67,
	0x35, 0x6c, 0x95, 0xec, 0x64, 0x34, 0x37, 0x62, 0x48, 0x4b, 0x0d, 0xbd, 0x0f, 0xae, 0x63, 0xe8,
	0x7d, 0x47, 0xcf, 0x5c, 0x74, 0x5b, 0xa1, 0xeb, 0x8a, 0xd5, 0xe5, 0x2a, 0x67, 0x5f, 0xe8, 0xba,
	0x92, 0xca, 0x64, 0x1c, 0x3d, 0xdd, 0x10, 0x38, 0x9b, 0xa5, 0x4e, 0x2c, 0x35, 0xbe, 0xf2, 0xfe,
	0x1a, 0x7d, 0x55, 0xb4, 0x6d, 0x99, 0xd2, 0x96, 0x3f, 0xde, 0xb0, 0x52, 0x21, 0x64, 0x47, 0xbb,
	0x4a, 0x0a, 0xcb, 0x3e, 0xff, 0x3e, 0x01, 0x87, 0xef, 0xed, 0xd4, 0x1f, 0x87, 0xa1, 0xa3, 0x4b,
	0x70, 0xe8, 0x1a, 0x56, 0x82, 0x36, 0x6c, 0x1a, 0xcc, 0x82, 0xc5, 0xa3, 0x04, 0xfe, 0x26, 0x07,
	0x7f, 0xc8, 0x02, 0x9c, 0xdd, 0xad, 0x60, 0x51, 0x5b, 0x69, 0x58, 0xc8, 0x06, 0x6d, 0xc5, 0x2c,
	0x9f, 0xb8, 0x90, 0x4f, 0xb4, 0x61, 0x51, 0x01, 0x1e, 0xba, 0xbf, 0x99, 0x86, 0xb3, 0x60, 0xf1,
	0xf8, 0xd5, 0xb1, 0xb5, 0x43, 0x37, 0x37, 0xbc, 0xec, 0x54, 0x25, 0x78, 0x4a, 0xeb, 0x9e, 0x25,
	0x2f, 0xc7, 0xb6, 0x17, 0x60, 0xbe, 0xb3, 0x6d, 0xe9, 0x83, 0xa3, 0x2b, 0x70, 0x30, 0xfe, 0xfb,
	0x74, 0xb2, 0x47, 0xc3, 0xd9, 0xd8, 0x30, 0x03, 0x27, 0xff, 0x6d, 0x30, 0x7b, 0x98, 0x48, 0x4c,
	0x6f, 0xc9, 0xe7, 0x7d, 0x77, 0x8f, 0xde, 0x14, 0xbd, 0xee, 0x64, 0xc3, 0x94, 0x46, 0xdf, 0x1c,
	0xbc, 0xf1, 0xc7, 0x31, 0x6a, 0x06, 0x66, 0xeb, 0x58, 0x6e, 0x92, 0xbf, 0x01, 0x38, 0x2d, 0x64,
	0x03, 0x77, 0x9e, 0x4b, 0x12, 0x6d, 0xd5, 0x5d, 0x0c, 0xeb, 0x5d, 0x04, 0x57, 0x1f, 0xac, 0x91,
	0xcb, 0x9a, 0x0a, 0x0e, 0xa5, 0xe2, 0x88, 0x33, 0x31, 0x2e, 0x8f, 0xee, 0xc6, 0xbd, 0xe7, 0x7c,
	0xdf, 0x79, 0xf4, 0x23, 0x9c, 0x9c, 0x13, 0xf2, 0x33, 0x9c, 0x9f, 0x9b, 0x48, 0x52, 0x6a, 0x68,
	0xe0, 0x80, 0xd2, 0x18, 0x2e, 0x9d, 0xf2, 0x97, 0xd3, 0x64, 0xa4, 0xd4, 0x99, 0xd7, 0x64, 0x69,
	0x9c, 0x79, 0xcd, 0x6d, 0x78, 0x6a, 0x08, 0x8c, 0x49, 0xa9, 0x31, 0xf6, 0x2a, 0x8c, 0xd3, 0x18,
	0x63, 0xaf, 0xcb, 0x1f, 0x8c, 0xc3, 0xbe, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xfd, 0xce,
	0x66, 0x6a, 0x03, 0x00, 0x00,
}
