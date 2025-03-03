package model

type User struct {
	Id        string `json:"id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	AvatarUrl string `json:"avatar_url" binding:"required"`
	CreatedAt string `json:"created_at" binding:"required"`
}

type SignupUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	AvatarUrl string `json:"avatar_url" binding:"required"`
}

type SignupUserResponse struct {
	Id        string    `json:"id" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	AvatarUrl string    `json:"avatar_url" binding:"required"`
	TokenSet  *TokenSet `json:"token_set"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Id        string    `json:"id" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	AvatarUrl string    `json:"avatar_url" binding:"required"`
	TokenSet  *TokenSet `json:"token_set" binding:"required"`
}
