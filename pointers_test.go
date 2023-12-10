package pointers

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// NOTE: The different data types use cases are being tested directly without
// passing them to interface{}, this test is more verbose than usual.
func TestNew(t *testing.T) {
	testName := func(x any) string { return fmt.Sprintf("Test: New(%v)", x) }
	{
		var x string
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = "foo"
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = "bar\nfoo"
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x int
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = 5
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = -12
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x uint64
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = 5434892638576219563
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = 93826762376
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x float64
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = 45.99
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = -1.33
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x complex128
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = complex(10, 11)
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = complex(-10, -11)
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x bool
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = false
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = true
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x time.Time
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = time.Now()
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = time.Now().Add(24 * time.Hour)
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x time.Duration
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = 5 * time.Minute
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = time.Since(time.Date(2023, time.December, 9, 18, 40, 0, 0, time.UTC))
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x error
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = errors.New("foo")
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = fmt.Errorf("bar %v", x.Error())
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x rune
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = rune('ç')
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = rune('ä')
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
	{
		var x byte
		p := New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		p = New(x, Nilify)
		assert.Nil(t, p, testName(x))
		x = byte('[')
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
		x = byte('>')
		p = New(x)
		assert.NotNil(t, p, testName(x))
		assert.Equal(t, x, *p, testName(x))
	}
}
