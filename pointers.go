package ptrutils

// FromPointer returns the value held at the pointer address.
// If the value is nil, the zero value for the type is returned.
func FromPointer[T any](t *T) T {
	if t == nil {
		return *new(T)
	}
	return *t
}

// Return a pointer to the given value.
func ToPointer[T any](t T) *T {
	return &t
}
