package model

type EncryptData struct {
	SecretKey string `validate:"required,min=1,max=50" json:"secretKey"`
	Message   string `validate:"required,min=1,max=2000" json:"message"`
	Algorithm string `validate:"required, oneof=3des blowfish caesar vigenere"`
}
