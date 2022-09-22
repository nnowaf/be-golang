package utils

import (
	"gobe/api/model"
	"math"

	"github.com/go-playground/validator/v10"
)

func GetEnd(pageNumber int, size int) int {
	if pageNumber == 0 {
		return 0
	}

	return int(math.Abs(float64(((pageNumber-1)*size + 1) - 1)))
}

func ValidateStruct(err error) []*model.ErrorResponse {
	var errors []*model.ErrorResponse
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)

		}
	}

	return errors
}
