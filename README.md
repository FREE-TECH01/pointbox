

 # PointBox

 PointBox is a simple Go HTTP application for managing members and their points.

 The application currently provides an in-memory member store, allowing users to:

 - View all members
- Add new members
- Delete members
- Add 10 points to a member
- Subtract 10 points from a member
- Track when a member's information was last updated

 The project is intentionally simple and is currently implemented without an external database. All member data is stored in memory while the application is running.

---

 ## Table of Contents

 - Features
- Technology
- Project Structure
- Architecture
- Requirements
- Installation
- Running the Application
- Using the API
- API Routes
- Example Workflow
- How the Application Works
- Data Storage
- Important Behavior
- Development
- Stopping the Application
- Future Improvements
- License

---

 ## Features

 ### Member Management

 PointBox currently supports the following member operations:

 - Create a member
- Retrieve all members
- Delete a member
- Add points to a member
- Subtract points from a member

 ### Default Points

 Every newly created member starts with:

```
1500 points
```

 ### Automatic IDs

 Members are automatically assigned an integer ID starting from `1`.

 For example:

```
First member  → ID 1
Second member → ID 2
Third member  → ID 3
```

 ### Last Updated

 When a member is created or their points are changed, PointBox records the time of the operation using the `LastUpdated` field.

---

 # Technology

 PointBox is currently built with:

 - **Go**
- Go's standard `net/http` package
- Go's standard `time` package
- Go's standard `strconv` package
- Go's standard `strings` package

 No external database or third-party Go packages are currently required.

---

 # Project Structure

 The project currently contains the following files:

```
pointbox/
│
├── go.mod
├── main.go
├── member.go
├── member_handler.go
├── member_store.go
├── routes.go
└── README.md
```

 ## File Responsibilities

 ### `main.go`

 The entry point of the application.

 It is responsible for creating and connecting the major components:

```
MemberStore
     ↓
MemberHandler
     ↓
Router
     ↓
HTTP Server
```

 The server currently runs on port `8080`.

---

 ### `member.go`

 Contains the `Member` data structure.

 A member currently contains:

```
type Member struct {
	Name        string
	Points      int
	LastUpdated string
}
```

 The file defines the shape of member data but does not contain the logic for modifying members.

---

 ### `member_store.go`

 Contains `MemberStore`, which is the application's current in-memory data store.

 It is responsible for:

 - Creating members
- Retrieving members
- Deleting members
- Adding points
- Subtracting points
- Generating member IDs

 The data is stored using:

```
map[int]Member
```

 The integer is the member's ID, while the `Member` is the member's data.

---

 ### `member_handler.go`

 Contains the HTTP handlers.

 Handlers receive HTTP requests, extract information from the URL, call the appropriate store method, and return an HTTP response.

 Current handlers include:

```
GetMembers
AddMember
DeleteMember
AddPoints
DeletePoints
```

---

 ### `routes.go`

 Contains the application's HTTP routing configuration.

 It connects URL paths to the appropriate handler functions.

 For example:

```
/add/John
```

 is connected to:

```
AddMember
```

---

 # Architecture

 PointBox uses a simple separation between the data model, storage, HTTP handlers, and routing.

 The overall architecture is:

```
                    HTTP Client
                         │
                         ▼
                      Router
                    routes.go
                         │
                         ▼
                 MemberHandler
              member_handler.go
                         │
                         ▼
                  MemberStore
               member_store.go
                         │
                         ▼
                  Member Model
                   member.go
```

 ## Request Flow

 For example, if a client requests:

```
/points/add/1
```

 the request flows through the application like this:

```
HTTP Request
     │
     ▼
routes.go
     │
     ▼
MemberHandler.AddPoints()
     │
     ▼
MemberStore.AddPoints()
     │
     ▼
Member #1 is updated
     │
     ▼
HTTP Response
```

 This separation means the HTTP layer does not directly manage the member data.

---

 # Requirements

 Before installing PointBox, make sure you have Go installed.

 You can check whether Go is installed by running:

```
go version
```

 You should receive output similar to:

```
go version go1.x.x ...
```

 The exact version may differ depending on the Go version installed on your machine.

 You can install Go from the official Go website:

 https://go.dev/

