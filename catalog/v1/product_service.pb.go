// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: catalog/v1/product_service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	dtopb "github.com/zofs/protogen/dtopb"
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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId     string                `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ShowcaseId  *wrappers.StringValue `protobuf:"bytes,3,opt,name=showcase_id,json=showcaseId,proto3" json:"showcase_id,omitempty"`
	CategoryId  *wrappers.StringValue `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name        string                `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	MinQty      uint32                `protobuf:"varint,6,opt,name=min_qty,json=minQty,proto3" json:"min_qty,omitempty"`
	Description string                `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamp.Timestamp  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   *timestamp.Timestamp  `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_product_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_product_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_catalog_v1_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Product) GetShowcaseId() *wrappers.StringValue {
	if x != nil {
		return x.ShowcaseId
	}
	return nil
}

func (x *Product) GetCategoryId() *wrappers.StringValue {
	if x != nil {
		return x.CategoryId
	}
	return nil
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetMinQty() uint32 {
	if x != nil {
		return x.MinQty
	}
	return 0
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Product) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Product) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type ProductCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId     string                `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ShowcaseId  *wrappers.StringValue `protobuf:"bytes,2,opt,name=showcase_id,json=showcaseId,proto3" json:"showcase_id,omitempty"`
	CategoryId  *wrappers.StringValue `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name        string                `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	MinQty      uint32                `protobuf:"varint,5,opt,name=min_qty,json=minQty,proto3" json:"min_qty,omitempty"`
	Description string                `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ProductCreate) Reset() {
	*x = ProductCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_product_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreate) ProtoMessage() {}

func (x *ProductCreate) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_product_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreate.ProtoReflect.Descriptor instead.
func (*ProductCreate) Descriptor() ([]byte, []int) {
	return file_catalog_v1_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProductCreate) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *ProductCreate) GetShowcaseId() *wrappers.StringValue {
	if x != nil {
		return x.ShowcaseId
	}
	return nil
}

func (x *ProductCreate) GetCategoryId() *wrappers.StringValue {
	if x != nil {
		return x.CategoryId
	}
	return nil
}

func (x *ProductCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCreate) GetMinQty() uint32 {
	if x != nil {
		return x.MinQty
	}
	return 0
}

func (x *ProductCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ProductUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId     string                `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ShowcaseId  *wrappers.StringValue `protobuf:"bytes,3,opt,name=showcase_id,json=showcaseId,proto3" json:"showcase_id,omitempty"`
	CategoryId  *wrappers.StringValue `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name        *wrappers.StringValue `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	MinQty      *wrappers.Int32Value  `protobuf:"bytes,6,opt,name=min_qty,json=minQty,proto3" json:"min_qty,omitempty"`
	Description *wrappers.StringValue `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ProductUpdate) Reset() {
	*x = ProductUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_product_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdate) ProtoMessage() {}

func (x *ProductUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_product_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdate.ProtoReflect.Descriptor instead.
func (*ProductUpdate) Descriptor() ([]byte, []int) {
	return file_catalog_v1_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProductUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductUpdate) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *ProductUpdate) GetShowcaseId() *wrappers.StringValue {
	if x != nil {
		return x.ShowcaseId
	}
	return nil
}

func (x *ProductUpdate) GetCategoryId() *wrappers.StringValue {
	if x != nil {
		return x.CategoryId
	}
	return nil
}

func (x *ProductUpdate) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ProductUpdate) GetMinQty() *wrappers.Int32Value {
	if x != nil {
		return x.MinQty
	}
	return nil
}

func (x *ProductUpdate) GetDescription() *wrappers.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

type ListProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId    *wrappers.StringValue `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ShowcaseId *wrappers.StringValue `protobuf:"bytes,2,opt,name=showcase_id,json=showcaseId,proto3" json:"showcase_id,omitempty"`
	CategoryId *wrappers.StringValue `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Paginate   *dtopb.Pagination     `protobuf:"bytes,4,opt,name=paginate,proto3" json:"paginate,omitempty"`
	Products   []*Product            `protobuf:"bytes,5,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ListProduct) Reset() {
	*x = ListProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_product_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProduct) ProtoMessage() {}

func (x *ListProduct) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_product_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProduct.ProtoReflect.Descriptor instead.
func (*ListProduct) Descriptor() ([]byte, []int) {
	return file_catalog_v1_product_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListProduct) GetStoreId() *wrappers.StringValue {
	if x != nil {
		return x.StoreId
	}
	return nil
}

func (x *ListProduct) GetShowcaseId() *wrappers.StringValue {
	if x != nil {
		return x.ShowcaseId
	}
	return nil
}

func (x *ListProduct) GetCategoryId() *wrappers.StringValue {
	if x != nil {
		return x.CategoryId
	}
	return nil
}

func (x *ListProduct) GetPaginate() *dtopb.Pagination {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *ListProduct) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_catalog_v1_product_service_proto protoreflect.FileDescriptor

var file_catalog_v1_product_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x64,
	0x74, 0x6f, 0x70, 0x62, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x64,
	0x74, 0x6f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b,
	0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6d, 0x69, 0x6e, 0x51, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf7, 0x01, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e,
	0x5f, 0x71, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x51,
	0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe0, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49,
	0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x6d, 0x69, 0x6e, 0x51, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa7, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x3d, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2d,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x64, 0x74, 0x6f, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x32, 0xae, 0x04, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x2f, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a,
	0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x74, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x34, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01,
	0x2a, 0x1a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a,
	0x1a, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x2c, 0x92, 0x41, 0x0e,
	0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x61, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x0f, 0x2e, 0x64, 0x74, 0x6f, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x31, 0x92, 0x41, 0x0e, 0x62, 0x0c,
	0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x64, 0x74, 0x6f, 0x70, 0x62, 0x2e, 0x49,
	0x44, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x31, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x42, 0xd2, 0x01, 0x92, 0x41, 0xa6, 0x01, 0x12, 0x49, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2a, 0x31, 0x0a, 0x0b,
	0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x72, 0x6d, 0x73, 0x75, 0x62, 0x65, 0x6b, 0x74, 0x69, 0x2e, 0x6d,
	0x69, 0x74, 0x2d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6f, 0x66, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalog_v1_product_service_proto_rawDescOnce sync.Once
	file_catalog_v1_product_service_proto_rawDescData = file_catalog_v1_product_service_proto_rawDesc
)

