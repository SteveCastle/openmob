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
			Title: "Simple Hero",
			Path:  "/heroes/SimpleHero",
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
	componentID := componentRes.ID

	// Create a FieldType
	fieldTypeReq := v1.CreateFieldTypeRequest{
		Api: apiVersion,
		Item: &v1.CreateFieldType{
			Title:         "Title",
			DataType:      "string",
			PropName:      "title",
			ComponentType: componentTypeID,
		},
	}
	fieldTypeRes, err := c.CreateFieldType(ctx, &fieldTypeReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create fieldType Result: <%+v>\n\n", fieldTypeRes)
	fieldTypeID := fieldTypeRes.ID
	// Create a Field
	fieldReq := v1.CreateFieldRequest{
		Api: apiVersion,
		Item: &v1.CreateField{
			FieldType:   fieldTypeID,
			Component:   componentID,
			StringValue: "My HomePage",
		},
	}
	fieldRes, err := c.CreateField(ctx, &fieldReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create field Result: <%+v>\n\n", fieldRes)

	// Create a HomePage
	homepageReq := v1.CreateHomePageRequest{
		Api: apiVersion,
		Item: &v1.CreateHomePage{
			Layout: layoutID,
			Title:  "My Homepage",
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
		BuildStatic: true,
	}
	res, err := c.CreateCause(ctx, &causeRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)
}
