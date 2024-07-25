// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: rpc/worker.proto

package worker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_IDLE        Status = 0
	Status_IN_PROGRESS Status = 1
	Status_COMPLETED   Status = 2
	Status_FAILED      Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "IDLE",
		1: "IN_PROGRESS",
		2: "COMPLETED",
		3: "FAILED",
	}
	Status_value = map[string]int32{
		"IDLE":        0,
		"IN_PROGRESS": 1,
		"COMPLETED":   2,
		"FAILED":      3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_worker_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_rpc_worker_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{0}
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{0}
}

func (x *StatusRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	TaskId       string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{1}
}

func (x *StatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_IDLE
}

func (x *StatusResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *StatusResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type HealthcheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthcheckRequest) Reset() {
	*x = HealthcheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthcheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthcheckRequest) ProtoMessage() {}

func (x *HealthcheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthcheckRequest.ProtoReflect.Descriptor instead.
func (*HealthcheckRequest) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{2}
}

type MapTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Filename  string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	NumReduce int32  `protobuf:"varint,3,opt,name=num_reduce,json=numReduce,proto3" json:"num_reduce,omitempty"`
}

func (x *MapTask) Reset() {
	*x = MapTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapTask) ProtoMessage() {}

func (x *MapTask) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapTask.ProtoReflect.Descriptor instead.
func (*MapTask) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{3}
}

func (x *MapTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *MapTask) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *MapTask) GetNumReduce() int32 {
	if x != nil {
		return x.NumReduce
	}
	return 0
}

type ReduceTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    string         `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Datanodes *DataNodesInfo `protobuf:"bytes,2,opt,name=datanodes,proto3" json:"datanodes,omitempty"`
}

func (x *ReduceTask) Reset() {
	*x = ReduceTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceTask) ProtoMessage() {}

func (x *ReduceTask) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceTask.ProtoReflect.Descriptor instead.
func (*ReduceTask) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{4}
}

func (x *ReduceTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *ReduceTask) GetDatanodes() *DataNodesInfo {
	if x != nil {
		return x.Datanodes
	}
	return nil
}

type DataNodesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeFileInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *DataNodesInfo) Reset() {
	*x = DataNodesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataNodesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNodesInfo) ProtoMessage() {}

func (x *DataNodesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNodesInfo.ProtoReflect.Descriptor instead.
func (*DataNodesInfo) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{5}
}

func (x *DataNodesInfo) GetNodes() []*NodeFileInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type NodeFileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Ip    string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Files []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *NodeFileInfo) Reset() {
	*x = NodeFileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeFileInfo) ProtoMessage() {}

func (x *NodeFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeFileInfo.ProtoReflect.Descriptor instead.
func (*NodeFileInfo) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{6}
}

func (x *NodeFileInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *NodeFileInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *NodeFileInfo) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{7}
}

func (x *Ack) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type InterMediateDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *InterMediateDataRequest) Reset() {
	*x = InterMediateDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterMediateDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterMediateDataRequest) ProtoMessage() {}

func (x *InterMediateDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterMediateDataRequest.ProtoReflect.Descriptor instead.
func (*InterMediateDataRequest) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{8}
}

func (x *InterMediateDataRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type InterMediateDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InterMediateDataResponse) Reset() {
	*x = InterMediateDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterMediateDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterMediateDataResponse) ProtoMessage() {}

func (x *InterMediateDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterMediateDataResponse.ProtoReflect.Descriptor instead.
func (*InterMediateDataResponse) Descriptor() ([]byte, []int) {
	return file_rpc_worker_proto_rawDescGZIP(), []int{9}
}

func (x *InterMediateDataResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_rpc_worker_proto protoreflect.FileDescriptor

var file_rpc_worker_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a,
	0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x22, 0x53, 0x0a, 0x0a, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x48, 0x0a,
	0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x35, 0x0a, 0x17, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x2e, 0x0a, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a,
	0x3e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x44, 0x4c,
	0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x32,
	0xec, 0x01, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x4d, 0x61, 0x70, 0x12, 0x08, 0x2e, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x73,
	0x6b, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x12, 0x0b, 0x2e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x0b, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x04,
	0x2e, 0x41, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27,
	0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x6b,
	0x64, 0x61, 0x2f, 0x6d, 0x61, 0x70, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_worker_proto_rawDescOnce sync.Once
	file_rpc_worker_proto_rawDescData = file_rpc_worker_proto_rawDesc
)

func file_rpc_worker_proto_rawDescGZIP() []byte {
	file_rpc_worker_proto_rawDescOnce.Do(func() {
		file_rpc_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_worker_proto_rawDescData)
	})
	return file_rpc_worker_proto_rawDescData
}

var file_rpc_worker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rpc_worker_proto_goTypes = []any{
	(Status)(0),                      // 0: Status
	(*StatusRequest)(nil),            // 1: StatusRequest
	(*StatusResponse)(nil),           // 2: StatusResponse
	(*HealthcheckRequest)(nil),       // 3: HealthcheckRequest
	(*MapTask)(nil),                  // 4: MapTask
	(*ReduceTask)(nil),               // 5: ReduceTask
	(*DataNodesInfo)(nil),            // 6: DataNodesInfo
	(*NodeFileInfo)(nil),             // 7: NodeFileInfo
	(*Ack)(nil),                      // 8: Ack
	(*InterMediateDataRequest)(nil),  // 9: InterMediateDataRequest
	(*InterMediateDataResponse)(nil), // 10: InterMediateDataResponse
}
var file_rpc_worker_proto_depIdxs = []int32{
	0,  // 0: StatusResponse.status:type_name -> Status
	6,  // 1: ReduceTask.datanodes:type_name -> DataNodesInfo
	7,  // 2: DataNodesInfo.nodes:type_name -> NodeFileInfo
	4,  // 3: Worker.AssignMap:input_type -> MapTask
	5,  // 4: Worker.AssignReduce:input_type -> ReduceTask
	3,  // 5: Worker.HealthCheck:input_type -> HealthcheckRequest
	1,  // 6: Worker.GetStatus:input_type -> StatusRequest
	9,  // 7: Worker.GetIntermediateData:input_type -> InterMediateDataRequest
	8,  // 8: Worker.AssignMap:output_type -> Ack
	8,  // 9: Worker.AssignReduce:output_type -> Ack
	8,  // 10: Worker.HealthCheck:output_type -> Ack
	2,  // 11: Worker.GetStatus:output_type -> StatusResponse
	10, // 12: Worker.GetIntermediateData:output_type -> InterMediateDataResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_worker_proto_init() }
func file_rpc_worker_proto_init() {
	if File_rpc_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_worker_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HealthcheckRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MapTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ReduceTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DataNodesInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*NodeFileInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*InterMediateDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_worker_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*InterMediateDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_worker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_worker_proto_goTypes,
		DependencyIndexes: file_rpc_worker_proto_depIdxs,
		EnumInfos:         file_rpc_worker_proto_enumTypes,
		MessageInfos:      file_rpc_worker_proto_msgTypes,
	}.Build()
	File_rpc_worker_proto = out.File
	file_rpc_worker_proto_rawDesc = nil
	file_rpc_worker_proto_goTypes = nil
	file_rpc_worker_proto_depIdxs = nil
}
