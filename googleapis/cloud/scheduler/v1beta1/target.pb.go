// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/scheduler/v1beta1/target.proto

package scheduler

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

// The HTTP method used to execute the job.
type HttpMethod int32

const (
	// HTTP method unspecified. Defaults to POST.
	HttpMethod_HTTP_METHOD_UNSPECIFIED HttpMethod = 0
	// HTTP POST
	HttpMethod_POST HttpMethod = 1
	// HTTP GET
	HttpMethod_GET HttpMethod = 2
	// HTTP HEAD
	HttpMethod_HEAD HttpMethod = 3
	// HTTP PUT
	HttpMethod_PUT HttpMethod = 4
	// HTTP DELETE
	HttpMethod_DELETE HttpMethod = 5
	// HTTP PATCH
	HttpMethod_PATCH HttpMethod = 6
	// HTTP OPTIONS
	HttpMethod_OPTIONS HttpMethod = 7
)

var HttpMethod_name = map[int32]string{
	0: "HTTP_METHOD_UNSPECIFIED",
	1: "POST",
	2: "GET",
	3: "HEAD",
	4: "PUT",
	5: "DELETE",
	6: "PATCH",
	7: "OPTIONS",
}

var HttpMethod_value = map[string]int32{
	"HTTP_METHOD_UNSPECIFIED": 0,
	"POST":                    1,
	"GET":                     2,
	"HEAD":                    3,
	"PUT":                     4,
	"DELETE":                  5,
	"PATCH":                   6,
	"OPTIONS":                 7,
}

func (x HttpMethod) String() string {
	return proto.EnumName(HttpMethod_name, int32(x))
}

func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30cd09e73d799915, []int{0}
}

