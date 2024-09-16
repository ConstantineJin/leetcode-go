package main

func getSneakyNumbers(nums []int) (ans []int) {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := set[num]; ok {
			ans = append(ans, num)
		}
		set[num] = struct{}{}
	}
	return
}
