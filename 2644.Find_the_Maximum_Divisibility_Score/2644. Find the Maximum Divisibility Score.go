package main

func maxDivScore(nums, divisors []int) (ans int) {
	maxCnt := -1
	for _, divisor := range divisors {
		cnt := 0
		for _, num := range nums {
			if num%divisor == 0 {
				cnt++
			}
		}
		if cnt > maxCnt || cnt == maxCnt && divisor < ans {
			maxCnt, ans = cnt, divisor
		}
	}
	return
}
