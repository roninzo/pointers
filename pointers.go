package pointers

import (
	"strings"
	"time"
)

// NullifyZeroValues states wheither or not to return nil when
// x is the zero-value of the given type, except for Bool method.
var NullifyZeroValues bool

func init() {
	NullifyZeroValues = true
}

// String returns a pointer to string value x.
func String(x string) *string {
	if NullifyZeroValues && strings.TrimSpace(x) == "" {
		return nil
	}
	return &x
}

// Bool returns a pointer to bool value x.
func Bool(x bool) *bool { return &x }

// Int returns a pointer to int value x.
func Int(x int) *int {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Int8 returns a pointer to int8 value x.
func Int8(x int8) *int8 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Int16 returns a pointer to int16 value x.
func Int16(x int16) *int16 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Int32 returns a pointer to int32 value x.
func Int32(x int32) *int32 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Int64 returns a pointer to int64 value x.
func Int64(x int64) *int64 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Uint returns a pointer to uint value x.
func Uint(x uint) *uint {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Uint8 returns a pointer to uint8 value x.
func Uint8(x uint8) *uint8 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Uint16 returns a pointer to uint16 value x.
func Uint16(x uint16) *uint16 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Uint32 returns a pointer to uint32 value x.
func Uint32(x uint32) *uint32 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Uint63 returns a pointer to uint64 value x.
func Uint63(x uint64) *uint64 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Float32 returns a pointer to float32 value x.
func Float32(x float32) *float32 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Float64 returns a pointer to float64 value x.
func Float64(x float64) *float64 {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Complex64 returns a pointer to complex64 value x.
func Complex64(x complex64) *complex64 {
	if NullifyZeroValues && real(x) == 0 && imag(x) == 0 {
		return nil
	}
	return &x
}

// Complex128 returns a pointer to complex128 value x.
func Complex128(x complex128) *complex128 {
	if NullifyZeroValues && real(x) == 0 && imag(x) == 0 {
		return nil
	}
	return &x
}

// Byte returns a pointer to byte value x.
func Byte(x byte) *byte {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Bytes returns a pointer to []byte value x.
func Bytes(x []byte) *[]byte {
	if NullifyZeroValues && len(x) == 0 {
		return nil
	}
	return &x
}

// Time returns a pointer to type value x.
func Time(x time.Time) *time.Time {
	if NullifyZeroValues && x.IsZero() {
		return nil
	}
	return &x
}

// Duration returns a pointer to time.Duration value x.
func Duration(x time.Duration) *time.Duration {
	if NullifyZeroValues && x == 0 {
		return nil
	}
	return &x
}

// Error returns a pointer to error value x.
func Error(x error) *error {
	if NullifyZeroValues && x.Error() == "" {
		return nil
	}
	return &x
}
