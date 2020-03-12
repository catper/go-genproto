// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/data_driven_model_status.proto

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

// Enumerates data driven model statuses.
type DataDrivenModelStatusEnum_DataDrivenModelStatus int32

const (
	// Not specified.
	DataDrivenModelStatusEnum_UNSPECIFIED DataDrivenModelStatusEnum_DataDrivenModelStatus = 0
	// Used for return value only. Represents value unknown in this version.
	DataDrivenModelStatusEnum_UNKNOWN DataDrivenModelStatusEnum_DataDrivenModelStatus = 1
	// The data driven model is available.
	DataDrivenModelStatusEnum_AVAILABLE DataDrivenModelStatusEnum_DataDrivenModelStatus = 2
	// The data driven model is stale. It hasn't been updated for at least 7
	// days. It is still being used, but will become expired if it does not get
	// updated for 30 days.
	DataDrivenModelStatusEnum_STALE DataDrivenModelStatusEnum_DataDrivenModelStatus = 3
	// The data driven model expired. It hasn't been updated for at least 30
	// days and cannot be used. Most commonly this is because there hasn't been
	// the required number of events in a recent 30-day period.
	DataDrivenModelStatusEnum_EXPIRED DataDrivenModelStatusEnum_DataDrivenModelStatus = 4
	// The data driven model has never been generated. Most commonly this is
	// because there has never been the required number of events in any 30-day
	// period.
	DataDrivenModelStatusEnum_NEVER_GENERATED DataDrivenModelStatusEnum_DataDrivenModelStatus = 5
)

var DataDrivenModelStatusEnum_DataDrivenModelStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AVAILABLE",
	3: "STALE",
	4: "EXPIRED",
	5: "NEVER_GENERATED",
}

var DataDrivenModelStatusEnum_DataDrivenModelStatus_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"AVAILABLE":       2,
	"STALE":           3,
	"EXPIRED":         4,
	"NEVER_GENERATED": 5,
}

func (x DataDrivenModelStatusEnum_DataDrivenModelStatus) String() string {
	return proto.EnumName(DataDrivenModelStatusEnum_DataDrivenModelStatus_name, int32(x))
}

func (DataDrivenModelStatusEnum_DataDrivenModelStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c5ef206018e780c1, []int{0, 0}
}

// Container for enum indicating data driven model status.
type DataDrivenModelStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataDrivenModelStatusEnum) Reset()         { *m = DataDrivenModelStatusEnum{} }
func (m *DataDrivenModelStatusEnum) String() string { return proto.CompactTextString(m) }
func (*DataDrivenModelStatusEnum) ProtoMessage()    {}
func (*DataDrivenModelStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5ef206018e780c1, []int{0}
}

func (m *DataDrivenModelStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataDrivenModelStatusEnum.Unmarshal(m, b)
}
func (m *DataDrivenModelStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataDrivenModelStatusEnum.Marshal(b, m, deterministic)
}
func (m *DataDrivenModelStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataDrivenModelStatusEnum.Merge(m, src)
}
func (m *DataDrivenModelStatusEnum) XXX_Size() int {
	return xxx_messageInfo_DataDrivenModelStatusEnum.Size(m)
}
func (m *DataDrivenModelStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DataDrivenModelStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DataDrivenModelStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.DataDrivenModelStatusEnum_DataDrivenModelStatus", DataDrivenModelStatusEnum_DataDrivenModelStatus_name, DataDrivenModelStatusEnum_DataDrivenModelStatus_value)
	proto.RegisterType((*DataDrivenModelStatusEnum)(nil), "google.ads.googleads.v2.enums.DataDrivenModelStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/data_driven_model_status.proto", fileDescriptor_c5ef206018e780c1)
}

