// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/timestamps.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	timestamp "github.com/catper/protobuf/ptypes/timestamp"
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

// Timestamps about this resource according to a particular system.
type SystemTimestamps struct {
	// The creation time of the resource within the given system.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last-modified time of the resource within the given system.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. The expiration time of the resource within the given system.
	// Currently only apllicable to BigQuery resources.
	ExpireTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SystemTimestamps) Reset()         { *m = SystemTimestamps{} }
func (m *SystemTimestamps) String() string { return proto.CompactTextString(m) }
func (*SystemTimestamps) ProtoMessage()    {}
func (*SystemTimestamps) Descriptor() ([]byte, []int) {
	return fileDescriptor_a178faa5e507582b, []int{0}
}

func (m *SystemTimestamps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemTimestamps.Unmarshal(m, b)
}
func (m *SystemTimestamps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemTimestamps.Marshal(b, m, deterministic)
}
func (m *SystemTimestamps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemTimestamps.Merge(m, src)
}
func (m *SystemTimestamps) XXX_Size() int {
	return xxx_messageInfo_SystemTimestamps.Size(m)
}
func (m *SystemTimestamps) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemTimestamps.DiscardUnknown(m)
}

var xxx_messageInfo_SystemTimestamps proto.InternalMessageInfo

func (m *SystemTimestamps) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *SystemTimestamps) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *SystemTimestamps) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func init() {
	proto.RegisterType((*SystemTimestamps)(nil), "google.cloud.datacatalog.v1beta1.SystemTimestamps")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/timestamps.proto", fileDescriptor_a178faa5e507582b)
}

var fileDescriptor_a178faa5e507582b = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd1, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x99, 0x48, 0x0c, 0xce, 0x82, 0x3a, 0xa1, 0x2e, 0x54, 0x88, 0x81, 0xc9, 0x56,
	0x60, 0xec, 0x02, 0x5d, 0x59, 0x50, 0x61, 0x62, 0xa9, 0x2e, 0xc9, 0xd5, 0x58, 0x72, 0x7a, 0x96,
	0xe3, 0x54, 0xe5, 0x25, 0x79, 0x06, 0x1e, 0x83, 0x11, 0xc5, 0x36, 0x26, 0x0b, 0xca, 0x98, 0xcb,
	0xff, 0xdd, 0x6f, 0xd9, 0xbc, 0x52, 0x44, 0xca, 0xa0, 0x6c, 0x0c, 0x0d, 0xad, 0x6c, 0xc1, 0x43,
	0x03, 0x1e, 0x0c, 0x29, 0x79, 0xac, 0x6a, 0xf4, 0x50, 0x49, 0xaf, 0x3b, 0xec, 0x3d, 0x74, 0xb6,
	0x17, 0xd6, 0x91, 0xa7, 0xc5, 0x2a, 0x12, 0x11, 0x88, 0x98, 0x10, 0x91, 0xc8, 0xf2, 0x2a, 0x2d,
	0x05, 0xab, 0xe5, 0x5e, 0xa3, 0x69, 0x77, 0x35, 0xbe, 0xc3, 0x51, 0x93, 0x8b, 0x2b, 0x72, 0x20,
	0x7c, 0xd5, 0xc3, 0xfe, 0xaf, 0x24, 0x06, 0xae, 0x3f, 0x19, 0xbf, 0x78, 0xf9, 0xe8, 0x3d, 0x76,
	0xaf, 0xb9, 0x7e, 0xb1, 0xe6, 0x65, 0xe3, 0x10, 0x3c, 0xee, 0xc6, 0xf8, 0x25, 0x5b, 0xb1, 0xdb,
	0xf2, 0x6e, 0x29, 0xd2, 0x71, 0x7e, 0x77, 0x89, 0x2c, 0xb6, 0x3c, 0xc6, 0xc7, 0xc1, 0x88, 0x07,
	0xdb, 0x66, 0x7c, 0x36, 0x8f, 0x63, 0x3c, 0xe0, 0x07, 0x5e, 0xe2, 0xc9, 0x6a, 0x97, 0x70, 0x31,
	0x87, 0x37, 0xc5, 0xd7, 0x63, 0xb1, 0xe5, 0xd1, 0x8c, 0xd3, 0xcd, 0x89, 0xdf, 0x34, 0xd4, 0x89,
	0xb9, 0xab, 0x7b, 0x66, 0x6f, 0x4f, 0x29, 0xa3, 0xc8, 0xc0, 0x41, 0x09, 0x72, 0x4a, 0x2a, 0x3c,
	0x84, 0x0e, 0x19, 0x7f, 0x81, 0xd5, 0xfd, 0xff, 0x0f, 0xb6, 0x9e, 0xcc, 0xbe, 0x19, 0xab, 0xcf,
	0x03, 0xbd, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x08, 0x58, 0xc7, 0x0c, 0xea, 0x01, 0x00, 0x00,
}
