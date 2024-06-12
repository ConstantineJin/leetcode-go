package main

// NumArray 树状数组
type NumArray struct{ nums, tree []int }

// Constructor 构造NumArray
func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	for i, num := range nums {
		i++
		tree[i] += num
		if nxt := i + i&-i; nxt < len(tree) {
			tree[nxt] += tree[i]
		}
	}
	return NumArray{nums: nums, tree: tree}
}

// Update 将下标为index的元素的值更新为val
func (a *NumArray) Update(index, val int) {
	delta := val - a.nums[index]
	a.nums[index] = val
	for i := index + 1; i < len(a.tree); i += i & -i {
		a.tree[i] += delta
	}
}

// prefixSum 求前缀和
func (a *NumArray) prefixSum(i int) (s int) {
	for ; i > 0; i &= i - 1 { // i -= i & -i 的另一种写法
		s += a.tree[i]
	}
	return
}

// SumRange 返回[left, right]闭区间内的元素和
func (a *NumArray) SumRange(left, right int) (sum int) {
	return a.prefixSum(right+1) - a.prefixSum(left) // right+1的前缀和与left的前缀和之差即为所求
}
