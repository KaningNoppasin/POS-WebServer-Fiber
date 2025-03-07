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
📁 Web-Server-Fiber
    ├── 📁 cmd                    # Entry point of the application (main.go)
    │   └── main.go
    ├── 📁 internal               # Core business logic and adapters
    │   ├── 📁 adapter            # Adapters (implement interfaces for DB and HTTP)
    │   │   ├── 📁 db             # Database repositories (data access layer)
    │   │   │   ├── bill_details_repository.go
    │   │   │   ├── bill_repository.go
    │   │   │   ├── customer_repository.go
    │   │   │   ├── product_repository.go
    │   │   │   ├── stock_repository.go
    │   │   ├── 📁 http           # HTTP handlers and routes
    │   │       ├── bill_details_handler.go
    │   │       ├── bill_handler.go
    │   │       ├── bill_routes.go
    │   │       ├── customer_handler.go
    │   │       ├── customer_routes.go
    │   │       ├── product_handler.go
    │   │       ├── product_routes.go
    │   │       ├── routes.go
    │   │       ├── stock_handler.go
    │   │       ├── stock_routes.go
    │   ├── 📁 core               # Business logic (Entities, Ports, Services)
    │   │   ├── 📁 entity         # Domain entities (core business models)
    │   │   │   ├── bill_details.go
    │   │   │   ├── bill.go
    │   │   │   ├── customer.go
    │   │   │   ├── product.go
    │   │   │   ├── stock.go
    │   │   ├── 📁 port           # Ports (interfaces for repositories & services)
    │   │   │   ├── bill_details_repository.go
    │   │   │   ├── bill_details_service.go
    │   │   │   ├── bill_repository.go
    │   │   │   ├── bill_service.go
    │   │   │   ├── customer_repository.go
    │   │   │   ├── customer_service.go
    │   │   │   ├── product_repository.go
    │   │   │   ├── product_service.go
    │   │   │   ├── stock_repository.go
    │   │   │   ├── stock_service.go
    │   │   ├── 📁 service        # Service layer (business logic implementations)
    │   │       ├── bill_details_service.go
    │   │       ├── bill_service.go
    │   │       ├── customer_service.go
    │   │       ├── error.go
    │   │       ├── product_service.go
    │   │       ├── stock_service.go
    ├── 📁 pkg                    # Shared utilities and configurations
    │   ├── 📁 config             # App configuration
    │   │   ├── config.go
    │   ├── 📁 database           # Database connection setup
    │   │   ├── postgres.go
    │   ├── 📁 response           # Standardized API response helpers
    │   │   ├── response.go
    │   ├── 📁 util               # Utility functions (e.g., file uploads)
    │       ├── file_upload.go
    ├── 📁 seeds                  # Database seeding scripts
    │   ├── seed_data.go
    ├── 📁 uploads                # File uploads storage (e.g., images)
    │   ├── default_image.jpeg
    ├── .env                      # Environment variables
    ├── .gitignore                # Git ignore file
    ├── go.mod                     # Go module dependencies
    ├── go.sum                     # Go dependencies checksum
    ├── README.md                  # Project documentation
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

🔧 How to Build & Run
1️⃣ Build the Image

docker build -t web-server-fiber .

2️⃣ Run the Container
docker run --name web-server-fiber -p 8080:8080 web-server-fiber

docker-compose up -d
docker logs web-server-fiber

## License
MIT License
