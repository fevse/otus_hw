package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res strings.Builder
	sr := []rune(s)
	switch {
	case len(s) == 0:
		return "", nil
	case len(s) == 1:
		return string(sr[0]), nil
	case unicode.IsDigit(sr[0]):
		return "", ErrInvalidString
	}

	var pr rune

	for _, v := range sr {
		switch {
		case unicode.IsDigit(pr) && unicode.IsDigit(v):
			return "", ErrInvalidString
		case unicode.IsDigit(v) && !unicode.IsDigit(pr):
			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Printf("Atoi error: %v\n", err)
				return "", ErrInvalidString
			}
			c := string(pr)
			res.WriteString(strings.Repeat(c, n))
		case !unicode.IsDigit(v) && !unicode.IsDigit(pr) && pr != 0:
			res.WriteRune(pr)
		}
		pr = v
	}
	if !unicode.IsDigit(sr[len(sr)-1]) {
		res.WriteRune(sr[len(sr)-1])
	}
	return res.String(), nil
}
