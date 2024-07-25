// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: rpc/master.proto

package master

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

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type WorkerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Ip   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *WorkerInfo) Reset() {
	*x = WorkerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerInfo) ProtoMessage() {}

func (x *WorkerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerInfo.ProtoReflect.Descriptor instead.
func (*WorkerInfo) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *WorkerInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type MapResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Filenames []string `protobuf:"bytes,2,rep,name=filenames,proto3" json:"filenames,omitempty"`
}

func (x *MapResult) Reset() {
	*x = MapResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapResult) ProtoMessage() {}

func (x *MapResult) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapResult.ProtoReflect.Descriptor instead.
func (*MapResult) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{2}
}

func (x *MapResult) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MapResult) GetFilenames() []string {
	if x != nil {
		return x.Filenames
	}
	return nil
}

type ReduceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *ReduceResult) Reset() {
	*x = ReduceResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_master_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceResult) ProtoMessage() {}

func (x *ReduceResult) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceResult.ProtoReflect.Descriptor instead.
func (*ReduceResult) Descriptor() ([]byte, []int) {
	return file_rpc_master_proto_rawDescGZIP(), []int{3}
}

func (x *ReduceResult) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ReduceResult) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
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
		mi := &file_rpc_master_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[4]
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
	return file_rpc_master_proto_rawDescGZIP(), []int{4}
}

func (x *Ack) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
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
		mi := &file_rpc_master_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataNodesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNodesInfo) ProtoMessage() {}

func (x *DataNodesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[5]
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
	return file_rpc_master_proto_rawDescGZIP(), []int{5}
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
		mi := &file_rpc_master_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeFileInfo) ProtoMessage() {}

func (x *NodeFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_master_proto_msgTypes[6]
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
	return file_rpc_master_proto_rawDescGZIP(), []int{6}
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

var File_rpc_master_proto protoreflect.FileDescriptor

var file_rpc_master_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x21, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x3d, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x34, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x48, 0x0a,
	0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x32, 0xbd, 0x01, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x15, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x04, 0x2e,
	0x4a, 0x6f, 0x62, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x23,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x0a, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x04, 0x2e,
	0x41, 0x63, 0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x0e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x29, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x0d, 0x2e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x6b, 0x64, 0x61, 0x2f, 0x6d, 0x61, 0x70, 0x72,
	0x65, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_master_proto_rawDescOnce sync.Once
	file_rpc_master_proto_rawDescData = file_rpc_master_proto_rawDesc
)

func file_rpc_master_proto_rawDescGZIP() []byte {
	file_rpc_master_proto_rawDescOnce.Do(func() {
		file_rpc_master_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_master_proto_rawDescData)
	})
	return file_rpc_master_proto_rawDescData
}

var file_rpc_master_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_master_proto_goTypes = []any{
	(*Job)(nil),           // 0: Job
	(*WorkerInfo)(nil),    // 1: WorkerInfo
	(*MapResult)(nil),     // 2: MapResult
	(*ReduceResult)(nil),  // 3: ReduceResult
	(*Ack)(nil),           // 4: Ack
	(*DataNodesInfo)(nil), // 5: DataNodesInfo
	(*NodeFileInfo)(nil),  // 6: NodeFileInfo
}
var file_rpc_master_proto_depIdxs = []int32{
	6, // 0: DataNodesInfo.nodes:type_name -> NodeFileInfo
	0, // 1: Master.Trigger:input_type -> Job
	1, // 2: Master.RegisterWorker:input_type -> WorkerInfo
	2, // 3: Master.UpdateMapResult:input_type -> MapResult
	5, // 4: Master.UpdateDataNodes:input_type -> DataNodesInfo
	3, // 5: Master.UpdateReduceResult:input_type -> ReduceResult
	4, // 6: Master.Trigger:output_type -> Ack
	4, // 7: Master.RegisterWorker:output_type -> Ack
	4, // 8: Master.UpdateMapResult:output_type -> Ack
	4, // 9: Master.UpdateDataNodes:output_type -> Ack
	4, // 10: Master.UpdateReduceResult:output_type -> Ack
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_master_proto_init() }
func file_rpc_master_proto_init() {
	if File_rpc_master_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_master_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Job); i {
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
		file_rpc_master_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WorkerInfo); i {
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
		file_rpc_master_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MapResult); i {
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
		file_rpc_master_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReduceResult); i {
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
		file_rpc_master_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
		file_rpc_master_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
		file_rpc_master_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_master_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_master_proto_goTypes,
		DependencyIndexes: file_rpc_master_proto_depIdxs,
		MessageInfos:      file_rpc_master_proto_msgTypes,
	}.Build()
	File_rpc_master_proto = out.File
	file_rpc_master_proto_rawDesc = nil
	file_rpc_master_proto_goTypes = nil
	file_rpc_master_proto_depIdxs = nil
}