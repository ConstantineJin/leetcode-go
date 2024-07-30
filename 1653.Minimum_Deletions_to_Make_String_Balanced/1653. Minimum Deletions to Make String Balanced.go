package main

func minimumDeletions(s string) int {
	n := len(s)
	cntA, cntB := make([]int, n), make([]int, n) // cntA[i] 维护 s[i] 右侧有多少个 'a'，cntB[i] 维护 s[i] 左侧有多少个 'b'
	for i, c := range s[:n-1] {
		if c == 'b' {
			cntB[i+1] = cntB[i] + 1
		} else {
			cntB[i+1] = cntB[i]
		}
	}
	for i := n - 1; i > 0; i-- {
		if s[i] == 'a' {
			cntA[i-1] = cntA[i] + 1
		} else {
			cntA[i-1] = cntA[i]
		}
	}
	ans := n
	for i := range s {
		ans = min(ans, cntA[i]+cntB[i])
	}
	return ans
}
