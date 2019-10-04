package api

import (
  "net/http"

  "github.com/gorilla/mux"
)

func StartAPI(port string) {
  r := mux.NewRouter()
  r.HandleFunc("/create", createHandler).Methods("POST")
  r.HandleFunc("/delete", deleteHandler).Methods("POST")

  r.HandleFunc("/load", findHandler).Methods("GET")
  http.Handle("/", r)

  http.ListenAndServe(port, nil)
}
