package validator

import (
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type field struct {
	name  string
	value interface{}
	tag   string
}

type fields []field

func (v Validator) Validate() error {
	var err error

	for _, field := range v.fields {
		field.name = strings.ToLower(field.name)

		switch field.value.(type) {
		case int:
			err = validateInt(field)
			if err != nil {
				return err
			}

		case int64:
			err = validateInt64(field)
			if err != nil {
				return err
			}

		case float64:
			err = validateFloat(field)
			if err != nil {
				return err
			}

		case string:
			err = validateString(field)
			if err != nil {
				return err
			}
		}

	}

	return err
}

func validateInt(field field) error {
	more, err := regexp.Compile("^>\\d+")
	less, err := regexp.Compile("^<\\d+")
	equal, err := regexp.Compile("^=\\d+")
	notEqual, err := regexp.Compile("^!=\\d+")

	for _, validator := range strings.Split(field.tag, separator) {

		if more.MatchString(validator) {
			s := strings.Replace(validator, ">", "", -1)
			x, _ := strconv.Atoi(s)

			if field.value.(int) <= x {
				return ErrMustBeMore(field.name, x)
			}
		}

		if less.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.Atoi(s)

			if field.value.(int) >= x {
				return ErrMustBeLess(field.name, x)
			}
		}

		if equal.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.Atoi(s)

			if field.value.(int) != x {
				return ErrMustBeEqual(field.name, x)
			}
		}

		if notEqual.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.Atoi(s)

			if field.value.(int) == x {
				return ErrMustNotBeEqual(field.name, x)
			}
		}
	}

	return err
}

func validateInt64(field field) error {
	more, err := regexp.Compile("^>\\d+")
	less, err := regexp.Compile("^<\\d+")
	equal, err := regexp.Compile("^=\\d+")
	notEqual, err := regexp.Compile("^!=\\d+")

	for _, validator := range strings.Split(field.tag, separator) {

		if more.MatchString(validator) {
			s := strings.Replace(validator, ">", "", -1)
			x, _ := strconv.ParseInt(s, 10, 64)

			if field.value.(int64) <= x {
				return ErrMustBeMore(field.name, x)
			}
		}

		if less.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseInt(s, 10, 64)

			if field.value.(int64) >= x {
				return ErrMustBeLess(field.name, x)
			}
		}

		if equal.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseInt(s, 10, 64)

			if field.value.(int64) != x {
				return ErrMustBeEqual(field.name, x)
			}
		}

		if notEqual.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseInt(s, 10, 64)

			if field.value.(int64) == x {
				return ErrMustNotBeEqual(field.name, x)
			}
		}
	}

	return err
}

func validateFloat(field field) error {
	more, err := regexp.Compile("^>\\d+.\\d+")
	less, err := regexp.Compile("^<\\d+.\\d+")
	equal, err := regexp.Compile("^=\\d+.\\d+")
	notEqual, err := regexp.Compile("^!=\\d+.\\d+")

	for _, validator := range strings.Split(field.tag, separator) {

		if more.MatchString(validator) {
			s := strings.Replace(validator, ">", "", -1)
			x, _ := strconv.ParseFloat(s, 64)

			if field.value.(float64) <= x {
				return ErrMustBeMore(field.name, x)
			}
		}

		if less.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseFloat(s, 64)

			if field.value.(float64) >= x {
				return ErrMustBeLess(field.name, x)
			}
		}

		if equal.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseFloat(s, 64)

			if field.value.(float64) != x {
				return ErrMustBeEqual(field.name, x)
			}
		}

		if notEqual.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseFloat(s, 64)

			if field.value.(float64) == x {
				return ErrMustNotBeEqual(field.name, x)
			}
		}
	}

	return err
}

func validateString(field field) error {
	if strings.Contains(field.tag, "required") && field.value.(string) == "" {
		return ErrRequired(field.name)
	}

	lenMore, err := regexp.Compile("^len>\\d+")
	lenLess, err := regexp.Compile("^len<\\d+")
	lenEqual, err := regexp.Compile("^len=\\d+")
	formatEqual, err := regexp.Compile("^format=.+")

	if err != nil {
		return err
	}

	for _, validator := range strings.Split(field.tag, separator) {
		runeVal := []rune(field.value.(string))

		if lenMore.MatchString(validator) {
			s := strings.Replace(validator, "len>", "", -1)
			x, _ := strconv.Atoi(s)

			if len(runeVal) <= x {
				return ErrLenTooShort(field.name)
			}
		}

		if lenLess.MatchString(validator) {
			s := strings.Replace(validator, "len<", "", -1)
			x, _ := strconv.Atoi(s)

			if len(runeVal) >= x {
				return ErrLenTooLong(field.name)
			}
		}

		if lenEqual.MatchString(validator) {
			s := strings.Replace(validator, "len=", "", -1)
			x, _ := strconv.Atoi(s)

			if len(runeVal) != x {
				return ErrLenMustBe(field.name, x)
			}
		}

		if formatEqual.MatchString(validator) {
			s := strings.Replace(validator, "format=", "", -1)

			format, inFormats := formats[s]

			if !inFormats {
				format = s
			}

			formatRegex, err := regexp.Compile(format)

			if err != nil {
				return err
			}

			if !formatRegex.MatchString(field.value.(string)) && len(runeVal) > 0 {
				return ErrInvalidFormat(field.name)
			}
		}
	}

	return err
}

func getFields(s interface{}) fields {
	r := reflect.ValueOf(s)
	numfield := r.Elem().NumField()

	if r.Kind() != reflect.Ptr {
		log.Fatal("Wrong type struct")
	}

	var result fields
	for i := 0; i < numfield; i++ {
		f := reflect.TypeOf(s).Elem().Field(i)
		result = append(result,
			field{
				name:  f.Name,
				value: r.Elem().Field(i).Interface(),
				tag:   f.Tag.Get("validate"),
			})
	}

	return result
}
