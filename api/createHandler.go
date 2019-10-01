package api

import (
  "fmt"
  "net/http"
  "encoding/json"

  
)

type CreateResponse struct {
  ID string `json:"id"`
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  fmt.Printf("Query: %+v \n", query)

  rawName, ok := query["name"]
  if !ok || len(rawName) <= 0 {
    w.WriteHeader(400)

    return
  }

  name := rawName[0]

  fmt.Printf("Name: %+v \n", name)

  resp := CreateResponse{
    ID: "test",
  }

  jsonResponse, err := json.Marshal(resp)
  if err != nil {
    w.WriteHeader(400)

    return
  }

  w.WriteHeader(200)
  w.Write(jsonResponse)
}
