// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/policytroubleshooter/v1/explanations.proto

package policytroubleshooter

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1 "google.golang.org/genproto/googleapis/iam/v1"
	expr "google.golang.org/genproto/googleapis/type/expr"
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

// Whether a member has a permission for a resource.
type AccessState int32

const (
	// Reserved for future use.
	AccessState_ACCESS_STATE_UNSPECIFIED AccessState = 0
	// The member has the permission.
	AccessState_GRANTED AccessState = 1
	// The member does not have the permission.
	AccessState_NOT_GRANTED AccessState = 2
	// The member has the permission only if a condition expression evaluates to
	// `true`.
	AccessState_UNKNOWN_CONDITIONAL AccessState = 3
	// The sender of the request does not have access to all of the policies that
	// Policy Troubleshooter needs to evaluate.
	AccessState_UNKNOWN_INFO_DENIED AccessState = 4
)

var AccessState_name = map[int32]string{
	0: "ACCESS_STATE_UNSPECIFIED",
	1: "GRANTED",
	2: "NOT_GRANTED",
	3: "UNKNOWN_CONDITIONAL",
	4: "UNKNOWN_INFO_DENIED",
}

var AccessState_value = map[string]int32{
	"ACCESS_STATE_UNSPECIFIED": 0,
	"GRANTED":                  1,
	"NOT_GRANTED":              2,
	"UNKNOWN_CONDITIONAL":      3,
	"UNKNOWN_INFO_DENIED":      4,
}

func (x AccessState) String() string {
	return proto.EnumName(AccessState_name, int32(x))
}

func (AccessState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45bb0803ea44f1ae, []int{0}
}

// The extent to which a single data point contributes to an overall
// determination.
type HeuristicRelevance int32

const (
	// Reserved for future use.
	HeuristicRelevance_HEURISTIC_RELEVANCE_UNSPECIFIED HeuristicRelevance = 0
	// The data point has a limited effect on the result. Changing the data point
	// is unlikely to affect the overall determination.
	HeuristicRelevance_NORMAL HeuristicRelevance = 1
	// The data point has a strong effect on the result. Changing the data point
	// is likely to affect the overall determination.
	HeuristicRelevance_HIGH HeuristicRelevance = 2
)

var HeuristicRelevance_name = map[int32]string{
	0: "HEURISTIC_RELEVANCE_UNSPECIFIED",
	1: "NORMAL",
	2: "HIGH",
}

var HeuristicRelevance_value = map[string]int32{
	"HEURISTIC_RELEVANCE_UNSPECIFIED": 0,
	"NORMAL":                          1,
	"HIGH":                            2,
}

func (x HeuristicRelevance) String() string {
	return proto.EnumName(HeuristicRelevance_name, int32(x))
}

func (HeuristicRelevance) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45bb0803ea44f1ae, []int{1}
}

// Whether a role includes a specific permission.
type BindingExplanation_RolePermission int32

const (
	// Reserved for future use.
	BindingExplanation_ROLE_PERMISSION_UNSPECIFIED BindingExplanation_RolePermission = 0
	// The permission is included in the role.
	BindingExplanation_ROLE_PERMISSION_INCLUDED BindingExplanation_RolePermission = 1
	// The permission is not included in the role.
	BindingExplanation_ROLE_PERMISSION_NOT_INCLUDED BindingExplanation_RolePermission = 2
	// The sender of the request is not allowed to access the binding.
	BindingExplanation_ROLE_PERMISSION_UNKNOWN_INFO_DENIED BindingExplanation_RolePermission = 3
)

var BindingExplanation_RolePermission_name = map[int32]string{
	0: "ROLE_PERMISSION_UNSPECIFIED",
	1: "ROLE_PERMISSION_INCLUDED",
	2: "ROLE_PERMISSION_NOT_INCLUDED",
	3: "ROLE_PERMISSION_UNKNOWN_INFO_DENIED",
}

var BindingExplanation_RolePermission_value = map[string]int32{
	"ROLE_PERMISSION_UNSPECIFIED":         0,
	"ROLE_PERMISSION_INCLUDED":            1,
	"ROLE_PERMISSION_NOT_INCLUDED":        2,
	"ROLE_PERMISSION_UNKNOWN_INFO_DENIED": 3,
}

