package main

func mostCompetitive(nums []int, k int) (ans []int) {
	n := len(nums)
	for i, num := range nums {
		for len(ans) > 0 && num < ans[len(ans)-1] && k-len(ans) < n-i {
			ans = ans[:len(ans)-1]
		}
		if len(ans) < k {
			ans = append(ans, num)
		}
	}
	return ans
}
