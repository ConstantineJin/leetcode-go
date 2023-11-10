package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)                               // 对药水的能量强度进行升序排序
	spellCnt, potionCnt := len(spells), len(potions) // 咒语的数量和药水的数量
	res := make([]int, spellCnt)
	for i := 0; i < spellCnt; i++ {
		res[i] = potionCnt - sort.SearchInts(potions, (int(success)+spells[i]-1)/spells[i]) // 先求得每句咒语所需药水强度的最小值，在已排序的咒语数组中二分查找到对应下标，然后用药水数量减去之
	}
	return res
}
