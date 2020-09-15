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
  SQL database. A lot of this data, and domain types, are relational but my justification for using a document store like mongo is that there are a number of data objects that have arbitrary property values based on class/race/equipment etc. that may not be predictable as updates to 5E are rolled out. Also a lot of the aggregate queries would be incredibly slow, potentially.

  Identity management: currently the plan is to have boilerplate auth flows within the graphql api. But in the future may consider splitting into a federation (since gqlgen has support for that with a config) and having a dedicated identity microservice OR having a dedicated idp but that is only if my wildest dreams come true.

  Base product goals: 
  - have public graphql api
  - sign up
  - log in
  - update user details
  - create a character
  - load and update character


