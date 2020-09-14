# dnd5e-graphql-api
graphql api for DnD 5e built with Go

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
  A graphql api for DnD 5E written in Go. Support for a web app client to keep track of a user's DnD 5E character information and provide information about race, class, spells etc.
  A user may use the site to browse information, or may sign up and log in to use the character creation/tracking tool

Considerations:
  SQL database. A lot of this data, and domain types, are relational but my justification for using a document store like mongo is that there are a number of data objects that have arbitrary property values based on class/race/equipment etc. that may not be predictable as updates to 5E are rolled out.

  Identity management: currently the plan is to have boilerplate auth flows within the graphql api. But in the future may consider splitting into a federation (since gqlgen has support for that with a config) and having a dedicated identity microservice OR having a dedicated idp but that is only if my wildest dreams come true.