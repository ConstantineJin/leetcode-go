package main

func averageWaitingTime(customers [][]int) float64 {
	n := len(customers)
	var s, t int // s 代表总共等待时间，t 代表当前时刻
	for _, customer := range customers {
		arrival, time := customer[0], customer[1]
		if t > arrival { // 当前顾客需要排队等待
			t += time
			s += t - arrival
		} else { // 当前顾客不需要排队等待
			t = arrival + time
			s += time
		}
	}
	return float64(s) / float64(n)
}
