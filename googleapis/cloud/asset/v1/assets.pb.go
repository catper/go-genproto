// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/asset/v1/assets.proto

package asset

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	_ "github.com/catper/protobuf/ptypes/any"
	_struct "github.com/catper/protobuf/ptypes/struct"
	timestamp "github.com/catper/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1 "google.golang.org/genproto/googleapis/iam/v1"
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

// Temporal asset. In addition to the asset, the temporal asset includes the
// status of the asset and valid from and to time of it.
type TemporalAsset struct {
	// The time window when the asset data and state was observed.
	Window *TimeWindow `protobuf:"bytes,1,opt,name=window,proto3" json:"window,omitempty"`
	// If the asset is deleted or not.
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Asset.
	Asset                *Asset   `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemporalAsset) Reset()         { *m = TemporalAsset{} }
func (m *TemporalAsset) String() string { return proto.CompactTextString(m) }
func (*TemporalAsset) ProtoMessage()    {}
func (*TemporalAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_e71186cd3ee2fb90, []int{0}
}

func (m *TemporalAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemporalAsset.Unmarshal(m, b)
}
func (m *TemporalAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemporalAsset.Marshal(b, m, deterministic)
}
func (m *TemporalAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemporalAsset.Merge(m, src)
}
func (m *TemporalAsset) XXX_Size() int {
	return xxx_messageInfo_TemporalAsset.Size(m)
}
func (m *TemporalAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_TemporalAsset.DiscardUnknown(m)
}

var xxx_messageInfo_TemporalAsset proto.InternalMessageInfo

func (m *TemporalAsset) GetWindow() *TimeWindow {
	if m != nil {
		return m.Window
	}
	return nil
}

func (m *TemporalAsset) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *TemporalAsset) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

// A time window of (start_time, end_time].
type TimeWindow struct {
	// Start time of the time window (exclusive).
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time of the time window (inclusive).
	// Current timestamp if not specified.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeWindow) Reset()         { *m = TimeWindow{} }
func (m *TimeWindow) String() string { return proto.CompactTextString(m) }
func (*TimeWindow) ProtoMessage()    {}
func (*TimeWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_e71186cd3ee2fb90, []int{1}
}

func (m *TimeWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeWindow.Unmarshal(m, b)
}
func (m *TimeWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeWindow.Marshal(b, m, deterministic)
}
func (m *TimeWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeWindow.Merge(m, src)
}
func (m *TimeWindow) XXX_Size() int {
	return xxx_messageInfo_TimeWindow.Size(m)
}
func (m *TimeWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeWindow.DiscardUnknown(m)
}

var xxx_messageInfo_TimeWindow proto.InternalMessageInfo

func (m *TimeWindow) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeWindow) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// Cloud asset. This includes all Google Cloud Platform resources,
// Cloud IAM policies, and other non-GCP assets.
type Asset struct {
	// The full name of the asset. For example:
	// `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`.
	// See [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more information.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the asset. Example: "compute.googleapis.com/Disk".
	AssetType string `protobuf:"bytes,2,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	// Representation of the resource.
	Resource *Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// Representation of the actual Cloud IAM policy set on a cloud resource. For
	// each resource, there must be at most one Cloud IAM policy set on it.
	IamPolicy *v1.Policy `protobuf:"bytes,4,opt,name=iam_policy,json=iamPolicy,proto3" json:"iam_policy,omitempty"`
	// Asset's ancestry path in Cloud Resource Manager (CRM) hierarchy,
	// represented as a list of relative resource names. Ancestry path starts with
	// the closest CRM ancestor and ends at root. If the asset is a CRM
	// project/folder/organization, this starts from the asset itself.
	//
	// Example: ["projects/123456789", "folders/5432", "organizations/1234"]
	Ancestors            []string `protobuf:"bytes,10,rep,name=ancestors,proto3" json:"ancestors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_e71186cd3ee2fb90, []int{2}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *Asset) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Asset) GetIamPolicy() *v1.Policy {
	if m != nil {
		return m.IamPolicy
	}
	return nil
}

func (m *Asset) GetAncestors() []string {
	if m != nil {
		return m.Ancestors
	}
	return nil
}

// Representation of a cloud resource.
type Resource struct {
	// The API version. Example: "v1".
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The URL of the discovery document containing the resource's JSON schema.
	// For example:
	// `"https://www.googleapis.com/discovery/v1/apis/compute/v1/rest"`.
	// It will be left unspecified for resources without a discovery-based API,
	// such as Cloud Bigtable.
	DiscoveryDocumentUri string `protobuf:"bytes,2,opt,name=discovery_document_uri,json=discoveryDocumentUri,proto3" json:"discovery_document_uri,omitempty"`
	// The JSON schema name listed in the discovery document.
	// Example: "Project". It will be left unspecified for resources (such as
	// Cloud Bigtable) without a discovery-based API.
	DiscoveryName string `protobuf:"bytes,3,opt,name=discovery_name,json=discoveryName,proto3" json:"discovery_name,omitempty"`
	// The REST URL for accessing the resource. An HTTP GET operation using this
	// URL returns the resource itself.
	// Example:
	// `https://cloudresourcemanager.googleapis.com/v1/projects/my-project-123`.
	// It will be left unspecified for resources without a REST API.
	ResourceUrl string `protobuf:"bytes,4,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// The full name of the immediate parent of this resource. See
	// [Resource
	// Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	// for more information.
	//
	// For GCP assets, it is the parent resource defined in the [Cloud IAM policy
	// hierarchy](https://cloud.google.com/iam/docs/overview#policy_hierarchy).
	// For example:
	// `"//cloudresourcemanager.googleapis.com/projects/my_project_123"`.
	//
	// For third-party assets, it is up to the users to define.
	Parent string `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	// The content of the resource, in which some sensitive fields are scrubbed
	// away and may not be present.
	Data                 *_struct.Struct `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_e71186cd3ee2fb90, []int{3}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Resource) GetDiscoveryDocumentUri() string {
	if m != nil {
		return m.DiscoveryDocumentUri
	}
	return ""
}

func (m *Resource) GetDiscoveryName() string {
	if m != nil {
		return m.DiscoveryName
	}
	return ""
}

func (m *Resource) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *Resource) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Resource) GetData() *_struct.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TemporalAsset)(nil), "google.cloud.asset.v1.TemporalAsset")
	proto.RegisterType((*TimeWindow)(nil), "google.cloud.asset.v1.TimeWindow")
	proto.RegisterType((*Asset)(nil), "google.cloud.asset.v1.Asset")
	proto.RegisterType((*Resource)(nil), "google.cloud.asset.v1.Resource")
}

func init() {
	proto.RegisterFile("google/cloud/asset/v1/assets.proto", fileDescriptor_e71186cd3ee2fb90)
}

var fileDescriptor_e71186cd3ee2fb90 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0xfa, 0xb1, 0xdd, 0x4c, 0x29, 0x07, 0x8b, 0x96, 0x74, 0x55, 0xd4, 0x76, 0x05, 0xa2,
	0x02, 0x29, 0xd1, 0x96, 0x72, 0x68, 0x7b, 0x6a, 0x41, 0xe2, 0x86, 0xaa, 0xd0, 0x16, 0x09, 0xad,
	0x14, 0xb9, 0x89, 0x89, 0x2c, 0x25, 0x76, 0x64, 0x3b, 0x5b, 0xed, 0x85, 0xdf, 0xd1, 0xdf, 0xc0,
	0x4f, 0xe1, 0x9f, 0xc0, 0x99, 0x03, 0x47, 0x94, 0xb1, 0xbd, 0x2b, 0x2d, 0xad, 0xb8, 0x79, 0xe6,
	0xcd, 0x7b, 0x9e, 0xf7, 0xbc, 0x1b, 0x18, 0x96, 0x52, 0x96, 0x15, 0x4b, 0xf2, 0x4a, 0xb6, 0x45,
	0x42, 0xb5, 0x66, 0x26, 0x99, 0x8c, 0xec, 0x41, 0xc7, 0x8d, 0x92, 0x46, 0x92, 0x4d, 0x3b, 0x13,
	0xe3, 0x4c, 0x8c, 0x50, 0x3c, 0x19, 0x0d, 0xb6, 0x1d, 0x95, 0x36, 0x3c, 0x51, 0x4c, 0xcb, 0x56,
	0xe5, 0xcc, 0x32, 0x06, 0x03, 0x07, 0x71, 0x5a, 0x77, 0x6a, 0x8d, 0xac, 0x78, 0x3e, 0x75, 0x98,
	0xa7, 0x61, 0x75, 0xd3, 0x7e, 0x4d, 0xa8, 0xf0, 0xd0, 0xce, 0x22, 0xa4, 0x8d, 0x6a, 0x73, 0xe3,
	0xd0, 0xdd, 0x45, 0xd4, 0xf0, 0x9a, 0x69, 0x43, 0xeb, 0x66, 0x81, 0xde, 0x2d, 0x44, 0x85, 0x90,
	0x86, 0x1a, 0x2e, 0x85, 0x73, 0x31, 0xbc, 0x0b, 0x60, 0xe3, 0x92, 0xd5, 0x8d, 0x54, 0xb4, 0x3a,
	0xeb, 0x3c, 0x90, 0x63, 0xe8, 0xdd, 0x72, 0x51, 0xc8, 0xdb, 0x28, 0xd8, 0x0b, 0x0e, 0xd6, 0x0f,
	0xf7, 0xe3, 0x7b, 0x8d, 0xc6, 0x97, 0xbc, 0x66, 0x9f, 0x71, 0x30, 0x75, 0x04, 0x12, 0xc1, 0x5a,
	0xc1, 0x2a, 0x66, 0x58, 0x11, 0x2d, 0xed, 0x05, 0x07, 0xfd, 0xd4, 0x97, 0xe4, 0x10, 0x56, 0x91,
	0x18, 0x2d, 0xa3, 0xe6, 0xce, 0x03, 0x9a, 0xb8, 0x41, 0x6a, 0x47, 0x87, 0xdf, 0x00, 0xe6, 0x77,
	0x90, 0x63, 0x00, 0x6d, 0xa8, 0x32, 0x59, 0xe7, 0xcf, 0xad, 0x36, 0xf0, 0x32, 0xde, 0x3c, 0x2e,
	0x85, 0xe6, 0xd3, 0x10, 0xa7, 0xbb, 0x9a, 0xbc, 0x85, 0x3e, 0x13, 0x85, 0x25, 0x2e, 0xfd, 0x97,
	0xb8, 0xc6, 0x44, 0xd1, 0x55, 0xc3, 0xdf, 0x01, 0xac, 0xda, 0x48, 0x08, 0xac, 0x08, 0xea, 0x6e,
	0x0d, 0x53, 0x3c, 0x93, 0x67, 0x00, 0xb8, 0x66, 0x66, 0xa6, 0x8d, 0x95, 0x0d, 0xd3, 0x10, 0x3b,
	0x97, 0xd3, 0x86, 0x91, 0x53, 0xe8, 0xfb, 0xd7, 0x77, 0x9e, 0x77, 0x1f, 0xf0, 0x9c, 0xba, 0xb1,
	0x74, 0x46, 0x20, 0x47, 0x00, 0x9c, 0xd6, 0x99, 0xfd, 0x81, 0x44, 0x2b, 0x48, 0xdf, 0xf4, 0x74,
	0x4e, 0xeb, 0x8e, 0x76, 0x81, 0x60, 0x1a, 0x72, 0x5a, 0xdb, 0x23, 0xd9, 0x81, 0x90, 0x8a, 0x9c,
	0x69, 0x23, 0x95, 0x8e, 0x60, 0x6f, 0x19, 0x17, 0xf2, 0x8d, 0x93, 0x97, 0xbf, 0xce, 0x9e, 0xc3,
	0x2e, 0xde, 0x6d, 0xaf, 0xb6, 0x72, 0xb4, 0xe1, 0x3a, 0xce, 0x65, 0x9d, 0x58, 0xa7, 0xc1, 0xab,
	0xe1, 0xcf, 0x00, 0xfa, 0x7e, 0xa7, 0xee, 0x45, 0x27, 0x4c, 0x69, 0x2e, 0x85, 0x33, 0xef, 0x4b,
	0x72, 0x04, 0x5b, 0x05, 0xd7, 0xb9, 0x9c, 0x30, 0x35, 0xcd, 0x0a, 0x99, 0xb7, 0x35, 0x13, 0x26,
	0x6b, 0x15, 0x77, 0x59, 0x3c, 0x99, 0xa1, 0xef, 0x1d, 0x78, 0xa5, 0x38, 0x79, 0x01, 0x8f, 0xe7,
	0x2c, 0xcc, 0x74, 0x19, 0xa7, 0x37, 0x66, 0xdd, 0x8f, 0x5d, 0xb8, 0xfb, 0xf0, 0xc8, 0x87, 0x91,
	0xb5, 0xaa, 0xc2, 0x08, 0xc2, 0x74, 0xdd, 0xf7, 0xae, 0x54, 0x45, 0xb6, 0xa0, 0xd7, 0x50, 0xc5,
	0x84, 0x89, 0x56, 0x11, 0x74, 0x15, 0x79, 0x0d, 0x2b, 0x05, 0x35, 0x34, 0xea, 0x61, 0x6a, 0x4f,
	0xff, 0x79, 0xe8, 0x4f, 0xf8, 0xe7, 0x49, 0x71, 0xe8, 0xfc, 0x2e, 0x80, 0xed, 0x5c, 0xd6, 0xf7,
	0xbf, 0xcc, 0x39, 0x60, 0x26, 0x17, 0x1d, 0xf3, 0x22, 0xf8, 0x72, 0xe2, 0x86, 0x4a, 0x59, 0x51,
	0x51, 0xc6, 0x52, 0x95, 0x49, 0xc9, 0x04, 0xea, 0x26, 0xf3, 0x2c, 0x17, 0x3e, 0x19, 0xa7, 0x78,
	0xf8, 0x13, 0x04, 0xdf, 0x97, 0x36, 0x3f, 0x58, 0xfe, 0x3b, 0xbc, 0x04, 0xa5, 0xe3, 0xeb, 0xd1,
	0x0f, 0xdf, 0x1f, 0x63, 0x7f, 0x8c, 0xfd, 0xf1, 0xf5, 0xe8, 0xa6, 0x87, 0xca, 0x6f, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0x91, 0xff, 0xf4, 0x8b, 0x04, 0x00, 0x00,
}
