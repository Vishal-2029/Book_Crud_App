# Book CRUD App (Golang)

    A simple and modular Book CRUD (Create, Read, Update, Delete) API built with **Go**, using clean architecture principles. It includes configuration management, logging, dependency injection, and Swagger API documentation.

---

##  Features

    -  RESTful API for managing books
    -  CRUD operations (Create, Read, Update, Delete)
    -  Dependency Injection
    -  Modular package structure
    -  Swagger documentation (`swagger.yaml`)
    -  Docker support
    -  Logging utility
    -  Config management with `.env` file

---

### 1. Clone the Repository

```bash
git clone https://github.com/Vishal-2029/bookcrud_app.git
cd bookcrud_app
```

### 2. Setup Environment
  - Create a .env file in the root directory and add necessary configurations:
    ```bash
    DB_NAME=book_crud
    DB_USER=root
    DB_PASSWORD=root
    DB_HOST=localhost
    ```

### 3. Build & Run 
  - Using Go
    ```bash
    go mod tidy
    go run cmd/main.go
    ```
    
  - Using Makefile
    ```bash
     make run
    ```
    
  - Using Docker
    ```bash
    docker-compose up --build
    ```
    
### 4. API Documentation 
  - Swagger is used for API documentation.
      - Access Swagger UI at: ` http://localhost:8080/swagger/index.html `
      - OpenAPI spec: `docs/swagger.yaml`

## Running Tests
    `go test ./...`

## Dependencies  
     - Go modules
     - MySQL
     - Docker 
     - Swaggo for Swagger docs
     - logger

 ## Useful Commands
  ```bash
    make run          # Run the app
    make build        # Build the binary
    make tidy         # Tidy up Go modules
    make test         # Run tests
    make docker       # Run Docker images
  ```

### License
    MIT © 2025 – Vishal-2029.

