package api

import (
  "net/http"

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

  user, err := userService.FetchUser(sessionID)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  owner := general.Project_Owner{
    ID: user.ID,
  }

  project := general.Project{
    ID: id,
    Name: name,
    Owner: owner,
  }

  database.UpdateProject(project)

  worked, err = userService.AddProject(sessionID, id, name)
  if err != nil || !worked {
    w.WriteHeader(400)
    return
  }

  resp := CreateResponse{
    ID: id,
  }

  sendSuccessResult(resp, w)
}
