package api

import (
  "net/http"

  "po_projects/database"
  "po_projects/userService"
)

func doesUserOwnProject(sessionID, projectID string) (bool, error) {
  user, err := userService.FetchUser(sessionID)
  if err != nil {
    return false, err
  }

  project, err := database.FindProject_ID(projectID)
  if err != nil {
    return false, err
  }

  if project.Owner.ID != user.ID {
    return false, nil
  }

  return true, nil
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  sessionID, worked := getSessionID(query)
  if !worked {
    w.WriteHeader(400)
    return
  }

  rawID, ok := query["id"]
  if !ok || len(rawID) <= 0 {
    w.WriteHeader(400)
    return
  }

  id := rawID[0]

  owns, err := doesUserOwnProject(sessionID, id)
  if err != nil || !owns {
    w.WriteHeader(400)
    return
  }

  worked, err = database.DeleteProject(id)
  if err != nil || !worked {
    w.WriteHeader(400)
    return
  }

  worked, err = userService.RemoveProject(sessionID, id)
  if err != nil || !worked {
    w.WriteHeader(400)
    return
  }

  w.WriteHeader(200)
}
