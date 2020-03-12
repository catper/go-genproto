// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/common/matching_function.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
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

// Matching function associated with a
// CustomerFeed, CampaignFeed, or AdGroupFeed. The matching function is used
// to filter the set of feed items selected.
type MatchingFunction struct {
	// String representation of the Function.
	//
	// Examples:
	//
	// 1. IDENTITY(true) or IDENTITY(false). All or no feed items served.
	// 2. EQUALS(CONTEXT.DEVICE,"Mobile")
	// 3. IN(FEED_ITEM_ID,{1000001,1000002,1000003})
	// 4. CONTAINS_ANY(FeedAttribute[12345678,0],{"Mars cruise","Venus cruise"})
	// 5. AND(IN(FEED_ITEM_ID,{10001,10002}),EQUALS(CONTEXT.DEVICE,"Mobile"))
	//
	// For more details, visit
	// https://developers.google.com/adwords/api/docs/guides/feed-matching-functions
	//
	// Note that because multiple strings may represent the same underlying
	// function (whitespace and single versus double quotation marks, for
	// example), the value returned may not be identical to the string sent in a
	// mutate request.
	FunctionString *wrappers.StringValue `protobuf:"bytes,1,opt,name=function_string,json=functionString,proto3" json:"function_string,omitempty"`
	// Operator for a function.
	Operator enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator `protobuf:"varint,4,opt,name=operator,proto3,enum=google.ads.googleads.v3.enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator" json:"operator,omitempty"`
	// The operands on the left hand side of the equation. This is also the
	// operand to be used for single operand expressions such as NOT.
	LeftOperands []*Operand `protobuf:"bytes,2,rep,name=left_operands,json=leftOperands,proto3" json:"left_operands,omitempty"`
	// The operands on the right hand side of the equation.
	RightOperands        []*Operand `protobuf:"bytes,3,rep,name=right_operands,json=rightOperands,proto3" json:"right_operands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MatchingFunction) Reset()         { *m = MatchingFunction{} }
func (m *MatchingFunction) String() string { return proto.CompactTextString(m) }
func (*MatchingFunction) ProtoMessage()    {}
func (*MatchingFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73413c68281034a, []int{0}
}

func (m *MatchingFunction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchingFunction.Unmarshal(m, b)
}
func (m *MatchingFunction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchingFunction.Marshal(b, m, deterministic)
}
func (m *MatchingFunction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchingFunction.Merge(m, src)
}
func (m *MatchingFunction) XXX_Size() int {
	return xxx_messageInfo_MatchingFunction.Size(m)
}
func (m *MatchingFunction) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchingFunction.DiscardUnknown(m)
}

var xxx_messageInfo_MatchingFunction proto.InternalMessageInfo

func (m *MatchingFunction) GetFunctionString() *wrappers.StringValue {
	if m != nil {
		return m.FunctionString
	}
	return nil
}

func (m *MatchingFunction) GetOperator() enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator {
	if m != nil {
		return m.Operator
	}
	return enums.MatchingFunctionOperatorEnum_UNSPECIFIED
}

func (m *MatchingFunction) GetLeftOperands() []*Operand {
	if m != nil {
		return m.LeftOperands
	}
	return nil
}

func (m *MatchingFunction) GetRightOperands() []*Operand {
	if m != nil {
		return m.RightOperands
	}
	return nil
}

// An operand in a matching function.
type Operand struct {
	// Different operands that can be used in a matching function. Required.
	//
	// Types that are valid to be assigned to FunctionArgumentOperand:
	//	*Operand_ConstantOperand_
	//	*Operand_FeedAttributeOperand_
	//	*Operand_FunctionOperand_
	//	*Operand_RequestContextOperand_
	FunctionArgumentOperand isOperand_FunctionArgumentOperand `protobuf_oneof:"function_argument_operand"`
	XXX_NoUnkeyedLiteral    struct{}                          `json:"-"`
	XXX_unrecognized        []byte                            `json:"-"`
	XXX_sizecache           int32                             `json:"-"`
}

func (m *Operand) Reset()         { *m = Operand{} }
func (m *Operand) String() string { return proto.CompactTextString(m) }
func (*Operand) ProtoMessage()    {}
func (*Operand) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73413c68281034a, []int{1}
}

func (m *Operand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operand.Unmarshal(m, b)
}
func (m *Operand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operand.Marshal(b, m, deterministic)
}
func (m *Operand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operand.Merge(m, src)
}
func (m *Operand) XXX_Size() int {
	return xxx_messageInfo_Operand.Size(m)
}
func (m *Operand) XXX_DiscardUnknown() {
	xxx_messageInfo_Operand.DiscardUnknown(m)
}

var xxx_messageInfo_Operand proto.InternalMessageInfo

type isOperand_FunctionArgumentOperand interface {
	isOperand_FunctionArgumentOperand()
}

type Operand_ConstantOperand_ struct {
	ConstantOperand *Operand_ConstantOperand `protobuf:"bytes,1,opt,name=constant_operand,json=constantOperand,proto3,oneof"`
}

type Operand_FeedAttributeOperand_ struct {
	FeedAttributeOperand *Operand_FeedAttributeOperand `protobuf:"bytes,2,opt,name=feed_attribute_operand,json=feedAttributeOperand,proto3,oneof"`
}

type Operand_FunctionOperand_ struct {
	FunctionOperand *Operand_FunctionOperand `protobuf:"bytes,3,opt,name=function_operand,json=functionOperand,proto3,oneof"`
}

type Operand_RequestContextOperand_ struct {
	RequestContextOperand *Operand_RequestContextOperand `protobuf:"bytes,4,opt,name=request_context_operand,json=requestContextOperand,proto3,oneof"`
}

func (*Operand_ConstantOperand_) isOperand_FunctionArgumentOperand() {}

func (*Operand_FeedAttributeOperand_) isOperand_FunctionArgumentOperand() {}

func (*Operand_FunctionOperand_) isOperand_FunctionArgumentOperand() {}

func (*Operand_RequestContextOperand_) isOperand_FunctionArgumentOperand() {}

func (m *Operand) GetFunctionArgumentOperand() isOperand_FunctionArgumentOperand {
	if m != nil {
		return m.FunctionArgumentOperand
	}
	return nil
}

func (m *Operand) GetConstantOperand() *Operand_ConstantOperand {
	if x, ok := m.GetFunctionArgumentOperand().(*Operand_ConstantOperand_); ok {
		return x.ConstantOperand
	}
	return nil
}

func (m *Operand) GetFeedAttributeOperand() *Operand_FeedAttributeOperand {
	if x, ok := m.GetFunctionArgumentOperand().(*Operand_FeedAttributeOperand_); ok {
		return x.FeedAttributeOperand
	}
	return nil
}

func (m *Operand) GetFunctionOperand() *Operand_FunctionOperand {
	if x, ok := m.GetFunctionArgumentOperand().(*Operand_FunctionOperand_); ok {
		return x.FunctionOperand
	}
	return nil
}

func (m *Operand) GetRequestContextOperand() *Operand_RequestContextOperand {
	if x, ok := m.GetFunctionArgumentOperand().(*Operand_RequestContextOperand_); ok {
		return x.RequestContextOperand
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operand) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operand_ConstantOperand_)(nil),
		(*Operand_FeedAttributeOperand_)(nil),
		(*Operand_FunctionOperand_)(nil),
		(*Operand_RequestContextOperand_)(nil),
	}
}

// A constant operand in a matching function.
type Operand_ConstantOperand struct {
	// Constant operand values. Required.
	//
	// Types that are valid to be assigned to ConstantOperandValue:
	//	*Operand_ConstantOperand_StringValue
	//	*Operand_ConstantOperand_LongValue
	//	*Operand_ConstantOperand_BooleanValue
	//	*Operand_ConstantOperand_DoubleValue
	ConstantOperandValue isOperand_ConstantOperand_ConstantOperandValue `protobuf_oneof:"constant_operand_value"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *Operand_ConstantOperand) Reset()         { *m = Operand_ConstantOperand{} }
