package api

import (
  "net/http"

  "po_projects/database"
  "po_projects/userService"
)

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

  worked, err := database.DeleteProject(id)
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
