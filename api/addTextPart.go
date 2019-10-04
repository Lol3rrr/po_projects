package api

import (
  "net/http"
  "encoding/json"

  guuid "github.com/google/uuid"

  "po_projects/general"
  "po_projects/database"
)

type AddResponse struct {
  ID string `json:"id"`
}

type AddRequest struct {
  Name string `json:"name"`
  Content string `json:"content"`
}

func addTextHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  sessionID, found := getQueryElement(query, "sessionID")
  if !found {
    w.WriteHeader(400)
    return
  }

  projectID, found := getQueryElement(query, "id")
  if !found {
    w.WriteHeader(400)
    return
  }

  defer r.Body.Close()
  decoder := json.NewDecoder(r.Body)
  var reqBody AddRequest
  err := decoder.Decode(&reqBody)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  textPart := general.Project_Text_Part {
    ID: guuid.New().String(),
    Name: reqBody.Name,
    Content: reqBody.Content,
  }

  project, err := database.FindProject_ID(projectID)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  if project.TextParts == nil {
    project.TextParts = make([]Project_Text_Part, 0)
  }

  project.TextParts = append(project.TextParts, textPart)

  database.UpdateProject(project)

  resp := AddResponse{
    ID: textPart.ID,
  }

  jsonResponse, err := json.Marshal(resp)
  if err != nil {
    w.WriteHeader(400)

    return
  }

  w.WriteHeader(200)
  w.Write(jsonResponse)
}
