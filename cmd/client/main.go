package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	echov1 "github.com/Fco-Lucas/grpc-learn/gen/echo/v1"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("failed to create client", "err", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := echov1.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &echov1.HelloRequest{
		Message: "Lucas",
	})
	if err != nil {
		logger.Error("SayHello failed", "err", err)
		os.Exit(1)
	}

	logger.Info("got response", "message", resp.GetMessage())
}