func (x BindingExplanation_RolePermission) String() string {
	return proto.EnumName(BindingExplanation_RolePermission_name, int32(x))
}

func (BindingExplanation_RolePermission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45bb0803ea44f1ae, []int{2, 0}
}

// Whether the binding includes the member.
type BindingExplanation_Membership int32

const (
	// Reserved for future use.
	BindingExplanation_MEMBERSHIP_UNSPECIFIED BindingExplanation_Membership = 0
	// The binding includes the member. The member can be included directly
	// or indirectly. For example:
	//
	// * A member is included directly if that member is listed in the binding.
	// * A member is included indirectly if that member is in a Google group or
	//   G Suite domain that is listed in the binding.
	BindingExplanation_MEMBERSHIP_INCLUDED BindingExplanation_Membership = 1
	// The binding does not include the member.
	BindingExplanation_MEMBERSHIP_NOT_INCLUDED BindingExplanation_Membership = 2
	// The sender of the request is not allowed to access the binding.
	BindingExplanation_MEMBERSHIP_UNKNOWN_INFO_DENIED BindingExplanation_Membership = 3
	// The member is an unsupported type. Only Google Accounts and service
	// accounts are supported.
	BindingExplanation_MEMBERSHIP_UNKNOWN_UNSUPPORTED BindingExplanation_Membership = 4
)

var BindingExplanation_Membership_name = map[int32]string{
	0: "MEMBERSHIP_UNSPECIFIED",
	1: "MEMBERSHIP_INCLUDED",
	2: "MEMBERSHIP_NOT_INCLUDED",
	3: "MEMBERSHIP_UNKNOWN_INFO_DENIED",
	4: "MEMBERSHIP_UNKNOWN_UNSUPPORTED",
}

var BindingExplanation_Membership_value = map[string]int32{
	"MEMBERSHIP_UNSPECIFIED":         0,
	"MEMBERSHIP_INCLUDED":            1,
	"MEMBERSHIP_NOT_INCLUDED":        2,
	"MEMBERSHIP_UNKNOWN_INFO_DENIED": 3,
	"MEMBERSHIP_UNKNOWN_UNSUPPORTED": 4,
}

func (x BindingExplanation_Membership) String() string {
	return proto.EnumName(BindingExplanation_Membership_name, int32(x))
}

func (BindingExplanation_Membership) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45bb0803ea44f1ae, []int{2, 1}
}

