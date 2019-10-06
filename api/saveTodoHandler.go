package api

import (
  "net/http"
  "encoding/json"

  guuid "github.com/google/uuid"

  "po_projects/general"
  "po_projects/database"
  "po_projects/userService"
)

type AddTodoResponse struct {
  ID string `json:"id"`
}

type AddTodoItem struct {
  Done    bool   `json:"done"`
  Content string `json:"content"`
}

type AddTodoRequest struct {
  Name  string        `json:"name"`
  Items []AddTodoItem `json:"items"`
}

func getTodoBody(req *http.Request) (AddTodoRequest, error) {
  var reqBody AddTodoRequest

  defer req.Body.Close()
  decoder := json.NewDecoder(req.Body)

  err := decoder.Decode(&reqBody)
  if err != nil {
    return reqBody, err
  }

  return reqBody, nil
}

func saveTodoHandler(w http.ResponseWriter, r *http.Request) {
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

  reqBody, err := getTodoBody(r)
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

  parts := make([]general.TodoPoint, 0)
  for _, tmpPart := range reqBody.Items {
    tmp := general.TodoPoint {
      Done: tmpPart.Done,
      Content: tmpPart.Content,
    }

    parts = append(parts, tmp)
  }

  todoPart := general.Project_Todo_Part {
    ID: itemID,
    Name: reqBody.Name,
    Parts: parts,
  }

  project.AddTodoPart(todoPart)

  database.UpdateProject(project)

  resp := AddTodoResponse{
    ID: itemID,
  }

  sendSuccessResult(resp, w)
}
