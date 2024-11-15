package main

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 {
		return 0
	}
	ans := right
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		for right < n && arr[right] < arr[left] {
			right++
		}
		ans = min(ans, right-left-1) // 开区间 (left,right) 是要删除的子数组
	}
	return ans
}
