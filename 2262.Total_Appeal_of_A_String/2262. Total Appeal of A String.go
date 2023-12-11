package main

func appealSum(s string) (ans int64) {
	sumG, last := 0, [26]int{} // sumG记录以s[i]结尾的字串的引力总和，last记录每个字母最后出现的位置
	for i := range last {
		last[i] = -1 // last初始化为-1
	}
	for i, c := range s {
		c -= 'a'
		sumG += i - last[c] // 从当前位置到当前元素上一次出现的位置，这些字串的引力都要+1
		ans += int64(sumG)
		last[c] = i // 更新当前元素对应的字母最后出现的位置
	}
	return
}
