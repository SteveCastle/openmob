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

	// Data
	layoutTypes := []v1.CreateLayoutType{
		v1.CreateLayoutType{Title: "HomePage"},
	}

	componentImplementations := []v1.CreateComponentImplementation{
		v1.CreateComponentImplementation{Title: "SimpleHero", Path: "/heroes/SimpleHero"},
		v1.CreateComponentImplementation{Title: "SimpleFooter", Path: "/footers/SimpleFooter"},
		v1.CreateComponentImplementation{Title: "SimpleBoycott", Path: "/boycotts/SimpleBoycott"},
		v1.CreateComponentImplementation{Title: "SimpleContent", Path: "/content/SimpleContent"},
		v1.CreateComponentImplementation{Title: "SimpleDonationDrive", Path: "/donations/SimpleDonationDrive"},
		v1.CreateComponentImplementation{Title: "SimpleElection", Path: "/elections/SimpleElection"},
		v1.CreateComponentImplementation{Title: "SimplePetition", Path: "/petitions/SimplePetition"},
		v1.CreateComponentImplementation{Title: "SimpleShop", Path: "/shop/SimpleShop"},
		v1.CreateComponentImplementation{Title: "SimpleSignups", Path: "/signups/SimpleSignups"},
		v1.CreateComponentImplementation{Title: "SimpleVolunteering", Path: "/volunteers/SimpleVolunteering"},
		v1.CreateComponentImplementation{Title: "SimpleBlogPost", Path: "/blogs/SimpleBlogPost"},
	}

	componentTypes := []v1.CreateComponentType{
		v1.CreateComponentType{Title: "Simple Header"},
		v1.CreateComponentType{Title: "Simple Footer"},
		v1.CreateComponentType{Title: "SimpleBoycott"},
		v1.CreateComponentType{Title: "SimpleContent"},
		v1.CreateComponentType{Title: "SimpleDonationDrive"},
		v1.CreateComponentType{Title: "SimpleElection"},
		v1.CreateComponentType{Title: "SimplePetition"},
		v1.CreateComponentType{Title: "SimpleShop"},
		v1.CreateComponentType{Title: "SimpleSignups"},
		v1.CreateComponentType{Title: "SimpleVolunteering"},
		v1.CreateComponentType{Title: "SimpleBlogPost"},
	}

	// Create Layout Types
	var layoutTypeID string
	for _, item := range layoutTypes {
		layoutTypeReq := v1.CreateLayoutTypeRequest{
			Api:  apiVersion,
			Item: &item,
		}
		layoutTypeRes, err := c.CreateLayoutType(ctx, &layoutTypeReq)
		if err != nil {
			log.Fatalf("Create failed: %v", err)
		}
		log.Printf("Create layoutType Result: <%+v>\n\n", layoutTypeRes)

		layoutTypeID = layoutTypeRes.ID
	}
	// Create a ComponentType
	var componentTypeID string
	for _, item := range componentTypes {
		componentTypeReq := v1.CreateComponentTypeRequest{
			Api:  apiVersion,
			Item: &item,
		}
		componentTypeRes, err := c.CreateComponentType(ctx, &componentTypeReq)
		if err != nil {
			log.Fatalf("Create failed: %v", err)
		}
		log.Printf("Create componentType Result: <%+v>\n\n", componentTypeRes)

		componentTypeID = componentTypeRes.ID
	}
	// Create ComponentImplementations
	var componentImplementationID string
	for _, item := range componentImplementations {
		item.ComponentType = componentTypeID
		componentImplementationReq := v1.CreateComponentImplementationRequest{
			Api:  apiVersion,
			Item: &item,
		}
		componentImplementationRes, err := c.CreateComponentImplementation(ctx, &componentImplementationReq)
		if err != nil {
			log.Fatalf("Create failed: %v", err)
		}
		log.Printf("Create componentImplementation Result: <%+v>\n\n", componentImplementationRes)

		componentImplementationID = componentImplementationRes.ID
	}

	// Create a FieldType
	fieldTypeReq := v1.CreateFieldTypeRequest{
		Api: apiVersion,
		Item: &v1.CreateFieldType{
			Title:    "Title",
			DataType: "string",
			PropName: "title",
		},
	}
	fieldTypeRes, err := c.CreateFieldType(ctx, &fieldTypeReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create fieldType Result: <%+v>\n\n", fieldTypeRes)
	fieldTypeID := fieldTypeRes.ID

	// Cause specific data.
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
			Layout:    layoutID,
			Container: false,
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

	// Create a Photo
	photoRequest := v1.CreatePhotoRequest{
		Api: apiVersion,
		Item: &v1.CreatePhoto{
			URI:    "http://placeimg.com/400/400/nature",
			Width:  400,
			Height: 400,
		},
	}
	photoRes, err := c.CreatePhoto(ctx, &photoRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create photo Result: <%+v>\n\n", photoRes)
	photoID := photoRes.ID

	// Create a Cause
	causeTitle := faker.Company().Name()
	causeRequest := v1.CreateCauseRequest{
		Api: apiVersion,
		Item: &v1.CreateCause{
			Title:    causeTitle,
			Slug:     slugify.Slugify(causeTitle),
			Summary:  faker.Company().CatchPhrase(),
			HomePage: homepageID,
			Photo:    photoID,
		},
		BuildStatic: true,
	}
	res, err := c.CreateCause(ctx, &causeRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)
}
