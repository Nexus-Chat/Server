// @generated by protoc-gen-connect-es v1.4.0
// @generated from file proto/auth/v1/auth.proto (package proto.auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AuthenticateRequest, AuthenticateResponse, ChangePasswordRequest, ChangePasswordResponse, LoginRequest, LoginResponse, RefreshTokenRequest, RefreshTokenResponse, RegisterRequest, RegisterResponse, SendResetPasswordVerificationRequest, SendResetPasswordVerificationResponse, SubmitResetPasswordRequest, SubmitResetPasswordResponse, VerifyEmailRequest, VerifyEmailResponse } from "./auth_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service proto.auth.v1.AuthService
 */
export declare const AuthService: {
  readonly typeName: "proto.auth.v1.AuthService",
  readonly methods: {
    /**
     * @generated from rpc proto.auth.v1.AuthService.Login
     */
    readonly login: {
      readonly name: "Login",
      readonly I: typeof LoginRequest,
      readonly O: typeof LoginResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.auth.v1.AuthService.Register
     */
    readonly register: {
      readonly name: "Register",
      readonly I: typeof RegisterRequest,
      readonly O: typeof RegisterResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.auth.v1.AuthService.VerifyEmail
     */
    readonly verifyEmail: {
      readonly name: "VerifyEmail",
      readonly I: typeof VerifyEmailRequest,
      readonly O: typeof VerifyEmailResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.auth.v1.AuthService.SendResetPasswordVerification
     */
    readonly sendResetPasswordVerification: {
      readonly name: "SendResetPasswordVerification",
      readonly I: typeof SendResetPasswordVerificationRequest,
      readonly O: typeof SendResetPasswordVerificationResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.auth.v1.AuthService.SubmitResetPassword
     */
    readonly submitResetPassword: {
      readonly name: "SubmitResetPassword",
      readonly I: typeof SubmitResetPasswordRequest,
      readonly O: typeof SubmitResetPasswordResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.auth.v1.AuthService.ChangePassword
     */
    readonly changePassword: {
      readonly name: "ChangePassword",
      readonly I: typeof ChangePasswordRequest,
      readonly O: typeof ChangePasswordResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.auth.v1.AuthService.Authenticate
     */
    readonly authenticate: {
      readonly name: "Authenticate",
      readonly I: typeof AuthenticateRequest,
      readonly O: typeof AuthenticateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.auth.v1.AuthService.RefreshToken
     */
    readonly refreshToken: {
      readonly name: "RefreshToken",
      readonly I: typeof RefreshTokenRequest,
      readonly O: typeof RefreshTokenResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

