package ptrs

func Ptr[T any](value T) *T {
	return &value
}
