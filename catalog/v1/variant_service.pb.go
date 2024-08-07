// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: catalog/v1/variant_service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Sku       string  `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Name      string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Picture   string  `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	Price     float32 `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Variant) Reset() {
	*x = Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_variant_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant) ProtoMessage() {}

func (x *Variant) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_variant_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant.ProtoReflect.Descriptor instead.
func (*Variant) Descriptor() ([]byte, []int) {
	return file_catalog_v1_variant_service_proto_rawDescGZIP(), []int{0}
}

func (x *Variant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Variant) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Variant) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *Variant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Variant) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Variant) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type VariantCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string  `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Sku       string  `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Name      string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Picture   string  `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	Price     float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *VariantCreate) Reset() {
	*x = VariantCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_variant_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariantCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantCreate) ProtoMessage() {}

func (x *VariantCreate) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_variant_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantCreate.ProtoReflect.Descriptor instead.
func (*VariantCreate) Descriptor() ([]byte, []int) {
	return file_catalog_v1_variant_service_proto_rawDescGZIP(), []int{1}
}

func (x *VariantCreate) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *VariantCreate) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *VariantCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VariantCreate) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *VariantCreate) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type VariantDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *VariantDelete) Reset() {
	*x = VariantDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_variant_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariantDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantDelete) ProtoMessage() {}

func (x *VariantDelete) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_variant_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantDelete.ProtoReflect.Descriptor instead.
func (*VariantDelete) Descriptor() ([]byte, []int) {
	return file_catalog_v1_variant_service_proto_rawDescGZIP(), []int{2}
}

func (x *VariantDelete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VariantDelete) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type VariantUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string                `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Sku       *wrappers.StringValue `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Name      *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Picture   *wrappers.StringValue `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	Price     *wrappers.FloatValue  `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *VariantUpdate) Reset() {
	*x = VariantUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_variant_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariantUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantUpdate) ProtoMessage() {}

func (x *VariantUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_variant_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantUpdate.ProtoReflect.Descriptor instead.
func (*VariantUpdate) Descriptor() ([]byte, []int) {
	return file_catalog_v1_variant_service_proto_rawDescGZIP(), []int{3}
}

func (x *VariantUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VariantUpdate) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *VariantUpdate) GetSku() *wrappers.StringValue {
	if x != nil {
		return x.Sku
	}
	return nil
}

func (x *VariantUpdate) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *VariantUpdate) GetPicture() *wrappers.StringValue {
	if x != nil {
		return x.Picture
	}
	return nil
}

func (x *VariantUpdate) GetPrice() *wrappers.FloatValue {
	if x != nil {
		return x.Price
	}
	return nil
}

type ListVariant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variants  []*Variant `protobuf:"bytes,1,rep,name=variants,proto3" json:"variants,omitempty"`
	ProductId string     `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *ListVariant) Reset() {
	*x = ListVariant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_variant_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVariant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVariant) ProtoMessage() {}

func (x *ListVariant) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_variant_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVariant.ProtoReflect.Descriptor instead.
func (*ListVariant) Descriptor() ([]byte, []int) {
	return file_catalog_v1_variant_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListVariant) GetVariants() []*Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *ListVariant) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

var File_catalog_v1_variant_service_proto protoreflect.FileDescriptor

var file_catalog_v1_variant_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a,
	0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x84, 0x01,
	0x0a, 0x0d, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x0d, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x8b, 0x02, 0x0a, 0x0d, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x31, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x60, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x32, 0xac, 0x03, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x1a, 0x16, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62,
	0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22, 0x44, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a,
	0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x89,
	0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x22,
	0x49, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x3a, 0x01, 0x2a, 0x1a, 0x2d, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x86, 0x01, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x46, 0x92, 0x41, 0x0e,
	0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2f, 0x2a, 0x2d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0xd2, 0x01, 0x92, 0x41, 0xa6, 0x01, 0x12, 0x49, 0x0a, 0x0f, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2a, 0x31, 0x0a,
	0x0b, 0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x72, 0x6d, 0x73, 0x75, 0x62, 0x65, 0x6b, 0x74, 0x69, 0x2e,
	0x6d, 0x69, 0x74, 0x2d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x32, 0x03, 0x31, 0x2e, 0x30, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a,
	0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6f, 0x66,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalog_v1_variant_service_proto_rawDescOnce sync.Once
	file_catalog_v1_variant_service_proto_rawDescData = file_catalog_v1_variant_service_proto_rawDesc
)