func (m *Operand_ConstantOperand) String() string { return proto.CompactTextString(m) }
func (*Operand_ConstantOperand) ProtoMessage()    {}
func (*Operand_ConstantOperand) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73413c68281034a, []int{1, 0}
}

func (m *Operand_ConstantOperand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operand_ConstantOperand.Unmarshal(m, b)
}
func (m *Operand_ConstantOperand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operand_ConstantOperand.Marshal(b, m, deterministic)
}
func (m *Operand_ConstantOperand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operand_ConstantOperand.Merge(m, src)
}
func (m *Operand_ConstantOperand) XXX_Size() int {
	return xxx_messageInfo_Operand_ConstantOperand.Size(m)
}
func (m *Operand_ConstantOperand) XXX_DiscardUnknown() {
	xxx_messageInfo_Operand_ConstantOperand.DiscardUnknown(m)
}

var xxx_messageInfo_Operand_ConstantOperand proto.InternalMessageInfo

type isOperand_ConstantOperand_ConstantOperandValue interface {
	isOperand_ConstantOperand_ConstantOperandValue()
}

type Operand_ConstantOperand_StringValue struct {
	StringValue *wrappers.StringValue `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Operand_ConstantOperand_LongValue struct {
	LongValue *wrappers.Int64Value `protobuf:"bytes,2,opt,name=long_value,json=longValue,proto3,oneof"`
}

type Operand_ConstantOperand_BooleanValue struct {
	BooleanValue *wrappers.BoolValue `protobuf:"bytes,3,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Operand_ConstantOperand_DoubleValue struct {
	DoubleValue *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

func (*Operand_ConstantOperand_StringValue) isOperand_ConstantOperand_ConstantOperandValue() {}

func (*Operand_ConstantOperand_LongValue) isOperand_ConstantOperand_ConstantOperandValue() {}

func (*Operand_ConstantOperand_BooleanValue) isOperand_ConstantOperand_ConstantOperandValue() {}

func (*Operand_ConstantOperand_DoubleValue) isOperand_ConstantOperand_ConstantOperandValue() {}

func (m *Operand_ConstantOperand) GetConstantOperandValue() isOperand_ConstantOperand_ConstantOperandValue {
	if m != nil {
		return m.ConstantOperandValue
	}
	return nil
}

func (m *Operand_ConstantOperand) GetStringValue() *wrappers.StringValue {
	if x, ok := m.GetConstantOperandValue().(*Operand_ConstantOperand_StringValue); ok {
		return x.StringValue
	}
	return nil
}

func (m *Operand_ConstantOperand) GetLongValue() *wrappers.Int64Value {
	if x, ok := m.GetConstantOperandValue().(*Operand_ConstantOperand_LongValue); ok {
		return x.LongValue
	}
	return nil
}

func (m *Operand_ConstantOperand) GetBooleanValue() *wrappers.BoolValue {
	if x, ok := m.GetConstantOperandValue().(*Operand_ConstantOperand_BooleanValue); ok {
		return x.BooleanValue
	}
	return nil
}

func (m *Operand_ConstantOperand) GetDoubleValue() *wrappers.DoubleValue {
	if x, ok := m.GetConstantOperandValue().(*Operand_ConstantOperand_DoubleValue); ok {
		return x.DoubleValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operand_ConstantOperand) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operand_ConstantOperand_StringValue)(nil),
		(*Operand_ConstantOperand_LongValue)(nil),
		(*Operand_ConstantOperand_BooleanValue)(nil),
		(*Operand_ConstantOperand_DoubleValue)(nil),
	}
}

