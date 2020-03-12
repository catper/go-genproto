// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/spanner/v1/transaction.proto

package spanner

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	duration "github.com/catper/protobuf/ptypes/duration"
	timestamp "github.com/catper/protobuf/ptypes/timestamp"
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

// # Transactions
//
//
// Each session can have at most one active transaction at a time. After the
// active transaction is completed, the session can immediately be
// re-used for the next transaction. It is not necessary to create a
// new session for each transaction.
//
// # Transaction Modes
//
// Cloud Spanner supports three transaction modes:
//
//   1. Locking read-write. This type of transaction is the only way
//      to write data into Cloud Spanner. These transactions rely on
//      pessimistic locking and, if necessary, two-phase commit.
//      Locking read-write transactions may abort, requiring the
//      application to retry.
//
//   2. Snapshot read-only. This transaction type provides guaranteed
//      consistency across several reads, but does not allow
//      writes. Snapshot read-only transactions can be configured to
//      read at timestamps in the past. Snapshot read-only
//      transactions do not need to be committed.
//
//   3. Partitioned DML. This type of transaction is used to execute
//      a single Partitioned DML statement. Partitioned DML partitions
//      the key space and runs the DML statement over each partition
//      in parallel using separate, internal transactions that commit
//      independently. Partitioned DML transactions do not need to be
//      committed.
//
// For transactions that only read, snapshot read-only transactions
// provide simpler semantics and are almost always faster. In
// particular, read-only transactions do not take locks, so they do
// not conflict with read-write transactions. As a consequence of not
// taking locks, they also do not abort, so retry loops are not needed.
//
// Transactions may only read/write data in a single database. They
// may, however, read/write data in different tables within that
// database.
//
// ## Locking Read-Write Transactions
//
// Locking transactions may be used to atomically read-modify-write
// data anywhere in a database. This type of transaction is externally
// consistent.
//
// Clients should attempt to minimize the amount of time a transaction
// is active. Faster transactions commit with higher probability
// and cause less contention. Cloud Spanner attempts to keep read locks
// active as long as the transaction continues to do reads, and the
// transaction has not been terminated by
// [Commit][google.spanner.v1.Spanner.Commit] or
// [Rollback][google.spanner.v1.Spanner.Rollback].  Long periods of
// inactivity at the client may cause Cloud Spanner to release a
// transaction's locks and abort it.
//
// Conceptually, a read-write transaction consists of zero or more
// reads or SQL statements followed by
// [Commit][google.spanner.v1.Spanner.Commit]. At any time before
// [Commit][google.spanner.v1.Spanner.Commit], the client can send a
// [Rollback][google.spanner.v1.Spanner.Rollback] request to abort the
// transaction.
//
// ### Semantics
//
// Cloud Spanner can commit the transaction if all read locks it acquired
// are still valid at commit time, and it is able to acquire write
// locks for all writes. Cloud Spanner can abort the transaction for any
// reason. If a commit attempt returns `ABORTED`, Cloud Spanner guarantees
// that the transaction has not modified any user data in Cloud Spanner.
//
// Unless the transaction commits, Cloud Spanner makes no guarantees about
// how long the transaction's locks were held for. It is an error to
// use Cloud Spanner locks for any sort of mutual exclusion other than
// between Cloud Spanner transactions themselves.
//
// ### Retrying Aborted Transactions
//
// When a transaction aborts, the application can choose to retry the
// whole transaction again. To maximize the chances of successfully
// committing the retry, the client should execute the retry in the
// same session as the original attempt. The original session's lock
// priority increases with each consecutive abort, meaning that each
// attempt has a slightly better chance of success than the previous.
//
// Under some circumstances (e.g., many transactions attempting to
// modify the same row(s)), a transaction can abort many times in a
// short period before successfully committing. Thus, it is not a good
// idea to cap the number of retries a transaction can attempt;
// instead, it is better to limit the total amount of wall time spent
// retrying.
//
// ### Idle Transactions
//
// A transaction is considered idle if it has no outstanding reads or
// SQL queries and has not started a read or SQL query within the last 10
// seconds. Idle transactions can be aborted by Cloud Spanner so that they
// don't hold on to locks indefinitely. In that case, the commit will
// fail with error `ABORTED`.
//
// If this behavior is undesirable, periodically executing a simple
// SQL query in the transaction (e.g., `SELECT 1`) prevents the
// transaction from becoming idle.
//
// ## Snapshot Read-Only Transactions
//
// Snapshot read-only transactions provides a simpler method than
// locking read-write transactions for doing several consistent
// reads. However, this type of transaction does not support writes.
//
// Snapshot transactions do not take locks. Instead, they work by
// choosing a Cloud Spanner timestamp, then executing all reads at that
// timestamp. Since they do not acquire locks, they do not block
// concurrent read-write transactions.
//
// Unlike locking read-write transactions, snapshot read-only
// transactions never abort. They can fail if the chosen read
// timestamp is garbage collected; however, the default garbage
// collection policy is generous enough that most applications do not
// need to worry about this in practice.
//
// Snapshot read-only transactions do not need to call
// [Commit][google.spanner.v1.Spanner.Commit] or
// [Rollback][google.spanner.v1.Spanner.Rollback] (and in fact are not
// permitted to do so).
//
// To execute a snapshot transaction, the client specifies a timestamp
// bound, which tells Cloud Spanner how to choose a read timestamp.
//
// The types of timestamp bound are:
//
//   - Strong (the default).
//   - Bounded staleness.
//   - Exact staleness.
//
// If the Cloud Spanner database to be read is geographically distributed,
// stale read-only transactions can execute more quickly than strong
// or read-write transaction, because they are able to execute far
// from the leader replica.
//
// Each type of timestamp bound is discussed in detail below.
//
// ### Strong
//
// Strong reads are guaranteed to see the effects of all transactions
// that have committed before the start of the read. Furthermore, all
// rows yielded by a single read are consistent with each other -- if
// any part of the read observes a transaction, all parts of the read
// see the transaction.
//
// Strong reads are not repeatable: two consecutive strong read-only
// transactions might return inconsistent results if there are
// concurrent writes. If consistency across reads is required, the
// reads should be executed within a transaction or at an exact read
// timestamp.
//
// See [TransactionOptions.ReadOnly.strong][google.spanner.v1.TransactionOptions.ReadOnly.strong].
//
// ### Exact Staleness
//
// These timestamp bounds execute reads at a user-specified
// timestamp. Reads at a timestamp are guaranteed to see a consistent
// prefix of the global transaction history: they observe
// modifications done by all transactions with a commit timestamp <=
// the read timestamp, and observe none of the modifications done by
// transactions with a larger commit timestamp. They will block until
// all conflicting transactions that may be assigned commit timestamps
// <= the read timestamp have finished.
//
// The timestamp can either be expressed as an absolute Cloud Spanner commit
// timestamp or a staleness relative to the current time.
//
// These modes do not require a "negotiation phase" to pick a
// timestamp. As a result, they execute slightly faster than the
// equivalent boundedly stale concurrency modes. On the other hand,
// boundedly stale reads usually return fresher results.
//
// See [TransactionOptions.ReadOnly.read_timestamp][google.spanner.v1.TransactionOptions.ReadOnly.read_timestamp] and
// [TransactionOptions.ReadOnly.exact_staleness][google.spanner.v1.TransactionOptions.ReadOnly.exact_staleness].
//
// ### Bounded Staleness
//
// Bounded staleness modes allow Cloud Spanner to pick the read timestamp,
// subject to a user-provided staleness bound. Cloud Spanner chooses the
// newest timestamp within the staleness bound that allows execution
// of the reads at the closest available replica without blocking.
//
// All rows yielded are consistent with each other -- if any part of
// the read observes a transaction, all parts of the read see the
// transaction. Boundedly stale reads are not repeatable: two stale
// reads, even if they use the same staleness bound, can execute at
// different timestamps and thus return inconsistent results.
//
// Boundedly stale reads execute in two phases: the first phase
// negotiates a timestamp among all replicas needed to serve the
// read. In the second phase, reads are executed at the negotiated
// timestamp.
//
// As a result of the two phase execution, bounded staleness reads are
// usually a little slower than comparable exact staleness
// reads. However, they are typically able to return fresher
// results, and are more likely to execute at the closest replica.
//
// Because the timestamp negotiation requires up-front knowledge of
// which rows will be read, it can only be used with single-use
// read-only transactions.
//
// See [TransactionOptions.ReadOnly.max_staleness][google.spanner.v1.TransactionOptions.ReadOnly.max_staleness] and
// [TransactionOptions.ReadOnly.min_read_timestamp][google.spanner.v1.TransactionOptions.ReadOnly.min_read_timestamp].
//
// ### Old Read Timestamps and Garbage Collection
//
// Cloud Spanner continuously garbage collects deleted and overwritten data
// in the background to reclaim storage space. This process is known
// as "version GC". By default, version GC reclaims versions after they
// are one hour old. Because of this, Cloud Spanner cannot perform reads
// at read timestamps more than one hour in the past. This
// restriction also applies to in-progress reads and/or SQL queries whose
// timestamp become too old while executing. Reads and SQL queries with
// too-old read timestamps fail with the error `FAILED_PRECONDITION`.
//
// ## Partitioned DML Transactions
//
// Partitioned DML transactions are used to execute DML statements with a
// different execution strategy that provides different, and often better,
// scalability properties for large, table-wide operations than DML in a
// ReadWrite transaction. Smaller scoped statements, such as an OLTP workload,
// should prefer using ReadWrite transactions.
//
// Partitioned DML partitions the keyspace and runs the DML statement on each
// partition in separate, internal transactions. These transactions commit
// automatically when complete, and run independently from one another.
//
// To reduce lock contention, this execution strategy only acquires read locks
// on rows that match the WHERE clause of the statement. Additionally, the
// smaller per-partition transactions hold locks for less time.
//
// That said, Partitioned DML is not a drop-in replacement for standard DML used
// in ReadWrite transactions.
//
//  - The DML statement must be fully-partitionable. Specifically, the statement
//    must be expressible as the union of many statements which each access only
//    a single row of the table.
//
//  - The statement is not applied atomically to all rows of the table. Rather,
//    the statement is applied atomically to partitions of the table, in
//    independent transactions. Secondary index rows are updated atomically
//    with the base table rows.
//
//  - Partitioned DML does not guarantee exactly-once execution semantics
//    against a partition. The statement will be applied at least once to each
//    partition. It is strongly recommended that the DML statement should be
//    idempotent to avoid unexpected results. For instance, it is potentially
//    dangerous to run a statement such as
//    `UPDATE table SET column = column + 1` as it could be run multiple times
//    against some rows.
//
//  - The partitions are committed automatically - there is no support for
//    Commit or Rollback. If the call returns an error, or if the client issuing
//    the ExecuteSql call dies, it is possible that some rows had the statement
//    executed on them successfully. It is also possible that statement was
//    never executed against other rows.
//
//  - Partitioned DML transactions may only contain the execution of a single
//    DML statement via ExecuteSql or ExecuteStreamingSql.
//
//  - If any error is encountered during the execution of the partitioned DML
//    operation (for instance, a UNIQUE INDEX violation, division by zero, or a
//    value that cannot be stored due to schema constraints), then the
//    operation is stopped at that point and an error is returned. It is
//    possible that at this point, some partitions have been committed (or even
//    committed multiple times), and other partitions have not been run at all.
//
// Given the above, Partitioned DML is good fit for large, database-wide,
// operations that are idempotent, such as deleting old rows from a very large
// table.
type TransactionOptions struct {
	// Required. The type of transaction.
	//
	// Types that are valid to be assigned to Mode:
	//	*TransactionOptions_ReadWrite_
	//	*TransactionOptions_PartitionedDml_
	//	*TransactionOptions_ReadOnly_
	Mode                 isTransactionOptions_Mode `protobuf_oneof:"mode"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TransactionOptions) Reset()         { *m = TransactionOptions{} }
func (m *TransactionOptions) String() string { return proto.CompactTextString(m) }
func (*TransactionOptions) ProtoMessage()    {}
func (*TransactionOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5743daa0b72b00f, []int{0}
}

func (m *TransactionOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOptions.Unmarshal(m, b)
}
func (m *TransactionOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOptions.Marshal(b, m, deterministic)
}
func (m *TransactionOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOptions.Merge(m, src)
}
func (m *TransactionOptions) XXX_Size() int {
	return xxx_messageInfo_TransactionOptions.Size(m)
}
func (m *TransactionOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOptions proto.InternalMessageInfo

type isTransactionOptions_Mode interface {
	isTransactionOptions_Mode()
}

type TransactionOptions_ReadWrite_ struct {
	ReadWrite *TransactionOptions_ReadWrite `protobuf:"bytes,1,opt,name=read_write,json=readWrite,proto3,oneof"`
}

type TransactionOptions_PartitionedDml_ struct {
	PartitionedDml *TransactionOptions_PartitionedDml `protobuf:"bytes,3,opt,name=partitioned_dml,json=partitionedDml,proto3,oneof"`
}

type TransactionOptions_ReadOnly_ struct {
	ReadOnly *TransactionOptions_ReadOnly `protobuf:"bytes,2,opt,name=read_only,json=readOnly,proto3,oneof"`
}

func (*TransactionOptions_ReadWrite_) isTransactionOptions_Mode() {}

func (*TransactionOptions_PartitionedDml_) isTransactionOptions_Mode() {}

func (*TransactionOptions_ReadOnly_) isTransactionOptions_Mode() {}

func (m *TransactionOptions) GetMode() isTransactionOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *TransactionOptions) GetReadWrite() *TransactionOptions_ReadWrite {
	if x, ok := m.GetMode().(*TransactionOptions_ReadWrite_); ok {
		return x.ReadWrite
	}
	return nil
}

func (m *TransactionOptions) GetPartitionedDml() *TransactionOptions_PartitionedDml {
	if x, ok := m.GetMode().(*TransactionOptions_PartitionedDml_); ok {
		return x.PartitionedDml
	}
	return nil
}

func (m *TransactionOptions) GetReadOnly() *TransactionOptions_ReadOnly {
	if x, ok := m.GetMode().(*TransactionOptions_ReadOnly_); ok {
		return x.ReadOnly
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TransactionOptions) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TransactionOptions_ReadWrite_)(nil),
		(*TransactionOptions_PartitionedDml_)(nil),
		(*TransactionOptions_ReadOnly_)(nil),
	}
}

// Message type to initiate a read-write transaction. Currently this
// transaction type has no options.
type TransactionOptions_ReadWrite struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionOptions_ReadWrite) Reset()         { *m = TransactionOptions_ReadWrite{} }
func (m *TransactionOptions_ReadWrite) String() string { return proto.CompactTextString(m) }
func (*TransactionOptions_ReadWrite) ProtoMessage()    {}
func (*TransactionOptions_ReadWrite) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5743daa0b72b00f, []int{0, 0}
}

func (m *TransactionOptions_ReadWrite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOptions_ReadWrite.Unmarshal(m, b)
}
func (m *TransactionOptions_ReadWrite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOptions_ReadWrite.Marshal(b, m, deterministic)
}
func (m *TransactionOptions_ReadWrite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOptions_ReadWrite.Merge(m, src)
}
func (m *TransactionOptions_ReadWrite) XXX_Size() int {
	return xxx_messageInfo_TransactionOptions_ReadWrite.Size(m)
}
func (m *TransactionOptions_ReadWrite) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOptions_ReadWrite.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOptions_ReadWrite proto.InternalMessageInfo

// Message type to initiate a Partitioned DML transaction.
type TransactionOptions_PartitionedDml struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionOptions_PartitionedDml) Reset()         { *m = TransactionOptions_PartitionedDml{} }
func (m *TransactionOptions_PartitionedDml) String() string { return proto.CompactTextString(m) }
func (*TransactionOptions_PartitionedDml) ProtoMessage()    {}
func (*TransactionOptions_PartitionedDml) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5743daa0b72b00f, []int{0, 1}
}

func (m *TransactionOptions_PartitionedDml) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOptions_PartitionedDml.Unmarshal(m, b)
}
func (m *TransactionOptions_PartitionedDml) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOptions_PartitionedDml.Marshal(b, m, deterministic)
}
func (m *TransactionOptions_PartitionedDml) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOptions_PartitionedDml.Merge(m, src)
}
func (m *TransactionOptions_PartitionedDml) XXX_Size() int {
	return xxx_messageInfo_TransactionOptions_PartitionedDml.Size(m)
}
func (m *TransactionOptions_PartitionedDml) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOptions_PartitionedDml.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOptions_PartitionedDml proto.InternalMessageInfo

// Message type to initiate a read-only transaction.
type TransactionOptions_ReadOnly struct {
	// How to choose the timestamp for the read-only transaction.
	//
	// Types that are valid to be assigned to TimestampBound:
	//	*TransactionOptions_ReadOnly_Strong
	//	*TransactionOptions_ReadOnly_MinReadTimestamp
	//	*TransactionOptions_ReadOnly_MaxStaleness
	//	*TransactionOptions_ReadOnly_ReadTimestamp
	//	*TransactionOptions_ReadOnly_ExactStaleness
	TimestampBound isTransactionOptions_ReadOnly_TimestampBound `protobuf_oneof:"timestamp_bound"`
	// If true, the Cloud Spanner-selected read timestamp is included in
	// the [Transaction][google.spanner.v1.Transaction] message that describes the transaction.
	ReturnReadTimestamp  bool     `protobuf:"varint,6,opt,name=return_read_timestamp,json=returnReadTimestamp,proto3" json:"return_read_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionOptions_ReadOnly) Reset()         { *m = TransactionOptions_ReadOnly{} }
func (m *TransactionOptions_ReadOnly) String() string { return proto.CompactTextString(m) }
func (*TransactionOptions_ReadOnly) ProtoMessage()    {}
func (*TransactionOptions_ReadOnly) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5743daa0b72b00f, []int{0, 2}
}

