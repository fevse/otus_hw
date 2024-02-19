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
						switch s[0] {
						case "len":
							l, err := strconv.Atoi(s[1])
							if err != nil {
								return ErrorStrconvAtoi
							}
							if len(val) != l {
								ve = append(ve, ValidationError{field.Name, ErrorStringLen})
							}
						case "in":
							in := strings.Split(s[1], ",")
							ok := false
							for _, j := range in {
								if j == val {
									ok = true
								}
							}
							if !ok {
								ve = append(ve, ValidationError{field.Name, ErrorStringIn})
							}
						case "regexp":
							re, err := regexp.Compile(s[1])
							if err != nil {
								return ErrorRegexp
							}
							if !re.MatchString(val) {
								ve = append(ve, ValidationError{field.Name, ErrorStringRegexp})
							}
						}
					}
				case "string":
					val := value.String()
					switch s[0] {
					case "len":
						l, err := strconv.Atoi(s[1])
						if err != nil {
							return ErrorStrconvAtoi
						}
						if len(val) != l {
							ve = append(ve, ValidationError{field.Name, ErrorStringLen})
						}
					case "in":
						in := strings.Split(s[1], ",")
						ok := false
						for _, j := range in {
							if j == val {
								ok = true
							}
						}
						if !ok {
							ve = append(ve, ValidationError{field.Name, ErrorStringIn})
						}
					case "regexp":
						re, err := regexp.Compile(s[1])
						if err != nil {
							return ErrorRegexp
						}
						if !re.MatchString(val) {
							ve = append(ve, ValidationError{field.Name, ErrorStringRegexp})
						}
					}
				case "[]int":
					for j := 0; j < value.Len(); j++ {
						val := value.Int()
						switch s[0] {
						case "in":
							in := strings.Split(s[1], ",")
							ok := false
							for _, j := range in {
								if j == strconv.Itoa(int(val)) {
									ok = true
								}
							}
							if !ok {
								ve = append(ve, ValidationError{field.Name, ErrorIntIn})
							}
						case "max":
							max, err := strconv.Atoi(s[1])
							if err != nil {
								fmt.Println(s[1])
								return ErrorStrconvAtoi
							}
							if val > int64(max) {
								ve = append(ve, ValidationError{field.Name, ErrorIntMax})
							}
						case "min":
							min, err := strconv.Atoi(s[1])
							if err != nil {
								fmt.Println(s[1])
								return ErrorStrconvAtoi
							}
							if val < int64(min) {
								ve = append(ve, ValidationError{field.Name, ErrorIntMin})
							}
						}
					}
				case "int":
					val := value.Int()
					switch s[0] {
					case "in":
						in := strings.Split(s[1], ",")
						ok := false
						for _, j := range in {
							if j == strconv.Itoa(int(val)) {
								ok = true
							}
						}
						if !ok {
							ve = append(ve, ValidationError{field.Name, ErrorIntIn})
						}
					case "max":
						max, err := strconv.Atoi(s[1])
						if err != nil {
							fmt.Println(s[1])
							return ErrorStrconvAtoi
						}
						if val > int64(max) {
							ve = append(ve, ValidationError{field.Name, ErrorIntMax})
						}
					case "min":
						min, err := strconv.Atoi(s[1])
						if err != nil {
							fmt.Println(s[1])
							return ErrorStrconvAtoi
						}
						if val < int64(min) {
							ve = append(ve, ValidationError{field.Name, ErrorIntMin})
						}
					}
				}
			}
		}
	}

	if len(ve) > 0 {
		return ve
	}
	return nil
}
