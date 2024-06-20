package main

func countBeautifulPairs(nums []int) (ans int) {
	var cnt [10]int
	for _, x := range nums {
		for y := 1; y < 10; y++ {
			if cnt[y] > 0 && gcd(x%10, y) == 1 {
				ans += cnt[y]
			}
		}
		for x >= 10 {
			x /= 10
		}
		cnt[x]++
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