func (m *TransactionOptions_ReadOnly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOptions_ReadOnly.Unmarshal(m, b)
}
func (m *TransactionOptions_ReadOnly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOptions_ReadOnly.Marshal(b, m, deterministic)
}
func (m *TransactionOptions_ReadOnly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOptions_ReadOnly.Merge(m, src)
}
func (m *TransactionOptions_ReadOnly) XXX_Size() int {
	return xxx_messageInfo_TransactionOptions_ReadOnly.Size(m)
}
func (m *TransactionOptions_ReadOnly) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOptions_ReadOnly.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOptions_ReadOnly proto.InternalMessageInfo

type isTransactionOptions_ReadOnly_TimestampBound interface {
	isTransactionOptions_ReadOnly_TimestampBound()
}

type TransactionOptions_ReadOnly_Strong struct {
	Strong bool `protobuf:"varint,1,opt,name=strong,proto3,oneof"`
}

type TransactionOptions_ReadOnly_MinReadTimestamp struct {
	MinReadTimestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=min_read_timestamp,json=minReadTimestamp,proto3,oneof"`
}

type TransactionOptions_ReadOnly_MaxStaleness struct {
	MaxStaleness *duration.Duration `protobuf:"bytes,3,opt,name=max_staleness,json=maxStaleness,proto3,oneof"`
}

