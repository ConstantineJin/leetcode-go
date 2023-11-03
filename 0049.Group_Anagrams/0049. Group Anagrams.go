package main

import "sort"

// 直接统计
//func groupAnagrams(strs []string) (res [][]string) {
//	// 哈希表的key为每个字母出现的次数组合（如“abbz”对应[26]int{1,2,0,0...,0,1}，value为符合该字符串的单词组成的[]string
//	mp := map[[26]int][]string{}
//	for _, str := range strs {
//		cnt := [26]int{}
//		for _, c := range str {
//			cnt[c-'a']++
//		}
//		mp[cnt] = append(mp[cnt], str)
//	}
//	for _, m := range mp {
//		res = append(res, m)
//	}
//	return
//}

// 将排序后的单词作为map的key
func groupAnagrams(strs []string) (res [][]string) {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		mp[string(s)] = append(mp[string(s)], str)
	}
	for _, m := range mp {
		res = append(res, m)
	}
	return
}
