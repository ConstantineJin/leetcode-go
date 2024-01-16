package main

import "strings"

//func count(num1, num2 string, minSum, maxSum int) int {
//	const mod = 1e9 + 7
//	var calc = func(s string) int {
//		var mem = make([][]int, len(s))
//		for i := range mem {
//			mem[i] = make([]int, min(9*len(s), maxSum)+1)
//			for j := range mem[i] {
//				mem[i][j] = -1
//			}
//		}
//		var dfs func(int, int, bool) int
//		dfs = func(i, sum int, isLimit bool) (res int) {
//			if sum > maxSum { // 非法
//				return
//			}
//			if i == len(s) {
//				if sum >= minSum { // 合法
//					return 1
//				}
//				return
//			}
//			if !isLimit {
//				var p = &mem[i][sum]
//				if *p >= 0 {
//					return *p
//				}
//				defer func() { *p = res }()
//			}
//			var up = 9
//			if isLimit {
//				up = int(s[i] - '0')
//			}
//			for d := 0; d <= up; d++ { // 枚举当前数位填 d
//				res = (res + dfs(i+1, sum+d, isLimit && d == up)) % mod
//			}
//			return
//		}
//		return dfs(0, 0, true)
//	}
//	var ans = calc(num2) - calc(num1) + mod // 避免负数
//	var sum int
//	for _, c := range num1 {
//		sum += int(c - '0')
//	}
//	if minSum <= sum && sum <= maxSum { // num1 是合法的，补回来
//		ans++
//	}
//	return ans % mod
//}

func count(num1, num2 string, minSum, maxSum int) int {
	const mod = 1e9 + 7
	var n = len(num2)
	num1 = strings.Repeat("0", n-len(num1)) + num1 // 补前导零，和 num2 对齐
	var mem = make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, min(9*n, maxSum)+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(int, int, bool, bool) int
	dfs = func(i, sum int, limitLow, limitHigh bool) (res int) {
		if sum > maxSum { // 非法
			return
		}
		if i == n {
			if sum >= minSum { // 合法
				return 1
			}
			return
		}
		if !limitLow && !limitHigh {
			var p = &mem[i][sum]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		var lo, hi = 0, 9
		if limitLow {
			lo = int(num1[i] - '0')
		}
		if limitHigh {
			hi = int(num2[i] - '0')
		}
		for d := lo; d <= hi; d++ { // 枚举当前数位填 d
			res = (res + dfs(i+1, sum+d, limitLow && d == lo, limitHigh && d == hi)) % mod
		}
		return
	}
	return dfs(0, 0, true, true)
}
