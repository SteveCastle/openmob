# OpenMob

A free and open source control center and platform for the modern civil society
organization. Empowers best practice fund raising, campaigning, and resource
coordination.

## Project Goals

1. OpenMob is a Web Framework designed specifically for non profits, political organizations, and service based organizations.
2. Fund raising, election management, awareness raising, and volunteer coordination are the core goals of an OpenMob site and every feature is designed to meet those goals as easily and expertly as possible.
3. OpenMob avoids vendor lock in and is easy to deploy on any cloud, datacenter, or self owned server. (See Minimum Specs)
4. Make testing and analytics a first class citizen to enable powerful experimentation and optimization.
5. Openmob is easy to use out of the box, but also fully extensible meaning you can add on, or make changes to meet your needs.

## Technical Goals

1. Public pages are statically generated which makes them cheap and fast to host. This provides world class speed even on low cost devices and networks around the world, reduces costs, and increases fund raising conversion. 
2. All features of Openmob are plugin based so you can replace components at will.
3. Openmob is a containerized application to allow for the largest possible freedom in where you choose to run your OpenMob deployment.
4. OpenMob comes with a command line tool to help you manage and deploy your installation.
5. OpenMob meets all web accessability guidelines to reach the largest possible audience.
6. OpenMob, exposes an API for your application so you can build on the core platform any way you want.

## Project Structure

                        ┌────────────────────┐
                        │                    │
                        │                    │        ┌────────────────────┐
                        │ Shrike gRPC Server │ ◀──┐   │                    │
                        │                    │    │   │  Mobile and Third  │
                        │                    │    └───│   Party Platform   │
                        └────────────────────┘        │       Access       │
                                   │                  │                    │
                                   ▼                  └────────────────────┘
                        ┌────────────────────┐
                        │                    │
                        │  Wren GraphQL API  │
                        │       Server       │
                        │                    │
                        │                    │
                        └────────────────────┘
                                   │
                                   │
    ┌────────────────────┐         └─────┐
    │                    │               │
    │   Bluebird React   │               ▼
    │ Component Library  │─┐  ┌────────────────────┐
    │ and Design System  │ │  │                    │
    │                    │ │  │ Starling Renderer  │
    └────────────────────┘ └─▶│     (GatsbyJS)     │
                              │                    │
                              │                    │
                              └────────────────────┘

Open Mob is a mono-repo comprised of three core packages. A gRPC API service
called
[Shrike](https://github.com/SteveCastle/openmob/tree/master/packages/shrike), a
GraphQL api server named
[Wren](https://github.com/SteveCastle/openmob/tree/master/packages/wren) and a
GatsbyJS based rendering engine for content and admin ui called
[Starling](https://github.com/SteveCastle/openmob/tree/master/packages/starling).
For more information on building these packages refer to the package level
README.

1. [Shrike](https://github.com/SteveCastle/openmob/tree/master/packages/shrike)
2. [Starling](https://github.com/SteveCastle/openmob/tree/master/packages/starling)
3. [Wren](https://github.com/SteveCastle/openmob/tree/master/packages/wren)

Other Open Source supporting packages like the
[Bluebird](https://github.com/SteveCastle/openmob/tree/master/packages/bluebird)
react component library and design system are also found in this repo.

## Codegen

Openmob makes heavy use of codegen to generate consistent typing from the database, through the server, all the
way through to the javascript client.

                                                                      ┌────────────────────────┐
                                                                      │                        │
                                                                      │                        │
                                                                      │                        │
                                                             ┌────────│  PostgreSQL DB Schema  │────────┐
                                                             │        │                        │        │
                                                             │        │                        │        │
                                                             │        │                        │        │
                                                             │        └────────────────────────┘        │
                                                             │                                          │
                                                             ▼                                          ▼
                                                ┌────────────────────────┐                 ┌────────────────────────┐
                                                │                        │                 │                        │
                                                │                        │                 │  GNORM - Go template   │
                                                │  GNORM - Go template   │                 │   based templater. A   │
                                                │    based templater.    │                 │ second GNORM pass with │─────────────────────┐
                                                │                        │                 │settings for Javascript │                     │
                                                │                        │                 │     type mappings.     │                     ▼
                                                │                        │                 │                        │         ┌───────────────────────┐
                                                └────────────────────────┘                 └────────────────────────┘         │                       │
                                                             │                                          │                     │                       │
                                                             │                                          │                     │                       │
                                                             │                                          └─────┐               │                       │
                                                             │                                                │               │   GraphQL Resolvers   │
                            ┌──────────────────────────┬─────┴────────────────────┐                           ▼               │                       │
                            │                          │                          │               ┌───────────────────────┐   │                       │
                            │                          │                          │               │                       │   │                       │
                            │                          │                          │               │                       │   │                       │
                            │                          │                          │               │                       │   └───────────────────────┘
                            ▼                          ▼                          ▼               │                       │
                ┌───────────────────────┐  ┌───────────────────────┐  ┌───────────────────────┐   │    GraphQL Schema     │
                │                       │  │                       │  │                       │   │                       │
                │                       │  │                       │  │                       │   │                       │
                │                       │  │api/proto/shrike.proto │  │                       │   │                       │
                │  /pkg/service - CRUD  │  │   - A Protobuf3 API   │  │   /pkg/query - SQL    │   │                       │
                │ Handlers for the gRPC │  │specification and CRUD │  │   Builders for CRUD   │   └───────────────────────┘
                │       Service.        │  │     gRPC service.     │  │       Handlers.       │
                │                       │  │                       │  │                       │
                │                       │  │                       │  │                       │
                │                       │  │                       │  │                       │
                └───────────────────────┘  └───────────────────────┘  └───────────────────────┘
