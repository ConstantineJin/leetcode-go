package main

func findMatrix(nums []int) [][]int {
	var cnt = [201]int{}
	var mx int
	for _, num := range nums {
		cnt[num]++
		mx = max(mx, cnt[num])
	}
	var ans = make([][]int, mx)
	for i, v := range cnt {
		if v == 0 {
			continue
		}
		for j := 0; j < v; j++ {
			ans[j] = append(ans[j], i)
		}
	}
	return ans
}
