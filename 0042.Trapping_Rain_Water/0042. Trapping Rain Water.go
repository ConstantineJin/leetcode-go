package main

// 对于下标i，下雨后水能到达的最大高度等于下标i两边的最大高度的最小值
//func trap(height []int) (res int) {
//	n := len(height)
//	capacity := make([]int, n) // 每个柱子在正向遍历时的容量
//	m := 0                     // 初始遍历所得的柱子高度最大值
//	for i, h := range height { // 正向遍历，假设最后有一根无限高的柱子
//		if h > m { // 遇到更高的柱子就更新
//			m = h
//			continue // 此处容量必然是0，直接进入下一次循环
//		}
//		capacity[i] = m - h // 此处容量为先前遍历所得的最大值与此处柱子高度之差
//	}
//	m = 0                         // 重置最大值
//	for i := n - 1; i >= 0; i-- { // 反向遍历，假设开头有一根无限高的柱子
//		if height[i] > m {
//			m = height[i]
//			continue
//		}
//		res += min(capacity[i], m-height[i]) // 每一处的实际容量是两次遍历的最小值，将其加入结果中
//	}
//	return
//}

// 双指针法，空间复杂度从O(n)降至O(1)
func trap(height []int) (res int) {
	left, right, leftMax, rightMax := 0, len(height)-1, height[0], height[len(height)-1] // left和right维护两个从两端开始向中间遍历的指针，leftMax和rightMax记录正向与反向遍历所得的最大值
	for left < right {                                                                   // 左右指针重叠代表完成遍历，退出循环
		if height[left] < height[right] { // left指针所指出高度更低
			res += leftMax - height[left]        // height[left]<height[right]<=rightMax，即右侧一定有更高的柱子，因此此处容量为leftMax与此处柱子高度之差
			left++                               // left指针后移
			leftMax = max(leftMax, height[left]) // 如有必要更新leftMax
		} else { // 同上
			res += rightMax - height[right]
			right--
			rightMax = max(rightMax, height[right])
		}
	}
	return
}
