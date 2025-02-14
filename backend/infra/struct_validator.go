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

func ValidateCaesarKeyLength(fl validator.FieldLevel) bool {
	fieldName := fl.StructFieldName()
	fieldValue := fl.Field()
	fieldContent := fieldValue.String()
	if fieldName == "Algorithm" && !(CheckIfAlgorithmIsValid(fieldContent)) {
		return false
	}

	if fieldName == "SecretKey" && fieldValue.String() == "caesar" {
		keyLength = len(fieldValue.String())
	}

	if keyLength < 0 {
		return false
	}

	if keyLength == 0 || keyLength > 1 {
		return false
	}

	return true
}

func CheckIfAlgorithmIsValid(algorithm string) bool {
	validAlgorithms := []string{"caesar", "3des", "blowfish", "vigenere"}
	for _, validAlgorithm := range validAlgorithms {
		if algorithm == validAlgorithm {
			return true
		}
	}
	return false
}

type StructValidator struct{}

func init() {
	structValidator = validator.New()
	structValidator.RegisterValidation("utf8_english_and_numbers_and_special_characters", UTF8EnglishAndNumbersAndSpecialCharacters)
}

func (sv *StructValidator) ValidateStruct(vi validation.ValidationInterface) error {
	return structValidator.Struct(vi)
}
