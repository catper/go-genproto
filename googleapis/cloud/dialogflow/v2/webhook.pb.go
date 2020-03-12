// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2/webhook.proto

package dialogflow

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	_struct "github.com/catper/protobuf/ptypes/struct"
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

// The request message for a webhook call.
type WebhookRequest struct {
	// The unique identifier of detectIntent request session.
	// Can be used to identify end-user inside webhook implementation.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`, or
	// `projects/<Project ID>/agent/environments/<Environment ID>/users/<User
	// ID>/sessions/<Session ID>`.
	Session string `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
	// The unique identifier of the response. Contains the same value as
	// `[Streaming]DetectIntentResponse.response_id`.
	ResponseId string `protobuf:"bytes,1,opt,name=response_id,json=responseId,proto3" json:"response_id,omitempty"`
	// The result of the conversational query or event processing. Contains the
	// same value as `[Streaming]DetectIntentResponse.query_result`.
	QueryResult *QueryResult `protobuf:"bytes,2,opt,name=query_result,json=queryResult,proto3" json:"query_result,omitempty"`
	// Optional. The contents of the original request that was passed to
	// `[Streaming]DetectIntent` call.
	OriginalDetectIntentRequest *OriginalDetectIntentRequest `protobuf:"bytes,3,opt,name=original_detect_intent_request,json=originalDetectIntentRequest,proto3" json:"original_detect_intent_request,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                     `json:"-"`
	XXX_unrecognized            []byte                       `json:"-"`
	XXX_sizecache               int32                        `json:"-"`
}

func (m *WebhookRequest) Reset()         { *m = WebhookRequest{} }
func (m *WebhookRequest) String() string { return proto.CompactTextString(m) }
func (*WebhookRequest) ProtoMessage()    {}
func (*WebhookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba880626f278d96, []int{0}
}

