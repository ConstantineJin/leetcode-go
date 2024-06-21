package main

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	var temp int // 滑动窗口，老板使自己不生气的区间中，原本不满意的顾客数量
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			temp += customers[i]
		}
	}
	mx := temp
	for i := 0; i < n-minutes; i++ {
		if grumpy[i] == 1 {
			temp -= customers[i]
		}
		if grumpy[i+minutes] == 1 {
			temp += customers[i+minutes]
		}
		mx = max(mx, temp)
	}
	var sum int // 原本就满意的客户总数
	for i, customer := range customers {
		if grumpy[i] == 0 {
			sum += customer
		}
	}
	return sum + mx
}