// Information about the member, resource, and permission to check.
type AccessTuple struct {
	// Required. The member, or principal, whose access you want to check, in the form of
	// the email address that represents that member. For example,
	// `alice@example.com` or
	// `my-service-account@my-project.iam.gserviceaccount.com`.
	//
	// The member must be a Google Account or a service account. Other types of
	// members are not supported.
	Principal string `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	// Required. The full resource name that identifies the resource. For example,
	// `//compute.googleapis.com/projects/my-project/zones/us-central1-a/instances/my-instance`.
	//
	// For examples of full resource names for Google Cloud services, see
	// https://cloud.google.com/iam/help/troubleshooter/full-resource-names.
	FullResourceName string `protobuf:"bytes,2,opt,name=full_resource_name,json=fullResourceName,proto3" json:"full_resource_name,omitempty"`
	// Required. The IAM permission to check for the specified member and resource.
	//
	// For a complete list of IAM permissions, see
	// https://cloud.google.com/iam/help/permissions/reference.
	//
	// For a complete list of predefined IAM roles and the permissions in each
	// role, see https://cloud.google.com/iam/help/roles/reference.
	Permission           string   `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessTuple) Reset()         { *m = AccessTuple{} }
func (m *AccessTuple) String() string { return proto.CompactTextString(m) }
func (*AccessTuple) ProtoMessage()    {}
func (*AccessTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_45bb0803ea44f1ae, []int{0}
}

func (m *AccessTuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessTuple.Unmarshal(m, b)
}
func (m *AccessTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessTuple.Marshal(b, m, deterministic)
}
func (m *AccessTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTuple.Merge(m, src)
}
func (m *AccessTuple) XXX_Size() int {
	return xxx_messageInfo_AccessTuple.Size(m)
}
func (m *AccessTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTuple.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTuple proto.InternalMessageInfo

func (m *AccessTuple) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *AccessTuple) GetFullResourceName() string {
	if m != nil {
		return m.FullResourceName
	}
	return ""
}

func (m *AccessTuple) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

// Details about how a specific IAM [Policy][google.iam.v1.Policy] contributed
// to the access check.
type ExplainedPolicy struct {
	// Indicates whether _this policy_ provides the specified permission to the
	// specified member for the specified resource.
	//
	// This field does _not_ indicate whether the member actually has the
	// permission for the resource. There might be another policy that overrides
	// this policy. To determine whether the member actually has the permission,
	// use the `access` field in the
	// [TroubleshootIamPolicyResponse][IamChecker.TroubleshootIamPolicyResponse].
	Access AccessState `protobuf:"varint,1,opt,name=access,proto3,enum=google.cloud.policytroubleshooter.v1.AccessState" json:"access,omitempty"`
	// The full resource name that identifies the resource. For example,
	// `//compute.googleapis.com/projects/my-project/zones/us-central1-a/instances/my-instance`.
	//
	// If the sender of the request does not have access to the policy, this field
	// is omitted.
	//
	// For examples of full resource names for Google Cloud services, see
	// https://cloud.google.com/iam/help/troubleshooter/full-resource-names.
	FullResourceName string `protobuf:"bytes,2,opt,name=full_resource_name,json=fullResourceName,proto3" json:"full_resource_name,omitempty"`
	// The IAM policy attached to the resource.
	//
	// If the sender of the request does not have access to the policy, this field
	// is empty.
	Policy *v1.Policy `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
	// Details about how each binding in the policy affects the member's ability,
	// or inability, to use the permission for the resource.
	//
	// If the sender of the request does not have access to the policy, this field
	// is omitted.
	BindingExplanations []*BindingExplanation `protobuf:"bytes,4,rep,name=binding_explanations,json=bindingExplanations,proto3" json:"binding_explanations,omitempty"`
	// The relevance of this policy to the overall determination in the
	// [TroubleshootIamPolicyResponse][IamChecker.TroubleshootIamPolicyResponse].
	//
	// If the sender of the request does not have access to the policy, this field
	// is omitted.
	Relevance            HeuristicRelevance `protobuf:"varint,5,opt,name=relevance,proto3,enum=google.cloud.policytroubleshooter.v1.HeuristicRelevance" json:"relevance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExplainedPolicy) Reset()         { *m = ExplainedPolicy{} }
func (m *ExplainedPolicy) String() string { return proto.CompactTextString(m) }
func (*ExplainedPolicy) ProtoMessage()    {}
func (*ExplainedPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_45bb0803ea44f1ae, []int{1}
}

func (m *ExplainedPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExplainedPolicy.Unmarshal(m, b)
}
func (m *ExplainedPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExplainedPolicy.Marshal(b, m, deterministic)
}
func (m *ExplainedPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExplainedPolicy.Merge(m, src)
}
func (m *ExplainedPolicy) XXX_Size() int {
	return xxx_messageInfo_ExplainedPolicy.Size(m)
}
func (m *ExplainedPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ExplainedPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ExplainedPolicy proto.InternalMessageInfo

func (m *ExplainedPolicy) GetAccess() AccessState {
	if m != nil {
		return m.Access
	}
	return AccessState_ACCESS_STATE_UNSPECIFIED
}

func (m *ExplainedPolicy) GetFullResourceName() string {
	if m != nil {
		return m.FullResourceName
	}
	return ""
}

func (m *ExplainedPolicy) GetPolicy() *v1.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *ExplainedPolicy) GetBindingExplanations() []*BindingExplanation {
	if m != nil {
		return m.BindingExplanations
	}
	return nil
}

func (m *ExplainedPolicy) GetRelevance() HeuristicRelevance {
	if m != nil {
		return m.Relevance
	}
	return HeuristicRelevance_HEURISTIC_RELEVANCE_UNSPECIFIED
}

// Details about how a binding in a policy affects a member's ability to use a
// permission.
type BindingExplanation struct {
	// Required. Indicates whether _this binding_ provides the specified permission to the
	// specified member for the specified resource.
	//
	// This field does _not_ indicate whether the member actually has the
	// permission for the resource. There might be another binding that overrides
	// this binding. To determine whether the member actually has the permission,
	// use the `access` field in the
	// [TroubleshootIamPolicyResponse][IamChecker.TroubleshootIamPolicyResponse].
	Access AccessState `protobuf:"varint,1,opt,name=access,proto3,enum=google.cloud.policytroubleshooter.v1.AccessState" json:"access,omitempty"`
	// The role that this binding grants. For example,
	// `roles/compute.serviceAgent`.
	//
	// For a complete list of predefined IAM roles, as well as the permissions in
	// each role, see https://cloud.google.com/iam/help/roles/reference.
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	// Indicates whether the role granted by this binding contains the specified
	// permission.
	RolePermission BindingExplanation_RolePermission `protobuf:"varint,3,opt,name=role_permission,json=rolePermission,proto3,enum=google.cloud.policytroubleshooter.v1.BindingExplanation_RolePermission" json:"role_permission,omitempty"`
	// The relevance of the permission's existence, or nonexistence, in the role
	// to the overall determination for the entire policy.
	RolePermissionRelevance HeuristicRelevance `protobuf:"varint,4,opt,name=role_permission_relevance,json=rolePermissionRelevance,proto3,enum=google.cloud.policytroubleshooter.v1.HeuristicRelevance" json:"role_permission_relevance,omitempty"`
	// Indicates whether each member in the binding includes the member specified
	// in the request, either directly or indirectly. Each key identifies a member
	// in the binding, and each value indicates whether the member in the binding
	// includes the member in the request.
	//
	// For example, suppose that a binding includes the following members:
	//
	// * `user:alice@example.com`
	// * `group:product-eng@example.com`
	//
	// You want to troubleshoot access for `user:bob@example.com`. This user is a
	// member of the group `group:product-eng@example.com`.
	//
	// For the first member in the binding, the key is `user:alice@example.com`,
	// and the `membership` field in the value is set to
	// `MEMBERSHIP_NOT_INCLUDED`.
	//
	// For the second member in the binding, the key is
	// `group:product-eng@example.com`, and the `membership` field in the value is
	// set to `MEMBERSHIP_INCLUDED`.
	Memberships map[string]*BindingExplanation_AnnotatedMembership `protobuf:"bytes,5,rep,name=memberships,proto3" json:"memberships,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The relevance of this binding to the overall determination for the entire
	// policy.
	Relevance HeuristicRelevance `protobuf:"varint,6,opt,name=relevance,proto3,enum=google.cloud.policytroubleshooter.v1.HeuristicRelevance" json:"relevance,omitempty"`
	// A condition expression that prevents access unless the expression evaluates
	// to `true`.
	//
	// To learn about IAM Conditions, see
	// http://cloud.google.com/iam/help/conditions/overview.
	Condition            *expr.Expr `protobuf:"bytes,7,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BindingExplanation) Reset()         { *m = BindingExplanation{} }
func (m *BindingExplanation) String() string { return proto.CompactTextString(m) }
func (*BindingExplanation) ProtoMessage()    {}
func (*BindingExplanation) Descriptor() ([]byte, []int) {
	return fileDescriptor_45bb0803ea44f1ae, []int{2}
}

func (m *BindingExplanation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingExplanation.Unmarshal(m, b)
}
func (m *BindingExplanation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingExplanation.Marshal(b, m, deterministic)
}
func (m *BindingExplanation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingExplanation.Merge(m, src)
}
func (m *BindingExplanation) XXX_Size() int {
	return xxx_messageInfo_BindingExplanation.Size(m)
}
func (m *BindingExplanation) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingExplanation.DiscardUnknown(m)
}

var xxx_messageInfo_BindingExplanation proto.InternalMessageInfo

func (m *BindingExplanation) GetAccess() AccessState {
	if m != nil {
		return m.Access
	}
	return AccessState_ACCESS_STATE_UNSPECIFIED
}

func (m *BindingExplanation) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *BindingExplanation) GetRolePermission() BindingExplanation_RolePermission {
	if m != nil {
		return m.RolePermission
	}
	return BindingExplanation_ROLE_PERMISSION_UNSPECIFIED
}

