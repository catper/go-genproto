// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/new_resource_creation_error.proto

package errors

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

// Enum describing possible new resource creation errors.
type NewResourceCreationErrorEnum_NewResourceCreationError int32

const (
	// Enum unspecified.
	NewResourceCreationErrorEnum_UNSPECIFIED NewResourceCreationErrorEnum_NewResourceCreationError = 0
	// The received error code is not known in this version.
	NewResourceCreationErrorEnum_UNKNOWN NewResourceCreationErrorEnum_NewResourceCreationError = 1
	// Do not set the id field while creating new resources.
	NewResourceCreationErrorEnum_CANNOT_SET_ID_FOR_CREATE NewResourceCreationErrorEnum_NewResourceCreationError = 2
	// Creating more than one resource with the same temp ID is not allowed.
	NewResourceCreationErrorEnum_DUPLICATE_TEMP_IDS NewResourceCreationErrorEnum_NewResourceCreationError = 3
	// Parent resource with specified temp ID failed validation, so no
	// validation will be done for this child resource.
	NewResourceCreationErrorEnum_TEMP_ID_RESOURCE_HAD_ERRORS NewResourceCreationErrorEnum_NewResourceCreationError = 4
)

var NewResourceCreationErrorEnum_NewResourceCreationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CANNOT_SET_ID_FOR_CREATE",
	3: "DUPLICATE_TEMP_IDS",
	4: "TEMP_ID_RESOURCE_HAD_ERRORS",
}

var NewResourceCreationErrorEnum_NewResourceCreationError_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"CANNOT_SET_ID_FOR_CREATE":    2,
	"DUPLICATE_TEMP_IDS":          3,
	"TEMP_ID_RESOURCE_HAD_ERRORS": 4,
}

func (x NewResourceCreationErrorEnum_NewResourceCreationError) String() string {
	return proto.EnumName(NewResourceCreationErrorEnum_NewResourceCreationError_name, int32(x))
}

func (NewResourceCreationErrorEnum_NewResourceCreationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d40957e6c0c933f5, []int{0, 0}
}

// Container for enum describing possible new resource creation errors.
type NewResourceCreationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewResourceCreationErrorEnum) Reset()         { *m = NewResourceCreationErrorEnum{} }
func (m *NewResourceCreationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NewResourceCreationErrorEnum) ProtoMessage()    {}
func (*NewResourceCreationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d40957e6c0c933f5, []int{0}
}

func (m *NewResourceCreationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Unmarshal(m, b)
}
func (m *NewResourceCreationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Marshal(b, m, deterministic)
}
func (m *NewResourceCreationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResourceCreationErrorEnum.Merge(m, src)
}
func (m *NewResourceCreationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Size(m)
}
func (m *NewResourceCreationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResourceCreationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NewResourceCreationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.NewResourceCreationErrorEnum_NewResourceCreationError", NewResourceCreationErrorEnum_NewResourceCreationError_name, NewResourceCreationErrorEnum_NewResourceCreationError_value)
	proto.RegisterType((*NewResourceCreationErrorEnum)(nil), "google.ads.googleads.v1.errors.NewResourceCreationErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/new_resource_creation_error.proto", fileDescriptor_d40957e6c0c933f5)
}

