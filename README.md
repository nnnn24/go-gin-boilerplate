# Go Gin Boilerplate

A boilerplate project for building web applications using the [Gin](https://github.com/gin-gonic/gin) framework in Go.

## Features

- Lightweight and fast HTTP server with Gin.
- Structured project layout for scalability.
- Environment-based configuration management.
- Middleware support (e.g., logging, CORS, authentication).
- RESTful API design.
- Example routes and handlers.
- Docker support for containerized deployment.

## Prerequisites

- [Go](https://golang.org/dl/) 1.18 or later
- [Docker](https://www.docker.com/) (optional, for containerization)

## Getting Started

1. Clone the repository:
  ```bash
  git clone https://github.com/nnnn24/go-gin-boilerplate.git
  cd go-gin-boilerplate
  ```

2. Install dependencies:
  ```bash
  go mod tidy
  ```

3. Run the application:
  ```bash
  go run main.go
  ```

4. Access the application at `http://localhost:8080`.

## Project Structure

```
.
├── main.go             # Application entry point
├── config/             # Configuration files
├── controllers/        # Route handlers
├── middleware/         # Custom middleware
├── models/             # Data models
├── routes/             # API route definitions
├── utils/              # Utility functions
└── Dockerfile          # Docker configuration
```

## Configuration

Environment variables can be used to configure the application. Create a `.env` file in the root directory:

```
PORT=8080
GIN_MODE=debug
```

## Docker

Build and run the application using Docker:

```bash
docker build -t go-gin-boilerplate .
docker run -p 8080:8080 go-gin-boilerplate
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
