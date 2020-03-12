// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/spanner/v1/type.proto

package spanner

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

// `TypeCode` is used as part of [Type][google.spanner.v1.Type] to
// indicate the type of a Cloud Spanner value.
//
// Each legal value of a type can be encoded to or decoded from a JSON
// value, using the encodings described below. All Cloud Spanner values can
// be `null`, regardless of type; `null`s are always encoded as a JSON
// `null`.
type TypeCode int32

const (
	// Not specified.
	TypeCode_TYPE_CODE_UNSPECIFIED TypeCode = 0
	// Encoded as JSON `true` or `false`.
	TypeCode_BOOL TypeCode = 1
	// Encoded as `string`, in decimal format.
	TypeCode_INT64 TypeCode = 2
	// Encoded as `number`, or the strings `"NaN"`, `"Infinity"`, or
	// `"-Infinity"`.
	TypeCode_FLOAT64 TypeCode = 3
	// Encoded as `string` in RFC 3339 timestamp format. The time zone
	// must be present, and must be `"Z"`.
	//
	// If the schema has the column option
	// `allow_commit_timestamp=true`, the placeholder string
	// `"spanner.commit_timestamp()"` can be used to instruct the system
	// to insert the commit timestamp associated with the transaction
	// commit.
	TypeCode_TIMESTAMP TypeCode = 4
	// Encoded as `string` in RFC 3339 date format.
	TypeCode_DATE TypeCode = 5
	// Encoded as `string`.
	TypeCode_STRING TypeCode = 6
	// Encoded as a base64-encoded `string`, as described in RFC 4648,
	// section 4.
	TypeCode_BYTES TypeCode = 7
	// Encoded as `list`, where the list elements are represented
	// according to
	// [array_element_type][google.spanner.v1.Type.array_element_type].
	TypeCode_ARRAY TypeCode = 8
	// Encoded as `list`, where list element `i` is represented according
	// to [struct_type.fields[i]][google.spanner.v1.StructType.fields].
	TypeCode_STRUCT TypeCode = 9
)

var TypeCode_name = map[int32]string{
	0: "TYPE_CODE_UNSPECIFIED",
	1: "BOOL",
	2: "INT64",
	3: "FLOAT64",
	4: "TIMESTAMP",
	5: "DATE",
	6: "STRING",
	7: "BYTES",
	8: "ARRAY",
	9: "STRUCT",
}

var TypeCode_value = map[string]int32{
	"TYPE_CODE_UNSPECIFIED": 0,
	"BOOL":                  1,
	"INT64":                 2,
	"FLOAT64":               3,
	"TIMESTAMP":             4,
	"DATE":                  5,
	"STRING":                6,
	"BYTES":                 7,
	"ARRAY":                 8,
	"STRUCT":                9,
}

func (x TypeCode) String() string {
	return proto.EnumName(TypeCode_name, int32(x))
}

func (TypeCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dc1f2442a7aeba2a, []int{0}
}

