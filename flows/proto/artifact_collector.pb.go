// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifact_collector.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
	proto2 "www.velocidex.com/golang/velociraptor/crypto/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type ArtifactCollectorContext_State int32

const (
	ArtifactCollectorContext_UNSET      ArtifactCollectorContext_State = 0
	ArtifactCollectorContext_RUNNING    ArtifactCollectorContext_State = 1
	ArtifactCollectorContext_TERMINATED ArtifactCollectorContext_State = 2
	ArtifactCollectorContext_ERROR      ArtifactCollectorContext_State = 3
	ArtifactCollectorContext_ARCHIVED   ArtifactCollectorContext_State = 4
)

var ArtifactCollectorContext_State_name = map[int32]string{
	0: "UNSET",
	1: "RUNNING",
	2: "TERMINATED",
	3: "ERROR",
	4: "ARCHIVED",
}

var ArtifactCollectorContext_State_value = map[string]int32{
	"UNSET":      0,
	"RUNNING":    1,
	"TERMINATED": 2,
	"ERROR":      3,
	"ARCHIVED":   4,
}

func (x ArtifactCollectorContext_State) String() string {
	return proto.EnumName(ArtifactCollectorContext_State_name, int32(x))
}

func (ArtifactCollectorContext_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{4, 0}
}

type FlowContext_State int32

const (
	FlowContext_UNSET      FlowContext_State = 0
	FlowContext_RUNNING    FlowContext_State = 1
	FlowContext_TERMINATED FlowContext_State = 2
	FlowContext_ERROR      FlowContext_State = 3
	FlowContext_ARCHIVED   FlowContext_State = 4
)

var FlowContext_State_name = map[int32]string{
	0: "UNSET",
	1: "RUNNING",
	2: "TERMINATED",
	3: "ERROR",
	4: "ARCHIVED",
}

var FlowContext_State_value = map[string]int32{
	"UNSET":      0,
	"RUNNING":    1,
	"TERMINATED": 2,
	"ERROR":      3,
	"ARCHIVED":   4,
}

func (x FlowContext_State) String() string {
	return proto.EnumName(FlowContext_State_name, int32(x))
}

func (FlowContext_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{7, 0}
}

type ArtifactParameters struct {
	Env                  []*proto1.VQLEnv `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArtifactParameters) Reset()         { *m = ArtifactParameters{} }
func (m *ArtifactParameters) String() string { return proto.CompactTextString(m) }
func (*ArtifactParameters) ProtoMessage()    {}
func (*ArtifactParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{0}
}

func (m *ArtifactParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactParameters.Unmarshal(m, b)
}
func (m *ArtifactParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactParameters.Marshal(b, m, deterministic)
}
func (m *ArtifactParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactParameters.Merge(m, src)
}
func (m *ArtifactParameters) XXX_Size() int {
	return xxx_messageInfo_ArtifactParameters.Size(m)
}
func (m *ArtifactParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactParameters proto.InternalMessageInfo

func (m *ArtifactParameters) GetEnv() []*proto1.VQLEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

type ArtifactCollectorArgs struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// If set we send an urgent request to the client.
	Urgent               bool                `protobuf:"varint,21,opt,name=urgent,proto3" json:"urgent,omitempty"`
	Artifacts            []string            `protobuf:"bytes,2,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	Parameters           *ArtifactParameters `protobuf:"bytes,5,opt,name=parameters,proto3" json:"parameters,omitempty"`
	OpsPerSecond         float32             `protobuf:"fixed32,6,opt,name=ops_per_second,json=opsPerSecond,proto3" json:"ops_per_second,omitempty"`
	Timeout              uint64              `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	AllowCustomOverrides bool                `protobuf:"varint,8,opt,name=allow_custom_overrides,json=allowCustomOverrides,proto3" json:"allow_custom_overrides,omitempty"`
	// A place to cache the compiled request. If this is provided we
	// do not compile the artifacts at all, we just use it as is.
	CompiledCollectorArgs *proto1.VQLCollectorArgs `protobuf:"bytes,20,opt,name=compiled_collector_args,json=compiledCollectorArgs,proto3" json:"compiled_collector_args,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *ArtifactCollectorArgs) Reset()         { *m = ArtifactCollectorArgs{} }
func (m *ArtifactCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorArgs) ProtoMessage()    {}
func (*ArtifactCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{1}
}

func (m *ArtifactCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorArgs.Unmarshal(m, b)
}
func (m *ArtifactCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorArgs.Marshal(b, m, deterministic)
}
func (m *ArtifactCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorArgs.Merge(m, src)
}
func (m *ArtifactCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorArgs.Size(m)
}
func (m *ArtifactCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorArgs proto.InternalMessageInfo

func (m *ArtifactCollectorArgs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ArtifactCollectorArgs) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ArtifactCollectorArgs) GetUrgent() bool {
	if m != nil {
		return m.Urgent
	}
	return false
}

func (m *ArtifactCollectorArgs) GetArtifacts() []string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetParameters() *ArtifactParameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetOpsPerSecond() float32 {
	if m != nil {
		return m.OpsPerSecond
	}
	return 0
}

func (m *ArtifactCollectorArgs) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ArtifactCollectorArgs) GetAllowCustomOverrides() bool {
	if m != nil {
		return m.AllowCustomOverrides
	}
	return false
}

