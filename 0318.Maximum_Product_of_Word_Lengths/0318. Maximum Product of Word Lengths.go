package main

func maxProduct(words []string) (ans int) {
	n := len(words)
	masks := make([]int, n)
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a') // masks[i]的二进制每一位表明是否含有对应字母，如masks[0]==1表明含有字母a
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// 使用位运算将判断两个单词是否含有公共字母降低为O(1)时间复杂度
			if masks[i]&masks[j] == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}
