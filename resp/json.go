// Package resp is the helper package for handling writes to json response
package resp

import (
	"encoding/json"
	"net/http"

	"github.com/lordvidex/errs/v2"
)

type ErrorRes struct {
	Message []string `json:"message"`
	Error   string   `json:"error"`
	Details []string `json:"details,omitempty"`
}

// JSON returns a JSON response with v.
func JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// JSONC returns a JSON response with v and a custom status code.
func JSONC(w http.ResponseWriter, code int, v interface{}) {
	w.WriteHeader(code)
	JSON(w, v)
}

// Error returns a JSON response with schema ErrorRes.
// If error type is unknown or does not conform to *errs.Error, http.InternalServerError is returned by default.
func Error(w http.ResponseWriter, err error) {
	var r ErrorRes
	code := http.StatusInternalServerError
	switch x := err.(type) {
	case *errs.Error:
		r.Message = x.Msg
		r.Error = x.Code.String()
		code = x.Code.HTTP()
	default:
		r.Message = append(r.Message, err.Error())
		r.Error = errs.Unknown.String()
	}
	JSONC(w, code, r)
}