// Http target. The job will be pushed to the job handler by means of
// an HTTP request via an [http_method][google.cloud.scheduler.v1beta1.HttpTarget.http_method] such as HTTP
// POST, HTTP GET, etc. The job is acknowledged by means of an HTTP
// response code in the range [200 - 299]. A failure to receive a response
// constitutes a failed execution. For a redirected request, the response
// returned by the redirected request is considered.
type HttpTarget struct {
	// Required. The full URI path that the request will be sent to. This string
	// must begin with either "http://" or "https://". Some examples of
	// valid values for [uri][google.cloud.scheduler.v1beta1.HttpTarget.uri] are:
	// `http://acme.com` and `https://acme.com/sales:8080`. Cloud Scheduler will
	// encode some characters for safety and compatibility. The maximum allowed
	// URL length is 2083 characters after encoding.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Which HTTP method to use for the request.
	HttpMethod HttpMethod `protobuf:"varint,2,opt,name=http_method,json=httpMethod,proto3,enum=google.cloud.scheduler.v1beta1.HttpMethod" json:"http_method,omitempty"`
	// The user can specify HTTP request headers to send with the job's
	// HTTP request. This map contains the header field names and
	// values. Repeated headers are not supported, but a header value can
	// contain commas. These headers represent a subset of the headers
	// that will accompany the job's HTTP request. Some HTTP request
	// headers will be ignored or replaced. A partial list of headers that
	// will be ignored or replaced is below:
	// - Host: This will be computed by Cloud Scheduler and derived from
	// [uri][google.cloud.scheduler.v1beta1.HttpTarget.uri].
	// * `Content-Length`: This will be computed by Cloud Scheduler.
	// * `User-Agent`: This will be set to `"Google-Cloud-Scheduler"`.
	// * `X-Google-*`: Google internal use only.
	// * `X-AppEngine-*`: Google internal use only.
	//
	// The total size of headers must be less than 80KB.
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// HTTP request body. A request body is allowed only if the HTTP
	// method is POST, PUT, or PATCH. It is an error to set body on a job with an
	// incompatible [HttpMethod][google.cloud.scheduler.v1beta1.HttpMethod].
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	// The mode for generating an `Authorization` header for HTTP requests.
	//
	// If specified, all `Authorization` headers in the [HttpTarget.headers][google.cloud.scheduler.v1beta1.HttpTarget.headers]
	// field will be overridden.
	//
	// Types that are valid to be assigned to AuthorizationHeader:
	//	*HttpTarget_OauthToken
	//	*HttpTarget_OidcToken
	AuthorizationHeader  isHttpTarget_AuthorizationHeader `protobuf_oneof:"authorization_header"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *HttpTarget) Reset()         { *m = HttpTarget{} }
func (m *HttpTarget) String() string { return proto.CompactTextString(m) }
func (*HttpTarget) ProtoMessage()    {}
func (*HttpTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_30cd09e73d799915, []int{0}
}

func (m *HttpTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpTarget.Unmarshal(m, b)
}
func (m *HttpTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpTarget.Marshal(b, m, deterministic)
}
func (m *HttpTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpTarget.Merge(m, src)
}
func (m *HttpTarget) XXX_Size() int {
	return xxx_messageInfo_HttpTarget.Size(m)
}
func (m *HttpTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpTarget.DiscardUnknown(m)
}

var xxx_messageInfo_HttpTarget proto.InternalMessageInfo

func (m *HttpTarget) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *HttpTarget) GetHttpMethod() HttpMethod {
	if m != nil {
		return m.HttpMethod
	}
	return HttpMethod_HTTP_METHOD_UNSPECIFIED
}

func (m *HttpTarget) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpTarget) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type isHttpTarget_AuthorizationHeader interface {
	isHttpTarget_AuthorizationHeader()
}

type HttpTarget_OauthToken struct {
	OauthToken *OAuthToken `protobuf:"bytes,5,opt,name=oauth_token,json=oauthToken,proto3,oneof"`
}

type HttpTarget_OidcToken struct {
	OidcToken *OidcToken `protobuf:"bytes,6,opt,name=oidc_token,json=oidcToken,proto3,oneof"`
}

func (*HttpTarget_OauthToken) isHttpTarget_AuthorizationHeader() {}

func (*HttpTarget_OidcToken) isHttpTarget_AuthorizationHeader() {}

func (m *HttpTarget) GetAuthorizationHeader() isHttpTarget_AuthorizationHeader {
	if m != nil {
		return m.AuthorizationHeader
	}
	return nil
}

func (m *HttpTarget) GetOauthToken() *OAuthToken {
	if x, ok := m.GetAuthorizationHeader().(*HttpTarget_OauthToken); ok {
		return x.OauthToken
	}
	return nil
}

func (m *HttpTarget) GetOidcToken() *OidcToken {
	if x, ok := m.GetAuthorizationHeader().(*HttpTarget_OidcToken); ok {
		return x.OidcToken
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpTarget) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpTarget_OauthToken)(nil),
		(*HttpTarget_OidcToken)(nil),
	}
}

// App Engine target. The job will be pushed to a job handler by means
// of an HTTP request via an [http_method][google.cloud.scheduler.v1beta1.AppEngineHttpTarget.http_method] such
// as HTTP POST, HTTP GET, etc. The job is acknowledged by means of an
// HTTP response code in the range [200 - 299]. Error 503 is
// considered an App Engine system error instead of an application
// error. Requests returning error 503 will be retried regardless of
// retry configuration and not counted against retry counts. Any other
// response code, or a failure to receive a response before the
// deadline, constitutes a failed attempt.
type AppEngineHttpTarget struct {
	// The HTTP method to use for the request. PATCH and OPTIONS are not
	// permitted.
	HttpMethod HttpMethod `protobuf:"varint,1,opt,name=http_method,json=httpMethod,proto3,enum=google.cloud.scheduler.v1beta1.HttpMethod" json:"http_method,omitempty"`
	// App Engine Routing setting for the job.
	AppEngineRouting *AppEngineRouting `protobuf:"bytes,2,opt,name=app_engine_routing,json=appEngineRouting,proto3" json:"app_engine_routing,omitempty"`
	// The relative URI.
	//
	// The relative URL must begin with "/" and must be a valid HTTP relative URL.
	// It can contain a path, query string arguments, and `#` fragments.
	// If the relative URL is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083 characters.
	RelativeUri string `protobuf:"bytes,3,opt,name=relative_uri,json=relativeUri,proto3" json:"relative_uri,omitempty"`
	// HTTP request headers.
	//
	// This map contains the header field names and values. Headers can be set
	// when the job is created.
	//
	// Cloud Scheduler sets some headers to default values:
	//
	// * `User-Agent`: By default, this header is
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"`.
	//   This header can be modified, but Cloud Scheduler will append
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"` to the
	//   modified `User-Agent`.
	// * `X-CloudScheduler`: This header will be set to true.
	//
	// If the job has an [body][google.cloud.scheduler.v1beta1.AppEngineHttpTarget.body], Cloud Scheduler sets
	// the following headers:
	//
	// * `Content-Type`: By default, the `Content-Type` header is set to
	//   `"application/octet-stream"`. The default can be overridden by explictly
	//   setting `Content-Type` to a particular media type when the job is
	//   created.
	//   For example, `Content-Type` can be set to `"application/json"`.
	// * `Content-Length`: This is computed by Cloud Scheduler. This value is
	//   output only. It cannot be changed.
	//
	// The headers below are output only. They cannot be set or overridden:
	//
	// * `X-Google-*`: For Google internal use only.
	// * `X-AppEngine-*`: For Google internal use only.
	//
	// In addition, some App Engine headers, which contain
	// job-specific information, are also be sent to the job handler.
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Body.
	//
	// HTTP request body. A request body is allowed only if the HTTP method is
	// POST or PUT. It will result in invalid argument error to set a body on a
	// job with an incompatible [HttpMethod][google.cloud.scheduler.v1beta1.HttpMethod].
	Body                 []byte   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEngineHttpTarget) Reset()         { *m = AppEngineHttpTarget{} }
