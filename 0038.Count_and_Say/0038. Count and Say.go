package main

import "strconv"

func countAndSay(n int) (ans string) {
	if n == 1 {
		return "1"
	}
	pre := countAndSay(n - 1)
	n, i, j := len(pre), 0, 0
	for i < n {
		for j = i + 1; j < n && pre[j] == pre[i]; j++ {
		}
		ans, i = ans+strconv.Itoa(j-i)+string(pre[i]), j
	}
	return
}
