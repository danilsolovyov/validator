package validator

import (
	"reflect"
	"strings"
	"regexp"
	"log"
	"strconv"
)

type Field struct {
	Name  string
	Value interface{}
	Tag   string
}

type Fields []Field

var separator = ";"

// Функции: Validate, createErrors

func Validate(s interface{}) error{
	var err error
	fields := getFields(s)

	for _, field := range fields {
		field.Name = strings.ToLower(field.Name)

		switch field.Value.(type) {
		case int:
			err = validateInt(field)
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

func validateInt(field Field) error {
	more, err := regexp.Compile("^>\\d+")
	less, err := regexp.Compile("^<\\d+")
	equal, err := regexp.Compile("^=\\d+")
	notEqual, err := regexp.Compile("^!=\\d+")
	
	for _, validator := range strings.Split(field.Tag, separator) {

		if more.MatchString(validator) {
			s := strings.Replace(validator, ">", "", -1)
			x, _ := strconv.Atoi(s)

			if field.Value.(int) <= x {
				return ErrMustBeMore(field.Name, x)
			}
		}

		if less.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.Atoi(s)

			if field.Value.(int) >= x {
				return ErrMustBeLess(field.Name, x)
			}
		}

		if equal.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.Atoi(s)

			if field.Value.(int) != x {
				return ErrMustBeEqual(field.Name, x)
			}
		}

		if notEqual.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.Atoi(s)

			if field.Value.(int) == x {
				return ErrMustNotBeEqual(field.Name, x)
			}
		}
	}

	return err
}

func validateFloat(field Field) error {
	more, err := regexp.Compile("^>\\d+.\\d+")
	less, err := regexp.Compile("^<\\d+.\\d+")
	equal, err := regexp.Compile("^=\\d+.\\d+")
	notEqual, err := regexp.Compile("^!=\\d+.\\d+")
	
	for _, validator := range strings.Split(field.Tag, separator) {

		if more.MatchString(validator) {
			s := strings.Replace(validator, ">", "", -1)
			x, _ := strconv.ParseFloat(s, 64)

			if field.Value.(float64) <= x {
				return ErrMustBeMore(field.Name, x)
			}
		}

		if less.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseFloat(s, 64)

			if field.Value.(float64) >= x {
				return ErrMustBeLess(field.Name, x)
			}
		}

		if equal.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseFloat(s, 64)

			if field.Value.(float64) != x {
				return ErrMustBeEqual(field.Name, x)
			}
		}

		if notEqual.MatchString(validator) {
			s := strings.Replace(validator, "<", "", -1)
			x, _ := strconv.ParseFloat(s, 64)

			if field.Value.(float64) == x {
				return ErrMustNotBeEqual(field.Name, x)
			}
		}
	}

	return err
}

func validateString(field Field) error {
	if strings.Contains(field.Tag, "required") && field.Value.(string) == "" {
		return ErrRequired(field.Name)
	}

	lenMore, err := regexp.Compile("^len>\\d+")
	lenLess, err := regexp.Compile("^len<\\d+")
	lenEqual, err := regexp.Compile("^len=\\d+")
	formatEqual, err := regexp.Compile("^format=.+")

	if err != nil {
		return err
	}

	for _, validator := range strings.Split(field.Tag, separator) {
		runeVal := []rune(field.Value.(string))
		
		if lenMore.MatchString(validator) {
			s := strings.Replace(validator, "len>", "", -1)
			x, _ := strconv.Atoi(s)

			if len(runeVal) <= x {
				return ErrLenTooShort(field.Name)
			}
		}

		if lenLess.MatchString(validator) {
			s := strings.Replace(validator, "len<", "", -1)
			x, _ := strconv.Atoi(s)

			if len(runeVal) >= x {
				return ErrLenTooLong(field.Name)
			}
		}

		if lenEqual.MatchString(validator) {
			s := strings.Replace(validator, "len=", "", -1)
			x, _ := strconv.Atoi(s)

			if len(runeVal) != x {
				return ErrLenMustBe(field.Name, x)
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

			if !formatRegex.MatchString(field.Value.(string)) && len(runeVal) > 0 {
				return ErrInvalidFormat(field.Name)
			}
		}
	}


	return err
}

func getFields(s interface{}) Fields {
	r := reflect.ValueOf(s)
	numfield := r.Elem().NumField()

	if r.Kind() != reflect.Ptr {
		log.Fatal("Wrong type struct")
	}
	
	var result Fields
	for i := 0; i < numfield; i++ {
		field := reflect.TypeOf(s).Elem().Field(i)
		result = append(result,
			Field {
				Name:  field.Name,
				Value: r.Elem().Field(i).Interface(),
				Tag:   field.Tag.Get("validate"), 
		})
	}

	return result
}