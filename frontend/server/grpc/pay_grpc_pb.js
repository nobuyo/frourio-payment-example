// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var pay_pb = require('./pay_pb.js');

function serialize_paymentservice_ChargeRequest(arg) {
  if (!(arg instanceof pay_pb.ChargeRequest)) {
    throw new Error('Expected argument of type paymentservice.ChargeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_paymentservice_ChargeRequest(buffer_arg) {
  return pay_pb.ChargeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_paymentservice_ChargeResponse(arg) {
  if (!(arg instanceof pay_pb.ChargeResponse)) {
    throw new Error('Expected argument of type paymentservice.ChargeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_paymentservice_ChargeResponse(buffer_arg) {
  return pay_pb.ChargeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var PaymentManagerService = exports.PaymentManagerService = {
  charge: {
    path: '/paymentservice.PaymentManager/Charge',
    requestStream: false,
    responseStream: false,
    requestType: pay_pb.ChargeRequest,
    responseType: pay_pb.ChargeResponse,
    requestSerialize: serialize_paymentservice_ChargeRequest,
    requestDeserialize: deserialize_paymentservice_ChargeRequest,
    responseSerialize: serialize_paymentservice_ChargeResponse,
    responseDeserialize: deserialize_paymentservice_ChargeResponse,
  },
  capture: {
    path: '/paymentservice.PaymentManager/Capture',
    requestStream: false,
    responseStream: false,
    requestType: pay_pb.ChargeRequest,
    responseType: pay_pb.ChargeResponse,
    requestSerialize: serialize_paymentservice_ChargeRequest,
    requestDeserialize: deserialize_paymentservice_ChargeRequest,
    responseSerialize: serialize_paymentservice_ChargeResponse,
    responseDeserialize: deserialize_paymentservice_ChargeResponse,
  },
  refund: {
    path: '/paymentservice.PaymentManager/Refund',
    requestStream: false,
    responseStream: false,
    requestType: pay_pb.ChargeRequest,
    responseType: pay_pb.ChargeResponse,
    requestSerialize: serialize_paymentservice_ChargeRequest,
    requestDeserialize: deserialize_paymentservice_ChargeRequest,
    responseSerialize: serialize_paymentservice_ChargeResponse,
    responseDeserialize: deserialize_paymentservice_ChargeResponse,
  },
};

exports.PaymentManagerClient = grpc.makeGenericClientConstructor(PaymentManagerService);
