package challenges

import (
	"errors"
	"fmt"
	"testing"
)

// Message verbosity levels
const (
	LogInfo  = iota
	LogDebug = iota
)

// Check returned error value. Takes into account verbosity settings.
func check(e error) {
	if e != nil {
		panic(e) // TODO: should swap to log.Fatal(err)
	}
}

// Return an error with a specfic message
func newError(msg string) error {
	return errors.New(msg)
}

// Generic helper functions for tests
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf(
			"Expected %v - Got %v",
			a, b,
		)
	}
}

func expectErr(t *testing.T, a interface{}) {
	_, ok := a.(error)
	if !ok {
		t.Errorf("Expected error.")
	}
}

// Output helpers
func log(level int, msg string, args ...interface{}) {
	if Verbose > level {
		fmt.Printf(msg, args...)
	}
}
func print(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}
