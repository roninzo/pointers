package pointers

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	f := "Any"
	assert.Nil(t, Any(false, Nilify), testName(f, false))
	assert.Equal(t, false, *Any(false).(*bool), testName(f, false))
	assert.Equal(t, true, *Any(true).(*bool), testName(f, true))
	assert.Nil(t, Any(byte(0), Nilify), testName(f, byte(0)))
	assert.Equal(t, byte(0), *Any(byte(0)).(*byte), testName(f, byte(0)))
	assert.Equal(t, byte(10), *Any(byte(10)).(*byte), testName(f, byte(10)))
	assert.Nil(t, Any([]byte{}, Nilify), testName(f, []byte{}))
	assert.Equal(t, []byte{}, *Any([]byte{}).(*[]byte), testName(f, []byte{}))
	assert.Equal(t, []byte{10, 55}, *Any([]byte{10, 55}).(*[]byte), testName(f, []byte{10, 55}))
	assert.Nil(t, Any(int(0), Nilify), testName(f, int(0)))
	assert.Equal(t, int(-7), *Any(int(-7)).(*int), testName(f, int(-7)))
	assert.Equal(t, int(0), *Any(int(0)).(*int), testName(f, int(0)))
	assert.Equal(t, int(10), *Any(int(10)).(*int), testName(f, int(10)))
	assert.Nil(t, Any(int8(0), Nilify), testName(f, int8(0)))
	assert.Equal(t, int8(-7), *Any(int8(-7)).(*int8), testName(f, int8(-7)))
	assert.Equal(t, int8(0), *Any(int8(0)).(*int8), testName(f, int8(0)))
	assert.Equal(t, int8(10), *Any(int8(10)).(*int8), testName(f, int8(10)))
	assert.Nil(t, Any(int16(0), Nilify), testName(f, int16(0)))
	assert.Equal(t, int16(-7), *Any(int16(-7)).(*int16), testName(f, int16(-7)))
	assert.Equal(t, int16(0), *Any(int16(0)).(*int16), testName(f, int16(0)))
	assert.Equal(t, int16(10), *Any(int16(10)).(*int16), testName(f, int16(10)))
	assert.Nil(t, Any(int32(0), Nilify), testName(f, int32(0)))
	assert.Equal(t, int32(-7), *Any(int32(-7)).(*int32), testName(f, int32(-7)))
	assert.Equal(t, int32(0), *Any(int32(0)).(*int32), testName(f, int32(0)))
	assert.Equal(t, int32(10), *Any(int32(10)).(*int32), testName(f, int32(10)))
	assert.Nil(t, Any(int64(0), Nilify), testName(f, int64(0)))
	assert.Equal(t, int64(-7), *Any(int64(-7)).(*int64), testName(f, int64(-7)))
	assert.Equal(t, int64(0), *Any(int64(0)).(*int64), testName(f, int64(0)))
	assert.Equal(t, int64(10), *Any(int64(10)).(*int64), testName(f, int64(10)))
	assert.Nil(t, Any(uint(0), Nilify), testName(f, uint(0)))
	assert.Equal(t, uint(0), *Any(uint(0)).(*uint), testName(f, uint(0)))
	assert.Equal(t, uint(10), *Any(uint(10)).(*uint), testName(f, uint(10)))
	assert.Nil(t, Any(uint8(0), Nilify), testName(f, uint8(0)))
	assert.Equal(t, uint8(0), *Any(uint8(0)).(*uint8), testName(f, uint8(0)))
	assert.Equal(t, uint8(10), *Any(uint8(10)).(*uint8), testName(f, uint8(10)))
	assert.Nil(t, Any(uint16(0), Nilify), testName(f, uint16(0)))
	assert.Equal(t, uint16(0), *Any(uint16(0)).(*uint16), testName(f, uint16(0)))
	assert.Equal(t, uint16(10), *Any(uint16(10)).(*uint16), testName(f, uint16(10)))
	assert.Nil(t, Any(uint32(0), Nilify), testName(f, uint32(0)))
	assert.Equal(t, uint32(0), *Any(uint32(0)).(*uint32), testName(f, uint32(0)))
	assert.Equal(t, uint32(10), *Any(uint32(10)).(*uint32), testName(f, uint32(10)))
	assert.Nil(t, Any(uint64(0), Nilify), testName(f, uint64(0)))
	assert.Equal(t, uint64(0), *Any(uint64(0)).(*uint64), testName(f, uint64(0)))
	assert.Equal(t, uint64(10), *Any(uint64(10)).(*uint64), testName(f, uint64(10)))
	assert.Nil(t, Any(uintptr(0), Nilify), testName(f, uintptr(0)))
	assert.Equal(t, uintptr(0), *Any(uintptr(0)).(*uintptr), testName(f, uintptr(0)))
	assert.Equal(t, uintptr(10), *Any(uintptr(10)).(*uintptr), testName(f, uintptr(10)))
	assert.Nil(t, Any(float32(0), Nilify), testName(f, float32(0)))
	assert.Equal(t, float32(-9.50), *Any(float32(-9.50)).(*float32), testName(f, float32(-9.50)))
	assert.Equal(t, float32(-1), *Any(float32(-1)).(*float32), testName(f, float32(-1)))
	assert.Equal(t, float32(0), *Any(float32(0)).(*float32), testName(f, float32(0)))
	assert.Equal(t, float32(4), *Any(float32(4)).(*float32), testName(f, float32(4)))
	assert.Equal(t, float32(6.05), *Any(float32(6.05)).(*float32), testName(f, float32(6.05)))
	assert.Nil(t, Any(float64(0), Nilify), testName(f, float64(0)))
	assert.Equal(t, float64(-9.50), *Any(float64(-9.50)).(*float64), testName(f, float64(-9.50)))
	assert.Equal(t, float64(-1), *Any(float64(-1)).(*float64), testName(f, float64(-1)))
	assert.Equal(t, float64(0), *Any(float64(0)).(*float64), testName(f, float64(0)))
	assert.Equal(t, float64(4), *Any(float64(4)).(*float64), testName(f, float64(4)))
	assert.Equal(t, float64(6.05), *Any(float64(6.05)).(*float64), testName(f, float64(6.05)))
	assert.Nil(t, Any(complex64(0), Nilify), testName(f, complex64(0)))
	assert.Equal(t, complex64(-9), *Any(complex64(-9)).(*complex64), testName(f, complex64(-9)))
	assert.Equal(t, complex64(0), *Any(complex64(0)).(*complex64), testName(f, complex64(0)))
	assert.Equal(t, complex64(7), *Any(complex64(7)).(*complex64), testName(f, complex64(7)))
	assert.Nil(t, Any(complex128(0), Nilify), testName(f, complex128(0)))
	assert.Equal(t, complex128(-9), *Any(complex128(-9)).(*complex128), testName(f, complex128(-9)))
	assert.Equal(t, complex128(0), *Any(complex128(0)).(*complex128), testName(f, complex128(0)))
	assert.Equal(t, complex128(7), *Any(complex128(7)).(*complex128), testName(f, complex128(7)))
	assert.Nil(t, Any("", Nilify), testName(f, ""))
	assert.Equal(t, "", *Any("").(*string), testName(f, ""))
	assert.Equal(t, "foo", *Any("foo").(*string), testName(f, "foo"))
	assert.Equal(t, "bar", *Any("bar").(*string), testName(f, "bar"))
	assert.Nil(t, Any(rune(0), Nilify), testName(f, rune(0)))
	assert.Equal(t, rune(0), *Any(rune(0)).(*rune), testName(f, rune(0)))
	assert.Equal(t, rune(9), *Any(rune(9)).(*rune), testName(f, rune(9)))
	assert.Equal(t, rune(39), *Any(rune(39)).(*rune), testName(f, rune(39)))
	assert.Nil(t, Any(time.Time{}, Nilify), testName(f, time.Time{}))
	assert.Equal(t, time.Time{}, *Any(time.Time{}).(*time.Time), testName(f, time.Time{}))
	assert.Equal(t, time.Unix(3242345, 055), *Any(time.Unix(3242345, 055)).(*time.Time), testName(f, time.Unix(3242345, 055)))
	assert.Nil(t, Any(time.Duration(0), Nilify), testName(f, time.Duration(0)))
	assert.Equal(t, -5*time.Minute, *Any(-5 * time.Minute).(*time.Duration), testName(f, -5*time.Minute))
	assert.Equal(t, time.Duration(0), *Any(time.Duration(0)).(*time.Duration), testName(f, time.Duration(0)))
	assert.Equal(t, 2*time.Hour, *Any(2 * time.Hour).(*time.Duration), testName(f, 2*time.Hour))
	assert.Equal(t, struct{}{}, *Any(struct{}{}).(*struct{}), testName(f, struct{}{}))
	assert.Equal(t, reflect.Method{}, *Any(reflect.Method{}).(*reflect.Method), testName(f, reflect.Method{}))
	assert.Equal(t, reflect.Method{Name: "name"}, *Any(reflect.Method{Name: "name"}).(*reflect.Method), testName(f, reflect.Method{Name: "name"}))
	var stringer fmt.Stringer
	assert.Nil(t, Any(stringer, Nilify), testName(f, stringer))
}