// `Type` indicates the type of a Cloud Spanner value, as might be stored in a
// table cell or returned from an SQL query.
type Type struct {
	// Required. The [TypeCode][google.spanner.v1.TypeCode] for this type.
	Code TypeCode `protobuf:"varint,1,opt,name=code,proto3,enum=google.spanner.v1.TypeCode" json:"code,omitempty"`
	// If [code][google.spanner.v1.Type.code] == [ARRAY][google.spanner.v1.TypeCode.ARRAY], then `array_element_type`
	// is the type of the array elements.
	ArrayElementType *Type `protobuf:"bytes,2,opt,name=array_element_type,json=arrayElementType,proto3" json:"array_element_type,omitempty"`
	// If [code][google.spanner.v1.Type.code] == [STRUCT][google.spanner.v1.TypeCode.STRUCT], then `struct_type`
	// provides type information for the struct's fields.
	StructType           *StructType `protobuf:"bytes,3,opt,name=struct_type,json=structType,proto3" json:"struct_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1f2442a7aeba2a, []int{0}
}

func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

func (m *Type) GetCode() TypeCode {
	if m != nil {
		return m.Code
	}
	return TypeCode_TYPE_CODE_UNSPECIFIED
}

func (m *Type) GetArrayElementType() *Type {
	if m != nil {
		return m.ArrayElementType
	}
	return nil
}

func (m *Type) GetStructType() *StructType {
	if m != nil {
		return m.StructType
	}
	return nil
}

// `StructType` defines the fields of a [STRUCT][google.spanner.v1.TypeCode.STRUCT] type.
type StructType struct {
	// The list of fields that make up this struct. Order is
	// significant, because values of this struct type are represented as
	// lists, where the order of field values matches the order of
	// fields in the [StructType][google.spanner.v1.StructType]. In turn, the order of fields
	// matches the order of columns in a read request, or the order of
	// fields in the `SELECT` clause of a query.
	Fields               []*StructType_Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StructType) Reset()         { *m = StructType{} }
func (m *StructType) String() string { return proto.CompactTextString(m) }
func (*StructType) ProtoMessage()    {}
func (*StructType) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1f2442a7aeba2a, []int{1}
}

func (m *StructType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructType.Unmarshal(m, b)
}
func (m *StructType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructType.Marshal(b, m, deterministic)
}
func (m *StructType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructType.Merge(m, src)
}
func (m *StructType) XXX_Size() int {
	return xxx_messageInfo_StructType.Size(m)
}
func (m *StructType) XXX_DiscardUnknown() {
	xxx_messageInfo_StructType.DiscardUnknown(m)
}

var xxx_messageInfo_StructType proto.InternalMessageInfo

func (m *StructType) GetFields() []*StructType_Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

// Message representing a single field of a struct.
type StructType_Field struct {
	// The name of the field. For reads, this is the column name. For
	// SQL queries, it is the column alias (e.g., `"Word"` in the
	// query `"SELECT 'hello' AS Word"`), or the column name (e.g.,
	// `"ColName"` in the query `"SELECT ColName FROM Table"`). Some
	// columns might have an empty name (e.g., !"SELECT
	// UPPER(ColName)"`). Note that a query result can contain
	// multiple fields with the same name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the field.
	Type                 *Type    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StructType_Field) Reset()         { *m = StructType_Field{} }
func (m *StructType_Field) String() string { return proto.CompactTextString(m) }
func (*StructType_Field) ProtoMessage()    {}
func (*StructType_Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1f2442a7aeba2a, []int{1, 0}
}

func (m *StructType_Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructType_Field.Unmarshal(m, b)
}
func (m *StructType_Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructType_Field.Marshal(b, m, deterministic)
}
func (m *StructType_Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructType_Field.Merge(m, src)
}
func (m *StructType_Field) XXX_Size() int {
	return xxx_messageInfo_StructType_Field.Size(m)
}
func (m *StructType_Field) XXX_DiscardUnknown() {
	xxx_messageInfo_StructType_Field.DiscardUnknown(m)
}

var xxx_messageInfo_StructType_Field proto.InternalMessageInfo

func (m *StructType_Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StructType_Field) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.spanner.v1.TypeCode", TypeCode_name, TypeCode_value)
	proto.RegisterType((*Type)(nil), "google.spanner.v1.Type")
	proto.RegisterType((*StructType)(nil), "google.spanner.v1.StructType")
	proto.RegisterType((*StructType_Field)(nil), "google.spanner.v1.StructType.Field")
}

func init() {
	proto.RegisterFile("google/spanner/v1/type.proto", fileDescriptor_dc1f2442a7aeba2a)
}

