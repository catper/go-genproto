// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/remoteworkers/v1test2/command.proto

package remoteworkers

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	any "github.com/catper/protobuf/ptypes/any"
	duration "github.com/catper/protobuf/ptypes/duration"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Describes a shell-style task to execute, suitable for providing as the Bots
// interface's `Lease.payload` field.
type CommandTask struct {
	// The inputs to the task.
	Inputs *CommandTask_Inputs `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// The expected outputs from the task.
	ExpectedOutputs *CommandTask_Outputs `protobuf:"bytes,4,opt,name=expected_outputs,json=expectedOutputs,proto3" json:"expected_outputs,omitempty"`
	// The timeouts of this task.
	Timeouts             *CommandTask_Timeouts `protobuf:"bytes,5,opt,name=timeouts,proto3" json:"timeouts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CommandTask) Reset()         { *m = CommandTask{} }
func (m *CommandTask) String() string { return proto.CompactTextString(m) }
func (*CommandTask) ProtoMessage()    {}
func (*CommandTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{0}
}

func (m *CommandTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandTask.Unmarshal(m, b)
}
func (m *CommandTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandTask.Marshal(b, m, deterministic)
}
func (m *CommandTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandTask.Merge(m, src)
}
func (m *CommandTask) XXX_Size() int {
	return xxx_messageInfo_CommandTask.Size(m)
}
func (m *CommandTask) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandTask.DiscardUnknown(m)
}

var xxx_messageInfo_CommandTask proto.InternalMessageInfo

func (m *CommandTask) GetInputs() *CommandTask_Inputs {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CommandTask) GetExpectedOutputs() *CommandTask_Outputs {
	if m != nil {
		return m.ExpectedOutputs
	}
	return nil
}

func (m *CommandTask) GetTimeouts() *CommandTask_Timeouts {
	if m != nil {
		return m.Timeouts
	}
	return nil
}

