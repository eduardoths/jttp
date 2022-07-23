package datastructures_test

import (
	"testing"

	datastructures "github.com/eduardoths/jttp/internal/data_structures"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

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
