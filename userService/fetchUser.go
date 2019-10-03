package userService

import (
  "errors"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func FetchUser(sessionID string) (User, error) {
  var result User

  reqUrl := baseApiURL + "load"

  req, err := http.NewRequest("GET", reqUrl, nil)
  if err != nil {
    return result, err
  }

  q := req.URL.Query()
  q.Add("sessionID", sessionID)
  req.URL.RawQuery = q.Encode()

  client := http.Client{
    Timeout: defaultTimeout,
  }

  resp, err := client.Do(req)
  if err != nil {
    return result, err
  }

  if resp.StatusCode != 200 {
    return result, errors.New("Could not load User")
  }

  defer resp.Body.Close()
  content, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return result, err
  }

  type ResponseUser struct {
    User User `json:"user"`
  }
  var respUser ResponseUser

  err = json.Unmarshal(content, &respUser)
  if err != nil {
    return result, err
  }

  result = respUser.User

  return result, nil
}
