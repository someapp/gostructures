package gostructures

import (
	"reflect"
	"sort"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	if trie.ContainsWord("test") {
		t.Errorf("should not contain test")
	}
	if trie.ContainsWord("") {
		t.Errorf("should not contain")
	}

	trie.AddWord("redherring")
	trie.AddWord("test")
	trie.AddWord("tests")

	for _, word := range []string{"", "t", "te", "tes", "testy"} {
		if trie.ContainsWord(word) {
			t.Errorf("should not contain %v", word)
		}
	}

	for _, word := range []string{"test", "tests", "redherring"} {
		if !trie.ContainsWord(word) {
			t.Errorf("should contain %v", word)
		}
	}

	expected := []string{"redherring", "test", "tests"}
	sort.Strings(expected)
	actual := trie.Words()
	sort.Strings(actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected words with prefix %v, %v", expected, actual)
	}
	expected = []string{"test", "tests"}

	for _, prefix := range []string{"t", "te", "tes", "test"} {
		actual := trie.WordsWithPrefix(prefix)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected words with prefix %v, %v", expected, actual)
		}
	}
}
