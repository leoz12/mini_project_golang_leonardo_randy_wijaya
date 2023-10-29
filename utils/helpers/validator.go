package helpers

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

func ReqeustValidator(s interface{}) []string {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	enTranslations.RegisterDefaultTranslations(validate, trans)

	var errorMessages []string
	err := validate.Struct(s)
	for _, val := range translateError(err, trans) {
		errorMessages = append(errorMessages, val.Error())
	}

	return errorMessages
}
