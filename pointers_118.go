//go:build go1.18
// +build go1.18

package pointers

// New returns a pointer to any value x.
//
//	p := pointers.New(13)
//
// is like a convenient one-liner equivalent of the following 2 lines
// of code:
//
//	i := 13
//	p := &i
//
// The added value of New is merely to provide a simple helper to do the
// above operations all in one assignment. If your program needs this a
// lot, it could simplify and make your code more readable. For instance,
// assigning many fields within a struct initialization that are pointers
// to primitives.
//
// When x is zero value, New still returns a pointer to x by default.
// To return nil instead, there are 2 ways to achieve this:
//
// Method 1:
// Pass an extra bool argument true or pointers.Nilify.
// Example:
//
//	_ = pointers.New(0) // => new(int)
//	_ = pointers.New(0, pointers.Nilify)  // => nil
//
// Method 2:
// Overwrite AlwaysNilify value to true or pointers.Nilify before using New.
// Example:
//
//	_ = pointers.New(0)          // => new(int)
//	_ = pointers.New("")         // => new(string)
//	_ = pointers.New(float32(0)) // => new(float32)
//	pointers.AlwaysNilify = true
//	_ = pointers.New(0)          // => nil
//	_ = pointers.New("")         // => nil
//	_ = pointers.New(float32(0)) // => nil
func New[T comparable](x T, opts ...bool) *T {
	nilify := AlwaysNilify
	if len(opts) > 0 {
		nilify = opts[0]
	}
	if nilify && x == *new(T) {
		return nil
	}
	return &x
}
