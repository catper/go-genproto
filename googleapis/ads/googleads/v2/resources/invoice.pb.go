// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/invoice.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// An invoice. All invoice information is snapshotted to match the PDF invoice.
// For invoices older than the launch of InvoiceService, the snapshotted
// information may not match the PDF invoice.
type Invoice struct {
	// Immutable. The resource name of the invoice. Multiple customers can share a given
	// invoice, so multiple resource names may point to the same invoice.
	// Invoice resource names have the form:
	//
	// `customers/{customer_id}/invoices/{invoice_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the invoice. It appears on the invoice PDF as "Invoice number".
	Id *wrappers.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The type of invoice.
	Type enums.InvoiceTypeEnum_InvoiceType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.InvoiceTypeEnum_InvoiceType" json:"type,omitempty"`
	// Output only. The resource name of this invoice’s billing setup.
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	BillingSetup *wrappers.StringValue `protobuf:"bytes,4,opt,name=billing_setup,json=billingSetup,proto3" json:"billing_setup,omitempty"`
	// Output only. A 16 digit ID used to identify the payments account associated with the
	// billing setup, e.g. "1234-5678-9012-3456". It appears on the invoice PDF as
	// "Billing Account Number".
	PaymentsAccountId *wrappers.StringValue `protobuf:"bytes,5,opt,name=payments_account_id,json=paymentsAccountId,proto3" json:"payments_account_id,omitempty"`
	// Output only. A 12 digit ID used to identify the payments profile associated with the
	// billing setup, e.g. "1234-5678-9012". It appears on the invoice PDF as
	// "Billing ID".
	PaymentsProfileId *wrappers.StringValue `protobuf:"bytes,6,opt,name=payments_profile_id,json=paymentsProfileId,proto3" json:"payments_profile_id,omitempty"`
	// Output only. The issue date in yyyy-mm-dd format. It appears on the invoice PDF as
	// either "Issue date" or "Invoice date".
	IssueDate *wrappers.StringValue `protobuf:"bytes,7,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	// Output only. The due date in yyyy-mm-dd format.
	DueDate *wrappers.StringValue `protobuf:"bytes,8,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	// Output only. The service period date range of this invoice. The end date is inclusive.
	ServiceDateRange *common.DateRange `protobuf:"bytes,9,opt,name=service_date_range,json=serviceDateRange,proto3" json:"service_date_range,omitempty"`
	// Output only. The currency code. All costs are returned in this currency. A subset of the
	// currency codes derived from the ISO 4217 standard is supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,10,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Output only. The total amount of invoice level adjustments. These adjustments are made
	// on the invoice, not on a specific account budget.
	InvoiceLevelAdjustmentsMicros *wrappers.Int64Value `protobuf:"bytes,11,opt,name=invoice_level_adjustments_micros,json=invoiceLevelAdjustmentsMicros,proto3" json:"invoice_level_adjustments_micros,omitempty"`
	// Output only. The pretax subtotal amount, in micros. This equals the sum of the
	// AccountBudgetSummary subtotal amounts, plus the invoice level adjustments.
	SubtotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,12,opt,name=subtotal_amount_micros,json=subtotalAmountMicros,proto3" json:"subtotal_amount_micros,omitempty"`
	// Output only. The sum of all taxes on the invoice, in micros. This equals the sum of the
	// AccountBudgetSummary tax amounts, plus taxes not associated with a specific
	// account budget.
	TaxAmountMicros *wrappers.Int64Value `protobuf:"bytes,13,opt,name=tax_amount_micros,json=taxAmountMicros,proto3" json:"tax_amount_micros,omitempty"`
	// Output only. The total amount, in micros. This equals the sum of the invoice subtotal
	// amount and the invoice tax amount.
	TotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,14,opt,name=total_amount_micros,json=totalAmountMicros,proto3" json:"total_amount_micros,omitempty"`
	// Output only. The resource name of the original invoice corrected, wrote off, or canceled
	// by this invoice, if applicable. If `corrected_invoice` is set,
	// `replaced_invoices` will not be set.
	//
	// Invoice resource names have the form:
	//
	// `customers/{customer_id}/invoices/{invoice_id}`
	CorrectedInvoice *wrappers.StringValue `protobuf:"bytes,15,opt,name=corrected_invoice,json=correctedInvoice,proto3" json:"corrected_invoice,omitempty"`
	// Output only. The resource name of the original invoice(s) being rebilled or replaced by
	// this invoice, if applicable. There might be multiple replaced invoices due
	// to invoice consolidation. The replaced invoices may not belong to the same
	// payments account. If `replaced_invoices` is set, `corrected_invoice` will
	// not be set.
	// Invoice resource names have the form:
	//
	// `customers/{customer_id}/invoices/{invoice_id}`
	ReplacedInvoices []*wrappers.StringValue `protobuf:"bytes,16,rep,name=replaced_invoices,json=replacedInvoices,proto3" json:"replaced_invoices,omitempty"`
	// Output only. The URL to a PDF copy of the invoice. Users need to pass in their OAuth
	// token to request the PDF with this URL.
	PdfUrl *wrappers.StringValue `protobuf:"bytes,17,opt,name=pdf_url,json=pdfUrl,proto3" json:"pdf_url,omitempty"`
	// Output only. The list of summarized account budget information associated with this
	// invoice.
	AccountBudgetSummaries []*Invoice_AccountBudgetSummary `protobuf:"bytes,18,rep,name=account_budget_summaries,json=accountBudgetSummaries,proto3" json:"account_budget_summaries,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                        `json:"-"`
	XXX_unrecognized       []byte                          `json:"-"`
	XXX_sizecache          int32                           `json:"-"`
}

func (m *Invoice) Reset()         { *m = Invoice{} }
func (m *Invoice) String() string { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()    {}
func (*Invoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b8f0bbb4092e161, []int{0}
}

func (m *Invoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice.Unmarshal(m, b)
}
func (m *Invoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice.Marshal(b, m, deterministic)
}
func (m *Invoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice.Merge(m, src)
}
func (m *Invoice) XXX_Size() int {
	return xxx_messageInfo_Invoice.Size(m)
}
func (m *Invoice) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice proto.InternalMessageInfo

func (m *Invoice) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Invoice) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Invoice) GetType() enums.InvoiceTypeEnum_InvoiceType {
	if m != nil {
		return m.Type
	}
	return enums.InvoiceTypeEnum_UNSPECIFIED
}

func (m *Invoice) GetBillingSetup() *wrappers.StringValue {
	if m != nil {
		return m.BillingSetup
	}
	return nil
}

func (m *Invoice) GetPaymentsAccountId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountId
	}
	return nil
}

func (m *Invoice) GetPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileId
	}
	return nil
}

func (m *Invoice) GetIssueDate() *wrappers.StringValue {
	if m != nil {
		return m.IssueDate
	}
	return nil
}

func (m *Invoice) GetDueDate() *wrappers.StringValue {
	if m != nil {
		return m.DueDate
	}
	return nil
}

func (m *Invoice) GetServiceDateRange() *common.DateRange {
	if m != nil {
		return m.ServiceDateRange
	}
	return nil
}

func (m *Invoice) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *Invoice) GetInvoiceLevelAdjustmentsMicros() *wrappers.Int64Value {
	if m != nil {
		return m.InvoiceLevelAdjustmentsMicros
	}
	return nil
}

func (m *Invoice) GetSubtotalAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.SubtotalAmountMicros
	}
	return nil
}

func (m *Invoice) GetTaxAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TaxAmountMicros
	}
	return nil
}

func (m *Invoice) GetTotalAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TotalAmountMicros
	}
	return nil
}

func (m *Invoice) GetCorrectedInvoice() *wrappers.StringValue {
	if m != nil {
		return m.CorrectedInvoice
	}
	return nil
}

func (m *Invoice) GetReplacedInvoices() []*wrappers.StringValue {
	if m != nil {
		return m.ReplacedInvoices
	}
	return nil
}

func (m *Invoice) GetPdfUrl() *wrappers.StringValue {
	if m != nil {
		return m.PdfUrl
	}
	return nil
}

func (m *Invoice) GetAccountBudgetSummaries() []*Invoice_AccountBudgetSummary {
	if m != nil {
		return m.AccountBudgetSummaries
	}
	return nil
}

// Represents a summarized account budget billable cost.
type Invoice_AccountBudgetSummary struct {
	// Output only. The resource name of the customer associated with this account budget.
	// This contains the customer ID, which appears on the invoice PDF as
	// "Account ID".
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	Customer *wrappers.StringValue `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	// Output only. The descriptive name of the account budget’s customer. It appears on the
	// invoice PDF as "Account".
	CustomerDescriptiveName *wrappers.StringValue `protobuf:"bytes,2,opt,name=customer_descriptive_name,json=customerDescriptiveName,proto3" json:"customer_descriptive_name,omitempty"`
	// Output only. The resource name of the account budget associated with this summarized
	// billable cost.
	// AccountBudget resource names have the form:
	//
	// `customers/{customer_id}/accountBudgets/{account_budget_id}`
	AccountBudget *wrappers.StringValue `protobuf:"bytes,3,opt,name=account_budget,json=accountBudget,proto3" json:"account_budget,omitempty"`
	// Output only. The name of the account budget. It appears on the invoice PDF as "Account
	// budget".
	AccountBudgetName *wrappers.StringValue `protobuf:"bytes,4,opt,name=account_budget_name,json=accountBudgetName,proto3" json:"account_budget_name,omitempty"`
	// Output only. The purchase order number of the account budget. It appears on the
	// invoice PDF as "Purchase order".
	PurchaseOrderNumber *wrappers.StringValue `protobuf:"bytes,5,opt,name=purchase_order_number,json=purchaseOrderNumber,proto3" json:"purchase_order_number,omitempty"`
	// Output only. The pretax subtotal amount attributable to this budget during the service
	// period, in micros.
	SubtotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,6,opt,name=subtotal_amount_micros,json=subtotalAmountMicros,proto3" json:"subtotal_amount_micros,omitempty"`
	// Output only. The tax amount attributable to this budget during the service period, in
	// micros.
	TaxAmountMicros *wrappers.Int64Value `protobuf:"bytes,7,opt,name=tax_amount_micros,json=taxAmountMicros,proto3" json:"tax_amount_micros,omitempty"`
	// Output only. The total amount attributable to this budget during the service period,
	// in micros. This equals the sum of the account budget subtotal amount and
	// the account budget tax amount.
	TotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,8,opt,name=total_amount_micros,json=totalAmountMicros,proto3" json:"total_amount_micros,omitempty"`
	// Output only. The billable activity date range of the account budget, within the
	// service date range of this invoice. The end date is inclusive. This can
	// be different from the account budget's start and end time.
	BillableActivityDateRange *common.DateRange `protobuf:"bytes,9,opt,name=billable_activity_date_range,json=billableActivityDateRange,proto3" json:"billable_activity_date_range,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}          `json:"-"`
	XXX_unrecognized          []byte            `json:"-"`
	XXX_sizecache             int32             `json:"-"`
}

func (m *Invoice_AccountBudgetSummary) Reset()         { *m = Invoice_AccountBudgetSummary{} }
func (m *Invoice_AccountBudgetSummary) String() string { return proto.CompactTextString(m) }
func (*Invoice_AccountBudgetSummary) ProtoMessage()    {}
func (*Invoice_AccountBudgetSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b8f0bbb4092e161, []int{0, 0}
}

func (m *Invoice_AccountBudgetSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice_AccountBudgetSummary.Unmarshal(m, b)
}
func (m *Invoice_AccountBudgetSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice_AccountBudgetSummary.Marshal(b, m, deterministic)
}
func (m *Invoice_AccountBudgetSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice_AccountBudgetSummary.Merge(m, src)
}
func (m *Invoice_AccountBudgetSummary) XXX_Size() int {
	return xxx_messageInfo_Invoice_AccountBudgetSummary.Size(m)
}
func (m *Invoice_AccountBudgetSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice_AccountBudgetSummary.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice_AccountBudgetSummary proto.InternalMessageInfo

func (m *Invoice_AccountBudgetSummary) GetCustomer() *wrappers.StringValue {
	if m != nil {
		return m.Customer
	}
	return nil
}

func (m *Invoice_AccountBudgetSummary) GetCustomerDescriptiveName() *wrappers.StringValue {
	if m != nil {
		return m.CustomerDescriptiveName
	}
	return nil
}

func (m *Invoice_AccountBudgetSummary) GetAccountBudget() *wrappers.StringValue {
	if m != nil {
		return m.AccountBudget
	}
	return nil
}

func (m *Invoice_AccountBudgetSummary) GetAccountBudgetName() *wrappers.StringValue {
	if m != nil {
		return m.AccountBudgetName
	}
	return nil
}

func (m *Invoice_AccountBudgetSummary) GetPurchaseOrderNumber() *wrappers.StringValue {
	if m != nil {
		return m.PurchaseOrderNumber
	}
	return nil
}

func (m *Invoice_AccountBudgetSummary) GetSubtotalAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.SubtotalAmountMicros
	}
	return nil
}

func (m *Invoice_AccountBudgetSummary) GetTaxAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TaxAmountMicros
	}
	return nil
}

func (m *Invoice_AccountBudgetSummary) GetTotalAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TotalAmountMicros
	}
	return nil
}

func (m *Invoice_AccountBudgetSummary) GetBillableActivityDateRange() *common.DateRange {
	if m != nil {
		return m.BillableActivityDateRange
	}
	return nil
}

func init() {
	proto.RegisterType((*Invoice)(nil), "google.ads.googleads.v2.resources.Invoice")
	proto.RegisterType((*Invoice_AccountBudgetSummary)(nil), "google.ads.googleads.v2.resources.Invoice.AccountBudgetSummary")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/invoice.proto", fileDescriptor_2b8f0bbb4092e161)
}

var fileDescriptor_2b8f0bbb4092e161 = []byte{
	// 954 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x41, 0x73, 0xdb, 0x44,
	0x14, 0xc7, 0xc7, 0x4e, 0x1b, 0x27, 0xdb, 0x24, 0x8d, 0x95, 0x52, 0x94, 0x10, 0xc0, 0xed, 0x4c,
	0x87, 0xc0, 0x41, 0x02, 0xc3, 0x70, 0x10, 0x30, 0x20, 0x53, 0x28, 0x81, 0xd6, 0x84, 0x84, 0xfa,
	0xc0, 0x98, 0xd1, 0xac, 0xb5, 0xcf, 0xea, 0x32, 0xd2, 0xae, 0x66, 0x77, 0x65, 0xea, 0xe9, 0x94,
	0xef, 0xc0, 0x89, 0x3b, 0x47, 0x3e, 0x0a, 0x9f, 0xa2, 0xe7, 0x7e, 0x03, 0x38, 0x31, 0x5a, 0xed,
	0xca, 0x71, 0xd3, 0x34, 0x02, 0x7a, 0x7b, 0xce, 0xbe, 0xff, 0x6f, 0xff, 0x79, 0xbb, 0xef, 0xad,
	0x90, 0x9f, 0x70, 0x9e, 0xa4, 0xe0, 0x63, 0x22, 0x4d, 0x58, 0x46, 0xb3, 0xbe, 0x2f, 0x40, 0xf2,
	0x42, 0xc4, 0x20, 0x7d, 0xca, 0x66, 0x9c, 0xc6, 0xe0, 0xe5, 0x82, 0x2b, 0xee, 0xdc, 0xa8, 0xb2,
	0x3c, 0x4c, 0xa4, 0x57, 0x0b, 0xbc, 0x59, 0xdf, 0xab, 0x05, 0x7b, 0xef, 0x9c, 0xc7, 0x8c, 0x79,
	0x96, 0x71, 0xe6, 0x13, 0xac, 0x40, 0x56, 0xb8, 0xbd, 0x77, 0xcf, 0xcb, 0x05, 0x56, 0x64, 0xf5,
	0xde, 0x91, 0x9a, 0xe7, 0xc6, 0xc0, 0xde, 0x9b, 0x56, 0x91, 0x53, 0x7f, 0x4a, 0x21, 0x25, 0xd1,
	0x04, 0x1e, 0xe0, 0x19, 0xe5, 0xc2, 0x24, 0xec, 0x9e, 0x4a, 0xb0, 0xa6, 0xcc, 0xd2, 0x1b, 0x66,
	0x49, 0xff, 0x9a, 0x14, 0x53, 0xff, 0x67, 0x81, 0xf3, 0x1c, 0x84, 0x75, 0xb3, 0x7f, 0x4a, 0x8a,
	0x19, 0xe3, 0x0a, 0x2b, 0xca, 0x99, 0x59, 0xbd, 0xf9, 0x9b, 0x83, 0x3a, 0x87, 0x95, 0x21, 0xe7,
	0x1e, 0xda, 0xb4, 0xec, 0x88, 0xe1, 0x0c, 0xdc, 0x56, 0xaf, 0x75, 0xb0, 0x3e, 0x38, 0x78, 0x12,
	0x5e, 0xfe, 0x3b, 0xbc, 0x89, 0x7a, 0x8b, 0xd2, 0x98, 0x28, 0xa7, 0xd2, 0x8b, 0x79, 0xe6, 0x1b,
	0xc0, 0xf1, 0x86, 0x95, 0x0f, 0x71, 0x06, 0xce, 0x7b, 0xa8, 0x4d, 0x89, 0xdb, 0xee, 0xb5, 0x0e,
	0xae, 0xf4, 0xf7, 0x8d, 0xc4, 0xb3, 0x2e, 0xbd, 0x13, 0x25, 0x28, 0x4b, 0x46, 0x38, 0x2d, 0x60,
	0xb0, 0xf2, 0x24, 0x5c, 0x39, 0x6e, 0x53, 0xe2, 0x9c, 0xa0, 0x4b, 0x65, 0x55, 0xdc, 0x95, 0x5e,
	0xeb, 0x60, 0xab, 0x1f, 0x78, 0xe7, 0x9d, 0x8b, 0x2e, 0xa4, 0x67, 0xb6, 0xfd, 0x7e, 0x9e, 0xc3,
	0x17, 0xac, 0xc8, 0x4e, 0xff, 0xae, 0x90, 0x1a, 0xe6, 0x7c, 0x89, 0x36, 0x27, 0x34, 0x4d, 0x29,
	0x4b, 0x22, 0x09, 0xaa, 0xc8, 0xdd, 0x4b, 0x4d, 0x2d, 0x6d, 0x18, 0xdd, 0x49, 0x29, 0x73, 0xbe,
	0x43, 0x3b, 0x39, 0x9e, 0x67, 0xc0, 0x94, 0x8c, 0x70, 0x1c, 0xf3, 0x82, 0xa9, 0x88, 0x12, 0xf7,
	0x72, 0x53, 0x5a, 0xd7, 0xaa, 0xc3, 0x4a, 0x7c, 0x48, 0x96, 0x90, 0xb9, 0xe0, 0x53, 0x9a, 0x42,
	0x89, 0x5c, 0xfd, 0xd7, 0xc8, 0xa3, 0x4a, 0x7c, 0x48, 0x9c, 0xcf, 0x10, 0xa2, 0x52, 0x16, 0x10,
	0x95, 0x37, 0xd2, 0xed, 0x34, 0x25, 0xad, 0x6b, 0xd1, 0x6d, 0xac, 0xc0, 0xf9, 0x18, 0xad, 0x11,
	0xab, 0x5f, 0x6b, 0xaa, 0xef, 0x10, 0xa3, 0x1e, 0x23, 0x47, 0x82, 0x98, 0x95, 0x17, 0xbc, 0x24,
	0x44, 0x02, 0xb3, 0x04, 0xdc, 0x75, 0xcd, 0x79, 0xfb, 0xdc, 0x03, 0xad, 0xba, 0xc8, 0x2b, 0x09,
	0xc7, 0xa5, 0xa0, 0x82, 0x6e, 0x1b, 0x52, 0xfd, 0xe7, 0xf2, 0x2c, 0xe3, 0x42, 0x08, 0x60, 0xf1,
	0x3c, 0x8a, 0x39, 0x01, 0x17, 0x35, 0x3e, 0x4b, 0xab, 0xfb, 0x9c, 0x13, 0x70, 0x28, 0xea, 0xd9,
	0x36, 0x4c, 0x61, 0x06, 0x69, 0x84, 0xc9, 0x4f, 0x85, 0x54, 0xd5, 0x49, 0x64, 0x34, 0x16, 0x5c,
	0xba, 0x57, 0x34, 0xfa, 0xb5, 0x33, 0xe8, 0x43, 0xa6, 0x3e, 0xfc, 0xe0, 0x14, 0xf9, 0x75, 0x43,
	0xba, 0x5b, 0x82, 0xc2, 0x05, 0xe7, 0x9e, 0xc6, 0x38, 0x23, 0x74, 0x5d, 0x16, 0x13, 0xc5, 0x15,
	0x4e, 0x23, 0x9c, 0xe9, 0x5b, 0x63, 0x36, 0xd8, 0x68, 0xb8, 0xc1, 0x35, 0xab, 0x0f, 0xb5, 0xdc,
	0x70, 0xef, 0xa2, 0xae, 0xc2, 0x0f, 0x9f, 0x41, 0x6e, 0x36, 0x44, 0x5e, 0x55, 0xf8, 0xe1, 0x12,
	0xed, 0x08, 0xed, 0x3c, 0xcf, 0xe2, 0x56, 0x43, 0x5e, 0xf7, 0xac, 0xbf, 0x21, 0xea, 0xc6, 0x5c,
	0x08, 0x88, 0x15, 0x90, 0xc8, 0x94, 0xc8, 0xbd, 0xda, 0xf4, 0xb8, 0xb6, 0x6b, 0xad, 0x9d, 0x4e,
	0x43, 0xd4, 0x15, 0x90, 0xa7, 0x38, 0x5e, 0xe0, 0xa4, 0xbb, 0xdd, 0x5b, 0x69, 0xc8, 0xb3, 0x5a,
	0x83, 0x93, 0x4e, 0x80, 0x3a, 0x39, 0x99, 0x46, 0x85, 0x48, 0xdd, 0x6e, 0x53, 0x57, 0xab, 0x39,
	0x99, 0xde, 0x17, 0xa9, 0xf3, 0x0b, 0x72, 0xed, 0x04, 0x98, 0x14, 0x24, 0x01, 0x15, 0xc9, 0x22,
	0xcb, 0xb0, 0xa0, 0x20, 0x5d, 0x47, 0x5b, 0xfa, 0xd4, 0xbb, 0xf0, 0x4d, 0xb1, 0xf3, 0xca, 0x33,
	0xf3, 0x60, 0xa0, 0x49, 0x27, 0x1a, 0x34, 0xaf, 0xf6, 0xbb, 0x8e, 0xcf, 0x2e, 0x51, 0x90, 0x7b,
	0xbf, 0xae, 0xa2, 0x6b, 0xcf, 0x53, 0x39, 0x9f, 0xa0, 0xb5, 0xb8, 0x90, 0x8a, 0x67, 0x20, 0xf4,
	0xf4, 0x6e, 0xf4, 0x5f, 0xd5, 0x12, 0xe7, 0x47, 0xb4, 0x6b, 0xe3, 0x88, 0x80, 0x8c, 0x05, 0xcd,
	0x15, 0x9d, 0x99, 0xd7, 0xa0, 0xf1, 0x24, 0x7f, 0xd5, 0x32, 0x6e, 0x2f, 0x10, 0xfa, 0x45, 0xf8,
	0x0a, 0x6d, 0x2d, 0x97, 0x4d, 0x0f, 0xfa, 0x46, 0xcc, 0xcd, 0xa5, 0x4a, 0x94, 0x83, 0xf3, 0x99,
	0x03, 0xd0, 0x16, 0x1b, 0x4f, 0xf6, 0xee, 0x12, 0x4e, 0x9b, 0xbb, 0x8f, 0x5e, 0xc9, 0x0b, 0x11,
	0x3f, 0xc0, 0x12, 0x22, 0x2e, 0x08, 0x88, 0x88, 0x15, 0xd9, 0x04, 0x44, 0xf3, 0x01, 0xbf, 0x63,
	0xf5, 0xdf, 0x96, 0xf2, 0xa1, 0x56, 0xbf, 0xa0, 0xfd, 0x57, 0x5f, 0x7e, 0xfb, 0x77, 0x5e, 0x72,
	0xfb, 0xaf, 0xfd, 0xf7, 0xf6, 0xe7, 0x68, 0xbf, 0x7c, 0x3d, 0xf1, 0x24, 0x85, 0x08, 0xc7, 0x8a,
	0xce, 0xa8, 0x9a, 0xff, 0xff, 0x17, 0x61, 0xd7, 0x32, 0x43, 0x83, 0xac, 0xd7, 0x83, 0xe1, 0xd3,
	0xf0, 0x9b, 0x8b, 0xbf, 0x51, 0x9c, 0xb7, 0xec, 0xe5, 0x94, 0xfe, 0x23, 0x1b, 0x3e, 0xb6, 0xdf,
	0x64, 0xd2, 0x7f, 0x64, 0xa2, 0xc7, 0x83, 0xbf, 0x5a, 0xe8, 0x56, 0xcc, 0xb3, 0x8b, 0xfb, 0x78,
	0xb0, 0x61, 0xd8, 0x47, 0x65, 0x79, 0x8e, 0x5a, 0x3f, 0x7c, 0x6d, 0x24, 0x09, 0x4f, 0x31, 0x4b,
	0x3c, 0x2e, 0x12, 0x3f, 0x01, 0xa6, 0x8b, 0xe7, 0x2f, 0x5c, 0xbd, 0xe0, 0xf3, 0xf4, 0xa3, 0x3a,
	0xfa, 0xbd, 0xbd, 0x72, 0x27, 0x0c, 0xff, 0x68, 0xdf, 0xb8, 0x53, 0x21, 0x43, 0x22, 0xbd, 0x2a,
	0x2c, 0xa3, 0x51, 0xdf, 0x3b, 0xb6, 0x99, 0x7f, 0xda, 0x9c, 0x71, 0x48, 0xe4, 0xb8, 0xce, 0x19,
	0x8f, 0xfa, 0xe3, 0x3a, 0xe7, 0x69, 0xfb, 0x56, 0xb5, 0x10, 0x04, 0x21, 0x91, 0x41, 0x50, 0x67,
	0x05, 0xc1, 0xa8, 0x1f, 0x04, 0x75, 0xde, 0x64, 0x55, 0x9b, 0x7d, 0xff, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1b, 0x7b, 0x88, 0x2e, 0x4a, 0x0b, 0x00, 0x00,
}
