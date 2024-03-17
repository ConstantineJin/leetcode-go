package main

type NumArray struct{ prefixSums []int }

func Constructor(nums []int) NumArray {
	var prefixSums = make([]int, len(nums)+1)
	for i, num := range nums {
		prefixSums[i+1] = prefixSums[i] + num
	}
	return NumArray{prefixSums: prefixSums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSums[right+1] - this.prefixSums[left]
}
