package stack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errAbortTest = errors.New("foo")

func TestAbortIfResume(t *testing.T) {
	var test error

	func() {
		defer Resume(func(err error) {
			test = err
		})

		AbortIf(errAbortTest)
	}()

	assert.Equal(t, errAbortTest, test)
}

func TestPanic(t *testing.T) {
	defer func() {
		recover()
	}()

	var test error

	func() {
		defer Resume(func(err error) {
			test = err
		})

		panic(errAbortTest)
	}()

	assert.Nil(t, test)
}

func TestAbortIfNil(t *testing.T) {
	var test error

	func() {
		defer Resume(func(err error) {
			test = err
		})

		AbortIf(nil)
	}()

	assert.Nil(t, test)
}

func TestStack(t *testing.T) {
	var test error
	var stack string

	func() {
		defer Resume(func(err error) {
			test = err
			stack = string(Trace())
		})

		Abort(errAbortTest)
	}()

	assert.Equal(t, errAbortTest, test)
	assert.Contains(t, stack, "stack.Abort")
}
