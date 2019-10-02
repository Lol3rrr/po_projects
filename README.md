# po_projects
The Projects Service of my Project-Overview Website

## Endpoints
### Post
 - /create
   - sessionID : The Session ID of the current User
   - name : The Name of the new Project
 - /delete
   - sessionID : The Session ID of the current User
   - id : The ID of the Project

### Get
 - /find
   - One of:
     - id : The ID of a Project
     - name : The Name of a Project
