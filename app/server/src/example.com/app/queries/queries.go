package queries

import (
  "github.com/graphql-go/graphql"
  fields "example.com/app/queries/fields"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootQuery",
  Fields: graphql.Fields{
    "getNotTodos": fields.GetNotTodos,
  },
})

