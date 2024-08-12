package main

type Trie struct {
	children   [26]*Trie
	isFinished bool
}

type MagicDictionary struct{ *Trie }

func Constructor() MagicDictionary { return MagicDictionary{&Trie{}} }

func (md *MagicDictionary) BuildDict(dictionary []string) {
	for _, s := range dictionary {
		cur := md.Trie
		for _, c := range s {
			c -= 'a'
			if cur.children[c] == nil {
				cur.children[c] = &Trie{}
			}
			cur = cur.children[c]
		}
		cur.isFinished = true
	}
}

func (md *MagicDictionary) Search(searchWord string) bool {
	n := len(searchWord)
	var dfs func(node *Trie, i int, modified bool) bool
	dfs = func(node *Trie, i int, modified bool) bool {
		if i == n {
			return modified && node.isFinished
		}
		c := searchWord[i] - 'a'
		if node.children[c] != nil && dfs(node.children[c], i+1, modified) {
			return true
		}
		if !modified {
			for j, child := range node.children {
				if j != int(c) && child != nil && dfs(child, i+1, true) {
					return true
				}
			}
		}
		return false
	}
	return dfs(md.Trie, 0, false)
}
