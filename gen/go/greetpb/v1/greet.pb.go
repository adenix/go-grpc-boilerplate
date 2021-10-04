// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: greetpb/v1/greet.proto

package greetpb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Greeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Greeting) Reset() {
	*x = Greeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_v1_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeting) ProtoMessage() {}

func (x *Greeting) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_v1_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeting.ProtoReflect.Descriptor instead.
func (*Greeting) Descriptor() ([]byte, []int) {
	return file_greetpb_v1_greet_proto_rawDescGZIP(), []int{0}
}

func (x *Greeting) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Greeting) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_v1_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_v1_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_v1_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_v1_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_v1_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_v1_greet_proto_rawDescGZIP(), []int{2}
}

func (x *GreetResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GreetManyTimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetManyTimesRequest) Reset() {
	*x = GreetManyTimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_v1_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesRequest) ProtoMessage() {}

func (x *GreetManyTimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_v1_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesRequest.ProtoReflect.Descriptor instead.
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_v1_greet_proto_rawDescGZIP(), []int{3}
}

func (x *GreetManyTimesRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetManyTimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Count  uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GreetManyTimesResponse) Reset() {
	*x = GreetManyTimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_v1_greet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesResponse) ProtoMessage() {}

func (x *GreetManyTimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_v1_greet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesResponse.ProtoReflect.Descriptor instead.
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_v1_greet_proto_rawDescGZIP(), []int{4}
}

func (x *GreetManyTimesResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *GreetManyTimesResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type LongGreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *LongGreetRequest) Reset() {
	*x = LongGreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_v1_greet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongGreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongGreetRequest) ProtoMessage() {}

func (x *LongGreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_v1_greet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongGreetRequest.ProtoReflect.Descriptor instead.
func (*LongGreetRequest) Descriptor() ([]byte, []int) {
	return file_greetpb_v1_greet_proto_rawDescGZIP(), []int{5}
}

func (x *LongGreetRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type LongGreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LongGreetResponse) Reset() {
	*x = LongGreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetpb_v1_greet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongGreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongGreetResponse) ProtoMessage() {}

func (x *LongGreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetpb_v1_greet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongGreetResponse.ProtoReflect.Descriptor instead.
func (*LongGreetResponse) Descriptor() ([]byte, []int) {
	return file_greetpb_v1_greet_proto_rawDescGZIP(), []int{6}
}

func (x *LongGreetResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_greetpb_v1_greet_proto protoreflect.FileDescriptor

var file_greetpb_v1_greet_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0c, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x27, 0x0a, 0x0d,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61,
	0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x46, 0x0a, 0x16, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x10, 0x4c, 0x6f, 0x6e, 0x67,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2b,
	0x0a, 0x11, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xbc, 0x04, 0x0a, 0x0c,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a,
	0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x92, 0x41, 0x40,
	0x12, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x20, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x30,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x20, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0xc8, 0x01, 0x0a, 0x0e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61,
	0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6d, 0x92, 0x41, 0x4b, 0x12, 0x17, 0x47, 0x72, 0x65, 0x65, 0x74, 0x20, 0x43,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x20, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x6c, 0x79,
	0x1a, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61,
	0x20, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x20, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3a, 0x01,
	0x2a, 0x30, 0x01, 0x12, 0xc4, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6e,
	0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x78,
	0x92, 0x41, 0x58, 0x12, 0x24, 0x47, 0x72, 0x65, 0x65, 0x74, 0x20, 0x43, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x20, 0x41, 0x66, 0x74, 0x65, 0x72, 0x20, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x6c, 0x79, 0x20, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x28, 0x01, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x65, 0x6e, 0x69, 0x78, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greetpb_v1_greet_proto_rawDescOnce sync.Once
	file_greetpb_v1_greet_proto_rawDescData = file_greetpb_v1_greet_proto_rawDesc
)

func file_greetpb_v1_greet_proto_rawDescGZIP() []byte {
	file_greetpb_v1_greet_proto_rawDescOnce.Do(func() {
		file_greetpb_v1_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greetpb_v1_greet_proto_rawDescData)
	})
	return file_greetpb_v1_greet_proto_rawDescData
}

var file_greetpb_v1_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_greetpb_v1_greet_proto_goTypes = []interface{}{
	(*Greeting)(nil),               // 0: greetpb.v1.Greeting
	(*GreetRequest)(nil),           // 1: greetpb.v1.GreetRequest
	(*GreetResponse)(nil),          // 2: greetpb.v1.GreetResponse
	(*GreetManyTimesRequest)(nil),  // 3: greetpb.v1.GreetManyTimesRequest
	(*GreetManyTimesResponse)(nil), // 4: greetpb.v1.GreetManyTimesResponse
	(*LongGreetRequest)(nil),       // 5: greetpb.v1.LongGreetRequest
	(*LongGreetResponse)(nil),      // 6: greetpb.v1.LongGreetResponse
}
var file_greetpb_v1_greet_proto_depIdxs = []int32{
	0, // 0: greetpb.v1.GreetRequest.greeting:type_name -> greetpb.v1.Greeting
	0, // 1: greetpb.v1.GreetManyTimesRequest.greeting:type_name -> greetpb.v1.Greeting
	0, // 2: greetpb.v1.LongGreetRequest.greeting:type_name -> greetpb.v1.Greeting
	1, // 3: greetpb.v1.GreetService.Greet:input_type -> greetpb.v1.GreetRequest
	3, // 4: greetpb.v1.GreetService.GreetManyTimes:input_type -> greetpb.v1.GreetManyTimesRequest
	5, // 5: greetpb.v1.GreetService.LongGreet:input_type -> greetpb.v1.LongGreetRequest
	2, // 6: greetpb.v1.GreetService.Greet:output_type -> greetpb.v1.GreetResponse
	4, // 7: greetpb.v1.GreetService.GreetManyTimes:output_type -> greetpb.v1.GreetManyTimesResponse
	6, // 8: greetpb.v1.GreetService.LongGreet:output_type -> greetpb.v1.LongGreetResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_greetpb_v1_greet_proto_init() }
func file_greetpb_v1_greet_proto_init() {
	if File_greetpb_v1_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greetpb_v1_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeting); i {
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
		file_greetpb_v1_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_greetpb_v1_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_greetpb_v1_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesRequest); i {
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
		file_greetpb_v1_greet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesResponse); i {
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
		file_greetpb_v1_greet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongGreetRequest); i {
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
		file_greetpb_v1_greet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongGreetResponse); i {
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
			RawDescriptor: file_greetpb_v1_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greetpb_v1_greet_proto_goTypes,
		DependencyIndexes: file_greetpb_v1_greet_proto_depIdxs,
		MessageInfos:      file_greetpb_v1_greet_proto_msgTypes,
	}.Build()
	File_greetpb_v1_greet_proto = out.File
	file_greetpb_v1_greet_proto_rawDesc = nil
	file_greetpb_v1_greet_proto_goTypes = nil
	file_greetpb_v1_greet_proto_depIdxs = nil
}
