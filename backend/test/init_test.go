
package test

import (
    "github.com/go-playground/validator/v10"
)

// ประกาศ validate ในที่เดียว
var validate *validator.Validate

func init() {
    validate = validator.New()
}
