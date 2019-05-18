# OpenMob

OpenMob is a free open source web framework for the modern civil society
organization. Designed spefically with the needs of non profits, political organizations, and activist organizations in mind it includes easy to use out of the box features for fund raising, campaigning, volunteer coordination, and political messaging.
## Project Goals

1. Meet the needs of organizations to provide world class internet marketing tools to their employees and volunteers.
2. Lower organizational costs with an easy to use and low maintence system.
3. Avoid vendor lockin to provide maximum freedom to organizations wishing to use the software
4. Enable powerful experimentation and optimization so organizations can move fast and make measured improvements in their results.
5. Provide a fully extensible system meaning organizations can add on, or make changes to meet their unique needs.

## Technical Goals

1. Statically generate landing pages which makes them cheap and fast to host. This provides world class speed even on low cost devices and networks around the world, reduces costs, and increases fund raising conversion. 
2. All features of Openmob are plugin based so you can replace components at will.
3. Openmob is a containerized application to allow for the largest possible freedom in where you choose to run your OpenMob deployment while still allowing you to scale as your usage grows.
4. A command line tool to help you manage, deploy, and backup your installation.
5. Meet all web accessability guidelines to reach the largest possible audience.
6. Expose an API for your application so you can build on the core platform any way you want.

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
