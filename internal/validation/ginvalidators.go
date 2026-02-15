package validation

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var trans ut.Translator

// Central map: tag → custom message (empty string = default)
var customMessages = map[string]string{
	"e164PhoneNumbers": "{0} must be a valid E.164 phone number",
	"regexp":           "", // fallback to default
}

func RegisterCustomValidations() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	log.Println("Registering custom gin binding validations")

	// Translator setup
	locale := en.New()
	uni := ut.New(locale, locale)
	trans, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	// Register validators
	v.RegisterValidation("e164PhoneNumbers", e164Validator)
	// log.Println("RegisterValidation err", err)

	v.RegisterValidation("regexp", regexpValidator)

	// log.Println("RegisterValidation done ig")

	// Register messages from the map
	for tag, msg := range customMessages {
		if msg != "" {
			registerMessage(v, tag, msg)
		}
	}
}

// Register translation for a tag
func registerMessage(v *validator.Validate, tag string, msg string) {
	v.RegisterTranslation(tag, trans,
		func(ut ut.Translator) error {
			return ut.Add(tag, msg, true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			// Use full nesting path instead of just field name
			t, _ := ut.T(tag, fe.StructNamespace())
			return t
		})
}

// Validators
func regexpValidator(fl validator.FieldLevel) bool {
	// log.Println("regexpValidator")
	pattern := fl.Param()
	value := fl.Field().String()
	// log.Println("regexpValidator pattern, value", pattern, value)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(value)
}

func e164Validator(fl validator.FieldLevel) bool {
	// log.Println("e164Validator")
	value := fl.Field().String()
	// log.Println("e164Validator pattern, value", pattern, value)
	matched, _ := regexp.MatchString(`^\+[1-9][0-9]{1,14}$`, value)
	return matched
}
