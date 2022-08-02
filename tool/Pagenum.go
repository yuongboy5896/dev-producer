package tool

import (
	"dev-producer/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaseUrl(ctx *gin.Context) (model.DaoPage, error) {
	var dp model.DaoPage
	limit := ctx.DefaultQuery("pagesize", "")
	pageNumber := ctx.DefaultQuery("pagenum", "")
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 0 {
		return dp, nil
	}
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil || pageNumberInt < 0 {
		return dp, nil
	}
	if pageNumberInt != 0 {
		pageNumberInt--
	}
	offsetInt := limitInt * pageNumberInt
	dp.Pagenum = limitInt + 1
	dp.Pagesize = offsetInt

	return dp, err
}
