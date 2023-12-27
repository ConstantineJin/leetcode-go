package main

func sortColors(nums []int) {
	var zero, one int
	for _, num := range nums {
		switch num {
		case 0:
			zero++
		case 1:
			one++
		}
	}
	for i := 0; i < zero; i++ {
		nums[i] = 0
	}
	for i := zero; i < zero+one; i++ {
		nums[i] = 1
	}
	for i := zero + one; i < len(nums); i++ {
		nums[i] = 2
	}
}
