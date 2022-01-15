package sync_context

import "testing"

func TestWithCancel(t *testing.T) {
	WithCancel()
}

func TestWithValue(t *testing.T) {
	WithValue()
}

func TestWithTimeOut(t *testing.T) {
	WithTimeOut()
}

func TestWithDeadline(t *testing.T) {
	WithDeadline()
}
