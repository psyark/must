package must

import "fmt"

func Try[T any](fn func() T) (t T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return fn(), nil
}
