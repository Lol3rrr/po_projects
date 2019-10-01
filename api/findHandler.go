package api

import (
  "net/http"
  "encoding/json"

  "po_projects/general"
  "po_projects/database"
)

type FindResponse struct {
  Results []general.Project `json:"results"`
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

  result := make([]general.Project, 1)
  result[0] = project

  resp := FindResponse{
    Results: result,
  }

  sendResult(resp, w)
}

func findUsingName(name string, w http.ResponseWriter) {
  resp := FindResponse{
    Results: make([]general.Project, 0),
  }

  sendResult(resp, w)
}


func findHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  rawID, id_ok := query["id"]
  rawName, name_ok := query["name"]

  if id_ok && len(rawID) > 0 {
    findUsingID(rawID[0], w)
    return
  }

  if name_ok && len(rawName) > 0 {
    findUsingName(rawName[0], w)
    return
  }
}
