package main

func parseBoolExpr(expression string) bool {
	var st []rune // 使用栈模拟
	for _, c := range expression {
		switch c {
		case ',': // 忽略逗号
		case ')': // 遇到右括号
			var hasT, hasF bool
			for ; st[len(st)-1] != '('; st = st[:len(st)-1] { // 不断从栈中弹出元素，直到弹出左括号
				switch st[len(st)-1] {
				case 't':
					hasT = true
				case 'f':
					hasF = true
				}
			}
			op := st[len(st)-2] // 此时栈顶是左括号，第二个元素才是逻辑运算符
			st = st[:len(st)-2]
			switch op {
			case '!': // 对于非运算操作，直接取反
				if hasT {
					st = append(st, 'f')
				} else {
					st = append(st, 't')
				}
			case '&': // 对于与运算操作，只要有一个为假结果就为假
				if hasF {
					st = append(st, 'f')
				} else {
					st = append(st, 't')
				}
			case '|': // 对于或运算操作，只要有一个为真结果就为假
				if hasT {
					st = append(st, 't')
				} else {
					st = append(st, 'f')
				}
			}
		default:
			st = append(st, c)
		}

	}
	return st[len(st)-1] == 't'
}
