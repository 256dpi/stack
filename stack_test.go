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
	var set bool

	defer func() {
		recover()
		set = true
	}()

	var test error

	func() {
		defer Resume(func(err error) {
			test = err
		})

		panic(errAbortTest)
	}()

	assert.True(t, set)
	assert.Nil(t, test)
}

func TestAbortIfNil(t *testing.T) {
	var set bool

	func() {
		defer Resume(func(err error) {
			set = true
		})

		AbortIf(nil)
	}()

	assert.False(t, set)
}

func TestStack(t *testing.T) {
	var test error
	var trace string

	func() {
		defer Resume(func(err error) {
			test = err
			trace = string(Trace())
		})

		Abort(errAbortTest)
	}()

	assert.Equal(t, errAbortTest, test)
	assert.Contains(t, trace, "stack.Abort")
}

func BenchmarkAbortResume(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		func(){
			defer Resume(func(err error) {
				// do nothing
			})

			AbortIf(errAbortTest)
		}()
	}
}
