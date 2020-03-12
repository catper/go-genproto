// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/company_service.proto

package talent

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	empty "github.com/catper/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The Request of the CreateCompany method.
type CreateCompanyRequest struct {
	// Required. Resource name of the tenant under which the company is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenant/bar". If tenant id is unspecified, a default tenant
	// is created, for example, "projects/foo".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The company to be created.
	Company              *Company `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCompanyRequest) Reset()         { *m = CreateCompanyRequest{} }
func (m *CreateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCompanyRequest) ProtoMessage()    {}
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{0}
}

func (m *CreateCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompanyRequest.Unmarshal(m, b)
}
func (m *CreateCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompanyRequest.Marshal(b, m, deterministic)
}
func (m *CreateCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompanyRequest.Merge(m, src)
}
func (m *CreateCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCompanyRequest.Size(m)
}
func (m *CreateCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompanyRequest proto.InternalMessageInfo

func (m *CreateCompanyRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateCompanyRequest) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

// Request for getting a company by name.
type GetCompanyRequest struct {
	// Required. The resource name of the company to be retrieved.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/api-test-project/tenants/foo/companies/bar".
	//
	// If tenant id is unspecified, the default tenant is used, for
	// example, "projects/api-test-project/companies/bar".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCompanyRequest) Reset()         { *m = GetCompanyRequest{} }
func (m *GetCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*GetCompanyRequest) ProtoMessage()    {}
func (*GetCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{1}
}

func (m *GetCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCompanyRequest.Unmarshal(m, b)
}
func (m *GetCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCompanyRequest.Marshal(b, m, deterministic)
}
func (m *GetCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCompanyRequest.Merge(m, src)
}
func (m *GetCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_GetCompanyRequest.Size(m)
}
func (m *GetCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCompanyRequest proto.InternalMessageInfo

func (m *GetCompanyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for updating a specified company.
type UpdateCompanyRequest struct {
	// Required. The company resource to replace the current resource in the system.
	Company *Company `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	// Strongly recommended for the best service experience.
	//
	// If [update_mask][google.cloud.talent.v4beta1.UpdateCompanyRequest.update_mask] is provided, only the specified fields in
	// [company][google.cloud.talent.v4beta1.UpdateCompanyRequest.company] are updated. Otherwise all the fields are updated.
	//
	// A field mask to specify the company fields to be updated. Only
	// top level fields of [Company][google.cloud.talent.v4beta1.Company] are supported.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCompanyRequest) Reset()         { *m = UpdateCompanyRequest{} }
func (m *UpdateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCompanyRequest) ProtoMessage()    {}
func (*UpdateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{2}
}

func (m *UpdateCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCompanyRequest.Unmarshal(m, b)
}
func (m *UpdateCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCompanyRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCompanyRequest.Merge(m, src)
}
func (m *UpdateCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCompanyRequest.Size(m)
}
func (m *UpdateCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCompanyRequest proto.InternalMessageInfo

func (m *UpdateCompanyRequest) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *UpdateCompanyRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request to delete a company.
type DeleteCompanyRequest struct {
	// Required. The resource name of the company to be deleted.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/foo/tenants/bar/companies/baz".
	//
	// If tenant id is unspecified, the default tenant is used, for
	// example, "projects/foo/companies/bar".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCompanyRequest) Reset()         { *m = DeleteCompanyRequest{} }
func (m *DeleteCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCompanyRequest) ProtoMessage()    {}
func (*DeleteCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{3}
}

func (m *DeleteCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCompanyRequest.Unmarshal(m, b)
}
func (m *DeleteCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCompanyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCompanyRequest.Merge(m, src)
}
func (m *DeleteCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCompanyRequest.Size(m)
}
func (m *DeleteCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCompanyRequest proto.InternalMessageInfo

func (m *DeleteCompanyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// List companies for which the client has ACL visibility.
type ListCompaniesRequest struct {
	// Required. Resource name of the tenant under which the company is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenant/bar".
	//
	// If tenant id is unspecified, the default tenant will be used, for
	// example, "projects/foo".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The starting indicator from which to return results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The maximum number of companies to be returned, at most 100.
	// Default is 100 if a non-positive number is provided.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Set to true if the companies requested must have open jobs.
	//
	// Defaults to false.
	//
	// If true, at most [page_size][google.cloud.talent.v4beta1.ListCompaniesRequest.page_size] of companies are fetched, among which
	// only those with open jobs are returned.
	RequireOpenJobs      bool     `protobuf:"varint,4,opt,name=require_open_jobs,json=requireOpenJobs,proto3" json:"require_open_jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCompaniesRequest) Reset()         { *m = ListCompaniesRequest{} }
func (m *ListCompaniesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCompaniesRequest) ProtoMessage()    {}
func (*ListCompaniesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{4}
}

func (m *ListCompaniesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCompaniesRequest.Unmarshal(m, b)
}
func (m *ListCompaniesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCompaniesRequest.Marshal(b, m, deterministic)
}
func (m *ListCompaniesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCompaniesRequest.Merge(m, src)
}
func (m *ListCompaniesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCompaniesRequest.Size(m)
}
func (m *ListCompaniesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCompaniesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCompaniesRequest proto.InternalMessageInfo

func (m *ListCompaniesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListCompaniesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCompaniesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCompaniesRequest) GetRequireOpenJobs() bool {
	if m != nil {
		return m.RequireOpenJobs
	}
	return false
}

// The List companies response object.
type ListCompaniesResponse struct {
	// Companies for the current client.
	Companies []*Company `protobuf:"bytes,1,rep,name=companies,proto3" json:"companies,omitempty"`
	// A token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Additional information for the API invocation, such as the request
	// tracking id.
	Metadata             *ResponseMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListCompaniesResponse) Reset()         { *m = ListCompaniesResponse{} }
func (m *ListCompaniesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCompaniesResponse) ProtoMessage()    {}
func (*ListCompaniesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c499e6cf60b8944, []int{5}
}

func (m *ListCompaniesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCompaniesResponse.Unmarshal(m, b)
}
func (m *ListCompaniesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCompaniesResponse.Marshal(b, m, deterministic)
}
func (m *ListCompaniesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCompaniesResponse.Merge(m, src)
}
func (m *ListCompaniesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCompaniesResponse.Size(m)
}
func (m *ListCompaniesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCompaniesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCompaniesResponse proto.InternalMessageInfo

func (m *ListCompaniesResponse) GetCompanies() []*Company {
	if m != nil {
		return m.Companies
	}
	return nil
}

func (m *ListCompaniesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListCompaniesResponse) GetMetadata() *ResponseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCompanyRequest)(nil), "google.cloud.talent.v4beta1.CreateCompanyRequest")
	proto.RegisterType((*GetCompanyRequest)(nil), "google.cloud.talent.v4beta1.GetCompanyRequest")
	proto.RegisterType((*UpdateCompanyRequest)(nil), "google.cloud.talent.v4beta1.UpdateCompanyRequest")
	proto.RegisterType((*DeleteCompanyRequest)(nil), "google.cloud.talent.v4beta1.DeleteCompanyRequest")
	proto.RegisterType((*ListCompaniesRequest)(nil), "google.cloud.talent.v4beta1.ListCompaniesRequest")
	proto.RegisterType((*ListCompaniesResponse)(nil), "google.cloud.talent.v4beta1.ListCompaniesResponse")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/company_service.proto", fileDescriptor_1c499e6cf60b8944)
}

var fileDescriptor_1c499e6cf60b8944 = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0xd6, 0x64, 0x4b, 0xbb, 0x99, 0xd5, 0xb6, 0xea, 0x74, 0x81, 0xe0, 0xa5, 0x6a, 0x64, 0xd0,
	0x2a, 0x84, 0xd6, 0xc3, 0xa6, 0x48, 0xa0, 0xae, 0x38, 0x38, 0xe1, 0x87, 0x40, 0xad, 0xba, 0x72,
	0xca, 0x65, 0x85, 0x14, 0x4d, 0x9c, 0xb7, 0x59, 0x37, 0xf6, 0x8c, 0xeb, 0x99, 0xec, 0xd2, 0xa2,
	0x1e, 0xe8, 0xbf, 0xc0, 0x85, 0x33, 0xff, 0x06, 0x67, 0x2e, 0x48, 0x70, 0x28, 0xb7, 0x9c, 0x56,
	0x82, 0x2b, 0xff, 0x00, 0x27, 0x64, 0x7b, 0x9c, 0xc4, 0x49, 0x70, 0xcc, 0xaa, 0x37, 0x7b, 0xde,
	0xfb, 0x66, 0xbe, 0x6f, 0xde, 0x7b, 0x9f, 0x8d, 0xf7, 0x87, 0x42, 0x0c, 0x7d, 0xa0, 0xae, 0x2f,
	0xc6, 0x03, 0xaa, 0x98, 0x0f, 0x5c, 0xd1, 0xd3, 0x0f, 0xfb, 0xa0, 0xd8, 0x3e, 0x75, 0x45, 0x10,
	0x32, 0xfe, 0xb4, 0x27, 0x21, 0x3a, 0xf5, 0x5c, 0xb0, 0xc2, 0x48, 0x28, 0x41, 0x76, 0x53, 0x88,
	0x95, 0x40, 0xac, 0x14, 0x62, 0x69, 0x88, 0xf1, 0xb6, 0xde, 0x8f, 0x85, 0x1e, 0x65, 0x9c, 0x0b,
	0xc5, 0x94, 0x27, 0xb8, 0x4c, 0xa1, 0xc6, 0x9b, 0x73, 0x51, 0xd7, 0xf7, 0x62, 0x60, 0x1a, 0xb8,
	0x35, 0x17, 0x38, 0xf6, 0xc0, 0x1f, 0xf4, 0xfa, 0x70, 0xc2, 0x4e, 0x3d, 0x11, 0xe9, 0x84, 0xb7,
	0xe6, 0x12, 0x22, 0x90, 0x62, 0x1c, 0x65, 0x7c, 0x8c, 0xc6, 0x1a, 0x09, 0x81, 0xe0, 0x3a, 0xf3,
	0xbd, 0x12, 0x62, 0x75, 0xaa, 0x16, 0x49, 0x93, 0xb7, 0xfe, 0xf8, 0x98, 0x42, 0x10, 0xaa, 0x2c,
	0x58, 0x5f, 0x0c, 0xa6, 0x94, 0x03, 0x26, 0x47, 0x69, 0x86, 0xf9, 0x23, 0xc2, 0x3b, 0x9d, 0x08,
	0x98, 0x82, 0x4e, 0xba, 0xad, 0x03, 0x4f, 0xc6, 0x20, 0x15, 0x39, 0xc0, 0x97, 0x43, 0x16, 0x01,
	0x57, 0x35, 0x54, 0x47, 0x8d, 0x6a, 0xfb, 0x9d, 0x73, 0xbb, 0xf2, 0x8f, 0x7d, 0x93, 0xec, 0x3e,
	0x16, 0x7d, 0x69, 0xa5, 0xfb, 0xb2, 0xd0, 0x93, 0x96, 0x2b, 0x02, 0x9a, 0x61, 0x35, 0x84, 0x74,
	0xf0, 0x15, 0xcd, 0xb2, 0x56, 0xa9, 0xa3, 0xc6, 0x56, 0xeb, 0x5d, 0xab, 0xa0, 0x16, 0x96, 0x86,
	0xb7, 0x37, 0xce, 0xed, 0x8a, 0x93, 0x21, 0xcd, 0xfb, 0xf8, 0xfa, 0x17, 0xa0, 0x16, 0x68, 0x7d,
	0x84, 0x2f, 0x71, 0x16, 0x40, 0x9e, 0x14, 0x2e, 0x24, 0x95, 0x00, 0x12, 0xa1, 0x5f, 0x87, 0x83,
	0x65, 0xa1, 0x73, 0x5c, 0xd1, 0x45, 0xb9, 0x92, 0x03, 0xbc, 0x35, 0x4e, 0x36, 0x4f, 0xee, 0x56,
	0x8b, 0x36, 0xb2, 0x8d, 0xb2, 0xeb, 0xb7, 0x3e, 0x8f, 0xaf, 0xff, 0x01, 0x93, 0x23, 0x07, 0xa7,
	0xe9, 0xf1, 0xb3, 0xf9, 0x10, 0xef, 0x7c, 0x0a, 0x3e, 0x2c, 0x31, 0xbb, 0xb0, 0xd6, 0x9f, 0x11,
	0xde, 0xb9, 0xef, 0x49, 0x7d, 0x77, 0x1e, 0xc8, 0x57, 0x52, 0xd4, 0x9b, 0x18, 0x87, 0x6c, 0x08,
	0x3d, 0x25, 0x46, 0xc0, 0x13, 0x89, 0x55, 0xa7, 0x1a, 0xaf, 0x3c, 0x8a, 0x17, 0xc8, 0x2e, 0x4e,
	0x5e, 0x7a, 0xd2, 0x7b, 0x06, 0xb5, 0x8d, 0x3a, 0x6a, 0xbc, 0xe6, 0x6c, 0xc6, 0x0b, 0x5d, 0xef,
	0x19, 0x90, 0x26, 0xbe, 0x1e, 0xc1, 0x93, 0xb1, 0x17, 0x41, 0x4f, 0x84, 0xc0, 0x7b, 0xf1, 0x79,
	0xb5, 0x4b, 0x75, 0xd4, 0xd8, 0x74, 0xae, 0xe9, 0xc0, 0xc3, 0x10, 0xf8, 0x57, 0xa2, 0x2f, 0xcd,
	0xdf, 0x11, 0x7e, 0x7d, 0x81, 0xbd, 0x0c, 0x05, 0x97, 0x40, 0xda, 0xb8, 0xea, 0x66, 0x8b, 0x35,
	0x54, 0xdf, 0x28, 0x5b, 0x2c, 0x67, 0x06, 0x23, 0x7b, 0xf8, 0x1a, 0x87, 0x6f, 0x55, 0x6f, 0x49,
	0xca, 0x76, 0xbc, 0x7c, 0x38, 0x95, 0xf3, 0x25, 0xde, 0x0c, 0x40, 0xb1, 0x01, 0x53, 0x2c, 0x51,
	0xb3, 0xd5, 0xba, 0x53, 0x78, 0x54, 0x46, 0xf2, 0x81, 0x06, 0x39, 0x53, 0x78, 0xeb, 0x65, 0x15,
	0x5f, 0xd5, 0x4c, 0xba, 0xa9, 0x41, 0x91, 0x3f, 0x11, 0xde, 0xce, 0x8d, 0x1d, 0xd9, 0x2f, 0x16,
	0xb2, 0x62, 0x44, 0x8d, 0x52, 0xda, 0xcd, 0xa7, 0x13, 0xfb, 0x6a, 0x5a, 0xc1, 0xdb, 0xba, 0x5b,
	0x5f, 0xfc, 0xf1, 0xd7, 0x0f, 0x95, 0xbe, 0xf9, 0xc1, 0xd4, 0x51, 0xbe, 0x4b, 0xe3, 0x9f, 0x84,
	0x91, 0x78, 0x0c, 0xae, 0x92, 0xb4, 0x49, 0x15, 0x70, 0xc6, 0xe3, 0xa7, 0xe7, 0x74, 0x7a, 0x77,
	0xf7, 0x50, 0xf3, 0xe8, 0x7d, 0x73, 0xaf, 0x00, 0x96, 0x4f, 0x26, 0xbf, 0x21, 0x8c, 0x67, 0x23,
	0x4c, 0xac, 0x42, 0xbe, 0x4b, 0xb3, 0x5e, 0x52, 0x9f, 0x37, 0xb1, 0x93, 0xae, 0x4f, 0x54, 0x7d,
	0x43, 0xe6, 0x54, 0xc5, 0xab, 0x2b, 0x35, 0xcd, 0x58, 0xd2, 0xe6, 0xf3, 0xa3, 0x06, 0xd9, 0xfb,
	0x6f, 0xcc, 0x7c, 0x26, 0xf9, 0x1b, 0xe1, 0xed, 0x9c, 0x87, 0xac, 0xa9, 0xda, 0x2a, 0xbf, 0x29,
	0xa9, 0xea, 0x05, 0x9a, 0xd8, 0x57, 0xe6, 0xeb, 0x15, 0xb5, 0x3e, 0x9e, 0xb1, 0xcc, 0x3e, 0x01,
	0xe5, 0x14, 0xc6, 0x75, 0xbb, 0xdb, 0xb2, 0xd6, 0xc3, 0x17, 0x40, 0xe4, 0x17, 0x84, 0xb7, 0x73,
	0xce, 0xb4, 0x46, 0xef, 0x2a, 0x17, 0x33, 0xde, 0x58, 0x72, 0xc1, 0xcf, 0xe2, 0x2f, 0xd4, 0x42,
	0xdd, 0x9a, 0x17, 0xa8, 0x5b, 0xb3, 0x6c, 0xdd, 0xe2, 0x69, 0xcb, 0x39, 0xca, 0x1a, 0x1d, 0xab,
	0xbc, 0xd3, 0x68, 0xfd, 0x1f, 0x48, 0xea, 0x05, 0xe6, 0x68, 0x62, 0x6b, 0xf7, 0x5c, 0xee, 0xce,
	0x72, 0x33, 0x97, 0xef, 0xce, 0xa2, 0x81, 0x33, 0xfc, 0x5f, 0xed, 0x1b, 0x2b, 0x7c, 0xfc, 0xa5,
	0xdd, 0x3d, 0x51, 0x2a, 0x94, 0xf7, 0x28, 0x3d, 0x3b, 0x3b, 0x5b, 0x34, 0x79, 0x36, 0x56, 0x27,
	0xe9, 0x6f, 0xc6, 0x9d, 0xd0, 0x67, 0xea, 0x58, 0x44, 0xc1, 0xed, 0x75, 0xe9, 0xf1, 0x21, 0xed,
	0xef, 0x11, 0xbe, 0xe5, 0x8a, 0xa0, 0xe8, 0x52, 0xda, 0x37, 0xf2, 0xae, 0x77, 0x18, 0x37, 0xc0,
	0x21, 0x3a, 0xb2, 0x35, 0x66, 0x28, 0x7c, 0xc6, 0x87, 0x96, 0x88, 0x86, 0x74, 0x08, 0x3c, 0x69,
	0x0f, 0x3a, 0x3b, 0x6f, 0xe5, 0xcf, 0xcf, 0x41, 0xfa, 0xfa, 0x53, 0x65, 0xa3, 0xf3, 0xa8, 0xdb,
	0xbf, 0x9c, 0x60, 0xee, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x85, 0x9f, 0x8f, 0x1c, 0x0a,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyServiceClient interface {
	// Creates a new company entity.
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Retrieves specified company.
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Updates specified company.
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Deletes specified company.
	// Prerequisite: The company has no jobs associated with it.
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all companies associated with the project.
	ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error) {
	out := new(ListCompaniesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/ListCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
type CompanyServiceServer interface {
	// Creates a new company entity.
	CreateCompany(context.Context, *CreateCompanyRequest) (*Company, error)
	// Retrieves specified company.
	GetCompany(context.Context, *GetCompanyRequest) (*Company, error)
	// Updates specified company.
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*Company, error)
	// Deletes specified company.
	// Prerequisite: The company has no jobs associated with it.
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*empty.Empty, error)
	// Lists all companies associated with the project.
	ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error)
}

// UnimplementedCompanyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (*UnimplementedCompanyServiceServer) CreateCompany(ctx context.Context, req *CreateCompanyRequest) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) GetCompany(ctx context.Context, req *GetCompanyRequest) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) UpdateCompany(ctx context.Context, req *UpdateCompanyRequest) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) DeleteCompany(ctx context.Context, req *DeleteCompanyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) ListCompanies(ctx context.Context, req *ListCompaniesRequest) (*ListCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanies not implemented")
}

func RegisterCompanyServiceServer(s *grpc.Server, srv CompanyServiceServer) {
	s.RegisterService(&_CompanyService_serviceDesc, srv)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/ListCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListCompanies(ctx, req.(*ListCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompanyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4beta1.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _CompanyService_GetCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyService_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _CompanyService_DeleteCompany_Handler,
		},
		{
			MethodName: "ListCompanies",
			Handler:    _CompanyService_ListCompanies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4beta1/company_service.proto",
}
