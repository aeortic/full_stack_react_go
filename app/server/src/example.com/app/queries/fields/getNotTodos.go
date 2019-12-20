package queries

import (
  "context"
  "github.com/graphql-go/graphql"
  "github.com/mongodb/mongo-go-driver/bson"

  "example.com/app/data"
  types "example.com/app/types"
)

type todoStruct struct {
  NAME string `json:"name"`
  DESCRIPTION string `json:"description"`
}

var GetNotTodos = &graphql.Field {
  Type: graphql.NewList(types.NotTodo),
  Description: "Get all not todos",
  Resolve: func(params graphql.ResolveParams) (interface{}, error) {
    notTodoCollection := mongo.Client.Database("medium-app").Collection("Not_Todos")

    todos, err := notTodoCollection.Find(context.Background(), bson.D{})
    if err != nil { panic(err) }

    var todosList []todoStruct

    for todos.Next(context.Background()) {
      var doc bson.M

      err := todos.Decode(&doc)
      if err != nil { panic(err)}

      // convert BSON to struct
      todo := todoStruct{}
      for key := range doc {

        value := doc[key].(string)

        switch (key) {
          case "name":
            todo.NAME = value
          case "description":
            todo.DESCRIPTION = value
          default:
        }
      }
      todosList = append(todosList, todo)
    }

    return todosList, nil
  },
}
