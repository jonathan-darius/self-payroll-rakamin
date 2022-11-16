package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"self-payroll/models"
	"self-payroll/utils"
	"strconv"
)

func CreateEmployees() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var employees models.User

		if err := ctx.BindJSON(&employees); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		var position models.Position
		if err := dbClient.Where("id = ?", employees.PositionID).First(&position).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}

		employees.Position = &position
		if err := dbClient.Create(&employees).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"success": false,
				"message": "failed",
				"data":    err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
			"data": utils.UserUtils{
				ID:         int(employees.ID),
				CreatedAt:  employees.CreatedAt.String(),
				UpdatedAt:  employees.UpdatedAt.String(),
				SecretID:   employees.SecretID,
				Name:       employees.Name,
				Email:      employees.Email,
				Phone:      employees.Phone,
				Address:    employees.Address,
				PositionID: employees.PositionID,
				Position: &utils.PositionUtils{
					Name:      employees.Position.Name,
					Salary:    employees.Position.Salary,
					ID:        int(employees.Position.ID),
					CreatedAt: employees.Position.CreatedAt.String(),
					UpdatedAt: employees.Position.UpdatedAt.String(),
				},
			},
		})
	}
}

func DeleteEmployees() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if err := dbClient.First(&models.User{}, id).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		if err := dbClient.Delete(&models.User{}, id).Error; err != nil {
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

func AllEmployees() gin.HandlerFunc {
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
		var data []models.User
		query := dbClient.Limit(limit).Offset(offset)
		res := query.Model(&models.User{}).Find(&data)
		if res.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  "Failed",
				"Message": "Data Not Found",
			})
			return
		}
		var position models.Position
		var response []utils.UserUtils
		for _, datum := range data {

			if err := dbClient.Where("id = ?", datum.PositionID).First(&position).Error; err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"message": "failed",
					"data":    "",
				})
				return
			}
			datum.Position = &position
			response = append(response, utils.UserUtils{
				ID:         int(datum.ID),
				CreatedAt:  datum.CreatedAt.String(),
				UpdatedAt:  datum.UpdatedAt.String(),
				SecretID:   datum.SecretID,
				Name:       datum.Name,
				Email:      datum.Email,
				Phone:      datum.Phone,
				Address:    datum.Address,
				PositionID: datum.PositionID,
				Position: &utils.PositionUtils{
					Name:      datum.Position.Name,
					Salary:    datum.Position.Salary,
					ID:        int(datum.Position.ID),
					CreatedAt: datum.Position.CreatedAt.String(),
					UpdatedAt: datum.Position.UpdatedAt.String(),
				},
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
			"data":    response,
		})
	}
}

func DetailEmployees() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var employeesDetail models.User
		if err := dbClient.Where("id = ?", id).First(&employeesDetail).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		var position models.Position
		if err := dbClient.Where("id = ?", employeesDetail.PositionID).First(&position).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		employeesDetail.Position = &position
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "success",
			"data":    employeesDetail,
		})
	}
}

func WithdrawBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var withDraw utils.WithdrawRequest
		if err := ctx.BindJSON(&withDraw); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		var employeesDetail models.User
		if err := dbClient.Where("id = ?", withDraw.Id).First(&employeesDetail).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		if employeesDetail.SecretID != withDraw.Secret_id {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		var position models.Position
		if err := dbClient.Where("id = ?", employeesDetail.PositionID).First(&position).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		salary := position.Salary
		company := new(models.Company)
		if err := dbClient.WithContext(ctx).First(company).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		if company.Balance < salary {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		tmp := models.Company{
			Balance: company.Balance - salary,
		}
		if err := dbClient.Model(&company).Updates(&tmp).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		trans := new(models.Transaction)
		trans.Amount = salary
		trans.Note = "Withdraw Salary"
		trans.Type = "debit"
		if err := dbClient.Create(&trans).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Success withdraw salary",
			"data":    "",
		})
	}
}
