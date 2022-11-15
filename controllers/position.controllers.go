package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"self-payroll/databases"
	"self-payroll/models"
	"self-payroll/utils"
	"strconv"
)

var (
	dbClient = databases.DBClient
)

func CreatePosition() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var position models.Position
		if err := ctx.BindJSON(&position); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		if err := dbClient.Create(&position).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"success": false,
				"message": "failed",
				"data":    nil,
			})
			return
		}
		response := &utils.PositionUtils{
			Name:      position.Name,
			Salary:    position.Salary,
			ID:        int(position.ID),
			CreatedAt: position.CreatedAt.String(),
			UpdatedAt: position.UpdatedAt.String(),
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
			"data":    response,
		})
	}
}

func DeletePosition() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if err := dbClient.First(&models.Position{}, id).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		if err := dbClient.Delete(&models.Position{}, id).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
			"data":    "",
		})
	}
}

func SearchPosition() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var position models.Position
		id := ctx.Param("id")
		if err := dbClient.Where("id = ?", id).First(&position).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  "Failed",
				"Message": "Data Not Found",
			})
			return
		}
		response := &utils.PositionUtils{
			Name:      position.Name,
			Salary:    position.Salary,
			ID:        int(position.ID),
			CreatedAt: position.CreatedAt.String(),
			UpdatedAt: position.UpdatedAt.String(),
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
			"data":    response,
		})
	}
}

func AllPosition() gin.HandlerFunc {
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
		var position []*models.Position
		query := dbClient.Limit(limit).Offset(offset)
		res := query.Model(&models.Position{}).Find(&position)
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
			"data":    position,
		})
	}
}
