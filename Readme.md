## TAG-Onboarding-Go

### Getting Started:
You can run the application 2 ways:
 - locally using go:  `go run cmd/tag-onboarding/main.go` 
 - using docker : `docker-compose up --build`

### API Usage
Once the application is running, you can interact with it using curl commands:

#### Save a User
```bash
curl -X POST http://localhost:8080/save \
  -H "Content-Type: application/json" \
  -d @data/user1.json
```

Or with inline JSON:
```bash
curl -X POST http://localhost:8080/save \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "age": 25
  }'
```

#### Find a User
```bash
curl -X GET http://localhost:8080/find/1
```

#### Example Response
Both endpoints return JSON responses in the following format:
```json
{
  "id": "1",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "age": 25
}
```

### Testing
This project comes with comprehensive testing
- unit tests: `go test -v ./...`
- integration tests and unit tests:
    * `docker compose -f 'docker-compose.yml' up -d --build 'mongo' `
    * `go test -tags=integration ./...`

### Project Layout:
``` 
- cmd
    - tag-onboarding
- data
- internal
    - application
    - domain
    - infrastructure
    - interface
```
* cmd - is the entrypoint to the application 
* data - holds any data files used for testing 
* internal/application - holds the buisness logic for the application 
* internal/domain - holds the domain specific logic 
* internal/infrastructure - holds logic for persistance and third party apis
* internal/interface - holds the logic to interface with external users id: api definition