// A feed attribute operand in a matching function.
// Used to represent a feed attribute in feed.
type Operand_FeedAttributeOperand struct {
	// The associated feed. Required.
	FeedId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	// Id of the referenced feed attribute. Required.
	FeedAttributeId      *wrappers.Int64Value `protobuf:"bytes,2,opt,name=feed_attribute_id,json=feedAttributeId,proto3" json:"feed_attribute_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Operand_FeedAttributeOperand) Reset()         { *m = Operand_FeedAttributeOperand{} }
func (m *Operand_FeedAttributeOperand) String() string { return proto.CompactTextString(m) }
func (*Operand_FeedAttributeOperand) ProtoMessage()    {}
func (*Operand_FeedAttributeOperand) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73413c68281034a, []int{1, 1}
}

func (m *Operand_FeedAttributeOperand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operand_FeedAttributeOperand.Unmarshal(m, b)
}
func (m *Operand_FeedAttributeOperand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operand_FeedAttributeOperand.Marshal(b, m, deterministic)
}
func (m *Operand_FeedAttributeOperand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operand_FeedAttributeOperand.Merge(m, src)
}
func (m *Operand_FeedAttributeOperand) XXX_Size() int {
	return xxx_messageInfo_Operand_FeedAttributeOperand.Size(m)
}
func (m *Operand_FeedAttributeOperand) XXX_DiscardUnknown() {
	xxx_messageInfo_Operand_FeedAttributeOperand.DiscardUnknown(m)
}

var xxx_messageInfo_Operand_FeedAttributeOperand proto.InternalMessageInfo

func (m *Operand_FeedAttributeOperand) GetFeedId() *wrappers.Int64Value {
	if m != nil {
		return m.FeedId
	}
	return nil
}

func (m *Operand_FeedAttributeOperand) GetFeedAttributeId() *wrappers.Int64Value {
	if m != nil {
		return m.FeedAttributeId
	}
	return nil
}

// A function operand in a matching function.
// Used to represent nested functions.
type Operand_FunctionOperand struct {
	// The matching function held in this operand.
	MatchingFunction     *MatchingFunction `protobuf:"bytes,1,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Operand_FunctionOperand) Reset()         { *m = Operand_FunctionOperand{} }
func (m *Operand_FunctionOperand) String() string { return proto.CompactTextString(m) }
func (*Operand_FunctionOperand) ProtoMessage()    {}
func (*Operand_FunctionOperand) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73413c68281034a, []int{1, 2}
}