type TransactionOptions_ReadOnly_ReadTimestamp struct {
	ReadTimestamp *timestamp.Timestamp `protobuf:"bytes,4,opt,name=read_timestamp,json=readTimestamp,proto3,oneof"`
}

type TransactionOptions_ReadOnly_ExactStaleness struct {
	ExactStaleness *duration.Duration `protobuf:"bytes,5,opt,name=exact_staleness,json=exactStaleness,proto3,oneof"`
}

func (*TransactionOptions_ReadOnly_Strong) isTransactionOptions_ReadOnly_TimestampBound() {}

func (*TransactionOptions_ReadOnly_MinReadTimestamp) isTransactionOptions_ReadOnly_TimestampBound() {}

func (*TransactionOptions_ReadOnly_MaxStaleness) isTransactionOptions_ReadOnly_TimestampBound() {}

func (*TransactionOptions_ReadOnly_ReadTimestamp) isTransactionOptions_ReadOnly_TimestampBound() {}

func (*TransactionOptions_ReadOnly_ExactStaleness) isTransactionOptions_ReadOnly_TimestampBound() {}

func (m *TransactionOptions_ReadOnly) GetTimestampBound() isTransactionOptions_ReadOnly_TimestampBound {
	if m != nil {
		return m.TimestampBound
	}
	return nil
}

