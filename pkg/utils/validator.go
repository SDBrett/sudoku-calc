package utils

import (
	"strings"
)

const (
	ErrExceedRange       = ValidationErr("The number give exceeds the allowable range")
	ErrInvalidSearchType = ValidationErr("Invalid search type, only 'valueonly' or 'full' permitted")
)

type ValidationErr string

func ValidateNumberRange(min, max, given int) error {
	if given < min || given > max {
		return ErrExceedRange
	}
	return nil
}

func ValidateSearchType(s string) error {
	s = strings.ToLower(s)
	if strings.ToLower(s) != "valueonly" && strings.ToLower(s) != "full" {
		return ErrInvalidSearchType
	}
	return nil
}

func (e ValidationErr) Error() string {
	return string(e)
}
