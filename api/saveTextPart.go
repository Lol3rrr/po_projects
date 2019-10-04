package api

import (
  "net/http"
  "encoding/json"

  guuid "github.com/google/uuid"

  "po_projects/general"
  "po_projects/database"
  "po_projects/userService"
)

type AddResponse struct {
  ID string `json:"id"`
}

type AddRequest struct {
  Name string `json:"name"`
  Content string `json:"content"`
}

func getBody(req *http.Request) (AddRequest, error) {
  var reqBody AddRequest

  defer req.Body.Close()
  decoder := json.NewDecoder(req.Body)

  err := decoder.Decode(&reqBody)
  if err != nil {
    return reqBody, err
  }

  return reqBody, nil
}

func saveTextHandler(w http.ResponseWriter, r *http.Request) {
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

  reqBody, err := getBody(r)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  id := guuid.New().String()

  textPart := general.Project_Text_Part {
    ID: id,
    Name: reqBody.Name,
    Content: reqBody.Content,
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

  if project.TextParts == nil {
    project.TextParts = make([]general.Project_Text_Part, 0)
  }

  updated := false
  for _, tmpPart := range project.TextParts {
    if tmpPart.ID == itemID {
      id = tmpPart.ID
      tmpPart.Name = reqBody.Name
      tmpPart.Content = reqBody.Content

      updated = true
    }
  }

  if !updated {
    project.TextParts = append(project.TextParts, textPart)
  }

  database.UpdateProject(project)

  resp := AddResponse{
    ID: id,
  }

  sendSuccessResult(resp, w)
}