func (m *TransactionOptions_ReadOnly) GetStrong() bool {
	if x, ok := m.GetTimestampBound().(*TransactionOptions_ReadOnly_Strong); ok {
		return x.Strong
	}
	return false
}

func (m *TransactionOptions_ReadOnly) GetMinReadTimestamp() *timestamp.Timestamp {
	if x, ok := m.GetTimestampBound().(*TransactionOptions_ReadOnly_MinReadTimestamp); ok {
		return x.MinReadTimestamp
	}
	return nil
}

func (m *TransactionOptions_ReadOnly) GetMaxStaleness() *duration.Duration {
	if x, ok := m.GetTimestampBound().(*TransactionOptions_ReadOnly_MaxStaleness); ok {
		return x.MaxStaleness
	}
	return nil
}

func (m *TransactionOptions_ReadOnly) GetReadTimestamp() *timestamp.Timestamp {
	if x, ok := m.GetTimestampBound().(*TransactionOptions_ReadOnly_ReadTimestamp); ok {
		return x.ReadTimestamp
	}
	return nil
}

func (m *TransactionOptions_ReadOnly) GetExactStaleness() *duration.Duration {
	if x, ok := m.GetTimestampBound().(*TransactionOptions_ReadOnly_ExactStaleness); ok {
		return x.ExactStaleness
	}
	return nil
}

