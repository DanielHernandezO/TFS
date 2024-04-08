// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: datanode_file.proto

package config

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

type ChunkMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ChunkId   int32  `protobuf:"varint,2,opt,name=chunkId,proto3" json:"chunkId,omitempty"`
	ReplicaId int32  `protobuf:"varint,3,opt,name=replicaId,proto3" json:"replicaId,omitempty"`
}

func (x *ChunkMetadata) Reset() {
	*x = ChunkMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkMetadata) ProtoMessage() {}

func (x *ChunkMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkMetadata.ProtoReflect.Descriptor instead.
func (*ChunkMetadata) Descriptor() ([]byte, []int) {
	return file_datanode_file_proto_rawDescGZIP(), []int{0}
}

func (x *ChunkMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChunkMetadata) GetChunkId() int32 {
	if x != nil {
		return x.ChunkId
	}
	return 0
}

func (x *ChunkMetadata) GetReplicaId() int32 {
	if x != nil {
		return x.ReplicaId
	}
	return 0
}

type DataNodeSocket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *DataNodeSocket) Reset() {
	*x = DataNodeSocket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataNodeSocket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNodeSocket) ProtoMessage() {}

func (x *DataNodeSocket) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNodeSocket.ProtoReflect.Descriptor instead.
func (*DataNodeSocket) Descriptor() ([]byte, []int) {
	return file_datanode_file_proto_rawDescGZIP(), []int{1}
}

func (x *DataNodeSocket) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *DataNodeSocket) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type ChunkWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkMetadata     *ChunkMetadata    `protobuf:"bytes,1,opt,name=chunkMetadata,proto3" json:"chunkMetadata,omitempty"`
	PipelineDataNodes []*DataNodeSocket `protobuf:"bytes,2,rep,name=pipelineDataNodes,proto3" json:"pipelineDataNodes,omitempty"`
	ChunkBytes        []byte            `protobuf:"bytes,3,opt,name=chunkBytes,proto3" json:"chunkBytes,omitempty"`
}

func (x *ChunkWriteRequest) Reset() {
	*x = ChunkWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkWriteRequest) ProtoMessage() {}

func (x *ChunkWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkWriteRequest.ProtoReflect.Descriptor instead.
func (*ChunkWriteRequest) Descriptor() ([]byte, []int) {
	return file_datanode_file_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkWriteRequest) GetChunkMetadata() *ChunkMetadata {
	if x != nil {
		return x.ChunkMetadata
	}
	return nil
}

func (x *ChunkWriteRequest) GetPipelineDataNodes() []*DataNodeSocket {
	if x != nil {
		return x.PipelineDataNodes
	}
	return nil
}

func (x *ChunkWriteRequest) GetChunkBytes() []byte {
	if x != nil {
		return x.ChunkBytes
	}
	return nil
}

type ChunkReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkMetadata *ChunkMetadata `protobuf:"bytes,1,opt,name=chunkMetadata,proto3" json:"chunkMetadata,omitempty"`
}

func (x *ChunkReadRequest) Reset() {
	*x = ChunkReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkReadRequest) ProtoMessage() {}

func (x *ChunkReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkReadRequest.ProtoReflect.Descriptor instead.
func (*ChunkReadRequest) Descriptor() ([]byte, []int) {
	return file_datanode_file_proto_rawDescGZIP(), []int{3}
}

func (x *ChunkReadRequest) GetChunkMetadata() *ChunkMetadata {
	if x != nil {
		return x.ChunkMetadata
	}
	return nil
}

type ChunkWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code          int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ChunkMetadata *ChunkMetadata `protobuf:"bytes,3,opt,name=chunkMetadata,proto3" json:"chunkMetadata,omitempty"`
}

func (x *ChunkWriteResponse) Reset() {
	*x = ChunkWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkWriteResponse) ProtoMessage() {}