var fileDescriptor_dc1f2442a7aeba2a = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x9d, 0x6d, 0xda, 0x6d, 0x4e, 0x51, 0xc6, 0x81, 0x65, 0xeb, 0xaa, 0x50, 0xd6, 0x9b,
	0xa2, 0x90, 0xd0, 0x2a, 0x22, 0x2c, 0x08, 0x69, 0x3a, 0x5d, 0x03, 0xbb, 0x6d, 0x48, 0x66, 0x17,
	0x2a, 0x85, 0x32, 0xb6, 0x63, 0x28, 0xa4, 0x33, 0x21, 0xc9, 0x2e, 0xf4, 0x25, 0xbc, 0xd0, 0xb7,
	0xf0, 0x21, 0x7c, 0x00, 0x9f, 0x4a, 0x66, 0x92, 0xaa, 0xb0, 0x2a, 0xde, 0x9d, 0xe4, 0xfb, 0xbf,
	0x33, 0x67, 0x86, 0x03, 0x4f, 0x12, 0xa5, 0x92, 0x54, 0xb8, 0x45, 0xc6, 0xa5, 0x14, 0xb9, 0x7b,
	0x3b, 0x70, 0xcb, 0x5d, 0x26, 0x9c, 0x2c, 0x57, 0xa5, 0x22, 0x0f, 0x2b, 0xea, 0xd4, 0xd4, 0xb9,
	0x1d, 0x9c, 0xec, 0x05, 0x9e, 0x6d, 0x5c, 0x2e, 0xa5, 0x2a, 0x79, 0xb9, 0x51, 0xb2, 0xa8, 0x84,
	0xd3, 0x6f, 0x08, 0x2c, 0xb6, 0xcb, 0x04, 0x71, 0xc1, 0x5a, 0xa9, 0xb5, 0xe8, 0xa2, 0x1e, 0xea,
	0x3f, 0x18, 0x3e, 0x76, 0xee, 0x34, 0x72, 0x74, 0xcc, 0x57, 0x6b, 0x11, 0x99, 0x20, 0xa1, 0x40,
	0x78, 0x9e, 0xf3, 0xdd, 0x52, 0xa4, 0x62, 0x2b, 0x64, 0xb9, 0xd4, 0x63, 0x74, 0x0f, 0x7a, 0xa8,
	0xdf, 0x19, 0x1e, 0xff, 0x45, 0x8f, 0xb0, 0x51, 0x68, 0x65, 0x98, 0x73, 0xdf, 0x42, 0xa7, 0x28,
	0xf3, 0x9b, 0x55, 0xed, 0x37, 0x8c, 0xff, 0xf4, 0x0f, 0x7e, 0x6c, 0x52, 0xa6, 0x0b, 0x14, 0x3f,
	0xeb, 0xd3, 0x2f, 0x08, 0xe0, 0x17, 0x22, 0x67, 0xd0, 0xfa, 0xb8, 0x11, 0xe9, 0xba, 0xe8, 0xa2,
	0x5e, 0xa3, 0xdf, 0x19, 0x3e, 0xfb, 0x67, 0x27, 0x67, 0xa2, 0xb3, 0x51, 0xad, 0x9c, 0xbc, 0x83,
	0xa6, 0xf9, 0x41, 0x08, 0x58, 0x92, 0x6f, 0xab, 0xc7, 0xb0, 0x23, 0x53, 0x93, 0x17, 0x60, 0xfd,
	0xcf, 0x0d, 0x4d, 0xe8, 0xf9, 0x27, 0x04, 0xed, 0xfd, 0x7b, 0x91, 0x47, 0x70, 0xc4, 0xe6, 0x21,
	0x5d, 0xfa, 0xb3, 0x31, 0x5d, 0x5e, 0x4d, 0xe3, 0x90, 0xfa, 0xc1, 0x24, 0xa0, 0x63, 0x7c, 0x8f,
	0xb4, 0xc1, 0x1a, 0xcd, 0x66, 0x17, 0x18, 0x11, 0x1b, 0x9a, 0xc1, 0x94, 0xbd, 0x7e, 0x85, 0x0f,
	0x48, 0x07, 0x0e, 0x27, 0x17, 0x33, 0x4f, 0x7f, 0x34, 0xc8, 0x7d, 0xb0, 0x59, 0x70, 0x49, 0x63,
	0xe6, 0x5d, 0x86, 0xd8, 0xd2, 0xc2, 0xd8, 0x63, 0x14, 0x37, 0x09, 0x40, 0x2b, 0x66, 0x51, 0x30,
	0x3d, 0xc7, 0x2d, 0x2d, 0x8f, 0xe6, 0x8c, 0xc6, 0xf8, 0x50, 0x97, 0x5e, 0x14, 0x79, 0x73, 0xdc,
	0xae, 0x13, 0x57, 0x3e, 0xc3, 0xf6, 0xe8, 0x33, 0x82, 0xa3, 0x95, 0xda, 0xde, 0x9d, 0x7a, 0x64,
	0xeb, 0x39, 0x43, 0xbd, 0x0c, 0x21, 0x7a, 0xff, 0xa6, 0xe6, 0x89, 0x4a, 0xb9, 0x4c, 0x1c, 0x95,
	0x27, 0x6e, 0x22, 0xa4, 0x59, 0x15, 0xb7, 0x42, 0x3c, 0xdb, 0x14, 0xbf, 0x2d, 0xdf, 0x59, 0x5d,
	0x7e, 0x3d, 0x38, 0x3e, 0xaf, 0x54, 0x3f, 0x55, 0x37, 0x6b, 0x27, 0xae, 0x0f, 0xb8, 0x1e, 0x7c,
	0xdf, 0x93, 0x85, 0x21, 0x8b, 0x9a, 0x2c, 0xae, 0x07, 0x1f, 0x5a, 0xa6, 0xf1, 0xcb, 0x1f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x55, 0xc4, 0x6e, 0xd4, 0xd4, 0x02, 0x00, 0x00,
}
