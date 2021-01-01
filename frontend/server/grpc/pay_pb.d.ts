// package: paymentservice
// file: pay.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class ChargeRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): ChargeRequest;

    getToken(): string;
    setToken(value: string): ChargeRequest;

    getAmount(): number;
    setAmount(value: number): ChargeRequest;

    getName(): string;
    setName(value: string): ChargeRequest;

    getDesc(): string;
    setDesc(value: string): ChargeRequest;

    getChargeId(): string;
    setChargeId(value: string): ChargeRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ChargeRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ChargeRequest): ChargeRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ChargeRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ChargeRequest;
    static deserializeBinaryFromReader(message: ChargeRequest, reader: jspb.BinaryReader): ChargeRequest;
}

export namespace ChargeRequest {
    export type AsObject = {
        id: number,
        token: string,
        amount: number,
        name: string,
        desc: string,
        chargeId: string,
    }
}

export class ChargeResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): ChargeResponse;

    getPaid(): boolean;
    setPaid(value: boolean): ChargeResponse;

    getRefunded(): boolean;
    setRefunded(value: boolean): ChargeResponse;

    getCaptured(): boolean;
    setCaptured(value: boolean): ChargeResponse;

    getAmount(): number;
    setAmount(value: number): ChargeResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ChargeResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ChargeResponse): ChargeResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ChargeResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ChargeResponse;
    static deserializeBinaryFromReader(message: ChargeResponse, reader: jspb.BinaryReader): ChargeResponse;
}

export namespace ChargeResponse {
    export type AsObject = {
        id: string,
        paid: boolean,
        refunded: boolean,
        captured: boolean,
        amount: number,
    }
}
