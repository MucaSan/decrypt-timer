package infrastructure

import (
	validation "decrypter-timer/domain/ifaces"
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

func UTF8EnglishAndNumbersAndSpecialCharacters(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9!#$%&'()*+,-.:;<=>?@[\\]^_`{|}~]*$")
	return regex.MatchString(fl.Field().String())
}

var structValidator *validator.Validate

type StructValidator struct{}

func init() {
	structValidator = validator.New()
	structValidator.RegisterValidation("utf8_english_and_numbers_and_special_characters", UTF8EnglishAndNumbersAndSpecialCharacters)
}
func (sv *StructValidator) ValidateStruct(vi validation.ValidationInterface) error {
	return structValidator.Struct(vi)
}
