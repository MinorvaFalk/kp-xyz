package validation

import (
	"encoding/json"
	"errors"
	"kp/pkg/exception"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	v.RegisterValidation("nik", func(fl validator.FieldLevel) bool {
		nik := fl.Field().String()
		regexNik := regexp.MustCompile(`^(1[1-9]|21|[37][1-6]|5[1-3]|6[1-5]|[89][12])\d{2}\d{2}([04][1-9]|[1256][0-9]|[37][01])(0[1-9]|1[0-2])\d{2}\d{4}$`)
		return regexNik.MatchString(nik)
	})

	v.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		date := fl.Field().String()
		_, err := time.Parse(time.DateOnly, date)
		return err == nil
	})

	return &Validator{
		validator: v,
	}
}

func (v *Validator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		var messages []map[string]interface{}

		if errors.As(err, &validationErrors) {
			for _, ve := range validationErrors {
				fieldName := ve.StructField()
				messages = append(messages, map[string]any{
					"field":   fieldName,
					"message": msgForTag(ve.Tag()),
				})
			}

			msg, err := json.Marshal(messages)
			if err != nil {
				return err
			}

			return exception.NewValidatonError(string(msg), validationErrors)
		}

		return err
	}

	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	}
	return ""
}
