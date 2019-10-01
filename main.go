package main

import (
  "os"
  "fmt"
  "strings"

  "po_projects/api"
  "po_projects/database"
)

func parseEnvVariables() (map[string]string) {
  result := make(map[string]string)

  for _, e := range os.Environ() {
    pair := strings.Split(e, "=")

    result[pair[0]] = pair[1]
  }

  return result
}

func printMissingEnv(envName string) {
  fmt.Printf("[Error] Missing '%s' as Env Variable \n", envName)
}

func main() {
  fmt.Printf("[Main] Starting... \n")

  envVariables := parseEnvVariables()

  dbUrl, ok := envVariables["MONGO_URL"]
  if !ok {
    printMissingEnv("MONGO_URL")
    return
  }

  database.SetDBUrl(dbUrl)

  err := database.Connect()
  if err != nil {
    fmt.Printf("[Error][Connect] %v \n", err)
    return
  }

  fmt.Printf("[Main][DB] Connected... \n")

  api.StartAPI(":8080")
}