func TestBool(t *testing.T) {
	f := "Bool"
	assert.Nil(t, Bool(false, Nilify), testName(f, false))
	assert.Equal(t, false, *Bool(false), testName(f, false))
	assert.Equal(t, true, *Bool(true), testName(f, true))
}

func TestByte(t *testing.T) {
	f := "Byte"
	assert.Nil(t, Byte(byte(0), Nilify), testName(f, byte(0)))
	assert.Equal(t, byte(0), *Byte(byte(0)), testName(f, byte(0)))
	assert.Equal(t, byte(10), *Byte(byte(10)), testName(f, byte(10)))
}

func TestBytes(t *testing.T) {
	f := "Bytes"
	assert.Nil(t, Bytes([]byte{}, Nilify), testName(f, []byte{}))
	assert.Equal(t, []byte{}, *Bytes([]byte{}), testName(f, []byte{}))
	assert.Equal(t, []byte{10, 55}, *Bytes([]byte{10, 55}), testName(f, []byte{10, 55}))
}

func TestInt(t *testing.T) {
	f := "Int"
	assert.Nil(t, Int(int(0), Nilify), testName(f, int(0)))
	assert.Equal(t, int(-7), *Int(int(-7)), testName(f, int(-7)))
	assert.Equal(t, int(0), *Int(int(0)), testName(f, int(0)))
	assert.Equal(t, int(10), *Int(int(10)), testName(f, int(10)))
}

