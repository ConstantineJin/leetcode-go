package main

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func partition(s string) (ans [][]string) {
	n := len(s)
	var path []string
	var dfs func(i, start int)
	dfs = func(i, start int) {
		if i == n {
			temp := make([]string, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		if i < n-1 { // 不在当前位置分割（i==n-1时必须分割）
			dfs(i+1, start)
		}
		if isPalindrome(s, start, i) { // 在当前位置分割
			path = append(path, s[start:i+1])
			dfs(i+1, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}