func file_catalog_v1_variant_service_proto_rawDescGZIP() []byte {
	file_catalog_v1_variant_service_proto_rawDescOnce.Do(func() {
		file_catalog_v1_variant_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_v1_variant_service_proto_rawDescData)
	})
	return file_catalog_v1_variant_service_proto_rawDescData
}

var file_catalog_v1_variant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_catalog_v1_variant_service_proto_goTypes = []interface{}{
	(*Variant)(nil),              // 0: catalog.v1.pb.Variant
	(*VariantCreate)(nil),        // 1: catalog.v1.pb.VariantCreate
	(*VariantDelete)(nil),        // 2: catalog.v1.pb.VariantDelete
	(*VariantUpdate)(nil),        // 3: catalog.v1.pb.VariantUpdate
	(*ListVariant)(nil),          // 4: catalog.v1.pb.ListVariant
	(*wrappers.StringValue)(nil), // 5: google.protobuf.StringValue
	(*wrappers.FloatValue)(nil),  // 6: google.protobuf.FloatValue
	(*empty.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_catalog_v1_variant_service_proto_depIdxs = []int32{
	5, // 0: catalog.v1.pb.VariantUpdate.sku:type_name -> google.protobuf.StringValue
	5, // 1: catalog.v1.pb.VariantUpdate.name:type_name -> google.protobuf.StringValue
	5, // 2: catalog.v1.pb.VariantUpdate.picture:type_name -> google.protobuf.StringValue
	6, // 3: catalog.v1.pb.VariantUpdate.price:type_name -> google.protobuf.FloatValue
	0, // 4: catalog.v1.pb.ListVariant.variants:type_name -> catalog.v1.pb.Variant
	1, // 5: catalog.v1.pb.VariantService.Create:input_type -> catalog.v1.pb.VariantCreate
	3, // 6: catalog.v1.pb.VariantService.Update:input_type -> catalog.v1.pb.VariantUpdate
	2, // 7: catalog.v1.pb.VariantService.Delete:input_type -> catalog.v1.pb.VariantDelete
	0, // 8: catalog.v1.pb.VariantService.Create:output_type -> catalog.v1.pb.Variant
	0, // 9: catalog.v1.pb.VariantService.Update:output_type -> catalog.v1.pb.Variant
	7, // 10: catalog.v1.pb.VariantService.Delete:output_type -> google.protobuf.Empty
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_catalog_v1_variant_service_proto_init() }
func file_catalog_v1_variant_service_proto_init() {
	if File_catalog_v1_variant_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_v1_variant_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variant); i {
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
		file_catalog_v1_variant_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariantCreate); i {
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
		file_catalog_v1_variant_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariantDelete); i {
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
		file_catalog_v1_variant_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariantUpdate); i {
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
		file_catalog_v1_variant_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVariant); i {
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
			RawDescriptor: file_catalog_v1_variant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_v1_variant_service_proto_goTypes,
		DependencyIndexes: file_catalog_v1_variant_service_proto_depIdxs,
		MessageInfos:      file_catalog_v1_variant_service_proto_msgTypes,
	}.Build()
	File_catalog_v1_variant_service_proto = out.File
	file_catalog_v1_variant_service_proto_rawDesc = nil
	file_catalog_v1_variant_service_proto_goTypes = nil
	file_catalog_v1_variant_service_proto_depIdxs = nil
}
