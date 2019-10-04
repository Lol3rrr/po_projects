package api

func getQueryElement(query map[string][]string, key string) (string, bool) {
  result, found := query[key]
  if !found || len(result) <= 0 {
    return "", false
  }

  return result[0], true
}
