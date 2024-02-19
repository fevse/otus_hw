package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			// Place your code here.
			Validate(100),
			ErrorUnsupportedType,
		},
		{
			Validate(User{
				ID:    "12345678901234567890123456789012345", // 35 - validation error
				Name:  "DefaultUser",
				Age:   55,              // validation error
				Email: "unknown email", // validation error
				Role:  "user",
				Phones: []string{
					"2128505",
					"2128506",
					"2128507",
				}, // validation error
			}),
			ValidationErrors{
				ValidationError{"ID", ErrorStringLen},
				ValidationError{"Age", ErrorIntMax},
				ValidationError{"Email", ErrorStringRegexp},
				ValidationError{"Phones", ErrorStringLen},
				ValidationError{"Phones", ErrorStringLen},
				ValidationError{"Phones", ErrorStringLen},
			},
		},
		{
			Validate(Response{
				Code: 200,
				Body: "hello",
			}),
			nil,
		},
		{
			Validate(Response{
				Code: 101,
				Body: "oh, no",
			}),
			ValidationErrors{ValidationError{"Code", ErrorIntIn}},
		},
		{
			Validate(App{
				Version: "hello",
			}),
			nil,
		},
		{
			Validate(App{
				Version: "goodbye",
			}),
			ValidationErrors{ValidationError{"Version", ErrorStringLen}},
		},
		// ...
		// Place your code here.
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			// Place your code here.
			e := tt.in
			if e != nil {
				require.Equal(t, e, tt.expectedErr)
			}
			_ = tt
		})
	}
}
