package models

import (
	"github.com/lauslim12/vigenere"
	"gorm.io/gorm"
)

type Vigenere struct {
	gorm.Model
	Cipher      *vigenere.Vigenere
	Description string
	SecretKey   string
}
