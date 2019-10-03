package general

type Project_Owner struct {
  ID    string `json:"id"`
}

type Project_Info_Part struct {
  Name    string `json:"name"`
  Content string `json:"content"`
}

type ProgressPoint struct {
  Content string  `json:"content"`
  Done    bool    `json:"done"`
}

type Project_Progress_Part struct {
  Name  string          `json:"name"`
  Parts []ProgressPoint `json:"progress_points"`
}

type Project struct {
  ID            string                  `json:"id"`
  Name          string                  `json:"name"`
  Owner         Project_Owner           `json:"owner"`
  InfoParts     []Project_Info_Part     `json:"info_parts"`
  ProgressParts []Project_Progress_Part `json:"progress_parts"`
  Tags          []string                `json:"tags"`
}
