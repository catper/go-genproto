// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/account_budget_proposal_error.proto

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

// Enum describing possible account budget proposal errors.
type AccountBudgetProposalErrorEnum_AccountBudgetProposalError int32

const (
	// Enum unspecified.
	AccountBudgetProposalErrorEnum_UNSPECIFIED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 0
	// The received error code is not known in this version.
	AccountBudgetProposalErrorEnum_UNKNOWN AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 1
	// The field mask must be empty for create/end/remove proposals.
	AccountBudgetProposalErrorEnum_FIELD_MASK_NOT_ALLOWED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 2
	// The field cannot be set because of the proposal type.
	AccountBudgetProposalErrorEnum_IMMUTABLE_FIELD AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 3
	// The field is required because of the proposal type.
	AccountBudgetProposalErrorEnum_REQUIRED_FIELD_MISSING AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 4
	// Proposals that have been approved cannot be cancelled.
	AccountBudgetProposalErrorEnum_CANNOT_CANCEL_APPROVED_PROPOSAL AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 5
	// Budgets that haven't been approved cannot be removed.
	AccountBudgetProposalErrorEnum_CANNOT_REMOVE_UNAPPROVED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 6
	// Budgets that are currently running cannot be removed.
	AccountBudgetProposalErrorEnum_CANNOT_REMOVE_RUNNING_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 7
	// Budgets that haven't been approved cannot be truncated.
	AccountBudgetProposalErrorEnum_CANNOT_END_UNAPPROVED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 8
	// Only budgets that are currently running can be truncated.
	AccountBudgetProposalErrorEnum_CANNOT_END_INACTIVE_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 9
	// All budgets must have names.
	AccountBudgetProposalErrorEnum_BUDGET_NAME_REQUIRED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 10
	// Expired budgets cannot be edited after a sufficient amount of time has
	// passed.
	AccountBudgetProposalErrorEnum_CANNOT_UPDATE_OLD_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 11
	// It is not permissible a propose a new budget that ends in the past.
	AccountBudgetProposalErrorEnum_CANNOT_END_IN_PAST AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 12
	// An expired budget cannot be extended to overlap with the running budget.
	AccountBudgetProposalErrorEnum_CANNOT_EXTEND_END_TIME AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 13
	// A purchase order number is required.
	AccountBudgetProposalErrorEnum_PURCHASE_ORDER_NUMBER_REQUIRED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 14
	// Budgets that have a pending update cannot be updated.
	AccountBudgetProposalErrorEnum_PENDING_UPDATE_PROPOSAL_EXISTS AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 15
	// Cannot propose more than one budget when the corresponding billing setup
	// hasn't been approved.
	AccountBudgetProposalErrorEnum_MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 16
	// Cannot update the start time of a budget that has already started.
	AccountBudgetProposalErrorEnum_CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 17
	// Cannot update the spending limit of a budget with an amount lower than
	// what has already been spent.
	AccountBudgetProposalErrorEnum_SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 18
	// Cannot propose a budget update without actually changing any fields.
	AccountBudgetProposalErrorEnum_UPDATE_IS_NO_OP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 19
	// The end time must come after the start time.
	AccountBudgetProposalErrorEnum_END_TIME_MUST_FOLLOW_START_TIME AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 20
	// The budget's date range must fall within the date range of its billing
	// setup.
	AccountBudgetProposalErrorEnum_BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 21
	// The user is not authorized to mutate budgets for the given billing setup.
	AccountBudgetProposalErrorEnum_NOT_AUTHORIZED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 22
	// Mutates are not allowed for the given billing setup.
	AccountBudgetProposalErrorEnum_INVALID_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 23
)

var AccountBudgetProposalErrorEnum_AccountBudgetProposalError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "FIELD_MASK_NOT_ALLOWED",
	3:  "IMMUTABLE_FIELD",
	4:  "REQUIRED_FIELD_MISSING",
	5:  "CANNOT_CANCEL_APPROVED_PROPOSAL",
	6:  "CANNOT_REMOVE_UNAPPROVED_BUDGET",
	7:  "CANNOT_REMOVE_RUNNING_BUDGET",
	8:  "CANNOT_END_UNAPPROVED_BUDGET",
	9:  "CANNOT_END_INACTIVE_BUDGET",
	10: "BUDGET_NAME_REQUIRED",
	11: "CANNOT_UPDATE_OLD_BUDGET",
	12: "CANNOT_END_IN_PAST",
	13: "CANNOT_EXTEND_END_TIME",
	14: "PURCHASE_ORDER_NUMBER_REQUIRED",
	15: "PENDING_UPDATE_PROPOSAL_EXISTS",
	16: "MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP",
	17: "CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET",
	18: "SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED",
	19: "UPDATE_IS_NO_OP",
	20: "END_TIME_MUST_FOLLOW_START_TIME",
	21: "BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP",
	22: "NOT_AUTHORIZED",
	23: "INVALID_BILLING_SETUP",
}

