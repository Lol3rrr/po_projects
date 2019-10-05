package api

import (
  "net/http"
  "encoding/json"

  guuid "github.com/google/uuid"

  "po_projects/general"
  "po_projects/database"
  "po_projects/userService"
)

type AddTextResponse struct {
  ID string `json:"id"`
}

type AddTextRequest struct {
  Name string `json:"name"`
  Content string `json:"content"`
}

func getTextBody(req *http.Request) (AddTextRequest, error) {
  var reqBody AddTextRequest

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

  reqBody, err := getTextBody(r)
  if err != nil {
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

  if itemID == "" {
    itemID = guuid.New().String()
  }

  textPart := general.Project_Text_Part {
    ID: itemID,
    Name: reqBody.Name,
    Content: reqBody.Content,
  }

  project.AddTextPart(textPart)

  database.UpdateProject(project)

  resp := AddTextResponse{
    ID: itemID,
  }

  sendSuccessResult(resp, w)
}