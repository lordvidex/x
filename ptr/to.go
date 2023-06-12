package ptr

// ToObj returns the value of the pointer x, or a zero value if x is nil.
func ToObj[T any](x *T) T {
	if x == nil {
		return *new(T)
	}
	return *x
}

// ToString converts a pointer to a string to a string.
func ToString(x *string) string {
	if x == nil {
		return ""
	}
	return *x
}
