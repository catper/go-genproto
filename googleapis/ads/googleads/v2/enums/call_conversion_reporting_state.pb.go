// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/call_conversion_reporting_state.proto

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

// Possible data types for a call conversion action state.
type CallConversionReportingStateEnum_CallConversionReportingState int32

const (
	// Not specified.
	CallConversionReportingStateEnum_UNSPECIFIED CallConversionReportingStateEnum_CallConversionReportingState = 0
	// Used for return value only. Represents value unknown in this version.
	CallConversionReportingStateEnum_UNKNOWN CallConversionReportingStateEnum_CallConversionReportingState = 1
	// Call conversion action is disabled.
	CallConversionReportingStateEnum_DISABLED CallConversionReportingStateEnum_CallConversionReportingState = 2
	// Call conversion action will use call conversion type set at the
	// account level.
	CallConversionReportingStateEnum_USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION CallConversionReportingStateEnum_CallConversionReportingState = 3
	// Call conversion action will use call conversion type set at the resource
	// (call only ads/call extensions) level.
	CallConversionReportingStateEnum_USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION CallConversionReportingStateEnum_CallConversionReportingState = 4
)

var CallConversionReportingStateEnum_CallConversionReportingState_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DISABLED",
	3: "USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION",
	4: "USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION",
}

var CallConversionReportingStateEnum_CallConversionReportingState_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"DISABLED":    2,
	"USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION":  3,
	"USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION": 4,
}

func (x CallConversionReportingStateEnum_CallConversionReportingState) String() string {
	return proto.EnumName(CallConversionReportingStateEnum_CallConversionReportingState_name, int32(x))
}

func (CallConversionReportingStateEnum_CallConversionReportingState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce9993055890e451, []int{0, 0}
}

// Container for enum describing possible data types for call conversion
// reporting state.
type CallConversionReportingStateEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallConversionReportingStateEnum) Reset()         { *m = CallConversionReportingStateEnum{} }
func (m *CallConversionReportingStateEnum) String() string { return proto.CompactTextString(m) }
func (*CallConversionReportingStateEnum) ProtoMessage()    {}
func (*CallConversionReportingStateEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce9993055890e451, []int{0}
}

func (m *CallConversionReportingStateEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallConversionReportingStateEnum.Unmarshal(m, b)
}
func (m *CallConversionReportingStateEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallConversionReportingStateEnum.Marshal(b, m, deterministic)
}
func (m *CallConversionReportingStateEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallConversionReportingStateEnum.Merge(m, src)
}
func (m *CallConversionReportingStateEnum) XXX_Size() int {
	return xxx_messageInfo_CallConversionReportingStateEnum.Size(m)
}
func (m *CallConversionReportingStateEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CallConversionReportingStateEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CallConversionReportingStateEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.CallConversionReportingStateEnum_CallConversionReportingState", CallConversionReportingStateEnum_CallConversionReportingState_name, CallConversionReportingStateEnum_CallConversionReportingState_value)
	proto.RegisterType((*CallConversionReportingStateEnum)(nil), "google.ads.googleads.v2.enums.CallConversionReportingStateEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/call_conversion_reporting_state.proto", fileDescriptor_ce9993055890e451)
}

