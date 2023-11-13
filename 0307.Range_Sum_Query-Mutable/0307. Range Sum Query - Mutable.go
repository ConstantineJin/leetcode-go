package main

import "fmt"

// 方法1：树状数组
type NumArray struct {
	nums, tree []int
}

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
func (this *NumArray) Update(index, val int) {
	delta := val - this.nums[index]
	this.nums[index] = val
	for i := index + 1; i < len(this.tree); i += i & -i {
		this.tree[i] += delta
	}
}

// prefixSum 求前缀和
func (this *NumArray) prefixSum(i int) (s int) {
	for ; i > 0; i &= i - 1 { // i -= i & -i 的另一种写法
		s += this.tree[i]
	}
	return
}

// SumRange 返回[left, right]闭区间内的元素和
func (this *NumArray) SumRange(left, right int) (sum int) {
	return this.prefixSum(right+1) - this.prefixSum(left) // right+1的前缀和与left的前缀和之差即为所求
}

func main() {
	numArray := Constructor([]int{1, 3, 5})
	fmt.Println(numArray.SumRange(0, 2))
	numArray.Update(1, 2)
	fmt.Println(numArray.SumRange(0, 2))
}