// Describes the inputs to a shell-style task.
type CommandTask_Inputs struct {
	// The command itself to run (e.g., argv).
	//
	// This field should be passed directly to the underlying operating system,
	// and so it must be sensible to that operating system. For example, on
	// Windows, the first argument might be "C:\Windows\System32\ping.exe" -
	// that is, using drive letters and backslashes. A command for a *nix
	// system, on the other hand, would use forward slashes.
	//
	// All other fields in the RWAPI must consistently use forward slashes,
	// since those fields may be interpretted by both the service and the bot.
	Arguments []string `protobuf:"bytes,1,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// The input filesystem to be set up prior to the task beginning. The
	// contents should be a repeated set of FileMetadata messages though other
	// formats are allowed if better for the implementation (eg, a LUCI-style
	// .isolated file).
	//
	// This field is repeated since implementations might want to cache the
	// metadata, in which case it may be useful to break up portions of the
	// filesystem that change frequently (eg, specific input files) from those
	// that don't (eg, standard header files).
	Files []*Digest `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	// Inline contents for blobs expected to be needed by the bot to execute the
	// task. For example, contents of entries in `files` or blobs that are
	// indirectly referenced by an entry there.
	//
	// The bot should check against this list before downloading required task
	// inputs to reduce the number of communications between itself and the
	// remote CAS server.
	InlineBlobs []*Blob `protobuf:"bytes,4,rep,name=inline_blobs,json=inlineBlobs,proto3" json:"inline_blobs,omitempty"`
	// All environment variables required by the task.
	EnvironmentVariables []*CommandTask_Inputs_EnvironmentVariable `protobuf:"bytes,3,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty"`
	// Directory from which a command is executed. It is a relative directory
	// with respect to the bot's working directory (i.e., "./"). If it is
	// non-empty, then it must exist under "./". Otherwise, "./" will be used.
	WorkingDirectory     string   `protobuf:"bytes,5,opt,name=working_directory,json=workingDirectory,proto3" json:"working_directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandTask_Inputs) Reset()         { *m = CommandTask_Inputs{} }
func (m *CommandTask_Inputs) String() string { return proto.CompactTextString(m) }
func (*CommandTask_Inputs) ProtoMessage()    {}
func (*CommandTask_Inputs) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{0, 0}
}

func (m *CommandTask_Inputs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandTask_Inputs.Unmarshal(m, b)
}
func (m *CommandTask_Inputs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandTask_Inputs.Marshal(b, m, deterministic)
}
func (m *CommandTask_Inputs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandTask_Inputs.Merge(m, src)
}
func (m *CommandTask_Inputs) XXX_Size() int {
	return xxx_messageInfo_CommandTask_Inputs.Size(m)
}
func (m *CommandTask_Inputs) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandTask_Inputs.DiscardUnknown(m)
}

var xxx_messageInfo_CommandTask_Inputs proto.InternalMessageInfo

func (m *CommandTask_Inputs) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *CommandTask_Inputs) GetFiles() []*Digest {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CommandTask_Inputs) GetInlineBlobs() []*Blob {
	if m != nil {
		return m.InlineBlobs
	}
	return nil
}

func (m *CommandTask_Inputs) GetEnvironmentVariables() []*CommandTask_Inputs_EnvironmentVariable {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

func (m *CommandTask_Inputs) GetWorkingDirectory() string {
	if m != nil {
		return m.WorkingDirectory
	}
	return ""
}

// An environment variable required by this task.
type CommandTask_Inputs_EnvironmentVariable struct {
	// The envvar name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The envvar value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandTask_Inputs_EnvironmentVariable) Reset() {
	*m = CommandTask_Inputs_EnvironmentVariable{}
}
func (m *CommandTask_Inputs_EnvironmentVariable) String() string { return proto.CompactTextString(m) }
func (*CommandTask_Inputs_EnvironmentVariable) ProtoMessage()    {}
func (*CommandTask_Inputs_EnvironmentVariable) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{0, 0, 0}
}

func (m *CommandTask_Inputs_EnvironmentVariable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandTask_Inputs_EnvironmentVariable.Unmarshal(m, b)
}
func (m *CommandTask_Inputs_EnvironmentVariable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandTask_Inputs_EnvironmentVariable.Marshal(b, m, deterministic)
}
func (m *CommandTask_Inputs_EnvironmentVariable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandTask_Inputs_EnvironmentVariable.Merge(m, src)
}
func (m *CommandTask_Inputs_EnvironmentVariable) XXX_Size() int {
	return xxx_messageInfo_CommandTask_Inputs_EnvironmentVariable.Size(m)
}
func (m *CommandTask_Inputs_EnvironmentVariable) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandTask_Inputs_EnvironmentVariable.DiscardUnknown(m)
}

var xxx_messageInfo_CommandTask_Inputs_EnvironmentVariable proto.InternalMessageInfo

func (m *CommandTask_Inputs_EnvironmentVariable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommandTask_Inputs_EnvironmentVariable) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Describes the expected outputs of the command.
type CommandTask_Outputs struct {
	// A list of expected files, relative to the execution root. All paths
	// MUST be delimited by forward slashes.
	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// A list of expected directories, relative to the execution root. All paths
	// MUST be delimited by forward slashes.
	Directories []string `protobuf:"bytes,2,rep,name=directories,proto3" json:"directories,omitempty"`
	// The destination to which any stdout should be sent. The method by which
	// the bot should send the stream contents to that destination is not
	// defined in this API. As examples, the destination could be a file
	// referenced in the `files` field in this message, or it could be a URI
	// that must be written via the ByteStream API.
	StdoutDestination string `protobuf:"bytes,3,opt,name=stdout_destination,json=stdoutDestination,proto3" json:"stdout_destination,omitempty"`
	// The destination to which any stderr should be sent. The method by which
	// the bot should send the stream contents to that destination is not
	// defined in this API. As examples, the destination could be a file
	// referenced in the `files` field in this message, or it could be a URI
	// that must be written via the ByteStream API.
	StderrDestination    string   `protobuf:"bytes,4,opt,name=stderr_destination,json=stderrDestination,proto3" json:"stderr_destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandTask_Outputs) Reset()         { *m = CommandTask_Outputs{} }
func (m *CommandTask_Outputs) String() string { return proto.CompactTextString(m) }
func (*CommandTask_Outputs) ProtoMessage()    {}
func (*CommandTask_Outputs) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{0, 1}
}

func (m *CommandTask_Outputs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandTask_Outputs.Unmarshal(m, b)
}
func (m *CommandTask_Outputs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandTask_Outputs.Marshal(b, m, deterministic)
}
func (m *CommandTask_Outputs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandTask_Outputs.Merge(m, src)
}
func (m *CommandTask_Outputs) XXX_Size() int {
	return xxx_messageInfo_CommandTask_Outputs.Size(m)
}
func (m *CommandTask_Outputs) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandTask_Outputs.DiscardUnknown(m)
}

var xxx_messageInfo_CommandTask_Outputs proto.InternalMessageInfo

func (m *CommandTask_Outputs) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CommandTask_Outputs) GetDirectories() []string {
	if m != nil {
		return m.Directories
	}
	return nil
}

func (m *CommandTask_Outputs) GetStdoutDestination() string {
	if m != nil {
		return m.StdoutDestination
	}
	return ""
}

func (m *CommandTask_Outputs) GetStderrDestination() string {
	if m != nil {
		return m.StderrDestination
	}
	return ""
}

