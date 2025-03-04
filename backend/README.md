# TaskHub API

A modern task management system built with Go and Gin framework.

## Prerequisites

- Go 1.19 or higher
- PostgreSQL
- Air (for hot reload) - optional

## Project Setup

1. Clone the repository

```bash
git clone https://github.com/TheCodeGhinux/TaskHub.git
cd TaskHub
```

2. Install dependencies

```bash
go mod tidy
```
OR
```
go get ./...   
```

3. Install required Go modules

```bash
# Gin framework for routing
go get -u github.com/gin-gonic/gin

# GORM and PostgreSQL driver
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# JWT authentication
go get -u github.com/dgrijalva/jwt-go

# Environment variables
go get -u github.com/joho/godotenv

# UUID handling
go get -u github.com/google/uuid

# Password hashing
go get -u golang.org/x/crypto/bcrypt
```

4. Configure environment variables by creating a `.env` file:

```env
# App Configuration
APP_NAME=TaskHub
APP_ENV=development

# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8081

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=taskhub_db

# JWT Configuration
TOKEN_SECRET=your_jwt_secret
TOKEN_DURATION=24h
```

5. Initialize the database

```bash
# Create PostgreSQL database
psql -U postgres
CREATE DATABASE taskhub_db;
```

6. Run migrations

```bash
# Migrations will run automatically when starting the application
go run main.go
```

7. Run the application

Development mode with hot reload:

```bash
# Install Air first
go install github.com/cosmtrek/air@latest

# Run with Air
air
```

Production mode:

```bash
go run main.go
```

## Project Structure

```
TaskHub/
├── internal/
│   ├── database/          # Database connection and utilities
│   ├── models/            # Data models and migrations
│   └── middleware/        # Custom middleware
├── pkg/
│   ├── config/           # Configuration management
│   ├── controllers/      # Request handlers
│   ├── repository/       # Database operations
│   └── routers/         # Route definitions
├── services/            # Business logic
│   ├── auth/
│   └── user/
└── utils/              # Helper functions and utilities
```

## API Endpoints

### Authentication

- `POST /api/v1/auth/register` - Register a new user

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "password": "secure_password"
}
```

- `POST /api/v1/auth/login` - Login user

```json
{
  "email": "john@example.com",
  "password": "secure_password"
}
```

### Response Format

Success Response:

```json
{
  "status": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

Error Response:

```json
{
  "status": "error",
  "message": "Invalid credentials"
}
```

## Development

### Adding New Features

1. Create new models in `internal/models`
2. Add migrations in `internal/models/migrations`
3. Create services in `services/`
4. Add controllers in `pkg/controllers`
5. Define routes in `pkg/routers`

### Testing

Run tests:

```bash
go test ./...
```

## Troubleshooting

Common issues and solutions:

1. Database connection issues:

   - Verify PostgreSQL is running
   - Check database credentials in .env
   - Ensure database exists

2. Module errors:
   - Run `go mod tidy`
   - Clear module cache: `go clean -modcache`

## Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details
