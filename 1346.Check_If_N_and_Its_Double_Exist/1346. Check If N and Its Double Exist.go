package main

func checkIfExist(arr []int) bool {
	set := make(map[int]struct{}, len(arr))
	for _, v := range arr {
		if _, ok := set[v*2]; ok {
			return true
		}
		if v%2 == 0 {
			if _, ok := set[v/2]; ok {
				return true
			}
		}
		set[v] = struct{}{}
	}
	return false
}
