package userService

import (
  "net/http"
)

func RemoveProject(sessionID, projectID string) (bool, error) {
  reqUrl := baseApiURL + "projects/delete"

  req, err := http.NewRequest("POST", reqUrl, nil)
  if err != nil {
    return false, err
  }

  q := req.URL.Query()
  q.Add("sessionID", sessionID)
  q.Add("project_id", projectID)
  req.URL.RawQuery = q.Encode()

  client := http.Client{
    Timeout: defaultTimeout,
  }

  resp, err := client.Do(req)
  if err != nil {
    return false, err
  }

  if resp.StatusCode != 200 {
    return false, nil
  }

  return true, nil
}
