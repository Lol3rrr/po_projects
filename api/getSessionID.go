package api

func getSessionID(query map[string][]string) (string, bool) {
  rawSessionID, s_ok := query["sessionID"]
  if !s_ok || len(rawSessionID) <= 0 {
    return "", false
  }

  return rawSessionID[0], true
}