func (m *ArtifactCollectorArgs) GetCompiledCollectorArgs() *proto1.VQLCollectorArgs {
	if m != nil {
		return m.CompiledCollectorArgs
	}
	return nil
}

type ArtifactCollectorResponse struct {
	FlowId               string                 `protobuf:"bytes,1,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Request              *ArtifactCollectorArgs `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ArtifactCollectorResponse) Reset()         { *m = ArtifactCollectorResponse{} }
func (m *ArtifactCollectorResponse) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorResponse) ProtoMessage()    {}
func (*ArtifactCollectorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{2}
}

func (m *ArtifactCollectorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorResponse.Unmarshal(m, b)
}
func (m *ArtifactCollectorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorResponse.Marshal(b, m, deterministic)
}
func (m *ArtifactCollectorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorResponse.Merge(m, src)
}
func (m *ArtifactCollectorResponse) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorResponse.Size(m)
}
func (m *ArtifactCollectorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorResponse proto.InternalMessageInfo

func (m *ArtifactCollectorResponse) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *ArtifactCollectorResponse) GetRequest() *ArtifactCollectorArgs {
	if m != nil {
		return m.Request
	}
	return nil
}

type ArtifactUploadedFileInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VfsPath              string   `protobuf:"bytes,2,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactUploadedFileInfo) Reset()         { *m = ArtifactUploadedFileInfo{} }
func (m *ArtifactUploadedFileInfo) String() string { return proto.CompactTextString(m) }
func (*ArtifactUploadedFileInfo) ProtoMessage()    {}
func (*ArtifactUploadedFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{3}
}

func (m *ArtifactUploadedFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactUploadedFileInfo.Unmarshal(m, b)
}
func (m *ArtifactUploadedFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactUploadedFileInfo.Marshal(b, m, deterministic)
}
func (m *ArtifactUploadedFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactUploadedFileInfo.Merge(m, src)
}
func (m *ArtifactUploadedFileInfo) XXX_Size() int {
	return xxx_messageInfo_ArtifactUploadedFileInfo.Size(m)
}
func (m *ArtifactUploadedFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactUploadedFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactUploadedFileInfo proto.InternalMessageInfo

func (m *ArtifactUploadedFileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactUploadedFileInfo) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *ArtifactUploadedFileInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// This context is serialized into the data store.
type ArtifactCollectorContext struct {
	// Where we are stored in the data store.
	Urn       string                 `protobuf:"bytes,5,opt,name=urn,proto3" json:"urn,omitempty"`
	SessionId string                 `protobuf:"bytes,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Request   *ArtifactCollectorArgs `protobuf:"bytes,11,opt,name=request,proto3" json:"request,omitempty"`
	// If an error occurs this is the backtrace.
	Backtrace     string `protobuf:"bytes,1,opt,name=backtrace,proto3" json:"backtrace,omitempty"`
	CreateTime    uint64 `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	ActiveTime    uint64 `protobuf:"varint,17,opt,name=active_time,json=activeTime,proto3" json:"active_time,omitempty"`
	KillTimestamp uint64 `protobuf:"varint,4,opt,name=kill_timestamp,json=killTimestamp,proto3" json:"kill_timestamp,omitempty"`
	// Uploads are now permanently stored in a csv file. This field is
	// never serialized - it is just a place holder during processing.
	UploadedFiles []*ArtifactUploadedFileInfo `protobuf:"bytes,24,rep,name=uploaded_files,json=uploadedFiles,proto3" json:"uploaded_files,omitempty"`
	// A total count of files uploaded by this query.
	TotalUploadedFiles         uint64 `protobuf:"varint,23,opt,name=total_uploaded_files,json=totalUploadedFiles,proto3" json:"total_uploaded_files,omitempty"`
	TotalExpectedUploadedBytes uint64 `protobuf:"varint,25,opt,name=total_expected_uploaded_bytes,json=totalExpectedUploadedBytes,proto3" json:"total_expected_uploaded_bytes,omitempty"`
	TotalUploadedBytes         uint64 `protobuf:"varint,26,opt,name=total_uploaded_bytes,json=totalUploadedBytes,proto3" json:"total_uploaded_bytes,omitempty"`
	// Logs are stored in their own CSV file. This is just a
	// placeholder during processing.
	Logs         []*proto2.LogMessage           `protobuf:"bytes,20,rep,name=logs,proto3" json:"logs,omitempty"`
	State        ArtifactCollectorContext_State `protobuf:"varint,14,opt,name=state,proto3,enum=proto.ArtifactCollectorContext_State" json:"state,omitempty"`
	Status       string                         `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	UserNotified bool                           `protobuf:"varint,16,opt,name=user_notified,json=userNotified,proto3" json:"user_notified,omitempty"`
	// Some of the collected artifacts may not have results.
	ArtifactsWithResults []string `protobuf:"bytes,22,rep,name=artifacts_with_results,json=artifactsWithResults,proto3" json:"artifacts_with_results,omitempty"`
	Dirty                bool     `protobuf:"varint,2,opt,name=dirty,proto3" json:"dirty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactCollectorContext) Reset()         { *m = ArtifactCollectorContext{} }
func (m *ArtifactCollectorContext) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorContext) ProtoMessage()    {}
func (*ArtifactCollectorContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{4}
}

func (m *ArtifactCollectorContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorContext.Unmarshal(m, b)
}
func (m *ArtifactCollectorContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorContext.Marshal(b, m, deterministic)
}
func (m *ArtifactCollectorContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorContext.Merge(m, src)
}
func (m *ArtifactCollectorContext) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorContext.Size(m)
}
func (m *ArtifactCollectorContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorContext.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorContext proto.InternalMessageInfo

func (m *ArtifactCollectorContext) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *ArtifactCollectorContext) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *ArtifactCollectorContext) GetRequest() *ArtifactCollectorArgs {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ArtifactCollectorContext) GetBacktrace() string {
	if m != nil {
		return m.Backtrace
	}
	return ""
}

func (m *ArtifactCollectorContext) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *ArtifactCollectorContext) GetActiveTime() uint64 {
	if m != nil {
		return m.ActiveTime
	}
	return 0
}

func (m *ArtifactCollectorContext) GetKillTimestamp() uint64 {
	if m != nil {
		return m.KillTimestamp
	}
	return 0
}

func (m *ArtifactCollectorContext) GetUploadedFiles() []*ArtifactUploadedFileInfo {
	if m != nil {
		return m.UploadedFiles
	}
	return nil
}

func (m *ArtifactCollectorContext) GetTotalUploadedFiles() uint64 {
	if m != nil {
		return m.TotalUploadedFiles
	}
	return 0
}

func (m *ArtifactCollectorContext) GetTotalExpectedUploadedBytes() uint64 {
	if m != nil {
		return m.TotalExpectedUploadedBytes
	}
	return 0
}

func (m *ArtifactCollectorContext) GetTotalUploadedBytes() uint64 {
	if m != nil {
		return m.TotalUploadedBytes
	}
	return 0
}

func (m *ArtifactCollectorContext) GetLogs() []*proto2.LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *ArtifactCollectorContext) GetState() ArtifactCollectorContext_State {
	if m != nil {
		return m.State
	}
	return ArtifactCollectorContext_UNSET
}

func (m *ArtifactCollectorContext) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ArtifactCollectorContext) GetUserNotified() bool {
	if m != nil {
		return m.UserNotified
	}
	return false
}

func (m *ArtifactCollectorContext) GetArtifactsWithResults() []string {
	if m != nil {
		return m.ArtifactsWithResults
	}
	return nil
}

func (m *ArtifactCollectorContext) GetDirty() bool {
	if m != nil {
		return m.Dirty
	}
	return false
}

// This is stored in the ArtifactCollector state.
type ClientEventTable struct {
	Version              uint64                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Artifacts            *ArtifactCollectorArgs `protobuf:"bytes,2,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ClientEventTable) Reset()         { *m = ClientEventTable{} }
