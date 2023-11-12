package main

import (
	"fmt"
	"strconv"
)

func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 1<<n) // 含有n个元素的集合幂集数量为2^n个
	for i := range res {
		bin := fmt.Sprintf("%0*s", n, strconv.FormatInt(int64(i), 2)) // 求得当前第i个子集的二进制表示，不足n位的补0
		for j, v := range bin {
			if v == '1' { // 如果第j位为1，则第i个子集包含nums[j]
				res[i] = append(res[i], nums[j])
			}
		}
	}
	return res
}
