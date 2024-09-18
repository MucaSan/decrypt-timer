package controllers

import (
	"api/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EncryptDescription(c *gin.Context) {
	var caesarEncrypt models.Caesar
	c.ShouldBindBodyWithJSON(&caesarEncrypt)
	if err := models.ValidateCaesar(caesarEncrypt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
		return
	}
	statusCode, encryptedDescription := models.ReturnDecryptedDescriptionWithStatusCode(&caesarEncrypt)
	if statusCode != http.StatusBadRequest {
		c.JSON(statusCode, gin.H{"Decrypted description:": encryptedDescription})
	}
	c.JSON(statusCode, gin.H{"Error: ": encryptedDescription})
}
func DecryptDescription(c *gin.Context) {
	var caesarDecrypt models.Caesar
	c.ShouldBindBodyWithJSON(&caesarDecrypt)
	if err := models.ValidateCaesar(caesarDecrypt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
		return
	}
	statusCode, decryptedDescription := models.ReturnDecryptedDescriptionWithStatusCode(&caesarDecrypt)
	if statusCode != http.StatusBadRequest {
		c.JSON(statusCode, gin.H{"Decrypted description:": decryptedDescription})
	}
	c.JSON(statusCode, gin.H{"Error: ": decryptedDescription})
}
