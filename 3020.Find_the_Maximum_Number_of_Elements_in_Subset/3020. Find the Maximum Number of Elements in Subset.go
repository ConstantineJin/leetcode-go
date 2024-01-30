package main

func pow(x, y int) int {
	var ans = 1
	for ; y > 0; y-- {
		ans *= x
	}
	return ans
}

func maximumLength(nums []int) (ans int) {
	var mp = make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	var dfs func(x, k int) // x是底数，k是指数对2的对数
	dfs = func(x, k int) {
		temp := pow(x, pow(2, k))
		switch mp[temp] {
		case 0:
			ans = max(ans, k*2-1)
		case 1:
			ans = max(ans, k*2+1)
		default:
			dfs(x, k+1)
		}
	}
	for _, num := range nums {
		if num != 1 {
			dfs(num, 0)
		}
	}
	if mp[1]%2 == 0 {
		mp[1]--
	}
	return max(mp[1], ans)
}
