// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/text_sentiment.proto

package automl

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

// Contains annotation details specific to text sentiment.
type TextSentimentAnnotation struct {
	// Output only. The sentiment with the semantic, as given to the
	// [AutoMl.ImportData][google.cloud.automl.v1.AutoMl.ImportData] when populating the dataset from which the model used
	// for the prediction had been trained.
	// The sentiment values are between 0 and
	// Dataset.text_sentiment_dataset_metadata.sentiment_max (inclusive),
	// with higher value meaning more positive sentiment. They are completely
	// relative, i.e. 0 means least positive sentiment and sentiment_max means
	// the most positive from the sentiments present in the train data. Therefore
	//  e.g. if train data had only negative sentiment, then sentiment_max, would
	// be still negative (although least negative).
	// The sentiment shouldn't be confused with "score" or "magnitude"
	// from the previous Natural Language Sentiment Analysis API.
	Sentiment            int32    `protobuf:"varint,1,opt,name=sentiment,proto3" json:"sentiment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentAnnotation) Reset()         { *m = TextSentimentAnnotation{} }
func (m *TextSentimentAnnotation) String() string { return proto.CompactTextString(m) }
func (*TextSentimentAnnotation) ProtoMessage()    {}
func (*TextSentimentAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0ef9f1a6c8dafd9, []int{0}
}

func (m *TextSentimentAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentAnnotation.Unmarshal(m, b)
}
func (m *TextSentimentAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentAnnotation.Marshal(b, m, deterministic)
}
func (m *TextSentimentAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentAnnotation.Merge(m, src)
}
func (m *TextSentimentAnnotation) XXX_Size() int {
	return xxx_messageInfo_TextSentimentAnnotation.Size(m)
}
func (m *TextSentimentAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentAnnotation proto.InternalMessageInfo

func (m *TextSentimentAnnotation) GetSentiment() int32 {
	if m != nil {
		return m.Sentiment
	}
	return 0
}

// Model evaluation metrics for text sentiment problems.
type TextSentimentEvaluationMetrics struct {
	// Output only. Precision.
	Precision float32 `protobuf:"fixed32,1,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. Recall.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,3,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	// Output only. Mean absolute error. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	MeanAbsoluteError float32 `protobuf:"fixed32,4,opt,name=mean_absolute_error,json=meanAbsoluteError,proto3" json:"mean_absolute_error,omitempty"`
	// Output only. Mean squared error. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	MeanSquaredError float32 `protobuf:"fixed32,5,opt,name=mean_squared_error,json=meanSquaredError,proto3" json:"mean_squared_error,omitempty"`
	// Output only. Linear weighted kappa. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	LinearKappa float32 `protobuf:"fixed32,6,opt,name=linear_kappa,json=linearKappa,proto3" json:"linear_kappa,omitempty"`
	// Output only. Quadratic weighted kappa. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	QuadraticKappa float32 `protobuf:"fixed32,7,opt,name=quadratic_kappa,json=quadraticKappa,proto3" json:"quadratic_kappa,omitempty"`
	// Output only. Confusion matrix of the evaluation.
	// Only set for the overall model evaluation, not for evaluation of a single
	// annotation spec.
	ConfusionMatrix      *ClassificationEvaluationMetrics_ConfusionMatrix `protobuf:"bytes,8,opt,name=confusion_matrix,json=confusionMatrix,proto3" json:"confusion_matrix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *TextSentimentEvaluationMetrics) Reset()         { *m = TextSentimentEvaluationMetrics{} }
func (m *TextSentimentEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*TextSentimentEvaluationMetrics) ProtoMessage()    {}
func (*TextSentimentEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0ef9f1a6c8dafd9, []int{1}
}

func (m *TextSentimentEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentEvaluationMetrics.Unmarshal(m, b)
}
func (m *TextSentimentEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *TextSentimentEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentEvaluationMetrics.Merge(m, src)
}
func (m *TextSentimentEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_TextSentimentEvaluationMetrics.Size(m)
}
func (m *TextSentimentEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentEvaluationMetrics proto.InternalMessageInfo

func (m *TextSentimentEvaluationMetrics) GetPrecision() float32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetRecall() float32 {
	if m != nil {
		return m.Recall
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetF1Score() float32 {
	if m != nil {
		return m.F1Score
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetMeanAbsoluteError() float32 {
	if m != nil {
		return m.MeanAbsoluteError
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetMeanSquaredError() float32 {
	if m != nil {
		return m.MeanSquaredError
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetLinearKappa() float32 {
	if m != nil {
		return m.LinearKappa
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetQuadraticKappa() float32 {
	if m != nil {
		return m.QuadraticKappa
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetConfusionMatrix() *ClassificationEvaluationMetrics_ConfusionMatrix {
	if m != nil {
		return m.ConfusionMatrix
	}
	return nil
}

func init() {
	proto.RegisterType((*TextSentimentAnnotation)(nil), "google.cloud.automl.v1.TextSentimentAnnotation")
	proto.RegisterType((*TextSentimentEvaluationMetrics)(nil), "google.cloud.automl.v1.TextSentimentEvaluationMetrics")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/text_sentiment.proto", fileDescriptor_b0ef9f1a6c8dafd9)
}

var fileDescriptor_b0ef9f1a6c8dafd9 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0x8e, 0x75, 0xc3, 0x43, 0x6c, 0x18, 0xa9, 0x64, 0xd5, 0x84, 0xc6, 0x6e, 0x98,
	0x04, 0x72, 0x14, 0xb8, 0x40, 0x0a, 0xdc, 0x74, 0xd5, 0xb4, 0x0b, 0xa8, 0x34, 0xb5, 0xa8, 0x17,
	0xa8, 0x52, 0x74, 0xe6, 0xba, 0x91, 0x85, 0x63, 0x67, 0xb6, 0x53, 0xf5, 0x2d, 0x78, 0x0f, 0x1e,
	0x82, 0x07, 0xe0, 0x51, 0x78, 0x0a, 0x94, 0xe3, 0xb4, 0x53, 0x61, 0xbb, 0x8b, 0xcf, 0xf7, 0x9d,
	0xff, 0x24, 0x39, 0x26, 0x6f, 0x0a, 0x63, 0x0a, 0x25, 0x12, 0xae, 0x4c, 0x3d, 0x4f, 0xa0, 0xf6,
	0xa6, 0x54, 0xc9, 0x32, 0x4d, 0xbc, 0x58, 0xf9, 0xdc, 0x09, 0xed, 0x65, 0x29, 0xb4, 0x67, 0x95,
	0x35, 0xde, 0xd0, 0x5e, 0x90, 0x19, 0xca, 0x2c, 0xc8, 0x6c, 0x99, 0xf6, 0x1f, 0x0a, 0xe1, 0x0a,
	0x9c, 0x93, 0x0b, 0xc9, 0xc1, 0x4b, 0xa3, 0x43, 0x48, 0xff, 0xa4, 0x95, 0xa1, 0x92, 0x09, 0x68,
	0x6d, 0x3c, 0x42, 0x17, 0xe8, 0xd9, 0x07, 0xf2, 0xe2, 0xab, 0x58, 0xf9, 0xc9, 0x7a, 0xf2, 0x60,
	0x63, 0xd0, 0x13, 0xf2, 0x78, 0xf3, 0x42, 0x71, 0x74, 0x1a, 0x9d, 0xef, 0x8e, 0xef, 0x0a, 0x67,
	0x3f, 0x76, 0xc8, 0xcb, 0xad, 0xce, 0xcb, 0x25, 0xa8, 0x1a, 0x3b, 0x47, 0xc2, 0x5b, 0xc9, 0x5d,
	0x13, 0x50, 0x59, 0xc1, 0xa5, 0x93, 0x46, 0x63, 0x40, 0x67, 0x7c, 0x57, 0xa0, 0x3d, 0xd2, 0xb5,
	0x82, 0x83, 0x52, 0x71, 0x07, 0x51, 0x7b, 0xa2, 0xc7, 0x64, 0x7f, 0x91, 0xe6, 0x8e, 0x1b, 0x2b,
	0xe2, 0x1d, 0x24, 0x7b, 0x8b, 0x74, 0xd2, 0x1c, 0x29, 0x23, 0xcf, 0x4b, 0x01, 0x3a, 0x87, 0x1b,
	0x67, 0x54, 0xed, 0x45, 0x2e, 0xac, 0x35, 0x36, 0x7e, 0x84, 0xd6, 0xb3, 0x06, 0x0d, 0x5a, 0x72,
	0xd9, 0x00, 0xfa, 0x96, 0x50, 0xf4, 0xdd, 0x6d, 0x0d, 0x56, 0xcc, 0x5b, 0x7d, 0x17, 0xf5, 0xa3,
	0x86, 0x4c, 0x02, 0x08, 0xf6, 0x2b, 0xf2, 0x44, 0x49, 0x2d, 0xc0, 0xe6, 0xdf, 0xa1, 0xaa, 0x20,
	0xee, 0xa2, 0x77, 0x10, 0x6a, 0x9f, 0x9b, 0x12, 0x7d, 0x4d, 0x0e, 0x6f, 0x6b, 0x98, 0x5b, 0xf0,
	0x92, 0xb7, 0xd6, 0x1e, 0x5a, 0x4f, 0x37, 0xe5, 0x20, 0x5a, 0x72, 0xc4, 0x8d, 0x5e, 0xd4, 0xcd,
	0x97, 0xe6, 0x25, 0x78, 0x2b, 0x57, 0xf1, 0xfe, 0x69, 0x74, 0x7e, 0xf0, 0xee, 0x8a, 0xdd, 0xbf,
	0x54, 0x36, 0xdc, 0x5a, 0xde, 0x7f, 0x7f, 0x93, 0x0d, 0xd7, 0x79, 0x23, 0x8c, 0x1b, 0x1f, 0xf2,
	0xed, 0xc2, 0xc5, 0xaf, 0x88, 0xf4, 0xb9, 0x29, 0x1f, 0xc8, 0xbf, 0xa0, 0x5b, 0xdb, 0xba, 0x6e,
	0xb6, 0x7f, 0x1d, 0x7d, 0xfb, 0xd4, 0xda, 0x85, 0x51, 0xa0, 0x0b, 0x66, 0x6c, 0x91, 0x14, 0x42,
	0xe3, 0xdd, 0x48, 0x02, 0x82, 0x4a, 0xba, 0x7f, 0x6f, 0xda, 0xc7, 0xf0, 0xf4, 0xb3, 0xd3, 0xbb,
	0x0a, 0xed, 0x43, 0x1c, 0x36, 0xa8, 0xbd, 0x19, 0x7d, 0x61, 0xd3, 0xf4, 0xf7, 0x1a, 0xcc, 0x10,
	0xcc, 0x10, 0xa8, 0xd9, 0x34, 0xfd, 0xd3, 0x39, 0x0e, 0x20, 0xcb, 0x90, 0x64, 0x59, 0xe8, 0xc9,
	0xb2, 0x69, 0x7a, 0xd3, 0xc5, 0xb1, 0xef, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xd4, 0x4b,
	0x62, 0x24, 0x03, 0x00, 0x00,
}
