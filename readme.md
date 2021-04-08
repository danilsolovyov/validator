# Validator

## How to install

`go get github.com/danilsolovyov/validator`
## How to use

```golang
package main

import (
	"encoding/json"
	"net/http"
	"validator"
)

type RequestUser struct {
	ID    int    `validate:">0"`
	Name  string `validate:"len>0;len<30"`
	Email string `validate:"format=email;required"`
}

func main() {
	// ...
	userValidator := validator.NewValidator(&RequestUser{})
	mux := http.NewServeMux()
	mux.Handle("/user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			UserHandler(w, r, userValidator)
	}))
	// ...
	data := RequestUser{
		ID:    3,
		Name:  "Mr. Somebody",
		Email: "somebody.example.com",
	}
	UserExample(userValidator, data)
}

// Example 1
func UserHandler(w http.ResponseWriter, r *http.Request, v validator.Validator) {
	var request RequestUser
	decoder := validator.Decoder{json.NewDecoder(r.Body)}
	err := decoder.DecodeAndValidate(&request, v)
	if err != nil {
		// ...
    }
}

// Example 2
func UserExample (v validator.Validator, data RequestUser) {
	v.AddValues(data)
	err := v.Validate()
	if err != nil {
		// ...
    }
	// ...
}
```

## Validation expressions
Separator for expressions: `;`
- Numeric (int, float64) `>, <, =, !=`
- String
    - `required`
    - Length `len` `>, <, =`
    - Formats `format=` `email, phone, ipv4, ipv6, date, time24, datetime`

