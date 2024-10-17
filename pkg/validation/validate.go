package validation

import "errors"

func ValidateName(name string) (error, string) {
	if len(name) == 0 {
		return errors.New("name cannot be empty"), ""
	}

	return nil, name
}
