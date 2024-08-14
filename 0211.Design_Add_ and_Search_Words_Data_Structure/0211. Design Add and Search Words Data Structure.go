package main

type node struct {
	children [26]*node
	isEnd    bool
}

type WordDictionary struct{ *node }

func Constructor() WordDictionary { return WordDictionary{&node{}} }

func (wd *WordDictionary) AddWord(word string) {
	cur := wd.node
	for _, c := range word {
		c -= 'a'
		if cur.children[c] == nil {
			cur.children[c] = &node{}
		}
		cur = cur.children[c]
	}
	cur.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
	var dfs func(i int, node *node) bool
	dfs = func(i int, node *node) bool {
		if i == len(word) {
			return node.isEnd
		}
		c := word[i]
		if c == '.' {
			for _, child := range node.children {
				if child != nil && dfs(i+1, child) {
					return true
				}
			}
			return false
		} else {
			c -= 'a'
			child := node.children[c]
			return child != nil && dfs(i+1, child)
		}
	}
	return dfs(0, wd.node)
}
