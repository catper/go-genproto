// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/spanner/v1/result_set.proto

package spanner

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

// Results from [Read][google.spanner.v1.Spanner.Read] or
// [ExecuteSql][google.spanner.v1.Spanner.ExecuteSql].
type ResultSet struct {
	// Metadata about the result set, such as row type information.
	Metadata *ResultSetMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Each element in `rows` is a row whose format is defined by
	// [metadata.row_type][google.spanner.v1.ResultSetMetadata.row_type]. The ith element
	// in each row matches the ith field in
	// [metadata.row_type][google.spanner.v1.ResultSetMetadata.row_type]. Elements are
	// encoded based on type as described
	// [here][google.spanner.v1.TypeCode].
	Rows []*_struct.ListValue `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	// Query plan and execution statistics for the SQL statement that
	// produced this result set. These can be requested by setting
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode].
	// DML statements always produce stats containing the number of rows
	// modified, unless executed using the
	// [ExecuteSqlRequest.QueryMode.PLAN][google.spanner.v1.ExecuteSqlRequest.QueryMode.PLAN] [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode].
	// Other fields may or may not be populated, based on the
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode].
	Stats                *ResultSetStats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResultSet) Reset()         { *m = ResultSet{} }
func (m *ResultSet) String() string { return proto.CompactTextString(m) }
func (*ResultSet) ProtoMessage()    {}
func (*ResultSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c1400ddbf24213, []int{0}
}

func (m *ResultSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultSet.Unmarshal(m, b)
}
func (m *ResultSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultSet.Marshal(b, m, deterministic)
}
func (m *ResultSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultSet.Merge(m, src)
}
func (m *ResultSet) XXX_Size() int {
	return xxx_messageInfo_ResultSet.Size(m)
}
func (m *ResultSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultSet.DiscardUnknown(m)
}

var xxx_messageInfo_ResultSet proto.InternalMessageInfo

func (m *ResultSet) GetMetadata() *ResultSetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ResultSet) GetRows() []*_struct.ListValue {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *ResultSet) GetStats() *ResultSetStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Partial results from a streaming read or SQL query. Streaming reads and
// SQL queries better tolerate large result sets, large rows, and large
// values, but are a little trickier to consume.
type PartialResultSet struct {
	// Metadata about the result set, such as row type information.
	// Only present in the first response.
	Metadata *ResultSetMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// A streamed result set consists of a stream of values, which might
	// be split into many `PartialResultSet` messages to accommodate
	// large rows and/or large values. Every N complete values defines a
	// row, where N is equal to the number of entries in
	// [metadata.row_type.fields][google.spanner.v1.StructType.fields].
	//
	// Most values are encoded based on type as described
	// [here][google.spanner.v1.TypeCode].
	//
	// It is possible that the last value in values is "chunked",
	// meaning that the rest of the value is sent in subsequent
	// `PartialResultSet`(s). This is denoted by the [chunked_value][google.spanner.v1.PartialResultSet.chunked_value]
	// field. Two or more chunked values can be merged to form a
	// complete value as follows:
	//
	//   * `bool/number/null`: cannot be chunked
	//   * `string`: concatenate the strings
	//   * `list`: concatenate the lists. If the last element in a list is a
	//     `string`, `list`, or `object`, merge it with the first element in
	//     the next list by applying these rules recursively.
	//   * `object`: concatenate the (field name, field value) pairs. If a
	//     field name is duplicated, then apply these rules recursively
	//     to merge the field values.
	//
	// Some examples of merging:
	//
	//     # Strings are concatenated.
	//     "foo", "bar" => "foobar"
	//
	//     # Lists of non-strings are concatenated.
	//     [2, 3], [4] => [2, 3, 4]
	//
	//     # Lists are concatenated, but the last and first elements are merged
	//     # because they are strings.
	//     ["a", "b"], ["c", "d"] => ["a", "bc", "d"]
	//
	//     # Lists are concatenated, but the last and first elements are merged
	//     # because they are lists. Recursively, the last and first elements
	//     # of the inner lists are merged because they are strings.
	//     ["a", ["b", "c"]], [["d"], "e"] => ["a", ["b", "cd"], "e"]
	//
	//     # Non-overlapping object fields are combined.
	//     {"a": "1"}, {"b": "2"} => {"a": "1", "b": 2"}
	//
	//     # Overlapping object fields are merged.
	//     {"a": "1"}, {"a": "2"} => {"a": "12"}
	//
	//     # Examples of merging objects containing lists of strings.
	//     {"a": ["1"]}, {"a": ["2"]} => {"a": ["12"]}
	//
	// For a more complete example, suppose a streaming SQL query is
	// yielding a result set whose rows contain a single string
	// field. The following `PartialResultSet`s might be yielded:
	//
	//     {
	//       "metadata": { ... }
	//       "values": ["Hello", "W"]
	//       "chunked_value": true
	//       "resume_token": "Af65..."
	//     }
	//     {
	//       "values": ["orl"]
	//       "chunked_value": true
	//       "resume_token": "Bqp2..."
	//     }
	//     {
	//       "values": ["d"]
	//       "resume_token": "Zx1B..."
	//     }
	//
	// This sequence of `PartialResultSet`s encodes two rows, one
	// containing the field value `"Hello"`, and a second containing the
	// field value `"World" = "W" + "orl" + "d"`.
	Values []*_struct.Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	// If true, then the final value in [values][google.spanner.v1.PartialResultSet.values] is chunked, and must
	// be combined with more values from subsequent `PartialResultSet`s
	// to obtain a complete field value.
	ChunkedValue bool `protobuf:"varint,3,opt,name=chunked_value,json=chunkedValue,proto3" json:"chunked_value,omitempty"`
	// Streaming calls might be interrupted for a variety of reasons, such
	// as TCP connection loss. If this occurs, the stream of results can
	// be resumed by re-sending the original request and including
	// `resume_token`. Note that executing any other transaction in the
	// same session invalidates the token.
	ResumeToken []byte `protobuf:"bytes,4,opt,name=resume_token,json=resumeToken,proto3" json:"resume_token,omitempty"`
	// Query plan and execution statistics for the statement that produced this
	// streaming result set. These can be requested by setting
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode] and are sent
	// only once with the last response in the stream.
	// This field will also be present in the last response for DML
	// statements.
	Stats                *ResultSetStats `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PartialResultSet) Reset()         { *m = PartialResultSet{} }