func (m *ClientEventTable) String() string { return proto.CompactTextString(m) }
func (*ClientEventTable) ProtoMessage()    {}
func (*ClientEventTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{5}
}

func (m *ClientEventTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventTable.Unmarshal(m, b)
}
func (m *ClientEventTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventTable.Marshal(b, m, deterministic)
}
func (m *ClientEventTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventTable.Merge(m, src)
}
func (m *ClientEventTable) XXX_Size() int {
	return xxx_messageInfo_ClientEventTable.Size(m)
}
func (m *ClientEventTable) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventTable.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventTable proto.InternalMessageInfo

func (m *ClientEventTable) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ClientEventTable) GetArtifacts() *ArtifactCollectorArgs {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type UploadedFileInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VfsPath              string   `protobuf:"bytes,2,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadedFileInfo) Reset()         { *m = UploadedFileInfo{} }
func (m *UploadedFileInfo) String() string { return proto.CompactTextString(m) }
func (*UploadedFileInfo) ProtoMessage()    {}
func (*UploadedFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{6}
}

func (m *UploadedFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadedFileInfo.Unmarshal(m, b)
}
func (m *UploadedFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadedFileInfo.Marshal(b, m, deterministic)
}
func (m *UploadedFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadedFileInfo.Merge(m, src)
}
func (m *UploadedFileInfo) XXX_Size() int {
	return xxx_messageInfo_UploadedFileInfo.Size(m)
}
func (m *UploadedFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadedFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UploadedFileInfo proto.InternalMessageInfo

func (m *UploadedFileInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadedFileInfo) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *UploadedFileInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// The flow context.
// Next field: 19
type FlowContext struct {
	Backtrace        string `protobuf:"bytes,1,opt,name=backtrace,proto3" json:"backtrace,omitempty"`
	CreateTime       uint64 `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	KillTimestamp    uint64 `protobuf:"varint,6,opt,name=kill_timestamp,json=killTimestamp,proto3" json:"kill_timestamp,omitempty"`
	NetworkBytesSent uint64 `protobuf:"varint,7,opt,name=network_bytes_sent,json=networkBytesSent,proto3" json:"network_bytes_sent,omitempty"`
	// Uploads are now permanently stored in a csv file. This field is
	// never serialized - it is just a place holder during processing.
	UploadedFiles              []*UploadedFileInfo `protobuf:"bytes,24,rep,name=uploaded_files,json=uploadedFiles,proto3" json:"uploaded_files,omitempty"`
	TotalUploadedFiles         uint64              `protobuf:"varint,23,opt,name=total_uploaded_files,json=totalUploadedFiles,proto3" json:"total_uploaded_files,omitempty"`
	TotalExpectedUploadedBytes uint64              `protobuf:"varint,25,opt,name=total_expected_uploaded_bytes,json=totalExpectedUploadedBytes,proto3" json:"total_expected_uploaded_bytes,omitempty"`
	TotalUploadedBytes         uint64              `protobuf:"varint,26,opt,name=total_uploaded_bytes,json=totalUploadedBytes,proto3" json:"total_uploaded_bytes,omitempty"`
	// Logs are stored in their own CSV file. This is just a
	// placeholder during processing.
	Logs                 []*proto2.LogMessage `protobuf:"bytes,20,rep,name=logs,proto3" json:"logs,omitempty"`
	SessionId            string               `protobuf:"bytes,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	State                FlowContext_State    `protobuf:"varint,14,opt,name=state,proto3,enum=proto.FlowContext_State" json:"state,omitempty"`
	Status               string               `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	UserNotified         bool                 `protobuf:"varint,16,opt,name=user_notified,json=userNotified,proto3" json:"user_notified,omitempty"`
	ActiveTime           uint64               `protobuf:"varint,17,opt,name=active_time,json=activeTime,proto3" json:"active_time,omitempty"`
	Artifacts            []string             `protobuf:"bytes,21,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	ArtifactsWithResults []string             `protobuf:"bytes,22,rep,name=artifacts_with_results,json=artifactsWithResults,proto3" json:"artifacts_with_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FlowContext) Reset()         { *m = FlowContext{} }
func (m *FlowContext) String() string { return proto.CompactTextString(m) }
func (*FlowContext) ProtoMessage()    {}
func (*FlowContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecbffc830d4bd8f6, []int{7}
}

func (m *FlowContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowContext.Unmarshal(m, b)
}
func (m *FlowContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowContext.Marshal(b, m, deterministic)
}
func (m *FlowContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowContext.Merge(m, src)
}
func (m *FlowContext) XXX_Size() int {
	return xxx_messageInfo_FlowContext.Size(m)
}
func (m *FlowContext) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowContext.DiscardUnknown(m)
}

