package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var mp, queue, ans = make(map[string]bool), []string{beginWord}, 1
	for _, w := range wordList {
		mp[w] = true
	}
	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			var word = queue[0]
			queue = queue[1:]
			if word == endWord {
				return ans
			}
			for j := 0; j < len(word); j++ {
				for k := 'a'; k <= 'z'; k++ {
					var newWord = word[:j] + string(k) + word[j+1:]
					if mp[newWord] == true {
						queue, mp[newWord] = append(queue, newWord), false
					}
				}
			}
		}
		ans++
	}
	return 0
}
