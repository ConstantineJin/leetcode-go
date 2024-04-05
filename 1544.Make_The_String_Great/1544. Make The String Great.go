package main

func makeGood(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		for len(stack) > 1 && stack[len(stack)-1]^' ' == stack[len(stack)-2] {
			stack = stack[:len(stack)-2]
		}
	}
	return string(stack)
}
