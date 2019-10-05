package api

import (
  "net/http"
  "encoding/json"

  guuid "github.com/google/uuid"

  "po_projects/general"
  "po_projects/database"
  "po_projects/userService"
)

type AddListResponse struct {
  ID string `json:"id"`
}

type AddListRequest struct {
  Name string `json:"name"`
  Items []string `json:"items"`
}

func getListBody(req *http.Request) (AddListRequest, error) {
  var reqBody AddListRequest

  defer req.Body.Close()
  decoder := json.NewDecoder(req.Body)

  err := decoder.Decode(&reqBody)
  if err != nil {
    return reqBody, err
  }

  return reqBody, nil
}

func saveListHandler(w http.ResponseWriter, r *http.Request) {
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

  reqBody, err := getListBody(r)
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

  parts := make([]general.ListPoint, 0)
  for _, tmpPart := range reqBody.Items {
    tmp := general.ListPoint {
      Content: tmpPart,
    }

    parts = append(parts, tmp)
  }

  listPart := general.Project_List_Part {
    ID: itemID,
    Name: reqBody.Name,
    Parts: parts,
  }

  project.AddListPart(listPart)

  database.UpdateProject(project)

  resp := AddListResponse{
    ID: itemID,
  }

  sendSuccessResult(resp, w)
}
