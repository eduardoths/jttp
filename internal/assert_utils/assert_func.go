package assertutils

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertFunc(t *testing.T, want, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	wantFunc := runtime.FuncForPC(reflect.ValueOf(want).Pointer()).Name()
	actualFunc := runtime.FuncForPC(reflect.ValueOf(actual).Pointer()).Name()
	return assert.Equal(t, wantFunc, actualFunc, msgAndArgs...)
}
