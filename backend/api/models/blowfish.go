package models

import (
	"golang.org/x/crypto/blowfish"
	"gorm.io/gorm"
)

type Blowfish struct {
	gorm.Model
	Cipher      *blowfish.Cipher
	Description string
	SecretKey   string
}
