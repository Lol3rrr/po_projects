package userService

import (
  "time"
)

const baseApiURL = "http://users:8080/"
const defaultTimeout = time.Duration(5 * time.Second)

type User struct {
  ID    string `json:"id"`
  Name  string `json:"name"`
}
