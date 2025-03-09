package handler

import (
	"net/http"

	"github.com/404th/Ink/model"
	"github.com/404th/Ink/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var (
		data model.CreatePostRequest
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		var errResp model.Response
		err = helper.ChangeErrorForm(err)
		errResp.Message = helper.SplitErrorMessage(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	resp, err := h.service.PostService().CreatePost(c.Request.Context(), &data)
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

func (u *Handler) GetAllPosts(c *gin.Context) {
	id, _ := c.GetQuery("id")

	resp, err := u.service.PostService().GetAllPosts(
		c.Request.Context(),
		&model.GetAllPostsRequest{
			Id: id,
		},
	)
	if err != nil {
		var errResp model.Response
		u.sugar.Infoln(err)
		err = helper.ChangeErrorForm(err)
		errResp.Message = err.Error()
		errResp.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	if len(resp.Posts) < 1 {
		var errResp model.Response
		errResp.Message = "Foydalanuvchi topilmadi"
		errResp.Data = nil
		u.sugar.Infoln(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
