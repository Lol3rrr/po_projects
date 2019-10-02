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

  rawSessionID, s_ok := query["sessionID"]
  if !s_ok || len(rawSessionID) <= 0 {
    w.WriteHeader(400)
    return
  }
  sessionID := rawSessionID[0]

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
