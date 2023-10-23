// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: order.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Delivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone   string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Zip     string `protobuf:"bytes,3,opt,name=zip,proto3" json:"zip,omitempty"`
	City    string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Region  string `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	Email   string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Delivery) Reset() {
	*x = Delivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delivery) ProtoMessage() {}

func (x *Delivery) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delivery.ProtoReflect.Descriptor instead.
func (*Delivery) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Delivery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Delivery) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Delivery) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *Delivery) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Delivery) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Delivery) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Delivery) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction  string                 `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	RequestId    string                 `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Currency     string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Provider     string                 `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	Amount       int32                  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	PaymentDt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=payment_dt,json=paymentDt,proto3" json:"payment_dt,omitempty"`
	Bank         string                 `protobuf:"bytes,7,opt,name=bank,proto3" json:"bank,omitempty"`
	DeliveryCost int32                  `protobuf:"varint,8,opt,name=delivery_cost,json=deliveryCost,proto3" json:"delivery_cost,omitempty"`
	GoodsTotal   int32                  `protobuf:"varint,9,opt,name=goods_total,json=goodsTotal,proto3" json:"goods_total,omitempty"`
	CustomFee    int32                  `protobuf:"varint,10,opt,name=custom_fee,json=customFee,proto3" json:"custom_fee,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *Payment) GetTransaction() string {
	if x != nil {
		return x.Transaction
	}
	return ""
}

func (x *Payment) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Payment) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Payment) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Payment) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Payment) GetPaymentDt() *timestamppb.Timestamp {
	if x != nil {
		return x.PaymentDt
	}
	return nil
}

func (x *Payment) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *Payment) GetDeliveryCost() int32 {
	if x != nil {
		return x.DeliveryCost
	}
	return 0
}

func (x *Payment) GetGoodsTotal() int32 {
	if x != nil {
		return x.GoodsTotal
	}
	return 0
}

func (x *Payment) GetCustomFee() int32 {
	if x != nil {
		return x.CustomFee
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChrtId      int32   `protobuf:"varint,1,opt,name=chrt_id,json=chrtId,proto3" json:"chrt_id,omitempty"`
	TrackNumber string  `protobuf:"bytes,2,opt,name=track_number,json=trackNumber,proto3" json:"track_number,omitempty"`
	Price       float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Rid         string  `protobuf:"bytes,4,opt,name=rid,proto3" json:"rid,omitempty"`
	Name        string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Sale        float32 `protobuf:"fixed32,6,opt,name=sale,proto3" json:"sale,omitempty"`
	Size        string  `protobuf:"bytes,7,opt,name=size,proto3" json:"size,omitempty"`
	TotalPrice  float32 `protobuf:"fixed32,8,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	NmId        int32   `protobuf:"varint,9,opt,name=nm_id,json=nmId,proto3" json:"nm_id,omitempty"`
	Brand       string  `protobuf:"bytes,10,opt,name=brand,proto3" json:"brand,omitempty"`
	Status      int32   `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetChrtId() int32 {
	if x != nil {
		return x.ChrtId
	}
	return 0
}

func (x *Item) GetTrackNumber() string {
	if x != nil {
		return x.TrackNumber
	}
	return ""
}

func (x *Item) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Item) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetSale() float32 {
	if x != nil {
		return x.Sale
	}
	return 0
}

func (x *Item) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *Item) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Item) GetNmId() int32 {
	if x != nil {
		return x.NmId
	}
	return 0
}

func (x *Item) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Item) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackNumber     string                 `protobuf:"bytes,1,opt,name=TrackNumber,proto3" json:"TrackNumber,omitempty"`
	Entry           string                 `protobuf:"bytes,2,opt,name=Entry,proto3" json:"Entry,omitempty"`
	Locale          string                 `protobuf:"bytes,3,opt,name=Locale,proto3" json:"Locale,omitempty"`
	InternalSig     string                 `protobuf:"bytes,4,opt,name=InternalSig,proto3" json:"InternalSig,omitempty"`
	CustomerID      string                 `protobuf:"bytes,5,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	DeliveryService string                 `protobuf:"bytes,6,opt,name=DeliveryService,proto3" json:"DeliveryService,omitempty"`
	ShardKey        string                 `protobuf:"bytes,7,opt,name=ShardKey,proto3" json:"ShardKey,omitempty"`
	SMID            int32                  `protobuf:"varint,8,opt,name=SMID,proto3" json:"SMID,omitempty"`
	OofShard        string                 `protobuf:"bytes,9,opt,name=OofShard,proto3" json:"OofShard,omitempty"`
	DateCreated     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=DateCreated,proto3" json:"DateCreated,omitempty"`
	Delivery        *Delivery              `protobuf:"bytes,11,opt,name=Delivery,proto3" json:"Delivery,omitempty"`
	Payment         *Payment               `protobuf:"bytes,12,opt,name=Payment,proto3" json:"Payment,omitempty"`
	Items           []*Item                `protobuf:"bytes,13,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetTrackNumber() string {
	if x != nil {
		return x.TrackNumber
	}
	return ""
}

func (x *Order) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

func (x *Order) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Order) GetInternalSig() string {
	if x != nil {
		return x.InternalSig
	}
	return ""
}

func (x *Order) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *Order) GetDeliveryService() string {
	if x != nil {
		return x.DeliveryService
	}
	return ""
}

func (x *Order) GetShardKey() string {
	if x != nil {
		return x.ShardKey
	}
	return ""
}

func (x *Order) GetSMID() int32 {
	if x != nil {
		return x.SMID
	}
	return 0
}

func (x *Order) GetOofShard() string {
	if x != nil {
		return x.OofShard
	}
	return ""
}

func (x *Order) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

func (x *Order) GetDelivery() *Delivery {
	if x != nil {
		return x.Delivery
	}
	return nil
}

func (x *Order) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x7a, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xce, 0x02, 0x0a, 0x07, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x22, 0x8a, 0x02, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x72, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x61, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc7, 0x03, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x53, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x4d,
	0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x4d, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x4f, 0x6f, 0x66, 0x53, 0x68, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4f, 0x6f, 0x66, 0x53, 0x68, 0x61, 0x72, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x44, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x08, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x6d, 0x79, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_order_proto_goTypes = []interface{}{
	(*Delivery)(nil),              // 0: order.Delivery
	(*Payment)(nil),               // 1: order.Payment
	(*Item)(nil),                  // 2: order.Item
	(*Order)(nil),                 // 3: order.Order
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_order_proto_depIdxs = []int32{
	4, // 0: order.Payment.payment_dt:type_name -> google.protobuf.Timestamp
	4, // 1: order.Order.DateCreated:type_name -> google.protobuf.Timestamp
	0, // 2: order.Order.Delivery:type_name -> order.Delivery
	1, // 3: order.Order.Payment:type_name -> order.Payment
	2, // 4: order.Order.Items:type_name -> order.Item
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delivery); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
