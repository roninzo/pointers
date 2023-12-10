package pointers

import (
	"reflect"
	"time"
)

// Any returns a pointer to interface{} value x.
func Any(x interface{}, opts ...bool) interface{} {
	if nilify(x, opts...) {
		return nil
	}
	r := reflect.New(reflect.TypeOf(x))
	reflect.ValueOf(r.Interface()).Elem().Set(reflect.ValueOf(x))
	return r.Interface()
}

// Bool returns a pointer to bool value x.
func Bool(x bool, opts ...bool) *bool {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Byte returns a pointer to byte value x.
func Byte(x byte, opts ...bool) *byte {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Bytes returns a pointer to []byte value x.
func Bytes(x []byte, opts ...bool) *[]byte {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Int returns a pointer to int value x.
func Int(x int, opts ...bool) *int {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Int8 returns a pointer to int8 value x.
func Int8(x int8, opts ...bool) *int8 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Int16 returns a pointer to int16 value x.
func Int16(x int16, opts ...bool) *int16 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Int32 returns a pointer to int32 value x.
func Int32(x int32, opts ...bool) *int32 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Int64 returns a pointer to int64 value x.
func Int64(x int64, opts ...bool) *int64 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Uint returns a pointer to uint value x.
func Uint(x uint, opts ...bool) *uint {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Uint8 returns a pointer to uint8 value x.
func Uint8(x uint8, opts ...bool) *uint8 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Uint16 returns a pointer to uint16 value x.
func Uint16(x uint16, opts ...bool) *uint16 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Uint32 returns a pointer to uint32 value x.
func Uint32(x uint32, opts ...bool) *uint32 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Uint64 returns a pointer to uint64 value x.
func Uint64(x uint64, opts ...bool) *uint64 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Uintptr returns a pointer to uintptr value x.
func Uintptr(x uintptr, opts ...bool) *uintptr {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Float32 returns a pointer to float32 value x.
func Float32(x float32, opts ...bool) *float32 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Float64 returns a pointer to float64 value x.
func Float64(x float64, opts ...bool) *float64 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Complex64 returns a pointer to complex64 value x.
func Complex64(x complex64, opts ...bool) *complex64 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Complex128 returns a pointer to complex128 value x.
func Complex128(x complex128, opts ...bool) *complex128 {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// String returns a pointer to string value x.
func String(x string, opts ...bool) *string {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Rune returns a pointer to rune value x.
func Rune(x rune, opts ...bool) *rune {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Time returns a pointer to type value x.
func Time(x time.Time, opts ...bool) *time.Time {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}

// Duration returns a pointer to time.Duration value x.
func Duration(x time.Duration, opts ...bool) *time.Duration {
	if nilify(x, opts...) {
		return nil
	}
	return &x
}
