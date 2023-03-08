package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// Define a HTTP server to serve a health check endpoint
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	// Start the HTTP server in a separate goroutine
	go func() {
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("failed to start http server: %v", err)
		}
	}()

	// Create a gRPC server and register the Greeter service
	grpcServer := grpc.NewServer()
	RegisterGreeterServer(grpcServer, &server{})

	// Define a listener for the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the gRPC server in a separate goroutine
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	// Wait for a signal to stop the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Stop the gRPC server gracefully
	log.Println("Stopping gRPC server...")
	grpcServer.GracefulStop()

	// Stop the HTTP server gracefully
	log.Println("Stopping HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("failed to stop http server: %v", err)
	}
	log.Println("Server stopped")
}