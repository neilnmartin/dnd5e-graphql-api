# dnd5e-graphql-api
WIP: graphql api for DnD 5e built with Go

## Instructions to run locally: 
get, generate, build and run:

`make all`

### Description:
  A graphql api for Dungeons and Dragons 5th Edition written in Go. 
  Support for a web app client user to sign up with an email and create a character. A user will be able to keep track of a their DnD 5E character information and access information about 5E race, class, spells, equipment etc.

### Tools and packages: 
  - graphql-go
  - gqlgen for generation of schema
  - mgo for mongodb

### Considerations:

  Domain Driven: DDD is a design pattern that revolves around the definition of domain entities and aggregates, and the decoupling of outside dependencies. The scale of this service is still small enough not to require heavily domain driven architecture but basic ideas like decoupling of infrastructure and persistence layers are still useful. 

  SQL database. A lot of the data types are relational (i.e. SubRace -> Race, etc.) BUT there are a number of data objects that have arbitrary properties based on class/race/equipment etc. that may not be predictable as updates to 5E are rolled out. I need the schema to be easily mutable, therefore need a document store. Also a lot of the aggregate queries could be incredibly slow, potentially.

  Identity management: currently the plan is to have boilerplate auth flows within the graphql api. But in the future may consider splitting into a federation (since gqlgen has support for that with a config) and having a dedicated identity microservice.

  Deployment: Google App Engine? Terraform + AWS?
  
  Base product goals: 
  - have public graphql api for dungeons and dragons 5th edition information
  - basic signup/login auth functionality
  - user detail management
  - character creation
  - character details management


A list of query-able Dungeons & Dragons 5th Edition entities (not exhaustive):
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


[There is a tendency in tech to focus on the tool instead of the outcome](https://www.youtube.com/watch?v=GBTdnfD6s5Q)
So why would I choose to build a graphql server in Go, when one in TypeScript would be so much faster?
One of the goals of this project is precisely to become more familiar with a tool (in this case the Go programming language) and be able to explore writing idiomatic, efficient code. As a web developer, I to be comfy in with it and expand into writing CLI tools and perhaps contributing to open source Go projects.
