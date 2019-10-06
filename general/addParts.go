package general

func (project *Project) AddTextPart(part Project_Text_Part) {
  if project.TextParts == nil {
    project.TextParts = make([]Project_Text_Part, 0)
  }

  index := project.IndexOfTextPart(part.ID)
  if index == -1 {
    project.TextParts = append(project.TextParts, part)
    return
  }

  project.TextParts[index] = part
}

func (project *Project) AddListPart(part Project_List_Part) {
  if project.ListParts == nil {
    project.ListParts = make([]Project_List_Part, 0)
  }

  index := project.IndexOfListPart(part.ID)
  if index == -1 {
    project.ListParts = append(project.ListParts, part)
    return
  }

  project.ListParts[index] = part
}

func (project *Project) AddTodoPart(part Project_Todo_Part) {
  if project.TodoParts == nil {
    project.TodoParts = make([]Project_Todo_Part, 0)
  }

  index := project.IndexOfTodoPart(part.ID)
  if index == -1 {
    project.TodoParts = append(project.TodoParts, part)
    return
  }

  project.TodoParts[index] = part
}
