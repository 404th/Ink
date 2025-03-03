package service

import (
	"context"

	"github.com/404th/Ink/config"
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
}

type NewsServiceI interface {
	// news service

	// news title service

	// news content type service

	// news content service

	// news
}

type UserServiceI interface {
	// user service
	SignupUser(ctx context.Context, req *model.SignupUserRequest) (resp *model.SignupUserResponse, err error)
	LoginUser(ctx context.Context, req *model.LoginUserRequest) (resp *model.LoginUserResponse, err error)

	// user role service

	// user data service
}

func (s *service) UserService() UserServiceI {
	userServiceN := userService.NewUserService(s.cfg, s.sugar, s.strg.User())
	if userServiceN != nil {
		return userServiceN
	}

	return s.strg.User()
}
