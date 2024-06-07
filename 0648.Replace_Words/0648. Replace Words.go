package main

import "strings"

// Trie 字典树
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// Insert 向字典树中插入单词 word
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// GetRoot 获取单词 word 的最短词根，如果字典树中不包含 word 的词根返回空字符串
func (t *Trie) GetRoot(word string) string {
	node := t
	for i, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			return ""
		}
		node = node.children[ch]
		if node.isEnd {
			return word[:i+1]
		}
	}
	return ""
}

func replaceWords(dictionary []string, sentence string) string {
	var trie Trie
	for _, word := range dictionary {
		trie.Insert(word)
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if root := trie.GetRoot(word); root != "" {
			words[i] = root
		} else {
			words[i] = word
		}
	}
	return strings.Join(words, " ")
}
