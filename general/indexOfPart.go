package general

func (project *Project) IndexOfTextPart(textID string) (int) {
  if project.TextParts == nil {
    project.TextParts = make([]Project_Text_Part, 0)
    return -1
  }

  for index, tmpValue := range project.TextParts {
    if tmpValue.ID == textID {
      return index
    }
  }

  return -1
}

func (project *Project) IndexOfListPart(listID string) (int) {
  if project.ListParts == nil {
    project.ListParts = make([]Project_List_Part, 0)
    return -1
  }

  for index, tmpValue := range project.ListParts {
    if tmpValue.ID == listID {
      return index
    }
  }

  return -1
}

func (project *Project) IndexOfTodoPart(todoID string) (int) {
  if project.TodoParts == nil {
    project.TodoParts = make([]Project_Todo_Part, 0)
    return -1
  }

  for index, tmpValue := range project.TodoParts {
    if tmpValue.ID == todoID {
      return index
    }
  }

  return -1
}
