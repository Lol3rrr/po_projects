package api

import (
  "net/http"
  "encoding/json"

  "po_projects/general"
  "po_projects/database"
)

type FindResponse struct {
  Project general.Project `json:"project"`
}

func sendResult(result FindResponse, w http.ResponseWriter) {
  jsonResponse, err := json.Marshal(result)
  if err != nil {
    w.WriteHeader(400)

    return
  }

  w.WriteHeader(200)
  w.Write(jsonResponse)
}

func findUsingID(id string, w http.ResponseWriter) {
  project, err := database.FindProject_ID(id)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  resp := FindResponse{
    Project: project,
  }

  sendResult(resp, w)
}


func loadHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  rawID, id_ok := query["id"]
  if id_ok && len(rawID) > 0 {
    findUsingID(rawID[0], w)
    return
  }
}
