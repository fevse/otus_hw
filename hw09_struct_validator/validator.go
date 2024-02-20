package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrorUnsupportedType = errors.New("unsupported type")
	ErrorRegexp          = errors.New("regexp error")
	ErrorStrconvAtoi     = errors.New("strconvAtoi error")
)

var (
	ErrorIntIn        = errors.New("int in err")
	ErrorIntMax       = errors.New("int max err")
	ErrorIntMin       = errors.New("int min err")
	ErrorStringLen    = errors.New("string len err")
	ErrorStringIn     = errors.New("string in err")
	ErrorStringRegexp = errors.New("string regexp err")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	err := ""
	for _, e := range v {
		err += fmt.Sprintf("Field: %v, error: %v", e.Field, e.Err)
	}
	return err
}

func Validate(v interface{}) error {
	vv := reflect.TypeOf(v)
	if vv.Kind().String() != "struct" {
		return ErrorUnsupportedType
	}
	p := reflect.ValueOf(v)
	var ve ValidationErrors
	for i := 0; i < vv.NumField(); i++ {
		field := vv.Field(i)
		value := p.Field(i)
		if validate, ok := field.Tag.Lookup("validate"); ok {
			x := strings.Split(validate, "|")
			for i := range x {
				s := strings.Split(x[i], ":")
				switch field.Type.String() {
				case "[]string":
					for j := 0; j < value.Len(); j++ {
						val := value.Index(j).String()

						x, err := stringVal(s[0], s[1], val, field.Name)
						if err != nil {
							return err
						}
						ve = append(ve, x...)
					}
				case "string":
					val := value.String()

					x, err := stringVal(s[0], s[1], val, field.Name)
					if err != nil {
						return err
					}
					ve = append(ve, x...)
				case "[]int":
					for j := 0; j < value.Len(); j++ {
						val := value.Int()
						x, err := intVal(s[0], s[1], field.Name, val)
						if err != nil {
							return err
						}
						ve = append(ve, x...)
					}
				case "int":
					val := value.Int()

					x, err := intVal(s[0], s[1], field.Name, val)
					if err != nil {
						return err
					}
					ve = append(ve, x...)
				}
			}
		}
	}
	if len(ve) > 0 {
		return ve
	}
	return nil
}

func stringVal(t, s, val, field string) (ValidationErrors, error) {
	var ve ValidationErrors
	switch t {
	case "len":
		l, err := strconv.Atoi(s)
		if err != nil {
			return nil, ErrorStrconvAtoi
		}
		if len(val) != l {
			ve = append(ve, ValidationError{field, ErrorStringLen})
		}
	case "in":
		in := strings.Split(s, ",")
		ok := false
		for _, j := range in {
			if j == val {
				ok = true
			}
		}
		if !ok {
			ve = append(ve, ValidationError{field, ErrorStringIn})
		}
	case "regexp":
		re, err := regexp.Compile(s)
		if err != nil {
			return nil, ErrorRegexp
		}
		if !re.MatchString(val) {
			ve = append(ve, ValidationError{field, ErrorStringRegexp})
		}
	}
	return ve, nil
}

func intVal(t, s, field string, val int64) (ValidationErrors, error) {
	var ve ValidationErrors
	switch t {
	case "in":
		in := strings.Split(s, ",")
		ok := false
		for _, j := range in {
			if j == strconv.Itoa(int(val)) {
				ok = true
			}
		}
		if !ok {
			ve = append(ve, ValidationError{field, ErrorIntIn})
		}
	case "max":
		max, err := strconv.Atoi(s)
		if err != nil {
			return nil, ErrorStrconvAtoi
		}
		if val > int64(max) {
			ve = append(ve, ValidationError{field, ErrorIntMax})
		}
	case "min":
		min, err := strconv.Atoi(s)
		if err != nil {
			return nil, ErrorStrconvAtoi
		}
		if val < int64(min) {
			ve = append(ve, ValidationError{field, ErrorIntMin})
		}
	}
	return ve, nil
}