func file_catalog_v1_product_service_proto_rawDescGZIP() []byte {
	file_catalog_v1_product_service_proto_rawDescOnce.Do(func() {
		file_catalog_v1_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_v1_product_service_proto_rawDescData)
	})
	return file_catalog_v1_product_service_proto_rawDescData
}

var file_catalog_v1_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_catalog_v1_product_service_proto_goTypes = []interface{}{
	(*Product)(nil),              // 0: catalog.v1.pb.Product
	(*ProductCreate)(nil),        // 1: catalog.v1.pb.ProductCreate
	(*ProductUpdate)(nil),        // 2: catalog.v1.pb.ProductUpdate
	(*ListProduct)(nil),          // 3: catalog.v1.pb.ListProduct
	(*wrappers.StringValue)(nil), // 4: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*wrappers.Int32Value)(nil),  // 6: google.protobuf.Int32Value
	(*dtopb.Pagination)(nil),     // 7: dtopb.Pagination
	(*dtopb.IDString)(nil),       // 8: dtopb.IDString
	(*empty.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_catalog_v1_product_service_proto_depIdxs = []int32{
	4,  // 0: catalog.v1.pb.Product.showcase_id:type_name -> google.protobuf.StringValue
	4,  // 1: catalog.v1.pb.Product.category_id:type_name -> google.protobuf.StringValue
	5,  // 2: catalog.v1.pb.Product.created_at:type_name -> google.protobuf.Timestamp
	5,  // 3: catalog.v1.pb.Product.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 4: catalog.v1.pb.Product.deleted_at:type_name -> google.protobuf.Timestamp
	4,  // 5: catalog.v1.pb.ProductCreate.showcase_id:type_name -> google.protobuf.StringValue
	4,  // 6: catalog.v1.pb.ProductCreate.category_id:type_name -> google.protobuf.StringValue
	4,  // 7: catalog.v1.pb.ProductUpdate.showcase_id:type_name -> google.protobuf.StringValue
	4,  // 8: catalog.v1.pb.ProductUpdate.category_id:type_name -> google.protobuf.StringValue
	4,  // 9: catalog.v1.pb.ProductUpdate.name:type_name -> google.protobuf.StringValue
	6,  // 10: catalog.v1.pb.ProductUpdate.min_qty:type_name -> google.protobuf.Int32Value
	4,  // 11: catalog.v1.pb.ProductUpdate.description:type_name -> google.protobuf.StringValue
	4,  // 12: catalog.v1.pb.ListProduct.store_id:type_name -> google.protobuf.StringValue
	4,  // 13: catalog.v1.pb.ListProduct.showcase_id:type_name -> google.protobuf.StringValue
	4,  // 14: catalog.v1.pb.ListProduct.category_id:type_name -> google.protobuf.StringValue
	7,  // 15: catalog.v1.pb.ListProduct.paginate:type_name -> dtopb.Pagination
	0,  // 16: catalog.v1.pb.ListProduct.products:type_name -> catalog.v1.pb.Product
	1,  // 17: catalog.v1.pb.ProductService.Create:input_type -> catalog.v1.pb.ProductCreate
	2,  // 18: catalog.v1.pb.ProductService.Update:input_type -> catalog.v1.pb.ProductUpdate
	3,  // 19: catalog.v1.pb.ProductService.List:input_type -> catalog.v1.pb.ListProduct
	8,  // 20: catalog.v1.pb.ProductService.Get:input_type -> dtopb.IDString
	8,  // 21: catalog.v1.pb.ProductService.Delete:input_type -> dtopb.IDString
	0,  // 22: catalog.v1.pb.ProductService.Create:output_type -> catalog.v1.pb.Product
	0,  // 23: catalog.v1.pb.ProductService.Update:output_type -> catalog.v1.pb.Product
	3,  // 24: catalog.v1.pb.ProductService.List:output_type -> catalog.v1.pb.ListProduct
	0,  // 25: catalog.v1.pb.ProductService.Get:output_type -> catalog.v1.pb.Product
	9,  // 26: catalog.v1.pb.ProductService.Delete:output_type -> google.protobuf.Empty
	22, // [22:27] is the sub-list for method output_type
	17, // [17:22] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_catalog_v1_product_service_proto_init() }
func file_catalog_v1_product_service_proto_init() {
	if File_catalog_v1_product_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_v1_product_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_catalog_v1_product_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreate); i {
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
		file_catalog_v1_product_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductUpdate); i {
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
		file_catalog_v1_product_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProduct); i {
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
			RawDescriptor: file_catalog_v1_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_v1_product_service_proto_goTypes,
		DependencyIndexes: file_catalog_v1_product_service_proto_depIdxs,
		MessageInfos:      file_catalog_v1_product_service_proto_msgTypes,
	}.Build()
	File_catalog_v1_product_service_proto = out.File
	file_catalog_v1_product_service_proto_rawDesc = nil
	file_catalog_v1_product_service_proto_goTypes = nil
	file_catalog_v1_product_service_proto_depIdxs = nil
}
