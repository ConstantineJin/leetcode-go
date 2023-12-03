package main

func carPooling(trips [][]int, capacity int) bool {
	diff, count := [1001]int{}, 0 // trips长度最大为1000，设diff数组记录每个位置人数增减情况；count记录当前车上人数
	for _, trip := range trips {  // 根据trips数组初始化diff数组
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
	}
	for _, p := range diff { // 计算每个位置车上人数
		count += p
		if count > capacity { // 超载就返回false
			return false
		}
	}
	return true
}
