package must

// Deprecated: Use github.com/samber/lo instead.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
