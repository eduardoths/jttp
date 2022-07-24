package datastructures

type TrieNode[K comparable, V any] struct {
	Children   map[K]*TrieNode[K, V]
	IsTerminal bool
	Key        K
	Value      *V
}

func NewTrie[K comparable, V any](key K, value *V) *TrieNode[K, V] {
	return &TrieNode[K, V]{
		Children:   make(map[K]*TrieNode[K, V]),
		IsTerminal: false,
		Key:        key,
		Value:      value,
	}
}

func (tn *TrieNode[K, V]) Find(key []K) *TrieNode[K, V] {
	x := tn
	for i := 0; i < len(key); i++ {
		if x.Children[key[i]] == nil {
			return nil
		}
		x = x.Children[key[i]]
	}
	return x
}

func (tn *TrieNode[K, V]) Insert(key []K, value V) {
	x := tn
	for i := 0; i < len(key); i++ {
		if x.Children[key[i]] == nil {
			x.Children[key[i]] = NewTrie[K, V](key[i], nil)
		}
		x = x.Children[key[i]]
	}
	x.Value = &value
	x.IsTerminal = true
}

func (tn *TrieNode[K, V]) ClosestMatch(key []K) []K {
	keys := []K{}
	x := tn
	for i := 0; i < len(key); i++ {
		if x.Children[key[i]] == nil {
			return keys
		}
		keys = append(keys, key[i])
		x = x.Children[key[i]]
	}
	return keys
}
