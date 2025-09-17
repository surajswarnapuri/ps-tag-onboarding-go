## TAG-Onboarding-Go

### Getting Started:
You can run the application 2 ways:
 - locally using go:  `go run cmd/tag-onboarding/main.go` 
 - using docker : `docker-compose up --build`

### Testing
This project comes with comprehensive testing
- unit tests: `go test -v ./...`
- integration tests and unit tests: `go test -tags=integration ./...`

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