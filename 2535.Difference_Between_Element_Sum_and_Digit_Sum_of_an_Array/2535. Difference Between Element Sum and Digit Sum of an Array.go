package main

func differenceOfSum(nums []int) (ans int) {
	for _, num := range nums {
		ans += num
		for num > 0 {
			ans -= num % 10
			num /= 10
		}
	}
	return
}
