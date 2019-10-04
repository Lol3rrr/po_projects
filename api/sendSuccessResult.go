package api

import (
  "net/http"
  "encoding/json"
)

func sendSuccessResult(content interface{}, w http.ResponseWriter) {
  jsonResponse, err := json.Marshal(content)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  w.WriteHeader(200)
  w.Write(jsonResponse)
}
