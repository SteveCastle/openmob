# OpenMob

A free and open source control center and platform for the modern civil society
organization. Empowers best practice fund raising, campaigning, and resource
coordination.

## Project Goals

1. A Web platform for non profits and socially beneficial organizations.
2. Make best practices for e-commerce, fund raising, lead generation, and
   Contact Management available for free.
3. Provide all the features required for an organization to raise funds,
   increase participation, and raise awareness out of the box.
4. Provide a means to coordinate volunteers for sales and order fulfillment.
5. Make testing and analytics a first class citizen.
6. Do not require users to lock in to any third-party service.

## Technical Goals

1. Statically build landing pages for fastest possible response times,
   accessability in low bandwidth regions, and low hosting costs.
2. Plugin based hosting, phone contact, and payment fullfilment to avoid
   corporate lock in.
3. Simple lock in free self deployment.
4. One command CLI deploy and application manager.
5. Provide best practice accessibility out of the box.

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
