package helper

import (
	"errors"
	"strconv"

	"github.com/404th/Ink/config"
	"github.com/gin-gonic/gin"
)

func MakeLimit(c *gin.Context) (int32, error) {
	limit, exists := c.GetQuery("limit")
	if !exists {
		limit = config.DefaultLimit
	}

	var resultLimit int32

	nL2Int, err := strconv.Atoi(limit)
	if err != nil {
		return 0, errors.New("Limit noto'g'ri tipda kiritilgan")
	}

	switch {
	case nL2Int > 100:
		nL2Int = 100
	case nL2Int < 1:
		nL2Int = 10
	}

	resultLimit = int32(nL2Int)

	return resultLimit, nil
}

func MakePage(c *gin.Context) (int32, error) {
	page, exists := c.GetQuery("page")
	if !exists {
		page = config.DefaultPage
	}

	var resultPage int32

	nL2Int, err := strconv.Atoi(page)
	if err != nil {
		return 0, errors.New("Page noto'g'ri tipda kiritilgan")
	}

	switch {
	case nL2Int < 1:
		nL2Int = 10
	}

	resultPage = int32(nL2Int)

	return resultPage, nil
}
