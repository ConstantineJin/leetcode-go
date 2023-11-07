package main

func trap(height []int) (res int) {
	n := len(height)
	capacity := make([]int, n) // 每个柱子在正向遍历时的容量
	m := 0                     // 初始遍历所得的柱子高度最大值
	for i, h := range height { // 正向遍历，假设最后有一根无限高的柱子
		if h > m { // 遇到更高的柱子就更新
			m = h
			continue // 此处容量必然是0，直接进入下一次循环
		}
		capacity[i] = m - h // 此处容量为先前遍历所得的最大值与此处柱子高度之差
	}
	m = 0                         // 重置最大值
	for i := n - 1; i >= 0; i-- { // 反向遍历，假设开头有一根无限高的柱子
		if height[i] > m {
			m = height[i]
			continue
		}
		res += min(capacity[i], m-height[i]) // 每一处的实际容量是两次遍历的最小值，将其加入结果中
	}
	return
}
