package main

func maximumGain(s string, x, y int) int {
	var stackCnt0, stackCnt1 int // 不需要真的模拟一个栈，只需要统计字符 a 和 b 在栈中出现的次数
	f := func(p, q rune, u, v int) (res int) {
		for _, c := range s + "#" { // 在 s 结尾添加一个无关字符避免特判
			if c == 'a' || c == 'b' {
				if c == q { // 优先匹配分值较高的子串
					if stackCnt0 > 0 {
						stackCnt0--
						res += u
					} else {
						stackCnt1++
					}
				} else {
					stackCnt0++
				}
			} else { // 其他字符可以被视作分隔符
				res += v * min(stackCnt0, stackCnt1)
				stackCnt0, stackCnt1 = 0, 0
			}
		}
		return
	}
	if x > y {
		return f('a', 'b', x, y)
	} else {
		return f('b', 'a', y, x)
	}
}
