// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: protos/device.proto

package go_device_microservice_proto

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

type SendCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId       string          `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	CommandRequest *CommandRequest `protobuf:"bytes,2,opt,name=command_request,json=commandRequest,proto3" json:"command_request,omitempty"`
}

func (x *SendCommandRequest) Reset() {
	*x = SendCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCommandRequest) ProtoMessage() {}

func (x *SendCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCommandRequest.ProtoReflect.Descriptor instead.
func (*SendCommandRequest) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{0}
}

func (x *SendCommandRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *SendCommandRequest) GetCommandRequest() *CommandRequest {
	if x != nil {
		return x.CommandRequest
	}
	return nil
}

type DeviceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*DeviceMessage_Heartbeat
	//	*DeviceMessage_CommandRequest
	//	*DeviceMessage_CommandResponse
	//	*DeviceMessage_StatusUpdate
	Payload isDeviceMessage_Payload `protobuf_oneof:"payload"`
}

func (x *DeviceMessage) Reset() {
	*x = DeviceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceMessage) ProtoMessage() {}

func (x *DeviceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceMessage.ProtoReflect.Descriptor instead.
func (*DeviceMessage) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceMessage) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (m *DeviceMessage) GetPayload() isDeviceMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *DeviceMessage) GetHeartbeat() *Heartbeat {
	if x, ok := x.GetPayload().(*DeviceMessage_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (x *DeviceMessage) GetCommandRequest() *CommandRequest {
	if x, ok := x.GetPayload().(*DeviceMessage_CommandRequest); ok {
		return x.CommandRequest
	}
	return nil
}

func (x *DeviceMessage) GetCommandResponse() *CommandResponse {
	if x, ok := x.GetPayload().(*DeviceMessage_CommandResponse); ok {
		return x.CommandResponse
	}
	return nil
}

func (x *DeviceMessage) GetStatusUpdate() *StatusUpdate {
	if x, ok := x.GetPayload().(*DeviceMessage_StatusUpdate); ok {
		return x.StatusUpdate
	}
	return nil
}

type isDeviceMessage_Payload interface {
	isDeviceMessage_Payload()
}

type DeviceMessage_Heartbeat struct {
	Heartbeat *Heartbeat `protobuf:"bytes,2,opt,name=heartbeat,proto3,oneof"`
}

type DeviceMessage_CommandRequest struct {
	CommandRequest *CommandRequest `protobuf:"bytes,3,opt,name=command_request,json=commandRequest,proto3,oneof"`
}

type DeviceMessage_CommandResponse struct {
	CommandResponse *CommandResponse `protobuf:"bytes,4,opt,name=command_response,json=commandResponse,proto3,oneof"`
}

type DeviceMessage_StatusUpdate struct {
	StatusUpdate *StatusUpdate `protobuf:"bytes,5,opt,name=status_update,json=statusUpdate,proto3,oneof"`
}

func (*DeviceMessage_Heartbeat) isDeviceMessage_Payload() {}

func (*DeviceMessage_CommandRequest) isDeviceMessage_Payload() {}

func (*DeviceMessage_CommandResponse) isDeviceMessage_Payload() {}

func (*DeviceMessage_StatusUpdate) isDeviceMessage_Payload() {}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{2}
}

func (x *Heartbeat) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId   string `protobuf:"bytes,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	CommandText string `protobuf:"bytes,2,opt,name=command_text,json=commandText,proto3" json:"command_text,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{3}
}

func (x *CommandRequest) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *CommandRequest) GetCommandText() string {
	if x != nil {
		return x.CommandText
	}
	return ""
}

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId    string `protobuf:"bytes,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Output       string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Success      bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage string `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // Optional
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{4}
}

func (x *CommandResponse) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *CommandResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *CommandResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CommandResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type StatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusUpdate) Reset() {
	*x = StatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUpdate) ProtoMessage() {}

func (x *StatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUpdate.ProtoReflect.Descriptor instead.
func (*StatusUpdate) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{5}
}

func (x *StatusUpdate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GenerateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigText   string `protobuf:"bytes,1,opt,name=config_text,json=configText,proto3" json:"config_text,omitempty"` // Long string for configuration
	InstanceText string `protobuf:"bytes,2,opt,name=instance_text,json=instanceText,proto3" json:"instance_text,omitempty"`
}

func (x *GenerateConfigRequest) Reset() {
	*x = GenerateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateConfigRequest) ProtoMessage() {}

