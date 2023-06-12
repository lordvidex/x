package req

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/lordvidex/errs"

	"github.com/lordvidex/x/auth"
	"github.com/lordvidex/x/ptr"

	"github.com/stretchr/testify/assert"
)

func ExampleWill_Err() {
	// struct definition with validation
	type t struct {
		X    int    `json:"x" validate:"required,gt=0"`
		Name string `json:"name" validate:"required"`
	}
	// json request body
	json := `{"x":1,"name":"hello"}`

	// mock request
	r, _ := http.NewRequestWithContext(context.TODO(), "POST", "http://localhost:8080", strings.NewReader(json))

	// dto
	v := &t{}

	// chained requests
	err := I.Will().Bind(r, v).Validate(v).Err()

	// expect
	fmt.Println(err, *v)
	// Output: <nil> {1 hello}
}

func TestWill_Token(t *testing.T) {
	testcases := []struct {
		name string
		ctx  context.Context
		want auth.Token
	}{
		{
			name: "get token from context",
			ctx:  auth.WithToken(context.Background(), auth.Token("hidden_token")),
			want: auth.Token("hidden_token"),
		},
		{
			name: "context without token",
			ctx:  context.Background(),
			want: auth.Token(""),
		},
		{
			name: "nested context with token",
			ctx:  context.WithValue(auth.WithToken(context.Background(), auth.Token("hidden_token")), ptr.String("key"), "value"),
			want: auth.Token("hidden_token"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var got auth.Token
			I.Will().Token(tc.ctx, &got)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWill_Validate(t *testing.T) {
	testcases := []struct {
		name   string
		v      interface{}
		expect error
	}{
		{
			name: "valid",
			v: struct {
				X    int    `validate:"required,gt=0"`
				Name string `validate:"required"`
			}{
				X:    1,
				Name: "hello",
			},
			expect: nil,
		},
		{
			name: "one missing field",
			v: struct {
				X int `validate:"required"`
			}{},
			expect: errs.B().Code(errs.InvalidArgument).Msg("X is a required field").Err(),
		},
		{
			name: "multiple errors from a single field returns first", // it would have been nice to receive all the validation errors from a single field though, but one is returned by the package used.
			v: struct {
				X int `validate:"required,gt=0"`
			}{},
			expect: errs.B().Code(errs.InvalidArgument).Msg("X is a required field").Err(),
		},
		{
			name: "multiple field errors returns first of each field",
			v: struct {
				X    int    `validate:"required,gt=0"`
				Name string `validate:"required"`
			}{
				X:    0,
				Name: "",
			},
			expect: errs.B().Code(errs.InvalidArgument).Msg("X is a required field").Msg("Name is a required field").Err(),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := I.Will().Validate(tc.v).Err()
			assert.True(t, errors.Is(err, tc.expect), "expected %v, got %v", tc.expect, err)
		})
	}
}

func TestWill_AbandonChainWhenErr(t *testing.T) {
	var token auth.Token
	v := struct {
		X int `validate:"required"`
	}{X: 0}
	err := I.Will().Validate(v).Token(context.Background(), &token).Err()
	assert.Error(t, err)
	assert.Equal(t, auth.Token(""), token, "token should not be reached")
}
