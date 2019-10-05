# po_projects
The Projects Service of my Project-Overview Website

## Endpoints
### Post
 - /create
   - sessionID : The Session-ID of the current User
   - name : The Name of the new Project
 - /delete
   - sessionID : The Session-ID of the current User
   - id : The ID of the Project

 - /save/text
   - sessionID : The Session-ID of the current User
   - projectID : The ID of the project
   - itemID : The ID of the Text-Part, can be empty if it is a new Part
   - Body: as json
     - name : The Name of the Text-Part
     - content : The Content of the Text-Part
 - /delete/text
   - sessionID : The Session-ID of the current User
   - projectID : The ID of the project
   - itemID : The ID of the Text-Part

 - /save/list
   - sessionID : The Session-ID of the current User
   - projectID : The ID of the project
   - itemID : The ID of the List-Part, can be empty if it is a new Part
   - Body: as json
     - name : The Name of the Text-Part
     - items : An array of strings that hold the Value of the Item in the List
 - /delete/list
   - sessionID : The Session-ID of the current User
   - projectID : The ID of the project
   - itemID : The ID of the List-Part

### Get
 - /load
   - One of:
     - id : The ID of a Project
     - name : The Name of a Project
