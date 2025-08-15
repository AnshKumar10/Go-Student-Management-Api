# Go Student Management API

A RESTful API for managing student records built with Go. This API provides endpoints to perform CRUD operations on student data with proper error handling and data validation.

## Features

- **Student Management**
  - Create new student records
  - Retrieve student details by ID
  - Update existing student information
  - Delete student records
  - List all students with pagination support

- **Data Validation**
  - Input validation for all API endpoints
  - Proper error handling and meaningful status codes
  - Data sanitization to prevent injection attacks

- **Database**
  - SQLite database for data persistence
  - Database migrations support
  - Connection pooling for better performance

- **API Documentation**
  - RESTful endpoints with proper HTTP methods
  - JSON request/response format
  - Consistent error response structure

## Prerequisites

- Go 1.16 or higher
- SQLite3

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/AnshKumar10/Go-Student-Management-Api.git
   cd Go-Student-Management-Api
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the root directory and configure your environment variables:
   ```
   DB_PATH=./students.db
   PORT=8080
   ```

## Running the Application

```bash
# Build the application
go build -o student-api ./cmd/student-api

# Run the application
./student-api
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Students

- `GET /api/students` - Get all students (with pagination)
- `POST /api/students` - Create a new student
- `GET /api/students/{id}` - Get a specific student by ID
- `PUT /api/students/{id}` - Update a student
- `DELETE /api/students/{id}` - Delete a student

## Request/Response Examples

### Create a Student

**Request:**
```http
POST /api/students
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 20,
  "grade": "A"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 20,
  "grade": "A",
  "created_at": "2025-02-15T12:00:00Z",
  "updated_at": "2025-02-15T12:00:00Z"
}
```

### Get All Students

**Request:**
```http
GET /api/students?page=1&limit=10
```

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "age": 20,
      "grade": "A",
      "created_at": "2025-02-15T12:00:00Z",
      "updated_at": "2025-02-15T12:00:00Z"
    }
  ],
  "pagination": {
    "current_page": 1,
    "per_page": 10,
    "total_items": 1,
    "total_pages": 1
  }
}
```

## Error Handling

All error responses follow this format:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": {
      "field": "Additional error details"
    }
  }
}
```

### Common Error Status Codes

- `400 Bad Request` - Invalid request data
- `404 Not Found` - Resource not found
- `409 Conflict` - Resource already exists
- `422 Unprocessable Entity` - Validation errors
- `500 Internal Server Error` - Server error

## Testing

Run the test suite:

```bash
go test -v ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with ❤️ using Go
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Viper](https://github.com/spf13/viper) for configuration management
