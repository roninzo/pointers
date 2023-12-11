//go:build go1.20
// +build go1.20

package pointers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew120(t *testing.T) {
	type foo struct {
		bar string
	}
	f := "[go1.20]New"
	assert.Nil(t, New(struct{}{}, Nilify), testName(f, struct{}{}))
	assert.Nil(t, New(foo{}, Nilify), testName(f, foo{}))
	assert.Nil(t, New(foo{bar: ""}, Nilify), testName(f, foo{}))
	assert.Equal(t, struct{}{}, *New(struct{}{}), testName(f, struct{}{}))
	assert.Equal(t, foo{}, *New(foo{}), testName(f, foo{}))
	assert.Equal(t, foo{bar: "foo"}, *New(foo{bar: "foo"}), testName(f, foo{bar: "foo"}))
	assert.NotEqual(t, foo{bar: "bar"}, *New(foo{bar: "foo"}), testName(f, foo{bar: "foo"}))
	var err error
	assert.Nil(t, New(err, Nilify), testName(f, err))
	assert.Equal(t, err, *New(err), testName(f, err))
	assert.Equal(t, errors.New("foo"), *New(errors.New("foo")), testName(f, errors.New("foo")))
	assert.Equal(t, fmt.Errorf("foo %s", "bar"), *New(fmt.Errorf("foo %s", "bar")), testName(f, fmt.Errorf("foo %s", "bar")))
}
