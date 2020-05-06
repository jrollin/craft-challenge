// https://blog.andreiavram.ro/abstract-dependencies/
package validator

import (
	"strings"

	validate "github.com/go-playground/validator"
)

// Validator offers struct validation
type Validator interface {
	Validate(s interface{}) error
}

// Error provides validation detailed information
type Error struct {
	Err    error
	Fields []Field
}

// Error returns the string representation of the error validation error
func (e *Error) Error() string {
	return e.Err.Error()
}

// Field representation of each struct field with validation error
type Field struct {
	Name  string      `json:"name"`
	Error string      `json:"error"`
	Value interface{} `json:"value"`
}

type validator struct {
	v *validate.Validate
}

func (v *validator) Validate(s interface{}) (err error) {
	if err = v.v.Struct(s); err == nil {
		return
	}
	if errs, ok := err.(validate.ValidationErrors); ok {
		fields := make([]Field, 0, len(errs))
		for _, e := range errs {
			fields = append(fields, Field{
				Name:  field(e.Namespace()),
				Value: e.Value(),
				Error: translate(e),
			})
		}

		return &Error{
			Err:    err,
			Fields: fields,
		}
	}

	return
}

// New returns a new Validator instance
func NewValidator() Validator {
	return &validator{
		v: validate.New(),
	}
}

func field(s string) (field string) {
	field = s
	if index := strings.Index(s, "."); index > -1 && index < len(s) {
		field = s[index+1:]
	}

	return strings.ToLower(field)
}

func translate(e validate.FieldError) string {
	translations := map[string]string{
		"required": "Is empty.",
		"gte":      "Must be equal to or greater than {param}.",
		"lt":       "Must be lower than {param}.",
		"lte":      "Must be lower than or equal to {param}.",
		"alpha":    "Must contain only letters.",
	}

	if t, ok := translations[e.Tag()]; ok {
		return strings.ReplaceAll(t, "{param}", e.Param())
	}

	return "invalid value"
}
