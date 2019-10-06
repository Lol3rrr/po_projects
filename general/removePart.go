package general

func (project *Project) RemoveTextPart(itemID string) {
  index := project.IndexOfTextPart(itemID)
  if index == -1 {
    return
  }

  project.TextParts = append(project.TextParts[:index], project.TextParts[index+1:]...)
}

func (project *Project) RemoveListPart(itemID string) {
  index := project.IndexOfListPart(itemID)
  if index == -1 {
    return
  }

  project.ListParts = append(project.ListParts[:index], project.ListParts[index+1:]...)
}

func (project *Project) RemoveTodoPart(itemID string) {
  index := project.IndexOfTodoPart(itemID)
  if index == -1 {
    return
  }

  project.TodoParts = append(project.TodoParts[:index], project.TodoParts[index+1:]...)
}