func (m *BindingExplanation) GetRolePermissionRelevance() HeuristicRelevance {
	if m != nil {
		return m.RolePermissionRelevance
	}
	return HeuristicRelevance_HEURISTIC_RELEVANCE_UNSPECIFIED
}

func (m *BindingExplanation) GetMemberships() map[string]*BindingExplanation_AnnotatedMembership {
	if m != nil {
		return m.Memberships
	}
	return nil
}

func (m *BindingExplanation) GetRelevance() HeuristicRelevance {
	if m != nil {
		return m.Relevance
	}
	return HeuristicRelevance_HEURISTIC_RELEVANCE_UNSPECIFIED
}

func (m *BindingExplanation) GetCondition() *expr.Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

// Details about whether the binding includes the member.
type BindingExplanation_AnnotatedMembership struct {
	// Indicates whether the binding includes the member.
	Membership BindingExplanation_Membership `protobuf:"varint,1,opt,name=membership,proto3,enum=google.cloud.policytroubleshooter.v1.BindingExplanation_Membership" json:"membership,omitempty"`
	// The relevance of the member's status to the overall determination for the
	// binding.
	Relevance            HeuristicRelevance `protobuf:"varint,2,opt,name=relevance,proto3,enum=google.cloud.policytroubleshooter.v1.HeuristicRelevance" json:"relevance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BindingExplanation_AnnotatedMembership) Reset() {
	*m = BindingExplanation_AnnotatedMembership{}
}
func (m *BindingExplanation_AnnotatedMembership) String() string { return proto.CompactTextString(m) }
func (*BindingExplanation_AnnotatedMembership) ProtoMessage()    {}
func (*BindingExplanation_AnnotatedMembership) Descriptor() ([]byte, []int) {
	return fileDescriptor_45bb0803ea44f1ae, []int{2, 0}
}

func (m *BindingExplanation_AnnotatedMembership) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingExplanation_AnnotatedMembership.Unmarshal(m, b)
}
func (m *BindingExplanation_AnnotatedMembership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingExplanation_AnnotatedMembership.Marshal(b, m, deterministic)
}
func (m *BindingExplanation_AnnotatedMembership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingExplanation_AnnotatedMembership.Merge(m, src)
}
func (m *BindingExplanation_AnnotatedMembership) XXX_Size() int {
	return xxx_messageInfo_BindingExplanation_AnnotatedMembership.Size(m)
}
func (m *BindingExplanation_AnnotatedMembership) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingExplanation_AnnotatedMembership.DiscardUnknown(m)
}

var xxx_messageInfo_BindingExplanation_AnnotatedMembership proto.InternalMessageInfo

func (m *BindingExplanation_AnnotatedMembership) GetMembership() BindingExplanation_Membership {
	if m != nil {
		return m.Membership
	}
	return BindingExplanation_MEMBERSHIP_UNSPECIFIED
}

func (m *BindingExplanation_AnnotatedMembership) GetRelevance() HeuristicRelevance {
	if m != nil {
		return m.Relevance
	}
	return HeuristicRelevance_HEURISTIC_RELEVANCE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("google.cloud.policytroubleshooter.v1.AccessState", AccessState_name, AccessState_value)
	proto.RegisterEnum("google.cloud.policytroubleshooter.v1.HeuristicRelevance", HeuristicRelevance_name, HeuristicRelevance_value)
	proto.RegisterEnum("google.cloud.policytroubleshooter.v1.BindingExplanation_RolePermission", BindingExplanation_RolePermission_name, BindingExplanation_RolePermission_value)
	proto.RegisterEnum("google.cloud.policytroubleshooter.v1.BindingExplanation_Membership", BindingExplanation_Membership_name, BindingExplanation_Membership_value)
	proto.RegisterType((*AccessTuple)(nil), "google.cloud.policytroubleshooter.v1.AccessTuple")
	proto.RegisterType((*ExplainedPolicy)(nil), "google.cloud.policytroubleshooter.v1.ExplainedPolicy")
	proto.RegisterType((*BindingExplanation)(nil), "google.cloud.policytroubleshooter.v1.BindingExplanation")
	proto.RegisterMapType((map[string]*BindingExplanation_AnnotatedMembership)(nil), "google.cloud.policytroubleshooter.v1.BindingExplanation.MembershipsEntry")
	proto.RegisterType((*BindingExplanation_AnnotatedMembership)(nil), "google.cloud.policytroubleshooter.v1.BindingExplanation.AnnotatedMembership")
}

func init() {
	proto.RegisterFile("google/cloud/policytroubleshooter/v1/explanations.proto", fileDescriptor_45bb0803ea44f1ae)
}

var fileDescriptor_45bb0803ea44f1ae = []byte{
	// 849 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdf, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x71, 0x92, 0x66, 0xe9, 0x89, 0xd4, 0x9a, 0x29, 0x6c, 0x43, 0x76, 0x45, 0x4b, 0x8a,
	0xc4, 0x6a, 0x05, 0x8e, 0x1a, 0x2e, 0x58, 0xc1, 0x95, 0x9b, 0xcc, 0x36, 0x16, 0x89, 0x1d, 0x8d,
	0x9d, 0x05, 0xed, 0x8d, 0xe5, 0x38, 0xb3, 0xe9, 0xa8, 0x8e, 0xc7, 0x1a, 0x3b, 0xd1, 0x56, 0xdc,
	0x70, 0xc9, 0x05, 0x4f, 0xc0, 0x03, 0xf0, 0x40, 0x3c, 0x01, 0x4f, 0x82, 0x90, 0xff, 0xad, 0x1d,
	0x37, 0x8b, 0xaa, 0x6e, 0xaf, 0xe2, 0xcc, 0x39, 0xf3, 0x9d, 0xdf, 0x7c, 0xe7, 0x64, 0x1c, 0xf8,
	0x7e, 0xc9, 0xf9, 0xd2, 0xa3, 0x3d, 0xd7, 0xe3, 0xeb, 0x45, 0x2f, 0xe0, 0x1e, 0x73, 0x6f, 0x22,
	0xc1, 0xd7, 0x73, 0x8f, 0x86, 0x57, 0x9c, 0x47, 0x54, 0xf4, 0x36, 0xe7, 0x3d, 0xfa, 0x36, 0xf0,
	0x1c, 0xdf, 0x89, 0x18, 0xf7, 0x43, 0x25, 0x10, 0x3c, 0xe2, 0xe8, 0xab, 0x74, 0xa3, 0x92, 0x6c,
	0x54, 0x76, 0x6d, 0x54, 0x36, 0xe7, 0x9d, 0x93, 0x4c, 0xde, 0x09, 0x58, 0xef, 0x0d, 0xa3, 0xde,
	0xc2, 0x9e, 0xd3, 0x2b, 0x67, 0xc3, 0xb8, 0x48, 0x65, 0x3a, 0x9d, 0x2c, 0x81, 0x39, 0xab, 0xb8,
	0x50, 0xaa, 0x93, 0xc5, 0x1e, 0x67, 0xb1, 0xe8, 0x26, 0xa0, 0x31, 0x42, 0xb6, 0xa7, 0xfb, 0xbb,
	0x04, 0x2d, 0xd5, 0x75, 0x69, 0x18, 0x5a, 0xeb, 0xc0, 0xa3, 0xe8, 0x4b, 0xd8, 0x0f, 0x04, 0xf3,
	0x5d, 0x16, 0x38, 0x5e, 0x5b, 0x3a, 0x95, 0x9e, 0xed, 0x5f, 0xd4, 0xff, 0x51, 0x6b, 0xa4, 0x58,
	0x45, 0xe7, 0x80, 0xde, 0xac, 0x3d, 0xcf, 0x16, 0x34, 0xe4, 0x6b, 0xe1, 0x52, 0xdb, 0x77, 0x56,
	0xb4, 0x5d, 0x2b, 0x72, 0xe5, 0x38, 0x4c, 0xb2, 0xa8, 0xee, 0xac, 0x28, 0x3a, 0x03, 0x08, 0xa8,
	0x58, 0xb1, 0x30, 0x64, 0xdc, 0x6f, 0xd7, 0x8b, 0xd4, 0xd2, 0x72, 0xf7, 0xdf, 0x1a, 0x1c, 0xe2,
	0xd8, 0x1c, 0xe6, 0xd3, 0xc5, 0x34, 0x81, 0x47, 0x1a, 0x34, 0x9d, 0x84, 0x2e, 0x61, 0x39, 0xe8,
	0x9f, 0x2b, 0x77, 0xb1, 0x4a, 0x49, 0x4f, 0x64, 0x46, 0x4e, 0x44, 0x49, 0x26, 0x80, 0xbe, 0x79,
	0x3f, 0xf6, 0x0e, 0xe2, 0x6f, 0xa1, 0x99, 0x8a, 0x27, 0xb4, 0xad, 0xfe, 0x67, 0x79, 0x61, 0xe6,
	0xac, 0xe2, 0x0a, 0x29, 0x1f, 0xc9, 0x92, 0xd0, 0x35, 0x7c, 0x3a, 0x67, 0xfe, 0x82, 0xf9, 0x4b,
	0xbb, 0xdc, 0xdf, 0x76, 0xe3, 0xb4, 0xfe, 0xac, 0xd5, 0x7f, 0x71, 0x37, 0xea, 0x8b, 0x54, 0x01,
	0x17, 0x02, 0xe4, 0x68, 0x7e, 0x6b, 0x2d, 0x44, 0xaf, 0x60, 0x5f, 0x50, 0x8f, 0x6e, 0x1c, 0xdf,
	0xa5, 0xed, 0xbd, 0xc4, 0x97, 0x3b, 0x56, 0x18, 0xd1, 0xb5, 0x60, 0x61, 0xc4, 0x5c, 0x92, 0xef,
	0x27, 0x85, 0x54, 0xf7, 0x37, 0x00, 0x74, 0x9b, 0x01, 0xe9, 0x1f, 0xdc, 0x83, 0xb4, 0xd7, 0x79,
	0x23, 0x10, 0x34, 0x04, 0xf7, 0x72, 0xeb, 0x93, 0x67, 0x14, 0xc0, 0x61, 0xfc, 0x69, 0x57, 0xa6,
	0xe4, 0xa0, 0x7f, 0x79, 0x5f, 0xeb, 0x14, 0xc2, 0x3d, 0x3a, 0x7d, 0x27, 0x47, 0x0e, 0xc4, 0xd6,
	0x77, 0x14, 0xc1, 0xe7, 0x95, 0x8a, 0x76, 0x61, 0x6a, 0xe3, 0x03, 0x4d, 0x3d, 0xde, 0x2e, 0xf6,
	0x2e, 0x80, 0xae, 0xa1, 0xb5, 0xa2, 0xab, 0x39, 0x15, 0xe1, 0x15, 0x0b, 0xc2, 0xf6, 0x5e, 0x32,
	0x1e, 0xda, 0xbd, 0xcf, 0x38, 0x29, 0xb4, 0xb0, 0x1f, 0x89, 0x1b, 0x52, 0x56, 0xdf, 0x9e, 0x93,
	0xe6, 0x83, 0xcd, 0x09, 0xea, 0xc1, 0xbe, 0xcb, 0xfd, 0x05, 0x8b, 0x11, 0xda, 0x8f, 0x92, 0x9f,
	0xc7, 0x27, 0xb9, 0x6e, 0x7c, 0xbf, 0x28, 0xf8, 0x6d, 0x20, 0x48, 0x91, 0xd3, 0xf9, 0x5b, 0x82,
	0x23, 0xd5, 0xf7, 0x79, 0x3c, 0x0b, 0x8b, 0x82, 0x19, 0xb9, 0x00, 0x05, 0x6f, 0x36, 0x5d, 0x83,
	0x07, 0x30, 0x83, 0x94, 0x64, 0xb7, 0x5d, 0xa8, 0x3d, 0x98, 0x0b, 0x9d, 0x3f, 0x24, 0x90, 0xab,
	0xfe, 0x23, 0x19, 0xea, 0xd7, 0xf4, 0x26, 0xbd, 0x38, 0x49, 0xfc, 0x88, 0xe6, 0xb0, 0xb7, 0x71,
	0xbc, 0x75, 0x5a, 0xba, 0xd5, 0x1f, 0xdf, 0xfb, 0x78, 0x3b, 0x0c, 0x24, 0xa9, 0xf4, 0x0f, 0xb5,
	0x17, 0x52, 0xf7, 0x4f, 0x09, 0x0e, 0xb6, 0x47, 0x1e, 0x9d, 0xc0, 0x13, 0x62, 0x8c, 0xb1, 0x3d,
	0xc5, 0x64, 0xa2, 0x99, 0xa6, 0x66, 0xe8, 0xf6, 0x4c, 0x37, 0xa7, 0x78, 0xa0, 0xbd, 0xd4, 0xf0,
	0x50, 0xfe, 0x08, 0x3d, 0x85, 0x76, 0x35, 0x41, 0xd3, 0x07, 0xe3, 0xd9, 0x10, 0x0f, 0x65, 0x09,
	0x9d, 0xc2, 0xd3, 0x6a, 0x54, 0x37, 0xac, 0x22, 0xa3, 0x86, 0xbe, 0x86, 0xb3, 0xdb, 0x05, 0x7e,
	0xd2, 0x8d, 0x9f, 0x63, 0x9d, 0x97, 0x86, 0x3d, 0xc4, 0x7a, 0x5c, 0xa8, 0xde, 0xfd, 0x4b, 0x02,
	0x28, 0xf5, 0xbd, 0x03, 0x8f, 0x27, 0x78, 0x72, 0x81, 0x89, 0x39, 0xd2, 0xa6, 0x15, 0xa6, 0x63,
	0x38, 0x2a, 0xc5, 0x4a, 0x38, 0x4f, 0xe0, 0xb8, 0x14, 0xa8, 0x90, 0x74, 0xe1, 0x8b, 0x2d, 0xc5,
	0x1d, 0x10, 0xef, 0xc9, 0x99, 0xe9, 0xe6, 0x6c, 0x3a, 0x35, 0x88, 0x85, 0x87, 0x72, 0xe3, 0xf9,
	0xaf, 0xf9, 0xdb, 0x30, 0xb9, 0xb7, 0x62, 0x83, 0xd4, 0xc1, 0x00, 0x9b, 0xa6, 0x6d, 0x5a, 0xaa,
	0x85, 0x2b, 0xa8, 0x2d, 0x78, 0x74, 0x49, 0x54, 0xdd, 0x4a, 0xf0, 0x0e, 0xa1, 0x15, 0x33, 0xe5,
	0x0b, 0xb5, 0xf8, 0x20, 0x79, 0x8d, 0x81, 0xa1, 0x0f, 0x35, 0x4b, 0x33, 0x74, 0x75, 0x2c, 0xd7,
	0xcb, 0x81, 0x32, 0x60, 0xe3, 0xb9, 0x01, 0xe8, 0xf6, 0xc8, 0xa1, 0x33, 0x38, 0x19, 0xe1, 0x19,
	0xd1, 0x4c, 0x4b, 0x1b, 0xd8, 0x04, 0x8f, 0xf1, 0x2b, 0x55, 0x1f, 0x54, 0x51, 0x00, 0x9a, 0xba,
	0x41, 0x26, 0xea, 0x58, 0x96, 0xd0, 0xc7, 0xd0, 0x18, 0x69, 0x97, 0x23, 0xb9, 0x76, 0xf1, 0xfa,
	0xf5, 0x2f, 0xd9, 0xb4, 0x2d, 0xb9, 0xe7, 0xf8, 0x4b, 0x85, 0x8b, 0x65, 0x6f, 0x49, 0xfd, 0xe4,
	0xd5, 0xdf, 0x4b, 0x43, 0x4e, 0xc0, 0xc2, 0xff, 0xff, 0xc7, 0xf2, 0xe3, 0xae, 0xf5, 0x79, 0x33,
	0x11, 0xf9, 0xee, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x45, 0x29, 0x05, 0xf5, 0x08, 0x00,
	0x00,
}
