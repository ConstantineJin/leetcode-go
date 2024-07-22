package main

func maxOperations(s string) (ans int) {
	if s[len(s)-1] == '0' {
		s += "1"
	}
	var arr []int
	for i, c := range s {
		if c == '1' {
			if i == 0 || i > 0 && s[i-1] == '0' {
				arr = append(arr, 1)
			}
			if i > 0 && s[i-1] == '1' {
				arr[len(arr)-1]++
			}
		}
	}
	m := len(arr)
	for i, v := range arr {
		ans += v * (m - i - 1)
	}
	return
}