var fileDescriptor_c5ef206018e780c1 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xb6, 0x9d, 0x53, 0x96, 0x21, 0x2b, 0x15, 0x0f, 0x0e, 0x77, 0xd8, 0x1e, 0x20, 0x85, 0x7a,
	0x8b, 0x5e, 0x52, 0x1b, 0x47, 0x71, 0xd6, 0xb2, 0x3f, 0x55, 0xa4, 0x50, 0xa2, 0x29, 0xa5, 0xb0,
	0x26, 0x73, 0xc9, 0xf6, 0x18, 0x3e, 0x84, 0x47, 0x1f, 0xc5, 0x47, 0xf1, 0xe2, 0x2b, 0x48, 0xd3,
	0x6d, 0xa7, 0xe9, 0x25, 0x7c, 0xe4, 0xfb, 0xc3, 0xef, 0xfb, 0xc0, 0x75, 0x2e, 0x44, 0x3e, 0xcf,
	0x1c, 0xca, 0xa4, 0x53, 0xc3, 0x0a, 0xad, 0x5d, 0x27, 0xe3, 0xab, 0x52, 0x3a, 0x8c, 0x2a, 0x9a,
	0xb2, 0x65, 0xb1, 0xce, 0x78, 0x5a, 0x0a, 0x96, 0xcd, 0x53, 0xa9, 0xa8, 0x5a, 0x49, 0xb8, 0x58,
	0x0a, 0x25, 0xec, 0x5e, 0x6d, 0x81, 0x94, 0x49, 0xb8, 0x73, 0xc3, 0xb5, 0x0b, 0xb5, 0xbb, 0x7b,
	0xb1, 0x0d, 0x5f, 0x14, 0x0e, 0xe5, 0x5c, 0x28, 0xaa, 0x0a, 0xc1, 0x37, 0xe6, 0xc1, 0xbb, 0x01,
	0xce, 0x7d, 0xaa, 0xa8, 0xaf, 0xe3, 0xef, 0xab, 0xf4, 0x89, 0x0e, 0x27, 0x7c, 0x55, 0x0e, 0xde,
	0xc0, 0xd9, 0x5e, 0xd2, 0xee, 0x80, 0xf6, 0x2c, 0x9c, 0x44, 0xe4, 0x26, 0xb8, 0x0d, 0x88, 0x6f,
	0x1d, 0xd8, 0x6d, 0x70, 0x3c, 0x0b, 0xef, 0xc2, 0x87, 0xc7, 0xd0, 0x32, 0xec, 0x13, 0xd0, 0xc2,
	0x31, 0x0e, 0x46, 0xd8, 0x1b, 0x11, 0xcb, 0xb4, 0x5b, 0xa0, 0x39, 0x99, 0xe2, 0x11, 0xb1, 0x1a,
	0x95, 0x8c, 0x3c, 0x45, 0xc1, 0x98, 0xf8, 0xd6, 0xa1, 0x7d, 0x0a, 0x3a, 0x21, 0x89, 0xc9, 0x38,
	0x1d, 0x92, 0x90, 0x8c, 0xf1, 0x94, 0xf8, 0x56, 0xd3, 0xfb, 0x31, 0x40, 0xff, 0x55, 0x94, 0xf0,
	0xdf, 0x52, 0x5e, 0x77, 0xef, 0x59, 0x51, 0x55, 0x29, 0x32, 0x9e, 0xbd, 0x8d, 0x39, 0x17, 0x73,
	0xca, 0x73, 0x28, 0x96, 0xb9, 0x93, 0x67, 0x5c, 0x17, 0xde, 0xee, 0xbb, 0x28, 0xe4, 0x1f, 0x73,
	0x5f, 0xe9, 0xf7, 0xc3, 0x6c, 0x0c, 0x31, 0xfe, 0x34, 0x7b, 0xc3, 0x3a, 0x0a, 0x33, 0x09, 0x6b,
	0x58, 0xa1, 0xd8, 0x85, 0xd5, 0x3e, 0xf2, 0x6b, 0xcb, 0x27, 0x98, 0xc9, 0x64, 0xc7, 0x27, 0xb1,
	0x9b, 0x68, 0xfe, 0xdb, 0xec, 0xd7, 0x9f, 0x08, 0x61, 0x26, 0x11, 0xda, 0x29, 0x10, 0x8a, 0x5d,
	0x84, 0xb4, 0xe6, 0xe5, 0x48, 0x1f, 0x76, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x97, 0xe3, 0xdc,
	0xa5, 0x06, 0x02, 0x00, 0x00,
}
