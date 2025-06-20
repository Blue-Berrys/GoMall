// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.3
// source: checkout_page.proto

package checkout

import (
	_ "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/api"
	common "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/common"
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

type CheckoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" form:"email"`
	Firstname       string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty" form:"firstname"`
	Lastname        string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty" form:"lastname"`
	Street          string `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty" form:"street"`
	Zipcode         string `protobuf:"bytes,5,opt,name=zipcode,proto3" json:"zipcode,omitempty" form:"zipcode"`
	Province        string `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty" form:"province"`
	Country         string `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty" form:"country"`
	City            string `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty" form:"city"`
	CardNum         string `protobuf:"bytes,9,opt,name=card_num,json=cardNum,proto3" json:"card_num,omitempty" form:"cardNum"`
	ExpirationMonth int32  `protobuf:"varint,10,opt,name=expiration_month,json=expirationMonth,proto3" json:"expiration_month,omitempty" form:"expirationMonth"`
	Cvv             int32  `protobuf:"varint,11,opt,name=cvv,proto3" json:"cvv,omitempty" form:"cvv"`
	ExpirationYear  int32  `protobuf:"varint,12,opt,name=expiration_year,json=expirationYear,proto3" json:"expiration_year,omitempty" form:"expirationYear"`
	Payment         string `protobuf:"bytes,13,opt,name=payment,proto3" json:"payment,omitempty" form:"payment"`
}

func (x *CheckoutReq) Reset() {
	*x = CheckoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutReq) ProtoMessage() {}

func (x *CheckoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutReq.ProtoReflect.Descriptor instead.
func (*CheckoutReq) Descriptor() ([]byte, []int) {
	return file_checkout_page_proto_rawDescGZIP(), []int{0}
}

func (x *CheckoutReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CheckoutReq) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *CheckoutReq) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CheckoutReq) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *CheckoutReq) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

func (x *CheckoutReq) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *CheckoutReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CheckoutReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CheckoutReq) GetCardNum() string {
	if x != nil {
		return x.CardNum
	}
	return ""
}

func (x *CheckoutReq) GetExpirationMonth() int32 {
	if x != nil {
		return x.ExpirationMonth
	}
	return 0
}

func (x *CheckoutReq) GetCvv() int32 {
	if x != nil {
		return x.Cvv
	}
	return 0
}

func (x *CheckoutReq) GetExpirationYear() int32 {
	if x != nil {
		return x.ExpirationYear
	}
	return 0
}

func (x *CheckoutReq) GetPayment() string {
	if x != nil {
		return x.Payment
	}
	return ""
}

var File_checkout_page_proto protoreflect.FileDescriptor

var file_checkout_page_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x1a, 0x15, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x04, 0x0a, 0x0b, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xbb, 0x18, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d,
	0xe2, 0xbb, 0x18, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe2, 0xbb, 0x18, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xe2, 0xbb, 0x18, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0xbb, 0x18, 0x07, 0x7a, 0x69, 0x70,
	0x63, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe2, 0xbb, 0x18, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0xbb, 0x18, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xbb,
	0x18, 0x04, 0x63, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x08,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xe2, 0xbb, 0x18, 0x07, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x52, 0x07, 0x63, 0x61, 0x72,
	0x64, 0x4e, 0x75, 0x6d, 0x12, 0x3e, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x13,
	0xe2, 0xbb, 0x18, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xe2, 0xbb, 0x18, 0x03, 0x63, 0x76, 0x76, 0x52, 0x03, 0x63, 0x76, 0x76, 0x12,
	0x3b, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x42, 0x12, 0xe2, 0xbb, 0x18, 0x0e, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x52, 0x0e, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x12, 0x25, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2,
	0xbb, 0x18, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x32, 0x96, 0x02, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x0d, 0xca, 0xc1, 0x18, 0x09, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x12, 0x60, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x57, 0x61,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0xd2,
	0xc1, 0x18, 0x11, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x77, 0x61, 0x69,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x56, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0xca, 0xc1, 0x18, 0x10, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x48, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6c, 0x75, 0x65, 0x2d,
	0x42, 0x65, 0x72, 0x72, 0x79, 0x73, 0x2f, 0x47, 0x6f, 0x4d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkout_page_proto_rawDescOnce sync.Once
	file_checkout_page_proto_rawDescData = file_checkout_page_proto_rawDesc
)

func file_checkout_page_proto_rawDescGZIP() []byte {
	file_checkout_page_proto_rawDescOnce.Do(func() {
		file_checkout_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkout_page_proto_rawDescData)
	})
	return file_checkout_page_proto_rawDescData
}

var file_checkout_page_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_checkout_page_proto_goTypes = []interface{}{
	(*CheckoutReq)(nil),  // 0: frontend.checkout.CheckoutReq
	(*common.Empty)(nil), // 1: frontend.common.Empty
}
var file_checkout_page_proto_depIdxs = []int32{
	1, // 0: frontend.checkout.CheckoutService.Checkout:input_type -> frontend.common.Empty
	0, // 1: frontend.checkout.CheckoutService.CheckoutWaiting:input_type -> frontend.checkout.CheckoutReq
	1, // 2: frontend.checkout.CheckoutService.CheckoutResult:input_type -> frontend.common.Empty
	1, // 3: frontend.checkout.CheckoutService.Checkout:output_type -> frontend.common.Empty
	1, // 4: frontend.checkout.CheckoutService.CheckoutWaiting:output_type -> frontend.common.Empty
	1, // 5: frontend.checkout.CheckoutService.CheckoutResult:output_type -> frontend.common.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_checkout_page_proto_init() }
func file_checkout_page_proto_init() {
	if File_checkout_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkout_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutReq); i {
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
			RawDescriptor: file_checkout_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkout_page_proto_goTypes,
		DependencyIndexes: file_checkout_page_proto_depIdxs,
		MessageInfos:      file_checkout_page_proto_msgTypes,
	}.Build()
	File_checkout_page_proto = out.File
	file_checkout_page_proto_rawDesc = nil
	file_checkout_page_proto_goTypes = nil
	file_checkout_page_proto_depIdxs = nil
}
