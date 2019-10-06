package api

import (
  "net/http"

  "po_projects/database"
  "po_projects/userService"
)

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  sessionID, found := getQueryElement(query, "sessionID")
  if !found {
    w.WriteHeader(400)
    return
  }

  projectID, found := getQueryElement(query, "projectID")
  if !found {
    w.WriteHeader(400)
    return
  }

  itemID, found := getQueryElement(query, "itemID")
  if !found {
    w.WriteHeader(400)
    return
  }


  project, err := database.FindProject_ID(projectID)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  user, err := userService.FetchUser(sessionID)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  if !project.IsOwner(user.ID) {
    w.WriteHeader(400)
    return
  }


  project.RemoveTodoPart(itemID)

  database.UpdateProject(project)

  w.WriteHeader(200)
}
