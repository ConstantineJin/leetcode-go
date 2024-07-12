package main

func countLargestGroup(n int) (ans int) {
	mp := make(map[int]int)
	for i := 1; i <= n; i++ {
		x, sum := i, 0
		for x > 0 {
			sum += x % 10
			x /= 10
		}
		mp[sum]++
	}
	var mx int
	for _, v := range mp {
		if v > mx {
			mx, ans = v, 1
		} else if v == mx {
			ans++
		}
	}
	return
}
