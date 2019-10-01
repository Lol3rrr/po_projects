package api

import (
  "fmt"
  "net/http"
  "encoding/json"

  guuid "github.com/google/uuid"

  "po_projects/general"
  "po_projects/database"
)

type CreateResponse struct {
  ID string `json:"id"`
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  fmt.Printf("Query: %+v \n", query)

  rawName, ok := query["name"]
  if !ok || len(rawName) <= 0 {
    w.WriteHeader(400)

    return
  }

  name := rawName[0]
  id := guuid.New().String()

  project := general.Project{
    ID: id,
    Name: name,
  }

  database.UpdateProject(project)

  // TODO: Also add the Project to the Users Project List

  resp := CreateResponse{
    ID: id,
  }

  jsonResponse, err := json.Marshal(resp)
  if err != nil {
    w.WriteHeader(400)

    return
  }

  w.WriteHeader(200)
  w.Write(jsonResponse)
}
