package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewShrikeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	req1 := v1.CreateCauseRequest{
		Api: apiVersion,
		Item: &v1.Cause{
			Title: "My Cool Cause",
		},
	}
	res1, err := c.CreateCause(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	id := res1.Id

	// Read
	req2 := v1.GetCauseRequest{
		Api: apiVersion,
		Id:  id,
	}
	res2, err := c.GetCause(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res2)

	// List
	req3 := v1.ListCauseRequest{
		Api:   apiVersion,
		Limit: 1,
	}
	res3, err := c.ListCause(ctx, &req3)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res3)
}
