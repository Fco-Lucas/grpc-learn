package main

import (
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	echov1 "github.com/Fco-Lucas/grpc-learn/gen/echo/v1"
	"github.com/Fco-Lucas/grpc-learn/internal/echoservice"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	const addr = ":50051"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("failed to listen", "err", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	echov1.RegisterEchoServiceServer(grpcServer, echoservice.New(logger))

	go func() {
		logger.Info("gRPC server listening", "addr", addr)
		if err := grpcServer.Serve(listener); err != nil {
			logger.Error("server error", "err", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down gracefully")
	grpcServer.GracefulStop()
}
