// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts,import_extension=.ts"
// @generated from file protobuf/message/message.proto (package protobuf.message.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FetchMessagesRequest, FetchMessagesResponse, SendTextMessageRequest, SendTextMessageResponse } from "./message_pb.ts";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service protobuf.message.v1.MessageService
 */
export const MessageService = {
  typeName: "protobuf.message.v1.MessageService",
  methods: {
    /**
     * @generated from rpc protobuf.message.v1.MessageService.FetchMessages
     */
    fetchMessages: {
      name: "FetchMessages",
      I: FetchMessagesRequest,
      O: FetchMessagesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc protobuf.message.v1.MessageService.SendTextMessage
     */
    sendTextMessage: {
      name: "SendTextMessage",
      I: SendTextMessageRequest,
      O: SendTextMessageResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

