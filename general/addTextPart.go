package general

func (project *Project) AddTextPart(part Project_Text_Part) {
  index := project.IndexOfTextPart(part.ID)
  if index == -1 {
    project.TextParts = append(project.TextParts, part)
    return
  }

  project.TextParts[index] = part
}
