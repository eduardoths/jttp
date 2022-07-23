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
				tn.Insert(test, 2131)
			},
			assert: func(t *testing.T, tn *datastructures.TrieNode[rune, int]) {
				test := []rune("test")
				v := tn.Find(test)
				want := &datastructures.TrieNode[rune, int]{
					Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
					IsTerminal: true,
					Key:        't',
					Value:      pointy.Int(2131),
				}
				assert.Equal(t, want, v)
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
