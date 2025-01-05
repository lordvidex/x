// will.go file contains the actual request utility functions.

package req

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/lordvidex/errs/v2"

	"github.com/lordvidex/x/auth"
)

// Will is a struct that is used to chain together a series of actions to be performed on the request.
// Once an error occurs, the chain is broken and no further action is performed.
// error can be retrieved by calling the Err() method.
type Will struct {
	err  error
	util *Util
}

// Err returns the error that occurred during the chain of actions if any.
func (w *Will) Err() error {
	return w.err
}

// Bind will bind the request body to the given interface.
func (w *Will) Bind(r *http.Request, v interface{}) *Will {
	if w.err != nil {
		return w
	}
	w.err = opt(json.NewDecoder(r.Body).Decode(v), "error parsing request body to json")
	return w
}

// Validate will validate the given interface with the given util validator.
func (w *Will) Validate(v interface{}) *Will {
	if w.err != nil {
		return w
	}
	err := w.util.Validate.Struct(v)
	if err != nil {
		switch err := err.(type) {
		case validator.ValidationErrors:
			msgs := make([]string, len(err))
			for i, ferr := range err {
				msgs[i] = ferr.Translate(w.util.Translator)
			}
			w.err = errs.B().Code(errs.InvalidArgument).Msg(msgs...).Err()
		default:
			w.err = errs.WrapCode(err, errs.Internal, "unknown validation error")
		}
	}
	return w
}

func (w *Will) Token(ctx context.Context, token *auth.Token) *Will {
	if w.err != nil {
		return w
	}
	t, _ := auth.GetToken(ctx)
	*token = t
	return w
}

func opt(err error, msg string) error {
	if err == nil {
		return nil
	}
	return errs.WrapCode(err, errs.InvalidArgument, msg)
}
