package main

func findAnagrams(s string, p string) (res []int) {
	if len(s) < len(p) { // s长度比p小，自然不可能包含字母异位词，返回0
		return
	}
	cnt, window := [26]int{}, [26]int{} // cnt记录p中每个字母出现的次数，window维护s中每一个长为len(p)的窗口中各字母的出现次数
	for i, ch := range p {              // 初始化cnt和window
		cnt[ch-'a']++
		window[s[i]-'a']++
	}
	for i := 0; i < len(s)-len(p); i++ {
		if window == cnt { // 如果当前window正好覆盖了字母异位词
			res = append(res, i)
		}
		window[s[i]-'a']-- // window后移一格
		window[s[i+len(p)]-'a']++
	}
	if window == cnt { // 再判断最后s中最后len(p)个是否是字母异位词
		res = append(res, len(s)-len(p))
	}
	return
}
