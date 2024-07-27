package main

func getSmallestString(s string, k int) string {
	t := []byte(s)
	for i, c := range s {
		disToA := int(min(c-'a', 'z'-c+1))
		if k >= disToA { // 贪心，顺序遍历 s，只要能将当前字符转换成'a'就转换
			t[i] = 'a'
			k -= disToA
		} else {
			t[i] = byte(c - int32(k)) // 剩余的 k 不足以将 s[i] 变成 'a'，就将其变为尽可能小的字符
			break
		}
	}
	return string(t)
}
