
# PointBox

 PointBox is a simple Go HTTP application for managing members and their points.

 ## Features

 - Add members
- View all members
- Delete members
- Add 10 points to a member
- Subtract 10 points from a member
- Track when a member was last updated
- Automatic member IDs

 > **Note:** Member data is currently stored in memory and will be lost when the application stops.

 ## Tech Stack

 - Go
- Go standard library (`net/http`, `time`, `strconv`, `strings`)
- In-memory storage using a Go map

 ## Getting Started

 ### Prerequisites

 Make sure Go is installed:

```
go version
```

 ### Installation

 Clone the repository:

```
git clone https://github.com/YOUR_USERNAME/pointbox.git
cd pointbox
```

 Run the application:

```
go run .
```

 The server will start on:

```
http://localhost:8080
```

 ## API Endpoints

 | Method | Endpoint | Description |
| --- | --- | --- |
| GET | `/` | Get all members |
| GET | `/add/{name}` | Add a new member |
| GET | `/delete/{id}` | Delete a member |
| GET | `/points/add/{id}` | Add 10 points |
| GET | `/points/delete/{id}` | Subtract 10 points |

### Examples

 Get all members:

```
curl http://localhost:8080/
```

 Add a member:

```
curl http://localhost:8080/add/John
```

 Add 10 points:

```
curl http://localhost:8080/points/add/1
```

 Subtract 10 points:

```
curl http://localhost:8080/points/delete/1
```

 Delete a member:

```
curl http://localhost:8080/delete/1
```

 ## Project Structure

```
pointbox/
├── main.go              # Application entry point
├── member.go            # Member data model
├── member_handler.go    # HTTP request handlers
├── member_store.go      # In-memory member storage and operations
├── routes.go            # HTTP route configuration
└── go.mod               # Go module definition
```

 ## Architecture

 The application follows a simple separation of responsibilities:

```
HTTP Request
     ↓
   Router
     ↓
  Handler
     ↓
Member Store
     ↓
 Member Data
```

 - **Routes** decide which handler receives a request.
- **Handlers** process HTTP requests and responses.
- **MemberStore** manages member data.
- **Member** defines the structure of a member.

 ## Development

 Format the code:

```
gofmt -w .
```

 Run tests:

```
go test ./...
```

 Build the project:

```
go build
```

 ## Current Limitations

 - Data is not persistent.
- No database is currently used.
- No authentication is implemented.
- Points are fixed at `+10` or `-10`.
- There are currently no automated tests.

 ## Future Improvements

 - Add a persistent database
- Use RESTful HTTP methods
- Return JSON responses
- Add automated tests
- Add authentication
- Add configurable point amounts

 ## License

 No license has been specified yet.
