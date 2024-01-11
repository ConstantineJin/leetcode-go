package main

func addMinimum(word string) (ans int) {
	var n = len(word)
	for i := 0; i < n-1; i++ {
		if word[i] == word[i+1] { // 相邻两个字母相同
			ans += 2 // 中间要补另外两个字母
		} else if word[i:i+2] == "ac" || word[i:i+2] == "ba" || word[i:i+2] == "cb" { // 相邻两个字母中间差一个
			ans += 1 // 中间补一个字母
		}
	}
	return ans + int(word[0]-'a'+'c'-word[n-1]) // 首位不是a、末位不是c都需要补相应数量的字母
}
