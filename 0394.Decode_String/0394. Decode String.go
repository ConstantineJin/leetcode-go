package main

import "strconv"

func decodeString(s string) string {
	var ans, stack, depth, num, cnt = make([]byte, 0), make([][]byte, 1), 0, make([]byte, 0), 0
	for _, ch := range s {
		var c = byte(ch)
		switch ch {
		case '[':
			depth++
			stack[depth] = make([]byte, 0)
		case ']':
			for j := 0; j < cnt; j++ {
				ans = append(ans, stack[depth]...)
			}
			clear(stack[depth])
			depth--
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = append(num, c)
		default:
			cnt, _ = strconv.Atoi(string(num))
			clear(num)
			stack[depth] = append(stack[depth], c)
		}
	}
	return string(ans)
}
