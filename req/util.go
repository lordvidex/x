// util.go contains the configuration for the request utility functions.

// Package req is the helper package for handling requests.
//
// It contains functions that help with parsing requests into structs, validating them, and returning errors.
package req

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

// I - the default instance of the req package for convenient usage.
// If custom registration is required for the Util components, it is recommended to create
// a new instance of the Util struct rather than registering on I directly.
// This is in order to prevent unexpected behavior when two or more modules use the req package.
var I = NewUtil()

type Util struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

// NewUtil creates a new instance of the Util struct with the default validator and an English translator.
func NewUtil() *Util {
	v := validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	entranslations.RegisterDefaultTranslations(v, trans)
	return &Util{Validate: v, Translator: trans}
}

func (u *Util) Will() *Will {
	return &Will{util: u}
}
