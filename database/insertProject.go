package database

import (
  "context"

  "po_projects/general"
)

func insertProject(project general.Project) (error) {
  collection := client.Database(dataBaseName).Collection(collectionName)

  _, err := collection.InsertOne(context.TODO(), project)
  if err != nil {
    return err
  }

  return nil
}
