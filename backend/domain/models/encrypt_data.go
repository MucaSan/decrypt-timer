package model

import validation "decrypter-timer/domain/ifaces"

type EncryptData struct {
	SecretKey string                              `validate:"required,utf8_english_and_numbers_and_special_characters,min=1,max=100" json:"secretKey"`
	Message   string                              `validate:"required,utf8_english_and_numbers_and_special_characters,min=1,max=2000" json:"message"`
	Algorithm string                              `validate:"required,utf8_english_and_numbers_and_special_characters, oneof=3des blowfish caesar vigenere" json:"algorithm"`
	validate  validation.StructValidatorInterface `validate: "-"`
}

func (e *EncryptData) IsValid() error {
	return e.validate.ValidateStruct(e)
}

func (e *EncryptData) InjectValidatorInterface(va validation.StructValidatorInterface) {
	e.validate = va
}
