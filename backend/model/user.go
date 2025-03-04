package model

type User struct {
	Id        string `json:"id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	AvatarUrl string `json:"avatarUrl" binding:"required"`
	CreatedAt string `json:"createdAt" binding:"required"`
}

type SignupUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	AvatarUrl string `json:"avatarUrl" binding:"required"`
}

type SignupUserResponse struct {
	Id        string    `json:"id" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	AvatarUrl string    `json:"avatarUrl" binding:"required"`
	TokenSet  *TokenSet `json:"tokenSet"`
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
	AvatarUrl string    `json:"avatarUrl" binding:"required"`
	TokenSet  *TokenSet `json:"tokenSet" binding:"required"`
}
