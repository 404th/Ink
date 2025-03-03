package userService

import (
	"context"
	"errors"

	"github.com/404th/Ink/config"
	"github.com/404th/Ink/internal/storage"
	"github.com/404th/Ink/model"
	"github.com/404th/Ink/pkg/helper"
	"github.com/404th/Ink/pkg/jwtToken"
	"go.uber.org/zap"
)

type userService struct {
	cfg   *config.Config
	sugar *zap.SugaredLogger
	strg  storage.UserPgI
}

func NewUserService(cfg *config.Config, sugar *zap.SugaredLogger, strg storage.UserPgI) *userService {
	return &userService{
		cfg:   cfg,
		sugar: sugar,
		strg:  strg,
	}
}

func (u *userService) SignupUser(ctx context.Context, req *model.SignupUserRequest) (resp *model.SignupUserResponse, err error) {
	resp = &model.SignupUserResponse{}
	u.sugar.Infoln("SignupUser ", config.InfoSplitter, "req", config.InfoSplitter, req)

	req.Password, err = helper.HashPassword(req.Password)
	if err != nil {
		u.sugar.Errorln("LoginUser ", config.ErrorSplitter, "helper.HashPassword ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return
	}

	resp, err = u.strg.SignupUser(ctx, req)
	if err != nil {
		u.sugar.Errorln("SignupUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return
	}

	accessToken, err := jwtToken.GenerateAccessJWT(req.Username, resp.Id, u.cfg.AccessTokenSecret, u.cfg.AccessTokenExpiryMinute)
	if err != nil {
		err = errors.New("Tizimga kirishga qayta urinib ko'ring. (Access token)")
		u.sugar.Infoln("LoginUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return nil, err
	}

	refreshToken, err := jwtToken.GenerateRefreshJWT(req.Username, resp.Id, u.cfg.RefreshTokenSecret, u.cfg.RefreshTokenExpiryHour)
	if err != nil {
		err = errors.New("Tizimga kirishga qayta urinib ko'ring! (Refresh token)")
		u.sugar.Infoln("LoginUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return nil, err
	}

	var tokenSet model.TokenSet
	tokenSet.AccessToken = accessToken
	tokenSet.RefreshToken = refreshToken
	resp.TokenSet = &tokenSet

	u.sugar.Infoln("SignupUser ", config.InfoSplitter, "info", config.InfoSplitter, config.ServiceSuccess)
	return
}

func (u *userService) LoginUser(ctx context.Context, req *model.LoginUserRequest) (resp *model.LoginUserResponse, err error) {
	resp = &model.LoginUserResponse{}
	u.sugar.Infoln("LoginUser ", config.InfoSplitter, "req", config.InfoSplitter, req)

	resp, err = u.strg.LoginUser(ctx, req)
	if err != nil {
		u.sugar.Errorln("LoginUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return
	}

	if !helper.CheckPasswordHash(req.Password, resp.Password) {
		err = errors.New("c'g'ri")
		u.sugar.Infoln("LoginUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return nil, err
	}

	var tokenSet model.TokenSet

	accessToken, err := jwtToken.GenerateAccessJWT(req.Username, resp.Id, u.cfg.AccessTokenSecret, u.cfg.AccessTokenExpiryMinute)
	if err != nil {
		err = errors.New("Tizimga kirishga qayta urinib ko'ring. (Access Token)")
		u.sugar.Infoln("LoginUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return nil, err
	}

	refreshToken, err := jwtToken.GenerateRefreshJWT(req.Username, resp.Id, u.cfg.RefreshTokenSecret, u.cfg.RefreshTokenExpiryHour)
	if err != nil {
		err = errors.New("Tizimga kirishga qayta urinib ko'ring! (Refresh Token)")
		u.sugar.Infoln("LoginUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return nil, err
	}

	tokenSet.AccessToken = accessToken
	tokenSet.RefreshToken = refreshToken

	resp.TokenSet = &tokenSet

	resp.Password = ""

	u.sugar.Infoln("LoginUser ", config.InfoSplitter, "info", config.InfoSplitter, config.ServiceSuccess)
	return
}

// func (s *userService) UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (resp *model.Message, err error) {
// 	resp = &model.Message{}
// 	return
// }

// func (u *userService) GetAllUsers(ctx context.Context, req *model.GetAllUsersRequest) (resp *model.GetAllUsersResponse, err error) {
// 	resp = &model.GetAllUsersResponse{}
// 	u.sugar.Infoln("GetAllUsers ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.GetAllUsers(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("GetAllUsers ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
// 		return
// 	}

// 	u.sugar.Infoln("GetAllUsers ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }

// func (s *userService) DeactivateUser() {
// 	return
// }

// // userService
// func (u *userService) CreateUserData(ctx context.Context, req *model.CreateUserDataRequest) (resp *model.Id, err error) {
// 	resp = &model.Id{}
// 	u.sugar.Infoln("CreateUserData ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.CreateUserData(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("CreateUserData ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
// 		return
// 	}

// 	u.sugar.Infoln("CreateUserData ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }

// func (u *userService) UpdateUserData(ctx context.Context, req *model.UpdateUserDataRequest) (resp *model.Message, err error) {
// 	resp = &model.Message{}
// 	u.sugar.Infoln("UpdateUserData ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.UpdateUserData(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("UpdateUserData ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
// 		return
// 	}

// 	u.sugar.Infoln("UpdateUserData ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }

// func (u *userService) GetAllUserDatas(ctx context.Context, req *model.GetAllUserDatasRequest) (resp *model.GetAllUserDatasResponse, err error) {
// 	resp = &model.GetAllUserDatasResponse{}
// 	u.sugar.Infoln("GetAllUserDatas ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.GetAllUserDatas(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("GetAllUserDatas ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
// 		return
// 	}

// 	u.sugar.Infoln("GetAllUserDatas ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }

// // userService
// func (u *userService) CreateUserRole(ctx context.Context, req *model.CreateUserRoleRequest) (resp *model.Id, err error) {
// 	resp = &model.Id{}
// 	u.sugar.Infoln("CreateUserRole ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.CreateUserRole(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("CreateUserRole ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())

// 		return
// 	}

// 	u.sugar.Infoln("CreateUserRole ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }

// func (u *userService) UpdateUserRole(ctx context.Context, req *model.UpdateUserRoleRequest) (resp *model.Message, err error) {
// 	resp = &model.Message{}
// 	u.sugar.Infoln("UpdateUserRole ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.UpdateUserRole(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("UpdateUserRole ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
// 		return
// 	}

// 	u.sugar.Infoln("UpdateUserRole ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }

// func (u *userService) GetAllUserRoles(ctx context.Context, req *model.GetAllUserRolesRequest) (resp *model.GetAllUserRolesResponse, err error) {
// 	resp = &model.GetAllUserRolesResponse{}
// 	u.sugar.Infoln("GetAllUserRoles ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.GetAllUserRoles(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("GetAllUserRoles ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
// 		return
// 	}

// 	u.sugar.Infoln("GetAllUserRoles ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }

// func (u *userService) DeleteUserRole(ctx context.Context, req *model.Id) (resp *model.Message, err error) {
// 	resp = &model.Message{}
// 	u.sugar.Infoln("DeleteUserRole ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.DeleteUserRole(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("DeleteUserRole ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
// 		return
// 	}

// 	u.sugar.Infoln("DeleteUserRole ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }

// func (u *userService) DeleteUserData(ctx context.Context, req *model.Id) (resp *model.Message, err error) {
// 	resp = &model.Message{}
// 	u.sugar.Infoln("DeleteUserData ", config.InfoSplitter, "req", config.InfoSplitter, req)

// 	resp, err = u.strg.DeleteUserData(ctx, req)
// 	if err != nil {
// 		u.sugar.Errorln("DeleteUserData ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
// 		return
// 	}

// 	u.sugar.Infoln("DeleteUserData ", config.InfoSplitter, "", config.InfoSplitter, config.ServiceSuccess)
// 	return
// }
