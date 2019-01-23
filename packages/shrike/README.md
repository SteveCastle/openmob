# Shrike Server

Shrike is the API server for OpenMob. It provides a robust simple to use
platform for building social action apps.

## Build from source.

### Prerequisites

1. Linux or OSX, Windows is not currently supported.
2. Install and set up docker.
3. Install Go.
4. Add GOPATH to your cli path.

### Bootstrap

1. From the project root run `make bootstrap`.

### Build one step at a time.

The bootstrap command will do all of this for you but for more control you can
also do each step seperately.

1. Run `make startdb` to start a docker container running Postgres with PostGis
   extensions.
2. Run `make install` to install system Go dependencies to the go path.
3. Run `make init` to run the initial database migrations.
4. Run `make-generate` to rebuild the protobuffers.
5. Run `make build` to build the go binaries.

# Run the development server

1. Run `make start`. Rebuilds the binary and starts server.

# Run test client.

2. Run `make test-client` to run api test client.
