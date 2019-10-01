package database

import (
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (error) {
  // Set client options
  authUrl := "mongodb://" + dbURL + ":27017"
  clientOptions := options.Client().ApplyURI(authUrl)

  // Connect to MongoDB
  var err error
  client, err = mongo.Connect(context.TODO(), clientOptions)
  if err != nil {
    return err
  }

  // Check the connection
  err = client.Ping(context.TODO(), nil)
  if err != nil {
    return err
  }

  return nil
}
