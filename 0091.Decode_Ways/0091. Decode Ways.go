package main

// 递归
//func numDecodings(s string) int {
//	var n = len(s)
//	var mem = make([]int, n+2)
//	for i := range mem {
//		mem[i] = -1
//	}
//	mem[n], mem[n+1] = 1, 0
//	var dfs func(i int) int
//	dfs = func(i int) int {
//		if mem[i] != -1 {
//			return mem[i]
//		}
//		switch s[i] {
//		case '0':
//			mem[i] = 0
//		case '1':
//			mem[i] = dfs(i+1) + dfs(i+2)
//		case '2':
//			if i < n-1 && s[i+1] <= '6' {
//				mem[i] = dfs(i+1) + dfs(i+2)
//				break
//			}
//			mem[i] = dfs(i + 1)
//		default:
//			mem[i] = dfs(i + 1)
//		}
//		return mem[i]
//	}
//	return dfs(0)
//}

// 递推，常量空间
func numDecodings(s string) int {
	var n = len(s)
	var a, b, c = 0, 1, 0 // a = f[i-2], b = f[i-1], c = f[i]
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			c += a
		}
		a, b = b, c
	}
	return c
}