func (m *PartialResultSet) String() string { return proto.CompactTextString(m) }
func (*PartialResultSet) ProtoMessage()    {}
func (*PartialResultSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c1400ddbf24213, []int{1}
}

func (m *PartialResultSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartialResultSet.Unmarshal(m, b)
}
func (m *PartialResultSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartialResultSet.Marshal(b, m, deterministic)
}
func (m *PartialResultSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialResultSet.Merge(m, src)
}
func (m *PartialResultSet) XXX_Size() int {
	return xxx_messageInfo_PartialResultSet.Size(m)
}
func (m *PartialResultSet) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialResultSet.DiscardUnknown(m)
}

var xxx_messageInfo_PartialResultSet proto.InternalMessageInfo

func (m *PartialResultSet) GetMetadata() *ResultSetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PartialResultSet) GetValues() []*_struct.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *PartialResultSet) GetChunkedValue() bool {
	if m != nil {
		return m.ChunkedValue
	}
	return false
}

func (m *PartialResultSet) GetResumeToken() []byte {
	if m != nil {
		return m.ResumeToken
	}
	return nil
}

func (m *PartialResultSet) GetStats() *ResultSetStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Metadata about a [ResultSet][google.spanner.v1.ResultSet] or [PartialResultSet][google.spanner.v1.PartialResultSet].
type ResultSetMetadata struct {
	// Indicates the field names and types for the rows in the result
	// set.  For example, a SQL query like `"SELECT UserId, UserName FROM
	// Users"` could return a `row_type` value like:
	//
	//     "fields": [
	//       { "name": "UserId", "type": { "code": "INT64" } },
	//       { "name": "UserName", "type": { "code": "STRING" } },
	//     ]
	RowType *StructType `protobuf:"bytes,1,opt,name=row_type,json=rowType,proto3" json:"row_type,omitempty"`
	// If the read or SQL query began a transaction as a side-effect, the
	// information about the new transaction is yielded here.
	Transaction          *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResultSetMetadata) Reset()         { *m = ResultSetMetadata{} }
