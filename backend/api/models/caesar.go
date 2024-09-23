package models

import (
	"net/http"
	"strings"

	caesarCipher "github.com/theTardigrade/golang-caesarCipher"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

const EnglishAlphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Caesar struct {
	gorm.Model
	Description string `json: "description" validate:"required max=200, regexp=^[a-zA-Z]*$"`
	SecretKey   string `json: "secretKey" validate:"required len=1, regexp=^[a-zA-Z]*$"`
}

func ValidateCaesar(c Caesar) error {
	if err := validator.Validate(c); err != nil {
		return err
	}
	return nil
}

func ReturnDecryptedDescriptionWithStatusCode(caesarDecrypt *Caesar) (int, string) {
	caesarDecrypt.SecretKey = strings.ToUpper(caesarDecrypt.SecretKey)
	for i, letter := range EnglishAlphabet {
		if string(letter) == caesarDecrypt.SecretKey {
			return http.StatusOK, caesarCipher.Decrypt(caesarDecrypt.Description, uint(i+1))
		}
	}
	return http.StatusBadRequest, "Couldn't parse from provided Secret Key."
}

func ReturnEncryptedDescriptionWithStatusCode(caesarEncrypt *Caesar) (int, string) {
	caesarEncrypt.SecretKey = strings.ToUpper(caesarEncrypt.SecretKey)
	for i, letter := range EnglishAlphabet {
		if string(letter) == caesarEncrypt.SecretKey {
			return http.StatusOK, caesarCipher.Decrypt(caesarEncrypt.Description, uint(i+1))
		}
	}
	return http.StatusBadRequest, "Couldn't parse from provided Secret Key."
}
