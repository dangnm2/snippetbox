package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

//Form holds form data and validation errors
type Form struct {
	url.Values
	Errors errors
}

//New initializes a custom Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required check that specific fields are present and not blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//MaxLength check that a field in the form contains a maximum number of characters
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field cannot be longer than %d characters", d))
	}
}

//PermittedValues checks that a field is in a list of values
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

//Valid returns true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}