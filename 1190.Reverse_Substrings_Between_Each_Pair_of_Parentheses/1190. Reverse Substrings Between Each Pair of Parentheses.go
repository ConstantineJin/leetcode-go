package main

func reverseParentheses(s string) string {
	n := len(s)

	// 使用栈预处理每一对括号的位置
	pair := make([]int, n)
	var stack []int
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i], pair[j] = j, i
		}
	}

	var ans []byte
	for i, step := 0, 1; i < n; i += step {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]  // 跳转到对应的括号
			step = -step // 遍历顺序反向
		} else {
			ans = append(ans, s[i]) // ans 中只包含非括号字符
		}
	}
	return string(ans)
}