---

 # Installation

 ## 1\. Clone the repository

 Clone the PointBox repository from GitHub:

```
git clone https://github.com/YOUR_USERNAME/pointbox.git
```

 Replace `YOUR_USERNAME` with the GitHub username that owns the repository.

 Then move into the project directory:

```
cd pointbox
```

---

 ## 2\. Verify the project files

 Run:

```
ls
```

 You should see files similar to:

```
go.mod
main.go
member.go
member_handler.go
member_store.go
routes.go
README.md
```

---

 ## 3\. Download/verify Go dependencies

 PointBox currently uses Go's standard library, so there should be no external dependencies to install.

 However, you can run:

```
go mod tidy
```

 This ensures the `go.mod` file is consistent with the code.

---

 ## 4\. Verify that the project builds

 Run:

```
go build
```

 If there are no errors, the project should compile successfully.

 You can also test the project with:

```
go test ./...
```

 At the current stage of development, there may not yet be automated tests.

---

 # Running the Application

 There are two simple ways to run PointBox.

 ## Option 1: Run directly with Go

 From the project directory:

```
go run .
```

 You should see:

```
Starting server on :8080
```

 The application will then be available at:

```
http://localhost:8080
```

---

 ## Option 2: Build and run the application

 Build the application:

```
go build
```

 This will create an executable.

 You can then run the generated executable.

 On Linux/macOS, for example:

```
./pointbox
```

 The server will start on port `8080`.

---

 # Using the API

 PointBox currently exposes several HTTP routes.

 ## Get All Members

 ### Endpoint

```
GET /
```

 ### Example

 Open:

```
http://localhost:8080/
```

 or use:

```
curl http://localhost:8080/
```

 ### Example response

```
ID: 1
Name: John
Points: 1500
Last Updated: 22-09-2026 12:30PM

ID: 2
Name: Sarah
Points: 1500
Last Updated: 22-09-2026 12:35PM
```

 The exact output depends on the members currently stored.

---

 # Add a Member

 ### Endpoint

```
/add/{name}
```

 ### Example

```
curl http://localhost:8080/add/John
```

 The application creates a member named `John`.

 The new member starts with:

```
Points: 1500
```

 ### Example response

```
John has been added successfully!
```

 The member will also receive the next available ID.

---

 # Delete a Member

 ### Endpoint

```
/delete/{id}
```

 ### Example

```
curl http://localhost:8080/delete/1
```

 This attempts to delete the member with ID `1`.

 ### Successful response

```
member with 1 was successfully deleted
```

 If the member does not exist, the application returns:

```
404 Not Found
```

---

 # Add Points

 ### Endpoint

```
/points/add/{id}
```

 ### Example

```
curl http://localhost:8080/points/add/1
```

 This adds:

```
10 points
```

 to member `1`.

 For example:

```
Before: 1500 points
After:  1510 points
```

 The member's `LastUpdated` value is also updated.

 ### Successful response

```
10 points successfully added to member 1
```

---

 # Subtract Points

 ### Endpoint

```
/points/delete/{id}
```

 ### Example

```
curl http://localhost:8080/points/delete/1
```

 This subtracts:

```
10 points
```

 from member `1`.

 For example:

```
Before: 1510 points
After:  1500 points
```

 The member's `LastUpdated` value is also updated.

 ### Successful response

```
10 points successfully subtracted from member 1
```

---

 # API Summary

 | Method | Endpoint | Purpose |
| --- | --- | --- |
| GET | `/` | Get all members |
| GET | `/add/{name}` | Add a member |
| GET | `/delete/{id}` | Delete a member |
| GET | `/points/add/{id}` | Add 10 points |
| GET | `/points/delete/{id}` | Subtract 10 points |

> **Note:** The current implementation does not explicitly check HTTP methods. The routes are primarily distinguished by their URL paths.

---

 # Example Workflow

 The following demonstrates a complete basic workflow.

 ## 1\. Start PointBox

```
go run .
```

---

 ## 2\. Create a member

```
curl http://localhost:8080/add/John
```

 Response:

```
John has been added successfully!
```

 John now has:

