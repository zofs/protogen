// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: catalog/v1/showcase_service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Showcase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId string `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Showcase) Reset() {
	*x = Showcase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_showcase_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Showcase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Showcase) ProtoMessage() {}

func (x *Showcase) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_showcase_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Showcase.ProtoReflect.Descriptor instead.
func (*Showcase) Descriptor() ([]byte, []int) {
	return file_catalog_v1_showcase_service_proto_rawDescGZIP(), []int{0}
}

func (x *Showcase) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Showcase) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Showcase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ShowcaseCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId string `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ShowcaseCreate) Reset() {
	*x = ShowcaseCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_showcase_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowcaseCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowcaseCreate) ProtoMessage() {}

func (x *ShowcaseCreate) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_showcase_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowcaseCreate.ProtoReflect.Descriptor instead.
func (*ShowcaseCreate) Descriptor() ([]byte, []int) {
	return file_catalog_v1_showcase_service_proto_rawDescGZIP(), []int{1}
}

func (x *ShowcaseCreate) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *ShowcaseCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ShowcaseUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId string `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ShowcaseUpdate) Reset() {
	*x = ShowcaseUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_showcase_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowcaseUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowcaseUpdate) ProtoMessage() {}

func (x *ShowcaseUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_showcase_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowcaseUpdate.ProtoReflect.Descriptor instead.
func (*ShowcaseUpdate) Descriptor() ([]byte, []int) {
	return file_catalog_v1_showcase_service_proto_rawDescGZIP(), []int{2}
}

func (x *ShowcaseUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShowcaseUpdate) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *ShowcaseUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ShowcaseDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId string `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
}

func (x *ShowcaseDelete) Reset() {
	*x = ShowcaseDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_showcase_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowcaseDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowcaseDelete) ProtoMessage() {}

func (x *ShowcaseDelete) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_showcase_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowcaseDelete.ProtoReflect.Descriptor instead.
func (*ShowcaseDelete) Descriptor() ([]byte, []int) {
	return file_catalog_v1_showcase_service_proto_rawDescGZIP(), []int{3}
}

func (x *ShowcaseDelete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShowcaseDelete) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

type ListShowcase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId   string      `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Showcases []*Showcase `protobuf:"bytes,2,rep,name=showcases,proto3" json:"showcases,omitempty"`
}

func (x *ListShowcase) Reset() {
	*x = ListShowcase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_showcase_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShowcase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShowcase) ProtoMessage() {}

func (x *ListShowcase) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_showcase_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShowcase.ProtoReflect.Descriptor instead.
func (*ListShowcase) Descriptor() ([]byte, []int) {
	return file_catalog_v1_showcase_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListShowcase) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *ListShowcase) GetShowcases() []*Showcase {
	if x != nil {
		return x.Showcases
	}
	return nil
}

var File_catalog_v1_showcase_service_proto protoreflect.FileDescriptor

var file_catalog_v1_showcase_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x08, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0e, 0x53, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x0e, 0x53, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x35, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x52, 0x09,
	0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x73, 0x32, 0xac, 0x04, 0x0a, 0x0f, 0x53, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x22, 0x41, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x7b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x17, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x22, 0x46, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a,
	0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a,
	0x01, 0x2a, 0x1a, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x7b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x80,
	0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x22, 0x3e, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x7b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x12, 0x84, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x43, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x2a, 0x2a, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x7b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xd3, 0x01, 0x92, 0x41, 0xa7, 0x01, 0x12,
	0x4a, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x20, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2a, 0x31, 0x0a, 0x0b, 0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x72, 0x6d, 0x73, 0x75,
	0x62, 0x65, 0x6b, 0x74, 0x69, 0x2e, 0x6d, 0x69, 0x74, 0x2d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x5a, 0x59, 0x0a, 0x57, 0x0a,
	0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x6f, 0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalog_v1_showcase_service_proto_rawDescOnce sync.Once
	file_catalog_v1_showcase_service_proto_rawDescData = file_catalog_v1_showcase_service_proto_rawDesc
)

func file_catalog_v1_showcase_service_proto_rawDescGZIP() []byte {
	file_catalog_v1_showcase_service_proto_rawDescOnce.Do(func() {
		file_catalog_v1_showcase_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_v1_showcase_service_proto_rawDescData)
	})
	return file_catalog_v1_showcase_service_proto_rawDescData
}

var file_catalog_v1_showcase_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_catalog_v1_showcase_service_proto_goTypes = []interface{}{
	(*Showcase)(nil),       // 0: catalog.v1.pb.Showcase
	(*ShowcaseCreate)(nil), // 1: catalog.v1.pb.ShowcaseCreate
	(*ShowcaseUpdate)(nil), // 2: catalog.v1.pb.ShowcaseUpdate
	(*ShowcaseDelete)(nil), // 3: catalog.v1.pb.ShowcaseDelete
	(*ListShowcase)(nil),   // 4: catalog.v1.pb.ListShowcase
	(*empty.Empty)(nil),    // 5: google.protobuf.Empty
}
var file_catalog_v1_showcase_service_proto_depIdxs = []int32{
	0, // 0: catalog.v1.pb.ListShowcase.showcases:type_name -> catalog.v1.pb.Showcase
	1, // 1: catalog.v1.pb.ShowcaseService.Create:input_type -> catalog.v1.pb.ShowcaseCreate
	2, // 2: catalog.v1.pb.ShowcaseService.Update:input_type -> catalog.v1.pb.ShowcaseUpdate
	4, // 3: catalog.v1.pb.ShowcaseService.List:input_type -> catalog.v1.pb.ListShowcase
	3, // 4: catalog.v1.pb.ShowcaseService.Delete:input_type -> catalog.v1.pb.ShowcaseDelete
	0, // 5: catalog.v1.pb.ShowcaseService.Create:output_type -> catalog.v1.pb.Showcase
	0, // 6: catalog.v1.pb.ShowcaseService.Update:output_type -> catalog.v1.pb.Showcase
	4, // 7: catalog.v1.pb.ShowcaseService.List:output_type -> catalog.v1.pb.ListShowcase
	5, // 8: catalog.v1.pb.ShowcaseService.Delete:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_catalog_v1_showcase_service_proto_init() }
func file_catalog_v1_showcase_service_proto_init() {
	if File_catalog_v1_showcase_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_v1_showcase_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Showcase); i {
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
		file_catalog_v1_showcase_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowcaseCreate); i {
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
		file_catalog_v1_showcase_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowcaseUpdate); i {
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
		file_catalog_v1_showcase_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowcaseDelete); i {
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
		file_catalog_v1_showcase_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShowcase); i {
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
			RawDescriptor: file_catalog_v1_showcase_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_v1_showcase_service_proto_goTypes,
		DependencyIndexes: file_catalog_v1_showcase_service_proto_depIdxs,
		MessageInfos:      file_catalog_v1_showcase_service_proto_msgTypes,
	}.Build()
	File_catalog_v1_showcase_service_proto = out.File
	file_catalog_v1_showcase_service_proto_rawDesc = nil
	file_catalog_v1_showcase_service_proto_goTypes = nil
	file_catalog_v1_showcase_service_proto_depIdxs = nil
}
