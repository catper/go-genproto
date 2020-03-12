// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/schema.proto

package datacatalog

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

// Represents a schema (e.g. BigQuery, GoogleSQL, Avro schema).
type Schema struct {
	// Required. Schema of columns. A maximum of 10,000 columns and sub-columns can be
	// specified.
	Columns              []*ColumnSchema `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_834760e0ed3be1cb, []int{0}
}

func (m *Schema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schema.Unmarshal(m, b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return xxx_messageInfo_Schema.Size(m)
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetColumns() []*ColumnSchema {
	if m != nil {
		return m.Columns
	}
	return nil
}

// Representation of a column within a schema. Columns could be nested inside
// other columns.
type ColumnSchema struct {
	// Required. Name of the column.
	Column string `protobuf:"bytes,6,opt,name=column,proto3" json:"column,omitempty"`
	// Required. Type of the column.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. Description of the column. Default value is an empty string.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. A column's mode indicates whether the values in this column are required,
	// nullable, etc. Only `NULLABLE`, `REQUIRED` and `REPEATED` are supported.
	// Default mode is `NULLABLE`.
	Mode string `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	// Optional. Schema of sub-columns. A column can have zero or more sub-columns.
	Subcolumns           []*ColumnSchema `protobuf:"bytes,7,rep,name=subcolumns,proto3" json:"subcolumns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ColumnSchema) Reset()         { *m = ColumnSchema{} }
func (m *ColumnSchema) String() string { return proto.CompactTextString(m) }
func (*ColumnSchema) ProtoMessage()    {}
func (*ColumnSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_834760e0ed3be1cb, []int{1}
}

func (m *ColumnSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnSchema.Unmarshal(m, b)
}
func (m *ColumnSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnSchema.Marshal(b, m, deterministic)
}
func (m *ColumnSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnSchema.Merge(m, src)
}
func (m *ColumnSchema) XXX_Size() int {
	return xxx_messageInfo_ColumnSchema.Size(m)
}
func (m *ColumnSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnSchema proto.InternalMessageInfo

func (m *ColumnSchema) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *ColumnSchema) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ColumnSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ColumnSchema) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ColumnSchema) GetSubcolumns() []*ColumnSchema {
	if m != nil {
		return m.Subcolumns
	}
	return nil
}

func init() {
	proto.RegisterType((*Schema)(nil), "google.cloud.datacatalog.v1beta1.Schema")
	proto.RegisterType((*ColumnSchema)(nil), "google.cloud.datacatalog.v1beta1.ColumnSchema")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/schema.proto", fileDescriptor_834760e0ed3be1cb)
}

var fileDescriptor_834760e0ed3be1cb = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4b, 0xfc, 0x30,
	0x18, 0xc6, 0x49, 0xfb, 0xa7, 0xc7, 0x3f, 0xe7, 0x94, 0xc5, 0xa2, 0x83, 0xe5, 0x50, 0xb8, 0xc5,
	0x84, 0xd3, 0xd1, 0xc9, 0x73, 0x14, 0x41, 0x7a, 0x83, 0xe0, 0x22, 0x6f, 0xd3, 0x98, 0x0b, 0xa4,
	0x7d, 0x4b, 0x9b, 0x1e, 0xfa, 0x51, 0xdd, 0xfc, 0x18, 0x8e, 0x62, 0xd2, 0x72, 0x5d, 0xe4, 0xc0,
	0xf5, 0x79, 0x7e, 0xcf, 0x0f, 0xde, 0x84, 0x5e, 0x6a, 0x44, 0x6d, 0x95, 0x90, 0x16, 0xfb, 0x52,
	0x94, 0xe0, 0x40, 0x82, 0x03, 0x8b, 0x5a, 0xec, 0x56, 0x85, 0x72, 0xb0, 0x12, 0x9d, 0xdc, 0xaa,
	0x0a, 0x78, 0xd3, 0xa2, 0x43, 0x96, 0x05, 0x9c, 0x7b, 0x9c, 0x4f, 0x70, 0x3e, 0xe0, 0x27, 0x67,
	0x83, 0x10, 0x1a, 0x23, 0x5e, 0x8d, 0xb2, 0xe5, 0x4b, 0xa1, 0xb6, 0xb0, 0x33, 0xd8, 0x06, 0xc5,
	0xe2, 0x89, 0x26, 0x1b, 0xaf, 0x64, 0x0f, 0x74, 0x26, 0xd1, 0xf6, 0x55, 0xdd, 0xa5, 0x51, 0x16,
	0x2f, 0xe7, 0x57, 0x9c, 0x1f, 0xd2, 0xf3, 0x3b, 0x3f, 0x08, 0x82, 0x75, 0xfc, 0x79, 0x1b, 0xe5,
	0xa3, 0x63, 0xf1, 0x41, 0xe8, 0xd1, 0xb4, 0x66, 0xa7, 0x34, 0x09, 0x5d, 0x9a, 0x64, 0x64, 0xf9,
	0x3f, 0xe0, 0x43, 0xc4, 0x8e, 0xe9, 0x3f, 0xf7, 0xde, 0xa8, 0x94, 0xec, 0x2b, 0x1f, 0xb0, 0x0b,
	0x3a, 0x2f, 0x55, 0x27, 0x5b, 0xd3, 0x38, 0x83, 0x75, 0x1a, 0x8d, 0x3d, 0xc9, 0xa7, 0xf9, 0xcf,
	0xbe, 0xc2, 0x52, 0xa5, 0xf1, 0xbe, 0xf7, 0x01, 0xdb, 0x50, 0xda, 0xf5, 0xc5, 0x78, 0xd8, 0xec,
	0xaf, 0x87, 0x91, 0x7c, 0xa2, 0x59, 0xbf, 0xd1, 0x73, 0x89, 0xd5, 0x41, 0xcb, 0x23, 0x79, 0xbe,
	0x1f, 0x18, 0x8d, 0x16, 0x6a, 0xcd, 0xb1, 0xd5, 0x42, 0xab, 0xda, 0x3f, 0xbd, 0x08, 0x15, 0x34,
	0xa6, 0xfb, 0xfd, 0xbf, 0x6f, 0x26, 0xd9, 0x17, 0x21, 0x45, 0xe2, 0xa7, 0xd7, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xc1, 0x01, 0x3a, 0xa1, 0x29, 0x02, 0x00, 0x00,
}
