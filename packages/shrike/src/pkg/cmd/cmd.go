package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	// mysql driver
	_ "github.com/lib/pq"

	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/protocol/grpc"
	v1 "github.com/SteveCastle/openmob/packages/shrike/src/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "5432", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "tern", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "tern", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "tern", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "tern", "Database schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}
	// Datbase connection string.
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	v1API := v1.NewCauseServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
