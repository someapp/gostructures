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

type runeWord = []rune

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
	}
}

func (t *Trie) AddWord(word string) {
	t.root.addWord(runeWord(word))
}

func (t *Trie) ContainsPrefix(word string) bool {
	_, ok := t.root.nodeForWord(runeWord(word))
	return ok
}

func (t *Trie) ContainsWord(word string) bool {
	if node, ok := t.root.nodeForWord(runeWord(word)); ok {
		return node.isWord
	}
	return false
}

func (t *Trie) Words() []string {
	return t.WordsWithPrefix("")
}

func (t *Trie) WordsWithPrefix(prefix string) []string {
	runePrefix := runeWord(prefix)
	node, ok := t.root.nodeForWord(runePrefix)
	if !ok {
		return []string{}
	}

	words := make([]runeWord, 0)
	node.words(runePrefix, &words)

	rtn := make([]string, len(words))
	for i, word := range words {
		rtn[i] = string(word)
	}
	return rtn
}

func (t *trieNode) addWord(word runeWord) {
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

func (t *trieNode) containsWord(word runeWord) bool {
	if len(word) == 0 {
		return t.isWord
	}
	head, tail := word[0], word[1:]
	if node, ok := t.children[head]; ok {
		return node.containsWord(tail)
	}
	return false
}

func (t *trieNode) nodeForWord(word runeWord) (*trieNode, bool) {
	if len(word) == 0 {
		return t, true
	}
	head, tail := word[0], word[1:]
	if node, ok := t.children[head]; ok {
		return node.nodeForWord(tail)
	}
	return nil, false
}

func (t *trieNode) words(prefix runeWord, output *[]runeWord) {
	if t.isWord {
		*output = append(*output, prefix)
	}

	for head, child := range t.children {
		childPrefix := append(prefix, head)
		child.words(childPrefix, output)
	}
}
