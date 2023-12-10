//go:build go1.20
// +build go1.20

package pointers

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew120(t *testing.T) {
	var err error
	f := "[go1.20]New"
	assert.Nil(t, New(err, Nilify), testName(f, err))
	assert.Equal(t, err, *New(err), testName(f, err))
	assert.Equal(t, errors.New("foo"), *New(errors.New("foo")), testName(f, errors.New("foo")))
	assert.Equal(t, fmt.Errorf("foo %s", "bar"), *New(fmt.Errorf("foo %s", "bar")), testName(f, fmt.Errorf("foo %s", "bar")))
	assert.Nil(t, New(struct{}{}, Nilify), testName(f, struct{}{}))
	assert.Nil(t, New(reflect.Method{}, Nilify), testName(f, reflect.Method{}))
	assert.Equal(t, struct{}{}, *New(struct{}{}), testName(f, struct{}{}))
	assert.Equal(t, reflect.Method{}, *New(reflect.Method{}), testName(f, reflect.Method{}))
	assert.Equal(t, reflect.Method{Name: "name"}, *New(reflect.Method{Name: "name"}), testName(f, reflect.Method{Name: "name"}))
	assert.NotEqual(t, reflect.Method{Name: "foo"}, *New(reflect.Method{Name: "name"}), testName(f, reflect.Method{Name: "name"}))
	assert.Equal(t, os.Interrupt, *New(os.Interrupt), testName(f, os.Interrupt))
}
