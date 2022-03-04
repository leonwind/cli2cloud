import * as jspb from 'google-protobuf'



export class Client extends jspb.Message {
  getEncrypted(): boolean;
  setEncrypted(value: boolean): Client;

  getSalt(): Uint8Array | string;
  getSalt_asU8(): Uint8Array;
  getSalt_asB64(): string;
  setSalt(value: Uint8Array | string): Client;
  hasSalt(): boolean;
  clearSalt(): Client;

  getIv(): Uint8Array | string;
  getIv_asU8(): Uint8Array;
  getIv_asB64(): string;
  setIv(value: Uint8Array | string): Client;
  hasIv(): boolean;
  clearIv(): Client;

  getTimestamp(): number;
  setTimestamp(value: number): Client;
  hasTimestamp(): boolean;
  clearTimestamp(): Client;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Client.AsObject;
  static toObject(includeInstance: boolean, msg: Client): Client.AsObject;
  static serializeBinaryToWriter(message: Client, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Client;
  static deserializeBinaryFromReader(message: Client, reader: jspb.BinaryReader): Client;
}

export namespace Client {
  export type AsObject = {
    encrypted: boolean,
    salt?: Uint8Array | string,
    iv?: Uint8Array | string,
    timestamp?: number,
  }

  export enum SaltCase { 
    _SALT_NOT_SET = 0,
    SALT = 2,
  }

  export enum IvCase { 
    _IV_NOT_SET = 0,
    IV = 3,
  }

  export enum TimestampCase { 
    _TIMESTAMP_NOT_SET = 0,
    TIMESTAMP = 4,
  }
}

export class ClientId extends jspb.Message {
  getId(): string;
  setId(value: string): ClientId;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientId.AsObject;
  static toObject(includeInstance: boolean, msg: ClientId): ClientId.AsObject;
  static serializeBinaryToWriter(message: ClientId, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientId;
  static deserializeBinaryFromReader(message: ClientId, reader: jspb.BinaryReader): ClientId;
}

export namespace ClientId {
  export type AsObject = {
    id: string,
  }
}

export class PublishRequest extends jspb.Message {
  getClientid(): ClientId | undefined;
  setClientid(value?: ClientId): PublishRequest;
  hasClientid(): boolean;
  clearClientid(): PublishRequest;

  getPayload(): Payload | undefined;
  setPayload(value?: Payload): PublishRequest;
  hasPayload(): boolean;
  clearPayload(): PublishRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PublishRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PublishRequest): PublishRequest.AsObject;
  static serializeBinaryToWriter(message: PublishRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PublishRequest;
  static deserializeBinaryFromReader(message: PublishRequest, reader: jspb.BinaryReader): PublishRequest;
}

export namespace PublishRequest {
  export type AsObject = {
    clientid?: ClientId.AsObject,
    payload?: Payload.AsObject,
  }
}

export class Payload extends jspb.Message {
  getBody(): string;
  setBody(value: string): Payload;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Payload.AsObject;
  static toObject(includeInstance: boolean, msg: Payload): Payload.AsObject;
  static serializeBinaryToWriter(message: Payload, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Payload;
  static deserializeBinaryFromReader(message: Payload, reader: jspb.BinaryReader): Payload;
}

export namespace Payload {
  export type AsObject = {
    body: string,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

