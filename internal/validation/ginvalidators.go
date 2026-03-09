package validation

import (
	"log"
	"regexp"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	log.Println("Registering custom gin binding validations")

	v.RegisterValidation("regexp", regexpValidator)
	v.RegisterValidation("isoDate", isoDateValidator)
	v.RegisterValidation("e164PhoneNumbers", e164Validator)
	v.RegisterValidation("isoDateTime", isoDateTimeValidator)
}

// Validators
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
	value := fl.Field().String()
	matched, _ := regexp.MatchString(`^\+[1-9][0-9]{1,14}$`, value)
	return matched
}

func isoDateValidator(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	_, err := time.Parse("2006-01-02", val)
	return err == nil
}

func isoDateTimeValidator(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	_, err := time.Parse(time.RFC3339, val)
	return err == nil
}
