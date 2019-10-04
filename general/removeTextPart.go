package general

func (project *Project) RemoveTextPart(itemID string) {
  index := project.IndexOfTextPart(itemID)
  if index == -1 {
    return
  }

  project.TextParts = append(project.TextParts[:index], project.TextParts[index+1:]...)
}
