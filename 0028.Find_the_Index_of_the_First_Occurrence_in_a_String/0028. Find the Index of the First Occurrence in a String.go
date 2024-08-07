package main

// 方法1: 标准库
//func strStr(haystack, needle string) int {
//	return strings.Index(haystack, needle)
//}

// 方法2: 切片
//func strStr(haystack, needle string) int {
//	n, m := len(haystack), len(needle)
//	if m > n { // 模式串比主串长，则模式串不可能是主串的子串
//		return -1
//	}
//	for i := 0; i <= n-m; i++ {
//		if haystack[i:i+m] == needle {
//			return i
//		}
//	}
//	return -1
//}

// 方法3: 暴力
//func strStr(haystack, needle string) int {
//	n, m := len(haystack), len(needle)
//	if m > n { // 模式串比主串长，则模式串不可能是主串的子串
//		return -1
//	}
//next:
//	for i := 0; i <= n-m; i++ {
//		for j := 0; j < m; j++ {
//			if haystack[i+j] != needle[j] {
//				continue next
//			}
//		}
//		return i
//	}
//	return -1
//}

// 方法4: KMP算法
func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m > n { // 模式串比主串长，则模式串不可能是主串的子串
		return -1
	}

	// 线性时空复杂度求 next 数组
	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ { // j 表示当前匹配的公共前后缀的长度
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
