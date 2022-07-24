package jttp

import (
	"strings"

	datastructures "github.com/eduardoths/jttp/internal/data_structures"
)

type muxNode *datastructures.TrieNode[string, Handler]

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
	splittenRoute := m.methodAndRouteToArray(method, route)
	// pure text patterns
	if node := m.trie.Find(splittenRoute); node != nil {
		return *node.Value
	}

	if handler := m.searchAfterPattern(m.trie, splittenRoute); handler != nil {
		return handler
	}

	return NotFoundHandler
}

func (m *mux) searchAfterPattern(currentNode *datastructures.TrieNode[string, Handler], url []string) Handler {
	closestMatch := m.trie.ClosestMatch(url)
	node := m.trie.Find(closestMatch)
	for k, v := range node.Children {
		isPatternSubstring := strings.HasPrefix(k, ":")
		if isPatternSubstring {
			return *v.Value
		}
	}
	return nil
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
