// @generated by protoc-gen-es v1.10.0
// @generated from file protobuf/chat/v1/chat.proto (package protobuf.chat.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { Chat } from "../../model/chat/v1/chat_pb.js";

/**
 * @generated from message protobuf.chat.v1.GetChatRequest
 */
export const GetChatRequest = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.GetChatRequest",
  () => [
    { no: 1, name: "chat_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message protobuf.chat.v1.GetChatResponse
 */
export const GetChatResponse = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.GetChatResponse",
  () => [
    { no: 1, name: "chat", kind: "message", T: Chat },
  ],
);

/**
 * @generated from message protobuf.chat.v1.GetUserChatsRequest
 */
export const GetUserChatsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.GetUserChatsRequest",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message protobuf.chat.v1.GetUserChatsResponse
 */
export const GetUserChatsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.GetUserChatsResponse",
  () => [
    { no: 1, name: "chats", kind: "message", T: Chat, repeated: true },
  ],
);

/**
 * @generated from message protobuf.chat.v1.CreateDirectRequest
 */
export const CreateDirectRequest = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.CreateDirectRequest",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "recipient_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message protobuf.chat.v1.CreateDirectResponse
 */
export const CreateDirectResponse = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.CreateDirectResponse",
  () => [
    { no: 1, name: "chat", kind: "message", T: Chat },
  ],
);

/**
 * @generated from message protobuf.chat.v1.CreateGroupRequest
 */
export const CreateGroupRequest = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.CreateGroupRequest",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message protobuf.chat.v1.CreateGroupResponse
 */
export const CreateGroupResponse = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.CreateGroupResponse",
  () => [
    { no: 1, name: "chat", kind: "message", T: Chat },
  ],
);

/**
 * @generated from message protobuf.chat.v1.CreateChannelRequest
 */
export const CreateChannelRequest = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.CreateChannelRequest",
  () => [
    { no: 1, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message protobuf.chat.v1.CreateChannelResponse
 */
export const CreateChannelResponse = /*@__PURE__*/ proto3.makeMessageType(
  "protobuf.chat.v1.CreateChannelResponse",
  () => [
    { no: 1, name: "chat", kind: "message", T: Chat },
  ],
);

