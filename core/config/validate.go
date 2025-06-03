package config

import "github.com/meta-atman/metarpc/core/validation"

// validate validates the value if it implements the Validator interface.
func validate(v any) error {
	if val, ok := v.(validation.Validator); ok {
		return val.Validate()
	}

	return nil
}
