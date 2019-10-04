package general

func (project *Project) IsOwner(userID string) bool {
  return (project.Owner.ID == userID)
}