```
ID: 1
Points: 1500
```

---

 ## 3\. View all members

```
curl http://localhost:8080/
```

 You should see something similar to:

```
ID: 1
Name: John
Points: 1500
Last Updated: 22-09-2026 12:30PM
```

---

 ## 4\. Add points

```
curl http://localhost:8080/points/add/1
```

 John now has:

```
1510 points
```

---

 ## 5\. Subtract points

```
curl http://localhost:8080/points/delete/1
```

 John now has:

```
1500 points
```

---

 ## 6\. Delete the member

```
curl http://localhost:8080/delete/1
```

 The member is removed from the store.

---

 # How the Application Works

 ## Creating the Store

 When the application starts, `main.go` creates a `MemberStore`:

```
store := NewMemberStore()
```

 The store starts with an empty map:

```
members = {}
```

 and the next member ID is:

```
nextID = 1
```

---

 ## Creating the Handler

 The store is passed to the handler:

```
handler := NewMemberHandler(store)
```

 This allows the handler to communicate with the member store.

 The relationship is:

```
MemberHandler
      │
      │ has access to
      ▼
MemberStore
```

---

 ## Creating the Router

 The handler is then passed to:

```
router := SetupRoutes(handler)
```

 The router connects URL paths to handler methods.

 For example:

```
/add/
    ↓
AddMember()

/delete/
    ↓
DeleteMember()

/points/add/
    ↓
AddPoints()
```

---

 ## Starting the Server

 Finally, the router is given to Go's HTTP server:

```
http.ListenAndServe(":8080", router)
```

 The application then waits for HTTP requests on port `8080`.

---

 # Data Storage

 PointBox currently uses an in-memory map:

```
map[int]Member
```

 For example, the data might look conceptually like:

```
members
│
├── 1 → Member{Name: "John", Points: 1510}
├── 2 → Member{Name: "Sarah", Points: 1500}
└── 3 → Member{Name: "Mike", Points: 1600}
```

 The integer is the member ID.

 The value is the corresponding `Member` struct.

---

 # Important: Data Is Not Persistent

 The current version of PointBox does **not** use a database.

 All members exist only in memory.

 For example:

```
Start application
       ↓
Add John
       ↓
John exists
       ↓
Stop application
       ↓
John is gone
```

 When the application is restarted, the store starts empty again:

```
members: make(map[int]Member)
```

 and member IDs start again from:

```
1
```

 This is an important limitation of the current version.

 A future version can replace the in-memory `MemberStore` with a persistent database.

---

 # Important Implementation Details

 ## New Members Start With 1500 Points

 The current implementation gives every new member:

```
1500 points
```

 This happens in `MemberStore.AddMember()`.

---

 ## Adding Points Adds 10

 The current implementation always adds:

```
+10 points
```

 It does not currently allow the caller to specify a different amount.

---

 ## Removing Points Subtracts 10

 The current implementation always subtracts:

```
-10 points
```

 There is currently no protection against a member's points becoming negative.

 For example:

```
Member has 0 points
       ↓
Delete 10 points
       ↓
Member has -10 points
```

 This could be addressed in a future version depending on the intended business rules.

---

 ## Member IDs

 Member IDs start at `1` and increase whenever a member is created.

 For example:

```
John  → 1
Sarah → 2
Mike  → 3
```

 If member `2` is deleted:

```
John  → 1
Sarah → deleted
Mike  → 3
```

 The application does not currently reuse deleted IDs.

---

 # Development

 When modifying the project, it is recommended to format the Go code using:

```
gofmt -w .
```

 You can verify that the project still builds:

```
go build
```

 You can run tests with:

```
go test ./...
```

 And run the application with:

```
go run .
```

---

 # Suggested Development Workflow

 A typical development workflow is:

```
# Get the latest code
git pull

# Make your changes

# Format the Go code
gofmt -w .

# Run tests
go test ./...

# Build the application
go build

# Check what changed
git status

# Stage changes
git add .

# Commit
git commit -m "Describe your changes"

# Push to GitHub
git push
```

---

 # Troubleshooting

 ## Port 8080 Is Already in Use

 If the application cannot start because port `8080` is already being used, another application may already be running on that port.

 On Linux, you can check with:

