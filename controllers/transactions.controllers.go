package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"self-payroll/models"
	"strconv"
)

func TransactionsGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit, err := strconv.Atoi(ctx.Query("limit"))

		if err != nil || limit < 1 {
			limit = 10
		}
		skip, err1 := strconv.Atoi(ctx.Query("skip"))
		if err1 != nil || skip < 1 {
			skip = 1
		}
		offset := (skip - 1) * limit
		var transaction []*models.Transaction
		query := dbClient.Limit(limit).Offset(offset)
		res := query.Model(&models.Transaction{}).Find(&transaction)
		if res.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  "Failed",
				"Message": "Data Not Found",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
			"data":    transaction,
		})
	}

}
