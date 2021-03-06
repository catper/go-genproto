// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/campaign_draft.proto

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

// A campaign draft.
type CampaignDraft struct {
	// The resource name of the campaign draft.
	// Campaign draft resource names have the form:
	//
	// `customers/{customer_id}/campaignDrafts/{base_campaign_id}~{draft_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the draft.
	//
	// This field is read-only.
	DraftId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=draft_id,json=draftId,proto3" json:"draft_id,omitempty"`
	// The base campaign to which the draft belongs.
	BaseCampaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=base_campaign,json=baseCampaign,proto3" json:"base_campaign,omitempty"`
	// The name of the campaign draft.
	//
	// This field is required and should not be empty when creating new
	// campaign drafts.
	//
	// It must not contain any null (code point 0x0), NL line feed
	// (code point 0xA) or carriage return (code point 0xD) characters.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name of the Campaign that results from overlaying the draft
	// changes onto the base campaign.
	//
	// This field is read-only.
	DraftCampaign *wrappers.StringValue `protobuf:"bytes,5,opt,name=draft_campaign,json=draftCampaign,proto3" json:"draft_campaign,omitempty"`
	// The status of the campaign draft. This field is read-only.
	//
	// When a new campaign draft is added, the status defaults to PROPOSED.
	Status enums.CampaignDraftStatusEnum_CampaignDraftStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.CampaignDraftStatusEnum_CampaignDraftStatus" json:"status,omitempty"`
	// Whether there is an experiment based on this draft currently serving.
	HasExperimentRunning *wrappers.BoolValue `protobuf:"bytes,7,opt,name=has_experiment_running,json=hasExperimentRunning,proto3" json:"has_experiment_running,omitempty"`
	// The resource name of the long-running operation that can be used to poll
	// for completion of draft promotion. This is only set if the draft promotion
	// is in progress or finished.
	LongRunningOperation *wrappers.StringValue `protobuf:"bytes,8,opt,name=long_running_operation,json=longRunningOperation,proto3" json:"long_running_operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CampaignDraft) Reset()         { *m = CampaignDraft{} }
func (m *CampaignDraft) String() string { return proto.CompactTextString(m) }
func (*CampaignDraft) ProtoMessage()    {}
func (*CampaignDraft) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba592967968d7ea, []int{0}
}

func (m *CampaignDraft) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignDraft.Unmarshal(m, b)
}
func (m *CampaignDraft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignDraft.Marshal(b, m, deterministic)
}
func (m *CampaignDraft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignDraft.Merge(m, src)
}
func (m *CampaignDraft) XXX_Size() int {
	return xxx_messageInfo_CampaignDraft.Size(m)
}
func (m *CampaignDraft) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignDraft.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignDraft proto.InternalMessageInfo

func (m *CampaignDraft) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignDraft) GetDraftId() *wrappers.Int64Value {
	if m != nil {
		return m.DraftId
	}
	return nil
}

func (m *CampaignDraft) GetBaseCampaign() *wrappers.StringValue {
	if m != nil {
		return m.BaseCampaign
	}
	return nil
}

func (m *CampaignDraft) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CampaignDraft) GetDraftCampaign() *wrappers.StringValue {
	if m != nil {
		return m.DraftCampaign
	}
	return nil
}

func (m *CampaignDraft) GetStatus() enums.CampaignDraftStatusEnum_CampaignDraftStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignDraftStatusEnum_UNSPECIFIED
}

func (m *CampaignDraft) GetHasExperimentRunning() *wrappers.BoolValue {
	if m != nil {
		return m.HasExperimentRunning
	}
	return nil
}

func (m *CampaignDraft) GetLongRunningOperation() *wrappers.StringValue {
	if m != nil {
		return m.LongRunningOperation
	}
	return nil
}

func init() {
	proto.RegisterType((*CampaignDraft)(nil), "google.ads.googleads.v3.resources.CampaignDraft")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/campaign_draft.proto", fileDescriptor_8ba592967968d7ea)
}

var fileDescriptor_8ba592967968d7ea = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0x6e, 0xeb, 0x86, 0x59, 0x77, 0x11, 0x4d, 0x53, 0x28, 0x13, 0xea, 0x40, 0x43,
	0xbd, 0x72, 0xd0, 0x32, 0x55, 0x22, 0x5c, 0xa5, 0x63, 0x9a, 0xb6, 0x0b, 0xa8, 0x32, 0xa9, 0x17,
	0xa8, 0x22, 0x72, 0x13, 0xcf, 0x8b, 0x94, 0xd8, 0x91, 0xed, 0x0c, 0x24, 0xb4, 0xa7, 0xe0, 0x0d,
	0xb8, 0xe4, 0x51, 0x78, 0x94, 0xbd, 0x04, 0x28, 0x76, 0xec, 0x51, 0x46, 0xa1, 0x77, 0x27, 0x39,
	0xff, 0x77, 0xce, 0x7f, 0x7c, 0x6c, 0x30, 0x22, 0x8c, 0x91, 0x02, 0xfb, 0x28, 0x13, 0xbe, 0x0e,
	0x9b, 0xe8, 0x26, 0xf0, 0x39, 0x16, 0xac, 0xe6, 0x29, 0x16, 0x7e, 0x8a, 0xca, 0x0a, 0xe5, 0x84,
	0x26, 0x19, 0x47, 0x57, 0x12, 0x56, 0x9c, 0x49, 0xe6, 0x1e, 0x68, 0x31, 0x44, 0x99, 0x80, 0x96,
	0x83, 0x37, 0x01, 0xb4, 0x5c, 0xff, 0xf5, 0xb2, 0xd2, 0x98, 0xd6, 0xe5, 0x9f, 0x65, 0x13, 0x21,
	0x91, 0xac, 0x85, 0xae, 0xde, 0x7f, 0x62, 0xd0, 0x2a, 0xb7, 0x46, 0xda, 0xd4, 0xb3, 0x36, 0xa5,
	0xbe, 0xe6, 0xf5, 0x95, 0xff, 0x89, 0xa3, 0xaa, 0xc2, 0xdc, 0xa0, 0xfb, 0xbf, 0xa1, 0x88, 0x52,
	0x26, 0x91, 0xcc, 0x19, 0x6d, 0xb3, 0xcf, 0xbf, 0x6e, 0x80, 0xde, 0x49, 0xdb, 0xf8, 0x6d, 0xd3,
	0xd7, 0x7d, 0x01, 0x7a, 0xa6, 0x43, 0x42, 0x51, 0x89, 0x3d, 0x67, 0xe0, 0x0c, 0x1f, 0xc5, 0xdb,
	0xe6, 0xe7, 0x3b, 0x54, 0x62, 0x77, 0x04, 0xb6, 0xb4, 0xcb, 0x3c, 0xf3, 0x3a, 0x03, 0x67, 0xf8,
	0xf8, 0xe8, 0x69, 0x3b, 0x35, 0x34, 0x3e, 0xe0, 0x39, 0x95, 0xa3, 0xe3, 0x29, 0x2a, 0x6a, 0x1c,
	0x6f, 0x2a, 0xf1, 0x79, 0xe6, 0x46, 0xa0, 0x37, 0x47, 0x02, 0x27, 0x66, 0x56, 0x6f, 0x4d, 0xc1,
	0xfb, 0x0f, 0xe0, 0x4b, 0xc9, 0x73, 0x4a, 0x34, 0xbd, 0xdd, 0x20, 0xc6, 0xa4, 0xfb, 0x0a, 0xac,
	0x2b, 0x5b, 0xeb, 0x2b, 0x90, 0x4a, 0xe9, 0x9e, 0x80, 0x1d, 0x6d, 0xd6, 0x76, 0xdd, 0x58, 0x81,
	0xed, 0x29, 0xc6, 0xb6, 0x9d, 0x83, 0xae, 0xde, 0x88, 0xd7, 0x1d, 0x38, 0xc3, 0x9d, 0xa3, 0x0b,
	0xb8, 0x6c, 0xe1, 0x6a, 0x9b, 0x70, 0xe1, 0x50, 0x2f, 0x15, 0x79, 0x4a, 0xeb, 0xf2, 0x6f, 0xff,
	0xe3, 0xb6, 0xb2, 0x3b, 0x01, 0x7b, 0xd7, 0x48, 0x24, 0xf8, 0x73, 0x85, 0x79, 0x5e, 0x62, 0x2a,
	0x13, 0x5e, 0x53, 0x9a, 0x53, 0xe2, 0x6d, 0x2a, 0xc3, 0xfd, 0x07, 0x86, 0xc7, 0x8c, 0x15, 0xda,
	0xee, 0xee, 0x35, 0x12, 0xa7, 0x16, 0x8c, 0x35, 0xe7, 0xc6, 0x60, 0xaf, 0x60, 0x94, 0x98, 0x3a,
	0x09, 0xab, 0x30, 0x57, 0xfb, 0xf7, 0xb6, 0x56, 0x38, 0x82, 0xdd, 0x86, 0x6d, 0x4b, 0xbd, 0x37,
	0x64, 0x88, 0xee, 0xa2, 0x8f, 0xe0, 0xe5, 0xfd, 0xc8, 0x6d, 0x54, 0xe5, 0x02, 0xa6, 0xac, 0xf4,
	0x17, 0x6f, 0xd3, 0x71, 0x5a, 0x0b, 0xc9, 0x4a, 0xcc, 0x85, 0xff, 0xc5, 0x84, 0xb7, 0xf6, 0xaa,
	0x2b, 0x4d, 0x93, 0x59, 0xb8, 0xfa, 0xb7, 0xe3, 0x9f, 0x0e, 0x38, 0x4c, 0x59, 0x09, 0xff, 0xfb,
	0xa6, 0xc6, 0xee, 0x42, 0xbb, 0x49, 0x33, 0xc5, 0xc4, 0xf9, 0x70, 0xd1, 0x82, 0x84, 0x15, 0x88,
	0x12, 0xc8, 0x38, 0xf1, 0x09, 0xa6, 0x6a, 0x46, 0xff, 0xde, 0xee, 0x3f, 0xde, 0xf8, 0x1b, 0x1b,
	0x7d, 0xeb, 0xac, 0x9d, 0x45, 0xd1, 0xf7, 0xce, 0xc1, 0x99, 0x2e, 0x19, 0x65, 0x02, 0xea, 0xb0,
	0x89, 0xa6, 0x01, 0x8c, 0x8d, 0xf2, 0x87, 0xd1, 0xcc, 0xa2, 0x4c, 0xcc, 0xac, 0x66, 0x36, 0x0d,
	0x66, 0x56, 0x73, 0xd7, 0x39, 0xd4, 0x89, 0x30, 0x8c, 0x32, 0x11, 0x86, 0x56, 0x15, 0x86, 0xd3,
	0x20, 0x0c, 0xad, 0x6e, 0xde, 0x55, 0x66, 0x83, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0xbb,
	0x5e, 0x61, 0x8f, 0x04, 0x00, 0x00,
}
