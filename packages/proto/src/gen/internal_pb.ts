// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file internal.proto (package river, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { MiniblockHeader, StreamEvent, SyncCookie } from "./protocol_pb.js";

/**
 * @generated from message river.PersistedEvent
 */
export class PersistedEvent extends Message<PersistedEvent> {
  /**
   * @generated from field: river.StreamEvent event = 1;
   */
  event?: StreamEvent;

  /**
   * @generated from field: bytes hash = 2;
   */
  hash = new Uint8Array(0);

  /**
   * @generated from field: string prev_miniblock_hash_str = 3;
   */
  prevMiniblockHashStr = "";

  /**
   * @generated from field: string creator_user_id = 4;
   */
  creatorUserId = "";

  constructor(data?: PartialMessage<PersistedEvent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "river.PersistedEvent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "event", kind: "message", T: StreamEvent },
    { no: 2, name: "hash", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "prev_miniblock_hash_str", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "creator_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersistedEvent {
    return new PersistedEvent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersistedEvent {
    return new PersistedEvent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersistedEvent {
    return new PersistedEvent().fromJsonString(jsonString, options);
  }

  static equals(a: PersistedEvent | PlainMessage<PersistedEvent> | undefined, b: PersistedEvent | PlainMessage<PersistedEvent> | undefined): boolean {
    return proto3.util.equals(PersistedEvent, a, b);
  }
}

/**
 * @generated from message river.PersistedMiniblock
 */
export class PersistedMiniblock extends Message<PersistedMiniblock> {
  /**
   * @generated from field: bytes hash = 1;
   */
  hash = new Uint8Array(0);

  /**
   * @generated from field: river.MiniblockHeader header = 2;
   */
  header?: MiniblockHeader;

  /**
   * @generated from field: repeated river.PersistedEvent events = 3;
   */
  events: PersistedEvent[] = [];

  constructor(data?: PartialMessage<PersistedMiniblock>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "river.PersistedMiniblock";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "hash", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "header", kind: "message", T: MiniblockHeader },
    { no: 3, name: "events", kind: "message", T: PersistedEvent, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersistedMiniblock {
    return new PersistedMiniblock().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersistedMiniblock {
    return new PersistedMiniblock().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersistedMiniblock {
    return new PersistedMiniblock().fromJsonString(jsonString, options);
  }

  static equals(a: PersistedMiniblock | PlainMessage<PersistedMiniblock> | undefined, b: PersistedMiniblock | PlainMessage<PersistedMiniblock> | undefined): boolean {
    return proto3.util.equals(PersistedMiniblock, a, b);
  }
}

/**
 * @generated from message river.PersistedSyncedStream
 */
export class PersistedSyncedStream extends Message<PersistedSyncedStream> {
  /**
   * @generated from field: river.SyncCookie sync_cookie = 1;
   */
  syncCookie?: SyncCookie;

  /**
   * @generated from field: uint64 last_snapshot_miniblock_num = 2;
   */
  lastSnapshotMiniblockNum = protoInt64.zero;

  /**
   * @generated from field: uint64 last_miniblock_num = 3;
   */
  lastMiniblockNum = protoInt64.zero;

  /**
   * @generated from field: repeated river.PersistedEvent minipoolEvents = 4;
   */
  minipoolEvents: PersistedEvent[] = [];

  constructor(data?: PartialMessage<PersistedSyncedStream>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "river.PersistedSyncedStream";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sync_cookie", kind: "message", T: SyncCookie },
    { no: 2, name: "last_snapshot_miniblock_num", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "last_miniblock_num", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 4, name: "minipoolEvents", kind: "message", T: PersistedEvent, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PersistedSyncedStream {
    return new PersistedSyncedStream().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PersistedSyncedStream {
    return new PersistedSyncedStream().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PersistedSyncedStream {
    return new PersistedSyncedStream().fromJsonString(jsonString, options);
  }

  static equals(a: PersistedSyncedStream | PlainMessage<PersistedSyncedStream> | undefined, b: PersistedSyncedStream | PlainMessage<PersistedSyncedStream> | undefined): boolean {
    return proto3.util.equals(PersistedSyncedStream, a, b);
  }
}
