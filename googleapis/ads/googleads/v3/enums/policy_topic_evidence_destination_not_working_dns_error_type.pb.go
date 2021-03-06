// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/policy_topic_evidence_destination_not_working_dns_error_type.proto

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

// The possible policy topic evidence destination not working DNS error types.
type PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType int32

const (
	// No value has been specified.
	PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_UNSPECIFIED PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_UNKNOWN PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType = 1
	// Host name not found in DNS when fetching landing page.
	PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_HOSTNAME_NOT_FOUND PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType = 2
	// Google internal crawler issue when communicating with DNS. This error
	// doesn't mean the landing page doesn't work. Google will recrawl the
	// landing page.
	PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_GOOGLE_CRAWLER_DNS_ISSUE PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType = 3
)

var PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "HOSTNAME_NOT_FOUND",
	3: "GOOGLE_CRAWLER_DNS_ISSUE",
}

var PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"HOSTNAME_NOT_FOUND":       2,
	"GOOGLE_CRAWLER_DNS_ISSUE": 3,
}

func (x PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType) String() string {
	return proto.EnumName(PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType_name, int32(x))
}

func (PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_773e4ae416a24f1f, []int{0, 0}
}

// Container for enum describing possible policy topic evidence destination not
// working DNS error types.
type PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) Reset() {
	*m = PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum{}
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) String() string {
	return proto.CompactTextString(m)
}
func (*PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) ProtoMessage() {}
func (*PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_773e4ae416a24f1f, []int{0}
}

func (m *PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum.Unmarshal(m, b)
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum.Marshal(b, m, deterministic)
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum.Merge(m, src)
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum.Size(m)
}
func (m *PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType", PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType_name, PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum_PolicyTopicEvidenceDestinationNotWorkingDnsErrorType_value)
	proto.RegisterType((*PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum)(nil), "google.ads.googleads.v3.enums.PolicyTopicEvidenceDestinationNotWorkingDnsErrorTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/policy_topic_evidence_destination_not_working_dns_error_type.proto", fileDescriptor_773e4ae416a24f1f)
}

var fileDescriptor_773e4ae416a24f1f = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0xaa, 0xd3, 0x40,
	0x18, 0x85, 0x4d, 0x2e, 0x28, 0xcc, 0x5d, 0x18, 0xb2, 0x10, 0x91, 0x7b, 0x17, 0xb7, 0x0f, 0x30,
	0x59, 0xc4, 0x85, 0x8e, 0xab, 0xb4, 0x99, 0xc6, 0x62, 0x9d, 0x84, 0x26, 0x69, 0x41, 0x02, 0x63,
	0xcc, 0x0c, 0x61, 0xb0, 0x9d, 0x09, 0x99, 0x69, 0xa5, 0xaf, 0xe0, 0xc2, 0x87, 0x70, 0xe9, 0x8b,
	0x08, 0x3e, 0x8a, 0x4f, 0x21, 0x99, 0xb4, 0x75, 0xa5, 0x8b, 0x6e, 0xc2, 0x21, 0xff, 0xfc, 0xe7,
	0x7c, 0x9c, 0x1f, 0x7c, 0x6c, 0x95, 0x6a, 0xb7, 0x3c, 0xa8, 0x99, 0x0e, 0x46, 0x39, 0xa8, 0x43,
	0x18, 0x70, 0xb9, 0xdf, 0xe9, 0xa0, 0x53, 0x5b, 0xd1, 0x1c, 0xa9, 0x51, 0x9d, 0x68, 0x28, 0x3f,
	0x08, 0xc6, 0x65, 0xc3, 0x29, 0xe3, 0xda, 0x08, 0x59, 0x1b, 0xa1, 0x24, 0x95, 0xca, 0xd0, 0x2f,
	0xaa, 0xff, 0x2c, 0x64, 0x4b, 0x99, 0xd4, 0x94, 0xf7, 0xbd, 0xea, 0xa9, 0x39, 0x76, 0x1c, 0x76,
	0xbd, 0x32, 0xca, 0xbf, 0x1f, 0x6d, 0x61, 0xcd, 0x34, 0xbc, 0x24, 0xc0, 0x43, 0x08, 0x6d, 0xc2,
	0x8b, 0xbb, 0x33, 0x40, 0x27, 0x82, 0x5a, 0x4a, 0x65, 0xac, 0xad, 0x1e, 0x97, 0x27, 0x3f, 0x1d,
	0xf0, 0x2a, 0xb3, 0x0c, 0xc5, 0x80, 0x80, 0x4f, 0x04, 0xf1, 0x5f, 0x00, 0xa2, 0xcc, 0x66, 0x8c,
	0x8f, 0xa5, 0xc6, 0x43, 0x78, 0x71, 0xec, 0x38, 0x96, 0xfb, 0xdd, 0xe4, 0xab, 0x03, 0x5e, 0x5e,
	0xb3, 0xec, 0x3f, 0x05, 0xb7, 0x25, 0xc9, 0x33, 0x3c, 0x5b, 0xcc, 0x17, 0x38, 0xf6, 0x1e, 0xf9,
	0xb7, 0xe0, 0x49, 0x49, 0xde, 0x91, 0x74, 0x43, 0x3c, 0xc7, 0x7f, 0x06, 0xfc, 0xb7, 0x69, 0x5e,
	0x90, 0xe8, 0x3d, 0xa6, 0x24, 0x2d, 0xe8, 0x3c, 0x2d, 0x49, 0xec, 0xb9, 0xfe, 0x1d, 0x78, 0x9e,
	0xa4, 0x69, 0xb2, 0xc4, 0x74, 0xb6, 0x8a, 0x36, 0x4b, 0xbc, 0xa2, 0x31, 0xc9, 0xe9, 0x22, 0xcf,
	0x4b, 0xec, 0xdd, 0x4c, 0xbf, 0xb9, 0xe0, 0xa1, 0x51, 0x3b, 0xf8, 0xdf, 0x36, 0xa6, 0xaf, 0xaf,
	0xe1, 0xcd, 0x86, 0xaa, 0x32, 0xe7, 0xc3, 0xf4, 0xe4, 0xdd, 0xaa, 0x6d, 0x2d, 0x5b, 0xa8, 0xfa,
	0x36, 0x68, 0xb9, 0xb4, 0x45, 0x9e, 0x6f, 0xdb, 0x09, 0xfd, 0x8f, 0x53, 0xbf, 0xb1, 0xdf, 0xef,
	0xee, 0x4d, 0x12, 0x45, 0x3f, 0xdc, 0xfb, 0x64, 0xb4, 0x8a, 0x98, 0x86, 0xa3, 0x1c, 0xd4, 0x3a,
	0x84, 0x43, 0xb1, 0xfa, 0xd7, 0x79, 0x5e, 0x45, 0x4c, 0x57, 0x97, 0x79, 0xb5, 0x0e, 0x2b, 0x3b,
	0xff, 0xed, 0x3e, 0x8c, 0x3f, 0x11, 0x8a, 0x98, 0x46, 0xe8, 0xf2, 0x02, 0xa1, 0x75, 0x88, 0x90,
	0x7d, 0xf3, 0xe9, 0xb1, 0x05, 0x0b, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x16, 0x63, 0x89,
	0x82, 0x02, 0x00, 0x00,
}
