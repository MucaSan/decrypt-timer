package validation

type ValidationInterface interface {
	IsValid() error
}

type StructValidatorInterface interface {
	ValidateStruct(vs ValidationInterface) error
}
