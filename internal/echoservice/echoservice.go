package echoservice

import (
	"context"
	"log/slog"

	echov1 "github.com/Fco-Lucas/grpc-learn/gen/echo/v1"
)

type Server struct {
	echov1.UnimplementedEchoServiceServer
	logger *slog.Logger
}

func New(logger *slog.Logger) *Server {
	return &Server{logger: logger}
}

func (s *Server) SayHello(ctx context.Context, req *echov1.HelloRequest) (*echov1.HelloResponse, error) {
	s.logger.Info("SayHello called", "message", req.GetMessage())

	return &echov1.HelloResponse{
		Message: "Hello, " + req.GetMessage() + "!",
	}, nil
}
