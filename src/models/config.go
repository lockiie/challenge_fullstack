package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func ValidateIsFK(fl validator.FieldLevel) bool {
	if fl.Top().Type() != fl.Parent().Type() {
		if fl.Field().IsZero() {
			return false
		}
	}
	return true
}

func init() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ = uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	//custom validation
	// validate.RegisterValidation("isFK", ValidateIsFK)

	// validate.RegisterTranslation("isFK", trans, func(ut ut.Translator) error {
	// 	return ut.Add("isFK", "{0} is required", true)
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("isFK", fe.StructNamespace())
	// 	return t
	// })
}

func translateError(err error) error {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	var message string

	for _, e := range validatorErrs {
		if message != "" {
			message += ", "
		}
		message += e.Translate(trans)
	}
	return errors.New(message)
}
