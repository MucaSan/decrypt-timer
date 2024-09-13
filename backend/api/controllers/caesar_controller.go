package controllers

import (
	"api/api/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	caesarCipher "github.com/theTardigrade/golang-caesarCipher"
)

func EncryptDescription(c *gin.Context) {
	var caesarEncrypt models.Caesar
	c.ShouldBindBodyWithJSON(&caesarEncrypt)
	if err := models.ValidateCaesar(caesarEncrypt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
		return
	}
	caesarEncrypt.SecretKey = strings.ToUpper(caesarEncrypt.SecretKey)
	for i, letter := range models.EnglishAlphabet {
		if string(letter) == caesarEncrypt.SecretKey {
			c.JSON(http.StatusOK, gin.H{"Encrypted message": caesarCipher.Encrypt(caesarEncrypt.Description, uint(i+1))})
			fmt.Println("Passing here!")
		}
	}

}
