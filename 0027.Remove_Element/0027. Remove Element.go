package main

func removeElement(nums []int, val int) (cnt int) {
	for i := 0; i < len(nums); i++ { // i为快指针，cnt为慢指针
		if nums[i] != val { // 快指针当前指向的元素不需要删除
			nums[cnt] = nums[i] // 将快指针指向元素的值赋予慢指针当前指向元素
			cnt++               // 结果数组长度+1
		}
	}
	nums = nums[:cnt] // 只保留慢指针之前的数组
	return
}