func (x *ChunkWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkWriteResponse.ProtoReflect.Descriptor instead.
func (*ChunkWriteResponse) Descriptor() ([]byte, []int) {
	return file_datanode_file_proto_rawDescGZIP(), []int{4}
}

func (x *ChunkWriteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChunkWriteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChunkWriteResponse) GetChunkMetadata() *ChunkMetadata {
	if x != nil {
		return x.ChunkMetadata
	}
	return nil
}

type ChunkReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code          int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ChunkMetadata *ChunkMetadata `protobuf:"bytes,3,opt,name=chunkMetadata,proto3" json:"chunkMetadata,omitempty"`
	ChunkBytes    []byte         `protobuf:"bytes,4,opt,name=chunkBytes,proto3" json:"chunkBytes,omitempty"`
}

func (x *ChunkReadResponse) Reset() {
	*x = ChunkReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkReadResponse) ProtoMessage() {}

func (x *ChunkReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkReadResponse.ProtoReflect.Descriptor instead.
func (*ChunkReadResponse) Descriptor() ([]byte, []int) {
	return file_datanode_file_proto_rawDescGZIP(), []int{5}
}

func (x *ChunkReadResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChunkReadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChunkReadResponse) GetChunkMetadata() *ChunkMetadata {
	if x != nil {
		return x.ChunkMetadata
	}
	return nil
}

func (x *ChunkReadResponse) GetChunkBytes() []byte {
	if x != nil {
		return x.ChunkBytes
	}
	return nil
}

var File_datanode_file_proto protoreflect.FileDescriptor

var file_datanode_file_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x49, 0x64, 0x22, 0x34, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x11, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x11, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x10, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x78, 0x0a,
	0x12, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x34, 0x0a, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x32, 0x7c, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x12, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x52,
	0x65, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x11, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x46,
	0x53, 0x2f, 0x54, 0x46, 0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datanode_file_proto_rawDescOnce sync.Once
	file_datanode_file_proto_rawDescData = file_datanode_file_proto_rawDesc
)

func file_datanode_file_proto_rawDescGZIP() []byte {
	file_datanode_file_proto_rawDescOnce.Do(func() {
		file_datanode_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_datanode_file_proto_rawDescData)
	})
	return file_datanode_file_proto_rawDescData
}

var file_datanode_file_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_datanode_file_proto_goTypes = []interface{}{
	(*ChunkMetadata)(nil),      // 0: ChunkMetadata
	(*DataNodeSocket)(nil),     // 1: DataNodeSocket
	(*ChunkWriteRequest)(nil),  // 2: ChunkWriteRequest
	(*ChunkReadRequest)(nil),   // 3: ChunkReadRequest
	(*ChunkWriteResponse)(nil), // 4: ChunkWriteResponse
	(*ChunkReadResponse)(nil),  // 5: ChunkReadResponse
}
var file_datanode_file_proto_depIdxs = []int32{
	0, // 0: ChunkWriteRequest.chunkMetadata:type_name -> ChunkMetadata
	1, // 1: ChunkWriteRequest.pipelineDataNodes:type_name -> DataNodeSocket
	0, // 2: ChunkReadRequest.chunkMetadata:type_name -> ChunkMetadata
	0, // 3: ChunkWriteResponse.chunkMetadata:type_name -> ChunkMetadata
	0, // 4: ChunkReadResponse.chunkMetadata:type_name -> ChunkMetadata
	2, // 5: DataNodeService.WriteChunk:input_type -> ChunkWriteRequest
	3, // 6: DataNodeService.ReadChunk:input_type -> ChunkReadRequest
	4, // 7: DataNodeService.WriteChunk:output_type -> ChunkWriteResponse
	5, // 8: DataNodeService.ReadChunk:output_type -> ChunkReadResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_datanode_file_proto_init() }
func file_datanode_file_proto_init() {
	if File_datanode_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datanode_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkMetadata); i {
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
		file_datanode_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataNodeSocket); i {
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
		file_datanode_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkWriteRequest); i {
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
		file_datanode_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkReadRequest); i {
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
		file_datanode_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkWriteResponse); i {
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
		file_datanode_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkReadResponse); i {
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
			RawDescriptor: file_datanode_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datanode_file_proto_goTypes,
		DependencyIndexes: file_datanode_file_proto_depIdxs,
		MessageInfos:      file_datanode_file_proto_msgTypes,
	}.Build()
	File_datanode_file_proto = out.File
	file_datanode_file_proto_rawDesc = nil
	file_datanode_file_proto_goTypes = nil
	file_datanode_file_proto_depIdxs = nil
}
