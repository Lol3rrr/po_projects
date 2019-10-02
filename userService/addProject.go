package userService

import (
  "net/http"
)

func AddProject(sessionID, projectID, projectName string) (bool, error) {
  reqUrl := baseApiURL + "projects/add"

  req, err := http.NewRequest("POST", reqUrl, nil)
  if err != nil {
    return false, err
  }

  q := req.URL.Query()
  q.Add("sessionID", sessionID)
  q.Add("project_id", projectID)
  q.Add("project_name", projectName)
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
