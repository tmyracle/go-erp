package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmyracle/go-erp/database"
	"github.com/tmyracle/go-erp/models"
)

func GetAccounts(c *gin.Context) {
	var accounts []models.Account
	if err := database.DB.Find(&accounts).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retreive accounts"})
		return
	} else {
		c.JSON(200, accounts)
	}
}

func GetAccount(c *gin.Context) {
	var account models.Account
	id := c.Param("id")
	if err := database.DB.First(&account, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	} else {
		c.JSON(http.StatusOK, account)
	}
}

func CreateAccount(c *gin.Context) {
	var account models.Account

	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, account)
	}
}
	