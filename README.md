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

## Project Structure

```
                                                               +----------------+
                    +--------------------------------+         |                |
                    |Shrike gRPC Server              |         |                |
                    |                                +--------->Mobile          |
                    +---------------+----------------+         |and Third Party |
                                    |                          |Platform Access |
                                    |                          +-------^--------+
                                    v                                  |
                    +---------------+-----------------+                |
                    |Wren GraphQL API Server          +----------------+
                    |                                 |
                    +---------------+-----------------+
                                    |
+--------------+                    |
|Bluebird      |    +---------------+-----------------+
|Component     |    |Starling Renderer: GatsbyJS      |
|Library       +---->                                 |
|              |    +---------------------------------+
|              |
+--------------+

```

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
