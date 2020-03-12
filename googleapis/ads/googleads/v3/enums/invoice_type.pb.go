// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/invoice_type.proto

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

// The possible type of invoices.
type InvoiceTypeEnum_InvoiceType int32

const (
	// Not specified.
	InvoiceTypeEnum_UNSPECIFIED InvoiceTypeEnum_InvoiceType = 0
	// Used for return value only. Represents value unknown in this version.
	InvoiceTypeEnum_UNKNOWN InvoiceTypeEnum_InvoiceType = 1
	// An invoice with a negative amount. The account receives a credit.
	InvoiceTypeEnum_CREDIT_MEMO InvoiceTypeEnum_InvoiceType = 2
	// An invoice with a positive amount. The account owes a balance.
	InvoiceTypeEnum_INVOICE InvoiceTypeEnum_InvoiceType = 3
)

var InvoiceTypeEnum_InvoiceType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CREDIT_MEMO",
	3: "INVOICE",
}

var InvoiceTypeEnum_InvoiceType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CREDIT_MEMO": 2,
	"INVOICE":     3,
}

func (x InvoiceTypeEnum_InvoiceType) String() string {
	return proto.EnumName(InvoiceTypeEnum_InvoiceType_name, int32(x))
}

func (InvoiceTypeEnum_InvoiceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22c008ffe76f80f6, []int{0, 0}
}

// Container for enum describing the type of invoices.
type InvoiceTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceTypeEnum) Reset()         { *m = InvoiceTypeEnum{} }
func (m *InvoiceTypeEnum) String() string { return proto.CompactTextString(m) }
func (*InvoiceTypeEnum) ProtoMessage()    {}
func (*InvoiceTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_22c008ffe76f80f6, []int{0}
}

func (m *InvoiceTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceTypeEnum.Unmarshal(m, b)
}
func (m *InvoiceTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceTypeEnum.Marshal(b, m, deterministic)
}
func (m *InvoiceTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceTypeEnum.Merge(m, src)
}
func (m *InvoiceTypeEnum) XXX_Size() int {
	return xxx_messageInfo_InvoiceTypeEnum.Size(m)
}
func (m *InvoiceTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.InvoiceTypeEnum_InvoiceType", InvoiceTypeEnum_InvoiceType_name, InvoiceTypeEnum_InvoiceType_value)
	proto.RegisterType((*InvoiceTypeEnum)(nil), "google.ads.googleads.v3.enums.InvoiceTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/invoice_type.proto", fileDescriptor_22c008ffe76f80f6)
}

var fileDescriptor_22c008ffe76f80f6 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xbd, 0x6a, 0xf3, 0x30,
	0x14, 0xfd, 0xe2, 0xc0, 0x57, 0x50, 0x86, 0x98, 0x8c, 0xa5, 0x19, 0x92, 0x07, 0x90, 0x0a, 0xda,
	0xd4, 0xc9, 0x71, 0xd4, 0x20, 0x4a, 0x6c, 0xd3, 0x26, 0x2e, 0x14, 0x43, 0x70, 0x63, 0x21, 0x0c,
	0xb1, 0x64, 0x22, 0xc7, 0x90, 0xd7, 0xe9, 0xd8, 0x47, 0xe9, 0x7b, 0x74, 0xe9, 0x53, 0x14, 0x49,
	0xb5, 0xc9, 0xd2, 0x2e, 0xe2, 0x70, 0xcf, 0x8f, 0xce, 0xbd, 0xe0, 0x56, 0x28, 0x25, 0x0e, 0x1c,
	0xe5, 0x85, 0x46, 0x0e, 0x1a, 0xd4, 0x62, 0xc4, 0xe5, 0xa9, 0xd2, 0xa8, 0x94, 0xad, 0x2a, 0xf7,
	0x7c, 0xd7, 0x9c, 0x6b, 0x0e, 0xeb, 0xa3, 0x6a, 0xd4, 0x64, 0xea, 0x64, 0x30, 0x2f, 0x34, 0xec,
	0x1d, 0xb0, 0xc5, 0xd0, 0x3a, 0xae, 0x6f, 0xba, 0xc0, 0xba, 0x44, 0xb9, 0x94, 0xaa, 0xc9, 0x9b,
	0x52, 0x49, 0xed, 0xcc, 0xf3, 0x0c, 0x8c, 0x99, 0x8b, 0xdc, 0x9c, 0x6b, 0x4e, 0xe5, 0xa9, 0x9a,
	0x33, 0x30, 0xba, 0x18, 0x4d, 0xc6, 0x60, 0xb4, 0x8d, 0x9e, 0x12, 0x1a, 0xb2, 0x7b, 0x46, 0x97,
	0xfe, 0xbf, 0xc9, 0x08, 0x5c, 0x6d, 0xa3, 0x87, 0x28, 0x7e, 0x8e, 0xfc, 0x81, 0x61, 0xc3, 0x47,
	0xba, 0x64, 0x9b, 0xdd, 0x9a, 0xae, 0x63, 0xdf, 0x33, 0x2c, 0x8b, 0xd2, 0x98, 0x85, 0xd4, 0x1f,
	0x2e, 0x3e, 0x07, 0x60, 0xb6, 0x57, 0x15, 0xfc, 0xb3, 0xe1, 0xc2, 0xbf, 0xf8, 0x2e, 0x31, 0xad,
	0x92, 0xc1, 0xcb, 0xe2, 0xc7, 0x22, 0xd4, 0x21, 0x97, 0x02, 0xaa, 0xa3, 0x40, 0x82, 0x4b, 0xdb,
	0xb9, 0x3b, 0x4b, 0x5d, 0xea, 0x5f, 0xae, 0x74, 0x67, 0xdf, 0x37, 0x6f, 0xb8, 0x0a, 0x82, 0x77,
	0x6f, 0xba, 0x72, 0x51, 0x41, 0xa1, 0xa1, 0x83, 0x06, 0xa5, 0x18, 0x9a, 0x6d, 0xf5, 0x47, 0xc7,
	0x67, 0x41, 0xa1, 0xb3, 0x9e, 0xcf, 0x52, 0x9c, 0x59, 0xfe, 0xcb, 0x9b, 0xb9, 0x21, 0x21, 0x41,
	0xa1, 0x09, 0xe9, 0x15, 0x84, 0xa4, 0x98, 0x10, 0xab, 0x79, 0xfd, 0x6f, 0x8b, 0xe1, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x78, 0x11, 0xae, 0x31, 0xbd, 0x01, 0x00, 0x00,
}
