package src

import (
	"errors"
	"fmt"
	"sync"
	"unicode/utf8"
)

type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

type Trie struct {
	Root *TrieNode
	mu   sync.RWMutex
}

func NewTrie() *Trie {
	return &Trie{Root: &TrieNode{Children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) error {
	if word == "" {
		return errors.New("cannot insert empty word")
	}
	if !utf8.ValidString(word) {
		return errors.New("word contains invalid UTF-8 characters")
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.Root
	for _, c := range word {
		if _, exists := node.Children[c]; !exists {
			node.Children[c] = &TrieNode{
				Children: make(map[rune]*TrieNode),
				IsEnd:    false,
			}
		}
		node = node.Children[c]
	}

	// Only set isEnd at the last node of the word
	if node.IsEnd {
		return nil
	}
	node.IsEnd = true
	return nil
}

func (t *Trie) Search(word string) (bool, error) {
	if word == "" {
		return false, errors.New("cannot search empty word")
	}
	if !utf8.ValidString(word) {
		return false, errors.New("word contains invalid UTF-8 characters")
	}

	t.mu.RLock()
	defer t.mu.RUnlock()

	node := t.Root
	for _, c := range word {
		if _, exists := node.Children[c]; !exists {
			return false, nil
		}
		node = node.Children[c]
	}
	return node.IsEnd, nil
}

func (t *Trie) Delete(word string) error {
	if word == "" {
		return errors.New("cannot delete empty word")
	}
	if !utf8.ValidString(word) {
		return errors.New("word contains invalid UTF-8 characters")
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	// If the word doesn't exist, deleteHelper will return false, but we treat it as a no-op
	if deleted := t.deleteHelper(t.Root, word, 0); !deleted {
		fmt.Print("Word '%s' not found in the Trie during deletion", word)
	}
	return nil
}

func (t *Trie) deleteHelper(root *TrieNode, word string, index int) bool {
	// Base case: Reached the end of the word
	if index == len(word) {
		if !root.IsEnd {
			return false // Word doesn't exist
		}
		root.IsEnd = false             // Unmark the end of the word
		return len(root.Children) == 0 // Can delete if no children
	}

	// Get the current character
	c := rune(word[index])
	nextNode, exists := root.Children[c]
	if !exists {
		return false // Word doesn't exist
	}

	// Recursively delete in the subtree
	shouldDeleteCurrentNode := t.deleteHelper(nextNode, word, index+1)

	// If the child node should be deleted, remove it from the current node's children
	if shouldDeleteCurrentNode {
		delete(root.Children, c)
		return len(root.Children) == 0 && !root.IsEnd
	}
	return false
}

// create a trie, create a node. add logic, search logic, delete logic.
func TrieRun() {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("ap")
	trie.Insert("a")
	fmt.Println("apple exist? ", trie.Search("apple"))
	trie.Delete("apple")
	fmt.Println("apple exist? ", trie.Search("apple"))
}
