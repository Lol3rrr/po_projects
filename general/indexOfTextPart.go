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