func (m *TransactionOptions_ReadOnly) GetReturnReadTimestamp() bool {
	if m != nil {
		return m.ReturnReadTimestamp
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TransactionOptions_ReadOnly) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TransactionOptions_ReadOnly_Strong)(nil),
		(*TransactionOptions_ReadOnly_MinReadTimestamp)(nil),
		(*TransactionOptions_ReadOnly_MaxStaleness)(nil),
		(*TransactionOptions_ReadOnly_ReadTimestamp)(nil),
		(*TransactionOptions_ReadOnly_ExactStaleness)(nil),
	}
}

// A transaction.
type Transaction struct {
	// `id` may be used to identify the transaction in subsequent
	// [Read][google.spanner.v1.Spanner.Read],
	// [ExecuteSql][google.spanner.v1.Spanner.ExecuteSql],
	// [Commit][google.spanner.v1.Spanner.Commit], or
	// [Rollback][google.spanner.v1.Spanner.Rollback] calls.
	//
	// Single-use read-only transactions do not have IDs, because
	// single-use transactions do not support multiple requests.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// For snapshot read-only transactions, the read timestamp chosen
	// for the transaction. Not returned by default: see
	// [TransactionOptions.ReadOnly.return_read_timestamp][google.spanner.v1.TransactionOptions.ReadOnly.return_read_timestamp].
	//
	// A timestamp in RFC3339 UTC \"Zulu\" format, accurate to nanoseconds.
	// Example: `"2014-10-02T15:01:23.045123456Z"`.
	ReadTimestamp        *timestamp.Timestamp `protobuf:"bytes,2,opt,name=read_timestamp,json=readTimestamp,proto3" json:"read_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5743daa0b72b00f, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Transaction) GetReadTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.ReadTimestamp
	}
	return nil
}

// This message is used to select the transaction in which a
// [Read][google.spanner.v1.Spanner.Read] or
// [ExecuteSql][google.spanner.v1.Spanner.ExecuteSql] call runs.
//
// See [TransactionOptions][google.spanner.v1.TransactionOptions] for more information about transactions.
type TransactionSelector struct {
	// If no fields are set, the default is a single use transaction
	// with strong concurrency.
	//
	// Types that are valid to be assigned to Selector:
	//	*TransactionSelector_SingleUse
	//	*TransactionSelector_Id
	//	*TransactionSelector_Begin
	Selector             isTransactionSelector_Selector `protobuf_oneof:"selector"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TransactionSelector) Reset()         { *m = TransactionSelector{} }
