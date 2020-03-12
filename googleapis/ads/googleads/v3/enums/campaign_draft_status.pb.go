// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/campaign_draft_status.proto

package enums

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

// Possible statuses of a campaign draft.
type CampaignDraftStatusEnum_CampaignDraftStatus int32

const (
	// The status has not been specified.
	CampaignDraftStatusEnum_UNSPECIFIED CampaignDraftStatusEnum_CampaignDraftStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignDraftStatusEnum_UNKNOWN CampaignDraftStatusEnum_CampaignDraftStatus = 1
	// Initial state of the draft, the advertiser can start adding changes with
	// no effect on serving.
	CampaignDraftStatusEnum_PROPOSED CampaignDraftStatusEnum_CampaignDraftStatus = 2
	// The campaign draft is removed.
	CampaignDraftStatusEnum_REMOVED CampaignDraftStatusEnum_CampaignDraftStatus = 3
	// Advertiser requested to promote draft's changes back into the original
	// campaign. Advertiser can poll the long running operation returned by
	// the promote action to see the status of the promotion.
	CampaignDraftStatusEnum_PROMOTING CampaignDraftStatusEnum_CampaignDraftStatus = 5
	// The process to merge changes in the draft back to the original campaign
	// has completed successfully.
	CampaignDraftStatusEnum_PROMOTED CampaignDraftStatusEnum_CampaignDraftStatus = 4
	// The promotion failed after it was partially applied. Promote cannot be
	// attempted again safely, so the issue must be corrected in the original
	// campaign.
	CampaignDraftStatusEnum_PROMOTE_FAILED CampaignDraftStatusEnum_CampaignDraftStatus = 6
)

var CampaignDraftStatusEnum_CampaignDraftStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PROPOSED",
	3: "REMOVED",
	5: "PROMOTING",
	4: "PROMOTED",
	6: "PROMOTE_FAILED",
}

var CampaignDraftStatusEnum_CampaignDraftStatus_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"PROPOSED":       2,
	"REMOVED":        3,
	"PROMOTING":      5,
	"PROMOTED":       4,
	"PROMOTE_FAILED": 6,
}

func (x CampaignDraftStatusEnum_CampaignDraftStatus) String() string {
	return proto.EnumName(CampaignDraftStatusEnum_CampaignDraftStatus_name, int32(x))
}

func (CampaignDraftStatusEnum_CampaignDraftStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_70c9d2fa31d31347, []int{0, 0}
}

// Container for enum describing possible statuses of a campaign draft.
type CampaignDraftStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignDraftStatusEnum) Reset()         { *m = CampaignDraftStatusEnum{} }
func (m *CampaignDraftStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignDraftStatusEnum) ProtoMessage()    {}
func (*CampaignDraftStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_70c9d2fa31d31347, []int{0}
}

func (m *CampaignDraftStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignDraftStatusEnum.Unmarshal(m, b)
}
func (m *CampaignDraftStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignDraftStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignDraftStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignDraftStatusEnum.Merge(m, src)
}
func (m *CampaignDraftStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignDraftStatusEnum.Size(m)
}
func (m *CampaignDraftStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignDraftStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignDraftStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.CampaignDraftStatusEnum_CampaignDraftStatus", CampaignDraftStatusEnum_CampaignDraftStatus_name, CampaignDraftStatusEnum_CampaignDraftStatus_value)
	proto.RegisterType((*CampaignDraftStatusEnum)(nil), "google.ads.googleads.v3.enums.CampaignDraftStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/campaign_draft_status.proto", fileDescriptor_70c9d2fa31d31347)
}