func (x *GenerateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateConfigRequest.ProtoReflect.Descriptor instead.
func (*GenerateConfigRequest) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{6}
}

func (x *GenerateConfigRequest) GetConfigText() string {
	if x != nil {
		return x.ConfigText
	}
	return ""
}

func (x *GenerateConfigRequest) GetInstanceText() string {
	if x != nil {
		return x.InstanceText
	}
	return ""
}

type GenerateConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // Optional
}

func (x *GenerateConfigResponse) Reset() {
	*x = GenerateConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateConfigResponse) ProtoMessage() {}

func (x *GenerateConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateConfigResponse.ProtoReflect.Descriptor instead.
func (*GenerateConfigResponse) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateConfigResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GenerateConfigResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_protos_device_proto protoreflect.FileDescriptor

var file_protos_device_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x72, 0x0a,
	0x12, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x3f, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xb0, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48, 0x00, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x12, 0x41, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x29, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x52, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54,
	0x65, 0x78, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x22, 0x57, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xeb, 0x01,
	0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x41, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x15,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x44, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x1a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x72, 0x6e, 0x65, 0x74, 0x66,
	0x6f, 0x72, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_device_proto_rawDescOnce sync.Once
	file_protos_device_proto_rawDescData = file_protos_device_proto_rawDesc
)

func file_protos_device_proto_rawDescGZIP() []byte {
	file_protos_device_proto_rawDescOnce.Do(func() {
		file_protos_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_device_proto_rawDescData)
	})
	return file_protos_device_proto_rawDescData
}

var file_protos_device_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_device_proto_goTypes = []any{
	(*SendCommandRequest)(nil),     // 0: device.SendCommandRequest
	(*DeviceMessage)(nil),          // 1: device.DeviceMessage
	(*Heartbeat)(nil),              // 2: device.Heartbeat
	(*CommandRequest)(nil),         // 3: device.CommandRequest
	(*CommandResponse)(nil),        // 4: device.CommandResponse
	(*StatusUpdate)(nil),           // 5: device.StatusUpdate
	(*GenerateConfigRequest)(nil),  // 6: device.GenerateConfigRequest
	(*GenerateConfigResponse)(nil), // 7: device.GenerateConfigResponse
}
var file_protos_device_proto_depIdxs = []int32{
	3, // 0: device.SendCommandRequest.command_request:type_name -> device.CommandRequest
	2, // 1: device.DeviceMessage.heartbeat:type_name -> device.Heartbeat
	3, // 2: device.DeviceMessage.command_request:type_name -> device.CommandRequest
	4, // 3: device.DeviceMessage.command_response:type_name -> device.CommandResponse
	5, // 4: device.DeviceMessage.status_update:type_name -> device.StatusUpdate
	1, // 5: device.DeviceGateway.Communicate:input_type -> device.DeviceMessage
	0, // 6: device.DeviceGateway.SendCommand:input_type -> device.SendCommandRequest
	6, // 7: device.DeviceGateway.GenerateConfig:input_type -> device.GenerateConfigRequest
	1, // 8: device.DeviceGateway.Communicate:output_type -> device.DeviceMessage
	4, // 9: device.DeviceGateway.SendCommand:output_type -> device.CommandResponse
	7, // 10: device.DeviceGateway.GenerateConfig:output_type -> device.GenerateConfigResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_device_proto_init() }
func file_protos_device_proto_init() {
	if File_protos_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_device_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SendCommandRequest); i {
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
		file_protos_device_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceMessage); i {
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
		file_protos_device_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Heartbeat); i {
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
		file_protos_device_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CommandRequest); i {
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
		file_protos_device_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CommandResponse); i {
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
		file_protos_device_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*StatusUpdate); i {
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
		file_protos_device_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateConfigRequest); i {
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
		file_protos_device_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateConfigResponse); i {
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
	file_protos_device_proto_msgTypes[1].OneofWrappers = []any{
		(*DeviceMessage_Heartbeat)(nil),
		(*DeviceMessage_CommandRequest)(nil),
		(*DeviceMessage_CommandResponse)(nil),
		(*DeviceMessage_StatusUpdate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_device_proto_goTypes,
		DependencyIndexes: file_protos_device_proto_depIdxs,
		MessageInfos:      file_protos_device_proto_msgTypes,
	}.Build()
	File_protos_device_proto = out.File
	file_protos_device_proto_rawDesc = nil
	file_protos_device_proto_goTypes = nil
	file_protos_device_proto_depIdxs = nil
}