func (m *TransactionSelector) String() string { return proto.CompactTextString(m) }
func (*TransactionSelector) ProtoMessage()    {}
func (*TransactionSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5743daa0b72b00f, []int{2}
}

func (m *TransactionSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionSelector.Unmarshal(m, b)
}
func (m *TransactionSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionSelector.Marshal(b, m, deterministic)
}
func (m *TransactionSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionSelector.Merge(m, src)
}
func (m *TransactionSelector) XXX_Size() int {
	return xxx_messageInfo_TransactionSelector.Size(m)
}
func (m *TransactionSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionSelector.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionSelector proto.InternalMessageInfo

type isTransactionSelector_Selector interface {
	isTransactionSelector_Selector()
}

type TransactionSelector_SingleUse struct {
	SingleUse *TransactionOptions `protobuf:"bytes,1,opt,name=single_use,json=singleUse,proto3,oneof"`
}

type TransactionSelector_Id struct {
	Id []byte `protobuf:"bytes,2,opt,name=id,proto3,oneof"`
}

type TransactionSelector_Begin struct {
	Begin *TransactionOptions `protobuf:"bytes,3,opt,name=begin,proto3,oneof"`
}

func (*TransactionSelector_SingleUse) isTransactionSelector_Selector() {}

func (*TransactionSelector_Id) isTransactionSelector_Selector() {}

func (*TransactionSelector_Begin) isTransactionSelector_Selector() {}

func (m *TransactionSelector) GetSelector() isTransactionSelector_Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *TransactionSelector) GetSingleUse() *TransactionOptions {
	if x, ok := m.GetSelector().(*TransactionSelector_SingleUse); ok {
		return x.SingleUse
	}
	return nil
}

func (m *TransactionSelector) GetId() []byte {
	if x, ok := m.GetSelector().(*TransactionSelector_Id); ok {
		return x.Id
	}
	return nil
}

