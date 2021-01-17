package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form - form struct
type Form struct {
	url.Values
	Errors errors
}

// New - Init custom form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required - Check required field
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength - Check max length of field value
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field length cannot exceed %d", d))
	}
}

// PermittedValues - Check permitted values
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	for _, opt := range opts {
		if strings.TrimSpace(value) == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// Valid - Check form is valid
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}