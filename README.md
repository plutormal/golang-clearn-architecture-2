# Golang CRUD API

โปรเจค API CRUD แบบพื้นฐานที่เขียนด้วย Golang สำหรับการจัดการข้อมูล User

## Features

- ✅ RESTful API
- ✅ CRUD operations (Create, Read, Update, Delete)
- ✅ JSON response format
- ✅ CORS middleware
- ✅ Logging middleware
- ✅ In-memory storage
- ✅ Error handling

## API Endpoints

### Health Check
- `GET /health` - ตรวจสอบสถานะของ API

### User Management
- `GET /api/v1/users` - ดึงข้อมูล User ทั้งหมด
- `GET /api/v1/users/{id}` - ดึงข้อมูล User ตาม ID
- `POST /api/v1/users` - สร้าง User ใหม่
- `PUT /api/v1/users/{id}` - อัพเดทข้อมูล User
- `DELETE /api/v1/users/{id}` - ลบ User

## Project Structure

```
.
├── main.go              # Entry point
├── go.mod              # Go modules
├── models/
│   └── user.go         # User model
├── services/
│   └── user_service.go # Business logic
├── handlers/
│   └── user_handler.go # HTTP handlers
├── routes/
│   └── routes.go       # Route setup
└── middleware/
    └── middleware.go   # Middleware functions
```

## การรันโปรเจค

1. Install dependencies:
```bash
go mod tidy
```

2. รันโปรเจค:
```bash
go run main.go
```

3. API จะทำงานที่: `http://localhost:8080`

## การทดสอบ API

### 1. Health Check
```bash
curl http://localhost:8080/health
```

### 2. สร้าง User ใหม่
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
  }'
```

### 3. ดึงข้อมูล User ทั้งหมด
```bash
curl http://localhost:8080/api/v1/users
```

### 4. ดึงข้อมูل User ตาม ID
```bash
curl http://localhost:8080/api/v1/users/1
```

### 5. อัพเดทข้อมูล User
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "age": 35
  }'
```

### 6. ลบ User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Response Format

### Success Response
```json
{
  "success": true,
  "data": {...},
  "message": "Optional message"
}
```

### Error Response
```json
{
  "success": false,
  "error": "Error message"
}
```

## User Model
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com", 
  "age": 30,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```
