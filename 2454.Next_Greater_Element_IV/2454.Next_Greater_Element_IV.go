package main

func secondGreaterElement(nums []int) []int {
	var ans = make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	var s, t = make([]int, 0), make([]int, 0)
	for i, num := range nums {
		for len(t) > 0 && nums[t[len(t)-1]] < num {
			ans[t[len(t)-1]] = num
			t = t[:len(t)-1]
		}
		var j = len(s) - 1
		for j >= 0 && nums[s[j]] < num {
			j--
		}
		t = append(t, s[j+1:]...)
		s = append(s[:j+1], i)
	}
	return ans
}
