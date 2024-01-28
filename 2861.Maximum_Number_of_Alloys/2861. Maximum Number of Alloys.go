package main

import (
	"slices"
	"sort"
)

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) (ans int) {
	var mx = slices.Min(stock) + budget // 二分查找的上界，stock的最小值再加上预算数（假设composition[i]和cost[i]都是1）
	for _, c := range composition {     // 逐个判断每台机器所能生产的最大合金数
		ans += sort.Search(mx-ans, func(i int) bool { // 二分查找的下限是ans（低于ans的结果没有意义）
			i += ans + 1              // i代表制造合金的份数
			var money int             // 采购花费
			for j, s := range stock { // 遍历库存金属
				if s < c[j]*i { // 如果j类金属库存不足
					money += (c[j]*i - s) * cost[j] // 更新采购花费
					if money > budget {             // 如果采购花费超过预算
						return true
					}
				}
			}
			return false
		})
	}
	return
}
