package handler

import (
	"net/http"
	"time"

	"github.com/404th/Ink/model"
	"github.com/404th/Ink/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (h *Handler) SignupUser(c *gin.Context) {
	var (
		data model.SignupUserRequest
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		var errResp model.Response
		err = helper.ChangeErrorForm(err)
		errResp.Message = helper.SplitErrorMessage(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	resp, err := h.service.UserService().SignupUser(c.Request.Context(), &data)
	if err != nil {
		var errResp model.Response
		err = helper.ChangeErrorForm(err)
		errResp.Message = helper.SplitErrorMessage(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}

	var response model.Response

	response.Success = true
	response.Data = resp
	response.Message = "Muvaffaqiyatli yaratildi"

	c.JSON(http.StatusOK, response)
}

func (u *Handler) LoginUser(c *gin.Context) {
	var (
		data model.LoginUserRequest
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		var errResp model.Response
		u.sugar.Infoln(err)
		err = helper.ChangeErrorForm(err)
		errResp.Message = "Ma'lumotlar noto'g'ri kiritilgan"
		errResp.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	if data.Password == "" || data.Username == "" {
		var errResp model.Response
		errResp.Message = "Ma'lumotlar to'liq kiritilishi kerak!"
		errResp.Data = errResp
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	resp, err := u.service.UserService().LoginUser(c.Request.Context(), &data)
	if err != nil {
		var errResp model.Response
		u.sugar.Infoln(err)
		err = helper.ChangeErrorForm(err)
		errResp.Message = err.Error()
		errResp.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	if resp.Id == "" {
		var errResp model.Response
		u.sugar.Infoln(err)
		errResp.Message = "Foydalanuvchi topilmadi"
		errResp.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (u *Handler) GetUser(c *gin.Context) {
	id, exists := c.GetQuery("id")
	if !exists {
		var errResp model.Response
		errResp.Message = "Noto'g'ri ma'lumot kiritildi"
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	resp, err := u.service.UserService().GetUser(c.Request.Context(), &model.Id{Id: id})
	if err != nil {
		var errResp model.Response
		u.sugar.Infoln(err)
		err = helper.ChangeErrorForm(err)
		errResp.Message = err.Error()
		errResp.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	if resp == nil {
		var errResp model.Response
		errResp.Message = "Foydalanuvchi topilmadi"
		errResp.Data = nil
		u.sugar.Infoln(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (u *Handler) HandleRefreshJWT(c *gin.Context) {
	var data model.HandleRefreshJWTRequest

	if err := c.ShouldBindJSON(&data); err != nil || len(data.RefreshToken) < 1 {
		var errResp model.ErrorResponse
		errResp.Message = "Noto'g'ri ma'lumot yuborildi"
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	// Parse and verify the refresh token
	refreshToken, err := jwt.Parse(data.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(u.cfg.RefreshTokenSecret), nil
	})

	if err != nil || !refreshToken.Valid {
		var errResp model.ErrorResponse
		errResp.Message = "Tizimga qayta kirishga urinib ko'ring"
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
		return
	}

	// Generate a new access token if refresh token is valid
	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		_x_data_username := claims["_x_data_username"].(string)
		_x_data_id := claims["_x_data_id"].(string)

		newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"_x_data_username": _x_data_username,
			"_x_data_id":       _x_data_id,
			"exp":              time.Now().Add(time.Minute * time.Duration(u.cfg.AccessTokenExpiryMinute)).Unix(),
		})

		newAccessTokenString, err := newAccessToken.SignedString([]byte(u.cfg.AccessTokenSecret))
		if err != nil {
			var errResp model.ErrorResponse
			errResp.Message = "Tizimga qayta kirishga urinib ko'ring. Ichki xatolik yuzaga keldi"
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
			return
		}

		var resp model.SuccessResponse
		resp.Message = "Refresh token muvaffaqiyatli yaratildi"
		var respData model.HandleRefreshJWTResponse
		respData.NewAccessToken = newAccessTokenString
		resp.Data = respData
		c.JSON(http.StatusOK, resp)
	} else {
		var errResp model.ErrorResponse
		errResp.Message = "Tizimga qayta kirishga urinib ko'ring"
		c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
		return
	}
}

// func (u *Handler) GetAllUsers(c *gin.Context) {
// 	var (
// 		data model.GetAllUsersRequest
// 	)

// 	limit, err := helper.MakeLimit(c)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	page, err := helper.MakePage(c)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	id := c.Query("id")
// 	userStatusId := c.Query("user_status_id")
// 	userRoleId := c.Query("user_role_id")
// 	username := c.Query("username")
// 	passportNumber := c.Query("passport_number")
// 	passportPinfl := c.Query("passport_pinfl")

// 	if len(id) > 0 && !helper.IsValidUUIDv4(c.Query("id")) {
// 		var errResp model.ErrorResponse
// 		errResp.Message = "Noto'g'ri ID kiritilgan"
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
// 		return
// 	}

// 	data.Metadata.Limit = limit
// 	data.Metadata.Page = page
// 	data.Metadata.Count = 0
// 	data.Id = id
// 	data.UserStatusId = userStatusId
// 	data.UserRoleId = userRoleId
// 	data.Username = username
// 	data.PassportNumber = passportNumber
// 	data.PassportPinfl = passportPinfl

// 	resp, err := u.service.UserService().GetAllUsers(c.Request.Context(), &data)
// 	if err != nil {
// 		var errResp model.ErrorResponse
// 		err = helper.ChangeErrorForm(err)
// 		errResp.Message = helper.SplitErrorMessage(err.Error())
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
// 		return
// 	}

// 	var response model.SuccessResponse

// 	response.Data = resp
// 	response.Message = "Muvaffaqiyatli bajarildi"

// 	c.JSON(http.StatusOK, response)
// }
