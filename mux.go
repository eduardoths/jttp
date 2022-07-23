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
	fullPattern := m.methodAndRouteToArray(method, pattern)
	m.trie.Insert(fullPattern, handler)
}

func (m *mux) Search(method string, route string) Handler {
	fullPattern := m.methodAndRouteToArray(method, route)
	if node := m.trie.Find(fullPattern); node != nil {
		return *node.Value
	}
	return NotFoundHandler
}

func (m *mux) methodAndRouteToArray(method string, route string) []string {
	var fullPattern = []string{}
	fullPattern = append(fullPattern, strings.ToUpper(method))
	urlPatterns := strings.SplitAfter(route, "/")
	if urlPatterns[len(urlPatterns)-1] == "" {
		urlPatterns = urlPatterns[:len(urlPatterns)-1]
	}
	fullPattern = append(fullPattern, urlPatterns...)
	return fullPattern
}
