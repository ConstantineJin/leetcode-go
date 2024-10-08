package main

func minSwaps(s string) (ans int) {
	var cnt int // 未匹配的 '[' 的个数
	for _, c := range s {
		if c == '[' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else { // 使用最右侧的左括号来交换
			cnt++
			ans++
		}
	}
	return
}
