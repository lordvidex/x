// Package ptr provides functions for conveniently converting primitive types
// to their pointer equivalents and vice versa.
package ptr

// String returns a ptr from a string value
func String(x string) *string {
	return &x
}

// Float64 returns a ptr from a float64 value
func Float64(x float64) *float64 {
	return &x
}

// Bool returns a ptr from a bool value
func Bool(x bool) *bool {
	return &x
}

// Obj returns a ptr of any type T from its object value
func Obj[T any](x T) *T {
	return &x
}
