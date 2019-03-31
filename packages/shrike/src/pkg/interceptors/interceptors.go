package interceptors

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SteveCastle/structs"
	"google.golang.org/grpc"
)

type server struct{}

//LoggingInterceptor logs requests to standard output.
func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// Calls the handler
	h, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}
	// Logging with grpclog (grpclog.LoggerV2)
	fmt.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

// BuildInterceptor triggers gatsby builds if required
// by hitting a webhook for the build pipeline defined in config.
func BuildInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Calls the handler
	h, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	// Translate the request in to a map so we can safely check for BuildStatic key.
	s := structs.Map(req)
	// If BuildStatic is set and equals true, call the build webhook from config.
	if s["BuildStatic"] == true {
		fmt.Printf("Build requested: %v\n", s["BuildStatic"])
		message := map[string]interface{}{}

		webhookMessage, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
		}

		resp, err := http.Post("http://localhost:8000/__refresh", "application/json", bytes.NewBuffer(webhookMessage))
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Build finished with response: %v\n", resp)
	}

	return h, err
}