func (m *AppEngineHttpTarget) String() string { return proto.CompactTextString(m) }
func (*AppEngineHttpTarget) ProtoMessage()    {}
func (*AppEngineHttpTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_30cd09e73d799915, []int{1}
}

func (m *AppEngineHttpTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineHttpTarget.Unmarshal(m, b)
}
func (m *AppEngineHttpTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineHttpTarget.Marshal(b, m, deterministic)
}
func (m *AppEngineHttpTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineHttpTarget.Merge(m, src)
}
func (m *AppEngineHttpTarget) XXX_Size() int {
	return xxx_messageInfo_AppEngineHttpTarget.Size(m)
}
func (m *AppEngineHttpTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineHttpTarget.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineHttpTarget proto.InternalMessageInfo

func (m *AppEngineHttpTarget) GetHttpMethod() HttpMethod {
	if m != nil {
		return m.HttpMethod
	}
	return HttpMethod_HTTP_METHOD_UNSPECIFIED
}

func (m *AppEngineHttpTarget) GetAppEngineRouting() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRouting
	}
	return nil
}

func (m *AppEngineHttpTarget) GetRelativeUri() string {
	if m != nil {
		return m.RelativeUri
	}
	return ""
}

func (m *AppEngineHttpTarget) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AppEngineHttpTarget) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// Pub/Sub target. The job will be delivered by publishing a message to
// the given Pub/Sub topic.
type PubsubTarget struct {
	// Required. The name of the Cloud Pub/Sub topic to which messages will
	// be published when a job is delivered. The topic name must be in the
	// same format as required by PubSub's
	// [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest),
	// for example `projects/PROJECT_ID/topics/TOPIC_ID`.
	//
	// The topic must be in the same project as the Cloud Scheduler job.
	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	// The message payload for PubsubMessage.
	//
	// Pubsub message must contain either non-empty data, or at least one
	// attribute.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// Attributes for PubsubMessage.
	//
	// Pubsub message must contain either non-empty data, or at least one
	// attribute.
	Attributes           map[string]string `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PubsubTarget) Reset()         { *m = PubsubTarget{} }
func (m *PubsubTarget) String() string { return proto.CompactTextString(m) }
func (*PubsubTarget) ProtoMessage()    {}
func (*PubsubTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_30cd09e73d799915, []int{2}
}

func (m *PubsubTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubsubTarget.Unmarshal(m, b)
}
func (m *PubsubTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubsubTarget.Marshal(b, m, deterministic)
}
func (m *PubsubTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubsubTarget.Merge(m, src)
}
func (m *PubsubTarget) XXX_Size() int {
	return xxx_messageInfo_PubsubTarget.Size(m)
}
func (m *PubsubTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_PubsubTarget.DiscardUnknown(m)
}

var xxx_messageInfo_PubsubTarget proto.InternalMessageInfo

func (m *PubsubTarget) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *PubsubTarget) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PubsubTarget) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// App Engine Routing.
//
// For more information about services, versions, and instances see
// [An Overview of App
// Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine),
// [Microservices Architecture on Google App
// Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request
// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed),
// and [App Engine Flex request
// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
type AppEngineRouting struct {
	// App service.
	//
	// By default, the job is sent to the service which is the default
	// service when the job is attempted.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// App version.
	//
	// By default, the job is sent to the version which is the default
	// version when the job is attempted.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// App instance.
	//
	// By default, the job is sent to an instance which is available when
	// the job is attempted.
	//
	// Requests can only be sent to a specific instance if
	// [manual scaling is used in App Engine
	// Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes).
	// App Engine Flex does not support instances. For more information, see
	// [App Engine Standard request
	// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
	// and [App Engine Flex request
	// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
	Instance string `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	// Output only. The host that the job is sent to.
	//
	// For more information about how App Engine requests are routed, see
	// [here](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
	//
	// The host is constructed as:
	//
	//
	// * `host = [application_domain_name]`</br>
	//   `| [service] + '.' + [application_domain_name]`</br>
	//   `| [version] + '.' + [application_domain_name]`</br>
	//   `| [version_dot_service]+ '.' + [application_domain_name]`</br>
	//   `| [instance] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_service] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version_dot_service] + '.' + [application_domain_name]`
	//
	// * `application_domain_name` = The domain name of the app, for
	//   example <app-id>.appspot.com, which is associated with the
	//   job's project ID.
	//
	// * `service =` [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service]
	//
	// * `version =` [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version]
	//
	// * `version_dot_service =`
	//   [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version] `+ '.' +`
	//   [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service]
	//
	// * `instance =` [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance]
	//
	// * `instance_dot_service =`
	//   [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] `+ '.' +`
	//   [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service]
	//
	// * `instance_dot_version =`
	//   [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] `+ '.' +`
	//   [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version]
	//
	// * `instance_dot_version_dot_service =`
	//   [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] `+ '.' +`
	//   [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version] `+ '.' +`
	//   [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service]
	//
	//
	// If [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service] is empty, then the job will be sent
	// to the service which is the default service when the job is attempted.
	//
	// If [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version] is empty, then the job will be sent
	// to the version which is the default version when the job is attempted.
	//
	// If [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] is empty, then the job will be
	// sent to an instance which is available when the job is attempted.
	//
	// If [service][google.cloud.scheduler.v1beta1.AppEngineRouting.service],
	// [version][google.cloud.scheduler.v1beta1.AppEngineRouting.version], or
	// [instance][google.cloud.scheduler.v1beta1.AppEngineRouting.instance] is invalid, then the job will be sent
	// to the default version of the default service when the job is attempted.
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEngineRouting) Reset()         { *m = AppEngineRouting{} }
func (m *AppEngineRouting) String() string { return proto.CompactTextString(m) }
func (*AppEngineRouting) ProtoMessage()    {}
func (*AppEngineRouting) Descriptor() ([]byte, []int) {
	return fileDescriptor_30cd09e73d799915, []int{3}
}

