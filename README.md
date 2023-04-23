# dnd5e-graphql-api
WIP: graphql api for DnD 5e built with Go

## Instructions to run locally: 
get, generate, build and run:

`make all`

### Description:
  A GraphQL API for Dungeons and Dragons 5th Edition written in Go. 
  Support for a web app client user to sign up with an email and create a character. A user will be able to keep track of a their DnD 5E character information and access information about 5E race, class, spells, equipment etc.

### Tools and packages: 
  - graphql-go
  - [gqlgen](https://github.com/99designs/gqlgen) for schema-based generation of statically-typed models
  - mgo for mongodb

### Considerations:

  Domain Driven Design: a design pattern that revolves around the definition of domain entities and aggregates, and the decoupling of outside dependencies. The scale of this service is still small enough not to require heavily domain driven architecture but basic ideas like decoupling of infrastructure and persistence layers are still useful. 

  Document store. While a lot of the data is relational in nature (i.e. SubRace -> Race, etc.) there are a number of entities that can have arbitrary properties based on class/race/equipment etc. Also these could be updated as future versions of 5E are rolled out. Schema needs to be easily mutable.

  Identity management: boilerplate auth flows. May consider splitting into a federation (since gqlgen has support for that with a config) and having a dedicated identity microservice.

  Deployment: Google App Engine? Terraform + AWS? AWS Amplify also an option.
  
  Goals: 
  - public graphql api for dungeons and dragons 5th edition information
  - basic auth functionality
  - user profile management
  - character creation
  - character management


A public API of query-able Dungeons & Dragons 5th Edition entities (not exhaustive):
- Race
- Subrace
- Class
- ClassLevel
- Subclass
- Skills
- AbilityScores
- Equipment
  - Items
  - Weapons
- Armor
- Features

One of the goals of this project is to learn to write idiomatic, efficient Go. Hope to expand into writing my own CLI tools and contribute to open source down the line.
