package main

import "sort"

func numRescueBoats(people []int, limit int) (ans int) {
	sort.Ints(people)        // 贪心，根据体重升序排序
	i, j := 0, len(people)-1 // i和j分别指向剩下的最瘦的人和最胖的人
	for i <= j {
		if people[i]+people[j] <= limit { // 如果当前剩下的最瘦和最胖的人不超载，那么最瘦的人可以上船
			i++
		}
		j--
		ans++
	}
	return
}
