package general

type Project_Owner struct {
  ID    string
  Name  string
}

type Project_Part struct {
  Name    string
  Content string
}

type Project struct {
  ID    string
  Name  string
  Owner Project_Owner
  Parts []Project_Part
  Tags  []string
}
