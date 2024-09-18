package main

import "sort"

func latestTimeCatchTheBus(buses, passengers []int, capacity int) (ans int) {
	sort.Ints(buses)
	sort.Ints(passengers)
	m, n := len(buses), len(passengers)
	var j, cnt int
	for _, bus := range buses { // 模拟乘客上车的过程
		for cnt = 0; cnt < capacity && j < n && passengers[j] <= bus; cnt++ {
			j++
		}
	}
	if cnt < capacity { // 如果最后一班公交车仍有空位，就在最后一班公交车发车的时候抵达
		ans = buses[m-1]
	} else { // 否则在最后一个上车的乘客到达的时间到达
		ans = passengers[j-1]
	}
	for j--; j >= 0 && ans == passengers[j]; j-- { // 倒序寻找第一个没有其他乘客到达的时刻
		ans--
	}
	return ans
}
