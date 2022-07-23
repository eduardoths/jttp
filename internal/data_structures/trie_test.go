package datastructures_test

import (
	"testing"

	datastructures "github.com/eduardoths/jttp/internal/data_structures"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	type TestCase struct {
		args   [][]rune
		before func(t *testing.T, tn *datastructures.TrieNode[rune, int], args ...[]rune)
		assert func(t *testing.T, tn *datastructures.TrieNode[rune, int], args ...[]rune)
	}

	testCases := map[string]TestCase{
		"it should find test": {
			args: [][]rune{[]rune("test")},
			before: func(t *testing.T, tn *datastructures.TrieNode[rune, int], args ...[]rune) {
				tn.Insert(args[0], 321)
			},
			assert: func(t *testing.T, tn *datastructures.TrieNode[rune, int], args ...[]rune) {
				v := tn.Find(args[0])
				lastKey := args[0][len(args[0])-1]
				want := &datastructures.TrieNode[rune, int]{
					Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
					IsTerminal: true,
					Key:        lastKey,
					Value:      pointy.Int(321),
				}
				assert.Equal(t, want, v)
			},
		},
		"it should find subparts of test": {
			args: [][]rune{[]rune("test")},
			before: func(t *testing.T, tn *datastructures.TrieNode[rune, int], args ...[]rune) {
				tn.Insert(args[0], 1234)
			},
			assert: func(t *testing.T, tn *datastructures.TrieNode[rune, int], args ...[]rune) {
				test := args[0]
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
				tc.before(t, root, tc.args...)
			}

			tc.assert(t, root, tc.args...)
		})
	}
}

func TestTrie_Insert(t *testing.T) {
	type args struct {
		key   []rune
		value int
	}

	type testCase struct {
		it   string
		in   []args
		want *datastructures.TrieNode[rune, int]
	}

	testCases := []testCase{
		{
			it: "should have a terminal children containing rune 't'",
			in: []args{{[]rune("t"), 0}},
			want: &datastructures.TrieNode[rune, int]{
				Children: map[rune]*datastructures.TrieNode[rune, int]{
					't': {
						Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
						IsTerminal: true,
						Key:        't',
						Value:      pointy.Int(0),
					},
				},
				IsTerminal: false,
				Key:        rune(0),
				Value:      nil,
			},
		},
		{
			it: "should have a terminal children with keys 't' and 'a'",
			in: []args{
				{[]rune("t"), 7},
				{[]rune("a"), -12},
			},
			want: &datastructures.TrieNode[rune, int]{
				Children: map[rune]*datastructures.TrieNode[rune, int]{
					't': {
						Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
						IsTerminal: true,
						Key:        't',
						Value:      pointy.Int(7),
					},
					'a': {
						Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
						IsTerminal: true,
						Key:        'a',
						Value:      pointy.Int(-12),
					},
				},
				IsTerminal: false,
				Key:        rune(0),
				Value:      nil,
			},
		},
		{
			it: "should have a terminal children containing key 'te' and containing key 'ta'",
			in: []args{
				{[]rune("te"), 1234},
				{[]rune("ta"), 4321},
			},
			want: &datastructures.TrieNode[rune, int]{
				Children: map[rune]*datastructures.TrieNode[rune, int]{
					't': {
						Children: map[rune]*datastructures.TrieNode[rune, int]{
							'e': {
								Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
								IsTerminal: true,
								Key:        'e',
								Value:      pointy.Int(1234),
							},
							'a': {
								Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
								IsTerminal: true,
								Key:        'a',
								Value:      pointy.Int(4321),
							},
						},
						IsTerminal: false,
						Key:        't',
						Value:      nil,
					},
				},
				IsTerminal: false,
				Key:        rune(0),
				Value:      nil,
			},
		},
		{
			it: "should have a terminal children for 't' and one for 'te'",
			in: []args{
				{[]rune("te"), 1234},
				{[]rune("t"), 4321},
			},
			want: &datastructures.TrieNode[rune, int]{
				Children: map[rune]*datastructures.TrieNode[rune, int]{
					't': {
						Children: map[rune]*datastructures.TrieNode[rune, int]{
							'e': {
								Children:   make(map[rune]*datastructures.TrieNode[rune, int]),
								IsTerminal: true,
								Key:        'e',
								Value:      pointy.Int(1234),
							},
						},
						IsTerminal: true,
						Key:        't',
						Value:      pointy.Int(4321),
					},
				},
				IsTerminal: false,
				Key:        rune(0),
				Value:      nil,
			},
		},
	}

	for _, scenario := range testCases {
		t.Run(scenario.it, func(t *testing.T) {
			root := datastructures.NewTrie[rune, int](rune(0), nil)
			for _, arg := range scenario.in {
				root.Insert(arg.key, arg.value)
			}

			assert.Equal(t, scenario.want, root)
		})
	}

}
