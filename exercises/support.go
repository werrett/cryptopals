package challenges 

import (
  "fmt"
  "testing"
  "errors"
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
func log(msg string, args ...interface{}) {
  if Verbose {
    fmt.Printf(msg, args...)
  }
}