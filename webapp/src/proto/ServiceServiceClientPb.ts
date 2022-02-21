/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as service_pb from './service_pb';


export class Cli2CloudClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorRegisterClient = new grpcWeb.MethodDescriptor(
    '/proto.Cli2Cloud/RegisterClient',
    grpcWeb.MethodType.UNARY,
    service_pb.Empty,
    service_pb.Client,
    (request: service_pb.Empty) => {
      return request.serializeBinary();
    },
    service_pb.Client.deserializeBinary
  );

  registerClient(
    request: service_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<service_pb.Client>;

  registerClient(
    request: service_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: service_pb.Client) => void): grpcWeb.ClientReadableStream<service_pb.Client>;

  registerClient(
    request: service_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: service_pb.Client) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.Cli2Cloud/RegisterClient',
        request,
        metadata || {},
        this.methodDescriptorRegisterClient,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.Cli2Cloud/RegisterClient',
    request,
    metadata || {},
    this.methodDescriptorRegisterClient);
  }

  methodDescriptorSubscribe = new grpcWeb.MethodDescriptor(
    '/proto.Cli2Cloud/Subscribe',
    grpcWeb.MethodType.SERVER_STREAMING,
    service_pb.Client,
    service_pb.Content,
    (request: service_pb.Client) => {
      return request.serializeBinary();
    },
    service_pb.Content.deserializeBinary
  );

  subscribe(
    request: service_pb.Client,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<service_pb.Content> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/proto.Cli2Cloud/Subscribe',
      request,
      metadata || {},
      this.methodDescriptorSubscribe);
  }

}