func (m *AppEngineRouting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineRouting.Unmarshal(m, b)
}
func (m *AppEngineRouting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineRouting.Marshal(b, m, deterministic)
}
func (m *AppEngineRouting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineRouting.Merge(m, src)
}
func (m *AppEngineRouting) XXX_Size() int {
	return xxx_messageInfo_AppEngineRouting.Size(m)
}
func (m *AppEngineRouting) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineRouting.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineRouting proto.InternalMessageInfo

func (m *AppEngineRouting) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AppEngineRouting) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AppEngineRouting) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *AppEngineRouting) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

// Contains information needed for generating an
// [OAuth token](https://developers.google.com/identity/protocols/OAuth2).
// This type of authorization should generally only be used when calling Google
// APIs hosted on *.googleapis.com.
type OAuthToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	// to be used for generating OAuth token.
	// The service account must be within the same project as the job. The caller
	// must have iam.serviceAccounts.actAs permission for the service account.
	ServiceAccountEmail string `protobuf:"bytes,1,opt,name=service_account_email,json=serviceAccountEmail,proto3" json:"service_account_email,omitempty"`
	// OAuth scope to be used for generating OAuth access token.
	// If not specified, "https://www.googleapis.com/auth/cloud-platform"
	// will be used.
	Scope                string   `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthToken) Reset()         { *m = OAuthToken{} }
func (m *OAuthToken) String() string { return proto.CompactTextString(m) }
func (*OAuthToken) ProtoMessage()    {}
func (*OAuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_30cd09e73d799915, []int{4}
}

func (m *OAuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthToken.Unmarshal(m, b)
}
func (m *OAuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthToken.Marshal(b, m, deterministic)
}
func (m *OAuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthToken.Merge(m, src)
}
func (m *OAuthToken) XXX_Size() int {
	return xxx_messageInfo_OAuthToken.Size(m)
}
func (m *OAuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthToken proto.InternalMessageInfo

func (m *OAuthToken) GetServiceAccountEmail() string {
	if m != nil {
		return m.ServiceAccountEmail
	}
	return ""
}

func (m *OAuthToken) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

// Contains information needed for generating an
// [OpenID Connect
// token](https://developers.google.com/identity/protocols/OpenIDConnect).
// This type of authorization can be used for many scenarios, including
// calling Cloud Run, or endpoints where you intend to validate the token
// yourself.
type OidcToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	// to be used for generating OIDC token.
	// The service account must be within the same project as the job. The caller
	// must have iam.serviceAccounts.actAs permission for the service account.
	ServiceAccountEmail string `protobuf:"bytes,1,opt,name=service_account_email,json=serviceAccountEmail,proto3" json:"service_account_email,omitempty"`
	// Audience to be used when generating OIDC token. If not specified, the URI
	// specified in target will be used.
	Audience             string   `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OidcToken) Reset()         { *m = OidcToken{} }