// Describes the timeouts associated with this task.
type CommandTask_Timeouts struct {
	// This specifies the maximum time that the task can run, excluding the
	// time required to download inputs or upload outputs. That is, the worker
	// will terminate the task if it runs longer than this.
	Execution *duration.Duration `protobuf:"bytes,1,opt,name=execution,proto3" json:"execution,omitempty"`
	// This specifies the maximum amount of time the task can be idle - that is,
	// go without generating some output in either stdout or stderr. If the
	// process is silent for more than the specified time, the worker will
	// terminate the task.
	Idle *duration.Duration `protobuf:"bytes,2,opt,name=idle,proto3" json:"idle,omitempty"`
	// If the execution or IO timeouts are exceeded, the worker will try to
	// gracefully terminate the task and return any existing logs. However,
	// tasks may be hard-frozen in which case this process will fail. This
	// timeout specifies how long to wait for a terminated task to shut down
	// gracefully (e.g. via SIGTERM) before we bring down the hammer (e.g.
	// SIGKILL on *nix, CTRL_BREAK_EVENT on Windows).
	Shutdown             *duration.Duration `protobuf:"bytes,3,opt,name=shutdown,proto3" json:"shutdown,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CommandTask_Timeouts) Reset()         { *m = CommandTask_Timeouts{} }
func (m *CommandTask_Timeouts) String() string { return proto.CompactTextString(m) }
func (*CommandTask_Timeouts) ProtoMessage()    {}
func (*CommandTask_Timeouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{0, 2}
}

func (m *CommandTask_Timeouts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandTask_Timeouts.Unmarshal(m, b)
}
func (m *CommandTask_Timeouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandTask_Timeouts.Marshal(b, m, deterministic)
}
func (m *CommandTask_Timeouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandTask_Timeouts.Merge(m, src)
}
func (m *CommandTask_Timeouts) XXX_Size() int {
	return xxx_messageInfo_CommandTask_Timeouts.Size(m)
}
func (m *CommandTask_Timeouts) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandTask_Timeouts.DiscardUnknown(m)
}

var xxx_messageInfo_CommandTask_Timeouts proto.InternalMessageInfo

func (m *CommandTask_Timeouts) GetExecution() *duration.Duration {
	if m != nil {
		return m.Execution
	}
	return nil
}

func (m *CommandTask_Timeouts) GetIdle() *duration.Duration {
	if m != nil {
		return m.Idle
	}
	return nil
}

func (m *CommandTask_Timeouts) GetShutdown() *duration.Duration {
	if m != nil {
		return m.Shutdown
	}
	return nil
}

// DEPRECATED - use CommandResult instead.
// Describes the actual outputs from the task.
type CommandOutputs struct {
	// exit_code is only fully reliable if the status' code is OK. If the task
	// exceeded its deadline or was cancelled, the process may still produce an
	// exit code as it is cancelled, and this will be populated, but a successful
	// (zero) is unlikely to be correct unless the status code is OK.
	ExitCode int32 `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	// The output files. The blob referenced by the digest should contain
	// one of the following (implementation-dependent):
	//    * A marshalled DirectoryMetadata of the returned filesystem
	//    * A LUCI-style .isolated file
	Outputs              *Digest  `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandOutputs) Reset()         { *m = CommandOutputs{} }
func (m *CommandOutputs) String() string { return proto.CompactTextString(m) }
func (*CommandOutputs) ProtoMessage()    {}
func (*CommandOutputs) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{1}
}

func (m *CommandOutputs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandOutputs.Unmarshal(m, b)
}
func (m *CommandOutputs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandOutputs.Marshal(b, m, deterministic)
}
func (m *CommandOutputs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandOutputs.Merge(m, src)
}
func (m *CommandOutputs) XXX_Size() int {
	return xxx_messageInfo_CommandOutputs.Size(m)
}
func (m *CommandOutputs) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandOutputs.DiscardUnknown(m)
}

var xxx_messageInfo_CommandOutputs proto.InternalMessageInfo

func (m *CommandOutputs) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *CommandOutputs) GetOutputs() *Digest {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// DEPRECATED - use CommandResult instead.
// Can be used as part of CompleteRequest.metadata, or are part of a more
// sophisticated message.
type CommandOverhead struct {
	// The elapsed time between calling Accept and Complete. The server will also
	// have its own idea of what this should be, but this excludes the overhead of
	// the RPCs and the bot response time.
	Duration *duration.Duration `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	// The amount of time *not* spent executing the command (ie
	// uploading/downloading files).
	Overhead             *duration.Duration `protobuf:"bytes,2,opt,name=overhead,proto3" json:"overhead,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CommandOverhead) Reset()         { *m = CommandOverhead{} }
func (m *CommandOverhead) String() string { return proto.CompactTextString(m) }
func (*CommandOverhead) ProtoMessage()    {}
func (*CommandOverhead) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{2}
}

func (m *CommandOverhead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandOverhead.Unmarshal(m, b)
}
func (m *CommandOverhead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandOverhead.Marshal(b, m, deterministic)
}
func (m *CommandOverhead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandOverhead.Merge(m, src)
}
func (m *CommandOverhead) XXX_Size() int {
	return xxx_messageInfo_CommandOverhead.Size(m)
}
func (m *CommandOverhead) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandOverhead.DiscardUnknown(m)
}

var xxx_messageInfo_CommandOverhead proto.InternalMessageInfo

func (m *CommandOverhead) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *CommandOverhead) GetOverhead() *duration.Duration {
	if m != nil {
		return m.Overhead
	}
	return nil
}

// All information about the execution of a command, suitable for providing as
// the Bots interface's `Lease.result` field.
type CommandResult struct {
	// An overall status for the command. For example, if the command timed out,
	// this might have a code of DEADLINE_EXCEEDED; if it was killed by the OS for
	// memory exhaustion, it might have a code of RESOURCE_EXHAUSTED.
	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// The exit code of the process. An exit code of "0" should only be trusted if
	// `status` has a code of OK (otherwise it may simply be unset).
	ExitCode int32 `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	// The output files. The blob referenced by the digest should contain
	// one of the following (implementation-dependent):
	//    * A marshalled DirectoryMetadata of the returned filesystem
	//    * A LUCI-style .isolated file
	Outputs *Digest `protobuf:"bytes,3,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// The elapsed time between calling Accept and Complete. The server will also
	// have its own idea of what this should be, but this excludes the overhead of
	// the RPCs and the bot response time.
	Duration *duration.Duration `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"` // Deprecated: Do not use.
	// The amount of time *not* spent executing the command (ie
	// uploading/downloading files).
	Overhead *duration.Duration `protobuf:"bytes,5,opt,name=overhead,proto3" json:"overhead,omitempty"` // Deprecated: Do not use.
	// Implementation-dependent metadata about the task. Both servers and bots
	// may define messages which can be encoded here; bots are free to provide
	// metadata in multiple formats, and servers are free to choose one or more
	// of the values to process and ignore others. In particular, it is *not*
	// considered an error for the bot to provide the server with a field that it
	// doesn't know about.
	Metadata             []*any.Any `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CommandResult) Reset()         { *m = CommandResult{} }
func (m *CommandResult) String() string { return proto.CompactTextString(m) }
func (*CommandResult) ProtoMessage()    {}
func (*CommandResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{3}
}

func (m *CommandResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResult.Unmarshal(m, b)
}
func (m *CommandResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResult.Marshal(b, m, deterministic)
}
func (m *CommandResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResult.Merge(m, src)
}
func (m *CommandResult) XXX_Size() int {
	return xxx_messageInfo_CommandResult.Size(m)
}
func (m *CommandResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResult.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResult proto.InternalMessageInfo

func (m *CommandResult) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CommandResult) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *CommandResult) GetOutputs() *Digest {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// Deprecated: Do not use.
func (m *CommandResult) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

// Deprecated: Do not use.
func (m *CommandResult) GetOverhead() *duration.Duration {
	if m != nil {
		return m.Overhead
	}
	return nil
}

func (m *CommandResult) GetMetadata() []*any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// The metadata for a file. Similar to the equivalent message in the Remote
// Execution API.
type FileMetadata struct {
	// The path of this file. If this message is part of the
	// CommandOutputs.outputs fields, the path is relative to the execution root
	// and must correspond to an entry in CommandTask.outputs.files. If this
	// message is part of a Directory message, then the path is relative to the
	// root of that directory. All paths MUST be delimited by forward slashes.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// A pointer to the contents of the file. The method by which a client
	// retrieves the contents from a CAS system is not defined here.
	Digest *Digest `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	// If the file is small enough, its contents may also or alternatively be
	// listed here.
	Contents []byte `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
	// Properties of the file
	IsExecutable         bool     `protobuf:"varint,4,opt,name=is_executable,json=isExecutable,proto3" json:"is_executable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileMetadata) Reset()         { *m = FileMetadata{} }
func (m *FileMetadata) String() string { return proto.CompactTextString(m) }
func (*FileMetadata) ProtoMessage()    {}
func (*FileMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{4}
}

func (m *FileMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileMetadata.Unmarshal(m, b)
}
func (m *FileMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileMetadata.Marshal(b, m, deterministic)
}
func (m *FileMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileMetadata.Merge(m, src)
}
func (m *FileMetadata) XXX_Size() int {
	return xxx_messageInfo_FileMetadata.Size(m)
}
func (m *FileMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_FileMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_FileMetadata proto.InternalMessageInfo

func (m *FileMetadata) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileMetadata) GetDigest() *Digest {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *FileMetadata) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *FileMetadata) GetIsExecutable() bool {
	if m != nil {
		return m.IsExecutable
	}
	return false
}

