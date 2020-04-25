package ensure

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func fatal(t *testing.T, desc string, a... interface{})  {
	_, file, line, _ := runtime.Caller(2)
	t.Fatalf("file:%v:%v:%s\n", file, line, fmt.Sprint(a))
}
// Nil ensures v is nil.
func Nil(t *testing.T, v interface{}, a ...interface{}) {
	if v != nil {
		fatal(t, "should be nil", a)
	}
}
func NotNil(t *testing.T, v interface{}, a ...interface{}) {
	if v == nil {
		fatal(t, "should not be nil", a)
	}
}
func DeepEqual(t *testing.T, actual, expected interface{}, a ...interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		fatal(t, "should equal", a)
	}
}

func True(t *testing.T, value bool, a ...interface{})  {
	if !value {
		fatal(t, "should be true", a)
	}
}