func (m *OidcToken) String() string { return proto.CompactTextString(m) }
func (*OidcToken) ProtoMessage()    {}
func (*OidcToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_30cd09e73d799915, []int{5}
}

func (m *OidcToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OidcToken.Unmarshal(m, b)
}
func (m *OidcToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OidcToken.Marshal(b, m, deterministic)
}
func (m *OidcToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OidcToken.Merge(m, src)
}
func (m *OidcToken) XXX_Size() int {
	return xxx_messageInfo_OidcToken.Size(m)
}
func (m *OidcToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OidcToken.DiscardUnknown(m)
}

var xxx_messageInfo_OidcToken proto.InternalMessageInfo

func (m *OidcToken) GetServiceAccountEmail() string {
	if m != nil {
		return m.ServiceAccountEmail
	}
	return ""
}

func (m *OidcToken) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.scheduler.v1beta1.HttpMethod", HttpMethod_name, HttpMethod_value)
	proto.RegisterType((*HttpTarget)(nil), "google.cloud.scheduler.v1beta1.HttpTarget")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.scheduler.v1beta1.HttpTarget.HeadersEntry")
	proto.RegisterType((*AppEngineHttpTarget)(nil), "google.cloud.scheduler.v1beta1.AppEngineHttpTarget")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.scheduler.v1beta1.AppEngineHttpTarget.HeadersEntry")
	proto.RegisterType((*PubsubTarget)(nil), "google.cloud.scheduler.v1beta1.PubsubTarget")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.scheduler.v1beta1.PubsubTarget.AttributesEntry")
	proto.RegisterType((*AppEngineRouting)(nil), "google.cloud.scheduler.v1beta1.AppEngineRouting")
	proto.RegisterType((*OAuthToken)(nil), "google.cloud.scheduler.v1beta1.OAuthToken")
	proto.RegisterType((*OidcToken)(nil), "google.cloud.scheduler.v1beta1.OidcToken")
}