// The metadata for a directory. Similar to the equivalent message in the Remote
// Execution API.
type DirectoryMetadata struct {
	// The path of the directory, as in [FileMetadata.path][google.devtools.remoteworkers.v1test2.FileMetadata.path].
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// A pointer to the contents of the directory, in the form of a marshalled
	// Directory message.
	Digest               *Digest  `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectoryMetadata) Reset()         { *m = DirectoryMetadata{} }
func (m *DirectoryMetadata) String() string { return proto.CompactTextString(m) }
func (*DirectoryMetadata) ProtoMessage()    {}
func (*DirectoryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{5}
}

func (m *DirectoryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectoryMetadata.Unmarshal(m, b)
}
func (m *DirectoryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectoryMetadata.Marshal(b, m, deterministic)
}
func (m *DirectoryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectoryMetadata.Merge(m, src)
}
func (m *DirectoryMetadata) XXX_Size() int {
	return xxx_messageInfo_DirectoryMetadata.Size(m)
}
func (m *DirectoryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectoryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DirectoryMetadata proto.InternalMessageInfo

func (m *DirectoryMetadata) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DirectoryMetadata) GetDigest() *Digest {
	if m != nil {
		return m.Digest
	}
	return nil
}

// The CommandTask and CommandResult messages assume the existence of a service
// that can serve blobs of content, identified by a hash and size known as a
// "digest." The method by which these blobs may be retrieved is not specified
// here, but a model implementation is in the Remote Execution API's
// "ContentAddressibleStorage" interface.
//
// In the context of the RWAPI, a Digest will virtually always refer to the
// contents of a file or a directory. The latter is represented by the
// byte-encoded Directory message.
type Digest struct {
	// A string-encoded hash (eg "1a2b3c", not the byte array [0x1a, 0x2b, 0x3c])
	// using an implementation-defined hash algorithm (eg SHA-256).
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// The size of the contents. While this is not strictly required as part of an
	// identifier (after all, any given hash will have exactly one canonical
	// size), it's useful in almost all cases when one might want to send or
	// retrieve blobs of content and is included here for this reason.
	SizeBytes            int64    `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Digest) Reset()         { *m = Digest{} }
