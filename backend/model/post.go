package model

type Post struct {
	Id        string `json:"id" binding:"require"`
	UserId    string `json:"user_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Rating    int16  `json:"rating" binding:"required"`
	CreatedAt string `json:"created_at" binding:"required"`
}

type CreatePostRequest struct {
	UserId  string `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type GetAllPostsRequest struct {
	Id string `json:"id"`
}

type GetAllPostsResponse struct {
	Posts []*Post `json:"posts" binding:"required"`
}
