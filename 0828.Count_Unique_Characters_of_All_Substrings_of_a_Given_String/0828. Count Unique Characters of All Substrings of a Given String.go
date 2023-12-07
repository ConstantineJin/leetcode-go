package main

func uniqueLetterString(s string) (ans int) {
	last0, last1, total := [26]int{}, [26]int{}, 0 // last0和last1分别记录字母出现的上个和上上个下标，total记录以s[i]结尾的子串的countUniqueChars值的和
	for i := range last0 {
		last0[i], last1[i] = -1, -1 // 初始化last0和last1为-1
	}
	for i, c := range s { // 遍历s，可以理解为从空串开始每次往末尾添加一个字符
		c -= 'A'
		total += i - 2*last0[c] + last1[c] // 增加了i-last0[c]，减少了last0[c]-last1[c]
		ans += total
		last1[c], last0[c] = last0[c], i // 更新last0和last1
	}
	return
}
