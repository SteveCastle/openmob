package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"

	"google.golang.org/grpc"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/interceptors"
	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
)

// RunServer runs gRPC service to publish Cause service
func RunServer(ctx context.Context, v1API v1.ShrikeServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service with Interceptors
	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer()),
    	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			interceptors.BuildInterceptor,
			interceptors.LoggingInterceptor,
			),
		),
	)
	v1.RegisterShrikeServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}