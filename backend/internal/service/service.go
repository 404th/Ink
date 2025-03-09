package service

import (
	"context"

	"github.com/404th/Ink/config"
	"github.com/404th/Ink/internal/service/postgresService/postService"
	"github.com/404th/Ink/internal/service/postgresService/userService"
	"github.com/404th/Ink/internal/storage"
	"github.com/404th/Ink/internal/tigres"
	"github.com/404th/Ink/model"
	"go.uber.org/zap"
)

type service struct {
	cfg   *config.Config
	sugar *zap.SugaredLogger
	strg  storage.StorageI
	tg    tigres.TigresI
}

func NewService(cfg *config.Config, sugar *zap.SugaredLogger, strg storage.StorageI, tg tigres.TigresI) ServiceI {
	s := &service{
		cfg:   cfg,
		sugar: sugar,
		strg:  strg,
		tg:    tg,
	}

	return s
}

type ServiceI interface {
	UserService() UserServiceI
	PostService() PostServiceI
}

type UserServiceI interface {
	// user service
	SignupUser(ctx context.Context, req *model.SignupUserRequest) (resp *model.SignupUserResponse, err error)
	LoginUser(ctx context.Context, req *model.LoginUserRequest) (resp *model.LoginUserResponse, err error)
	GetUser(ctx context.Context, req *model.Id) (resp *model.User, err error)
}

type PostServiceI interface {
	// post service
	CreatePost(ctx context.Context, req *model.CreatePostRequest) (resp *model.Post, err error)
	GetAllPosts(ctx context.Context, req *model.GetAllPostsRequest) (resp *model.GetAllPostsResponse, err error)
}

func (s *service) UserService() UserServiceI {
	userServiceN := userService.NewUserService(s.cfg, s.sugar, s.strg.User())
	if userServiceN != nil {
		return userServiceN
	}

	return s.strg.User()
}

func (s *service) PostService() PostServiceI {
	postServiceN := postService.NewPostService(s.cfg, s.sugar, s.strg.Post())
	if postServiceN != nil {
		return postServiceN
	}

	return s.strg.Post()
}
