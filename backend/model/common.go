package model

type Id struct {
	Id string `json:"id" binding:"required"`
}

type Metadata struct {
	Limit int32 `json:"limit" binding:"required"`
	Page  int32 `json:"page" binding:"required"`
	Count int32 `json:"count" binding:"required"`
}

type Empty struct {
}

type Message struct {
	Description string `json:"description" binding:"required"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type TokenSet struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// {
// 	"success": true,
// 	"message": "User created successfully",
// 	"data": {
// 		"userId": "12345abcde",
// 		"username": "johndoe",
// 		"email": "john.doe@example.com",
// 		"avatarUrl": "https://example.com/avatars/12345abcde.jpg",
// 		"createdAt": "2025-03-02T10:30:45Z"
// 	},
// 	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
// }
