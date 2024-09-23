package models

import (
	"github.com/h5law/vinegar"
	"gorm.io/gorm"
)

type Vigenere struct {
	gorm.Model
	Vigenere    vinegar.Vigenere
	Description string
	SecretKey   string
}
