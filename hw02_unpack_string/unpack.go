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
	case len(sr) == 0:
		return "", nil
	case len(sr) == 1:
		return string(sr[0]), nil
	case unicode.IsDigit(sr[0]):
		return "", ErrInvalidString
	}
	if !unicode.IsDigit(sr[0]) && !unicode.IsDigit(sr[1]) {
		fmt.Fprint(&res, string(sr[0]))
	}
	i := 1
	for i < len(sr)-1 {
		switch {
		case !unicode.IsDigit(sr[i]) && !unicode.IsDigit(sr[i+1]):
			fmt.Fprint(&res, string(sr[i]))
		case unicode.IsDigit(sr[i]) && !unicode.IsDigit(sr[i-1]) && !unicode.IsDigit(sr[i+1]):
			n, _ := strconv.Atoi(string(sr[i]))
			c := string(sr[i-1])
			fmt.Fprint(&res, strings.Repeat(c, n))
		case unicode.IsDigit(sr[i]) && unicode.IsDigit(sr[i+1]):
			return "", ErrInvalidString
		}
		i++
	}
	fmt.Fprint(&res, string(sr[len(sr)-1]))
	return res.String(), nil
}