var fileDescriptor_d40957e6c0c933f5 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x6e, 0xa3, 0x30,
	0x14, 0xc6, 0x07, 0x32, 0x9a, 0x91, 0x9c, 0xc5, 0x20, 0x16, 0xa3, 0x68, 0x26, 0x93, 0x91, 0x38,
	0x80, 0x11, 0x9a, 0x9d, 0x67, 0x33, 0x0e, 0x38, 0x19, 0xd4, 0x16, 0x90, 0x81, 0x54, 0xaa, 0x90,
	0x2c, 0x1a, 0x2c, 0x14, 0x29, 0xc1, 0x91, 0x9d, 0x3f, 0xc7, 0xe8, 0x1d, 0xba, 0xec, 0xaa, 0xe7,
	0xe8, 0x51, 0x7a, 0x82, 0x2e, 0x2b, 0x30, 0xc9, 0x2e, 0x5d, 0xf1, 0x61, 0xff, 0xde, 0xf7, 0xbd,
	0xe7, 0x07, 0xfe, 0xd5, 0x42, 0xd4, 0x6b, 0xee, 0x96, 0x95, 0x72, 0xb5, 0x6c, 0xd5, 0xc1, 0x73,
	0xb9, 0x94, 0x42, 0x2a, 0xb7, 0xe1, 0x47, 0x26, 0xb9, 0x12, 0x7b, 0xb9, 0xe4, 0x6c, 0x29, 0x79,
	0xb9, 0x5b, 0x89, 0x86, 0x75, 0x97, 0x70, 0x2b, 0xc5, 0x4e, 0xd8, 0x13, 0x5d, 0x06, 0xcb, 0x4a,
	0xc1, 0xb3, 0x03, 0x3c, 0x78, 0x50, 0x3b, 0xfc, 0x18, 0x9f, 0x12, 0xb6, 0x2b, 0xb7, 0x6c, 0x1a,
	0xb1, 0xeb, 0x2c, 0x94, 0xae, 0x76, 0x9e, 0x0d, 0x30, 0x8e, 0xf8, 0x91, 0xf6, 0x11, 0x7e, 0x9f,
	0x40, 0xda, 0x5a, 0xd2, 0xec, 0x37, 0xce, 0x83, 0x01, 0x46, 0x97, 0x00, 0xfb, 0x1b, 0x18, 0xe6,
	0x51, 0x9a, 0x10, 0x3f, 0x9c, 0x85, 0x24, 0xb0, 0x3e, 0xd9, 0x43, 0xf0, 0x35, 0x8f, 0xae, 0xa2,
	0xf8, 0x36, 0xb2, 0x0c, 0x7b, 0x0c, 0x46, 0x3e, 0x8e, 0xa2, 0x38, 0x63, 0x29, 0xc9, 0x58, 0x18,
	0xb0, 0x59, 0x4c, 0x99, 0x4f, 0x09, 0xce, 0x88, 0x65, 0xda, 0xdf, 0x81, 0x1d, 0xe4, 0xc9, 0x75,
	0xe8, 0xe3, 0x8c, 0xb0, 0x8c, 0xdc, 0x24, 0x2c, 0x0c, 0x52, 0x6b, 0x60, 0xff, 0x06, 0x3f, 0xfb,
	0x3f, 0x46, 0x49, 0x1a, 0xe7, 0xd4, 0x27, 0xec, 0x3f, 0x0e, 0x18, 0xa1, 0x34, 0xa6, 0xa9, 0xf5,
	0x79, 0xfa, 0x66, 0x00, 0x67, 0x29, 0x36, 0xf0, 0xe3, 0xb9, 0xa7, 0xbf, 0x2e, 0x75, 0x9d, 0xb4,
	0x83, 0x27, 0xc6, 0x5d, 0xd0, 0x1b, 0xd4, 0x62, 0x5d, 0x36, 0x35, 0x14, 0xb2, 0x76, 0x6b, 0xde,
	0x74, 0xcf, 0x72, 0x5a, 0xc5, 0x76, 0xa5, 0x2e, 0x6d, 0xe6, 0xaf, 0xfe, 0x3c, 0x9a, 0x83, 0x39,
	0xc6, 0x4f, 0xe6, 0x64, 0xae, 0xcd, 0x70, 0xa5, 0xa0, 0x96, 0xad, 0x5a, 0x78, 0xb0, 0x8b, 0x54,
	0x2f, 0x27, 0xa0, 0xc0, 0x95, 0x2a, 0xce, 0x40, 0xb1, 0xf0, 0x0a, 0x0d, 0xbc, 0x9a, 0x8e, 0x3e,
	0x45, 0x08, 0x57, 0x0a, 0xa1, 0x33, 0x82, 0xd0, 0xc2, 0x43, 0x48, 0x43, 0xf7, 0x5f, 0xba, 0xee,
	0xfe, 0xbc, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x68, 0x8c, 0x1b, 0x36, 0x02, 0x00, 0x00,
}
