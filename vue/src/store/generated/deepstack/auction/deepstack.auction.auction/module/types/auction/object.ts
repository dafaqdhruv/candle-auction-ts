/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "deepstack.auction.auction";

export interface Object {
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
  creator: string;
}

const baseObject: object = {
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
  creator: "",
};

export const Object = {
  encode(message: Object, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
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
    if (message.creator !== "") {
      writer.uint32(98).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Object {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseObject } as Object;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
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
        case 12:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Object {
    const message = { ...baseObject } as Object;
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Object): unknown {
    const obj: any = {};
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
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Object>): Object {
    const message = { ...baseObject } as Object;
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

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
