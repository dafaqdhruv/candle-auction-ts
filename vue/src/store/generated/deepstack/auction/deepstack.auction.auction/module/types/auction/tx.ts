/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "deepstack.auction.auction";

export interface MsgCreateObject {
  creator: string;
  description: string;
  owner: string;
  value: string;
  prevOwnSign: string;
  startBlockHeight: number;
  endBlockHeight: number;
  auctionDuration: number;
  marginBlocks: number;
  highestBid: string;
  highestBidder: string;
}

export interface MsgCreateObjectResponse {
  id: number;
}

export interface MsgUpdateObject {
  creator: string;
  id: number;
  description: string;
  owner: string;
  value: string;
  prevOwnSign: string;
  startBlockHeight: number;
  endBlockHeight: number;
  auctionDuration: number;
  marginBlocks: number;
  highestBid: string;
  highestBidder: string;
}

export interface MsgUpdateObjectResponse {}

export interface MsgDeleteObject {
  creator: string;
  id: number;
}

export interface MsgDeleteObjectResponse {}

const baseMsgCreateObject: object = {
  creator: "",
  description: "",
  owner: "",
  value: "",
  prevOwnSign: "",
  startBlockHeight: 0,
  endBlockHeight: 0,
  auctionDuration: 0,
  marginBlocks: 0,
  highestBid: "",
  highestBidder: "",
};

export const MsgCreateObject = {
  encode(message: MsgCreateObject, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.value !== "") {
      writer.uint32(34).string(message.value);
    }
    if (message.prevOwnSign !== "") {
      writer.uint32(42).string(message.prevOwnSign);
    }
    if (message.startBlockHeight !== 0) {
      writer.uint32(48).uint64(message.startBlockHeight);
    }
    if (message.endBlockHeight !== 0) {
      writer.uint32(56).uint64(message.endBlockHeight);
    }
    if (message.auctionDuration !== 0) {
      writer.uint32(64).uint64(message.auctionDuration);
    }
    if (message.marginBlocks !== 0) {
      writer.uint32(72).uint64(message.marginBlocks);
    }
    if (message.highestBid !== "") {
      writer.uint32(82).string(message.highestBid);
    }
    if (message.highestBidder !== "") {
      writer.uint32(90).string(message.highestBidder);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateObject } as MsgCreateObject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.value = reader.string();
          break;
        case 5:
          message.prevOwnSign = reader.string();
          break;
        case 6:
          message.startBlockHeight = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.endBlockHeight = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.auctionDuration = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.marginBlocks = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.highestBid = reader.string();
          break;
        case 11:
          message.highestBidder = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateObject {
    const message = { ...baseMsgCreateObject } as MsgCreateObject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.prevOwnSign !== undefined && object.prevOwnSign !== null) {
      message.prevOwnSign = String(object.prevOwnSign);
    } else {
      message.prevOwnSign = "";
    }
    if (
      object.startBlockHeight !== undefined &&
      object.startBlockHeight !== null
    ) {
      message.startBlockHeight = Number(object.startBlockHeight);
    } else {
      message.startBlockHeight = 0;
    }
    if (object.endBlockHeight !== undefined && object.endBlockHeight !== null) {
      message.endBlockHeight = Number(object.endBlockHeight);
    } else {
      message.endBlockHeight = 0;
    }
    if (
      object.auctionDuration !== undefined &&
      object.auctionDuration !== null
    ) {
      message.auctionDuration = Number(object.auctionDuration);
    } else {
      message.auctionDuration = 0;
    }
    if (object.marginBlocks !== undefined && object.marginBlocks !== null) {
      message.marginBlocks = Number(object.marginBlocks);
    } else {
      message.marginBlocks = 0;
    }
    if (object.highestBid !== undefined && object.highestBid !== null) {
      message.highestBid = String(object.highestBid);
    } else {
      message.highestBid = "";
    }
    if (object.highestBidder !== undefined && object.highestBidder !== null) {
      message.highestBidder = String(object.highestBidder);
    } else {
      message.highestBidder = "";
    }
    return message;
  },

  toJSON(message: MsgCreateObject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.description !== undefined &&
      (obj.description = message.description);
    message.owner !== undefined && (obj.owner = message.owner);
    message.value !== undefined && (obj.value = message.value);
    message.prevOwnSign !== undefined &&
      (obj.prevOwnSign = message.prevOwnSign);
    message.startBlockHeight !== undefined &&
      (obj.startBlockHeight = message.startBlockHeight);
    message.endBlockHeight !== undefined &&
      (obj.endBlockHeight = message.endBlockHeight);
    message.auctionDuration !== undefined &&
      (obj.auctionDuration = message.auctionDuration);
    message.marginBlocks !== undefined &&
      (obj.marginBlocks = message.marginBlocks);
    message.highestBid !== undefined && (obj.highestBid = message.highestBid);
    message.highestBidder !== undefined &&
      (obj.highestBidder = message.highestBidder);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateObject>): MsgCreateObject {
    const message = { ...baseMsgCreateObject } as MsgCreateObject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.prevOwnSign !== undefined && object.prevOwnSign !== null) {
      message.prevOwnSign = object.prevOwnSign;
    } else {
      message.prevOwnSign = "";
    }
    if (
      object.startBlockHeight !== undefined &&
      object.startBlockHeight !== null
    ) {
      message.startBlockHeight = object.startBlockHeight;
    } else {
      message.startBlockHeight = 0;
    }
    if (object.endBlockHeight !== undefined && object.endBlockHeight !== null) {
      message.endBlockHeight = object.endBlockHeight;
    } else {
      message.endBlockHeight = 0;
    }
    if (
      object.auctionDuration !== undefined &&
      object.auctionDuration !== null
    ) {
      message.auctionDuration = object.auctionDuration;
    } else {
      message.auctionDuration = 0;
    }
    if (object.marginBlocks !== undefined && object.marginBlocks !== null) {
      message.marginBlocks = object.marginBlocks;
    } else {
      message.marginBlocks = 0;
    }
    if (object.highestBid !== undefined && object.highestBid !== null) {
      message.highestBid = object.highestBid;
    } else {
      message.highestBid = "";
    }
    if (object.highestBidder !== undefined && object.highestBidder !== null) {
      message.highestBidder = object.highestBidder;
    } else {
      message.highestBidder = "";
    }
    return message;
  },
};

