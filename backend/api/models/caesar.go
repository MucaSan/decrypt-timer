package models

import (
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
