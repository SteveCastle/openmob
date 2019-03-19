package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/api/v1"
	"github.com/jaswdr/faker"
	slugify "github.com/mozillazg/go-slugify"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	faker := faker.New()
	// get configuration
	address := flag.String("server", "localhost:9090", "gRPC server in format host:port")
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

	// Create a Layout Type
	layoutTypeReq := v1.CreateLayoutTypeRequest{
		Api: apiVersion,
		Item: &v1.CreateLayoutType{
			Title: "Homepage",
		},
	}
	layoutTypeRes, err := c.CreateLayoutType(ctx, &layoutTypeReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create layoutType Result: <%+v>\n\n", layoutTypeRes)

	layoutTypeID := layoutTypeRes.ID

	// Create a Layout
	layoutReq := v1.CreateLayoutRequest{
		Api: apiVersion,
		Item: &v1.CreateLayout{
			LayoutType: layoutTypeID,
		},
	}
	layoutRes, err := c.CreateLayout(ctx, &layoutReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create layout Result: <%+v>\n\n", layoutRes)

	layoutID := layoutRes.ID

	// Create a LayoutRow
	layoutRowReq := v1.CreateLayoutRowRequest{
		Api: apiVersion,
		Item: &v1.CreateLayoutRow{
			Layout: layoutID,
		},
	}
	layoutRowRes, err := c.CreateLayoutRow(ctx, &layoutRowReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create layoutRow Result: <%+v>\n\n", layoutRowRes)

	layoutRowID := layoutRowRes.ID

	// Create a LayoutColumn
	layoutColumnReq := v1.CreateLayoutColumnRequest{
		Api: apiVersion,
		Item: &v1.CreateLayoutColumn{
			LayoutRow: layoutRowID,
			Width:     12,
		},
	}
	layoutColumnRes, err := c.CreateLayoutColumn(ctx, &layoutColumnReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create layoutColumn Result: <%+v>\n\n", layoutColumnRes)

	layoutColumnID := layoutColumnRes.ID

	// Create a ComponentImplemntation
	componentImplementationReq := v1.CreateComponentImplementationRequest{
		Api: apiVersion,
		Item: &v1.CreateComponentImplementation{
			Title: "Component",
		},
	}
	componentImplementationRes, err := c.CreateComponentImplementation(ctx, &componentImplementationReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create componentImplementation Result: <%+v>\n\n", componentImplementationRes)

	componentImplementationID := componentImplementationRes.ID

	// Create a ComponentType
	componentTypeReq := v1.CreateComponentTypeRequest{
		Api: apiVersion,
		Item: &v1.CreateComponentType{
			Title: "Header",
		},
	}
	componentTypeRes, err := c.CreateComponentType(ctx, &componentTypeReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create componentType Result: <%+v>\n\n", componentTypeRes)

	componentTypeID := componentTypeRes.ID

	// Create a Component
	componentReq := v1.CreateComponentRequest{
		Api: apiVersion,
		Item: &v1.CreateComponent{
			ComponentType:           componentTypeID,
			ComponentImplementation: componentImplementationID,
			LayoutColumn:            layoutColumnID,
		},
	}
	componentRes, err := c.CreateComponent(ctx, &componentReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create component Result: <%+v>\n\n", componentRes)

	// Create a HomePage
	homepageReq := v1.CreateHomePageRequest{
		Api: apiVersion,
		Item: &v1.CreateHomePage{
			Layout: layoutID,
		},
	}
	homepageRes, err := c.CreateHomePage(ctx, &homepageReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create homepage Result: <%+v>\n\n", homepageRes)

	homepageID := homepageRes.ID

	// Create a Cause
	causeTitle := faker.Company().Name()
	causeRequest := v1.CreateCauseRequest{
		Api: apiVersion,
		Item: &v1.CreateCause{
			Title:    causeTitle,
			Slug:     slugify.Slugify(causeTitle),
			Summary:  faker.Company().CatchPhrase(),
			HomePage: homepageID,
		},
	}
	res, err := c.CreateCause(ctx, &causeRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)
}
