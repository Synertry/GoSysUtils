/*
 *           GoSysUtils
 *     Copyright (c) Synertry 2022 - 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Package Trie implements a simple trie data structure for storing dictionary with strings.
// Source: https://youtu.be/H-6-8_p88r0 (JamieGo)
package Trie

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents a node in the trie
type Node struct {
	Children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create a new Trie
func InitTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // Assuming lowercase letters a-z
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{}
		}
		currentNode = currentNode.Children[charIndex]
	}
	currentNode.isEnd = true // Mark the end of the word
}

// Search will take in a word and RETURN true if that word is included in the trie
// same walking logic as Insert, but we don't need to create nodes
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false // character is not found, word does not exist
		}
		currentNode = currentNode.Children[charIndex]
	}
	return currentNode.isEnd // Return true if we reached the end of a word
}
