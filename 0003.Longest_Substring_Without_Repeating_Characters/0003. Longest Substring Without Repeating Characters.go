package main

func lengthOfLongestSubstring(s string) (ans int) {
	window := [128]bool{} // 也可以用 map，这里为了效率用的数组
	left := 0
	for right, c := range s {
		for window[c] { // 加入 c 后，窗口内会有重复元素
			window[s[left]] = false
			left++
		}
		window[c] = true
		ans = max(ans, right-left+1) // 更新窗口长度最大值
	}
	return
}
