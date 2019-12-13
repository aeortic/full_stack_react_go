package types

import (
  "github.com/graphql-go/graphql"
)

var NotToDo = graphql.NewObject(graphql.ObjectConfig {
  Name: "NotToDo",
  Fields: graphql.Fields{
    "name": &graphql.Field{
      Type: graphql.String,
    },
    "description": &graphql.Field{
      Type: graphql.String,
    },
  },
})