func (m *TransactionSelector) GetBegin() *TransactionOptions {
	if x, ok := m.GetSelector().(*TransactionSelector_Begin); ok {
		return x.Begin
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TransactionSelector) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TransactionSelector_SingleUse)(nil),
		(*TransactionSelector_Id)(nil),
		(*TransactionSelector_Begin)(nil),
	}
}

func init() {
	proto.RegisterType((*TransactionOptions)(nil), "google.spanner.v1.TransactionOptions")
	proto.RegisterType((*TransactionOptions_ReadWrite)(nil), "google.spanner.v1.TransactionOptions.ReadWrite")
	proto.RegisterType((*TransactionOptions_PartitionedDml)(nil), "google.spanner.v1.TransactionOptions.PartitionedDml")
	proto.RegisterType((*TransactionOptions_ReadOnly)(nil), "google.spanner.v1.TransactionOptions.ReadOnly")
	proto.RegisterType((*Transaction)(nil), "google.spanner.v1.Transaction")
	proto.RegisterType((*TransactionSelector)(nil), "google.spanner.v1.TransactionSelector")
}

func init() {
	proto.RegisterFile("google/spanner/v1/transaction.proto", fileDescriptor_a5743daa0b72b00f)
}

var fileDescriptor_a5743daa0b72b00f = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5f, 0x8b, 0xd3, 0x4c,
	0x14, 0xc6, 0xd3, 0xee, 0x6e, 0xe9, 0x9e, 0x76, 0xdb, 0xee, 0x2c, 0xcb, 0xdb, 0x37, 0x88, 0x4a,
	0x45, 0xf0, 0x2a, 0xa1, 0xab, 0x17, 0x82, 0x08, 0xda, 0x2d, 0x1a, 0x04, 0xd9, 0x92, 0xae, 0x2b,
	0x48, 0x21, 0x4e, 0x9b, 0x31, 0x0c, 0x24, 0x33, 0x61, 0x66, 0xb2, 0x76, 0xef, 0xfd, 0x12, 0x7e,
	0x05, 0x3f, 0x82, 0xd7, 0x5e, 0xf9, 0xa9, 0x24, 0x93, 0x49, 0xff, 0xe5, 0x62, 0x7b, 0x97, 0xe4,
	0x3c, 0xcf, 0x73, 0x7e, 0x73, 0x4e, 0x12, 0x78, 0x12, 0x71, 0x1e, 0xc5, 0xc4, 0x95, 0x29, 0x66,
	0x8c, 0x08, 0xf7, 0x76, 0xe8, 0x2a, 0x81, 0x99, 0xc4, 0x0b, 0x45, 0x39, 0x73, 0x52, 0xc1, 0x15,
	0x47, 0xa7, 0x85, 0xc8, 0x31, 0x22, 0xe7, 0x76, 0x68, 0x3f, 0x34, 0x3e, 0x2d, 0x98, 0x67, 0xdf,
	0xdc, 0x30, 0x13, 0x78, 0x6d, 0xb1, 0x1f, 0xed, 0xd6, 0x15, 0x4d, 0x88, 0x54, 0x38, 0x49, 0x8d,
	0xe0, 0x81, 0x11, 0xe0, 0x94, 0xba, 0x98, 0x31, 0xae, 0xb4, 0x5b, 0x16, 0xd5, 0xc1, 0x9f, 0x23,
	0x40, 0xd7, 0x6b, 0x8e, 0xab, 0x54, 0x17, 0xd1, 0x04, 0x40, 0x10, 0x1c, 0x06, 0xdf, 0x05, 0x55,
	0xa4, 0x5f, 0x7b, 0x5c, 0x7b, 0xd6, 0xba, 0x70, 0x9d, 0x0a, 0x9d, 0x53, 0xb5, 0x3a, 0x3e, 0xc1,
	0xe1, 0xe7, 0xdc, 0xe6, 0x59, 0xfe, 0xb1, 0x28, 0x6f, 0x50, 0x00, 0xdd, 0x14, 0x0b, 0x45, 0x73,
	0x11, 0x09, 0x83, 0x30, 0x89, 0xfb, 0x07, 0x3a, 0xf6, 0xc5, 0x7e, 0xb1, 0x93, 0xb5, 0x79, 0x9c,
	0xc4, 0x9e, 0xe5, 0x77, 0xd2, 0xad, 0x27, 0xe8, 0x23, 0xe8, 0x6e, 0x01, 0x67, 0xf1, 0x5d, 0xbf,
	0xae, 0xa3, 0x9d, 0xfd, 0x89, 0xaf, 0x58, 0x7c, 0xe7, 0x59, 0x7e, 0x53, 0x98, 0x6b, 0xbb, 0x05,
	0xc7, 0xab, 0x93, 0xd8, 0x3d, 0xe8, 0x6c, 0xf7, 0xb7, 0x7f, 0x1c, 0x40, 0xb3, 0xf4, 0xa1, 0x3e,
	0x34, 0xa4, 0x12, 0x9c, 0x45, 0x7a, 0x52, 0x4d, 0xcf, 0xf2, 0xcd, 0x3d, 0xfa, 0x00, 0x28, 0xa1,
	0x2c, 0xd0, 0x60, 0xab, 0xc5, 0x18, 0x3a, 0xbb, 0xa4, 0x2b, 0x57, 0xe7, 0x5c, 0x97, 0x0a, 0xcf,
	0xf2, 0x7b, 0x09, 0x65, 0x79, 0x83, 0xd5, 0x33, 0xf4, 0x06, 0x4e, 0x12, 0xbc, 0x0c, 0xa4, 0xc2,
	0x31, 0x61, 0x44, 0x4a, 0x33, 0xbf, 0xff, 0x2b, 0x31, 0x63, 0xf3, 0x86, 0x78, 0x96, 0xdf, 0x4e,
	0xf0, 0x72, 0x5a, 0x1a, 0xd0, 0x25, 0x74, 0x76, 0x48, 0x0e, 0xf7, 0x20, 0x39, 0x11, 0x5b, 0x18,
	0x63, 0xe8, 0x92, 0x25, 0x5e, 0xa8, 0x0d, 0x90, 0xa3, 0xfb, 0x41, 0x3a, 0xda, 0xb3, 0x46, 0xb9,
	0x80, 0x73, 0x41, 0x54, 0x26, 0x2a, 0xb3, 0x69, 0xe4, 0x13, 0xf4, 0xcf, 0x8a, 0xe2, 0xd6, 0x00,
	0x46, 0xa7, 0xd0, 0x5d, 0xe9, 0x82, 0x39, 0xcf, 0x58, 0x38, 0x6a, 0xc0, 0x61, 0xc2, 0x43, 0x32,
	0xf8, 0x0a, 0xad, 0x8d, 0xc5, 0xa2, 0x0e, 0xd4, 0x69, 0xa8, 0x97, 0xd1, 0xf6, 0xeb, 0x34, 0x44,
	0x6f, 0x2b, 0x07, 0xbf, 0x77, 0x05, 0x3b, 0xc7, 0x1e, 0xfc, 0xae, 0xc1, 0xd9, 0x46, 0x8b, 0x29,
	0x89, 0xc9, 0x42, 0x71, 0x81, 0xde, 0x01, 0x48, 0xca, 0xa2, 0x98, 0x04, 0x99, 0x2c, 0xbf, 0x94,
	0xa7, 0x7b, 0xbd, 0x77, 0xf9, 0xf7, 0x51, 0x58, 0x3f, 0x49, 0x82, 0x7a, 0x1a, 0x39, 0xc7, 0x6a,
	0x7b, 0x96, 0x86, 0x7e, 0x0d, 0x47, 0x73, 0x12, 0x51, 0x66, 0xf6, 0xbc, 0x77, 0x68, 0xe1, 0x1a,
	0x01, 0x34, 0xa5, 0x81, 0x1c, 0xfd, 0xac, 0xc1, 0xf9, 0x82, 0x27, 0xd5, 0x84, 0x51, 0x6f, 0x23,
	0x62, 0x92, 0x0f, 0x61, 0x52, 0xfb, 0xf2, 0xd2, 0xc8, 0x22, 0x1e, 0x63, 0x16, 0x39, 0x5c, 0x44,
	0x6e, 0x44, 0x98, 0x1e, 0x91, 0x5b, 0x94, 0x70, 0x4a, 0xe5, 0xc6, 0x9f, 0xec, 0x95, 0xb9, 0xfc,
	0x55, 0xff, 0xef, 0x7d, 0x61, 0xbd, 0x8c, 0x79, 0x16, 0x3a, 0x53, 0xd3, 0xe7, 0x66, 0xf8, 0xb7,
	0xac, 0xcc, 0x74, 0x65, 0x66, 0x2a, 0xb3, 0x9b, 0xe1, 0xbc, 0xa1, 0x83, 0x9f, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0x1b, 0xee, 0x19, 0x21, 0x05, 0x00, 0x00,
}
