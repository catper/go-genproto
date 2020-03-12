// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/invoice.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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
	// The resource name of the invoice. Multiple customers can share a given
	// invoice, so multiple resource names may point to the same invoice.
	// Invoice resource names have the form:
	//
	// `customers/{customer_id}/invoices/{invoice_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the invoice. It appears on the invoice PDF as "Invoice number".
	Id *wrappers.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The type of invoice.
	Type enums.InvoiceTypeEnum_InvoiceType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.InvoiceTypeEnum_InvoiceType" json:"type,omitempty"`
	// The resource name of this invoice’s billing setup.
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	BillingSetup *wrappers.StringValue `protobuf:"bytes,4,opt,name=billing_setup,json=billingSetup,proto3" json:"billing_setup,omitempty"`
	// A 16 digit ID used to identify the payments account associated with the
	// billing setup, e.g. "1234-5678-9012-3456". It appears on the invoice PDF as
	// "Billing Account Number".
	PaymentsAccountId *wrappers.StringValue `protobuf:"bytes,5,opt,name=payments_account_id,json=paymentsAccountId,proto3" json:"payments_account_id,omitempty"`
	// A 12 digit ID used to identify the payments profile associated with the
	// billing setup, e.g. "1234-5678-9012". It appears on the invoice PDF as
	// "Billing ID".
	PaymentsProfileId *wrappers.StringValue `protobuf:"bytes,6,opt,name=payments_profile_id,json=paymentsProfileId,proto3" json:"payments_profile_id,omitempty"`
	// The issue date in yyyy-mm-dd format. It appears on the invoice PDF as
	// either "Issue date" or "Invoice date".
	IssueDate *wrappers.StringValue `protobuf:"bytes,7,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	// The due date in yyyy-mm-dd format.
	DueDate *wrappers.StringValue `protobuf:"bytes,8,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	// The service period date range of this invoice. The end date is inclusive.
	ServiceDateRange *common.DateRange `protobuf:"bytes,9,opt,name=service_date_range,json=serviceDateRange,proto3" json:"service_date_range,omitempty"`
	// The currency code. All costs are returned in this currency. A subset of the
	// currency codes derived from the ISO 4217 standard is supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,10,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The total amount of invoice level adjustments. These adjustments are made
	// on the invoice, not on a specific account budget.
	InvoiceLevelAdjustmentsMicros *wrappers.Int64Value `protobuf:"bytes,11,opt,name=invoice_level_adjustments_micros,json=invoiceLevelAdjustmentsMicros,proto3" json:"invoice_level_adjustments_micros,omitempty"`
	// The pretax subtotal amount, in micros. This equals the sum of the
	// AccountBudgetSummary subtotal amounts, plus the invoice level adjustments.
	SubtotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,12,opt,name=subtotal_amount_micros,json=subtotalAmountMicros,proto3" json:"subtotal_amount_micros,omitempty"`
	// The sum of all taxes on the invoice, in micros. This equals the sum of the
	// AccountBudgetSummary tax amounts, plus taxes not associated with a specific
	// account budget.
	TaxAmountMicros *wrappers.Int64Value `protobuf:"bytes,13,opt,name=tax_amount_micros,json=taxAmountMicros,proto3" json:"tax_amount_micros,omitempty"`
	// The total amount, in micros. This equals the sum of the invoice subtotal
	// amount and the invoice tax amount.
	TotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,14,opt,name=total_amount_micros,json=totalAmountMicros,proto3" json:"total_amount_micros,omitempty"`
	// The resource name of the original invoice corrected, wrote off, or canceled
	// by this invoice, if applicable. If `corrected_invoice` is set,
	// `replaced_invoices` will not be set.
	//
	// Invoice resource names have the form:
	//
	// `customers/{customer_id}/invoices/{invoice_id}`
	CorrectedInvoice *wrappers.StringValue `protobuf:"bytes,15,opt,name=corrected_invoice,json=correctedInvoice,proto3" json:"corrected_invoice,omitempty"`
	// The resource name of the original invoice(s) being rebilled or replaced by
	// this invoice, if applicable. There might be multiple replaced invoices due
	// to invoice consolidation. The replaced invoices may not belong to the same
	// payments account. If `replaced_invoices` is set, `corrected_invoice` will
	// not be set.
	// Invoice resource names have the form:
	//
	// `customers/{customer_id}/invoices/{invoice_id}`
	ReplacedInvoices []*wrappers.StringValue `protobuf:"bytes,16,rep,name=replaced_invoices,json=replacedInvoices,proto3" json:"replaced_invoices,omitempty"`
	// The URL to a PDF copy of the invoice. Users need to pass in their OAuth
	// token to request the PDF with this URL.
	PdfUrl *wrappers.StringValue `protobuf:"bytes,17,opt,name=pdf_url,json=pdfUrl,proto3" json:"pdf_url,omitempty"`
	// The list of summarized account budget information associated with this
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
	return fileDescriptor_52f5271ace0f4785, []int{0}
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
	// The resource name of the customer associated with this account budget.
	// This contains the customer ID, which appears on the invoice PDF as
	// "Account ID".
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	Customer *wrappers.StringValue `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	// The descriptive name of the account budget’s customer. It appears on the
	// invoice PDF as "Account".
	CustomerDescriptiveName *wrappers.StringValue `protobuf:"bytes,2,opt,name=customer_descriptive_name,json=customerDescriptiveName,proto3" json:"customer_descriptive_name,omitempty"`
	// The resource name of the account budget associated with this summarized
	// billable cost.
	// AccountBudget resource names have the form:
	//
	// `customers/{customer_id}/accountBudgets/{account_budget_id}`
	AccountBudget *wrappers.StringValue `protobuf:"bytes,3,opt,name=account_budget,json=accountBudget,proto3" json:"account_budget,omitempty"`
	// The name of the account budget. It appears on the invoice PDF as "Account
	// budget".
	AccountBudgetName *wrappers.StringValue `protobuf:"bytes,4,opt,name=account_budget_name,json=accountBudgetName,proto3" json:"account_budget_name,omitempty"`
	// The purchase order number of the account budget. It appears on the
	// invoice PDF as "Purchase order".
	PurchaseOrderNumber *wrappers.StringValue `protobuf:"bytes,5,opt,name=purchase_order_number,json=purchaseOrderNumber,proto3" json:"purchase_order_number,omitempty"`
	// The pretax subtotal amount attributable to this budget during the service
	// period, in micros.
	SubtotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,6,opt,name=subtotal_amount_micros,json=subtotalAmountMicros,proto3" json:"subtotal_amount_micros,omitempty"`
	// The tax amount attributable to this budget during the service period, in
	// micros.
	TaxAmountMicros *wrappers.Int64Value `protobuf:"bytes,7,opt,name=tax_amount_micros,json=taxAmountMicros,proto3" json:"tax_amount_micros,omitempty"`
	// The total amount attributable to this budget during the service period,
	// in micros. This equals the sum of the account budget subtotal amount and
	// the account budget tax amount.
	TotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,8,opt,name=total_amount_micros,json=totalAmountMicros,proto3" json:"total_amount_micros,omitempty"`
	// The billable activity date range of the account budget, within the
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
	return fileDescriptor_52f5271ace0f4785, []int{0, 0}
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
	proto.RegisterType((*Invoice)(nil), "google.ads.googleads.v3.resources.Invoice")
	proto.RegisterType((*Invoice_AccountBudgetSummary)(nil), "google.ads.googleads.v3.resources.Invoice.AccountBudgetSummary")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/invoice.proto", fileDescriptor_52f5271ace0f4785)
}