func (m *ResultSetMetadata) String() string { return proto.CompactTextString(m) }
func (*ResultSetMetadata) ProtoMessage()    {}
func (*ResultSetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c1400ddbf24213, []int{2}
}

func (m *ResultSetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultSetMetadata.Unmarshal(m, b)
}
func (m *ResultSetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultSetMetadata.Marshal(b, m, deterministic)
}
func (m *ResultSetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultSetMetadata.Merge(m, src)
}
func (m *ResultSetMetadata) XXX_Size() int {
	return xxx_messageInfo_ResultSetMetadata.Size(m)
}
func (m *ResultSetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultSetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ResultSetMetadata proto.InternalMessageInfo

func (m *ResultSetMetadata) GetRowType() *StructType {
	if m != nil {
		return m.RowType
	}
	return nil
}

func (m *ResultSetMetadata) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// Additional statistics about a [ResultSet][google.spanner.v1.ResultSet] or [PartialResultSet][google.spanner.v1.PartialResultSet].
type ResultSetStats struct {
	// [QueryPlan][google.spanner.v1.QueryPlan] for the query associated with this result.
	QueryPlan *QueryPlan `protobuf:"bytes,1,opt,name=query_plan,json=queryPlan,proto3" json:"query_plan,omitempty"`
	// Aggregated statistics from the execution of the query. Only present when
	// the query is profiled. For example, a query could return the statistics as
	// follows:
	//
	//     {
	//       "rows_returned": "3",
	//       "elapsed_time": "1.22 secs",
	//       "cpu_time": "1.19 secs"
	//     }
	QueryStats *_struct.Struct `protobuf:"bytes,2,opt,name=query_stats,json=queryStats,proto3" json:"query_stats,omitempty"`
	// The number of rows modified by the DML statement.
	//
	// Types that are valid to be assigned to RowCount:
	//	*ResultSetStats_RowCountExact
	//	*ResultSetStats_RowCountLowerBound
	RowCount             isResultSetStats_RowCount `protobuf_oneof:"row_count"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ResultSetStats) Reset()         { *m = ResultSetStats{} }
func (m *ResultSetStats) String() string { return proto.CompactTextString(m) }
func (*ResultSetStats) ProtoMessage()    {}
func (*ResultSetStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8c1400ddbf24213, []int{3}
}

func (m *ResultSetStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultSetStats.Unmarshal(m, b)
}
func (m *ResultSetStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultSetStats.Marshal(b, m, deterministic)
}
func (m *ResultSetStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultSetStats.Merge(m, src)
}
func (m *ResultSetStats) XXX_Size() int {
	return xxx_messageInfo_ResultSetStats.Size(m)
}
func (m *ResultSetStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultSetStats.DiscardUnknown(m)
}

var xxx_messageInfo_ResultSetStats proto.InternalMessageInfo

func (m *ResultSetStats) GetQueryPlan() *QueryPlan {
	if m != nil {
		return m.QueryPlan
	}
	return nil
}

func (m *ResultSetStats) GetQueryStats() *_struct.Struct {
	if m != nil {
		return m.QueryStats
	}
	return nil
}

type isResultSetStats_RowCount interface {
	isResultSetStats_RowCount()
}

type ResultSetStats_RowCountExact struct {
	RowCountExact int64 `protobuf:"varint,3,opt,name=row_count_exact,json=rowCountExact,proto3,oneof"`
}

type ResultSetStats_RowCountLowerBound struct {
	RowCountLowerBound int64 `protobuf:"varint,4,opt,name=row_count_lower_bound,json=rowCountLowerBound,proto3,oneof"`
}

func (*ResultSetStats_RowCountExact) isResultSetStats_RowCount() {}

func (*ResultSetStats_RowCountLowerBound) isResultSetStats_RowCount() {}

func (m *ResultSetStats) GetRowCount() isResultSetStats_RowCount {
	if m != nil {
		return m.RowCount
	}
	return nil
}

func (m *ResultSetStats) GetRowCountExact() int64 {
	if x, ok := m.GetRowCount().(*ResultSetStats_RowCountExact); ok {
		return x.RowCountExact
	}
	return 0
}

func (m *ResultSetStats) GetRowCountLowerBound() int64 {
	if x, ok := m.GetRowCount().(*ResultSetStats_RowCountLowerBound); ok {
		return x.RowCountLowerBound
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResultSetStats) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResultSetStats_RowCountExact)(nil),
		(*ResultSetStats_RowCountLowerBound)(nil),
	}
}

func init() {
	proto.RegisterType((*ResultSet)(nil), "google.spanner.v1.ResultSet")
	proto.RegisterType((*PartialResultSet)(nil), "google.spanner.v1.PartialResultSet")
	proto.RegisterType((*ResultSetMetadata)(nil), "google.spanner.v1.ResultSetMetadata")
	proto.RegisterType((*ResultSetStats)(nil), "google.spanner.v1.ResultSetStats")
}

func init() {
	proto.RegisterFile("google/spanner/v1/result_set.proto", fileDescriptor_e8c1400ddbf24213)
}

var fileDescriptor_e8c1400ddbf24213 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0xd2, 0x96, 0x64, 0x9d, 0x02, 0x5d, 0xa9, 0x34, 0x8a, 0x02, 0x4a, 0x53, 0x0e,
	0x39, 0xd9, 0x4a, 0x7b, 0x20, 0x52, 0x2f, 0x55, 0x2a, 0x04, 0x87, 0x22, 0x85, 0x4d, 0x94, 0x03,
	0x8a, 0x64, 0x6d, 0x9c, 0xc5, 0x44, 0x75, 0x76, 0xdd, 0xdd, 0x75, 0x42, 0x1e, 0x80, 0x33, 0x77,
	0x1e, 0x81, 0x07, 0xe0, 0x21, 0x78, 0x1d, 0x2e, 0x1c, 0xd1, 0xfe, 0x71, 0x12, 0xb0, 0x85, 0x84,
	0xd4, 0x9b, 0x35, 0xf3, 0xfb, 0x66, 0xe6, 0x9b, 0x1d, 0x83, 0x76, 0xc4, 0x58, 0x14, 0x13, 0x5f,
	0x24, 0x98, 0x52, 0xc2, 0xfd, 0x65, 0xd7, 0xe7, 0x44, 0xa4, 0xb1, 0x0c, 0x04, 0x91, 0x5e, 0xc2,
	0x99, 0x64, 0xf0, 0xc8, 0x30, 0x9e, 0x65, 0xbc, 0x65, 0xb7, 0xd1, 0xb4, 0x32, 0x0d, 0x4c, 0xd3,
	0x0f, 0xbe, 0x90, 0x3c, 0x0d, 0xad, 0xa0, 0x51, 0x50, 0xf4, 0x2e, 0x25, 0x7c, 0x1d, 0x24, 0x31,
	0xa6, 0x96, 0x39, 0xcb, 0x33, 0x92, 0x63, 0x2a, 0x70, 0x28, 0xe7, 0x2c, 0x83, 0x9a, 0x05, 0xd0,
	0x3a, 0x21, 0x7f, 0x65, 0x71, 0x32, 0xf7, 0x31, 0xa5, 0x4c, 0x62, 0x25, 0x15, 0x26, 0xdb, 0xfe,
	0xee, 0x80, 0x2a, 0xd2, 0x56, 0x86, 0x44, 0xc2, 0x2b, 0x50, 0x59, 0x10, 0x89, 0x67, 0x58, 0xe2,
	0xba, 0xd3, 0x72, 0x3a, 0xee, 0xf9, 0x0b, 0x2f, 0x67, 0xcb, 0xdb, 0xf0, 0x6f, 0x2d, 0x8b, 0x36,
	0x2a, 0xe8, 0x81, 0x3d, 0xce, 0x56, 0xa2, 0x5e, 0x6a, 0x95, 0x3b, 0xee, 0x79, 0x23, 0x53, 0x67,
	0x1b, 0xf0, 0x6e, 0xe6, 0x42, 0x8e, 0x71, 0x9c, 0x12, 0xa4, 0x39, 0xf8, 0x12, 0xec, 0x0b, 0x89,
	0xa5, 0xa8, 0x97, 0x75, 0xbb, 0xd3, 0x7f, 0xb5, 0x1b, 0x2a, 0x10, 0x19, 0xbe, 0xfd, 0xb9, 0x04,
	0x9e, 0x0c, 0x30, 0x97, 0x73, 0x1c, 0xdf, 0xef, 0xfc, 0x07, 0x4b, 0x35, 0x5e, 0xe6, 0xe0, 0x69,
	0xce, 0x81, 0x99, 0xde, 0x52, 0xf0, 0x0c, 0x1c, 0x86, 0x1f, 0x53, 0x7a, 0x4b, 0x66, 0x81, 0x8e,
	0x68, 0x1f, 0x15, 0x54, 0xb3, 0x41, 0x0d, 0xc3, 0x53, 0x50, 0x53, 0xe7, 0xb2, 0x20, 0x81, 0x64,
	0xb7, 0x84, 0xd6, 0xf7, 0x5a, 0x4e, 0xa7, 0x86, 0x5c, 0x13, 0x1b, 0xa9, 0xd0, 0x76, 0x0f, 0xfb,
	0xff, 0xb9, 0x87, 0x2f, 0x0e, 0x38, 0xca, 0x19, 0x82, 0x3d, 0x50, 0xe1, 0x6c, 0x15, 0xa8, 0x33,
	0xb0, 0x8b, 0x78, 0x56, 0x50, 0x71, 0xa8, 0xcf, 0x71, 0xb4, 0x4e, 0x08, 0x7a, 0xc8, 0xd9, 0x4a,
	0x7d, 0xc0, 0x2b, 0xe0, 0xee, 0x5c, 0x58, 0xbd, 0xa4, 0xc5, 0xcf, 0x0b, 0xc4, 0xa3, 0x2d, 0x85,
	0x76, 0x25, 0xed, 0x9f, 0x0e, 0x78, 0xf4, 0xe7, 0xac, 0xf0, 0x12, 0x80, 0xed, 0x69, 0xdb, 0x81,
	0x9a, 0x05, 0x35, 0xdf, 0x29, 0x68, 0x10, 0x63, 0x8a, 0xaa, 0x77, 0xd9, 0x27, 0xec, 0x01, 0xd7,
	0x88, 0xcd, 0x82, 0xcc, 0x44, 0x27, 0xb9, 0x77, 0x31, 0x66, 0x90, 0x69, 0x64, 0xda, 0x76, 0xc0,
	0x63, 0xb5, 0x85, 0x90, 0xa5, 0x54, 0x06, 0xe4, 0x13, 0x0e, 0xa5, 0x7e, 0x9e, 0xf2, 0x9b, 0x07,
	0xe8, 0x90, 0xb3, 0xd5, 0xb5, 0x8a, 0xbf, 0x52, 0x61, 0x78, 0x01, 0x8e, 0xb7, 0x64, 0xcc, 0x56,
	0x84, 0x07, 0x53, 0x96, 0xd2, 0x99, 0x7e, 0x2a, 0xc5, 0xc3, 0x8c, 0xbf, 0x51, 0xc9, 0xbe, 0xca,
	0xf5, 0x5d, 0x50, 0xdd, 0x88, 0xfa, 0x5f, 0x1d, 0x70, 0x1c, 0xb2, 0x45, 0xde, 0x54, 0x7f, 0xbb,
	0x8c, 0x81, 0x9a, 0x75, 0xe0, 0xbc, 0xef, 0x59, 0x28, 0x62, 0x31, 0xa6, 0x91, 0xc7, 0x78, 0xe4,
	0x47, 0x84, 0x6a, 0x27, 0xbe, 0x49, 0xe1, 0x64, 0x2e, 0x76, 0xfe, 0xe7, 0x4b, 0xfb, 0xf9, 0xcb,
	0x71, 0xbe, 0x95, 0x4e, 0x5e, 0x1b, 0xf5, 0x75, 0xcc, 0xd2, 0x99, 0x37, 0xb4, 0x8d, 0xc6, 0xdd,
	0x1f, 0x59, 0x66, 0xa2, 0x33, 0x13, 0x9b, 0x99, 0x8c, 0xbb, 0xd3, 0x03, 0x5d, 0xfb, 0xe2, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0xc3, 0x82, 0xaa, 0xc8, 0x04, 0x00, 0x00,
}
