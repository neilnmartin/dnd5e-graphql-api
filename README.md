# dnd5e-graphql-api
WIP: graphql api for DnD 5e built with Go

instructions to run locally: 
  - git clone
  - go get
  - make generate
  - make start

Tools and packages: 
  - graphql-go
  - gqlgen for generation of schema
  - mongodb (current choice of persistence pending any issues with schema)

Description:
  A graphql api for DnD 5E written in Go. 
  Support for a web app client user to sign up with an email and create a character. User will be able to keep track of a their DnD 5E character information and access information about 5E race, class, spells, equipment etc.

Considerations:
  Domain Driven: the scale of this service is still small enough not to require heavy domain driven architecture but basic ideas like decoupling of infrastructure and persistence layers are still attractive. 

  SQL database. A lot of the data types are relational BUT there are a number of data objects that have arbitrary properties based on class/race/equipment etc. that may not be predictable as updates to 5E are rolled out. I need the schema to be easily mutable, therefore need a document store. Also a lot of the aggregate queries could be incredibly slow, potentially.

  Identity management: currently the plan is to have boilerplate auth flows within the graphql api. But in the future may consider splitting into a federation (since gqlgen has support for that with a config) and having a dedicated identity microservice OR having a dedicated idp but that is only if my wildest dreams come true.

  Deployment: Google App Engine?
  
  Base product goals: 
  - have public graphql api
  - sign up
  - log in
  - update user details
  - create a character
  - load and update character


