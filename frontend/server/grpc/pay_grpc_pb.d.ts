// package: paymentservice
// file: pay.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as pay_pb from "./pay_pb";

interface IPaymentManagerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    charge: IPaymentManagerService_ICharge;
    capture: IPaymentManagerService_ICapture;
    refund: IPaymentManagerService_IRefund;
}

interface IPaymentManagerService_ICharge extends grpc.MethodDefinition<pay_pb.ChargeRequest, pay_pb.ChargeResponse> {
    path: "/paymentservice.PaymentManager/Charge";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pay_pb.ChargeRequest>;
    requestDeserialize: grpc.deserialize<pay_pb.ChargeRequest>;
    responseSerialize: grpc.serialize<pay_pb.ChargeResponse>;
    responseDeserialize: grpc.deserialize<pay_pb.ChargeResponse>;
}
interface IPaymentManagerService_ICapture extends grpc.MethodDefinition<pay_pb.ChargeRequest, pay_pb.ChargeResponse> {
    path: "/paymentservice.PaymentManager/Capture";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pay_pb.ChargeRequest>;
    requestDeserialize: grpc.deserialize<pay_pb.ChargeRequest>;
    responseSerialize: grpc.serialize<pay_pb.ChargeResponse>;
    responseDeserialize: grpc.deserialize<pay_pb.ChargeResponse>;
}
interface IPaymentManagerService_IRefund extends grpc.MethodDefinition<pay_pb.ChargeRequest, pay_pb.ChargeResponse> {
    path: "/paymentservice.PaymentManager/Refund";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pay_pb.ChargeRequest>;
    requestDeserialize: grpc.deserialize<pay_pb.ChargeRequest>;
    responseSerialize: grpc.serialize<pay_pb.ChargeResponse>;
    responseDeserialize: grpc.deserialize<pay_pb.ChargeResponse>;
}

export const PaymentManagerService: IPaymentManagerService;

export interface IPaymentManagerServer {
    charge: grpc.handleUnaryCall<pay_pb.ChargeRequest, pay_pb.ChargeResponse>;
    capture: grpc.handleUnaryCall<pay_pb.ChargeRequest, pay_pb.ChargeResponse>;
    refund: grpc.handleUnaryCall<pay_pb.ChargeRequest, pay_pb.ChargeResponse>;
}

export interface IPaymentManagerClient {
    charge(request: pay_pb.ChargeRequest, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    charge(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    charge(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    capture(request: pay_pb.ChargeRequest, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    capture(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    capture(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    refund(request: pay_pb.ChargeRequest, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    refund(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    refund(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
}

export class PaymentManagerClient extends grpc.Client implements IPaymentManagerClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public charge(request: pay_pb.ChargeRequest, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    public charge(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    public charge(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    public capture(request: pay_pb.ChargeRequest, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    public capture(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    public capture(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    public refund(request: pay_pb.ChargeRequest, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    public refund(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
    public refund(request: pay_pb.ChargeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pay_pb.ChargeResponse) => void): grpc.ClientUnaryCall;
}
