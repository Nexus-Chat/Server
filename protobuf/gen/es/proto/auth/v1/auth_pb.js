// @generated by protoc-gen-es v1.10.0
// @generated from file proto/auth/v1/auth.proto (package proto.auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, proto3 } from "@bufbuild/protobuf";
import { User } from "../../model/user/v1/user_pb.js";

/**
 * @generated from message proto.auth.v1.LoginRequest
 */
export const LoginRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.LoginRequest",
  () => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.LoginResponse
 */
export const LoginResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.LoginResponse",
  () => [
    { no: 1, name: "user", kind: "message", T: User },
    { no: 2, name: "access_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "refresh_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.RegisterRequest
 */
export const RegisterRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.RegisterRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "last_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.RegisterResponse
 */
export const RegisterResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.RegisterResponse",
  () => [
    { no: 1, name: "user", kind: "message", T: User },
    { no: 2, name: "verify_email_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.VerifyEmailRequest
 */
export const VerifyEmailRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.VerifyEmailRequest",
  () => [
    { no: 1, name: "verify_email_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.VerifyEmailResponse
 */
export const VerifyEmailResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.VerifyEmailResponse",
  [],
);

/**
 * @generated from message proto.auth.v1.SendResetPasswordVerificationRequest
 */
export const SendResetPasswordVerificationRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.SendResetPasswordVerificationRequest",
  () => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.SendResetPasswordVerificationResponse
 */
export const SendResetPasswordVerificationResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.SendResetPasswordVerificationResponse",
  () => [
    { no: 1, name: "verify_email_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "timeout", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message proto.auth.v1.SubmitResetPasswordRequest
 */
export const SubmitResetPasswordRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.SubmitResetPasswordRequest",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "new_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.SubmitResetPasswordResponse
 */
export const SubmitResetPasswordResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.SubmitResetPasswordResponse",
  [],
);

/**
 * @generated from message proto.auth.v1.ChangePasswordRequest
 */
export const ChangePasswordRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.ChangePasswordRequest",
  () => [
    { no: 1, name: "access_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "old_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "new_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.ChangePasswordResponse
 */
export const ChangePasswordResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.ChangePasswordResponse",
  [],
);

/**
 * @generated from message proto.auth.v1.AuthenticateRequest
 */
export const AuthenticateRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.AuthenticateRequest",
  () => [
    { no: 1, name: "access_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.AuthenticateResponse
 */
export const AuthenticateResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.AuthenticateResponse",
  () => [
    { no: 1, name: "user", kind: "message", T: User },
  ],
);

/**
 * @generated from message proto.auth.v1.RefreshTokenRequest
 */
export const RefreshTokenRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.RefreshTokenRequest",
  () => [
    { no: 1, name: "refresh_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "access_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proto.auth.v1.RefreshTokenResponse
 */
export const RefreshTokenResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proto.auth.v1.RefreshTokenResponse",
  () => [
    { no: 1, name: "access_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

