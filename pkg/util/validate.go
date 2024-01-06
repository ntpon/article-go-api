package util

import "github.com/go-playground/validator/v10"


func GetValidationErrors(err error) map[string]string {
    errors := make(map[string]string)
    for _, err := range err.(validator.ValidationErrors) {
        errors[err.Field()] = err.Tag()
    }
    return errors
}






