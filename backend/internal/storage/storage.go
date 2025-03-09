package storage

import (
	"context"

	"github.com/404th/Ink/internal/storage/postgres/postPg"
	"github.com/404th/Ink/internal/storage/postgres/userPg"
	"github.com/404th/Ink/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) StorageI {
	return &storage{
		db: db,
	}
}

func (s *storage) User() UserPgI {
	return userPg.NewUserPg(s.db)
}

func (s *storage) Post() PostPgI {
	return postPg.NewPostPg(s.db)
}

// abstract interfaces
type StorageI interface {
	User() UserPgI
	Post() PostPgI
}

type UserPgI interface {
	// user storage
	SignupUser(ctx context.Context, req *model.SignupUserRequest) (resp *model.SignupUserResponse, err error)
	LoginUser(ctx context.Context, req *model.LoginUserRequest) (resp *model.LoginUserResponse, err error)
	GetUser(ctx context.Context, req *model.Id) (resp *model.User, err error)
}

type PostPgI interface {
	// post service
	CreatePost(ctx context.Context, req *model.CreatePostRequest) (resp *model.Post, err error)
	GetAllPosts(ctx context.Context, req *model.GetAllPostsRequest) (resp *model.GetAllPostsResponse, err error)
}
