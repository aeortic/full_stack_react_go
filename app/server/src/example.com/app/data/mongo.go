package mongo

import (
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions = options.Client().ApplyURI("mongodb+srv://aeortic:P1nchamyrasan@@cluster0-qcc6b.mongodb.net/test?retryWrites=true&w=majority")

var Client, err = mongo.Connect(context.Background(), clientOptions)