func (m *Digest) String() string { return proto.CompactTextString(m) }
func (*Digest) ProtoMessage()    {}
func (*Digest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{6}
}

func (m *Digest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Digest.Unmarshal(m, b)
}
func (m *Digest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Digest.Marshal(b, m, deterministic)
}
func (m *Digest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Digest.Merge(m, src)
}
func (m *Digest) XXX_Size() int {
	return xxx_messageInfo_Digest.Size(m)
}
func (m *Digest) XXX_DiscardUnknown() {
	xxx_messageInfo_Digest.DiscardUnknown(m)
}

var xxx_messageInfo_Digest proto.InternalMessageInfo

func (m *Digest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Digest) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

// Describes a blob of binary content with its digest.
type Blob struct {
	// The digest of the blob. This should be verified by the receiver.
	Digest *Digest `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	// The contents of the blob.
	Contents             []byte   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blob) Reset()         { *m = Blob{} }
func (m *Blob) String() string { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()    {}
func (*Blob) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{7}
}

func (m *Blob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blob.Unmarshal(m, b)
}
func (m *Blob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blob.Marshal(b, m, deterministic)
}
func (m *Blob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blob.Merge(m, src)
}
func (m *Blob) XXX_Size() int {
	return xxx_messageInfo_Blob.Size(m)
}
func (m *Blob) XXX_DiscardUnknown() {
	xxx_messageInfo_Blob.DiscardUnknown(m)
}

var xxx_messageInfo_Blob proto.InternalMessageInfo

func (m *Blob) GetDigest() *Digest {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Blob) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

// The contents of a directory. Similar to the equivalent message in the Remote
// Execution API.
type Directory struct {
	// The files in this directory
	Files []*FileMetadata `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// Any subdirectories
	Directories          []*DirectoryMetadata `protobuf:"bytes,2,rep,name=directories,proto3" json:"directories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Directory) Reset()         { *m = Directory{} }
func (m *Directory) String() string { return proto.CompactTextString(m) }
func (*Directory) ProtoMessage()    {}
func (*Directory) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b7f8597ab32e88e, []int{8}
}

func (m *Directory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Directory.Unmarshal(m, b)
}
func (m *Directory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Directory.Marshal(b, m, deterministic)
}
func (m *Directory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Directory.Merge(m, src)
}
func (m *Directory) XXX_Size() int {
	return xxx_messageInfo_Directory.Size(m)
}
func (m *Directory) XXX_DiscardUnknown() {
	xxx_messageInfo_Directory.DiscardUnknown(m)
}

var xxx_messageInfo_Directory proto.InternalMessageInfo

func (m *Directory) GetFiles() []*FileMetadata {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Directory) GetDirectories() []*DirectoryMetadata {
	if m != nil {
		return m.Directories
	}
	return nil
}

func init() {
	proto.RegisterType((*CommandTask)(nil), "google.devtools.remoteworkers.v1test2.CommandTask")
	proto.RegisterType((*CommandTask_Inputs)(nil), "google.devtools.remoteworkers.v1test2.CommandTask.Inputs")
	proto.RegisterType((*CommandTask_Inputs_EnvironmentVariable)(nil), "google.devtools.remoteworkers.v1test2.CommandTask.Inputs.EnvironmentVariable")
	proto.RegisterType((*CommandTask_Outputs)(nil), "google.devtools.remoteworkers.v1test2.CommandTask.Outputs")
	proto.RegisterType((*CommandTask_Timeouts)(nil), "google.devtools.remoteworkers.v1test2.CommandTask.Timeouts")
	proto.RegisterType((*CommandOutputs)(nil), "google.devtools.remoteworkers.v1test2.CommandOutputs")
	proto.RegisterType((*CommandOverhead)(nil), "google.devtools.remoteworkers.v1test2.CommandOverhead")
	proto.RegisterType((*CommandResult)(nil), "google.devtools.remoteworkers.v1test2.CommandResult")
	proto.RegisterType((*FileMetadata)(nil), "google.devtools.remoteworkers.v1test2.FileMetadata")
	proto.RegisterType((*DirectoryMetadata)(nil), "google.devtools.remoteworkers.v1test2.DirectoryMetadata")
	proto.RegisterType((*Digest)(nil), "google.devtools.remoteworkers.v1test2.Digest")
	proto.RegisterType((*Blob)(nil), "google.devtools.remoteworkers.v1test2.Blob")
	proto.RegisterType((*Directory)(nil), "google.devtools.remoteworkers.v1test2.Directory")
}

func init() {
	proto.RegisterFile("google/devtools/remoteworkers/v1test2/command.proto", fileDescriptor_7b7f8597ab32e88e)
}

var fileDescriptor_7b7f8597ab32e88e = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5b, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0xae, 0x2f, 0xb5, 0x8f, 0x53, 0xda, 0x0c, 0xa9, 0x70, 0xcd, 0x45, 0xd1, 0xa2, 0x4a,
	0x81, 0x2a, 0x6b, 0xea, 0x08, 0x41, 0xc8, 0x03, 0xc2, 0x71, 0xa8, 0xfa, 0x50, 0x2e, 0x43, 0xd4,
	0x48, 0x7d, 0xb1, 0xc6, 0xbb, 0xd3, 0xcd, 0xa8, 0xeb, 0x1d, 0x6b, 0x66, 0xd6, 0x6d, 0x78, 0x41,
	0xe2, 0x8f, 0x54, 0xe2, 0x8d, 0x3e, 0xf0, 0xc0, 0x7f, 0xe0, 0x81, 0x7f, 0x85, 0xe6, 0xb6, 0x59,
	0x37, 0xd0, 0x98, 0x80, 0xfa, 0xb6, 0xfe, 0xce, 0xf9, 0xbe, 0x39, 0xf3, 0x9d, 0x39, 0x47, 0x86,
	0xbd, 0x8c, 0xf3, 0x2c, 0xa7, 0xc3, 0x94, 0x2e, 0x15, 0xe7, 0xb9, 0x1c, 0x0a, 0x3a, 0xe7, 0x8a,
	0x3e, 0xe3, 0xe2, 0x29, 0x15, 0x72, 0xb8, 0xbc, 0xa7, 0xa8, 0x54, 0xa3, 0x61, 0xc2, 0xe7, 0x73,
	0x52, 0xa4, 0xf1, 0x42, 0x70, 0xc5, 0xd1, 0x1d, 0x4b, 0x8a, 0x3d, 0x29, 0x5e, 0x21, 0xc5, 0x8e,
	0x34, 0xb8, 0xed, 0xb4, 0x0d, 0x69, 0x56, 0x3e, 0x19, 0x92, 0xe2, 0xcc, 0x2a, 0x0c, 0x3e, 0x78,
	0x35, 0x94, 0x96, 0x82, 0x28, 0xc6, 0x0b, 0x17, 0x7f, 0xc7, 0xc5, 0xc5, 0x22, 0x19, 0x4a, 0x45,
	0x54, 0x29, 0x6d, 0x20, 0x7a, 0xd1, 0x81, 0xde, 0xa1, 0x2d, 0xe6, 0x98, 0xc8, 0xa7, 0xe8, 0x7b,
	0x68, 0xb3, 0x62, 0x51, 0x2a, 0xd9, 0x0f, 0xb6, 0x83, 0x9d, 0xde, 0x68, 0x3f, 0x5e, 0xab, 0xb6,
	0xb8, 0xa6, 0x11, 0x3f, 0x30, 0x02, 0xd8, 0x09, 0x21, 0x0a, 0x37, 0xe9, 0xf3, 0x05, 0x4d, 0x14,
	0x4d, 0xa7, 0xbc, 0x54, 0x46, 0xbc, 0x69, 0xc4, 0xbf, 0xb8, 0x82, 0xf8, 0xb7, 0x56, 0x01, 0xdf,
	0xf0, 0x9a, 0x0e, 0x40, 0x27, 0xd0, 0x51, 0x6c, 0x4e, 0xb9, 0x96, 0x6f, 0x19, 0xf9, 0x83, 0x2b,
	0xc8, 0x1f, 0x3b, 0x09, 0x5c, 0x89, 0x0d, 0x7e, 0x6b, 0x40, 0xdb, 0x5e, 0x09, 0xbd, 0x07, 0x5d,
	0x22, 0xb2, 0x72, 0x4e, 0x0b, 0x63, 0x50, 0x63, 0xa7, 0x8b, 0xcf, 0x01, 0x74, 0x08, 0xad, 0x27,
	0x2c, 0xa7, 0xb2, 0x1f, 0x6e, 0x37, 0x76, 0x7a, 0xa3, 0xdd, 0x35, 0x8f, 0x9f, 0xb0, 0x8c, 0x4a,
	0x85, 0x2d, 0x17, 0x7d, 0x03, 0x1b, 0xac, 0xc8, 0x59, 0x41, 0xa7, 0xb3, 0x9c, 0xcf, 0xb4, 0x53,
	0x5a, 0xeb, 0xee, 0x9a, 0x5a, 0xe3, 0x9c, 0xcf, 0x70, 0xcf, 0x0a, 0xe8, 0x6f, 0x89, 0x7e, 0x0e,
	0xe0, 0x16, 0x2d, 0x96, 0x4c, 0xf0, 0x42, 0x57, 0x39, 0x5d, 0x12, 0xc1, 0xc8, 0x4c, 0x57, 0xd9,
	0x30, 0xca, 0x0f, 0xaf, 0xdc, 0xe0, 0xf8, 0xe8, 0x5c, 0xf6, 0x91, 0x53, 0xc5, 0x5b, 0xf4, 0x22,
	0x28, 0xd1, 0x5d, 0xd8, 0xd4, 0x7a, 0xac, 0xc8, 0xa6, 0x29, 0x13, 0x34, 0x51, 0x5c, 0x9c, 0x99,
	0x26, 0x75, 0xf1, 0x4d, 0x17, 0x98, 0x78, 0x7c, 0xf0, 0x25, 0xbc, 0xfd, 0x37, 0xca, 0x08, 0x41,
	0xb3, 0x20, 0x73, 0x6a, 0xde, 0x65, 0x17, 0x9b, 0x6f, 0xb4, 0x05, 0xad, 0x25, 0xc9, 0x4b, 0xda,
	0x0f, 0x0d, 0x68, 0x7f, 0x0c, 0x5e, 0x04, 0x70, 0xcd, 0xbf, 0x8a, 0x2d, 0xdf, 0x13, 0xdb, 0x2d,
	0x67, 0xf2, 0x36, 0xf4, 0x7c, 0x1d, 0xcc, 0xf5, 0xab, 0x8b, 0xeb, 0x10, 0xda, 0x05, 0x24, 0x55,
	0xca, 0x4b, 0x35, 0x4d, 0xa9, 0x54, 0xac, 0x30, 0xc3, 0xd4, 0x6f, 0x98, 0x63, 0x36, 0x6d, 0x64,
	0x72, 0x1e, 0x70, 0xe9, 0x54, 0x88, 0x95, 0xf4, 0x66, 0x95, 0x4e, 0x85, 0xa8, 0xa5, 0x0f, 0x7e,
	0x0d, 0xa0, 0xe3, 0x5f, 0x1a, 0xfa, 0x0c, 0xba, 0xf4, 0x39, 0x4d, 0x4a, 0x43, 0xb1, 0x53, 0x77,
	0xdb, 0x37, 0xc5, 0xcf, 0x73, 0x3c, 0x71, 0xf3, 0x8c, 0xcf, 0x73, 0xd1, 0x2e, 0x34, 0x59, 0x9a,
	0xdb, 0xcb, 0xbf, 0x96, 0x63, 0xd2, 0xd0, 0xa7, 0xd0, 0x91, 0xa7, 0xa5, 0x4a, 0xf9, 0x33, 0x7b,
	0x91, 0xd7, 0x52, 0xaa, 0xd4, 0x68, 0x09, 0x6f, 0xb9, 0xde, 0x7b, 0x4f, 0xdf, 0xd5, 0x05, 0x33,
	0x35, 0x4d, 0x78, 0x6a, 0xdb, 0xd1, 0xc2, 0x1d, 0x0d, 0x1c, 0xf2, 0x94, 0xa2, 0xfb, 0x70, 0xcd,
	0x0f, 0xb9, 0xad, 0xeb, 0x5f, 0x8e, 0x81, 0x67, 0x47, 0x3f, 0xc1, 0x0d, 0x7f, 0xee, 0x92, 0x8a,
	0x53, 0x4a, 0x52, 0x7d, 0x03, 0xbf, 0xd7, 0x2e, 0x37, 0xaa, 0x4a, 0xd5, 0x34, 0xee, 0x24, 0x2e,
	0xf7, 0xaa, 0x4a, 0x8d, 0xfe, 0x0c, 0xe1, 0xba, 0xab, 0x00, 0x53, 0x59, 0xe6, 0x0a, 0x7d, 0x0c,
	0x6d, 0xbb, 0x3c, 0xdd, 0xe9, 0xc8, 0xcb, 0x88, 0x45, 0x12, 0xff, 0x60, 0x22, 0xd8, 0x65, 0xac,
	0x9a, 0x14, 0xfe, 0xb3, 0x49, 0x8d, 0xff, 0x62, 0x12, 0xda, 0xaf, 0x39, 0xd2, 0xbc, 0xe4, 0x6a,
	0xe3, 0xb0, 0x1f, 0xd4, 0x5c, 0xd9, 0xaf, 0xb9, 0xd2, 0x5a, 0x8b, 0xea, 0xd3, 0xd1, 0x27, 0xd0,
	0x99, 0x53, 0x45, 0x52, 0xa2, 0x48, 0xbf, 0x6d, 0xb6, 0xc8, 0xd6, 0x05, 0xea, 0x57, 0xc5, 0x19,
	0xae, 0xb2, 0xa2, 0x97, 0x01, 0x6c, 0x7c, 0xcd, 0x72, 0xfa, 0xd0, 0x01, 0x7a, 0x9a, 0x17, 0x44,
	0x9d, 0xfa, 0x69, 0xd6, 0xdf, 0xe8, 0x08, 0xda, 0xa9, 0xb9, 0xdf, 0xd5, 0x5e, 0x8e, 0x23, 0xa3,
	0x01, 0x74, 0x12, 0x5e, 0x28, 0xb3, 0xa3, 0xb5, 0xbb, 0x1b, 0xb8, 0xfa, 0x8d, 0x3e, 0x84, 0xeb,
	0x4c, 0x4e, 0xed, 0x08, 0xe9, 0xad, 0x62, 0x4c, 0xeb, 0xe0, 0x0d, 0x26, 0x8f, 0x2a, 0x2c, 0x2a,
	0x60, 0xb3, 0xda, 0x46, 0x6f, 0xa0, 0xe0, 0xe8, 0x00, 0xda, 0x16, 0xd1, 0x87, 0x9c, 0x12, 0x59,
	0x1d, 0xa2, 0xbf, 0xd1, 0xfb, 0x00, 0x92, 0xfd, 0x48, 0xa7, 0xb3, 0x33, 0x45, 0xed, 0x4c, 0x35,
	0x70, 0x57, 0x23, 0x63, 0x0d, 0x44, 0x0c, 0x9a, 0x7a, 0xd1, 0xd7, 0x6a, 0x09, 0xfe, 0x2f, 0xf3,
	0xc2, 0x55, 0xf3, 0xa2, 0xdf, 0x03, 0xe8, 0x56, 0xc6, 0xa0, 0x07, 0xf5, 0xcd, 0xda, 0x1b, 0xed,
	0xad, 0x79, 0x5e, 0xfd, 0x15, 0xf8, 0x75, 0xfc, 0xf8, 0xe2, 0x3a, 0xee, 0x8d, 0x3e, 0x5f, 0xfb,
	0x02, 0xaf, 0xb4, 0x6a, 0x65, 0x91, 0x8f, 0xff, 0x08, 0xe0, 0xa3, 0x84, 0xcf, 0xd7, 0x13, 0x1b,
	0xdf, 0xc2, 0x06, 0x3e, 0xb1, 0xb0, 0x9b, 0x7e, 0xf9, 0x5d, 0xf0, 0x18, 0x3b, 0x7e, 0xc6, 0x73,
	0x52, 0x64, 0x31, 0x17, 0xd9, 0x30, 0xa3, 0x85, 0x79, 0xed, 0x43, 0x1b, 0x22, 0x0b, 0x26, 0x2f,
	0xf9, 0xdb, 0x77, 0xb0, 0x82, 0xfe, 0x12, 0x86, 0xf8, 0xe4, 0x65, 0x78, 0xe7, 0xbe, 0x55, 0x9e,
	0xd0, 0xe5, 0xb1, 0xa9, 0x6c, 0xa5, 0x84, 0xf8, 0xd1, 0xbd, 0x63, 0x4d, 0x9d, 0xb5, 0xcd, 0x59,
	0x7b, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x79, 0x22, 0xd3, 0x61, 0x0a, 0x00, 0x00,
}