func TestInt8(t *testing.T) {
	f := "Int8"
	assert.Nil(t, Int8(int8(0), Nilify), testName(f, int8(0)))
	assert.Equal(t, int8(-7), *Int8(int8(-7)), testName(f, int8(-7)))
	assert.Equal(t, int8(0), *Int8(int8(0)), testName(f, int8(0)))
	assert.Equal(t, int8(10), *Int8(int8(10)), testName(f, int8(10)))
}

func TestInt16(t *testing.T) {
	f := "Int16"
	assert.Nil(t, Int16(int16(0), Nilify), testName(f, int16(0)))
	assert.Equal(t, int16(-7), *Int16(int16(-7)), testName(f, int16(-7)))
	assert.Equal(t, int16(0), *Int16(int16(0)), testName(f, int16(0)))
	assert.Equal(t, int16(10), *Int16(int16(10)), testName(f, int16(10)))
}

func TestInt32(t *testing.T) {
	f := "Int32"
	assert.Nil(t, Int32(int32(0), Nilify), testName(f, int32(0)))
	assert.Equal(t, int32(-7), *Int32(int32(-7)), testName(f, int32(-7)))
	assert.Equal(t, int32(0), *Int32(int32(0)), testName(f, int32(0)))
	assert.Equal(t, int32(10), *Int32(int32(10)), testName(f, int32(10)))
}

func TestInt64(t *testing.T) {
	f := "Int64"
	assert.Nil(t, Int64(int64(0), Nilify), testName(f, int64(0)))
	assert.Equal(t, int64(-7), *Int64(int64(-7)), testName(f, int64(-7)))
	assert.Equal(t, int64(0), *Int64(int64(0)), testName(f, int64(0)))
	assert.Equal(t, int64(10), *Int64(int64(10)), testName(f, int64(10)))
}

func TestUint(t *testing.T) {
	f := "Uint"
	assert.Nil(t, Uint(uint(0), Nilify), testName(f, uint(0)))
	assert.Equal(t, uint(0), *Uint(uint(0)), testName(f, uint(0)))
	assert.Equal(t, uint(10), *Uint(uint(10)), testName(f, uint(10)))
}

func TestUint8(t *testing.T) {
	f := "Uint8"
	assert.Nil(t, Uint8(uint8(0), Nilify), testName(f, uint8(0)))
	assert.Equal(t, uint8(0), *Uint8(uint8(0)), testName(f, uint8(0)))
	assert.Equal(t, uint8(10), *Uint8(uint8(10)), testName(f, uint8(10)))
}

func TestUint16(t *testing.T) {
	f := "Uint16"
	assert.Nil(t, Uint16(uint16(0), Nilify), testName(f, uint16(0)))
	assert.Equal(t, uint16(0), *Uint16(uint16(0)), testName(f, uint16(0)))
	assert.Equal(t, uint16(10), *Uint16(uint16(10)), testName(f, uint16(10)))
}

func TestUint32(t *testing.T) {
	f := "Uint32"
	assert.Nil(t, Uint32(uint32(0), Nilify), testName(f, uint32(0)))
	assert.Equal(t, uint32(0), *Uint32(uint32(0)), testName(f, uint32(0)))
	assert.Equal(t, uint32(10), *Uint32(uint32(10)), testName(f, uint32(10)))
}

