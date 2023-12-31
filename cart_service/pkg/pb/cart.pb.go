// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/proto/cart.proto

package pb

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

type CreateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int32 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *CreateCartRequest) Reset() {
	*x = CreateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartRequest) ProtoMessage() {}

func (x *CreateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartRequest.ProtoReflect.Descriptor instead.
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCartRequest) GetUserid() int32 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type CreateCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateCartResponse) Reset() {
	*x = CreateCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartResponse) ProtoMessage() {}

func (x *CreateCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartResponse.ProtoReflect.Descriptor instead.
func (*CreateCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCartResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ViewCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int32 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *ViewCartRequest) Reset() {
	*x = ViewCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartRequest) ProtoMessage() {}

func (x *ViewCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartRequest.ProtoReflect.Descriptor instead.
func (*ViewCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{2}
}

func (x *ViewCartRequest) GetUserid() int32 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type CartItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductDetailID string `protobuf:"bytes,1,opt,name=ProductDetailID,proto3" json:"ProductDetailID,omitempty"`
	Image           string `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
	ModelName       string `protobuf:"bytes,3,opt,name=ModelName,proto3" json:"ModelName,omitempty"`
	Price           uint32 `protobuf:"varint,4,opt,name=Price,proto3" json:"Price,omitempty"`
	Size            string `protobuf:"bytes,5,opt,name=Size,proto3" json:"Size,omitempty"`
	Colour          string `protobuf:"bytes,6,opt,name=Colour,proto3" json:"Colour,omitempty"`
	BrandName       string `protobuf:"bytes,7,opt,name=BrandName,proto3" json:"BrandName,omitempty"`
	Quantity        uint32 `protobuf:"varint,8,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Percentage      int32  `protobuf:"varint,9,opt,name=Percentage,proto3" json:"Percentage,omitempty"`
	Total           uint32 `protobuf:"varint,10,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *CartItems) Reset() {
	*x = CartItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItems) ProtoMessage() {}

func (x *CartItems) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItems.ProtoReflect.Descriptor instead.
func (*CartItems) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{3}
}

func (x *CartItems) GetProductDetailID() string {
	if x != nil {
		return x.ProductDetailID
	}
	return ""
}

func (x *CartItems) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CartItems) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *CartItems) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartItems) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *CartItems) GetColour() string {
	if x != nil {
		return x.Colour
	}
	return ""
}

func (x *CartItems) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *CartItems) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItems) GetPercentage() int32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *CartItems) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ViewCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     int32        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Grandtotal int32        `protobuf:"varint,2,opt,name=Grandtotal,proto3" json:"Grandtotal,omitempty"`
	Items      []*CartItems `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ViewCartResponse) Reset() {
	*x = ViewCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartResponse) ProtoMessage() {}

func (x *ViewCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartResponse.ProtoReflect.Descriptor instead.
func (*ViewCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{4}
}

func (x *ViewCartResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ViewCartResponse) GetGrandtotal() int32 {
	if x != nil {
		return x.Grandtotal
	}
	return 0
}

func (x *ViewCartResponse) GetItems() []*CartItems {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Productid string `protobuf:"bytes,1,opt,name=productid,proto3" json:"productid,omitempty"`
	Userid    uint32 `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *AddCartRequest) Reset() {
	*x = AddCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartRequest) ProtoMessage() {}

func (x *AddCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartRequest.ProtoReflect.Descriptor instead.
func (*AddCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{5}
}

func (x *AddCartRequest) GetProductid() string {
	if x != nil {
		return x.Productid
	}
	return ""
}

func (x *AddCartRequest) GetUserid() uint32 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type AddCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AddCartResponse) Reset() {
	*x = AddCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartResponse) ProtoMessage() {}

func (x *AddCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartResponse.ProtoReflect.Descriptor instead.
func (*AddCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{6}
}

func (x *AddCartResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddCartResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type CartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid uint32 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *CartRequest) Reset() {
	*x = CartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartRequest) ProtoMessage() {}

func (x *CartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartRequest.ProtoReflect.Descriptor instead.
func (*CartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{7}
}

func (x *CartRequest) GetUserid() uint32 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type CartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid     uint32 `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Grandtotal int32  `protobuf:"varint,3,opt,name=grandtotal,proto3" json:"grandtotal,omitempty"`
}

func (x *CartResponse) Reset() {
	*x = CartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartResponse) ProtoMessage() {}

func (x *CartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartResponse.ProtoReflect.Descriptor instead.
func (*CartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{8}
}

func (x *CartResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartResponse) GetUserid() uint32 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *CartResponse) GetGrandtotal() int32 {
	if x != nil {
		return x.Grandtotal
	}
	return 0
}

type CartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cartid uint32 `protobuf:"varint,1,opt,name=cartid,proto3" json:"cartid,omitempty"`
}

func (x *CartItemRequest) Reset() {
	*x = CartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemRequest) ProtoMessage() {}

func (x *CartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemRequest.ProtoReflect.Descriptor instead.
func (*CartItemRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{9}
}

func (x *CartItemRequest) GetCartid() uint32 {
	if x != nil {
		return x.Cartid
	}
	return 0
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartID          uint32 `protobuf:"varint,1,opt,name=CartID,proto3" json:"CartID,omitempty"`
	ProductDetailID uint32 `protobuf:"varint,2,opt,name=ProductDetailID,proto3" json:"ProductDetailID,omitempty"`
	Quantity        uint32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{10}
}

func (x *CartItem) GetCartID() uint32 {
	if x != nil {
		return x.CartID
	}
	return 0
}

func (x *CartItem) GetProductDetailID() uint32 {
	if x != nil {
		return x.ProductDetailID
	}
	return 0
}

func (x *CartItem) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CartItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CartItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CartItemResponse) Reset() {
	*x = CartItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemResponse) ProtoMessage() {}

func (x *CartItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemResponse.ProtoReflect.Descriptor instead.
func (*CartItemResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{11}
}

func (x *CartItemResponse) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grandtotal int32 `protobuf:"varint,1,opt,name=grandtotal,proto3" json:"grandtotal,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_cart_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cart_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cart_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateResponse) GetGrandtotal() int32 {
	if x != nil {
		return x.Grandtotal
	}
	return 0
}

var File_pkg_proto_cart_proto protoreflect.FileDescriptor

var file_pkg_proto_cart_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x2b, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x29, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x22, 0x9b, 0x02, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x28,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x6f,
	0x0a, 0x10, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x72,
	0x61, 0x6e, 0x64, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x47, 0x72, 0x61, 0x6e, 0x64, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x46, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25,
	0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x29, 0x0a,
	0x0f, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x36, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x30, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x67, 0x72, 0x61, 0x6e, 0x64, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x97, 0x03, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08,
	0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69,
	0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x74, 0x6f, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x08,
	0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x09, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_cart_proto_rawDescOnce sync.Once
	file_pkg_proto_cart_proto_rawDescData = file_pkg_proto_cart_proto_rawDesc
)

func file_pkg_proto_cart_proto_rawDescGZIP() []byte {
	file_pkg_proto_cart_proto_rawDescOnce.Do(func() {
		file_pkg_proto_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_cart_proto_rawDescData)
	})
	return file_pkg_proto_cart_proto_rawDescData
}

var file_pkg_proto_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_proto_cart_proto_goTypes = []interface{}{
	(*CreateCartRequest)(nil),  // 0: pb.CreateCartRequest
	(*CreateCartResponse)(nil), // 1: pb.CreateCartResponse
	(*ViewCartRequest)(nil),    // 2: pb.ViewCartRequest
	(*CartItems)(nil),          // 3: pb.CartItems
	(*ViewCartResponse)(nil),   // 4: pb.ViewCartResponse
	(*AddCartRequest)(nil),     // 5: pb.AddCartRequest
	(*AddCartResponse)(nil),    // 6: pb.AddCartResponse
	(*CartRequest)(nil),        // 7: pb.CartRequest
	(*CartResponse)(nil),       // 8: pb.CartResponse
	(*CartItemRequest)(nil),    // 9: pb.CartItemRequest
	(*CartItem)(nil),           // 10: pb.CartItem
	(*CartItemResponse)(nil),   // 11: pb.CartItemResponse
	(*UpdateResponse)(nil),     // 12: pb.UpdateResponse
}
var file_pkg_proto_cart_proto_depIdxs = []int32{
	3,  // 0: pb.ViewCartResponse.items:type_name -> pb.CartItems
	10, // 1: pb.CartItemResponse.items:type_name -> pb.CartItem
	2,  // 2: pb.CartService.ViewCart:input_type -> pb.ViewCartRequest
	5,  // 3: pb.CartService.AddtoCart:input_type -> pb.AddCartRequest
	0,  // 4: pb.CartService.CreateCart:input_type -> pb.CreateCartRequest
	7,  // 5: pb.CartService.FindCart:input_type -> pb.CartRequest
	9,  // 6: pb.CartService.CartItems:input_type -> pb.CartItemRequest
	8,  // 7: pb.CartService.UpdateCart:input_type -> pb.CartResponse
	9,  // 8: pb.CartService.DeleteCart:input_type -> pb.CartItemRequest
	4,  // 9: pb.CartService.ViewCart:output_type -> pb.ViewCartResponse
	6,  // 10: pb.CartService.AddtoCart:output_type -> pb.AddCartResponse
	1,  // 11: pb.CartService.CreateCart:output_type -> pb.CreateCartResponse
	8,  // 12: pb.CartService.FindCart:output_type -> pb.CartResponse
	11, // 13: pb.CartService.CartItems:output_type -> pb.CartItemResponse
	12, // 14: pb.CartService.UpdateCart:output_type -> pb.UpdateResponse
	12, // 15: pb.CartService.DeleteCart:output_type -> pb.UpdateResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_cart_proto_init() }
func file_pkg_proto_cart_proto_init() {
	if File_pkg_proto_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartRequest); i {
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
		file_pkg_proto_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartResponse); i {
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
		file_pkg_proto_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCartRequest); i {
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
		file_pkg_proto_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItems); i {
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
		file_pkg_proto_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCartResponse); i {
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
		file_pkg_proto_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartRequest); i {
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
		file_pkg_proto_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartResponse); i {
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
		file_pkg_proto_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartRequest); i {
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
		file_pkg_proto_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartResponse); i {
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
		file_pkg_proto_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemRequest); i {
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
		file_pkg_proto_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_pkg_proto_cart_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemResponse); i {
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
		file_pkg_proto_cart_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_pkg_proto_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_cart_proto_goTypes,
		DependencyIndexes: file_pkg_proto_cart_proto_depIdxs,
		MessageInfos:      file_pkg_proto_cart_proto_msgTypes,
	}.Build()
	File_pkg_proto_cart_proto = out.File
	file_pkg_proto_cart_proto_rawDesc = nil
	file_pkg_proto_cart_proto_goTypes = nil
	file_pkg_proto_cart_proto_depIdxs = nil
}