var fileDescriptor_52f5271ace0f4785 = []byte{
	// 916 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xc7, 0xb5, 0x9b, 0x36, 0x1f, 0xd3, 0x24, 0xcd, 0x3a, 0xa5, 0x38, 0x21, 0xa0, 0x2d, 0xa8,
	0x22, 0x20, 0x64, 0xa3, 0x2e, 0x5f, 0x72, 0x2f, 0x90, 0xd3, 0xa2, 0x68, 0x69, 0x09, 0x61, 0x03,
	0x01, 0xa1, 0x48, 0xd6, 0xac, 0xe7, 0xc4, 0x4c, 0x65, 0xcf, 0x58, 0x33, 0xe3, 0xa5, 0xab, 0xaa,
	0x2f, 0xc3, 0x25, 0x8f, 0xc2, 0x73, 0x70, 0x55, 0x9e, 0x80, 0x3b, 0x34, 0xe3, 0x19, 0x6f, 0xb7,
	0x69, 0xba, 0x06, 0x7a, 0x77, 0xd6, 0x73, 0xfe, 0x3f, 0x9f, 0x39, 0x33, 0xe7, 0xbf, 0x46, 0x61,
	0xc6, 0x79, 0x96, 0x43, 0x88, 0x89, 0xb4, 0xa1, 0x8e, 0x26, 0x83, 0x50, 0x80, 0xe4, 0x95, 0x48,
	0x41, 0x86, 0x94, 0x4d, 0x38, 0x4d, 0x21, 0x28, 0x05, 0x57, 0xdc, 0xbb, 0x55, 0x67, 0x05, 0x98,
	0xc8, 0xa0, 0x11, 0x04, 0x93, 0x41, 0xd0, 0x08, 0x76, 0x3f, 0xbc, 0x8c, 0x99, 0xf2, 0xa2, 0xe0,
	0x2c, 0x24, 0x58, 0x81, 0xac, 0x71, 0xbb, 0x1f, 0x5f, 0x96, 0x0b, 0xac, 0x2a, 0x9a, 0x77, 0x27,
	0x6a, 0x5a, 0xda, 0x02, 0x76, 0x77, 0x9c, 0xa2, 0xa4, 0x4d, 0x91, 0x76, 0xe9, 0x1d, 0xbb, 0x64,
	0x7e, 0x8d, 0xab, 0xf3, 0xf0, 0x57, 0x81, 0xcb, 0x12, 0x84, 0x7b, 0xd9, 0xde, 0x73, 0x52, 0xcc,
	0x18, 0x57, 0x58, 0x51, 0xce, 0xec, 0xea, 0xbb, 0x7f, 0x6d, 0xa1, 0x95, 0x61, 0xfd, 0x3e, 0xef,
	0x3d, 0xb4, 0xe1, 0xd8, 0x09, 0xc3, 0x05, 0xf8, 0x9d, 0x7e, 0x67, 0x7f, 0x6d, 0xb4, 0xee, 0x1e,
	0x1e, 0xe1, 0x02, 0xbc, 0x8f, 0x50, 0x97, 0x12, 0xbf, 0xdb, 0xef, 0xec, 0x5f, 0xbb, 0xb3, 0x67,
	0x9b, 0x11, 0xb8, 0x77, 0x07, 0x27, 0x4a, 0x50, 0x96, 0x9d, 0xe2, 0xbc, 0x82, 0x51, 0x97, 0x12,
	0xef, 0x08, 0x5d, 0xd1, 0xbb, 0xf0, 0x97, 0xfa, 0x9d, 0xfd, 0xcd, 0x3b, 0x51, 0x70, 0x59, 0x1f,
	0xcd, 0xc6, 0x03, 0x5b, 0xc8, 0xf7, 0xd3, 0x12, 0xbe, 0x62, 0x55, 0xf1, 0xfc, 0xef, 0x91, 0xe1,
	0x78, 0x31, 0xda, 0x18, 0xd3, 0x3c, 0xa7, 0x2c, 0x4b, 0x24, 0xa8, 0xaa, 0xf4, 0xaf, 0xb4, 0x28,
	0x64, 0xdd, 0x4a, 0x4e, 0xb4, 0xc2, 0x7b, 0x88, 0xb6, 0x4b, 0x3c, 0x2d, 0x80, 0x29, 0x99, 0xe0,
	0x34, 0xe5, 0x15, 0x53, 0x09, 0x25, 0xfe, 0xd5, 0x16, 0xa0, 0x9e, 0x13, 0xc6, 0xb5, 0x6e, 0x48,
	0xe6, 0x68, 0xa5, 0xe0, 0xe7, 0x34, 0x07, 0x4d, 0x5b, 0xfe, 0x37, 0xb4, 0xe3, 0x5a, 0x37, 0x24,
	0xde, 0x5d, 0x84, 0xa8, 0x94, 0x15, 0x24, 0xfa, 0xb6, 0xf8, 0x2b, 0x2d, 0x20, 0x6b, 0x26, 0xff,
	0x3e, 0x56, 0xe0, 0x7d, 0x8e, 0x56, 0x89, 0x93, 0xae, 0xb6, 0x90, 0xae, 0x10, 0x2b, 0xfc, 0x11,
	0x79, 0x12, 0xc4, 0x44, 0x5f, 0x39, 0x2d, 0x4e, 0x04, 0x66, 0x19, 0xf8, 0x6b, 0x06, 0xf1, 0xc1,
	0xa5, 0x47, 0x56, 0xdf, 0xeb, 0x40, 0x13, 0x46, 0x5a, 0x30, 0xda, 0xb2, 0x90, 0xe6, 0x89, 0x3e,
	0xad, 0xb4, 0x12, 0x02, 0x58, 0x3a, 0x4d, 0x52, 0x4e, 0xc0, 0x47, 0x6d, 0x4e, 0xcb, 0x49, 0xee,
	0x71, 0x02, 0x1e, 0x41, 0x7d, 0x37, 0x0e, 0x39, 0x4c, 0x20, 0x4f, 0x30, 0x79, 0x54, 0x49, 0x55,
	0x37, 0xbc, 0xa0, 0xa9, 0xe0, 0xd2, 0xbf, 0x66, 0xa8, 0x6f, 0x5d, 0xa0, 0x0e, 0x99, 0xfa, 0xec,
	0x93, 0x1a, 0xfa, 0xb6, 0x85, 0x3c, 0xd4, 0x8c, 0x78, 0x86, 0xf8, 0xc6, 0x10, 0xbc, 0xef, 0xd0,
	0x4d, 0x59, 0x8d, 0x15, 0x57, 0x38, 0x4f, 0x70, 0x61, 0xae, 0x84, 0x65, 0xaf, 0x2f, 0x66, 0xdf,
	0x70, 0xd2, 0xd8, 0x28, 0x2d, 0xf2, 0x10, 0xf5, 0x14, 0x7e, 0xfc, 0x02, 0x6d, 0x63, 0x31, 0xed,
	0xba, 0xc2, 0x8f, 0xe7, 0x40, 0x0f, 0xd0, 0xf6, 0xcb, 0x0a, 0xdb, 0x5c, 0x8c, 0xea, 0x5d, 0xac,
	0x6a, 0x88, 0x7a, 0x29, 0x17, 0x02, 0x52, 0x05, 0x24, 0xb1, 0x3d, 0xf1, 0xaf, 0xb7, 0x38, 0x95,
	0xad, 0x46, 0xe6, 0xdc, 0x62, 0x88, 0x7a, 0x02, 0xca, 0x1c, 0xa7, 0x33, 0x92, 0xf4, 0xb7, 0xfa,
	0x4b, 0x8b, 0x51, 0x4e, 0x66, 0x49, 0xd2, 0xfb, 0x14, 0xad, 0x94, 0xe4, 0x3c, 0xa9, 0x44, 0xee,
	0xf7, 0x5a, 0xd4, 0xb2, 0x5c, 0x92, 0xf3, 0x1f, 0x44, 0xee, 0x4d, 0x91, 0xef, 0x06, 0x78, 0x5c,
	0x91, 0x0c, 0x54, 0x22, 0xab, 0xa2, 0xc0, 0x82, 0x82, 0xf4, 0x3d, 0x53, 0xc8, 0x97, 0xc1, 0x42,
	0xe3, 0x76, 0x26, 0x13, 0xd8, 0x99, 0x3e, 0x30, 0xa4, 0x13, 0x03, 0x9a, 0x8e, 0x6e, 0xe2, 0x8b,
	0x4f, 0x29, 0xc8, 0xdd, 0x3f, 0xaf, 0xa2, 0x1b, 0x2f, 0x13, 0x78, 0x5f, 0xa0, 0xd5, 0xb4, 0x92,
	0x8a, 0x17, 0x20, 0x8c, 0x7d, 0x2e, 0xda, 0x4b, 0x93, 0xed, 0xfd, 0x84, 0x76, 0x5c, 0x9c, 0x10,
	0x90, 0xa9, 0xa0, 0xa5, 0xa2, 0x13, 0xeb, 0xc4, 0x6d, 0xfc, 0xf6, 0x4d, 0x27, 0xbf, 0x3f, 0x53,
	0x1b, 0xcb, 0xbe, 0x87, 0x36, 0xe7, 0xfb, 0x64, 0xec, 0x78, 0x11, 0x6e, 0x63, 0x6e, 0xeb, 0xda,
	0xe8, 0x5e, 0x68, 0xb6, 0x29, 0xac, 0x8d, 0xff, 0xf6, 0xe6, 0x48, 0xa6, 0xa4, 0x63, 0xf4, 0x46,
	0x59, 0x89, 0xf4, 0x17, 0x2c, 0x21, 0xe1, 0x82, 0x80, 0x48, 0x58, 0x55, 0x8c, 0x41, 0xb4, 0xb2,
	0xe1, 0x6d, 0x27, 0xfd, 0x56, 0x2b, 0x8f, 0x8c, 0xf0, 0x15, 0x23, 0xbc, 0xfc, 0x5a, 0x47, 0x78,
	0xe5, 0xf5, 0x8d, 0xf0, 0xea, 0x7f, 0x1a, 0xe1, 0x47, 0x68, 0x4f, 0xff, 0x9f, 0xe1, 0x71, 0x0e,
	0x09, 0x4e, 0x15, 0x9d, 0x50, 0x35, 0xfd, 0x5f, 0xbe, 0xbd, 0xe3, 0x70, 0xb1, 0xa5, 0x35, 0x4b,
	0xd1, 0xd1, 0xb3, 0xf8, 0x01, 0xea, 0xcf, 0xe4, 0x36, 0x2a, 0xa9, 0xd4, 0x98, 0xd0, 0x59, 0xc1,
	0xfb, 0xee, 0xe6, 0xc9, 0xf0, 0x89, 0x0b, 0x9f, 0xba, 0xcf, 0x18, 0x19, 0x3e, 0xb1, 0xd1, 0xd3,
	0x83, 0xbf, 0x3b, 0xe8, 0x76, 0xca, 0x8b, 0xc5, 0x53, 0x79, 0xb0, 0x6e, 0xd9, 0xc7, 0xba, 0x29,
	0xc7, 0x9d, 0x9f, 0xbf, 0xb6, 0x92, 0x8c, 0xe7, 0x98, 0x65, 0x01, 0x17, 0x59, 0x98, 0x01, 0x33,
	0x2d, 0x0b, 0x67, 0x55, 0xbd, 0xe2, 0x8b, 0xee, 0x6e, 0x13, 0xfd, 0xd6, 0x5d, 0x3a, 0x8c, 0xe3,
	0xdf, 0xbb, 0xb7, 0x0e, 0x6b, 0x64, 0x4c, 0x64, 0x50, 0x87, 0x3a, 0x3a, 0x1d, 0x04, 0x23, 0x97,
	0xf9, 0x87, 0xcb, 0x39, 0x8b, 0x89, 0x3c, 0x6b, 0x72, 0xce, 0x4e, 0x07, 0x67, 0x4d, 0xce, 0xb3,
	0xee, 0xed, 0x7a, 0x21, 0x8a, 0x62, 0x22, 0xa3, 0xa8, 0xc9, 0x8a, 0xa2, 0xd3, 0x41, 0x14, 0x35,
	0x79, 0xe3, 0x65, 0x53, 0xec, 0xe0, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x62, 0x31, 0x73,
	0x7d, 0x0a, 0x00, 0x00,
}
