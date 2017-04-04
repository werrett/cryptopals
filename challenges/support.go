package challenges

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
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
	val1 := reflect.ValueOf(a)
	val2 := reflect.ValueOf(b)

	if val1.Kind() != val2.Kind() {
		t.Errorf("Values diff kind %v - Got %v", val1.Kind(), val2.Kind())
	}

	switch val1.Kind() {
	case reflect.Slice:
		if !bytes.Equal(val1.Bytes(), val2.Bytes()) {
			t.Errorf("Arrays not equal %s - Got %s", a, b)
		}
	// FIXME: Not handling pointers

	// Types: Int, Bool, Float32, Float64
	default:
		if a != b {
			t.Errorf("Expected %v - Got %v", a, b)
		}
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
