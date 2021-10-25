package strutil

import (
	"fmt"
	"reflect"
	"testing"
)

func Assert(t *testing.T, exp interface{}, act interface{}, msg string, args ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(act, exp) {
		msg = fmt.Sprintf(msg, args...)
		t.Errorf("%s\nactual: %v\nexpected: %v", msg, act, exp)
	}
}

func AssertPanics(t *testing.T, fn func(), msg string, args ...interface{}) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(msg, args...)
		}
	}()
	fn()
}
