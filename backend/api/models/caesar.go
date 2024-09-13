package models

import (
	"gorm.io/gorm"
)

const EnglishAlphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Caesar struct {
	gorm.Model
	Description string `json: "description" validate:"max=200, regexp=^[a-zA-Z]*$`
	SecretKey   string `json: "secretKey" validate:"max=1, regexp=^[a-zA-Z]*$`
}
