// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: application/grpc/proto/notifications.proto

package grpc_protobuf

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NotificationChannel int32

const (
	NotificationChannel_EMAIL   NotificationChannel = 0
	NotificationChannel_SMS     NotificationChannel = 1
	NotificationChannel_BROWSER NotificationChannel = 2
)

// Enum value maps for NotificationChannel.
var (
	NotificationChannel_name = map[int32]string{
		0: "EMAIL",
		1: "SMS",
		2: "BROWSER",
	}
	NotificationChannel_value = map[string]int32{
		"EMAIL":   0,
		"SMS":     1,
		"BROWSER": 2,
	}
)

func (x NotificationChannel) Enum() *NotificationChannel {
	p := new(NotificationChannel)
	*p = x
	return p
}

func (x NotificationChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_application_grpc_proto_notifications_proto_enumTypes[0].Descriptor()
}

func (NotificationChannel) Type() protoreflect.EnumType {
	return &file_application_grpc_proto_notifications_proto_enumTypes[0]
}

func (x NotificationChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationChannel.Descriptor instead.
func (NotificationChannel) EnumDescriptor() ([]byte, []int) {
	return file_application_grpc_proto_notifications_proto_rawDescGZIP(), []int{0}
}

type ListNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListNotificationsRequest) Reset() {
	*x = ListNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_grpc_proto_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotificationsRequest) ProtoMessage() {}

func (x *ListNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_grpc_proto_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotificationsRequest.ProtoReflect.Descriptor instead.
func (*ListNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_application_grpc_proto_notifications_proto_rawDescGZIP(), []int{0}
}

type ListNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *ListNotificationsResponse) Reset() {
	*x = ListNotificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_grpc_proto_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotificationsResponse) ProtoMessage() {}

func (x *ListNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_application_grpc_proto_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotificationsResponse.ProtoReflect.Descriptor instead.
func (*ListNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_application_grpc_proto_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *ListNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type GetNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId string `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
}

func (x *GetNotificationRequest) Reset() {
	*x = GetNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_grpc_proto_notifications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationRequest) ProtoMessage() {}

func (x *GetNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_grpc_proto_notifications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationRequest) Descriptor() ([]byte, []int) {
	return file_application_grpc_proto_notifications_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotificationRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

type GetNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *GetNotificationResponse) Reset() {
	*x = GetNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_grpc_proto_notifications_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationResponse) ProtoMessage() {}

func (x *GetNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_application_grpc_proto_notifications_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationResponse) Descriptor() ([]byte, []int) {
	return file_application_grpc_proto_notifications_proto_rawDescGZIP(), []int{3}
}

func (x *GetNotificationResponse) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

type SendNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId string `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
}

func (x *SendNotificationRequest) Reset() {
	*x = SendNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_grpc_proto_notifications_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationRequest) ProtoMessage() {}

func (x *SendNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_grpc_proto_notifications_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationRequest.ProtoReflect.Descriptor instead.
func (*SendNotificationRequest) Descriptor() ([]byte, []int) {
	return file_application_grpc_proto_notifications_proto_rawDescGZIP(), []int{4}
}

func (x *SendNotificationRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AppKey    string              `protobuf:"bytes,3,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
	SegmentId string              `protobuf:"bytes,4,opt,name=segment_id,json=segmentId,proto3" json:"segment_id,omitempty"`
	Channel   NotificationChannel `protobuf:"varint,5,opt,name=channel,proto3,enum=grpc.NotificationChannel" json:"channel,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_grpc_proto_notifications_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_application_grpc_proto_notifications_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_application_grpc_proto_notifications_proto_rawDescGZIP(), []int{5}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Notification) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *Notification) GetSegmentId() string {
	if x != nil {
		return x.SegmentId
	}
	return ""
}

func (x *Notification) GetChannel() NotificationChannel {
	if x != nil {
		return x.Channel
	}
	return NotificationChannel_EMAIL
}

var File_application_grpc_proto_notifications_proto protoreflect.FileDescriptor

var file_application_grpc_proto_notifications_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55,
	0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x41, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x17, 0x53,
	0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x9f, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2a, 0x36, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49,
	0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x4d, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x42, 0x52, 0x4f, 0x57, 0x53, 0x45, 0x52, 0x10, 0x02, 0x32, 0x9d, 0x02, 0x0a, 0x1c, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x57, 0x0a, 0x1a, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x6b, 0x0a, 0x1a, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x63, 0x6e, 0x61, 0x73, 0x63, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x6f, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2d, 0x6b, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x2d,
	0x70, 0x6f, 0x63, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_application_grpc_proto_notifications_proto_rawDescOnce sync.Once
	file_application_grpc_proto_notifications_proto_rawDescData = file_application_grpc_proto_notifications_proto_rawDesc
)

func file_application_grpc_proto_notifications_proto_rawDescGZIP() []byte {
	file_application_grpc_proto_notifications_proto_rawDescOnce.Do(func() {
		file_application_grpc_proto_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_grpc_proto_notifications_proto_rawDescData)
	})
	return file_application_grpc_proto_notifications_proto_rawDescData
}

