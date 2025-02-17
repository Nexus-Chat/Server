package grpc

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"github.com/kavkaco/Kavka-Core/config"
	grpc_handlers "github.com/kavkaco/Kavka-Core/delivery/grpc/handlers"
	"github.com/kavkaco/Kavka-Core/delivery/grpc/interceptor"
	"github.com/kavkaco/Kavka-Core/infra/stream"
	"github.com/kavkaco/Kavka-Core/internal/service/auth"
	"github.com/kavkaco/Kavka-Core/internal/service/chat"
	"github.com/kavkaco/Kavka-Core/internal/service/message"
	"github.com/kavkaco/Kavka-Core/internal/service/search"
	"github.com/kavkaco/Kavka-Core/log"
	"github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/auth/v1/authv1connect"
	"github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/chat/v1/chatv1connect"
	"github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/events/v1/eventsv1connect"
	"github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/message/v1/messagev1connect"
	"github.com/kavkaco/Kavka-ProtoBuf/gen/go/protobuf/search/v1/searchv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func handleCORS(allowedOrigins []string, h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:      allowedOrigins,
		AllowedMethods:      []string{"POST"},
		AllowedHeaders:      append(connectcors.AllowedHeaders(), []string{"X-Access-Token"}...),
		AllowPrivateNetwork: true,
	}).Handler(h)
}

type Services struct {
	AuthService      *auth.AuthService
	ChatService      *chat.ChatService
	MessageService   *message.MessageService
	SearchService    *search.SearchService
	StreamSubscriber stream.StreamSubscriber
}

func NewGrpcServer(cfg *config.HTTP, router *http.ServeMux, services *Services) error {
	authInterceptor := interceptor.NewAuthInterceptor(services.AuthService)
	interceptors := connect.WithInterceptors(authInterceptor)

	authGrpcHandler := grpc_handlers.NewAuthGrpcHandler(services.AuthService)
	authGrpcRoute, authGrpcRouter := authv1connect.NewAuthServiceHandler(authGrpcHandler)

	chatGrpcHandler := grpc_handlers.NewChatGrpcHandler(log.NewSubLogger("chats-handler"), services.ChatService)
	chatGrpcRoute, chatGrpcRouter := chatv1connect.NewChatServiceHandler(chatGrpcHandler, interceptors)

	eventsGrpcHandler := grpc_handlers.NewEventsGrpcHandler(log.NewSubLogger("events-handler"), services.StreamSubscriber)
	eventsGrpcRoute, eventsGrpcRouter := eventsv1connect.NewEventsServiceHandler(eventsGrpcHandler, interceptors)

	messageGrpcHandler := grpc_handlers.NewMessageGrpcHandler(log.NewSubLogger("message-handler"), services.MessageService)
	messageGrpcRoute, messageGrpcRouter := messagev1connect.NewMessageServiceHandler(messageGrpcHandler, interceptors)

	searchGrpcHandler := grpc_handlers.NewSearchGrpcHandler(log.NewSubLogger("message-handler"), services.SearchService)
	searchGrpcRoute, searchGrpcRouter := searchv1connect.NewSearchServiceHandler(searchGrpcHandler, interceptors)

	router.Handle(authGrpcRoute, authGrpcRouter)
	router.Handle(chatGrpcRoute, chatGrpcRouter)
	router.Handle(eventsGrpcRoute, eventsGrpcRouter)
	router.Handle(messageGrpcRoute, messageGrpcRouter)
	router.Handle(searchGrpcRoute, searchGrpcRouter)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	handler := handleCORS(cfg.Cors.AllowOrigins, router)
	server := &http.Server{
		Addr:         addr,
		Handler:      h2c.NewHandler(handler, &http2.Server{}),
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}
	fmt.Println("Grpc server is listening at ", addr)
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
