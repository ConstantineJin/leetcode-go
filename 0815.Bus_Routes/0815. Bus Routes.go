package main

func numBusesToDestination(routes [][]int, source, target int) int {
	if source == target {
		return 0
	}
	mp := make(map[int][]int) // key 为站号，value 为经停该站的线路号
	for i, route := range routes {
		for _, stop := range route {
			mp[stop] = append(mp[stop], i)
		}
	}
	visitedStop, visitedRoute := make(map[int]struct{}), make(map[int]struct{})
	for ans, cur := 1, []int{source}; len(cur) > 0; ans++ {
		var nxt []int
		for _, stop := range cur {
			for _, route := range mp[stop] {
				if _, ok := visitedRoute[route]; ok {
					continue
				}
				visitedRoute[route] = struct{}{}
				for _, nextStop := range routes[route] {
					if nextStop == target {
						return ans
					}
					if _, ok := visitedStop[nextStop]; !ok {
						visitedStop[nextStop] = struct{}{}
						nxt = append(nxt, nextStop)
					}
				}
			}
		}
		cur = nxt
	}
	return -1
}
