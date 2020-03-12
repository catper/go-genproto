// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/monitored_resource.proto

package monitoredres

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	_struct "github.com/catper/protobuf/ptypes/struct"
	api "google.golang.org/genproto/googleapis/api"
	label "google.golang.org/genproto/googleapis/api/label"
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

// An object that describes the schema of a [MonitoredResource][google.api.MonitoredResource] object using a
// type name and a set of labels.  For example, the monitored resource
// descriptor for Google Compute Engine VM instances has a type of
// `"gce_instance"` and specifies the use of the labels `"instance_id"` and
// `"zone"` to identify particular VM instances.
//
// Different APIs can support different monitored resource types. APIs generally
// provide a `list` method that returns the monitored resource descriptors used
// by the API.
type MonitoredResourceDescriptor struct {
	// Optional. The resource name of the monitored resource descriptor:
	// `"projects/{project_id}/monitoredResourceDescriptors/{type}"` where
	// {type} is the value of the `type` field in this object and
	// {project_id} is a project ID that provides API-specific context for
	// accessing the type.  APIs that do not use project information can use the
	// resource name format `"monitoredResourceDescriptors/{type}"`.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The monitored resource type. For example, the type
	// `"cloudsql_database"` represents databases in Google Cloud SQL.
	// The maximum length of this value is 256 characters.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. A concise name for the monitored resource type that might be
	// displayed in user interfaces. It should be a Title Cased Noun Phrase,
	// without any article or other determiners. For example,
	// `"Google Cloud SQL Database"`.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. A detailed description of the monitored resource type that might
	// be used in documentation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Required. A set of labels used to describe instances of this monitored
	// resource type. For example, an individual Google Cloud SQL database is
	// identified by values for the labels `"database_id"` and `"zone"`.
	Labels []*label.LabelDescriptor `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	// Optional. The launch stage of the monitored resource definition.
	LaunchStage          api.LaunchStage `protobuf:"varint,7,opt,name=launch_stage,json=launchStage,proto3,enum=google.api.LaunchStage" json:"launch_stage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MonitoredResourceDescriptor) Reset()         { *m = MonitoredResourceDescriptor{} }
func (m *MonitoredResourceDescriptor) String() string { return proto.CompactTextString(m) }
func (*MonitoredResourceDescriptor) ProtoMessage()    {}
func (*MonitoredResourceDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd8bd738b08f2bf, []int{0}
}

func (m *MonitoredResourceDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoredResourceDescriptor.Unmarshal(m, b)
}
func (m *MonitoredResourceDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoredResourceDescriptor.Marshal(b, m, deterministic)
}
func (m *MonitoredResourceDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoredResourceDescriptor.Merge(m, src)
}
func (m *MonitoredResourceDescriptor) XXX_Size() int {
	return xxx_messageInfo_MonitoredResourceDescriptor.Size(m)
}
func (m *MonitoredResourceDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoredResourceDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoredResourceDescriptor proto.InternalMessageInfo

func (m *MonitoredResourceDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetLabels() []*label.LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MonitoredResourceDescriptor) GetLaunchStage() api.LaunchStage {
	if m != nil {
		return m.LaunchStage
	}
	return api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
}

// An object representing a resource that can be used for monitoring, logging,
// billing, or other purposes. Examples include virtual machine instances,
// databases, and storage devices such as disks. The `type` field identifies a
// [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] object that describes the resource's
// schema. Information in the `labels` field identifies the actual resource and
// its attributes according to the schema. For example, a particular Compute
// Engine VM instance could be represented by the following object, because the
// [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] for `"gce_instance"` has labels
// `"instance_id"` and `"zone"`:
//
//     { "type": "gce_instance",
//       "labels": { "instance_id": "12345678901234",
//                   "zone": "us-central1-a" }}
type MonitoredResource struct {
	// Required. The monitored resource type. This field must match
	// the `type` field of a [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] object. For
	// example, the type of a Compute Engine VM instance is `gce_instance`.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Required. Values for all of the labels listed in the associated monitored
	// resource descriptor. For example, Compute Engine VM instances use the
	// labels `"project_id"`, `"instance_id"`, and `"zone"`.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MonitoredResource) Reset()         { *m = MonitoredResource{} }
func (m *MonitoredResource) String() string { return proto.CompactTextString(m) }
func (*MonitoredResource) ProtoMessage()    {}
func (*MonitoredResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd8bd738b08f2bf, []int{1}
}

func (m *MonitoredResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoredResource.Unmarshal(m, b)
}
func (m *MonitoredResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoredResource.Marshal(b, m, deterministic)
}
func (m *MonitoredResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoredResource.Merge(m, src)
}
func (m *MonitoredResource) XXX_Size() int {
	return xxx_messageInfo_MonitoredResource.Size(m)
}
func (m *MonitoredResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoredResource.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoredResource proto.InternalMessageInfo

func (m *MonitoredResource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MonitoredResource) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Auxiliary metadata for a [MonitoredResource][google.api.MonitoredResource] object.
// [MonitoredResource][google.api.MonitoredResource] objects contain the minimum set of information to
// uniquely identify a monitored resource instance. There is some other useful
// auxiliary metadata. Monitoring and Logging use an ingestion
// pipeline to extract metadata for cloud resources of all types, and store
// the metadata in this message.
type MonitoredResourceMetadata struct {
	// Output only. Values for predefined system metadata labels.
	// System labels are a kind of metadata extracted by Google, including
	// "machine_image", "vpc", "subnet_id",
	// "security_group", "name", etc.
	// System label values can be only strings, Boolean values, or a list of
	// strings. For example:
	//
	//     { "name": "my-test-instance",
	//       "security_group": ["a", "b", "c"],
	//       "spot_instance": false }
	SystemLabels *_struct.Struct `protobuf:"bytes,1,opt,name=system_labels,json=systemLabels,proto3" json:"system_labels,omitempty"`
	// Output only. A map of user-defined metadata labels.
	UserLabels           map[string]string `protobuf:"bytes,2,rep,name=user_labels,json=userLabels,proto3" json:"user_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MonitoredResourceMetadata) Reset()         { *m = MonitoredResourceMetadata{} }
func (m *MonitoredResourceMetadata) String() string { return proto.CompactTextString(m) }
func (*MonitoredResourceMetadata) ProtoMessage()    {}
func (*MonitoredResourceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cd8bd738b08f2bf, []int{2}
}

func (m *MonitoredResourceMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoredResourceMetadata.Unmarshal(m, b)
}
func (m *MonitoredResourceMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoredResourceMetadata.Marshal(b, m, deterministic)
}
func (m *MonitoredResourceMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoredResourceMetadata.Merge(m, src)
}
func (m *MonitoredResourceMetadata) XXX_Size() int {
	return xxx_messageInfo_MonitoredResourceMetadata.Size(m)
}
func (m *MonitoredResourceMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoredResourceMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoredResourceMetadata proto.InternalMessageInfo

func (m *MonitoredResourceMetadata) GetSystemLabels() *_struct.Struct {
	if m != nil {
		return m.SystemLabels
	}
	return nil
}

func (m *MonitoredResourceMetadata) GetUserLabels() map[string]string {
	if m != nil {
		return m.UserLabels
	}
	return nil
}

func init() {
	proto.RegisterType((*MonitoredResourceDescriptor)(nil), "google.api.MonitoredResourceDescriptor")
	proto.RegisterType((*MonitoredResource)(nil), "google.api.MonitoredResource")
	proto.RegisterMapType((map[string]string)(nil), "google.api.MonitoredResource.LabelsEntry")
	proto.RegisterType((*MonitoredResourceMetadata)(nil), "google.api.MonitoredResourceMetadata")
	proto.RegisterMapType((map[string]string)(nil), "google.api.MonitoredResourceMetadata.UserLabelsEntry")
}

func init() {
	proto.RegisterFile("google/api/monitored_resource.proto", fileDescriptor_6cd8bd738b08f2bf)
}

var fileDescriptor_6cd8bd738b08f2bf = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0xa5, 0x67, 0x66, 0x57, 0xac, 0x8c, 0xab, 0x36, 0xb2, 0xc6, 0xac, 0x42, 0x1c, 0x2f, 0xe3,
	0x25, 0x81, 0x5d, 0x04, 0x5d, 0xf5, 0xb0, 0xab, 0x22, 0x82, 0x2b, 0x43, 0x16, 0x3d, 0x78, 0x09,
	0x3d, 0x49, 0x1b, 0x83, 0x49, 0x3a, 0x74, 0x77, 0x84, 0xfc, 0x1d, 0xc1, 0xff, 0xe7, 0x51, 0xf0,
	0x22, 0xfd, 0x91, 0x49, 0x62, 0x44, 0xf0, 0x56, 0xf5, 0xde, 0xab, 0xaa, 0xf7, 0xd2, 0x04, 0x1e,
	0x64, 0x8c, 0x65, 0x05, 0x0d, 0x49, 0x9d, 0x87, 0x25, 0xab, 0x72, 0xc9, 0x38, 0x4d, 0x63, 0x4e,
	0x05, 0x6b, 0x78, 0x42, 0x83, 0x9a, 0x33, 0xc9, 0x30, 0x18, 0x51, 0x40, 0xea, 0xdc, 0x3b, 0x1c,
	0x0c, 0x14, 0x64, 0x4b, 0x0b, 0xa3, 0xf1, 0xee, 0x8d, 0xf0, 0xa6, 0x4a, 0x3e, 0xc7, 0x42, 0x92,
	0xcc, 0xae, 0xf0, 0xee, 0x5a, 0x5a, 0x77, 0xdb, 0xe6, 0x53, 0x28, 0x24, 0x6f, 0x12, 0x69, 0xd8,
	0xd5, 0x2f, 0x04, 0x47, 0x17, 0xdd, 0xf5, 0xc8, 0x1e, 0x7f, 0x49, 0x45, 0xc2, 0xf3, 0x5a, 0x32,
	0x8e, 0x31, 0x2c, 0x2a, 0x52, 0x52, 0x77, 0xcf, 0x47, 0xeb, 0xab, 0x91, 0xae, 0x15, 0x26, 0xdb,
	0x9a, 0xba, 0xc8, 0x60, 0xaa, 0xc6, 0xf7, 0x61, 0x99, 0xe6, 0xa2, 0x2e, 0x48, 0x1b, 0x6b, 0xfd,
	0x4c, 0x73, 0x8e, 0xc5, 0xde, 0xa9, 0x31, 0x1f, 0x9c, 0xd4, 0x2e, 0xce, 0x59, 0xe5, 0xce, 0xad,
	0xa2, 0x87, 0xf0, 0x09, 0xec, 0xeb, 0x60, 0xc2, 0x5d, 0xf8, 0xf3, 0xb5, 0x73, 0x7c, 0x14, 0xf4,
	0xf1, 0x83, 0xb7, 0x8a, 0xe9, 0x9d, 0x45, 0x56, 0x8a, 0x4f, 0x61, 0x39, 0x4c, 0xed, 0x5e, 0xf1,
	0xd1, 0xfa, 0xe0, 0xf8, 0xf6, 0x78, 0x54, 0xf1, 0x97, 0x8a, 0x8e, 0x9c, 0xa2, 0x6f, 0x56, 0xdf,
	0x11, 0xdc, 0x9c, 0xa4, 0xff, 0x6b, 0xbe, 0xb3, 0x9d, 0xb5, 0x99, 0xb6, 0xf6, 0x70, 0xb8, 0x7f,
	0xb2, 0xc2, 0x98, 0x15, 0xaf, 0x2a, 0xc9, 0xdb, 0xce, 0xa8, 0xf7, 0x04, 0x9c, 0x01, 0x8c, 0x6f,
	0xc0, 0xfc, 0x0b, 0x6d, 0xed, 0x11, 0x55, 0xe2, 0x5b, 0xb0, 0xf7, 0x95, 0x14, 0x4d, 0xf7, 0xf1,
	0x4c, 0x73, 0x3a, 0x7b, 0x8c, 0x56, 0x3f, 0x10, 0xdc, 0x99, 0x1c, 0xb9, 0xa0, 0x92, 0xa4, 0x44,
	0x12, 0xfc, 0x0c, 0xae, 0x89, 0x56, 0x48, 0x5a, 0xc6, 0xd6, 0xa2, 0xda, 0xe9, 0xf4, 0x9f, 0xa0,
	0x7b, 0xf9, 0xe0, 0x52, 0xbf, 0x7c, 0xb4, 0x34, 0x6a, 0x63, 0x06, 0x7f, 0x00, 0xa7, 0x11, 0x94,
	0xc7, 0xa3, 0x78, 0x8f, 0xfe, 0x19, 0xaf, 0xbb, 0x1c, 0xbc, 0x17, 0x94, 0x0f, 0xa3, 0x42, 0xb3,
	0x03, 0xbc, 0xe7, 0x70, 0xfd, 0x0f, 0xfa, 0x7f, 0x22, 0x9f, 0xb7, 0x70, 0x90, 0xb0, 0x72, 0x60,
	0xe3, 0xfc, 0x70, 0xe2, 0x63, 0xa3, 0x82, 0x6d, 0xd0, 0xc7, 0x17, 0x56, 0x95, 0xb1, 0x82, 0x54,
	0x59, 0xc0, 0x78, 0x16, 0x66, 0xb4, 0xd2, 0xb1, 0x43, 0x43, 0x91, 0x3a, 0x17, 0xe3, 0x3f, 0x8d,
	0x53, 0xf1, 0x74, 0xd8, 0xfc, 0x44, 0xe8, 0xdb, 0x6c, 0xf1, 0xfa, 0x6c, 0xf3, 0x66, 0xbb, 0xaf,
	0x27, 0x4f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xa6, 0xca, 0xf1, 0xa2, 0x03, 0x00, 0x00,
}
