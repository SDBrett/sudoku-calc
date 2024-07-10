package utils

const (
	ErrExceedRange = ValidationErr("The number give exceeds the allowable range")
)

type ValidationErr string

func ValidateNumberRange(min, max, given int) error {
	if given < min || given > max {
		return ErrExceedRange
	}
	return nil
}

func (e ValidationErr) Error() string {
	return string(e)
}
