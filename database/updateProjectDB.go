package database

import (
  "context"

  "go.mongodb.org/mongo-driver/bson"

  "po_projects/general"
)

func updateProjectDB(projectID string, project general.Project) (error) {
  collection := client.Database(dataBaseName).Collection(collectionName)

  filter := bson.D{{"id", projectID}}
  update := bson.D{{"$set", project}}

  _, err := collection.UpdateOne(context.TODO(), filter, update)
  if err != nil {
    return err
  }

  return nil
}