var fileDescriptor_70c9d2fa31d31347 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x4e, 0xe3, 0x30,
	0x18, 0x9d, 0xa4, 0x33, 0x9d, 0x19, 0x97, 0x9f, 0xc8, 0x2c, 0x40, 0x88, 0x2e, 0xda, 0x03, 0x38,
	0x8b, 0xac, 0x30, 0xab, 0xb4, 0x71, 0xab, 0x08, 0x1a, 0x47, 0xfd, 0x09, 0x12, 0x8a, 0x54, 0x99,
	0xa6, 0x58, 0x91, 0x1a, 0x3b, 0xaa, 0xd3, 0x6e, 0xb9, 0x07, 0x4b, 0x96, 0x1c, 0x85, 0xa3, 0x20,
	0x71, 0x07, 0x14, 0x27, 0xed, 0xaa, 0xb0, 0x89, 0x5e, 0xfc, 0x7e, 0xf4, 0xbd, 0x07, 0xae, 0xb9,
	0x94, 0x7c, 0xb5, 0xb4, 0x59, 0xa2, 0xec, 0x0a, 0x96, 0x68, 0xeb, 0xd8, 0x4b, 0xb1, 0xc9, 0x94,
	0xbd, 0x60, 0x59, 0xce, 0x52, 0x2e, 0xe6, 0xc9, 0x9a, 0x3d, 0x15, 0x73, 0x55, 0xb0, 0x62, 0xa3,
	0x50, 0xbe, 0x96, 0x85, 0x84, 0xed, 0x4a, 0x8f, 0x58, 0xa2, 0xd0, 0xde, 0x8a, 0xb6, 0x0e, 0xd2,
	0xd6, 0xcb, 0xab, 0x5d, 0x72, 0x9e, 0xda, 0x4c, 0x08, 0x59, 0xb0, 0x22, 0x95, 0xa2, 0x36, 0x77,
	0x5f, 0x0c, 0x70, 0xde, 0xaf, 0xc3, 0xbd, 0x32, 0x7b, 0xa2, 0xa3, 0x89, 0xd8, 0x64, 0xdd, 0x67,
	0x70, 0x76, 0x80, 0x82, 0xa7, 0xa0, 0x35, 0x0b, 0x26, 0x21, 0xe9, 0xfb, 0x03, 0x9f, 0x78, 0xd6,
	0x2f, 0xd8, 0x02, 0x7f, 0x67, 0xc1, 0x6d, 0x40, 0xef, 0x03, 0xcb, 0x80, 0x47, 0xe0, 0x5f, 0x38,
	0xa6, 0x21, 0x9d, 0x10, 0xcf, 0x32, 0x4b, 0x6a, 0x4c, 0x46, 0x34, 0x22, 0x9e, 0xd5, 0x80, 0xc7,
	0xe0, 0x7f, 0x38, 0xa6, 0x23, 0x3a, 0xf5, 0x83, 0xa1, 0xf5, 0xa7, 0x56, 0x8e, 0xe8, 0x94, 0x78,
	0xd6, 0x6f, 0x08, 0xc1, 0x49, 0xfd, 0x37, 0x1f, 0xb8, 0xfe, 0x1d, 0xf1, 0xac, 0x66, 0xef, 0xd3,
	0x00, 0x9d, 0x85, 0xcc, 0xd0, 0x8f, 0x05, 0x7b, 0x17, 0x07, 0x8e, 0x0c, 0xcb, 0x72, 0xa1, 0xf1,
	0xd0, 0xab, 0xad, 0x5c, 0xae, 0x98, 0xe0, 0x48, 0xae, 0xb9, 0xcd, 0x97, 0x42, 0x57, 0xdf, 0xcd,
	0x9c, 0xa7, 0xea, 0x9b, 0xd5, 0x6f, 0xf4, 0xf7, 0xd5, 0x6c, 0x0c, 0x5d, 0xf7, 0xcd, 0x6c, 0x0f,
	0xab, 0x28, 0x37, 0x51, 0xa8, 0x82, 0x25, 0x8a, 0x1c, 0x54, 0x6e, 0xa5, 0xde, 0x77, 0x7c, 0xec,
	0x26, 0x2a, 0xde, 0xf3, 0x71, 0xe4, 0xc4, 0x9a, 0xff, 0x30, 0x3b, 0xd5, 0x23, 0xc6, 0x6e, 0xa2,
	0x30, 0xde, 0x2b, 0x30, 0x8e, 0x1c, 0x8c, 0xb5, 0xe6, 0xb1, 0xa9, 0x0f, 0x73, 0xbe, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x6e, 0x06, 0x85, 0xe8, 0x0d, 0x02, 0x00, 0x00,
}