var file_application_grpc_proto_notifications_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_application_grpc_proto_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_application_grpc_proto_notifications_proto_goTypes = []interface{}{
	(NotificationChannel)(0),          // 0: grpc.NotificationChannel
	(*ListNotificationsRequest)(nil),  // 1: grpc.ListNotificationsRequest
	(*ListNotificationsResponse)(nil), // 2: grpc.ListNotificationsResponse
	(*GetNotificationRequest)(nil),    // 3: grpc.GetNotificationRequest
	(*GetNotificationResponse)(nil),   // 4: grpc.GetNotificationResponse
	(*SendNotificationRequest)(nil),   // 5: grpc.SendNotificationRequest
	(*Notification)(nil),              // 6: grpc.Notification
	(*wrapperspb.BoolValue)(nil),      // 7: google.protobuf.BoolValue
}
var file_application_grpc_proto_notifications_proto_depIdxs = []int32{
	6, // 0: grpc.ListNotificationsResponse.notifications:type_name -> grpc.Notification
	6, // 1: grpc.GetNotificationResponse.notification:type_name -> grpc.Notification
	0, // 2: grpc.Notification.channel:type_name -> grpc.NotificationChannel
	3, // 3: grpc.NotificationsServiceFrontend.GetNotification:input_type -> grpc.GetNotificationRequest
	1, // 4: grpc.NotificationsServiceFrontend.ListNotifications:input_type -> grpc.ListNotificationsRequest
	5, // 5: grpc.NotificationsServiceFrontend.EnqueueSendingNotification:input_type -> grpc.SendNotificationRequest
	5, // 6: grpc.NotificationsServiceWorker.SendNotification:input_type -> grpc.SendNotificationRequest
	4, // 7: grpc.NotificationsServiceFrontend.GetNotification:output_type -> grpc.GetNotificationResponse
	2, // 8: grpc.NotificationsServiceFrontend.ListNotifications:output_type -> grpc.ListNotificationsResponse
	7, // 9: grpc.NotificationsServiceFrontend.EnqueueSendingNotification:output_type -> google.protobuf.BoolValue
	7, // 10: grpc.NotificationsServiceWorker.SendNotification:output_type -> google.protobuf.BoolValue
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_application_grpc_proto_notifications_proto_init() }
func file_application_grpc_proto_notifications_proto_init() {
	if File_application_grpc_proto_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_application_grpc_proto_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotificationsRequest); i {
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
		file_application_grpc_proto_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotificationsResponse); i {
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
		file_application_grpc_proto_notifications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationRequest); i {
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
		file_application_grpc_proto_notifications_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationResponse); i {
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
		file_application_grpc_proto_notifications_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotificationRequest); i {
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
		file_application_grpc_proto_notifications_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_application_grpc_proto_notifications_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_application_grpc_proto_notifications_proto_goTypes,
		DependencyIndexes: file_application_grpc_proto_notifications_proto_depIdxs,
		EnumInfos:         file_application_grpc_proto_notifications_proto_enumTypes,
		MessageInfos:      file_application_grpc_proto_notifications_proto_msgTypes,
	}.Build()
	File_application_grpc_proto_notifications_proto = out.File
	file_application_grpc_proto_notifications_proto_rawDesc = nil
	file_application_grpc_proto_notifications_proto_goTypes = nil
	file_application_grpc_proto_notifications_proto_depIdxs = nil
}
