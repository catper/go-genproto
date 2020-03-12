// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/app_store.proto

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

// App store type in an app extension.
type AppStoreEnum_AppStore int32

const (
	// Not specified.
	AppStoreEnum_UNSPECIFIED AppStoreEnum_AppStore = 0
	// Used for return value only. Represents value unknown in this version.
	AppStoreEnum_UNKNOWN AppStoreEnum_AppStore = 1
	// Apple iTunes.
	AppStoreEnum_APPLE_ITUNES AppStoreEnum_AppStore = 2
	// Google Play.
	AppStoreEnum_GOOGLE_PLAY AppStoreEnum_AppStore = 3
)

var AppStoreEnum_AppStore_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "APPLE_ITUNES",
	3: "GOOGLE_PLAY",
}

var AppStoreEnum_AppStore_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"APPLE_ITUNES": 2,
	"GOOGLE_PLAY":  3,
}

func (x AppStoreEnum_AppStore) String() string {
	return proto.EnumName(AppStoreEnum_AppStore_name, int32(x))
}

func (AppStoreEnum_AppStore) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa65b602be71fce1, []int{0, 0}
}

// Container for enum describing app store type in an app extension.
type AppStoreEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppStoreEnum) Reset()         { *m = AppStoreEnum{} }
func (m *AppStoreEnum) String() string { return proto.CompactTextString(m) }
func (*AppStoreEnum) ProtoMessage()    {}
func (*AppStoreEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa65b602be71fce1, []int{0}
}

func (m *AppStoreEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppStoreEnum.Unmarshal(m, b)
}
func (m *AppStoreEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppStoreEnum.Marshal(b, m, deterministic)
}
func (m *AppStoreEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppStoreEnum.Merge(m, src)
}
func (m *AppStoreEnum) XXX_Size() int {
	return xxx_messageInfo_AppStoreEnum.Size(m)
}
func (m *AppStoreEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppStoreEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppStoreEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AppStoreEnum_AppStore", AppStoreEnum_AppStore_name, AppStoreEnum_AppStore_value)
	proto.RegisterType((*AppStoreEnum)(nil), "google.ads.googleads.v3.enums.AppStoreEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/app_store.proto", fileDescriptor_fa65b602be71fce1)
}

var fileDescriptor_fa65b602be71fce1 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4f, 0xf3, 0x20,
	0x1c, 0xc6, 0xdf, 0x75, 0xc9, 0xab, 0x61, 0x33, 0x36, 0x3d, 0x1a, 0x77, 0xd8, 0xee, 0xc2, 0x81,
	0x1b, 0x9e, 0x98, 0x62, 0xb3, 0x6c, 0xe9, 0x48, 0xe6, 0x66, 0xd4, 0x26, 0x0b, 0xda, 0x86, 0x2c,
	0x59, 0x81, 0x94, 0x6e, 0x1f, 0xc8, 0xa3, 0x1f, 0xc5, 0x8f, 0xe1, 0xd1, 0x4f, 0x61, 0x00, 0xdb,
	0x9b, 0x5e, 0xc8, 0x03, 0xcf, 0x8f, 0x87, 0x87, 0x3f, 0xb8, 0x92, 0x5a, 0xcb, 0x7d, 0x89, 0x44,
	0x61, 0x51, 0x90, 0x4e, 0x1d, 0x31, 0x2a, 0xd5, 0xa1, 0xb2, 0x48, 0x18, 0xb3, 0xb5, 0x8d, 0xae,
	0x4b, 0x68, 0x6a, 0xdd, 0xe8, 0x64, 0x14, 0x18, 0x28, 0x0a, 0x0b, 0x3b, 0x1c, 0x1e, 0x31, 0xf4,
	0xf8, 0xc5, 0x65, 0x9b, 0x66, 0x76, 0x48, 0x28, 0xa5, 0x1b, 0xd1, 0xec, 0xb4, 0xb2, 0xe1, 0xf2,
	0xe4, 0x19, 0x0c, 0xa9, 0x31, 0x2b, 0x17, 0xc7, 0xd4, 0xa1, 0x9a, 0xcc, 0xc1, 0x69, 0xbb, 0x4f,
	0xce, 0xc1, 0x60, 0x9d, 0xad, 0x38, 0xbb, 0x99, 0xdd, 0xcd, 0xd8, 0x6d, 0xfc, 0x2f, 0x19, 0x80,
	0x93, 0x75, 0x36, 0xcf, 0x96, 0x0f, 0x59, 0xdc, 0x4b, 0x62, 0x30, 0xa4, 0x9c, 0x2f, 0xd8, 0x76,
	0x76, 0xbf, 0xce, 0xd8, 0x2a, 0x8e, 0x1c, 0x9f, 0x2e, 0x97, 0xe9, 0x82, 0x6d, 0xf9, 0x82, 0x3e,
	0xc6, 0xfd, 0xe9, 0x67, 0x0f, 0x8c, 0x5f, 0x75, 0x05, 0xff, 0x2c, 0x38, 0x3d, 0x6b, 0x1f, 0xe4,
	0xae, 0x11, 0xef, 0x3d, 0x4d, 0x7f, 0x78, 0xa9, 0xf7, 0x42, 0x49, 0xa8, 0x6b, 0x89, 0x64, 0xa9,
	0x7c, 0xdf, 0x76, 0x1e, 0x66, 0x67, 0x7f, 0x19, 0xcf, 0xb5, 0x5f, 0xdf, 0xa2, 0x7e, 0x4a, 0xe9,
	0x7b, 0x34, 0x4a, 0x43, 0x14, 0x2d, 0x2c, 0x0c, 0xd2, 0xa9, 0x0d, 0x86, 0xee, 0xb3, 0xf6, 0xa3,
	0xf5, 0x73, 0x5a, 0xd8, 0xbc, 0xf3, 0xf3, 0x0d, 0xce, 0xbd, 0xff, 0x15, 0x8d, 0xc3, 0x21, 0x21,
	0xb4, 0xb0, 0x84, 0x74, 0x04, 0x21, 0x1b, 0x4c, 0x88, 0x67, 0x5e, 0xfe, 0xfb, 0x62, 0xf8, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x25, 0xe8, 0xdf, 0x59, 0xb6, 0x01, 0x00, 0x00,
}