var fileDescriptor_ce9993055890e451 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0xeb, 0x9b, 0x30,
	0x1c, 0x9d, 0x76, 0x6c, 0x23, 0x1d, 0x4c, 0x3c, 0x8e, 0x16, 0xd6, 0x9e, 0x36, 0xd8, 0x22, 0xb8,
	0x5b, 0x76, 0x8a, 0x69, 0x56, 0x64, 0x12, 0x8b, 0x56, 0x07, 0x43, 0x08, 0x59, 0x15, 0x11, 0x6c,
	0x22, 0xc6, 0xf6, 0x03, 0xed, 0x32, 0xd8, 0x47, 0x19, 0xfb, 0x24, 0xfb, 0x00, 0x3b, 0x0f, 0xf5,
	0xaf, 0xb7, 0xbf, 0x97, 0xf0, 0xc8, 0x7b, 0xbf, 0xf7, 0x92, 0xf7, 0x03, 0xa4, 0x54, 0xaa, 0xac,
	0x0b, 0x47, 0xe4, 0xda, 0x19, 0x61, 0x8f, 0xee, 0xae, 0x53, 0xc8, 0xdb, 0x55, 0x3b, 0x17, 0x51,
	0xd7, 0xfc, 0xa2, 0xe4, 0xbd, 0x68, 0x75, 0xa5, 0x24, 0x6f, 0x8b, 0x46, 0xb5, 0x5d, 0x25, 0x4b,
	0xae, 0x3b, 0xd1, 0x15, 0xb0, 0x69, 0x55, 0xa7, 0xec, 0xed, 0x38, 0x09, 0x45, 0xae, 0xe1, 0x6c,
	0x02, 0xef, 0x2e, 0x1c, 0x4c, 0x5e, 0x6f, 0xa6, 0x8c, 0xa6, 0x72, 0x84, 0x94, 0xaa, 0x13, 0x5d,
	0xa5, 0xa4, 0x1e, 0x87, 0xf7, 0x7f, 0x0c, 0xf0, 0x86, 0x88, 0xba, 0x26, 0x73, 0x4a, 0x34, 0x85,
	0xc4, 0x7d, 0x06, 0x95, 0xb7, 0xeb, 0xfe, 0xa7, 0x01, 0x36, 0x4b, 0x22, 0xfb, 0x15, 0x58, 0x27,
	0x2c, 0x3e, 0x51, 0xe2, 0x7f, 0xf6, 0xe9, 0xc1, 0x7a, 0x62, 0xaf, 0xc1, 0xf3, 0x84, 0x7d, 0x61,
	0xe1, 0x57, 0x66, 0x19, 0xf6, 0x4b, 0xf0, 0xe2, 0xe0, 0xc7, 0xd8, 0x0b, 0xe8, 0xc1, 0x32, 0xed,
	0xf7, 0xe0, 0x6d, 0x12, 0x53, 0x8e, 0x09, 0x09, 0x13, 0x76, 0xe6, 0x01, 0x4d, 0x69, 0xc0, 0x09,
	0x0e, 0x02, 0x4e, 0x42, 0x96, 0xd2, 0x28, 0xf6, 0x43, 0xc6, 0x31, 0x39, 0xfb, 0x21, 0xb3, 0x56,
	0xf6, 0x07, 0xf0, 0xae, 0x57, 0x47, 0x34, 0x0e, 0x93, 0x88, 0xd0, 0x65, 0xf9, 0x53, 0xef, 0x9f,
	0x01, 0x76, 0x17, 0x75, 0x85, 0x8b, 0x95, 0x78, 0xbb, 0xa5, 0xcf, 0x9c, 0xfa, 0x5e, 0x4e, 0xc6,
	0x37, 0xef, 0xc1, 0xa3, 0x54, 0xb5, 0x90, 0x25, 0x54, 0x6d, 0xe9, 0x94, 0x85, 0x1c, 0x5a, 0x9b,
	0x76, 0xd5, 0x54, 0xfa, 0x91, 0xd5, 0x7d, 0x1a, 0xce, 0x1f, 0xe6, 0xea, 0x88, 0xf1, 0x2f, 0x73,
	0x7b, 0x1c, 0xad, 0x70, 0xae, 0xe1, 0x08, 0x7b, 0x94, 0xba, 0xb0, 0x6f, 0x57, 0xff, 0x9e, 0xf8,
	0x0c, 0xe7, 0x3a, 0x9b, 0xf9, 0x2c, 0x75, 0xb3, 0x81, 0xff, 0x6b, 0xee, 0xc6, 0x4b, 0x84, 0x70,
	0xae, 0x11, 0x9a, 0x15, 0x08, 0xa5, 0x2e, 0x42, 0x83, 0xe6, 0xfb, 0xb3, 0xe1, 0x61, 0x1f, 0xff,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x88, 0x3b, 0xea, 0xea, 0x52, 0x02, 0x00, 0x00,
}
