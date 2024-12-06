package main

func maxCount(banned []int, n, maxSum int) (ans int) {
	set := make(map[int]struct{}, len(banned))
	for _, v := range banned {
		set[v] = struct{}{}
	}
	for i := 1; i <= n && i <= maxSum; i++ {
		if _, ok := set[i]; !ok {
			maxSum -= i
			ans++
		}
	}
	return
}
