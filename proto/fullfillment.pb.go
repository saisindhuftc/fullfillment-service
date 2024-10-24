// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/fullfillment.proto

package proto

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

type AssignOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId          string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	DeliveryPersonId string `protobuf:"bytes,2,opt,name=deliveryPersonId,proto3" json:"deliveryPersonId,omitempty"`
}

func (x *AssignOrderRequest) Reset() {
	*x = AssignOrderRequest{}
	mi := &file_proto_fullfillment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignOrderRequest) ProtoMessage() {}

func (x *AssignOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignOrderRequest.ProtoReflect.Descriptor instead.
func (*AssignOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{0}
}

func (x *AssignOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *AssignOrderRequest) GetDeliveryPersonId() string {
	if x != nil {
		return x.DeliveryPersonId
	}
	return ""
}

type AssignOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AssignOrderResponse) Reset() {
	*x = AssignOrderResponse{}
	mi := &file_proto_fullfillment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignOrderResponse) ProtoMessage() {}

func (x *AssignOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignOrderResponse.ProtoReflect.Descriptor instead.
func (*AssignOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{1}
}

func (x *AssignOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetOrderStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *GetOrderStatusRequest) Reset() {
	*x = GetOrderStatusRequest{}
	mi := &file_proto_fullfillment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderStatusRequest) ProtoMessage() {}

func (x *GetOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*GetOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderStatusRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetOrderStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetOrderStatusResponse) Reset() {
	*x = GetOrderStatusResponse{}
	mi := &file_proto_fullfillment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderStatusResponse) ProtoMessage() {}

func (x *GetOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*GetOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderStatusResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GetOrderStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateOrderStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateOrderStatusRequest) Reset() {
	*x = UpdateOrderStatusRequest{}
	mi := &file_proto_fullfillment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusRequest) ProtoMessage() {}

func (x *UpdateOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOrderStatusRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateOrderStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateOrderStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateOrderStatusResponse) Reset() {
	*x = UpdateOrderStatusResponse{}
	mi := &file_proto_fullfillment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusResponse) ProtoMessage() {}

func (x *UpdateOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateOrderStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetOrdersByDeliveryPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryPersonId string `protobuf:"bytes,1,opt,name=deliveryPersonId,proto3" json:"deliveryPersonId,omitempty"`
}

func (x *GetOrdersByDeliveryPersonRequest) Reset() {
	*x = GetOrdersByDeliveryPersonRequest{}
	mi := &file_proto_fullfillment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersByDeliveryPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersByDeliveryPersonRequest) ProtoMessage() {}

func (x *GetOrdersByDeliveryPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersByDeliveryPersonRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersByDeliveryPersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrdersByDeliveryPersonRequest) GetDeliveryPersonId() string {
	if x != nil {
		return x.DeliveryPersonId
	}
	return ""
}

type GetOrdersByDeliveryPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrdersByDeliveryPersonResponse) Reset() {
	*x = GetOrdersByDeliveryPersonResponse{}
	mi := &file_proto_fullfillment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersByDeliveryPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersByDeliveryPersonResponse) ProtoMessage() {}

func (x *GetOrdersByDeliveryPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersByDeliveryPersonResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersByDeliveryPersonResponse) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{7}
}

func (x *GetOrdersByDeliveryPersonResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_proto_fullfillment_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fullfillment_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_fullfillment_proto_rawDescGZIP(), []int{8}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_fullfillment_proto protoreflect.FileDescriptor

var file_proto_fullfillment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x66, 0x69, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5a, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2d, 0x0a,
	0x13, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x31, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x4a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4c, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x33, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4e,
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x49,
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x39, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0xf1, 0x02, 0x0a, 0x12, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x56, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x42, 0x79, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fullfillment_proto_rawDescOnce sync.Once
	file_proto_fullfillment_proto_rawDescData = file_proto_fullfillment_proto_rawDesc
)

func file_proto_fullfillment_proto_rawDescGZIP() []byte {
	file_proto_fullfillment_proto_rawDescOnce.Do(func() {
		file_proto_fullfillment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fullfillment_proto_rawDescData)
	})
	return file_proto_fullfillment_proto_rawDescData
}

var file_proto_fullfillment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_fullfillment_proto_goTypes = []any{
	(*AssignOrderRequest)(nil),                // 0: proto.AssignOrderRequest
	(*AssignOrderResponse)(nil),               // 1: proto.AssignOrderResponse
	(*GetOrderStatusRequest)(nil),             // 2: proto.GetOrderStatusRequest
	(*GetOrderStatusResponse)(nil),            // 3: proto.GetOrderStatusResponse
	(*UpdateOrderStatusRequest)(nil),          // 4: proto.UpdateOrderStatusRequest
	(*UpdateOrderStatusResponse)(nil),         // 5: proto.UpdateOrderStatusResponse
	(*GetOrdersByDeliveryPersonRequest)(nil),  // 6: proto.GetOrdersByDeliveryPersonRequest
	(*GetOrdersByDeliveryPersonResponse)(nil), // 7: proto.GetOrdersByDeliveryPersonResponse
	(*Order)(nil),                             // 8: proto.Order
}
var file_proto_fullfillment_proto_depIdxs = []int32{
	8, // 0: proto.GetOrdersByDeliveryPersonResponse.orders:type_name -> proto.Order
	0, // 1: proto.FulfillmentService.AssignOrder:input_type -> proto.AssignOrderRequest
	2, // 2: proto.FulfillmentService.GetOrderStatus:input_type -> proto.GetOrderStatusRequest
	4, // 3: proto.FulfillmentService.UpdateOrderStatus:input_type -> proto.UpdateOrderStatusRequest
	6, // 4: proto.FulfillmentService.GetOrdersByDeliveryPerson:input_type -> proto.GetOrdersByDeliveryPersonRequest
	1, // 5: proto.FulfillmentService.AssignOrder:output_type -> proto.AssignOrderResponse
	3, // 6: proto.FulfillmentService.GetOrderStatus:output_type -> proto.GetOrderStatusResponse
	5, // 7: proto.FulfillmentService.UpdateOrderStatus:output_type -> proto.UpdateOrderStatusResponse
	7, // 8: proto.FulfillmentService.GetOrdersByDeliveryPerson:output_type -> proto.GetOrdersByDeliveryPersonResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_fullfillment_proto_init() }
func file_proto_fullfillment_proto_init() {
	if File_proto_fullfillment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_fullfillment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fullfillment_proto_goTypes,
		DependencyIndexes: file_proto_fullfillment_proto_depIdxs,
		MessageInfos:      file_proto_fullfillment_proto_msgTypes,
	}.Build()
	File_proto_fullfillment_proto = out.File
	file_proto_fullfillment_proto_rawDesc = nil
	file_proto_fullfillment_proto_goTypes = nil
	file_proto_fullfillment_proto_depIdxs = nil
}