package pointers

import "fmt"

// New returns a pointer to any value x.
func New[T comparable](x T) *T {
	if fmt.Sprintf("%T", x) != "bool" {
		if x == *new(T) {
			return nil
		}
	}
	return &x
}