```
sudo lsof -i :8080
```

 You can then stop the process using that port or change the port in `main.go`.

---

 ## Go Command Not Found

 If you see:

```
go: command not found
```

 Go is either not installed or is not available in your system's `PATH`.

 Install Go and verify the installation with:

```
go version
```

---

 ## Changes Are Not Appearing

 Remember that PointBox stores members in memory.

 If you restart the server, previously created members will disappear.

 This is expected behavior in the current version.

---

 # Current Limitations

 PointBox is currently a simple learning/project implementation and has several limitations:

 - Data is stored only in memory.
- There is no persistent database.
- There is no authentication or authorization.
- There is no user account system.
- There are no automated tests yet.
- Points are always changed by exactly 10.
- Points can currently become negative.
- HTTP methods are not explicitly separated between operations.
- Responses are plain text rather than JSON.
- There is no structured API error format.
- Member names are currently supplied through the URL.
- The `LastUpdated` value is stored as a formatted string rather than a native time value.
- There is currently no concurrency protection around the in-memory store.

 These limitations are not necessarily bugs; they reflect the current scope of the project.

---

 # Future Improvements

 Possible future versions of PointBox could introduce:

 ## Database

 Replace the in-memory `MemberStore` with a persistent database such as:

 - PostgreSQL
- MySQL
- SQLite

 This would allow members to survive application restarts.

---

 ## RESTful API

 The current API could be redesigned to use conventional HTTP methods.

 For example:

```
GET    /members
POST   /members
DELETE /members/{id}
POST   /members/{id}/points
DELETE /members/{id}/points
```

---

 ## JSON Responses

 Instead of returning plain text:

```
ID: 1
Name: John
Points: 1500
```

 the API could return JSON:

```
{
  "id": 1,
  "name": "John",
  "points": 1500,
  "lastUpdated": "2026-09-22T12:30:00Z"
}
```

 This would make the API easier for frontend applications and other services to consume.

---

 ## Validation

 Additional validation could be introduced for:

 - Member names
- Point balances
- IDs
- Request methods
- Request bodies

---

 ## Testing

 Automated unit and integration tests could be added for:

 - Creating members
- Retrieving members
- Deleting members
- Adding points
- Removing points
- Invalid IDs
- Missing names
- Negative point balances

---

 ## Persistent Timestamps

 The `LastUpdated` field could eventually use Go's `time.Time` instead of a formatted string.

 For example:

```
LastUpdated time.Time
```

 Formatting can then be handled when the data is displayed rather than when it is stored.

---

 # License

 No license has currently been specified for this project.

 If this project is intended to be publicly shared or used by other people, consider adding an appropriate open-source license.

---

 # Author

 PointBox was created as a Go project for learning and experimenting with HTTP servers, routing, handlers, data storage, and application structure.

---

 # Quick Start

 For someone who just wants to get PointBox running:

```
git clone https://github.com/YOUR_USERNAME/pointbox.git
cd pointbox
go mod tidy
go run .
```

 Then open:

```
http://localhost:8080
```

 Or test it from the terminal:

```
curl http://localhost:8080/
```

 To add a member:

```
curl http://localhost:8080/add/John
```

 To add 10 points:

```
curl http://localhost:8080/points/add/1
```

 To subtract 10 points:

```
curl http://localhost:8080/points/delete/1
```

 To delete a member:

```
curl http://localhost:8080/delete/1
```

---

 ## Project Architecture at a Glance

```
                         PointBox
                            │
                            ▼
                         main.go
                            │
             ┌──────────────┼──────────────┐
             │              │              │
             ▼              ▼              ▼
       MemberStore     MemberHandler    SetupRoutes
      (data/storage)   (HTTP logic)     (routing)
             │              │              │
             └──────────────┴──────────────┘
                            │
                            ▼
                         Member
                       (data model)
```

 The key idea behind the project is:

 > **Routes decide where a request goes, handlers deal with HTTP, the store manages data, and the Member struct defines the data itself.**

 That separation makes the project easier to understand now and gives it a foundation for adding a real database, a frontend, authentication, tests, and a more complete API in the future.