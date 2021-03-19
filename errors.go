package validator

import (
	"errors"
	"fmt")


func ErrMustBeMore(name string, value interface{}) error {
	return errors.New(fmt.Sprintf("%v must be more than %v", name, value))
}

func ErrMustBeLess(name string, value interface{}) error {
	return errors.New(fmt.Sprintf("%v must be less than %v", name, value))
}

func ErrMustBeEqual(name string, value interface{}) error {
	return errors.New(fmt.Sprintf("%v must be equal %v", name, value))
}

func ErrMustNotBeEqual(name string, value interface{}) error {
	return errors.New(fmt.Sprintf("%v must not be equal %v", name, value))
}

func ErrLenTooShort(name string) error {
	return errors.New(fmt.Sprintf("%v is too short", name))
}

func ErrLenTooLong(name string) error {
	return errors.New(fmt.Sprintf("%v is too long", name))
}

func ErrLenMustBe(name string, value interface{}) error {
	return errors.New(fmt.Sprintf("%v must be %v characters long", name, value))
}

func ErrInvalidFormat(name string) error {
	return errors.New(fmt.Sprintf("%v have invalid format", name))
}

func ErrRequired(name string) error {
	return errors.New(fmt.Sprintf("%v is required", name))
}