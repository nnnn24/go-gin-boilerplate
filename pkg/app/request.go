package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/nnnn24/go-gin-boilerplate/pkg/logging"
)

func MarkErrors(errors []*validator.ValidationErrors) {
	for _, err := range errors {
		logging.Info(err.Error())
	}
}
