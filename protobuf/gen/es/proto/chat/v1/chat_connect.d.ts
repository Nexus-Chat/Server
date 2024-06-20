// @generated by protoc-gen-connect-es v1.4.0
// @generated from file proto/chat/v1/chat.proto (package proto.chat.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateChannelRequest, CreateChannelResponse, CreateDirectRequest, CreateDirectResponse, CreateGroupRequest, CreateGroupResponse, GetChatRequest, GetChatResponse, GetUserChatsRequest, GetUserChatsResponse } from "./chat_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service proto.chat.v1.ChatService
 */
export declare const ChatService: {
  readonly typeName: "proto.chat.v1.ChatService",
  readonly methods: {
    /**
     * @generated from rpc proto.chat.v1.ChatService.GetChat
     */
    readonly getChat: {
      readonly name: "GetChat",
      readonly I: typeof GetChatRequest,
      readonly O: typeof GetChatResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.chat.v1.ChatService.GetUserChats
     */
    readonly getUserChats: {
      readonly name: "GetUserChats",
      readonly I: typeof GetUserChatsRequest,
      readonly O: typeof GetUserChatsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.chat.v1.ChatService.CreateDirect
     */
    readonly createDirect: {
      readonly name: "CreateDirect",
      readonly I: typeof CreateDirectRequest,
      readonly O: typeof CreateDirectResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.chat.v1.ChatService.CreateGroup
     */
    readonly createGroup: {
      readonly name: "CreateGroup",
      readonly I: typeof CreateGroupRequest,
      readonly O: typeof CreateGroupResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.chat.v1.ChatService.CreateChannel
     */
    readonly createChannel: {
      readonly name: "CreateChannel",
      readonly I: typeof CreateChannelRequest,
      readonly O: typeof CreateChannelResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

