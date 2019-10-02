package api

import (
  "net/http"
  "encoding/json"

  guuid "github.com/google/uuid"

  "po_projects/general"
  "po_projects/database"
  "po_projects/userService"
)

type CreateResponse struct {
  ID string `json:"id"`
}

func createHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  sessionID, worked := getSessionID(query)
  if !worked {
    w.WriteHeader(400)
    return
  }

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

  worked, err := userService.AddProject(sessionID, id, name)
  if err != nil || !worked {
    w.WriteHeader(400)
    return
  }

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
