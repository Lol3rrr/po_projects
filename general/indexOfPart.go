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
