package controllers

import (
	"api/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/h5law/vinegar"
)

func VigenereEncryptDescription(ctx *gin.Context) {
	var CipherVigenere models.Vigenere
	ctx.ShouldBindBodyWithJSON(&CipherVigenere)
	CipherVigenere.Vigenere = vinegar.NewVigenere("kyrptos")
	encryptedDescription := CipherVigenere.Vigenere.Encrypt(CipherVigenere.Description, CipherVigenere.SecretKey)
	ctx.JSON(http.StatusOK, gin.H{"Encrypted Description: ": encryptedDescription})
}