func (m *WebhookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookRequest.Unmarshal(m, b)
}
func (m *WebhookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookRequest.Marshal(b, m, deterministic)
}
func (m *WebhookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookRequest.Merge(m, src)
}
func (m *WebhookRequest) XXX_Size() int {
	return xxx_messageInfo_WebhookRequest.Size(m)
}
func (m *WebhookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookRequest proto.InternalMessageInfo

func (m *WebhookRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *WebhookRequest) GetResponseId() string {
	if m != nil {
		return m.ResponseId
	}
	return ""
}

func (m *WebhookRequest) GetQueryResult() *QueryResult {
	if m != nil {
		return m.QueryResult
	}
	return nil
}

func (m *WebhookRequest) GetOriginalDetectIntentRequest() *OriginalDetectIntentRequest {
	if m != nil {
		return m.OriginalDetectIntentRequest
	}
	return nil
}

// The response message for a webhook call.
//
// This response is validated by the Dialogflow server. If validation fails,
// an error will be returned in the [QueryResult.diagnostic_info][google.cloud.dialogflow.v2.QueryResult.diagnostic_info] field.
// Setting JSON fields to an empty value with the wrong type is a common error.
// To avoid this error:
//
// - Use `""` for empty strings
// - Use `{}` or `null` for empty objects
// - Use `[]` or `null` for empty arrays
//
// For more information, see the
// [Protocol Buffers Language
// Guide](https://developers.google.com/protocol-buffers/docs/proto3#json).
type WebhookResponse struct {
	// Optional. The text to be shown on the screen. This value is passed directly
	// to `QueryResult.fulfillment_text`.
	FulfillmentText string `protobuf:"bytes,1,opt,name=fulfillment_text,json=fulfillmentText,proto3" json:"fulfillment_text,omitempty"`
	// Optional. The collection of rich messages to present to the user. This
	// value is passed directly to `QueryResult.fulfillment_messages`.
	FulfillmentMessages []*Intent_Message `protobuf:"bytes,2,rep,name=fulfillment_messages,json=fulfillmentMessages,proto3" json:"fulfillment_messages,omitempty"`
	// Optional. This value is passed directly to `QueryResult.webhook_source`.
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. This value is passed directly to `QueryResult.webhook_payload`.
	// See the related `fulfillment_messages[i].payload field`, which may be used
	// as an alternative to this field.
	//
	// This field can be used for Actions on Google responses.
	// It should have a structure similar to the JSON message shown here. For more
	// information, see
	// [Actions on Google Webhook
	// Format](https://developers.google.com/actions/dialogflow/webhook)
	// <pre>{
	//   "google": {
	//     "expectUserResponse": true,
	//     "richResponse": {
	//       "items": [
	//         {
	//           "simpleResponse": {
	//             "textToSpeech": "this is a simple response"
	//           }
	//         }
	//       ]
	//     }
	//   }
	// }</pre>
	Payload *_struct.Struct `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional. The collection of output contexts. This value is passed directly
	// to `QueryResult.output_contexts`.
	OutputContexts []*Context `protobuf:"bytes,5,rep,name=output_contexts,json=outputContexts,proto3" json:"output_contexts,omitempty"`
	// Optional. Makes the platform immediately invoke another `DetectIntent` call
	// internally with the specified event as input.
	// When this field is set, Dialogflow ignores the `fulfillment_text`,
	// `fulfillment_messages`, and `payload` fields.
	FollowupEventInput *EventInput `protobuf:"bytes,6,opt,name=followup_event_input,json=followupEventInput,proto3" json:"followup_event_input,omitempty"`
	// Optional. Additional session entity types to replace or extend developer
	// entity types with. The entity synonyms apply to all languages and persist
	// for the session of this query. Setting the session entity types inside
	// webhook overwrites the session entity types that have been set through
	// `DetectIntentRequest.query_params.session_entity_types`.
	SessionEntityTypes   []*SessionEntityType `protobuf:"bytes,10,rep,name=session_entity_types,json=sessionEntityTypes,proto3" json:"session_entity_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WebhookResponse) Reset()         { *m = WebhookResponse{} }
func (m *WebhookResponse) String() string { return proto.CompactTextString(m) }
func (*WebhookResponse) ProtoMessage()    {}
func (*WebhookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba880626f278d96, []int{1}
}

func (m *WebhookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookResponse.Unmarshal(m, b)
}
func (m *WebhookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookResponse.Marshal(b, m, deterministic)
}
func (m *WebhookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookResponse.Merge(m, src)
}
func (m *WebhookResponse) XXX_Size() int {
	return xxx_messageInfo_WebhookResponse.Size(m)
}
func (m *WebhookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookResponse proto.InternalMessageInfo

func (m *WebhookResponse) GetFulfillmentText() string {
	if m != nil {
		return m.FulfillmentText
	}
	return ""
}

func (m *WebhookResponse) GetFulfillmentMessages() []*Intent_Message {
	if m != nil {
		return m.FulfillmentMessages
	}
	return nil
}

func (m *WebhookResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *WebhookResponse) GetPayload() *_struct.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *WebhookResponse) GetOutputContexts() []*Context {
	if m != nil {
		return m.OutputContexts
	}
	return nil
}

func (m *WebhookResponse) GetFollowupEventInput() *EventInput {
	if m != nil {
		return m.FollowupEventInput
	}
	return nil
}

func (m *WebhookResponse) GetSessionEntityTypes() []*SessionEntityType {
	if m != nil {
		return m.SessionEntityTypes
	}
	return nil
}

// Represents the contents of the original request that was passed to
// the `[Streaming]DetectIntent` call.
type OriginalDetectIntentRequest struct {
	// The source of this request, e.g., `google`, `facebook`, `slack`. It is set
	// by Dialogflow-owned servers.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. The version of the protocol used for this request.
	// This field is AoG-specific.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Optional. This field is set to the value of the `QueryParameters.payload`
	// field passed in the request. Some integrations that query a Dialogflow
	// agent may provide additional information in the payload.
	//
	// In particular for the Telephony Gateway this field has the form:
	// <pre>{
	//  "telephony": {
	//    "caller_id": "+18558363987"
	//  }
	// }</pre>
	// Note: The caller ID field (`caller_id`) will be redacted for Standard
	// Edition agents and populated with the caller ID in [E.164
	// format](https://en.wikipedia.org/wiki/E.164) for Enterprise Edition agents.
	Payload              *_struct.Struct `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OriginalDetectIntentRequest) Reset()         { *m = OriginalDetectIntentRequest{} }
func (m *OriginalDetectIntentRequest) String() string { return proto.CompactTextString(m) }
func (*OriginalDetectIntentRequest) ProtoMessage()    {}
func (*OriginalDetectIntentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba880626f278d96, []int{2}
}

func (m *OriginalDetectIntentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalDetectIntentRequest.Unmarshal(m, b)
}
func (m *OriginalDetectIntentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalDetectIntentRequest.Marshal(b, m, deterministic)
}
func (m *OriginalDetectIntentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalDetectIntentRequest.Merge(m, src)
}
func (m *OriginalDetectIntentRequest) XXX_Size() int {
	return xxx_messageInfo_OriginalDetectIntentRequest.Size(m)
}
func (m *OriginalDetectIntentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalDetectIntentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalDetectIntentRequest proto.InternalMessageInfo

func (m *OriginalDetectIntentRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *OriginalDetectIntentRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *OriginalDetectIntentRequest) GetPayload() *_struct.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*WebhookRequest)(nil), "google.cloud.dialogflow.v2.WebhookRequest")
	proto.RegisterType((*WebhookResponse)(nil), "google.cloud.dialogflow.v2.WebhookResponse")
	proto.RegisterType((*OriginalDetectIntentRequest)(nil), "google.cloud.dialogflow.v2.OriginalDetectIntentRequest")
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2/webhook.proto", fileDescriptor_2ba880626f278d96)
}

var fileDescriptor_2ba880626f278d96 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0x9d, 0xef, 0x6b, 0xd5, 0x4d, 0xd5, 0xa2, 0xa5, 0x02, 0x2b, 0x45, 0x25, 0x0a, 0x12,
	0x0d, 0x48, 0xd8, 0xc2, 0x20, 0x71, 0xe0, 0xd6, 0xa6, 0xa0, 0x20, 0x10, 0xc5, 0xad, 0x00, 0x21,
	0x21, 0xcb, 0xb1, 0x37, 0x66, 0xc5, 0x66, 0xc7, 0xf1, 0xee, 0x26, 0x8d, 0xc4, 0x89, 0xff, 0xc0,
	0x89, 0x1b, 0x47, 0xfe, 0x1c, 0x57, 0x8e, 0xc8, 0xbb, 0x6b, 0x12, 0xaa, 0xd6, 0x70, 0x9c, 0x99,
	0x37, 0x6f, 0xde, 0x3c, 0xcf, 0x1a, 0xf5, 0x73, 0x80, 0x9c, 0x91, 0x20, 0x65, 0xa0, 0xb2, 0x20,
	0xa3, 0x09, 0x83, 0x7c, 0xcc, 0x60, 0x1e, 0xcc, 0xc2, 0x60, 0x4e, 0x46, 0x1f, 0x00, 0x3e, 0xfa,
	0x45, 0x09, 0x12, 0x70, 0xc7, 0x20, 0x7d, 0x8d, 0xf4, 0x97, 0x48, 0x7f, 0x16, 0x76, 0x9a, 0x58,
	0x52, 0xe0, 0x92, 0x9c, 0x49, 0xc3, 0xd2, 0xd9, 0x6f, 0x40, 0x52, 0x2e, 0x09, 0xaf, 0x81, 0x4d,
	0x94, 0x82, 0x08, 0x41, 0x81, 0x5b, 0xe4, 0xc3, 0xbf, 0x23, 0x63, 0xc2, 0x25, 0x95, 0x8b, 0x58,
	0x2e, 0x0a, 0x62, 0xbb, 0x6e, 0xd8, 0x2e, 0x1d, 0x8d, 0xd4, 0x38, 0x10, 0xb2, 0x54, 0xa9, 0x3c,
	0x57, 0x4d, 0x0a, 0x1a, 0x24, 0x9c, 0x83, 0x4c, 0x24, 0x05, 0x2e, 0x4c, 0xb5, 0xf7, 0xc5, 0x45,
	0x5b, 0x6f, 0x8c, 0x39, 0x11, 0x99, 0x2a, 0x22, 0x24, 0xf6, 0xd0, 0xba, 0x9d, 0xe5, 0xfd, 0xd7,
	0x75, 0xfa, 0x1b, 0x51, 0x1d, 0xe2, 0x9b, 0xa8, 0x5d, 0x12, 0x51, 0x00, 0x17, 0x24, 0xa6, 0x99,
	0xe7, 0xe8, 0x2a, 0xaa, 0x53, 0xc3, 0x0c, 0x3f, 0x43, 0x9b, 0x53, 0x45, 0xca, 0x45, 0x5c, 0x12,
	0xa1, 0x98, 0xf4, 0xdc, 0xae, 0xd3, 0x6f, 0x87, 0xfb, 0xfe, 0xe5, 0x7e, 0xfb, 0xaf, 0x2a, 0x7c,
	0xa4, 0xe1, 0x51, 0x7b, 0xba, 0x0c, 0xf0, 0x27, 0xb4, 0x07, 0x25, 0xcd, 0x29, 0x4f, 0x58, 0x9c,
	0x11, 0x49, 0x52, 0x19, 0x1b, 0x57, 0xe3, 0xd2, 0x08, 0xf5, 0x5a, 0x9a, 0xfd, 0x51, 0x13, 0xfb,
	0x4b, 0xcb, 0x30, 0xd0, 0x04, 0x43, 0xdd, 0x6f, 0xf7, 0x8c, 0x76, 0xe1, 0xf2, 0x62, 0xef, 0x47,
	0x0b, 0x6d, 0xff, 0xf6, 0xc5, 0xec, 0x87, 0xef, 0xa0, 0x2b, 0x63, 0xc5, 0xc6, 0x94, 0xb1, 0x49,
	0x25, 0xa3, 0x3a, 0x05, 0xeb, 0xc1, 0xf6, 0x4a, 0xfe, 0x94, 0x9c, 0x49, 0xfc, 0x1e, 0xed, 0xac,
	0x42, 0x27, 0x44, 0x88, 0x24, 0x27, 0xc2, 0x73, 0xbb, 0xad, 0x7e, 0x3b, 0xbc, 0xdb, 0x24, 0xd9,
	0xe8, 0xf0, 0x5f, 0x98, 0x96, 0xe8, 0xea, 0x0a, 0x8f, 0xcd, 0x09, 0x7c, 0x0d, 0xad, 0x09, 0x50,
	0x65, 0x4a, 0xb4, 0x07, 0x1b, 0x91, 0x8d, 0xf0, 0x7d, 0xb4, 0x5e, 0x24, 0x0b, 0x06, 0x49, 0xa6,
	0x3f, 0x5d, 0x3b, 0xbc, 0x5e, 0x4f, 0xaa, 0x6f, 0xc3, 0x3f, 0xd1, 0xb7, 0x11, 0xd5, 0x38, 0xfc,
	0x1c, 0x6d, 0x83, 0x92, 0x85, 0x92, 0xb1, 0xbd, 0x6e, 0xe1, 0xfd, 0xaf, 0x45, 0xde, 0x6a, 0x12,
	0x79, 0x68, 0xb0, 0xd1, 0x96, 0xe9, 0xb5, 0xa1, 0xc0, 0x6f, 0xd1, 0xce, 0x18, 0x18, 0x83, 0xb9,
	0x2a, 0x62, 0x32, 0xab, 0x56, 0xa7, 0xbc, 0x50, 0xd2, 0x5b, 0xd3, 0x6a, 0x6e, 0x37, 0x51, 0x1e,
	0x55, 0xf0, 0x61, 0x85, 0x8e, 0x70, 0xcd, 0xb1, 0xcc, 0xe1, 0x18, 0xed, 0x5c, 0xf0, 0x02, 0x84,
	0x87, 0xb4, 0xd8, 0x7b, 0x4d, 0xcc, 0x27, 0xa6, 0xef, 0x48, 0xb7, 0x9d, 0x2e, 0x0a, 0x12, 0x61,
	0x71, 0x3e, 0x25, 0x7a, 0x9f, 0x1d, 0xb4, 0xdb, 0x70, 0x2e, 0x2b, 0x9e, 0x3b, 0x7f, 0x78, 0xee,
	0xa1, 0xf5, 0x19, 0x29, 0xf5, 0x73, 0x71, 0xcd, 0x73, 0xb1, 0xe1, 0xea, 0xd7, 0x68, 0xfd, 0xdb,
	0xd7, 0x38, 0xf8, 0xea, 0xa0, 0xbd, 0x14, 0x26, 0x0d, 0xdb, 0x1c, 0x6c, 0xda, 0xb3, 0x3c, 0xae,
	0x38, 0x8e, 0x9d, 0x77, 0x03, 0x8b, 0xcd, 0x81, 0x25, 0x3c, 0xf7, 0xa1, 0xcc, 0x83, 0x9c, 0x70,
	0x3d, 0x21, 0x30, 0xa5, 0xa4, 0xa0, 0xe2, 0xa2, 0x5f, 0xca, 0xe3, 0x65, 0xf4, 0xd3, 0x71, 0xbe,
	0xb9, 0xee, 0xe0, 0xc9, 0x77, 0xb7, 0xf3, 0xd4, 0xd0, 0x1d, 0xea, 0xd1, 0x83, 0xe5, 0xe8, 0xd7,
	0xe1, 0x68, 0x4d, 0xb3, 0x3e, 0xf8, 0x15, 0x00, 0x00, 0xff, 0xff, 0x41, 0x99, 0x74, 0xaa, 0x6a,
	0x05, 0x00, 0x00,
}
