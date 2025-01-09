package infrastructure

import (
	validation "decrypter-timer/domain/ifaces"
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

var keyLength int
var structValidator *validator.Validate

func UTF8EnglishAndNumbersAndSpecialCharacters(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9!#$%&'()*+,-.:;<=>?@[\\]^_`{|}~]*$")
	return regex.MatchString(fl.Field().String())
}

func CaesarKeyLength(fl validator.FieldLevel) bool {
	if fl.StructFieldName() == "SecretKey" {
		keyLength = len(fl.Field().String())
	}

	if VerifyStructFieldNotAlgorithmAndRightAlgorithmType(fl, "caesar") {
		return true
	}

	return ValidateCaesarKeyLength()

}

func VerifyStructFieldNotAlgorithmAndRightAlgorithmType(fl validator.FieldLevel, fieldValue string) bool {
	return (!(fl.StructFieldName() == "Algorithm") && !(fl.Field().String() == fieldValue)) || ((fl.StructFieldName() == "Algorithm") && !(fl.Field().String() == fieldValue))
}

func ValidateCaesarKeyLength() bool {
	if keyLength == 0 || keyLength > 1 {
		return false
	}

	return true
}

type StructValidator struct{}

func init() {
	structValidator = validator.New()
	structValidator.RegisterValidation("utf8_english_and_numbers_and_special_characters", UTF8EnglishAndNumbersAndSpecialCharacters)
}

func (sv *StructValidator) ValidateStruct(vi validation.ValidationInterface) error {
	return structValidator.Struct(vi)
}
