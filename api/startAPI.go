package api

import (
  "net/http"

  "github.com/gorilla/mux"
)

func StartAPI(port string) {
  r := mux.NewRouter()
  r.HandleFunc("/create", createHandler).Methods("POST")
  r.HandleFunc("/delete", deleteHandler).Methods("POST")

  r.HandleFunc("/save/text", saveTextHandler).Methods("POST")
  r.HandleFunc("/delete/text", deleteTextHandler).Methods("POST")

  r.HandleFunc("/save/list", saveListHandler).Methods("POST")
  r.HandleFunc("/delete/list", deleteListHandler).Methods("POST")

  r.HandleFunc("/save/todo", saveTodoHandler).Methods("POST")
  r.HandleFunc("/delete/todo", deleteTodoHandler).Methods("POST")

  r.HandleFunc("/load", loadHandler).Methods("GET")
  http.Handle("/", r)

  http.ListenAndServe(port, nil)
}
