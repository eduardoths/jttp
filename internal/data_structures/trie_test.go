package datastructures_test

import (
	"testing"

	datastructures "github.com/eduardoths/jttp/internal/data_structures"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	type TestCase struct {
		before func(t *testing.T, tn *datastructures.TrieNode[rune, int])
		assert func(t *testing.T, tn *datastructures.TrieNode[rune, int])
	}

	testCases := map[string]TestCase{
		"it should find test": {
			before: func(t *testing.T, tn *datastructures.TrieNode[rune, int]) {
				test := []rune("test")
				tn.Insert(test, 321)
			},
			assert: func(t *testing.T, tn *datastructures.TrieNode[rune, int]) {
				test := []rune("test")
				v := tn.Find(test)
				want := &datastructures.TrieNode[rune, int]{
					Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
					IsTerminal: true,
					Key:        't',
					Value:      pointy.Int(321),
				}
				assert.Equal(t, want, v)
			},
		},
		"it should find subparts of test": {
			before: func(t *testing.T, tn *datastructures.TrieNode[rune, int]) {
				test := []rune("test")
				tn.Insert(test, 1234)
			},
			assert: func(t *testing.T, tn *datastructures.TrieNode[rune, int]) {
				test := []rune("test")
				for i := range test[:len(test)-1] {
					actual := tn.Find(test[:i+1])
					t.Log(test[:i+1])
					assert.NotEmpty(t, actual)
					assert.Equal(t, false, actual.IsTerminal)
					assert.Equal(t, test[i], actual.Key, "asserts node key")
					assert.Nil(t, actual.Value, "node value is not nil")
					assert.NotNil(t, actual.Children[test[i+1]], "children is nil")
				}
			},
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			root := datastructures.NewTrie[rune, int](rune(0), nil)

			if tc.before != nil {
				tc.before(t, root)
			}

			tc.assert(t, root)
		})
	}
}
