package validator

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v *validator.Validate
}

func New() *Validator {
	return &Validator{
		v: validator.New(),
	}
}

func (va *Validator) Validate(_ context.Context, s any) error {
	err := va.v.Struct(s)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			es := make([]string, len(errors))
			for i, e := range errors {
				es[i] = e.Error()
			}
			return fmt.Errorf("validation errors: %v", es)
		} else {
			return fmt.Errorf("validation failed: %v", err)
		}
	}
	return nil
}
