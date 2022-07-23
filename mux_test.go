package jttp

import (
	"net/http"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMux_Add(t *testing.T) {
	type input struct {
		method  string
		pattern string
	}
	type testCase struct {
		it          string
		in          input
		testPattern []string
	}

	testCases := []testCase{
		{
			it:          "Should handle GET /",
			in:          input{http.MethodGet, "/"},
			testPattern: []string{"GET", "/"},
		},
		{
			it:          "Should handle GET /test",
			in:          input{http.MethodGet, "/test"},
			testPattern: []string{"GET", "/", "test"},
		},
		{
			it:          "Should handle POST /test/xpto",
			in:          input{http.MethodPost, "/test/xpto"},
			testPattern: []string{"POST", "/", "test/", "xpto"},
		},
		{
			it:          "Method input should be case insenstive",
			in:          input{"put", "/test/xpto"},
			testPattern: []string{"PUT", "/", "test/", "xpto"},
		},
	}

	for _, scenario := range testCases {
		t.Run(scenario.it, func(t *testing.T) {
			mux := NewMux()
			var callbackWasCalled = false

			var inHandler Handler = func() {
				callbackWasCalled = true
			}

			mux.Add(scenario.in.method, scenario.in.pattern, inHandler)

			handler := mux.trie.Find(scenario.testPattern)
			if !assert.NotNil(t, handler, "pattern doesn't match to a handler") {
				t.FailNow()
			}
			callback := *handler.Value
			callback()
			assert.True(t, callbackWasCalled, "function associated to pattern doesn't match")
		})
	}

}

func TestMux_Search(t *testing.T) {
	type inserts struct {
		method  string
		pattern string
		handler Handler
	}

	type input struct {
		method string
		route  string
	}

	type testCase struct {
		it      string
		inserts []inserts
		in      input
		want    Handler
	}

	var secondTestWantFunc = func() {}
	testCases := []testCase{
		{
			it:      "Should return not found handler",
			inserts: []inserts{},
			want:    NotFoundHandler,
		},
		{
			it: "Sould return the same handler",
			inserts: []inserts{
				{http.MethodPost, "/", secondTestWantFunc},
			},
			in:   input{http.MethodPost, "/"},
			want: secondTestWantFunc,
		},
	}

	for _, scenario := range testCases {
		t.Run(scenario.it, func(t *testing.T) {
			mux := NewMux()
			for _, insert := range scenario.inserts {
				mux.Add(insert.method, insert.pattern, insert.handler)
			}
			actual := mux.Search(scenario.in.method, scenario.in.route)

			wantPointer := runtime.FuncForPC(reflect.ValueOf(scenario.want).Pointer()).Name()
			actualPointer := runtime.FuncForPC(reflect.ValueOf(actual).Pointer()).Name()
			assert.Equal(t, wantPointer, actualPointer)
		})
	}
}
