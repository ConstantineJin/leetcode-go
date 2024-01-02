package main

import "sort"

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var n = len(nums)
	for i, num := range nums[:n-2] {
		if i > 0 && num == nums[i-1] { // 跳过重复数字
			continue
		}
		if num+nums[i+1]+nums[i+2] > 0 { // 连续三个数之和已经大于0，不可能再有解了
			break
		}
		if num+nums[n-1]+nums[n-2] < 0 { // 当前数与最大的两个数之和仍然为负，当前数无解
			continue
		}
		var j, k = i + 1, n - 1 // 双指针，初始指向当前数的下一个数和最大数，分别向右和向左单向移动
		for j < k {             // 指针不相遇
			var sum = num + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else { // 和为0
				ans = append(ans, []int{num, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ { // 跳过重复数字
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- { // 跳过重复数字
				}
			}
		}
	}
	return
}
