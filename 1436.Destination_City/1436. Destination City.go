package main

func destCity(paths [][]string) string {
	var dest = make(map[string]int)
	for _, path := range paths {
		dest[path[0]]--
		dest[path[1]]++
	}
	for k, v := range dest {
		if v == 1 {
			return k
		}
	}
	return ""
}
