package database

import (
  "context"

  "go.mongodb.org/mongo-driver/bson"

  "po_projects/general"
)

func FindProject(projectID string) (general.Project, error) {
  collection := client.Database(dataBaseName).Collection(collectionName)
  var result general.Project

  filter := bson.D{{"id", projectID}}

  err := collection.FindOne(context.TODO(), filter).Decode(&result)
  if err != nil {
    return result, err
  }

  return result, nil
}