var AccountBudgetProposalErrorEnum_AccountBudgetProposalError_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"FIELD_MASK_NOT_ALLOWED":          2,
	"IMMUTABLE_FIELD":                 3,
	"REQUIRED_FIELD_MISSING":          4,
	"CANNOT_CANCEL_APPROVED_PROPOSAL": 5,
	"CANNOT_REMOVE_UNAPPROVED_BUDGET": 6,
	"CANNOT_REMOVE_RUNNING_BUDGET":    7,
	"CANNOT_END_UNAPPROVED_BUDGET":    8,
	"CANNOT_END_INACTIVE_BUDGET":      9,
	"BUDGET_NAME_REQUIRED":            10,
	"CANNOT_UPDATE_OLD_BUDGET":        11,
	"CANNOT_END_IN_PAST":              12,
	"CANNOT_EXTEND_END_TIME":          13,
	"PURCHASE_ORDER_NUMBER_REQUIRED":  14,
	"PENDING_UPDATE_PROPOSAL_EXISTS":  15,
	"MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP": 16,
	"CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET":               17,
	"SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED":        18,
	"UPDATE_IS_NO_OP":                                   19,
	"END_TIME_MUST_FOLLOW_START_TIME":                   20,
	"BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP": 21,
	"NOT_AUTHORIZED":                                    22,
	"INVALID_BILLING_SETUP":                             23,
}

func (x AccountBudgetProposalErrorEnum_AccountBudgetProposalError) String() string {
	return proto.EnumName(AccountBudgetProposalErrorEnum_AccountBudgetProposalError_name, int32(x))
}

func (AccountBudgetProposalErrorEnum_AccountBudgetProposalError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b846ff75902c2ac4, []int{0, 0}
}

// Container for enum describing possible account budget proposal errors.
type AccountBudgetProposalErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBudgetProposalErrorEnum) Reset()         { *m = AccountBudgetProposalErrorEnum{} }
func (m *AccountBudgetProposalErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetProposalErrorEnum) ProtoMessage()    {}
func (*AccountBudgetProposalErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b846ff75902c2ac4, []int{0}
}

func (m *AccountBudgetProposalErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetProposalErrorEnum.Unmarshal(m, b)
}
func (m *AccountBudgetProposalErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetProposalErrorEnum.Marshal(b, m, deterministic)
}
func (m *AccountBudgetProposalErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetProposalErrorEnum.Merge(m, src)
}
func (m *AccountBudgetProposalErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetProposalErrorEnum.Size(m)
}
func (m *AccountBudgetProposalErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetProposalErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetProposalErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.AccountBudgetProposalErrorEnum_AccountBudgetProposalError", AccountBudgetProposalErrorEnum_AccountBudgetProposalError_name, AccountBudgetProposalErrorEnum_AccountBudgetProposalError_value)
	proto.RegisterType((*AccountBudgetProposalErrorEnum)(nil), "google.ads.googleads.v1.errors.AccountBudgetProposalErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/account_budget_proposal_error.proto", fileDescriptor_b846ff75902c2ac4)
}

