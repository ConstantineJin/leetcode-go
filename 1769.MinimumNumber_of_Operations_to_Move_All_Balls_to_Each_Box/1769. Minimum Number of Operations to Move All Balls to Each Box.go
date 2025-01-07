package main

func minOperations(boxes string) []int {
	n := len(boxes)
	ans := make([]int, n)
	for i, cnt := 1, 0; i < n; i++ {
		cnt += int(boxes[i-1] - '0')
		ans[i] = ans[i-1] + cnt
	}
	for i, cnt, s := n-2, 0, 0; i >= 0; i-- {
		cnt += int(boxes[i+1] - '0')
		s += cnt
		ans[i] += s
	}
	return ans
}
