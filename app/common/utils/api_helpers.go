package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"

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

func DynamicFilterTasks(c *gin.Context) (string, []interface{}) {
	// Get all query parameters dynamically
	queryParams := c.Request.URL.Query()
	whereClauses := []string{}
	args := []interface{}{}

	// Iterate over query parameters and apply filters
	for key, values := range queryParams {
		// Ignore invalid fields that don’t exist in the User model
		if !isValidField(key) {
			continue
		}

		// Use only the first value (ignoring multiple values for simplicity)
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", key))
		args = append(args, values[0])
	}
	return strings.Join(whereClauses, " AND "), args
}

// Function to validate fields against the User model
func isValidField(field string) bool {
	allowedFields := map[string]bool{
		"id":     true,
		"title":  true,
		"status": true,
	}
	return allowedFields[field]
}
