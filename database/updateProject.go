package database

import (
  "po_projects/general"
)

func UpdateProject(project general.Project) {
  projectID := project.ID
  _, err := FindProject_ID(projectID)
  if err != nil {
    insertProject(project)
  }else {
    updateProjectDB(projectID, project)
  }
}
