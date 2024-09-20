// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protobuf/chat/v1/chat.proto

package chatv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/kavkaco/Kavka-Core/protobuf/gen/go/protobuf/chat/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ChatServiceName is the fully-qualified name of the ChatService service.
	ChatServiceName = "protobuf.chat.v1.ChatService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChatServiceGetChatProcedure is the fully-qualified name of the ChatService's GetChat RPC.
	ChatServiceGetChatProcedure = "/protobuf.chat.v1.ChatService/GetChat"
	// ChatServiceGetDirectChatProcedure is the fully-qualified name of the ChatService's GetDirectChat
	// RPC.
	ChatServiceGetDirectChatProcedure = "/protobuf.chat.v1.ChatService/GetDirectChat"
	// ChatServiceGetUserChatsProcedure is the fully-qualified name of the ChatService's GetUserChats
	// RPC.
	ChatServiceGetUserChatsProcedure = "/protobuf.chat.v1.ChatService/GetUserChats"
	// ChatServiceCreateDirectProcedure is the fully-qualified name of the ChatService's CreateDirect
	// RPC.
	ChatServiceCreateDirectProcedure = "/protobuf.chat.v1.ChatService/CreateDirect"
	// ChatServiceCreateGroupProcedure is the fully-qualified name of the ChatService's CreateGroup RPC.
	ChatServiceCreateGroupProcedure = "/protobuf.chat.v1.ChatService/CreateGroup"
	// ChatServiceCreateChannelProcedure is the fully-qualified name of the ChatService's CreateChannel
	// RPC.
	ChatServiceCreateChannelProcedure = "/protobuf.chat.v1.ChatService/CreateChannel"
	// ChatServiceJoinChatProcedure is the fully-qualified name of the ChatService's JoinChat RPC.
	ChatServiceJoinChatProcedure = "/protobuf.chat.v1.ChatService/JoinChat"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	chatServiceServiceDescriptor             = v1.File_protobuf_chat_v1_chat_proto.Services().ByName("ChatService")
	chatServiceGetChatMethodDescriptor       = chatServiceServiceDescriptor.Methods().ByName("GetChat")
	chatServiceGetDirectChatMethodDescriptor = chatServiceServiceDescriptor.Methods().ByName("GetDirectChat")
	chatServiceGetUserChatsMethodDescriptor  = chatServiceServiceDescriptor.Methods().ByName("GetUserChats")
	chatServiceCreateDirectMethodDescriptor  = chatServiceServiceDescriptor.Methods().ByName("CreateDirect")
	chatServiceCreateGroupMethodDescriptor   = chatServiceServiceDescriptor.Methods().ByName("CreateGroup")
	chatServiceCreateChannelMethodDescriptor = chatServiceServiceDescriptor.Methods().ByName("CreateChannel")
	chatServiceJoinChatMethodDescriptor      = chatServiceServiceDescriptor.Methods().ByName("JoinChat")
)

// ChatServiceClient is a client for the protobuf.chat.v1.ChatService service.
type ChatServiceClient interface {
	GetChat(context.Context, *connect.Request[v1.GetChatRequest]) (*connect.Response[v1.GetChatResponse], error)
	GetDirectChat(context.Context, *connect.Request[v1.GetDirectChatRequest]) (*connect.Response[v1.GetDirectChatResponse], error)
	GetUserChats(context.Context, *connect.Request[v1.GetUserChatsRequest]) (*connect.Response[v1.GetUserChatsResponse], error)
	CreateDirect(context.Context, *connect.Request[v1.CreateDirectRequest]) (*connect.Response[v1.CreateDirectResponse], error)
	CreateGroup(context.Context, *connect.Request[v1.CreateGroupRequest]) (*connect.Response[v1.CreateGroupResponse], error)
	CreateChannel(context.Context, *connect.Request[v1.CreateChannelRequest]) (*connect.Response[v1.CreateChannelResponse], error)
	JoinChat(context.Context, *connect.Request[v1.JoinChatRequest]) (*connect.Response[v1.JoinChatResponse], error)
}