func TestUint64(t *testing.T) {
	f := "Uint64"
	assert.Nil(t, Uint64(uint64(0), Nilify), testName(f, uint64(0)))
	assert.Equal(t, uint64(0), *Uint64(uint64(0)), testName(f, uint64(0)))
	assert.Equal(t, uint64(10), *Uint64(uint64(10)), testName(f, uint64(10)))
}

func TestUintptr(t *testing.T) {
	f := "Uintptr"
	assert.Nil(t, Uintptr(uintptr(0), Nilify), testName(f, uintptr(0)))
	assert.Equal(t, uintptr(0), *Uintptr(uintptr(0)), testName(f, uintptr(0)))
	assert.Equal(t, uintptr(10), *Uintptr(uintptr(10)), testName(f, uintptr(10)))
}

func TestFloat32(t *testing.T) {
	f := "Float32"
	assert.Nil(t, Float32(float32(0), Nilify), testName(f, float32(0)))
	assert.Equal(t, float32(-9.50), *Float32(float32(-9.50)), testName(f, float32(-9.50)))
	assert.Equal(t, float32(-1), *Float32(float32(-1)), testName(f, float32(-1)))
	assert.Equal(t, float32(0), *Float32(float32(0)), testName(f, float32(0)))
	assert.Equal(t, float32(4), *Float32(float32(4)), testName(f, float32(4)))
	assert.Equal(t, float32(6.05), *Float32(float32(6.05)), testName(f, float32(6.05)))
}

func TestFloat64(t *testing.T) {
	f := "Float64"
	assert.Nil(t, Float64(float64(0), Nilify), testName(f, float64(0)))
	assert.Equal(t, float64(-9.50), *Float64(float64(-9.50)), testName(f, float64(-9.50)))
	assert.Equal(t, float64(-1), *Float64(float64(-1)), testName(f, float64(-1)))
	assert.Equal(t, float64(0), *Float64(float64(0)), testName(f, float64(0)))
	assert.Equal(t, float64(4), *Float64(float64(4)), testName(f, float64(4)))
	assert.Equal(t, float64(6.05), *Float64(float64(6.05)), testName(f, float64(6.05)))
}

func TestComplex64(t *testing.T) {
	f := "Complex64"
	assert.Nil(t, Complex64(complex64(0), Nilify), testName(f, complex64(0)))
	assert.Equal(t, complex64(-9), *Complex64(complex64(-9)), testName(f, complex64(-9)))
	assert.Equal(t, complex64(0), *Complex64(complex64(0)), testName(f, complex64(0)))
	assert.Equal(t, complex64(7), *Complex64(complex64(7)), testName(f, complex64(7)))
}

func TestComplex128(t *testing.T) {
	f := "Complex128"
	assert.Nil(t, Complex128(complex128(0), Nilify), testName(f, complex128(0)))
	assert.Equal(t, complex128(-9), *Complex128(complex128(-9)), testName(f, complex128(-9)))
	assert.Equal(t, complex128(0), *Complex128(complex128(0)), testName(f, complex128(0)))
	assert.Equal(t, complex128(7), *Complex128(complex128(7)), testName(f, complex128(7)))
}

func TestString(t *testing.T) {
	f := "String"
	assert.Nil(t, String("", Nilify), testName(f, ""))
	assert.Equal(t, "", *String(""), testName(f, ""))
	assert.Equal(t, "foo", *String("foo"), testName(f, "foo"))
	assert.Equal(t, "bar", *String("bar"), testName(f, "bar"))
}

func TestRune(t *testing.T) {
	f := "Rune"
	assert.Nil(t, Rune(rune(0), Nilify), testName(f, rune(0)))
	assert.Equal(t, rune(0), *Rune(rune(0)), testName(f, rune(0)))
	assert.Equal(t, rune(9), *Rune(rune(9)), testName(f, rune(9)))
	assert.Equal(t, rune(39), *Rune(rune(39)), testName(f, rune(39)))
}

func TestTime(t *testing.T) {
	f := "Time"
	assert.Nil(t, Time(time.Time{}, Nilify), testName(f, time.Time{}))
	assert.Equal(t, time.Time{}, *Time(time.Time{}), testName(f, time.Time{}))
	assert.Equal(t, time.Unix(3242345, 055), *Time(time.Unix(3242345, 055)), testName(f, time.Unix(3242345, 055)))
}

func TestDuration(t *testing.T) {
	f := "Duration"
	assert.Nil(t, Duration(time.Duration(0), Nilify), testName(f, time.Duration(0)))
	assert.Equal(t, -5*time.Minute, *Duration(-5 * time.Minute), testName(f, -5*time.Minute))
	assert.Equal(t, time.Duration(0), *Duration(time.Duration(0)), testName(f, time.Duration(0)))
	assert.Equal(t, 2*time.Hour, *Duration(2 * time.Hour), testName(f, 2*time.Hour))
}