var xxx_messageInfo_FlowContext proto.InternalMessageInfo

func (m *FlowContext) GetBacktrace() string {
	if m != nil {
		return m.Backtrace
	}
	return ""
}

func (m *FlowContext) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FlowContext) GetKillTimestamp() uint64 {
	if m != nil {
		return m.KillTimestamp
	}
	return 0
}

func (m *FlowContext) GetNetworkBytesSent() uint64 {
	if m != nil {
		return m.NetworkBytesSent
	}
	return 0
}

func (m *FlowContext) GetUploadedFiles() []*UploadedFileInfo {
	if m != nil {
		return m.UploadedFiles
	}
	return nil
}

func (m *FlowContext) GetTotalUploadedFiles() uint64 {
	if m != nil {
		return m.TotalUploadedFiles
	}
	return 0
}

func (m *FlowContext) GetTotalExpectedUploadedBytes() uint64 {
	if m != nil {
		return m.TotalExpectedUploadedBytes
	}
	return 0
}

func (m *FlowContext) GetTotalUploadedBytes() uint64 {
	if m != nil {
		return m.TotalUploadedBytes
	}
	return 0
}

func (m *FlowContext) GetLogs() []*proto2.LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *FlowContext) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *FlowContext) GetState() FlowContext_State {
	if m != nil {
		return m.State
	}
	return FlowContext_UNSET
}

func (m *FlowContext) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FlowContext) GetUserNotified() bool {
	if m != nil {
		return m.UserNotified
	}
	return false
}

func (m *FlowContext) GetActiveTime() uint64 {
	if m != nil {
		return m.ActiveTime
	}
	return 0
}

