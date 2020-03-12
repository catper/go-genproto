// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/billing_setup_status.proto

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

// The possible statuses of a BillingSetup.
type BillingSetupStatusEnum_BillingSetupStatus int32

const (
	// Not specified.
	BillingSetupStatusEnum_UNSPECIFIED BillingSetupStatusEnum_BillingSetupStatus = 0
	// Used for return value only. Represents value unknown in this version.
	BillingSetupStatusEnum_UNKNOWN BillingSetupStatusEnum_BillingSetupStatus = 1
	// The billing setup is pending approval.
	BillingSetupStatusEnum_PENDING BillingSetupStatusEnum_BillingSetupStatus = 2
	// The billing setup has been approved but the corresponding first budget
	// has not.  This can only occur for billing setups configured for monthly
	// invoicing.
	BillingSetupStatusEnum_APPROVED_HELD BillingSetupStatusEnum_BillingSetupStatus = 3
	// The billing setup has been approved.
	BillingSetupStatusEnum_APPROVED BillingSetupStatusEnum_BillingSetupStatus = 4
	// The billing setup was cancelled by the user prior to approval.
	BillingSetupStatusEnum_CANCELLED BillingSetupStatusEnum_BillingSetupStatus = 5
)

var BillingSetupStatusEnum_BillingSetupStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "APPROVED_HELD",
	4: "APPROVED",
	5: "CANCELLED",
}

var BillingSetupStatusEnum_BillingSetupStatus_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"PENDING":       2,
	"APPROVED_HELD": 3,
	"APPROVED":      4,
	"CANCELLED":     5,
}

func (x BillingSetupStatusEnum_BillingSetupStatus) String() string {
	return proto.EnumName(BillingSetupStatusEnum_BillingSetupStatus_name, int32(x))
}

func (BillingSetupStatusEnum_BillingSetupStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0274de631061209b, []int{0, 0}
}

// Message describing BillingSetup statuses.
type BillingSetupStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillingSetupStatusEnum) Reset()         { *m = BillingSetupStatusEnum{} }
func (m *BillingSetupStatusEnum) String() string { return proto.CompactTextString(m) }
func (*BillingSetupStatusEnum) ProtoMessage()    {}
func (*BillingSetupStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0274de631061209b, []int{0}
}

func (m *BillingSetupStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetupStatusEnum.Unmarshal(m, b)
}
func (m *BillingSetupStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetupStatusEnum.Marshal(b, m, deterministic)
}
func (m *BillingSetupStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetupStatusEnum.Merge(m, src)
}
func (m *BillingSetupStatusEnum) XXX_Size() int {
	return xxx_messageInfo_BillingSetupStatusEnum.Size(m)
}
func (m *BillingSetupStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetupStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetupStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.BillingSetupStatusEnum_BillingSetupStatus", BillingSetupStatusEnum_BillingSetupStatus_name, BillingSetupStatusEnum_BillingSetupStatus_value)
	proto.RegisterType((*BillingSetupStatusEnum)(nil), "google.ads.googleads.v3.enums.BillingSetupStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/billing_setup_status.proto", fileDescriptor_0274de631061209b)
}

var fileDescriptor_0274de631061209b = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4a, 0xc3, 0x30,
	0x18, 0xb5, 0x9d, 0xbf, 0x99, 0xc3, 0x9a, 0x0b, 0x05, 0x71, 0x17, 0xdb, 0x03, 0xa4, 0x17, 0xbd,
	0x91, 0x78, 0x95, 0xae, 0x71, 0x0e, 0x47, 0x56, 0x1c, 0xab, 0x20, 0x85, 0xd1, 0xd9, 0x12, 0x0a,
	0x5d, 0x52, 0x96, 0x76, 0xef, 0xe0, 0x6b, 0x78, 0xe9, 0xa3, 0xf8, 0x28, 0xe2, 0x43, 0x48, 0xd3,
	0xb5, 0x37, 0x43, 0x6f, 0xc2, 0xc9, 0x77, 0xbe, 0x73, 0xf8, 0xce, 0x01, 0x77, 0x5c, 0x4a, 0x9e,
	0x25, 0x76, 0x14, 0x2b, 0xbb, 0x86, 0x15, 0xda, 0x3a, 0x76, 0x22, 0xca, 0xb5, 0xb2, 0x57, 0x69,
	0x96, 0xa5, 0x82, 0x2f, 0x55, 0x52, 0x94, 0xf9, 0x52, 0x15, 0x51, 0x51, 0x2a, 0x94, 0x6f, 0x64,
	0x21, 0x61, 0xbf, 0x5e, 0x47, 0x51, 0xac, 0x50, 0xab, 0x44, 0x5b, 0x07, 0x69, 0xe5, 0xcd, 0x6d,
	0x63, 0x9c, 0xa7, 0x76, 0x24, 0x84, 0x2c, 0xa2, 0x22, 0x95, 0x62, 0x27, 0x1e, 0xbe, 0x1b, 0xe0,
	0xca, 0xad, 0xbd, 0xe7, 0x95, 0xf5, 0x5c, 0x3b, 0x53, 0x51, 0xae, 0x87, 0x12, 0xc0, 0x7d, 0x06,
	0x5e, 0x80, 0xee, 0x82, 0xcd, 0x7d, 0x3a, 0x9a, 0x3c, 0x4c, 0xa8, 0x67, 0x1d, 0xc0, 0x2e, 0x38,
	0x59, 0xb0, 0x27, 0x36, 0x7b, 0x61, 0x96, 0x51, 0x7d, 0x7c, 0xca, 0xbc, 0x09, 0x1b, 0x5b, 0x26,
	0xbc, 0x04, 0x3d, 0xe2, 0xfb, 0xcf, 0xb3, 0x80, 0x7a, 0xcb, 0x47, 0x3a, 0xf5, 0xac, 0x0e, 0x3c,
	0x07, 0xa7, 0xcd, 0xc8, 0x3a, 0x84, 0x3d, 0x70, 0x36, 0x22, 0x6c, 0x44, 0xa7, 0x53, 0xea, 0x59,
	0x47, 0xee, 0x8f, 0x01, 0x06, 0x6f, 0x72, 0x8d, 0xfe, 0xcd, 0xe3, 0x5e, 0xef, 0x1f, 0xe5, 0x57,
	0x51, 0x7c, 0xe3, 0xd5, 0xdd, 0x29, 0xb9, 0xcc, 0x22, 0xc1, 0x91, 0xdc, 0x70, 0x9b, 0x27, 0x42,
	0x07, 0x6d, 0x3a, 0xcd, 0x53, 0xf5, 0x47, 0xc5, 0xf7, 0xfa, 0xfd, 0x30, 0x3b, 0x63, 0x42, 0x3e,
	0xcd, 0xfe, 0xb8, 0xb6, 0x22, 0xb1, 0x42, 0x35, 0xac, 0x50, 0xe0, 0xa0, 0xaa, 0x1a, 0xf5, 0xd5,
	0xf0, 0x21, 0x89, 0x55, 0xd8, 0xf2, 0x61, 0xe0, 0x84, 0x9a, 0xff, 0x36, 0x07, 0xf5, 0x10, 0x63,
	0x12, 0x2b, 0x8c, 0xdb, 0x0d, 0x8c, 0x03, 0x07, 0x63, 0xbd, 0xb3, 0x3a, 0xd6, 0x87, 0x39, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x5e, 0x25, 0x32, 0xfa, 0x01, 0x00, 0x00,
}
