package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/404th/Ink/config"
	"github.com/404th/Ink/model"
	"github.com/gin-gonic/gin"
)

// ------------------- Tigres (SDK) HTTP Handlers -------------------

// UploadImageHandler handles image upload requests.
func (h *Handler) UploadImageHandler(c *gin.Context) {
	// Limit the size of the incoming request.
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, config.MaxImageSize+1024) // 5MB

	file, err := c.FormFile("file")
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = fmt.Sprintf("Ruxsat etilgan maksimal rasm hajmi: %dMB\n", config.MaxImageSize)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	if file.Size > config.MaxImageSize {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = fmt.Sprintf("Ruxsat etilgan maksimal rasm hajmi: %dMB.\n", config.MaxImageSize)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !config.AllowedImageExt[ext] {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Ruxsat etilgan rasm formatlari: jpg, jpeg, png, gif, bmp\n"
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	f, err := file.Open()
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Yuklangan ma'lumotni ochishni iloji bo'lmadi\n"
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}
	defer f.Close()

	// Generate a unique key.
	key := fmt.Sprintf("img_%d_%s", time.Now().Unix(), file.Filename)

	storedLocation, err := h.tg.UploadImage(c.Request.Context(), key, f)
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Ma'lumotni yuklashni iloji bo'lmadi\n"
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}

	var response model.SuccessResponse
	response.Data = gin.H{"file": storedLocation}
	response.Message = "Muvaffaqiyatli yuklandi"
	c.JSON(http.StatusOK, response)
}

// UploadVideoHandler handles video upload requests.
func (h *Handler) UploadVideoHandler(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, config.MaxVideoSize+1024) // 20MB

	file, err := c.FormFile("file")
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = fmt.Sprintf("Ruxsat etilgan maksimal video hajmi: %dMB.\n", config.MaxVideoSize)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	if file.Size > config.MaxVideoSize {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = fmt.Sprintf("Ruxsat etilgan maksimal video hajmi: %dMB\n", config.MaxVideoSize)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !config.AllowedVideoExt[ext] {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Ruxsat etilgan video formatlari: mp4, mov, avi, wmv, mkv\n"
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	f, err := file.Open()
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Yuklangan ma'lumotni ochishni iloji bo'lmadi\n"
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}
	defer f.Close()

	key := fmt.Sprintf("vid_%d_%s", time.Now().Unix(), file.Filename)

	storedLocation, err := h.tg.UploadVideo(c.Request.Context(), key, f)
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Ma'lumotni yuklashni iloji bo'lmadi\n"
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}

	var response model.SuccessResponse
	response.Data = gin.H{"file": storedLocation}
	response.Message = "Muvaffaqiyatli yuklandi"
	c.JSON(http.StatusOK, response)
}

// UploadFileHandler handles generic file upload requests.
func (h *Handler) UploadFileHandler(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, config.MaxFileSize+1024) // 20MB

	file, err := c.FormFile("file")
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = fmt.Sprintf("Ruxsat etilgan maksimal fayl hajmi: %dMB\n", config.MaxFileSize)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	if file.Size > config.MaxFileSize {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = fmt.Sprintf("Ruxsat etilgan maksimal fayl hajmi: %dMB.\n", config.MaxFileSize)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !config.AllowedFileExt[ext] {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Ruxsat etilgan fayl formatlari: pdf, xlsx, xls, doc, docx, ppt, pptx\n"
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	f, err := file.Open()
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Yuklangan ma'lumotni ochishni iloji bo'lmadi\n"
		c.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}
	defer f.Close()

	key := fmt.Sprintf("file_%d_%s", time.Now().Unix(), file.Filename)

	storedLocation, err := h.tg.UploadFile(c.Request.Context(), key, f)
	if err != nil {
		var errResp model.ErrorResponse
		errResp.Data = nil
		errResp.Message = "Ma'lumotni yuklashni iloji bo'lmadi\n"
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}

	var response model.SuccessResponse
	response.Data = gin.H{"file": storedLocation}
	response.Message = "Muvaffaqiyatli yuklandi"
	c.JSON(http.StatusOK, response)
}
