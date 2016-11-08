// Package stack provides a very small and clean API for using the built-in panic
// and recover functions to abort and resume execution of a goroutine.
package stack

import "runtime/debug"

// Cause wraps errors handled by Abort and Resume.
type Cause struct {
	Err error
}

// Abort will abort even if the supplied error is nil.
func Abort(err error) {
	panic(&Cause{err})
}

// AbortIf will only abort if the supplied error is present.
func AbortIf(err error) {
	if err != nil {
		Abort(err)
	}
}

// Resume will try to recover an earlier call to Abort and call fn if an error
// has been recovered. It will not recover direct calls to the built-in panic
// function.
//
// Note: If the built-in panic function has been called with nil, a call to
// Resume will discard that panic and continue execution.
func Resume(fn func(error)) {
	val := recover()
	if cause, ok := val.(*Cause); ok {
		fn(cause.Err)
		return
	} else if val != nil {
		panic(val)
	}
}

// Trace returns a formatted stack trace of the original call to Abort during
// a call to Resume.
func Trace() []byte {
	return debug.Stack()
}
