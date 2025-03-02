# Point of Sale (POS) System Web Server

## Overview
This project is a Point of Sale (POS) system web server built using Go and the Fiber framework. The system follows the Hexagonal Architecture pattern, ensuring a modular and scalable design.

## Prerequisites
Before starting the project, ensure that you have the following installed:

- **Go**: [Download Go](https://golang.org/dl/)
- **Nodemon** (for auto-reloading in development): 
  Install globally via npm:
  ```bash
  npm install -g nodemon
  ```

## Project Structure
```
.
├── README.md
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── adapter
│   │   ├── db
│   │   │   ├── bill_details_repository.go
│   │   │   ├── bill_repository.go
│   │   │   ├── product_repository.go
│   │   │   └── stock_repository.go
│   │   └── http
│   │       ├── bill_details_handler.go
│   │       ├── bill_handler.go
│   │       ├── product_handler.go
│   │       ├── product_routes.go
│   │       ├── routes.go
│   │       └── stock_handler.go
│   └── core
│       ├── entity
│       │   ├── bill.go
│       │   ├── bill_details.go
│       │   ├── product.go
│       │   └── stock.go
│       ├── port
│       │   ├── bill_details_repository.go
│       │   ├── bill_details_service.go
│       │   ├── bill_repository.go
│       │   ├── bill_service.go
│       │   ├── product_repository.go
│       │   ├── product_service.go
│       │   ├── stock_repository.go
│       │   └── stock_service.go
│       └── service
│           ├── bill_details_service.go
│           ├── bill_service.go
│           ├── error.go
│           ├── product_service.go
│           └── stock_service.go
├── pkg
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── postgres.go
│   ├── response
│   │   └── response.go
│   └── util
│       └── file_upload.go
└── uploads
    └── default_image.jpeg
```

## Architecture
The system is based on Hexagonal Architecture, also known as Ports and Adapters, which provides separation of concerns and flexibility for future modifications.

### Layers
1. **Core Layer:** Contains business logic and domain entities.
   - `entity/`: Defines domain entities (Bill, Product, etc.).
   - `port/`: Defines interfaces for repositories and services.
   - `service/`: Implements business logic.

2. **Adapter Layer:** Provides implementations for ports.
   - `db/`: Implements repositories interacting with PostgreSQL.
   - `http/`: Handles HTTP requests and routes.

3. **Infrastructure Layer:** Contains utilities and configurations.
   - `config/`: Manages application configurations.
   - `database/`: Handles database connection.
   - `response/`: Standardized API response structure.
   - `util/`: Utility functions (e.g., file uploads).

## Installation

### 1. Install Go Dependencies
Ensure that your Go dependencies are up to date by running the following command in your project directory:
```bash
go mod tidy
```

### 2. Set Up Swagger
Generate the Swagger documentation for your API with the following command:
```bash
swag init --parseDependency --parseInternal
swag init
```

## Running the Project

### Option 1: With Nodemon (Auto-reload during development)
Use Nodemon to automatically reload your application upon changes:
```bash
nodemon --exec go run cmd/main.go --signal SIGTERM
```

### Option 2: Directly with Go
You can also run the project directly with Go:
```bash
go run cmd/main.go
```

## API Endpoints
The system provides RESTful API endpoints for managing products, bills, stocks, etc. Routes are defined in `internal/adapter/http/routes.go`.

## Database
The application uses PostgreSQL for data storage. The connection is managed in `pkg/database/postgres.go`.

## License
MIT License
