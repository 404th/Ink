# Backend API Documentation

## Overview

This document outlines the API endpoints, request/response formats, and methods for handling user forms in the system, with particular focus on the signup process and user management.

## Authentication

Most endpoints require authentication via JWT token in the Authorization header:

```
Authorization: Bearer {token}
```

## API Endpoints

### User Management

#### 1. User Signup

Creates a new user account.

**Endpoint:** `POST /api/users/signup`

**Request Format:**

```json
{
  "username": "johndoe",
  "email": "john.doe@example.com",
  "password": "SecurePassword123",
  "avatarFile": "imageUrl.com/sdfnsdfsdf"  // before sending image 
}
```

**Response Format:**

```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "userId": "12345abcde",
    "username": "johndoe",
    "email": "john.doe@example.com",
    "avatarUrl": "https://example.com/avatars/12345abcde.jpg",
    "createdAt": "2025-03-02T10:30:45Z"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Notes:**

- The `avatarFile` field should be sent as multipart/form-data.
- Supported image formats: JPEG, PNG, GIF (max size: 2MB).
- If no avatar is provided, a default avatar will be assigned.

#### 2. User Login

Authenticates a user and returns a JWT token.

**Endpoint:** `POST /api/users/login`

**Request Format:**

```json
{
  "email": "john.doe@example.com",
  "password": "SecurePassword123"
}
```

**Response Format:**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "userId": "12345abcde",
    "username": "johndoe",
    "avatarUrl": "https://example.com/avatars/12345abcde.jpg"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 3. Update User Avatar

Updates a user's Avatar.

**Endpoint:** `PUT /api/users/avatar`

**Headers:**

```
Authorization: Bearer {token}
Content-Type: multipart/form-data
```

**Request Format:**
Form data with key:

- `avatarFile`: The image file to upload

**Response Format:**

```json
{
  "success": true,
  "message": "Avatar updated successfully",
  "data": {
    "avatarUrl": "https://example.com/avatars/12345abcde_updated.jpg",
    "updatedAt": "2025-03-02T14:25:10Z"
  }
}
```

#### 4. Get User Profile

Retrieves a user's profile information.

**Endpoint:** `GET /api/users/profile`

**Headers:**

```
Authorization: Bearer {token}
```

**Response Format:**

```json
{
  "success": true,
  "data": {
    "userId": "12345abcde",
    "username": "johndoe",
    "email": "john.doe@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "avatarUrl": "https://example.com/avatars/12345abcde.jpg",
    "createdAt": "2025-03-02T10:30:45Z",
    "lastLogin": "2025-03-02T14:20:30Z"
  }
}
```

#### 5. Update User Profile

Updates a user's profile information.

**Endpoint:** `PUT /api/users/profile`

**Headers:**

```
Authorization: Bearer {token}
```

**Request Format:**

```json
{
  "firstName": "Johnny",
  "lastName": "Doe",
  "email": "johnny.doe@example.com"
}
```

**Response Format:**

```json
{
  "success": true,
  "message": "Profile updated successfully",
  "data": {
    "userId": "12345abcde",
    "username": "johndoe",
    "email": "johnny.doe@example.com",
    "firstName": "Johnny",
    "lastName": "Doe",
    "avatarUrl": "https://example.com/avatars/12345abcde.jpg",
    "updatedAt": "2025-03-02T15:45:20Z"
  }
}
```

## Error Handling

All endpoints return standardized error responses:

```json
{
  "success": false,
  "error": {
    "code": "INVALID_REQUEST",
    "message": "Detailed error message",
    "details": [
      "Error detail 1",
      "Error detail 2"
    ]
  }
}
```

### Common Error Codes:

- `INVALID_REQUEST`: Request format or parameters are invalid
- `UNAUTHORIZED`: Authentication required or failed
- `FORBIDDEN`: User lacks permission for this action
- `NOT_FOUND`: Requested resource not found
- `CONFLICT`: Resource already exists or conflict occurred
- `SERVER_ERROR`: Internal server error

## File Upload Guidelines

### Avatars:

- **Supported formats:** JPEG, PNG, GIF
- **Maximum file size:** 2MB
- **Recommended dimensions:** 200x200 pixels (images will be resized if larger)
- **Content restrictions:** No explicit content allowed (uploads are moderated)

### Form Processing:

1. For endpoints that accept file uploads, use `multipart/form-data` encoding
2. Form fields can be sent alongside file uploads in the same request
3. File upload fields should include the file object directly, not a URL or base64 string

## Rate Limiting

API endpoints are subject to rate limiting:

- Authentication endpoints: 10 requests per minute
- Profile update endpoints: 30 requests per hour
- Avatar upload: 5 uploads per hour

Exceeding these limits will result in a `429 Too Many Requests` response.
