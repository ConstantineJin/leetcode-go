package main

// 类似142.环形链表II，使用快慢指针寻找链表环入口
func findDuplicate(nums []int) int {
	var slow, fast int
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
