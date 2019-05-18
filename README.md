*This project is in an early development stage and should only be installed by developers. Visit https://grassroots.dev or follow @grassrootsdev on twitter for updates on the progress and upcoming release date.*

# OpenMob

OpenMob is a free open source web framework for the modern civil society
organization. Designed specifically with the needs of non profits, political organizations, and activist organizations in mind it includes easy to use out of the box features for fund raising, campaigning, volunteer coordination, and political messaging.

## What does the name OpenMob mean?
In animal behavior, [mobbing](https://en.wikipedia.org/wiki/Mobbing_(animal_behavior)) is a form of defense used by smaller prey animals(especially birds) against larger predators. The smaller animals band together and harass the predator until it decides the meal is no longer worth the hassle. Since a fundamental goal of this project is to democratize cutting edge political technology by making it more freely available to anyone with the passion and energy to pursue a civic cause or goal, we thought it was an apt metaphor.

## Project Goals

1. Meet the needs of organizations to provide world class information technology tools to their employees and volunteers.
2. Lower organizational costs with an easy to use and low maintenance system.
3. Avoid vendor lock-in to provide maximum freedom to organizations wishing to use the software
4. Enable powerful experimentation and optimization so organizations can move fast and make measured improvements in their results.
5. Provide a fully extensible system allowing organizations to add on, or make changes to meet their unique needs.

## Technical Goals

1. Statically generate landing pages which makes them cheap and fast to host. This provides world class speed even on low cost devices and networks around the world, reduces costs, and increases fund raising conversion. 
2. All features of OpenMob are plugin based so you can replace components at will.
3. OpenMob is a containerized application to allow for the largest possible freedom in where you choose to run your OpenMob deployment while still allowing you to scale as your usage grows.
4. A command line tool to help you manage, deploy, and backup your installation.
5. Meet all web accessability guidelines to reach the largest possible audience.
6. Expose an API for your application so you can build on the core platform any way you want.

## Developing 
### Dependencies
You will need the following dependencies to build and run the project.
1. GO
2. Docker
3. Postgres
### Installation
1. Clone the repository `git clone https://github.com/SteveCastle/openmob.git`
2. Enter the project root directory `cd openmob`
3. Run `make bootstrap` to install all of the go dependencies and build the project.
4. See the make file for additional make commands you can use.

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
react component library and design system are also found in this repository.

## Code Generation

OpenMob makes heavy use of code generation to generate code for common tasks by introspecting the database and outputting database CRUD code, and API handlers. This helps keep the codebase simple, easy to understand, fast to build on, and reliable.

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
