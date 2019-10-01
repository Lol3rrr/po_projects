package api

import (
  "net/http"

  "github.com/gorilla/mux"
)

func StartAPI(port string) {
  r := mux.NewRouter()
  r.HandleFunc("/create", CreateHandler).Methods("POST")
  http.Handle("/", r)

  http.ListenAndServe(port, nil)
}