func (m *Operand_FunctionOperand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operand_FunctionOperand.Unmarshal(m, b)
}
func (m *Operand_FunctionOperand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operand_FunctionOperand.Marshal(b, m, deterministic)
}
func (m *Operand_FunctionOperand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operand_FunctionOperand.Merge(m, src)
}
func (m *Operand_FunctionOperand) XXX_Size() int {
	return xxx_messageInfo_Operand_FunctionOperand.Size(m)
}
func (m *Operand_FunctionOperand) XXX_DiscardUnknown() {
	xxx_messageInfo_Operand_FunctionOperand.DiscardUnknown(m)
}

var xxx_messageInfo_Operand_FunctionOperand proto.InternalMessageInfo

func (m *Operand_FunctionOperand) GetMatchingFunction() *MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

// An operand in a function referring to a value in the request context.
type Operand_RequestContextOperand struct {
	// Type of value to be referred in the request context.
	ContextType          enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType `protobuf:"varint,1,opt,name=context_type,json=contextType,proto3,enum=google.ads.googleads.v3.enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType" json:"context_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                          `json:"-"`
	XXX_unrecognized     []byte                                                            `json:"-"`
	XXX_sizecache        int32                                                             `json:"-"`
}

func (m *Operand_RequestContextOperand) Reset()         { *m = Operand_RequestContextOperand{} }
func (m *Operand_RequestContextOperand) String() string { return proto.CompactTextString(m) }
func (*Operand_RequestContextOperand) ProtoMessage()    {}
func (*Operand_RequestContextOperand) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73413c68281034a, []int{1, 3}
}

func (m *Operand_RequestContextOperand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operand_RequestContextOperand.Unmarshal(m, b)
}
func (m *Operand_RequestContextOperand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operand_RequestContextOperand.Marshal(b, m, deterministic)
}
func (m *Operand_RequestContextOperand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operand_RequestContextOperand.Merge(m, src)
}
func (m *Operand_RequestContextOperand) XXX_Size() int {
	return xxx_messageInfo_Operand_RequestContextOperand.Size(m)
}
func (m *Operand_RequestContextOperand) XXX_DiscardUnknown() {
	xxx_messageInfo_Operand_RequestContextOperand.DiscardUnknown(m)
}

var xxx_messageInfo_Operand_RequestContextOperand proto.InternalMessageInfo

func (m *Operand_RequestContextOperand) GetContextType() enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType {
	if m != nil {
		return m.ContextType
	}
	return enums.MatchingFunctionContextTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MatchingFunction)(nil), "google.ads.googleads.v3.common.MatchingFunction")
	proto.RegisterType((*Operand)(nil), "google.ads.googleads.v3.common.Operand")
	proto.RegisterType((*Operand_ConstantOperand)(nil), "google.ads.googleads.v3.common.Operand.ConstantOperand")
	proto.RegisterType((*Operand_FeedAttributeOperand)(nil), "google.ads.googleads.v3.common.Operand.FeedAttributeOperand")
	proto.RegisterType((*Operand_FunctionOperand)(nil), "google.ads.googleads.v3.common.Operand.FunctionOperand")
	proto.RegisterType((*Operand_RequestContextOperand)(nil), "google.ads.googleads.v3.common.Operand.RequestContextOperand")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/common/matching_function.proto", fileDescriptor_a73413c68281034a)
}

var fileDescriptor_a73413c68281034a = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x26, 0x09, 0x82, 0xcb, 0x24, 0x24, 0x60, 0x01, 0x37, 0x37, 0x20, 0x84, 0xb2, 0xb9, 0xac,
	0xc6, 0x55, 0x82, 0xa8, 0xe4, 0xd2, 0x4a, 0x0e, 0xff, 0x52, 0x5b, 0x90, 0x8b, 0xb2, 0x40, 0xa9,
	0x22, 0xc7, 0x33, 0x31, 0x96, 0xec, 0x19, 0xd7, 0x33, 0x86, 0xb2, 0xe8, 0x63, 0xb4, 0x0f, 0xd0,
	0x65, 0x97, 0xdd, 0xf5, 0x15, 0xfa, 0x26, 0xed, 0xaa, 0x8f, 0x50, 0xd9, 0x33, 0x76, 0x88, 0x93,
	0x40, 0xb2, 0xf2, 0x9c, 0x39, 0xe7, 0xfb, 0xbe, 0x39, 0x3f, 0x33, 0x06, 0xfb, 0x36, 0xa5, 0xb6,
	0x8b, 0x55, 0x13, 0x31, 0x55, 0x2c, 0xa3, 0xd5, 0x6d, 0x53, 0xb5, 0xa8, 0xe7, 0x51, 0xa2, 0x7a,
	0x26, 0xb7, 0x6e, 0x1c, 0x62, 0x77, 0xfb, 0x21, 0xb1, 0xb8, 0x43, 0x09, 0xf4, 0x03, 0xca, 0xa9,
	0xb2, 0x2d, 0x82, 0xa1, 0x89, 0x18, 0x4c, 0x71, 0xf0, 0xb6, 0x09, 0x05, 0xae, 0xd6, 0x9a, 0xc4,
	0x8b, 0x49, 0xe8, 0xb1, 0x51, 0xda, 0xae, 0x45, 0x09, 0xc7, 0x1f, 0x79, 0x97, 0xdf, 0xfb, 0x58,
	0x68, 0xd4, 0x5e, 0xcd, 0xca, 0x41, 0x7d, 0x1c, 0x98, 0x9c, 0x06, 0x12, 0x2f, 0xcf, 0xa8, 0xc6,
	0x56, 0x2f, 0xec, 0xab, 0x77, 0x81, 0xe9, 0xfb, 0x38, 0x60, 0xd2, 0xbf, 0x95, 0xf0, 0xfb, 0x8e,
	0x6a, 0x12, 0x42, 0xb9, 0x19, 0xb1, 0x48, 0x6f, 0xfd, 0x57, 0x1e, 0xac, 0xbc, 0x91, 0x12, 0x27,
	0x52, 0x41, 0x39, 0x06, 0x95, 0x54, 0x8d, 0xf1, 0xc0, 0x21, 0x76, 0x35, 0xb7, 0x93, 0xdb, 0x2d,
	0x36, 0xb6, 0x64, 0x15, 0x60, 0x22, 0x06, 0xdf, 0xc5, 0xee, 0xb6, 0xe9, 0x86, 0xd8, 0x28, 0x27,
	0x20, 0xb1, 0xa9, 0xf8, 0xe0, 0x9f, 0xe4, 0xac, 0xd5, 0xf9, 0x9d, 0xdc, 0x6e, 0xb9, 0x71, 0x05,
	0x27, 0x15, 0x34, 0x4e, 0x16, 0x66, 0x4f, 0x72, 0x21, 0xe1, 0xc7, 0x24, 0xf4, 0x26, 0x3a, 0x8d,
	0x54, 0x45, 0x79, 0x0d, 0x96, 0x5d, 0xdc, 0xe7, 0xa2, 0x44, 0x04, 0xb1, 0x6a, 0x7e, 0xa7, 0xb0,
	0x5b, 0x6c, 0xfc, 0x0f, 0x1f, 0xef, 0x23, 0xbc, 0x10, 0xf1, 0x46, 0x29, 0x42, 0x4b, 0x83, 0x29,
	0x6f, 0x41, 0x39, 0x70, 0xec, 0x9b, 0x07, 0x74, 0x85, 0xd9, 0xe8, 0x96, 0x63, 0x78, 0xc2, 0x57,
	0xff, 0xb1, 0x04, 0x16, 0xa5, 0xa1, 0x20, 0xb0, 0x62, 0x51, 0xc2, 0xb8, 0x49, 0x52, 0x7a, 0x59,
	0xe3, 0xe7, 0x53, 0xb2, 0xc3, 0x43, 0x89, 0x97, 0xf6, 0xd9, 0x9c, 0x51, 0xb1, 0x86, 0xb7, 0x14,
	0x0e, 0x36, 0xfa, 0x18, 0xa3, 0xae, 0xc9, 0x79, 0xe0, 0xf4, 0x42, 0x8e, 0x53, 0xad, 0x7c, 0xac,
	0x75, 0x30, 0xad, 0xd6, 0x09, 0xc6, 0x48, 0x4f, 0x48, 0x06, 0x82, 0x6b, 0xfd, 0x31, 0xfb, 0x51,
	0x6e, 0xc3, 0xc3, 0x4a, 0x50, 0xb5, 0x30, 0x5b, 0x6e, 0x43, 0x3d, 0x16, 0xb9, 0xf5, 0x87, 0xb7,
	0x94, 0x3b, 0xf0, 0x6f, 0x80, 0x3f, 0x84, 0x98, 0xf1, 0xf4, 0x56, 0x25, 0x62, 0xf3, 0xb1, 0xd8,
	0xcb, 0x69, 0xc5, 0x0c, 0x41, 0x73, 0x28, 0x58, 0x06, 0x92, 0xeb, 0xc1, 0x38, 0x47, 0xed, 0x7b,
	0x1e, 0x54, 0x32, 0xb5, 0x57, 0x74, 0x50, 0x12, 0x17, 0xa5, 0x7b, 0x1b, 0x5d, 0x85, 0x69, 0xae,
	0xcb, 0xd9, 0x9c, 0x51, 0x64, 0x03, 0x53, 0x39, 0x00, 0xc0, 0xa5, 0x29, 0x81, 0xe8, 0xcf, 0xe6,
	0x08, 0xc1, 0x39, 0xe1, 0xfb, 0x7b, 0x09, 0x7e, 0x29, 0x02, 0x08, 0xb4, 0x0e, 0x96, 0x7b, 0x94,
	0xba, 0xd8, 0x24, 0x92, 0x40, 0x14, 0xbc, 0x36, 0x42, 0xd0, 0xa2, 0xd4, 0x4d, 0xf0, 0x25, 0x09,
	0x49, 0x28, 0x4a, 0x88, 0x86, 0x3d, 0x17, 0x4b, 0x86, 0xf9, 0x09, 0x39, 0x1c, 0xc5, 0x41, 0x69,
	0x0e, 0x68, 0x60, 0xb6, 0xaa, 0x60, 0x23, 0x3b, 0xd5, 0x82, 0xac, 0xf6, 0x39, 0x07, 0xd6, 0xc6,
	0x0d, 0x91, 0xb2, 0x07, 0x16, 0xe3, 0x11, 0x75, 0x92, 0xf9, 0x7f, 0x2c, 0x67, 0x63, 0x21, 0x8a,
	0x3d, 0x47, 0xca, 0x29, 0x58, 0xcd, 0x0c, 0xb6, 0x83, 0xa6, 0xa8, 0x99, 0x51, 0x19, 0x1a, 0xd8,
	0x73, 0x54, 0xf3, 0x41, 0x25, 0x33, 0x6b, 0xca, 0x7b, 0xb0, 0x3a, 0xf2, 0xe8, 0xca, 0xb3, 0x3d,
	0x7b, 0x6a, 0xa4, 0xb2, 0x6f, 0x94, 0xb1, 0xe2, 0x65, 0x76, 0x6a, 0x5f, 0x72, 0x60, 0x7d, 0xec,
	0xc4, 0x29, 0x9f, 0x40, 0xe9, 0xe1, 0xff, 0x21, 0xd6, 0x2c, 0x37, 0xae, 0x67, 0x7c, 0x33, 0x25,
	0xe9, 0xd5, 0xbd, 0x8f, 0xc7, 0x3e, 0x9b, 0x0f, 0xfc, 0x46, 0xd1, 0x1a, 0x18, 0xad, 0x4d, 0xf0,
	0x5f, 0x7a, 0x6d, 0xcd, 0xc0, 0x0e, 0x3d, 0x3c, 0xe8, 0x62, 0xeb, 0x4f, 0x0e, 0xd4, 0x2d, 0xea,
	0x3d, 0x91, 0x7f, 0x6b, 0x3d, 0xab, 0x76, 0x19, 0x35, 0xe1, 0x32, 0x77, 0x7d, 0x24, 0x81, 0x36,
	0x75, 0x4d, 0x62, 0x43, 0x1a, 0xd8, 0xaa, 0x8d, 0x49, 0xdc, 0xa2, 0xe4, 0xaf, 0xe7, 0x3b, 0x6c,
	0xd2, 0x0f, 0xfa, 0x85, 0xf8, 0x7c, 0xcd, 0x17, 0x4e, 0x75, 0xfd, 0x5b, 0x7e, 0xfb, 0x54, 0x90,
	0xe9, 0x88, 0x41, 0xb1, 0x8c, 0x56, 0xed, 0x26, 0x3c, 0x8c, 0xc3, 0x7e, 0x26, 0x01, 0x1d, 0x1d,
	0xb1, 0x4e, 0x1a, 0xd0, 0x69, 0x37, 0x3b, 0x22, 0xe0, 0x77, 0xbe, 0x2e, 0x76, 0x35, 0x4d, 0x47,
	0x4c, 0xd3, 0xd2, 0x10, 0x4d, 0x6b, 0x37, 0x35, 0x4d, 0x04, 0xf5, 0x16, 0xe2, 0xd3, 0x35, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x29, 0xdc, 0x21, 0x0f, 0x3d, 0x08, 0x00, 0x00,
}
