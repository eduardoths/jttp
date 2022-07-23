package jttp

import (
	"net/http"
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
