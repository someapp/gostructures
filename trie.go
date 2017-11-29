package gostructures

type Trie struct {
	root *trieNode
}

func NewTrie() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

type trieNode struct {
	isWord   bool
	children map[rune]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
	}
}

func (t *Trie) AddWord(word string) {
	t.root.addWord([]rune(word))
}

func (t *Trie) ContainsWord(word string) bool {
	return t.root.containsWord([]rune(word))
}

func (t *trieNode) addWord(word []rune) {
	if len(word) == 0 {
		t.isWord = true
		return
	}
	head, tail := word[0], word[1:]
	node, ok := t.children[head]
	if !ok {
		node = newTrieNode()
		t.children[head] = node
	}
	node.addWord(tail)
}

func (t *trieNode) containsWord(word []rune) bool {
	if len(word) == 0 {
		return t.isWord
	}
	head, tail := word[0], word[1:]
	if node, ok := t.children[head]; ok {
		return node.containsWord(tail)
	}
	return false
}
