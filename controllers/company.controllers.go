package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"self-payroll/models"
	"self-payroll/utils"
)

func GetCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		company := new(models.Company)
		if err := dbClient.WithContext(ctx).First(company).Error; err != nil {
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
			"data": utils.CompanyUtils{
				Name:      company.Name,
				Address:   company.Address,
				Balance:   company.Balance,
				ID:        int(company.ID),
				CreatedAt: company.CreatedAt.String(),
				UpdatedAt: company.UpdatedAt.String(),
			},
		})
	}
}

func CreateOrUpdateCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		companyModel := models.Company{}
		newCompany := models.Company{}
		if err := ctx.BindJSON(&newCompany); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := dbClient.WithContext(ctx).First(&companyModel).Error; err != nil {
			if err := dbClient.WithContext(ctx).Create(&newCompany).Find(&companyModel).Error; err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"success": true,
					"message": "success",
					"data": utils.CompanyUtils{
						Name:      newCompany.Name,
						Address:   newCompany.Address,
						Balance:   newCompany.Balance,
						ID:        int(newCompany.ID),
						CreatedAt: newCompany.CreatedAt.String(),
						UpdatedAt: newCompany.UpdatedAt.String(),
					},
				})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}

		if err := dbClient.Model(&companyModel).Updates(&newCompany).Error; err != nil {
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
			"data": utils.CompanyUtils{
				Name:      companyModel.Name,
				Address:   companyModel.Address,
				Balance:   companyModel.Balance,
				ID:        int(companyModel.ID),
				CreatedAt: companyModel.CreatedAt.String(),
				UpdatedAt: companyModel.UpdatedAt.String(),
			},
		})
	}

}

func TopUpCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		company := new(models.Company)
		topup := utils.TopUpUtils{}
		if err := ctx.BindJSON(&topup); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		if err := dbClient.WithContext(ctx).First(company).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "failed",
				"data":    "",
			})
			return
		}
		tmp := models.Company{
			Balance: topup.Balance + company.Balance,
		}
		if err := dbClient.Model(&company).Updates(&tmp).Error; err != nil {
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
			"data": utils.CompanyUtils{
				Name:      company.Name,
				Address:   company.Address,
				Balance:   company.Balance,
				ID:        int(company.ID),
				CreatedAt: company.CreatedAt.String(),
				UpdatedAt: company.UpdatedAt.String(),
			},
		})

	}
}
