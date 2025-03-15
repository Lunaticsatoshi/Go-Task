package utils

import (
	"math"
	"strconv"

	"github.com/Lunaticsatoshi/go-task/app/common/interfaces"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GeneratePaginationMeta(db *gorm.DB, intPage int, intLimit int, model interface{}) interfaces.PaginationMeta {
	var totalCount int64
	db.Model(model).Count(&totalCount)

	totalPages := int(math.Ceil(float64(totalCount) / float64(intLimit)))

	var prevPage, nextPage *int
	if intPage > 1 {
		pp := intPage - 1
		prevPage = &pp
	}
	if intPage < totalPages {
		np := intPage + 1
		nextPage = &np
	}

	return interfaces.PaginationMeta{
		CurrentPage: intPage,
		PrevPage:    prevPage,
		NextPage:    nextPage,
		TotalPages:  totalPages,
		TotalCount:  totalCount,
		Limit:       intLimit,
	}
}

func GetRequestPaginationData(ctx *gin.Context) (int, int, int, string, string) {
	var page = ctx.DefaultQuery("p", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	sortKey := ctx.DefaultQuery("sort_key", "created_at")
	sortOrder := ctx.DefaultQuery("sort_order", "desc")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	if sortKey != "created_at" && sortKey != "updated_at" {
		sortKey = "created_at"
	}

	return intPage, intLimit, offset, sortKey, sortOrder
}