var fileDescriptor_b846ff75902c2ac4 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcb, 0x6e, 0x13, 0x3d,
	0x14, 0xfe, 0x9b, 0xfe, 0xa4, 0xe0, 0x42, 0x6b, 0xdc, 0x0b, 0x25, 0xaa, 0x52, 0x14, 0x96, 0x48,
	0x13, 0x05, 0x04, 0x12, 0x83, 0x58, 0x38, 0x33, 0x4e, 0x62, 0x75, 0xc6, 0x36, 0xbe, 0xa4, 0x55,
	0x15, 0xc9, 0x4a, 0x9b, 0x28, 0xaa, 0xd4, 0x66, 0xa2, 0x4c, 0xda, 0x07, 0x62, 0xc9, 0x8e, 0xd7,
	0x60, 0xcd, 0x53, 0xf0, 0x08, 0xac, 0x90, 0x67, 0x32, 0x69, 0x52, 0x51, 0x16, 0x51, 0xce, 0xd8,
	0xdf, 0xe5, 0x9c, 0x63, 0x1f, 0x83, 0xe6, 0x28, 0x49, 0x46, 0x57, 0xc3, 0x7a, 0x7f, 0x90, 0xd6,
	0xf3, 0xd0, 0x45, 0xb7, 0x8d, 0xfa, 0x70, 0x3a, 0x4d, 0xa6, 0x69, 0xbd, 0x7f, 0x71, 0x91, 0xdc,
	0x8c, 0x67, 0xf6, 0xfc, 0x66, 0x30, 0x1a, 0xce, 0xec, 0x64, 0x9a, 0x4c, 0x92, 0xb4, 0x7f, 0x65,
	0xb3, 0x6d, 0x6f, 0x32, 0x4d, 0x66, 0x09, 0xaa, 0xe6, 0x44, 0xaf, 0x3f, 0x48, 0xbd, 0x85, 0x86,
	0x77, 0xdb, 0xf0, 0x72, 0x8d, 0xca, 0x61, 0xe1, 0x31, 0xb9, 0xac, 0xf7, 0xc7, 0xe3, 0x64, 0xd6,
	0x9f, 0x5d, 0x26, 0xe3, 0x34, 0x67, 0xd7, 0x7e, 0x96, 0x41, 0x15, 0xe7, 0x2e, 0xcd, 0xcc, 0x44,
	0xcc, 0x3d, 0x88, 0x63, 0x93, 0xf1, 0xcd, 0x75, 0xed, 0x7b, 0x19, 0x54, 0x1e, 0x86, 0xa0, 0x6d,
	0xb0, 0x69, 0x98, 0x12, 0x24, 0xa0, 0x2d, 0x4a, 0x42, 0xf8, 0x1f, 0xda, 0x04, 0x1b, 0x86, 0x1d,
	0x33, 0x7e, 0xc2, 0xe0, 0x1a, 0xaa, 0x80, 0xfd, 0x16, 0x25, 0x51, 0x68, 0x63, 0xac, 0x8e, 0x2d,
	0xe3, 0xda, 0xe2, 0x28, 0xe2, 0x27, 0x24, 0x84, 0x25, 0xb4, 0x03, 0xb6, 0x69, 0x1c, 0x1b, 0x8d,
	0x9b, 0x11, 0xb1, 0x19, 0x0a, 0xae, 0x3b, 0x82, 0x24, 0x5f, 0x0c, 0x95, 0x24, 0xb4, 0x73, 0x26,
	0x55, 0x8a, 0xb2, 0x36, 0xfc, 0x1f, 0xbd, 0x06, 0x47, 0x01, 0x66, 0x4e, 0x24, 0xc0, 0x2c, 0x20,
	0x91, 0xc5, 0x42, 0x48, 0xde, 0x25, 0xa1, 0x15, 0x92, 0x0b, 0xae, 0x70, 0x04, 0x1f, 0x2d, 0x81,
	0x24, 0x89, 0x79, 0x97, 0x58, 0xc3, 0x16, 0xb0, 0xa6, 0x09, 0xdb, 0x44, 0xc3, 0x32, 0x7a, 0x05,
	0x0e, 0x57, 0x41, 0xd2, 0x30, 0x46, 0x59, 0xbb, 0x40, 0x6c, 0x2c, 0x21, 0x08, 0x0b, 0xff, 0xa2,
	0xf1, 0x18, 0x55, 0x41, 0x65, 0x09, 0x41, 0x19, 0x0e, 0x34, 0xed, 0x92, 0x62, 0xff, 0x09, 0x3a,
	0x00, 0xbb, 0x79, 0x6c, 0x19, 0x8e, 0x89, 0x2d, 0xaa, 0x82, 0x00, 0x1d, 0x82, 0x83, 0x39, 0xd3,
	0x88, 0x10, 0x6b, 0x62, 0x79, 0xb4, 0xd0, 0xdd, 0x44, 0xfb, 0x00, 0xad, 0xe8, 0x5a, 0x81, 0x95,
	0x86, 0x4f, 0x5d, 0x67, 0x8a, 0xf5, 0x53, 0xed, 0xb6, 0xdc, 0x4f, 0xd3, 0x98, 0xc0, 0x67, 0xa8,
	0x06, 0xaa, 0xc2, 0xc8, 0xa0, 0x83, 0x15, 0xb1, 0x5c, 0x86, 0x44, 0x5a, 0x66, 0xe2, 0x26, 0x91,
	0x77, 0xae, 0x5b, 0x19, 0x86, 0xb0, 0xd0, 0x55, 0x39, 0xb7, 0x2d, 0xba, 0x66, 0xc9, 0x29, 0x55,
	0x5a, 0xc1, 0x6d, 0xf4, 0x19, 0x7c, 0x8c, 0x4d, 0xa4, 0xa9, 0x88, 0x8a, 0x42, 0xd4, 0xf2, 0xa1,
	0xd9, 0x16, 0x97, 0x2b, 0xbd, 0xa0, 0x51, 0xe4, 0xf4, 0x14, 0xd1, 0x46, 0x40, 0x88, 0xea, 0xe0,
	0xcd, 0x6a, 0x61, 0x4a, 0x63, 0xa9, 0xb3, 0x24, 0x33, 0x6a, 0xf6, 0x79, 0xd7, 0xc3, 0xe7, 0xe8,
	0x03, 0x78, 0xab, 0x8a, 0xa4, 0x22, 0x1a, 0x53, 0x6d, 0x9d, 0x8f, 0xb4, 0xba, 0x83, 0x99, 0xc5,
	0x41, 0x20, 0x0d, 0x09, 0x6d, 0xc0, 0x95, 0x5e, 0xb9, 0x3a, 0xc8, 0x5d, 0x9d, 0xb9, 0x03, 0x75,
	0x09, 0x5a, 0x2e, 0xe0, 0x8e, 0x3b, 0xf9, 0xa2, 0x25, 0x36, 0x36, 0x4a, 0xdb, 0x16, 0x77, 0x84,
	0xa5, 0x24, 0xe0, 0x2e, 0x7a, 0x0f, 0x1a, 0xf3, 0x53, 0xc9, 0xe8, 0x12, 0xb3, 0x36, 0xb1, 0x94,
	0x05, 0x3c, 0x16, 0x58, 0x53, 0x77, 0x13, 0x4f, 0xa8, 0xee, 0xdc, 0xab, 0x6c, 0x0f, 0x21, 0xb0,
	0x95, 0x65, 0x60, 0x74, 0x87, 0x4b, 0x7a, 0x46, 0x42, 0xb8, 0x8f, 0x5e, 0x82, 0x3d, 0xca, 0xba,
	0x38, 0xa2, 0xf7, 0x1b, 0xf1, 0xa2, 0xf9, 0x7b, 0x0d, 0xd4, 0x2e, 0x92, 0x6b, 0xef, 0xdf, 0xb3,
	0xd9, 0x3c, 0x7a, 0x78, 0xae, 0x84, 0x1b, 0x4f, 0xb1, 0x76, 0x16, 0xce, 0x25, 0x46, 0xc9, 0x55,
	0x7f, 0x3c, 0xf2, 0x92, 0xe9, 0xa8, 0x3e, 0x1a, 0x8e, 0xb3, 0xe1, 0x2d, 0x9e, 0x8c, 0xc9, 0x65,
	0xfa, 0xd0, 0x0b, 0xf2, 0x29, 0xff, 0xfb, 0x5a, 0x5a, 0x6f, 0x63, 0xfc, 0xad, 0x54, 0x6d, 0xe7,
	0x62, 0x78, 0x90, 0x7a, 0x79, 0xe8, 0xa2, 0x6e, 0xc3, 0xcb, 0x2c, 0xd3, 0x1f, 0x05, 0xa0, 0x87,
	0x07, 0x69, 0x6f, 0x01, 0xe8, 0x75, 0x1b, 0xbd, 0x1c, 0xf0, 0xab, 0x54, 0xcb, 0x57, 0x7d, 0x1f,
	0x0f, 0x52, 0xdf, 0x5f, 0x40, 0x7c, 0xbf, 0xdb, 0xf0, 0xfd, 0x1c, 0x74, 0x5e, 0xce, 0xb2, 0x7b,
	0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x81, 0x26, 0x78, 0xde, 0x04, 0x00, 0x00,
}
