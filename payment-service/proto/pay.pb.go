// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.2
// source: payment-service/proto/pay.proto

package paymentservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ChargeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Amount   int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	ChargeId string `protobuf:"bytes,6,opt,name=charge_id,json=chargeId,proto3" json:"charge_id,omitempty"`
}

func (x *ChargeRequest) Reset() {
	*x = ChargeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_service_proto_pay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeRequest) ProtoMessage() {}

func (x *ChargeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_service_proto_pay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeRequest.ProtoReflect.Descriptor instead.
func (*ChargeRequest) Descriptor() ([]byte, []int) {
	return file_payment_service_proto_pay_proto_rawDescGZIP(), []int{0}
}

func (x *ChargeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChargeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ChargeRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ChargeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChargeRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ChargeRequest) GetChargeId() string {
	if x != nil {
		return x.ChargeId
	}
	return ""
}

type ChargeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Paid     bool   `protobuf:"varint,2,opt,name=paid,proto3" json:"paid,omitempty"`
	Refunded bool   `protobuf:"varint,3,opt,name=refunded,proto3" json:"refunded,omitempty"`
	Captured bool   `protobuf:"varint,4,opt,name=captured,proto3" json:"captured,omitempty"`
	Amount   int64  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ChargeResponse) Reset() {
	*x = ChargeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_service_proto_pay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeResponse) ProtoMessage() {}

func (x *ChargeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_service_proto_pay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeResponse.ProtoReflect.Descriptor instead.
func (*ChargeResponse) Descriptor() ([]byte, []int) {
	return file_payment_service_proto_pay_proto_rawDescGZIP(), []int{1}
}

func (x *ChargeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChargeResponse) GetPaid() bool {
	if x != nil {
		return x.Paid
	}
	return false
}

func (x *ChargeResponse) GetRefunded() bool {
	if x != nil {
		return x.Refunded
	}
	return false
}

func (x *ChargeResponse) GetCaptured() bool {
	if x != nil {
		return x.Captured
	}
	return false
}

func (x *ChargeResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_payment_service_proto_pay_proto protoreflect.FileDescriptor

var file_payment_service_proto_pay_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x92, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x61, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xec, 0x01,
	0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x47, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x43, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x1d, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_service_proto_pay_proto_rawDescOnce sync.Once
	file_payment_service_proto_pay_proto_rawDescData = file_payment_service_proto_pay_proto_rawDesc
)

func file_payment_service_proto_pay_proto_rawDescGZIP() []byte {
	file_payment_service_proto_pay_proto_rawDescOnce.Do(func() {
		file_payment_service_proto_pay_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_service_proto_pay_proto_rawDescData)
	})
	return file_payment_service_proto_pay_proto_rawDescData
}

var file_payment_service_proto_pay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payment_service_proto_pay_proto_goTypes = []interface{}{
	(*ChargeRequest)(nil),  // 0: paymentservice.ChargeRequest
	(*ChargeResponse)(nil), // 1: paymentservice.ChargeResponse
}
var file_payment_service_proto_pay_proto_depIdxs = []int32{
	0, // 0: paymentservice.PaymentManager.Charge:input_type -> paymentservice.ChargeRequest
	0, // 1: paymentservice.PaymentManager.Capture:input_type -> paymentservice.ChargeRequest
	0, // 2: paymentservice.PaymentManager.Refund:input_type -> paymentservice.ChargeRequest
	1, // 3: paymentservice.PaymentManager.Charge:output_type -> paymentservice.ChargeResponse
	1, // 4: paymentservice.PaymentManager.Capture:output_type -> paymentservice.ChargeResponse
	1, // 5: paymentservice.PaymentManager.Refund:output_type -> paymentservice.ChargeResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_service_proto_pay_proto_init() }
func file_payment_service_proto_pay_proto_init() {
	if File_payment_service_proto_pay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_service_proto_pay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeRequest); i {
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
		file_payment_service_proto_pay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeResponse); i {
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
			RawDescriptor: file_payment_service_proto_pay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_service_proto_pay_proto_goTypes,
		DependencyIndexes: file_payment_service_proto_pay_proto_depIdxs,
		MessageInfos:      file_payment_service_proto_pay_proto_msgTypes,
	}.Build()
	File_payment_service_proto_pay_proto = out.File
	file_payment_service_proto_pay_proto_rawDesc = nil
	file_payment_service_proto_pay_proto_goTypes = nil
	file_payment_service_proto_pay_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaymentManagerClient is the client API for PaymentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentManagerClient interface {
	Charge(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error)
	Capture(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error)
	Refund(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error)
}

type paymentManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentManagerClient(cc grpc.ClientConnInterface) PaymentManagerClient {
	return &paymentManagerClient{cc}
}

func (c *paymentManagerClient) Charge(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error) {
	out := new(ChargeResponse)
	err := c.cc.Invoke(ctx, "/paymentservice.PaymentManager/Charge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentManagerClient) Capture(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error) {
	out := new(ChargeResponse)
	err := c.cc.Invoke(ctx, "/paymentservice.PaymentManager/Capture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentManagerClient) Refund(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error) {
	out := new(ChargeResponse)
	err := c.cc.Invoke(ctx, "/paymentservice.PaymentManager/Refund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentManagerServer is the server API for PaymentManager service.
type PaymentManagerServer interface {
	Charge(context.Context, *ChargeRequest) (*ChargeResponse, error)
	Capture(context.Context, *ChargeRequest) (*ChargeResponse, error)
	Refund(context.Context, *ChargeRequest) (*ChargeResponse, error)
}

// UnimplementedPaymentManagerServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentManagerServer struct {
}

func (*UnimplementedPaymentManagerServer) Charge(context.Context, *ChargeRequest) (*ChargeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Charge not implemented")
}
func (*UnimplementedPaymentManagerServer) Capture(context.Context, *ChargeRequest) (*ChargeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Capture not implemented")
}
func (*UnimplementedPaymentManagerServer) Refund(context.Context, *ChargeRequest) (*ChargeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refund not implemented")
}

func RegisterPaymentManagerServer(s *grpc.Server, srv PaymentManagerServer) {
	s.RegisterService(&_PaymentManager_serviceDesc, srv)
}

func _PaymentManager_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentManagerServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentservice.PaymentManager/Charge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentManagerServer).Charge(ctx, req.(*ChargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentManager_Capture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentManagerServer).Capture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentservice.PaymentManager/Capture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentManagerServer).Capture(ctx, req.(*ChargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentManager_Refund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentManagerServer).Refund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentservice.PaymentManager/Refund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentManagerServer).Refund(ctx, req.(*ChargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "paymentservice.PaymentManager",
	HandlerType: (*PaymentManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Charge",
			Handler:    _PaymentManager_Charge_Handler,
		},
		{
			MethodName: "Capture",
			Handler:    _PaymentManager_Capture_Handler,
		},
		{
			MethodName: "Refund",
			Handler:    _PaymentManager_Refund_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment-service/proto/pay.proto",
}
