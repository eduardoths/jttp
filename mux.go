package jttp

import (
	"strings"

	datastructures "github.com/eduardoths/jttp/internal/data_structures"
)

type mux struct {
	trie *datastructures.TrieNode[string, Handler]
}

func NewMux() *mux {
	var defaultHandler Handler
	return &mux{
		trie: datastructures.NewTrie("", &defaultHandler),
	}
}

func (m *mux) Add(method string, pattern string, handler Handler) {
	var fullPattern = make([]string, 0)
	fullPattern = append(fullPattern, strings.ToUpper(method))
	urlPatterns := strings.SplitAfter(pattern, "/")
	if urlPatterns[len(urlPatterns)-1] == "" {
		urlPatterns = urlPatterns[:len(urlPatterns)-1]
	}
	fullPattern = append(fullPattern, urlPatterns...)
	m.trie.Insert(fullPattern, handler)
}

func (m *mux) Search(method string, route string) Handler {
	var fullPattern = make([]string, 0)
	fullPattern = append(fullPattern, strings.ToUpper(method))
	urlPatterns := strings.SplitAfter(route, "/")
	if urlPatterns[len(urlPatterns)-1] == "" {
		urlPatterns = urlPatterns[:len(urlPatterns)-1]
	}
	fullPattern = append(fullPattern, urlPatterns...)

	if node := m.trie.Find(fullPattern); node != nil {
		return *node.Value
	}
	return NotFoundHandler
}
