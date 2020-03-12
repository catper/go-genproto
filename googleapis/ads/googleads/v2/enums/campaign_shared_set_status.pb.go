// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/campaign_shared_set_status.proto

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

// Enum listing the possible campaign shared set statuses.
type CampaignSharedSetStatusEnum_CampaignSharedSetStatus int32

const (
	// Not specified.
	CampaignSharedSetStatusEnum_UNSPECIFIED CampaignSharedSetStatusEnum_CampaignSharedSetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignSharedSetStatusEnum_UNKNOWN CampaignSharedSetStatusEnum_CampaignSharedSetStatus = 1
	// The campaign shared set is enabled.
	CampaignSharedSetStatusEnum_ENABLED CampaignSharedSetStatusEnum_CampaignSharedSetStatus = 2
	// The campaign shared set is removed and can no longer be used.
	CampaignSharedSetStatusEnum_REMOVED CampaignSharedSetStatusEnum_CampaignSharedSetStatus = 3
)

var CampaignSharedSetStatusEnum_CampaignSharedSetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var CampaignSharedSetStatusEnum_CampaignSharedSetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x CampaignSharedSetStatusEnum_CampaignSharedSetStatus) String() string {
	return proto.EnumName(CampaignSharedSetStatusEnum_CampaignSharedSetStatus_name, int32(x))
}

func (CampaignSharedSetStatusEnum_CampaignSharedSetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3e55b7a419c0b9db, []int{0, 0}
}

// Container for enum describing types of campaign shared set statuses.
type CampaignSharedSetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignSharedSetStatusEnum) Reset()         { *m = CampaignSharedSetStatusEnum{} }
func (m *CampaignSharedSetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSetStatusEnum) ProtoMessage()    {}
func (*CampaignSharedSetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e55b7a419c0b9db, []int{0}
}

func (m *CampaignSharedSetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSetStatusEnum.Unmarshal(m, b)
}
func (m *CampaignSharedSetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSetStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSetStatusEnum.Merge(m, src)
}
func (m *CampaignSharedSetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSetStatusEnum.Size(m)
}
func (m *CampaignSharedSetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus", CampaignSharedSetStatusEnum_CampaignSharedSetStatus_name, CampaignSharedSetStatusEnum_CampaignSharedSetStatus_value)
	proto.RegisterType((*CampaignSharedSetStatusEnum)(nil), "google.ads.googleads.v2.enums.CampaignSharedSetStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/campaign_shared_set_status.proto", fileDescriptor_3e55b7a419c0b9db)
}

var fileDescriptor_3e55b7a419c0b9db = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4a, 0xf3, 0x30,
	0x18, 0xfd, 0xd7, 0xc1, 0x2f, 0x74, 0x17, 0x96, 0xdd, 0x08, 0xba, 0x5d, 0x6c, 0x0f, 0x90, 0x40,
	0xbd, 0x8b, 0x20, 0xa4, 0x5b, 0x1c, 0x43, 0xed, 0xa6, 0x65, 0x15, 0xa4, 0x30, 0xe2, 0x12, 0x62,
	0x61, 0x4d, 0xca, 0xbe, 0x6c, 0x0f, 0xe4, 0xa5, 0x8f, 0xe2, 0xa3, 0x78, 0xe7, 0x1b, 0x48, 0xd2,
	0x6d, 0x77, 0xf3, 0xa6, 0x9c, 0xaf, 0xe7, 0x3b, 0xe7, 0x3b, 0x39, 0xe1, 0xad, 0x32, 0x46, 0xad,
	0x25, 0xe6, 0x02, 0x70, 0x03, 0x1d, 0xda, 0xc5, 0x58, 0xea, 0x6d, 0x05, 0x78, 0xc5, 0xab, 0x9a,
	0x97, 0x4a, 0x2f, 0xe1, 0x9d, 0x6f, 0xa4, 0x58, 0x82, 0xb4, 0x4b, 0xb0, 0xdc, 0x6e, 0x01, 0xd5,
	0x1b, 0x63, 0x4d, 0xb7, 0xdf, 0x88, 0x10, 0x17, 0x80, 0x8e, 0x7a, 0xb4, 0x8b, 0x91, 0xd7, 0x5f,
	0xf6, 0x0e, 0xf6, 0x75, 0x89, 0xb9, 0xd6, 0xc6, 0x72, 0x5b, 0x1a, 0xbd, 0x17, 0x0f, 0xeb, 0xf0,
	0x6a, 0xb4, 0x3f, 0x90, 0x79, 0xff, 0x4c, 0xda, 0xcc, 0xbb, 0x33, 0xbd, 0xad, 0x86, 0x4f, 0xe1,
	0xc5, 0x09, 0xba, 0x7b, 0x1e, 0x76, 0x16, 0x69, 0x36, 0x67, 0xa3, 0xe9, 0xdd, 0x94, 0x8d, 0xa3,
	0x7f, 0xdd, 0x4e, 0x78, 0xb6, 0x48, 0xef, 0xd3, 0xd9, 0x4b, 0x1a, 0xb5, 0xdc, 0xc0, 0x52, 0x9a,
	0x3c, 0xb0, 0x71, 0x14, 0xb8, 0xe1, 0x99, 0x3d, 0xce, 0x72, 0x36, 0x8e, 0xda, 0xc9, 0x4f, 0x2b,
	0x1c, 0xac, 0x4c, 0x85, 0xfe, 0x4c, 0x9d, 0xf4, 0x4e, 0x9c, 0x9d, 0xbb, 0xd4, 0xf3, 0xd6, 0x6b,
	0xb2, 0x97, 0x2b, 0xb3, 0xe6, 0x5a, 0x21, 0xb3, 0x51, 0x58, 0x49, 0xed, 0xdf, 0x74, 0x28, 0xb1,
	0x2e, 0xe1, 0x44, 0xa7, 0x37, 0xfe, 0xfb, 0x11, 0xb4, 0x27, 0x94, 0x7e, 0x06, 0xfd, 0x49, 0x63,
	0x45, 0x05, 0xa0, 0x06, 0x3a, 0x94, 0xc7, 0xc8, 0x35, 0x00, 0x5f, 0x07, 0xbe, 0xa0, 0x02, 0x8a,
	0x23, 0x5f, 0xe4, 0x71, 0xe1, 0xf9, 0xef, 0x60, 0xd0, 0xfc, 0x24, 0x84, 0x0a, 0x20, 0xe4, 0xb8,
	0x41, 0x48, 0x1e, 0x13, 0xe2, 0x77, 0xde, 0xfe, 0xfb, 0x60, 0xd7, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa1, 0x2e, 0xc2, 0x0b, 0xeb, 0x01, 0x00, 0x00,
}
