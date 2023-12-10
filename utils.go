package pointers

import (
	"fmt"
	"time"
)

const Nilify bool = true   // New sets to nil when input is zero value.
var AlwaysNilify = !Nilify // Controls default behaviour for func New.

// nilify returns true input x is zero value and one the following is
// also true: opts[0] or AlwaysNilify. Else, returns false.
func nilify(x interface{}, opts ...bool) bool {
	nilify := AlwaysNilify
	if len(opts) > 0 {
		nilify = opts[0]
	}
	return nilify && isZero(x)
}

// isZero returns true if the underlying value of x is
// the zero-value of its type.
func isZero(x interface{}) bool {
	switch t := x.(type) {
	case bool:
		return !t
	case int:
		return t == 0
	case int8:
		return t == 0
	case int16:
		return t == 0
	case int32: // <=> rune
		return t == 0
	case int64:
		return t == 0
	case uint:
		return t == 0
	case uint8: // <=> byte
		return t == 0
	case uint16:
		return t == 0
	case uint32:
		return t == 0
	case uint64:
		return t == 0
	case uintptr:
		return t == 0
	case float32:
		return t == 0
	case float64:
		return t == 0
	case complex64:
		return real(t) == 0 && imag(t) == 0
	case complex128:
		return real(t) == 0 && imag(t) == 0
	case string:
		return t == ""
	case time.Time:
		return t.IsZero()
	case time.Duration:
		return t == 0
	case []byte:
		return len(t) == 00
	default:
		return x == nil
	}
}

// testName describe the current test.
func testName(f string, x interface{}) string {
	return fmt.Sprintf("Test %s(x %T = %v)", f, x, x)
}