const baseMsgCreateObjectResponse: object = { id: 0 };

export const MsgCreateObjectResponse = {
  encode(
    message: MsgCreateObjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateObjectResponse,
    } as MsgCreateObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateObjectResponse {
    const message = {
      ...baseMsgCreateObjectResponse,
    } as MsgCreateObjectResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateObjectResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateObjectResponse>
  ): MsgCreateObjectResponse {
    const message = {
      ...baseMsgCreateObjectResponse,
    } as MsgCreateObjectResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateObject: object = {
  creator: "",
  id: 0,
  description: "",
  owner: "",
  value: "",
  prevOwnSign: "",
  startBlockHeight: 0,
  endBlockHeight: 0,
  auctionDuration: 0,
  marginBlocks: 0,
  highestBid: "",
  highestBidder: "",
};

export const MsgUpdateObject = {
  encode(message: MsgUpdateObject, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.owner !== "") {
      writer.uint32(34).string(message.owner);
    }
    if (message.value !== "") {
      writer.uint32(42).string(message.value);
    }
    if (message.prevOwnSign !== "") {
      writer.uint32(50).string(message.prevOwnSign);
    }
    if (message.startBlockHeight !== 0) {
      writer.uint32(56).uint64(message.startBlockHeight);
    }
    if (message.endBlockHeight !== 0) {
      writer.uint32(64).uint64(message.endBlockHeight);
    }
    if (message.auctionDuration !== 0) {
      writer.uint32(72).uint64(message.auctionDuration);
    }
    if (message.marginBlocks !== 0) {
      writer.uint32(80).uint64(message.marginBlocks);
    }
    if (message.highestBid !== "") {
      writer.uint32(90).string(message.highestBid);
    }
    if (message.highestBidder !== "") {
      writer.uint32(98).string(message.highestBidder);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateObject } as MsgUpdateObject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.owner = reader.string();
          break;
        case 5:
          message.value = reader.string();
          break;
        case 6:
          message.prevOwnSign = reader.string();
          break;
        case 7:
          message.startBlockHeight = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.endBlockHeight = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.auctionDuration = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.marginBlocks = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.highestBid = reader.string();
          break;
        case 12:
          message.highestBidder = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateObject {
    const message = { ...baseMsgUpdateObject } as MsgUpdateObject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.prevOwnSign !== undefined && object.prevOwnSign !== null) {
      message.prevOwnSign = String(object.prevOwnSign);
    } else {
      message.prevOwnSign = "";
    }
    if (
      object.startBlockHeight !== undefined &&
      object.startBlockHeight !== null
    ) {
      message.startBlockHeight = Number(object.startBlockHeight);
    } else {
      message.startBlockHeight = 0;
    }
    if (object.endBlockHeight !== undefined && object.endBlockHeight !== null) {
      message.endBlockHeight = Number(object.endBlockHeight);
    } else {
      message.endBlockHeight = 0;
    }
    if (
      object.auctionDuration !== undefined &&
      object.auctionDuration !== null
    ) {
      message.auctionDuration = Number(object.auctionDuration);
    } else {
      message.auctionDuration = 0;
    }
    if (object.marginBlocks !== undefined && object.marginBlocks !== null) {
      message.marginBlocks = Number(object.marginBlocks);
    } else {
      message.marginBlocks = 0;
    }
    if (object.highestBid !== undefined && object.highestBid !== null) {
      message.highestBid = String(object.highestBid);
    } else {
      message.highestBid = "";
    }
    if (object.highestBidder !== undefined && object.highestBidder !== null) {
      message.highestBidder = String(object.highestBidder);
    } else {
      message.highestBidder = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateObject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.description !== undefined &&
      (obj.description = message.description);
    message.owner !== undefined && (obj.owner = message.owner);
    message.value !== undefined && (obj.value = message.value);
    message.prevOwnSign !== undefined &&
      (obj.prevOwnSign = message.prevOwnSign);
    message.startBlockHeight !== undefined &&
      (obj.startBlockHeight = message.startBlockHeight);
    message.endBlockHeight !== undefined &&
      (obj.endBlockHeight = message.endBlockHeight);
    message.auctionDuration !== undefined &&
      (obj.auctionDuration = message.auctionDuration);
    message.marginBlocks !== undefined &&
      (obj.marginBlocks = message.marginBlocks);
    message.highestBid !== undefined && (obj.highestBid = message.highestBid);
    message.highestBidder !== undefined &&
      (obj.highestBidder = message.highestBidder);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateObject>): MsgUpdateObject {
    const message = { ...baseMsgUpdateObject } as MsgUpdateObject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.prevOwnSign !== undefined && object.prevOwnSign !== null) {
      message.prevOwnSign = object.prevOwnSign;
    } else {
      message.prevOwnSign = "";
    }
    if (
      object.startBlockHeight !== undefined &&
      object.startBlockHeight !== null
    ) {
      message.startBlockHeight = object.startBlockHeight;
    } else {
      message.startBlockHeight = 0;
    }
    if (object.endBlockHeight !== undefined && object.endBlockHeight !== null) {
      message.endBlockHeight = object.endBlockHeight;
    } else {
      message.endBlockHeight = 0;
    }
    if (
      object.auctionDuration !== undefined &&
      object.auctionDuration !== null
    ) {
      message.auctionDuration = object.auctionDuration;
    } else {
      message.auctionDuration = 0;
    }
    if (object.marginBlocks !== undefined && object.marginBlocks !== null) {
      message.marginBlocks = object.marginBlocks;
    } else {
      message.marginBlocks = 0;
    }
    if (object.highestBid !== undefined && object.highestBid !== null) {
      message.highestBid = object.highestBid;
    } else {
      message.highestBid = "";
    }
    if (object.highestBidder !== undefined && object.highestBidder !== null) {
      message.highestBidder = object.highestBidder;
    } else {
      message.highestBidder = "";
    }
    return message;
  },
};

const baseMsgUpdateObjectResponse: object = {};

export const MsgUpdateObjectResponse = {
  encode(_: MsgUpdateObjectResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateObjectResponse,
    } as MsgUpdateObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateObjectResponse {
    const message = {
      ...baseMsgUpdateObjectResponse,
    } as MsgUpdateObjectResponse;
    return message;
  },

  toJSON(_: MsgUpdateObjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateObjectResponse>
  ): MsgUpdateObjectResponse {
    const message = {
      ...baseMsgUpdateObjectResponse,
    } as MsgUpdateObjectResponse;
    return message;
  },
};

const baseMsgDeleteObject: object = { creator: "", id: 0 };

export const MsgDeleteObject = {
  encode(message: MsgDeleteObject, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteObject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteObject } as MsgDeleteObject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteObject {
    const message = { ...baseMsgDeleteObject } as MsgDeleteObject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteObject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteObject>): MsgDeleteObject {
    const message = { ...baseMsgDeleteObject } as MsgDeleteObject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteObjectResponse: object = {};

export const MsgDeleteObjectResponse = {
  encode(_: MsgDeleteObjectResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteObjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteObjectResponse,
    } as MsgDeleteObjectResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteObjectResponse {
    const message = {
      ...baseMsgDeleteObjectResponse,
    } as MsgDeleteObjectResponse;
    return message;
  },

  toJSON(_: MsgDeleteObjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteObjectResponse>
  ): MsgDeleteObjectResponse {
    const message = {
      ...baseMsgDeleteObjectResponse,
    } as MsgDeleteObjectResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateObject(request: MsgCreateObject): Promise<MsgCreateObjectResponse>;
  UpdateObject(request: MsgUpdateObject): Promise<MsgUpdateObjectResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteObject(request: MsgDeleteObject): Promise<MsgDeleteObjectResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateObject(request: MsgCreateObject): Promise<MsgCreateObjectResponse> {
    const data = MsgCreateObject.encode(request).finish();
    const promise = this.rpc.request(
      "deepstack.auction.auction.Msg",
      "CreateObject",
      data
    );
    return promise.then((data) =>
      MsgCreateObjectResponse.decode(new Reader(data))
    );
  }

  UpdateObject(request: MsgUpdateObject): Promise<MsgUpdateObjectResponse> {
    const data = MsgUpdateObject.encode(request).finish();
    const promise = this.rpc.request(
      "deepstack.auction.auction.Msg",
      "UpdateObject",
      data
    );
    return promise.then((data) =>
      MsgUpdateObjectResponse.decode(new Reader(data))
    );
  }

  DeleteObject(request: MsgDeleteObject): Promise<MsgDeleteObjectResponse> {
    const data = MsgDeleteObject.encode(request).finish();
    const promise = this.rpc.request(
      "deepstack.auction.auction.Msg",
      "DeleteObject",
      data
    );
    return promise.then((data) =>
      MsgDeleteObjectResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
