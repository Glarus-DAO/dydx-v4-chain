import { IndexerOrder, IndexerOrderAmino, IndexerOrderSDKType } from "../protocol/v1/clob";
import { BinaryReader, BinaryWriter } from "../../../binary";
/** Enum for the ticker type, PERPETUAL or SPOT. */
export enum RedisOrder_TickerType {
  /**
   * TICKER_TYPE_UNSPECIFIED - Default value for the enum. Should never be used in an initialized
   * `RedisOrder`.
   */
  TICKER_TYPE_UNSPECIFIED = 0,
  /** TICKER_TYPE_PERPETUAL - Ticker is for a perpetual pair. */
  TICKER_TYPE_PERPETUAL = 1,
  /** TICKER_TYPE_SPOT - Ticker is for a spot pair. */
  TICKER_TYPE_SPOT = 2,
  UNRECOGNIZED = -1,
}
export const RedisOrder_TickerTypeSDKType = RedisOrder_TickerType;
export const RedisOrder_TickerTypeAmino = RedisOrder_TickerType;
export function redisOrder_TickerTypeFromJSON(object: any): RedisOrder_TickerType {
  switch (object) {
    case 0:
    case "TICKER_TYPE_UNSPECIFIED":
      return RedisOrder_TickerType.TICKER_TYPE_UNSPECIFIED;
    case 1:
    case "TICKER_TYPE_PERPETUAL":
      return RedisOrder_TickerType.TICKER_TYPE_PERPETUAL;
    case 2:
    case "TICKER_TYPE_SPOT":
      return RedisOrder_TickerType.TICKER_TYPE_SPOT;
    case -1:
    case "UNRECOGNIZED":
    default:
      return RedisOrder_TickerType.UNRECOGNIZED;
  }
}
export function redisOrder_TickerTypeToJSON(object: RedisOrder_TickerType): string {
  switch (object) {
    case RedisOrder_TickerType.TICKER_TYPE_UNSPECIFIED:
      return "TICKER_TYPE_UNSPECIFIED";
    case RedisOrder_TickerType.TICKER_TYPE_PERPETUAL:
      return "TICKER_TYPE_PERPETUAL";
    case RedisOrder_TickerType.TICKER_TYPE_SPOT:
      return "TICKER_TYPE_SPOT";
    case RedisOrder_TickerType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
/**
 * RedisOrder is a proto for orders stored in Redis. This proto holds some
 * human-readable values such as price, size and ticker as well as the original
 * `Order` proto from the dYdX application.
 */
export interface RedisOrder {
  /** uuid of the Order generated by the Indexer based on the `OrderId`. */
  id: string;
  /** Order proto from the protocol. */
  order?: IndexerOrder;
  /** Ticker for the exchange pair for the order. */
  ticker: string;
  /** Type of the ticker, PERPETUAL or SPOT. */
  tickerType: RedisOrder_TickerType;
  /** Human-readable price of the order. */
  price: string;
  /** Human-readable size of the order. */
  size: string;
}
export interface RedisOrderProtoMsg {
  typeUrl: "/dydxprotocol.indexer.redis.RedisOrder";
  value: Uint8Array;
}
/**
 * RedisOrder is a proto for orders stored in Redis. This proto holds some
 * human-readable values such as price, size and ticker as well as the original
 * `Order` proto from the dYdX application.
 */
export interface RedisOrderAmino {
  /** uuid of the Order generated by the Indexer based on the `OrderId`. */
  id?: string;
  /** Order proto from the protocol. */
  order?: IndexerOrderAmino;
  /** Ticker for the exchange pair for the order. */
  ticker?: string;
  /** Type of the ticker, PERPETUAL or SPOT. */
  ticker_type?: RedisOrder_TickerType;
  /** Human-readable price of the order. */
  price?: string;
  /** Human-readable size of the order. */
  size?: string;
}
export interface RedisOrderAminoMsg {
  type: "/dydxprotocol.indexer.redis.RedisOrder";
  value: RedisOrderAmino;
}
/**
 * RedisOrder is a proto for orders stored in Redis. This proto holds some
 * human-readable values such as price, size and ticker as well as the original
 * `Order` proto from the dYdX application.
 */
export interface RedisOrderSDKType {
  id: string;
  order?: IndexerOrderSDKType;
  ticker: string;
  ticker_type: RedisOrder_TickerType;
  price: string;
  size: string;
}
function createBaseRedisOrder(): RedisOrder {
  return {
    id: "",
    order: undefined,
    ticker: "",
    tickerType: 0,
    price: "",
    size: ""
  };
}
export const RedisOrder = {
  typeUrl: "/dydxprotocol.indexer.redis.RedisOrder",
  encode(message: RedisOrder, writer: BinaryWriter = BinaryWriter.create()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.order !== undefined) {
      IndexerOrder.encode(message.order, writer.uint32(18).fork()).ldelim();
    }
    if (message.ticker !== "") {
      writer.uint32(26).string(message.ticker);
    }
    if (message.tickerType !== 0) {
      writer.uint32(32).int32(message.tickerType);
    }
    if (message.price !== "") {
      writer.uint32(42).string(message.price);
    }
    if (message.size !== "") {
      writer.uint32(50).string(message.size);
    }
    return writer;
  },
  decode(input: BinaryReader | Uint8Array, length?: number): RedisOrder {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRedisOrder();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.order = IndexerOrder.decode(reader, reader.uint32());
          break;
        case 3:
          message.ticker = reader.string();
          break;
        case 4:
          message.tickerType = (reader.int32() as any);
          break;
        case 5:
          message.price = reader.string();
          break;
        case 6:
          message.size = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromPartial(object: Partial<RedisOrder>): RedisOrder {
    const message = createBaseRedisOrder();
    message.id = object.id ?? "";
    message.order = object.order !== undefined && object.order !== null ? IndexerOrder.fromPartial(object.order) : undefined;
    message.ticker = object.ticker ?? "";
    message.tickerType = object.tickerType ?? 0;
    message.price = object.price ?? "";
    message.size = object.size ?? "";
    return message;
  },
  fromAmino(object: RedisOrderAmino): RedisOrder {
    const message = createBaseRedisOrder();
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    }
    if (object.order !== undefined && object.order !== null) {
      message.order = IndexerOrder.fromAmino(object.order);
    }
    if (object.ticker !== undefined && object.ticker !== null) {
      message.ticker = object.ticker;
    }
    if (object.ticker_type !== undefined && object.ticker_type !== null) {
      message.tickerType = redisOrder_TickerTypeFromJSON(object.ticker_type);
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = object.size;
    }
    return message;
  },
  toAmino(message: RedisOrder): RedisOrderAmino {
    const obj: any = {};
    obj.id = message.id;
    obj.order = message.order ? IndexerOrder.toAmino(message.order) : undefined;
    obj.ticker = message.ticker;
    obj.ticker_type = redisOrder_TickerTypeToJSON(message.tickerType);
    obj.price = message.price;
    obj.size = message.size;
    return obj;
  },
  fromAminoMsg(object: RedisOrderAminoMsg): RedisOrder {
    return RedisOrder.fromAmino(object.value);
  },
  fromProtoMsg(message: RedisOrderProtoMsg): RedisOrder {
    return RedisOrder.decode(message.value);
  },
  toProto(message: RedisOrder): Uint8Array {
    return RedisOrder.encode(message).finish();
  },
  toProtoMsg(message: RedisOrder): RedisOrderProtoMsg {
    return {
      typeUrl: "/dydxprotocol.indexer.redis.RedisOrder",
      value: RedisOrder.encode(message).finish()
    };
  }
};