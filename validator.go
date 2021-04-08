package validator

import (
	"reflect"
)

type Validator struct {
	name string
	fields
}

var separator = ";"

func NewValidator(s interface{}) Validator {
	return Validator{
		name:   reflect.TypeOf(s).Elem().Name(),
		fields: getFields(s),
	}
}

func (v *Validator) AddValues(s interface{}) {
	for i, _ := range v.fields {
		v.fields[i].value = reflect.ValueOf(s).Elem().Field(i)
	}
}

func (v *Validator) GetName() string {
	return v.name
}