func init() {
	proto.RegisterFile("google/cloud/scheduler/v1beta1/target.proto", fileDescriptor_30cd09e73d799915)
}

var fileDescriptor_30cd09e73d799915 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0xaf, 0xcf, 0xc9, 0xe5, 0x32, 0x8e, 0xc0, 0xda, 0x16, 0x08, 0x29, 0xa0, 0x34, 0x4f, 0xa1,
	0x48, 0x36, 0x3d, 0x1e, 0x40, 0x05, 0x44, 0x7d, 0x3d, 0xd3, 0x1c, 0x70, 0x17, 0xe3, 0xf3, 0xf1,
	0x50, 0x10, 0xd6, 0xc6, 0x5e, 0xd9, 0xcb, 0x25, 0x5e, 0x6b, 0xbd, 0x8e, 0x74, 0x54, 0x95, 0xf8,
	0x0a, 0x7c, 0x22, 0xc4, 0xd7, 0xe1, 0x23, 0xf0, 0x84, 0x76, 0xfd, 0x27, 0x51, 0x24, 0x9a, 0x83,
	0x3e, 0x65, 0x66, 0x67, 0xe6, 0xb7, 0x33, 0xbf, 0xf9, 0x6d, 0x0c, 0x1f, 0x25, 0x8c, 0x25, 0x4b,
	0x62, 0x47, 0x4b, 0x56, 0xc6, 0x76, 0x11, 0xa5, 0x24, 0x2e, 0x97, 0x84, 0xdb, 0xeb, 0x47, 0x0b,
	0x22, 0xf0, 0x23, 0x5b, 0x60, 0x9e, 0x10, 0x61, 0xe5, 0x9c, 0x09, 0x86, 0x3e, 0xa8, 0x92, 0x2d,
	0x95, 0x6c, 0xb5, 0xc9, 0x56, 0x9d, 0x3c, 0x7a, 0xb7, 0x06, 0xc3, 0x39, 0xb5, 0x39, 0x29, 0x58,
	0xc9, 0x23, 0x52, 0x95, 0x8e, 0xde, 0xdb, 0x0a, 0xe1, 0x2c, 0x63, 0x02, 0x0b, 0xca, 0xb2, 0xa2,
	0x8a, 0x4e, 0xfe, 0xd4, 0x01, 0x66, 0x42, 0xe4, 0x81, 0xba, 0x0d, 0x99, 0xa0, 0x97, 0x9c, 0x0e,
	0xb5, 0xb1, 0x36, 0xed, 0xfb, 0xd2, 0x44, 0xdf, 0x82, 0x91, 0x0a, 0x91, 0x87, 0x2b, 0x22, 0x52,
	0x16, 0x0f, 0x0f, 0xc6, 0xda, 0xf4, 0x8d, 0xe3, 0x87, 0xd6, 0xab, 0xfb, 0xb1, 0x24, 0xe4, 0xb9,
	0xaa, 0xf0, 0x21, 0x6d, 0x6d, 0xf4, 0x3d, 0xf4, 0x52, 0x82, 0x63, 0xc2, 0x8b, 0xa1, 0x3e, 0xd6,
	0xa7, 0xc6, 0xf1, 0xa7, 0xb7, 0x01, 0xaa, 0x7a, 0xb3, 0x66, 0x55, 0xa5, 0x9b, 0x09, 0x7e, 0xe3,
	0x37, 0x38, 0x08, 0x41, 0x67, 0xc1, 0xe2, 0x9b, 0x61, 0x67, 0xac, 0x4d, 0x07, 0xbe, 0xb2, 0xd1,
	0x39, 0x18, 0x0c, 0x97, 0x22, 0x0d, 0x05, 0xbb, 0x26, 0xd9, 0xb0, 0x3b, 0xd6, 0xa6, 0xc6, 0xfe,
	0x9e, 0xe7, 0x4e, 0x29, 0xd2, 0x40, 0x56, 0xcc, 0xee, 0xf8, 0xa0, 0x00, 0x94, 0x87, 0xbe, 0x01,
	0x60, 0x34, 0x8e, 0x6a, 0xb4, 0x43, 0x85, 0xf6, 0xe1, 0x5e, 0x34, 0x1a, 0x47, 0x0d, 0x58, 0x9f,
	0x35, 0xce, 0xe8, 0x31, 0x0c, 0xb6, 0xe7, 0x90, 0x84, 0x5f, 0x93, 0x9b, 0x86, 0xf0, 0x6b, 0x72,
	0x83, 0xee, 0x41, 0x77, 0x8d, 0x97, 0x25, 0x51, 0x54, 0xf7, 0xfd, 0xca, 0x79, 0x7c, 0xf0, 0x99,
	0x76, 0xf2, 0x36, 0xdc, 0x93, 0x4d, 0x31, 0x4e, 0x7f, 0x55, 0x3b, 0x0c, 0x2b, 0x0e, 0x26, 0xbf,
	0xeb, 0x70, 0xd7, 0xc9, 0x73, 0x37, 0x4b, 0x68, 0x46, 0xb6, 0x96, 0xb9, 0xb3, 0x3a, 0xed, 0xb5,
	0x56, 0xf7, 0x33, 0x20, 0x9c, 0xe7, 0x21, 0x51, 0x97, 0x84, 0x9c, 0x95, 0x82, 0x66, 0x89, 0xea,
	0xd1, 0x38, 0xfe, 0x78, 0x1f, 0x66, 0xdb, 0x9d, 0x5f, 0xd5, 0xf9, 0x26, 0xde, 0x39, 0x41, 0x0f,
	0x60, 0xc0, 0xc9, 0x12, 0x0b, 0xba, 0x26, 0xa1, 0x94, 0xa0, 0xae, 0xa6, 0x37, 0x9a, 0xb3, 0x2b,
	0x4e, 0xd1, 0xf3, 0x8d, 0x7a, 0x3a, 0x4a, 0x3d, 0x4f, 0x6e, 0x7d, 0xef, 0xad, 0x65, 0xd4, 0xdd,
	0xc8, 0xe8, 0x75, 0x76, 0x35, 0xf9, 0xed, 0x00, 0x06, 0x5e, 0xb9, 0x28, 0xca, 0x45, 0xbd, 0x8c,
	0xaf, 0x00, 0x04, 0xcb, 0x69, 0x14, 0x66, 0x78, 0x45, 0x2a, 0x8c, 0x93, 0xf1, 0xdf, 0xce, 0xfb,
	0x70, 0x3f, 0x57, 0x69, 0xf5, 0x28, 0x38, 0xa7, 0x85, 0x15, 0xb1, 0x95, 0x1d, 0xc8, 0x64, 0xbf,
	0xaf, 0x6a, 0x2e, 0xf0, 0x8a, 0xc8, 0x0e, 0x63, 0x2c, 0xb0, 0x22, 0x66, 0xe0, 0x2b, 0x1b, 0xfd,
	0x04, 0x80, 0x85, 0xe0, 0x74, 0x51, 0x0a, 0xd2, 0x90, 0xf2, 0xc5, 0x3e, 0x52, 0xb6, 0xdb, 0xb2,
	0x9c, 0xb6, 0xbc, 0x22, 0x64, 0x0b, 0x6f, 0xf4, 0x25, 0xbc, 0xb9, 0x13, 0xfe, 0x4f, 0x14, 0xac,
	0xc1, 0xdc, 0xdd, 0x3b, 0x1a, 0x42, 0xaf, 0x20, 0x7c, 0x4d, 0xa3, 0x9a, 0x02, 0xbf, 0x71, 0x65,
	0x64, 0x4d, 0x78, 0x41, 0x59, 0x56, 0x23, 0x35, 0x2e, 0x1a, 0xc1, 0x11, 0xcd, 0x0a, 0x81, 0xb3,
	0x88, 0xd4, 0xaa, 0x68, 0x7d, 0x49, 0x4a, 0xca, 0x0a, 0xa1, 0x5e, 0x7f, 0xdf, 0x57, 0xf6, 0xe4,
	0x07, 0x80, 0xcd, 0x53, 0x46, 0xc7, 0xf0, 0x56, 0x7d, 0x45, 0x88, 0xa3, 0x88, 0x95, 0x99, 0x08,
	0xc9, 0x0a, 0xd3, 0x65, 0x7d, 0xff, 0xdd, 0x3a, 0xe8, 0x54, 0x31, 0x57, 0x86, 0xe4, 0x4c, 0x45,
	0xc4, 0xf2, 0x76, 0x26, 0xe5, 0x4c, 0x7e, 0x84, 0x7e, 0xfb, 0xa8, 0xff, 0x17, 0xec, 0x08, 0x8e,
	0x70, 0x19, 0x53, 0x22, 0x07, 0xa9, 0x90, 0x5b, 0xff, 0x61, 0x51, 0xfd, 0x0d, 0xd7, 0x8f, 0xed,
	0x3e, 0xbc, 0x33, 0x0b, 0x02, 0x2f, 0x3c, 0x77, 0x83, 0xd9, 0xfc, 0x34, 0xbc, 0xba, 0xb8, 0xf4,
	0xdc, 0xa7, 0x67, 0x5f, 0x9f, 0xb9, 0xa7, 0xe6, 0x1d, 0x74, 0x04, 0x1d, 0x6f, 0x7e, 0x19, 0x98,
	0x1a, 0xea, 0x81, 0xfe, 0xcc, 0x0d, 0xcc, 0x03, 0x79, 0x34, 0x73, 0x9d, 0x53, 0x53, 0x97, 0x47,
	0xde, 0x55, 0x60, 0x76, 0x10, 0xc0, 0xe1, 0xa9, 0xfb, 0x9d, 0x1b, 0xb8, 0x66, 0x17, 0xf5, 0xa1,
	0xeb, 0x39, 0xc1, 0xd3, 0x99, 0x79, 0x88, 0x0c, 0xe8, 0xcd, 0xbd, 0xe0, 0x6c, 0x7e, 0x71, 0x69,
	0xf6, 0x4e, 0xfe, 0xd0, 0xfe, 0x72, 0x9e, 0xbc, 0x52, 0x80, 0xe8, 0x41, 0xce, 0xd9, 0x2f, 0x24,
	0x12, 0x85, 0xfd, 0xa2, 0xb6, 0x5e, 0xda, 0x4a, 0x93, 0x85, 0xfd, 0x42, 0xfd, 0xbe, 0x84, 0x49,
	0xc4, 0x56, 0x7b, 0x34, 0x77, 0x62, 0x54, 0x72, 0xf3, 0xe4, 0x37, 0xc7, 0xd3, 0x9e, 0x3f, 0xab,
	0xd3, 0x13, 0xb6, 0xc4, 0x59, 0x62, 0x31, 0x9e, 0xd8, 0x09, 0xc9, 0xd4, 0x17, 0xc9, 0xde, 0xb4,
	0xf1, 0x6f, 0x9f, 0xc6, 0xcf, 0xdb, 0x93, 0xc5, 0xa1, 0xaa, 0xf9, 0xe4, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x53, 0x2a, 0x3d, 0xd9, 0x4d, 0x07, 0x00, 0x00,
}
