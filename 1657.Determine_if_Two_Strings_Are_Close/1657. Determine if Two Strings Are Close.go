package main

import (
	"slices"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) { // 两种操作都不改变字符串长度，如果原始长度不一致之间则两字符串必不接近
		return false
	}
	var mask1, mask2 int   // 用二进制的每个bit记录每个字母是否出现过
	var cnt1, cnt2 [26]int // 记录每个字母出现的次数
	for _, c := range word1 {
		mask1 |= 1 << (c - 'a')
		cnt1[c-'a']++
	}
	for _, c := range word2 {
		mask2 |= 1 << (c - 'a')
		cnt2[c-'a']++
	}
	sort.Ints(cnt1[:])
	sort.Ints(cnt2[:])
	return mask1 == mask2 && slices.Equal(cnt1[:], cnt2[:]) // 两种操作都不能创造未出现的字符，因此两字符串中出现过的字符必须严格相同，同时出现次数也必须一致
}