func (m *FlowContext) GetArtifacts() []string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *FlowContext) GetArtifactsWithResults() []string {
	if m != nil {
		return m.ArtifactsWithResults
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.ArtifactCollectorContext_State", ArtifactCollectorContext_State_name, ArtifactCollectorContext_State_value)
	proto.RegisterEnum("proto.FlowContext_State", FlowContext_State_name, FlowContext_State_value)
	proto.RegisterType((*ArtifactParameters)(nil), "proto.ArtifactParameters")
	proto.RegisterType((*ArtifactCollectorArgs)(nil), "proto.ArtifactCollectorArgs")
	proto.RegisterType((*ArtifactCollectorResponse)(nil), "proto.ArtifactCollectorResponse")
	proto.RegisterType((*ArtifactUploadedFileInfo)(nil), "proto.ArtifactUploadedFileInfo")
	proto.RegisterType((*ArtifactCollectorContext)(nil), "proto.ArtifactCollectorContext")
	proto.RegisterType((*ClientEventTable)(nil), "proto.ClientEventTable")
	proto.RegisterType((*UploadedFileInfo)(nil), "proto.UploadedFileInfo")
	proto.RegisterType((*FlowContext)(nil), "proto.FlowContext")
}

func init() {
	proto.RegisterFile("artifact_collector.proto", fileDescriptor_ecbffc830d4bd8f6)
}

var fileDescriptor_ecbffc830d4bd8f6 = []byte{
	// 1523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xed, 0x6e, 0x1b, 0x45,
	0x17, 0x7e, 0x1d, 0x27, 0x76, 0x3c, 0x69, 0xf2, 0xba, 0xf3, 0x26, 0xcd, 0x26, 0x6f, 0xab, 0x0e,
	0x6e, 0x0b, 0xa6, 0x2a, 0xeb, 0x34, 0xfd, 0x40, 0x2a, 0x02, 0x14, 0x27, 0x0e, 0xb5, 0x48, 0x93,
	0x74, 0xe3, 0xb4, 0x02, 0x81, 0xac, 0xf1, 0xee, 0x59, 0x7b, 0x9a, 0xf5, 0xce, 0x76, 0x66, 0xd6,
	0x6e, 0x7a, 0x0f, 0x5c, 0x03, 0x37, 0xc3, 0x7f, 0x7e, 0x71, 0x03, 0x70, 0x1b, 0x20, 0xa1, 0x99,
	0xdd, 0x8d, 0xed, 0x7c, 0x15, 0x24, 0x90, 0x9a, 0x3f, 0xde, 0x3d, 0x5f, 0xcf, 0x9c, 0x33, 0xe7,
	0x3c, 0x67, 0x83, 0x2c, 0x2a, 0x14, 0xf3, 0xa9, 0xab, 0xda, 0x2e, 0x0f, 0x02, 0x70, 0x15, 0x17,
	0x76, 0x24, 0xb8, 0xe2, 0x78, 0xc6, 0xfc, 0xac, 0x2e, 0x9a, 0x9f, 0x9a, 0x84, 0x3e, 0x0d, 0x15,
	0x73, 0x13, 0xe5, 0xea, 0x32, 0x75, 0x15, 0xe3, 0xa1, 0xac, 0x25, 0xda, 0xc1, 0xeb, 0x20, 0x53,
	0xb8, 0xe2, 0x38, 0x52, 0x3c, 0x95, 0xbf, 0xe2, 0x1d, 0x99, 0x28, 0x2a, 0x6f, 0x11, 0xde, 0x48,
	0xa1, 0xf6, 0xa9, 0xa0, 0x7d, 0x50, 0x20, 0x24, 0xf6, 0x50, 0x1e, 0xc2, 0x81, 0x95, 0x27, 0xf9,
	0xea, 0xdc, 0xfa, 0x7c, 0x62, 0x6a, 0xbf, 0x78, 0xbe, 0xd3, 0x08, 0x07, 0xf5, 0xcd, 0x5f, 0x7f,
	0xff, 0xed, 0xa7, 0xdc, 0xe7, 0xf8, 0x41, 0x23, 0x1c, 0x30, 0xc1, 0xc3, 0x3e, 0x84, 0x8a, 0x0c,
	0xa8, 0x60, 0xb4, 0x13, 0x80, 0x24, 0x8a, 0x93, 0x0e, 0x90, 0x48, 0xf0, 0x01, 0xf3, 0xc0, 0x23,
	0x3e, 0x17, 0x44, 0xf5, 0x80, 0xbc, 0x8e, 0x41, 0x1c, 0xdb, 0x95, 0x82, 0x01, 0x91, 0x8e, 0x0e,
	0x5f, 0xf9, 0x79, 0x06, 0x2d, 0x65, 0xe0, 0x9b, 0x59, 0x9a, 0x1b, 0xa2, 0x2b, 0xb1, 0x85, 0x8a,
	0xae, 0x00, 0xaa, 0xb8, 0xb0, 0x72, 0x24, 0x57, 0x2d, 0x39, 0xd9, 0x2b, 0xfe, 0x3f, 0x2a, 0xb9,
	0x01, 0x83, 0x50, 0xb5, 0x99, 0x67, 0xe5, 0x8d, 0x6e, 0x36, 0x11, 0x34, 0x3d, 0x7c, 0x0d, 0x15,
	0x62, 0xd1, 0x85, 0x50, 0x59, 0x4b, 0x24, 0x57, 0x9d, 0x75, 0xd2, 0x37, 0xbc, 0x8d, 0x4a, 0x59,
	0x3d, 0xa5, 0x35, 0x45, 0xf2, 0xd5, 0x52, 0xbd, 0x6a, 0xb2, 0xa8, 0x60, 0xab, 0xd5, 0x03, 0x72,
	0xa2, 0xd4, 0xa7, 0x0f, 0x68, 0x1c, 0xba, 0x3d, 0xbb, 0x52, 0xd8, 0x31, 0x0f, 0xce, 0xc8, 0x15,
	0x1f, 0x21, 0x14, 0x9d, 0x14, 0xc9, 0x9a, 0x21, 0xb9, 0xea, 0xdc, 0xfa, 0x4a, 0x5a, 0x9d, 0xb3,
	0x55, 0xac, 0xaf, 0x19, 0x8c, 0xbb, 0xf8, 0xfa, 0x48, 0xa6, 0x01, 0xd4, 0x38, 0xa2, 0x5d, 0x41,
	0x23, 0xad, 0x33, 0x16, 0x1e, 0xbf, 0x42, 0x0b, 0x3c, 0x92, 0xed, 0x08, 0x44, 0x5b, 0x82, 0xcb,
	0x43, 0xcf, 0x2a, 0x90, 0x5c, 0x75, 0xaa, 0xbe, 0x65, 0xa2, 0x7e, 0x81, 0x6f, 0xed, 0x45, 0x20,
	0xa8, 0xb9, 0x6e, 0x12, 0x81, 0x20, 0x89, 0x11, 0xa9, 0xb6, 0x7a, 0x82, 0x2b, 0x15, 0xb0, 0xb0,
	0xfb, 0xb1, 0x5d, 0x59, 0xd8, 0x8b, 0x24, 0xd9, 0x07, 0x41, 0x0e, 0x8c, 0x76, 0xbd, 0x78, 0x7f,
	0xcd, 0xfc, 0x39, 0x57, 0x78, 0x24, 0xf7, 0x41, 0x24, 0x62, 0x0c, 0xa8, 0xa8, 0x58, 0x1f, 0x78,
	0xac, 0xac, 0x22, 0xc9, 0x55, 0xa7, 0xeb, 0x5f, 0x1b, 0x90, 0x06, 0x7e, 0xb4, 0x1b, 0xf7, 0x3b,
	0x20, 0x08, 0xf7, 0xd3, 0xf8, 0x26, 0x03, 0x11, 0x87, 0xa4, 0x03, 0x3e, 0x17, 0x40, 0x5c, 0x1a,
	0xba, 0x10, 0x68, 0xb4, 0xf1, 0x6b, 0x2e, 0xb6, 0x92, 0x68, 0xeb, 0xf9, 0xc7, 0x6b, 0x6b, 0x4e,
	0x16, 0x1b, 0xff, 0x90, 0x43, 0xd7, 0x68, 0x10, 0xf0, 0x61, 0xdb, 0x8d, 0xa5, 0xe2, 0xfd, 0x36,
	0x1f, 0x80, 0x10, 0xcc, 0x03, 0x69, 0xcd, 0xea, 0x0b, 0xab, 0xbf, 0x34, 0xb0, 0xcf, 0xf1, 0x5e,
	0xd3, 0x27, 0x4a, 0xc4, 0x40, 0x86, 0x40, 0x86, 0x2c, 0x08, 0x48, 0x2c, 0x81, 0x50, 0x92, 0x78,
	0x9d, 0x14, 0x8f, 0x30, 0x9f, 0x44, 0x02, 0xa4, 0x6e, 0x40, 0x16, 0x4a, 0x05, 0xd4, 0xd3, 0x07,
	0xd5, 0xe7, 0x08, 0x69, 0x1f, 0xbc, 0x13, 0x43, 0xdb, 0x59, 0x34, 0xb0, 0x9b, 0xc6, 0x7f, 0x2f,
	0x03, 0xc5, 0x7b, 0x68, 0xd9, 0xe5, 0xfd, 0x88, 0x05, 0xe0, 0x8d, 0xe6, 0xac, 0x4d, 0x45, 0x57,
	0x5a, 0x8b, 0xe6, 0x72, 0x97, 0x47, 0xad, 0x3f, 0xd1, 0xa0, 0xce, 0x52, 0xe6, 0x37, 0x21, 0xae,
	0x04, 0x68, 0xe5, 0x4c, 0x43, 0x3b, 0x20, 0x23, 0x1e, 0x4a, 0xc0, 0xcb, 0xa8, 0xe8, 0xeb, 0xdc,
	0x99, 0x97, 0x36, 0x75, 0x41, 0xbf, 0x36, 0x3d, 0xfc, 0x18, 0x15, 0x05, 0xbc, 0x8e, 0x41, 0x2a,
	0x6b, 0xca, 0xc0, 0x5e, 0x3f, 0xd5, 0x53, 0x93, 0xd8, 0x99, 0x71, 0xe5, 0x7b, 0x64, 0x65, 0x16,
	0x87, 0x51, 0xc0, 0xa9, 0x07, 0xde, 0x36, 0x0b, 0xa0, 0x19, 0xfa, 0x1c, 0x63, 0x34, 0xad, 0x6b,
	0x90, 0x22, 0x99, 0x67, 0xbc, 0x82, 0x66, 0x07, 0xbe, 0x6c, 0x47, 0x54, 0xf5, 0x0c, 0x50, 0xc9,
	0x29, 0x0e, 0x7c, 0xb9, 0x4f, 0x55, 0x4f, 0x9b, 0x4b, 0xf6, 0x16, 0xcc, 0x44, 0x4d, 0x3b, 0xe6,
	0xb9, 0xf2, 0x4b, 0x71, 0x14, 0xff, 0xe4, 0x04, 0x9b, 0x3c, 0x54, 0xf0, 0x46, 0xe1, 0x32, 0xca,
	0xc7, 0x22, 0x34, 0x33, 0x50, 0x72, 0xf4, 0x23, 0xbe, 0x81, 0x90, 0x04, 0x29, 0x19, 0x0f, 0x75,
	0x86, 0xf3, 0x46, 0x51, 0x4a, 0x25, 0x93, 0x49, 0xce, 0xfd, 0x8d, 0x24, 0xf1, 0x75, 0x54, 0xea,
	0x50, 0xf7, 0x48, 0x09, 0xea, 0x66, 0xd9, 0x8c, 0x04, 0xf8, 0x26, 0x9a, 0x33, 0xcc, 0x00, 0x6d,
	0xdd, 0x63, 0xe9, 0xf1, 0x51, 0x22, 0xd2, 0x3d, 0xa8, 0x0d, 0x34, 0x27, 0x0e, 0x52, 0x83, 0xab,
	0x89, 0x41, 0x22, 0x32, 0x06, 0x77, 0xd0, 0xc2, 0x11, 0x0b, 0x02, 0xa3, 0x96, 0x8a, 0xf6, 0x23,
	0x6b, 0xda, 0xd8, 0xcc, 0x6b, 0x69, 0x2b, 0x13, 0xe2, 0x6d, 0xb4, 0x10, 0xa7, 0x35, 0x6e, 0xfb,
	0x2c, 0x00, 0x69, 0x59, 0x86, 0x1c, 0x6f, 0x9e, 0xca, 0xe2, 0xf4, 0x45, 0x38, 0xf3, 0xf1, 0x98,
	0x44, 0xe2, 0x35, 0xb4, 0xa8, 0xb8, 0xa2, 0x41, 0xfb, 0x54, 0xb4, 0x65, 0x03, 0x8a, 0x8d, 0xee,
	0x70, 0xc2, 0x63, 0x03, 0xdd, 0x48, 0x3c, 0xe0, 0x4d, 0x04, 0xae, 0x02, 0x6f, 0xe4, 0xda, 0x39,
	0x56, 0x20, 0xad, 0x15, 0xe3, 0xba, 0x6a, 0x8c, 0x1a, 0xa9, 0x4d, 0x16, 0xa2, 0xae, 0x2d, 0xce,
	0x01, 0x4d, 0x3c, 0x57, 0xcf, 0x01, 0x4d, 0x3c, 0xee, 0xa0, 0xe9, 0x80, 0x9b, 0x31, 0xd0, 0x49,
	0x5e, 0x4d, 0x93, 0xdc, 0xe1, 0xdd, 0x67, 0x20, 0x25, 0xed, 0x82, 0x63, 0xd4, 0xf8, 0x33, 0x34,
	0x23, 0x15, 0x55, 0x60, 0x2d, 0x90, 0x5c, 0x75, 0x61, 0xfd, 0xce, 0x45, 0x57, 0x9a, 0x76, 0x8d,
	0x7d, 0xa0, 0x8d, 0x9d, 0xc4, 0x07, 0xef, 0xa1, 0x82, 0x7e, 0x88, 0xa5, 0xf5, 0x5f, 0x7d, 0xad,
	0xf5, 0x4f, 0xcd, 0xf0, 0xdf, 0xc7, 0x35, 0x63, 0x1d, 0x2a, 0xa9, 0x87, 0x99, 0x86, 0x04, 0x84,
	0xe0, 0x82, 0x24, 0xa6, 0xc4, 0x0c, 0x7b, 0xe7, 0xd8, 0x0c, 0x79, 0xc2, 0xfb, 0xb6, 0x93, 0x86,
	0xc1, 0xb7, 0xd0, 0x7c, 0x2c, 0x41, 0xb4, 0x43, 0xae, 0x98, 0xcf, 0xc0, 0xb3, 0xca, 0x66, 0x0b,
	0x5c, 0xd1, 0xc2, 0xdd, 0x54, 0x86, 0x7f, 0xd4, 0x1c, 0x94, 0xb1, 0x6f, 0x7b, 0xc8, 0x54, 0xaf,
	0x2d, 0x40, 0xc6, 0x81, 0x92, 0xd6, 0x35, 0xb3, 0x19, 0x98, 0x39, 0x86, 0x8b, 0xa9, 0xde, 0x0c,
	0x7e, 0x1c, 0x04, 0x44, 0x0f, 0xcc, 0x19, 0xe2, 0x26, 0xda, 0x55, 0x8b, 0x98, 0x20, 0x69, 0x00,
	0x9b, 0xb4, 0x7a, 0x4c, 0x12, 0x45, 0x8f, 0xf4, 0x1e, 0xd4, 0xde, 0x5c, 0x8c, 0xb1, 0x95, 0x9e,
	0xc1, 0x9a, 0xe4, 0xb1, 0x70, 0x13, 0x7e, 0xd2, 0xac, 0x94, 0x45, 0x7b, 0xc9, 0x54, 0xcf, 0x49,
	0xa2, 0xe0, 0x45, 0x34, 0xe3, 0x31, 0xa1, 0x8e, 0xcd, 0x8c, 0xce, 0x3a, 0xc9, 0x4b, 0xe5, 0x29,
	0x9a, 0x31, 0xd5, 0xc3, 0x25, 0x34, 0x73, 0xb8, 0x7b, 0xd0, 0x68, 0x95, 0xff, 0x83, 0xe7, 0x50,
	0xd1, 0x39, 0xdc, 0xdd, 0x6d, 0xee, 0x7e, 0x55, 0xce, 0xe1, 0x05, 0x84, 0x5a, 0x0d, 0xe7, 0x59,
	0x73, 0x77, 0xa3, 0xd5, 0xd8, 0x2a, 0x4f, 0x69, 0xbb, 0x86, 0xe3, 0xec, 0x39, 0xe5, 0x3c, 0xbe,
	0x82, 0x66, 0x37, 0x9c, 0xcd, 0xa7, 0xcd, 0x17, 0x8d, 0xad, 0xf2, 0x74, 0xa5, 0x87, 0xca, 0x9b,
	0xa6, 0x72, 0x8d, 0x01, 0x84, 0xaa, 0xa5, 0xd7, 0xb6, 0x5e, 0xb8, 0x03, 0x10, 0x7a, 0x54, 0xcd,
	0x8c, 0x4d, 0x3b, 0xd9, 0x2b, 0x7e, 0x32, 0xb9, 0x3b, 0xdf, 0x3d, 0xb9, 0x23, 0xf3, 0xca, 0x21,
	0x2a, 0xff, 0x1b, 0xc4, 0xf4, 0x47, 0x09, 0xcd, 0x6d, 0x6b, 0x3a, 0x4f, 0xb9, 0xe8, 0x72, 0x8a,
	0x78, 0x78, 0x0e, 0x45, 0xd4, 0xff, 0x67, 0x2e, 0x79, 0x1e, 0xcd, 0x39, 0x5b, 0xdb, 0x5b, 0x54,
	0x81, 0x56, 0x4d, 0xf0, 0xc6, 0x93, 0x33, 0xb4, 0x50, 0xb8, 0xd8, 0xf1, 0x14, 0x57, 0xdc, 0x43,
	0x38, 0x04, 0x35, 0xe4, 0xe2, 0x28, 0x99, 0xb3, 0xb6, 0x6e, 0xda, 0x64, 0xb1, 0x3a, 0xe5, 0x54,
	0x63, 0xc6, 0xec, 0x40, 0x7f, 0x9c, 0xc4, 0x17, 0x30, 0x4b, 0xb6, 0x7b, 0x4e, 0x57, 0xb0, 0xfe,
	0xc8, 0x1c, 0xa1, 0x86, 0x3f, 0xd9, 0x20, 0x01, 0x93, 0x4a, 0x4f, 0x89, 0xf1, 0x23, 0x59, 0x18,
	0x42, 0x25, 0x89, 0xa8, 0x50, 0xd9, 0x32, 0xd4, 0x6b, 0xc6, 0x3e, 0x4d, 0x44, 0xdf, 0x5c, 0x46,
	0x44, 0xf5, 0x8f, 0x0c, 0xc6, 0x07, 0xf8, 0x66, 0x4b, 0xdb, 0x90, 0xf0, 0xe4, 0x2b, 0xe0, 0x04,
	0xc3, 0x58, 0xdb, 0xe7, 0x32, 0x16, 0xff, 0x4b, 0x8c, 0x55, 0xbf, 0x67, 0x30, 0x3e, 0xc4, 0xb7,
	0x9f, 0xf2, 0x21, 0xe9, 0xd3, 0xf0, 0x98, 0x18, 0xad, 0xde, 0xf9, 0x89, 0xa3, 0xf9, 0xd4, 0x00,
	0x17, 0xd8, 0x00, 0xec, 0x4b, 0xf9, 0xed, 0xbb, 0xcb, 0xf8, 0xad, 0x7e, 0xd7, 0xe0, 0xdc, 0xc6,
	0x95, 0xb3, 0x38, 0x69, 0x74, 0x8f, 0x48, 0x4e, 0x7c, 0x2a, 0xec, 0x73, 0xb9, 0xf0, 0xcb, 0x77,
	0x70, 0x61, 0xdd, 0x32, 0x00, 0x18, 0xcf, 0xef, 0xa4, 0xd7, 0xa1, 0xcd, 0xed, 0xd5, 0xdc, 0x54,
	0xca, 0x92, 0xef, 0xd8, 0x8c, 0xf6, 0x24, 0x89, 0x5a, 0x29, 0xc0, 0x58, 0x87, 0xbf, 0x0f, 0xbc,
	0xf9, 0xf0, 0x9c, 0x45, 0x7a, 0xc1, 0x18, 0x8d, 0x6d, 0xd7, 0xee, 0x38, 0x7b, 0x2c, 0x19, 0x7e,
	0x6d, 0x1a, 0x9f, 0x4d, 0xbc, 0xd1, 0x1c, 0x75, 0x27, 0x11, 0xa0, 0x62, 0x11, 0x4a, 0x42, 0x89,
	0x04, 0x53, 0xc2, 0xb1, 0x4f, 0x72, 0xcd, 0xa9, 0x3e, 0x83, 0xc0, 0x4b, 0x3e, 0x05, 0x4d, 0xcf,
	0xab, 0x1e, 0xf4, 0xed, 0xf1, 0x4f, 0xf3, 0xf7, 0x9e, 0xd6, 0xff, 0x39, 0x02, 0xaf, 0x3f, 0xf8,
	0xf6, 0xfe, 0x70, 0x38, 0xb4, 0x07, 0x10, 0x70, 0x97, 0x79, 0xf0, 0xc6, 0x76, 0x79, 0xbf, 0xd6,
	0xe5, 0x01, 0x0d, 0xbb, 0xb5, 0x44, 0x28, 0x68, 0xa4, 0xb8, 0xa8, 0xe9, 0xd2, 0xa6, 0xff, 0x0b,
	0x76, 0x0a, 0xe6, 0xe7, 0xc1, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xb1, 0x3a, 0x88, 0x5a,
	0x0e, 0x00, 0x00,
}