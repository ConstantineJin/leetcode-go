package main

import "strconv"

func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	mem := make([][1 << 10]int, m)
	for i := range mem {
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, mask int, isLimit, isNum bool) (res int)
	dfs = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		if !isLimit && isNum {
			p := &mem[i][mask]
			if *p >= 0 { // 之前计算过
				return *p
			}
			defer func() { *p = res }()
		}
		if !isNum { // 可以跳过当前位
			res += dfs(i+1, mask, false, false)
		}
		var d int
		if !isNum {
			d = 1
		}
		up := 9
		if isLimit { // 如果前面所填数字和 n 都一样，则当前位至多填 s[i]
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ {
			if mask>>d&1 == 0 { // 如果 d 不在 mask 中，说明之前没有填过 d
				res += dfs(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return dfs(0, 0, true, false)
}
