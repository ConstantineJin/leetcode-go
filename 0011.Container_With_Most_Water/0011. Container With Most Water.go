package main

func maxArea(height []int) (res int) {
	left, right := 0, len(height)-1                        // 左右双指针
	leftHeight, rightHeight := height[left], height[right] // 记录容器左右侧的高度
	res = (right - left) * min(leftHeight, rightHeight)    // 结果的初始值
	for left < right {                                     // 保证左右指针不交叉不重叠
		if leftHeight < rightHeight { // 始终向内移动较矮一侧的指针
			for left < right && height[left] <= leftHeight { // 直到找到更高的容器壁
				left++
			}
			leftHeight = height[left]
		} else {
			for left < right && height[right] <= rightHeight {
				right--
			}
			rightHeight = height[right]
		}
		res = max(res, (right-left)*min(leftHeight, rightHeight)) // 如果新构成的容器容积更大，则更新结果
	}
	return
}
