package validation

type ValidationInterface interface {
	IsValid() error
	InjectValidatorInterface(va StructValidatorInterface)
}

type StructValidatorInterface interface {
	ValidateStruct(vi ValidationInterface) error
}
