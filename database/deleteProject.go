package database

import (
  "context"

  "go.mongodb.org/mongo-driver/bson"
)

func DeleteProject(id string) (bool, error)  {
  collection := client.Database(dataBaseName).Collection(collectionName)

  filter := bson.D{{"id", id}}

  deleteResult, err := collection.DeleteOne(context.TODO(), filter)
  if err != nil {
    return false, err
  }

  if deleteResult.DeletedCount <= 0 {
    return false, nil
  }

  return true, nil
}
