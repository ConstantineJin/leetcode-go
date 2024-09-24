package main

func maximumSubsequenceCount(text, pattern string) (ans int64) {
	a, b := pattern[0], pattern[1]
	var cntA, cntB int
	for i := range text {
		c := text[i]
		if c == b { // 必须先判断 c 与 b 是否相等，否则在处理 a == b 的时候会出错
			ans += int64(cntA)
			cntB++
		}
		if c == a {
			cntA++
		}
	}
	return ans + int64(max(cntA, cntB)) // 额外字母 a 如果插入在 text 首，则子序列增加 cntB 个；额外字母 b 如果插入在 text 尾，则子序列增加 cntA 个
}
