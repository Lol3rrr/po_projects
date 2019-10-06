package general

type Project_Owner struct {
  ID    string `json:"id"`
}

type Project_Text_Part struct {
  ID      string `json:"id"`
  Name    string `json:"name"`
  Content string `json:"content"`
}

type TodoPoint struct {
  Done    bool    `json:"done"`
  Content string  `json:"content"`
}

type Project_Todo_Part struct {
  ID    string      `json:"id"`
  Name  string      `json:"name"`
  Parts []TodoPoint `json:"todo_points"`
}

type ListPoint struct {
  Content string `json:"content"`
}

type Project_List_Part struct {
  ID    string        `json:"id"`
  Name  string        `json:"name"`
  Parts []ListPoint   `json:"list_points"`
}

type Project struct {
  ID        string              `json:"id"`
  Name      string              `json:"name"`
  Owner     Project_Owner       `json:"owner"`
  TextParts []Project_Text_Part `json:"text_parts"`
  ListParts []Project_List_Part `json:"list_parts"`
  TodoParts []Project_Todo_Part `json:"todo_parts"`
  Tags      []string            `json:"tags"`
}
