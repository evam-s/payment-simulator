package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

func RegisterCustomValidations() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	log.Println("Registering custom gin binding validations")

	v.RegisterValidation("e164PhoneNumbers", e164Validator)
	v.RegisterValidation("regexp", regexpValidator)
}

func regexpValidator(fl validator.FieldLevel) bool {
	pattern := fl.Param()
	value := fl.Field().String()

	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(value)
}

func e164Validator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// E.164: + followed by 1–15 digits
	matched, _ := regexp.MatchString(`^\+[1-9][0-9]{1,14}$`, phone)
	return matched
}