// NewChatServiceClient constructs a client for the protobuf.chat.v1.ChatService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChatServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ChatServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &chatServiceClient{
		getChat: connect.NewClient[v1.GetChatRequest, v1.GetChatResponse](
			httpClient,
			baseURL+ChatServiceGetChatProcedure,
			connect.WithSchema(chatServiceGetChatMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getDirectChat: connect.NewClient[v1.GetDirectChatRequest, v1.GetDirectChatResponse](
			httpClient,
			baseURL+ChatServiceGetDirectChatProcedure,
			connect.WithSchema(chatServiceGetDirectChatMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getUserChats: connect.NewClient[v1.GetUserChatsRequest, v1.GetUserChatsResponse](
			httpClient,
			baseURL+ChatServiceGetUserChatsProcedure,
			connect.WithSchema(chatServiceGetUserChatsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createDirect: connect.NewClient[v1.CreateDirectRequest, v1.CreateDirectResponse](
			httpClient,
			baseURL+ChatServiceCreateDirectProcedure,
			connect.WithSchema(chatServiceCreateDirectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createGroup: connect.NewClient[v1.CreateGroupRequest, v1.CreateGroupResponse](
			httpClient,
			baseURL+ChatServiceCreateGroupProcedure,
			connect.WithSchema(chatServiceCreateGroupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createChannel: connect.NewClient[v1.CreateChannelRequest, v1.CreateChannelResponse](
			httpClient,
			baseURL+ChatServiceCreateChannelProcedure,
			connect.WithSchema(chatServiceCreateChannelMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		joinChat: connect.NewClient[v1.JoinChatRequest, v1.JoinChatResponse](
			httpClient,
			baseURL+ChatServiceJoinChatProcedure,
			connect.WithSchema(chatServiceJoinChatMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// chatServiceClient implements ChatServiceClient.
type chatServiceClient struct {
	getChat       *connect.Client[v1.GetChatRequest, v1.GetChatResponse]
	getDirectChat *connect.Client[v1.GetDirectChatRequest, v1.GetDirectChatResponse]
	getUserChats  *connect.Client[v1.GetUserChatsRequest, v1.GetUserChatsResponse]
	createDirect  *connect.Client[v1.CreateDirectRequest, v1.CreateDirectResponse]
	createGroup   *connect.Client[v1.CreateGroupRequest, v1.CreateGroupResponse]
	createChannel *connect.Client[v1.CreateChannelRequest, v1.CreateChannelResponse]
	joinChat      *connect.Client[v1.JoinChatRequest, v1.JoinChatResponse]
}

// GetChat calls protobuf.chat.v1.ChatService.GetChat.
func (c *chatServiceClient) GetChat(ctx context.Context, req *connect.Request[v1.GetChatRequest]) (*connect.Response[v1.GetChatResponse], error) {
	return c.getChat.CallUnary(ctx, req)
}

// GetDirectChat calls protobuf.chat.v1.ChatService.GetDirectChat.
func (c *chatServiceClient) GetDirectChat(ctx context.Context, req *connect.Request[v1.GetDirectChatRequest]) (*connect.Response[v1.GetDirectChatResponse], error) {
	return c.getDirectChat.CallUnary(ctx, req)
}

// GetUserChats calls protobuf.chat.v1.ChatService.GetUserChats.
func (c *chatServiceClient) GetUserChats(ctx context.Context, req *connect.Request[v1.GetUserChatsRequest]) (*connect.Response[v1.GetUserChatsResponse], error) {
	return c.getUserChats.CallUnary(ctx, req)
}

// CreateDirect calls protobuf.chat.v1.ChatService.CreateDirect.
func (c *chatServiceClient) CreateDirect(ctx context.Context, req *connect.Request[v1.CreateDirectRequest]) (*connect.Response[v1.CreateDirectResponse], error) {
	return c.createDirect.CallUnary(ctx, req)
}

// CreateGroup calls protobuf.chat.v1.ChatService.CreateGroup.
func (c *chatServiceClient) CreateGroup(ctx context.Context, req *connect.Request[v1.CreateGroupRequest]) (*connect.Response[v1.CreateGroupResponse], error) {
	return c.createGroup.CallUnary(ctx, req)
}

// CreateChannel calls protobuf.chat.v1.ChatService.CreateChannel.
func (c *chatServiceClient) CreateChannel(ctx context.Context, req *connect.Request[v1.CreateChannelRequest]) (*connect.Response[v1.CreateChannelResponse], error) {
	return c.createChannel.CallUnary(ctx, req)
}

// JoinChat calls protobuf.chat.v1.ChatService.JoinChat.
func (c *chatServiceClient) JoinChat(ctx context.Context, req *connect.Request[v1.JoinChatRequest]) (*connect.Response[v1.JoinChatResponse], error) {
	return c.joinChat.CallUnary(ctx, req)
}

// ChatServiceHandler is an implementation of the protobuf.chat.v1.ChatService service.
type ChatServiceHandler interface {
	GetChat(context.Context, *connect.Request[v1.GetChatRequest]) (*connect.Response[v1.GetChatResponse], error)
	GetDirectChat(context.Context, *connect.Request[v1.GetDirectChatRequest]) (*connect.Response[v1.GetDirectChatResponse], error)
	GetUserChats(context.Context, *connect.Request[v1.GetUserChatsRequest]) (*connect.Response[v1.GetUserChatsResponse], error)
	CreateDirect(context.Context, *connect.Request[v1.CreateDirectRequest]) (*connect.Response[v1.CreateDirectResponse], error)
	CreateGroup(context.Context, *connect.Request[v1.CreateGroupRequest]) (*connect.Response[v1.CreateGroupResponse], error)
	CreateChannel(context.Context, *connect.Request[v1.CreateChannelRequest]) (*connect.Response[v1.CreateChannelResponse], error)
	JoinChat(context.Context, *connect.Request[v1.JoinChatRequest]) (*connect.Response[v1.JoinChatResponse], error)
}

// NewChatServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChatServiceHandler(svc ChatServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	chatServiceGetChatHandler := connect.NewUnaryHandler(
		ChatServiceGetChatProcedure,
		svc.GetChat,
		connect.WithSchema(chatServiceGetChatMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chatServiceGetDirectChatHandler := connect.NewUnaryHandler(
		ChatServiceGetDirectChatProcedure,
		svc.GetDirectChat,
		connect.WithSchema(chatServiceGetDirectChatMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chatServiceGetUserChatsHandler := connect.NewUnaryHandler(
		ChatServiceGetUserChatsProcedure,
		svc.GetUserChats,
		connect.WithSchema(chatServiceGetUserChatsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chatServiceCreateDirectHandler := connect.NewUnaryHandler(
		ChatServiceCreateDirectProcedure,
		svc.CreateDirect,
		connect.WithSchema(chatServiceCreateDirectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chatServiceCreateGroupHandler := connect.NewUnaryHandler(
		ChatServiceCreateGroupProcedure,
		svc.CreateGroup,
		connect.WithSchema(chatServiceCreateGroupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chatServiceCreateChannelHandler := connect.NewUnaryHandler(
		ChatServiceCreateChannelProcedure,
		svc.CreateChannel,
		connect.WithSchema(chatServiceCreateChannelMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chatServiceJoinChatHandler := connect.NewUnaryHandler(
		ChatServiceJoinChatProcedure,
		svc.JoinChat,
		connect.WithSchema(chatServiceJoinChatMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/protobuf.chat.v1.ChatService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ChatServiceGetChatProcedure:
			chatServiceGetChatHandler.ServeHTTP(w, r)
		case ChatServiceGetDirectChatProcedure:
			chatServiceGetDirectChatHandler.ServeHTTP(w, r)
		case ChatServiceGetUserChatsProcedure:
			chatServiceGetUserChatsHandler.ServeHTTP(w, r)
		case ChatServiceCreateDirectProcedure:
			chatServiceCreateDirectHandler.ServeHTTP(w, r)
		case ChatServiceCreateGroupProcedure:
			chatServiceCreateGroupHandler.ServeHTTP(w, r)
		case ChatServiceCreateChannelProcedure:
			chatServiceCreateChannelHandler.ServeHTTP(w, r)
		case ChatServiceJoinChatProcedure:
			chatServiceJoinChatHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedChatServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChatServiceHandler struct{}

func (UnimplementedChatServiceHandler) GetChat(context.Context, *connect.Request[v1.GetChatRequest]) (*connect.Response[v1.GetChatResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.chat.v1.ChatService.GetChat is not implemented"))
}

func (UnimplementedChatServiceHandler) GetDirectChat(context.Context, *connect.Request[v1.GetDirectChatRequest]) (*connect.Response[v1.GetDirectChatResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.chat.v1.ChatService.GetDirectChat is not implemented"))
}

func (UnimplementedChatServiceHandler) GetUserChats(context.Context, *connect.Request[v1.GetUserChatsRequest]) (*connect.Response[v1.GetUserChatsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.chat.v1.ChatService.GetUserChats is not implemented"))
}

func (UnimplementedChatServiceHandler) CreateDirect(context.Context, *connect.Request[v1.CreateDirectRequest]) (*connect.Response[v1.CreateDirectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.chat.v1.ChatService.CreateDirect is not implemented"))
}

func (UnimplementedChatServiceHandler) CreateGroup(context.Context, *connect.Request[v1.CreateGroupRequest]) (*connect.Response[v1.CreateGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.chat.v1.ChatService.CreateGroup is not implemented"))
}

func (UnimplementedChatServiceHandler) CreateChannel(context.Context, *connect.Request[v1.CreateChannelRequest]) (*connect.Response[v1.CreateChannelResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.chat.v1.ChatService.CreateChannel is not implemented"))
}

func (UnimplementedChatServiceHandler) JoinChat(context.Context, *connect.Request[v1.JoinChatRequest]) (*connect.Response[v1.JoinChatResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("protobuf.chat.v1.ChatService.JoinChat is not implemented"))
}
