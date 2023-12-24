package main

type EST struct{}

func isPathCrossing(path string) bool {
	var set = map[[2]int]EST{[2]int{0, 0}: {}}
	var pos [2]int
	for _, p := range path {
		switch p {
		case 'N':
			pos[0]++
		case 'S':
			pos[0]--
		case 'E':
			pos[1]++
		case 'W':
			pos[1]--
		}
		if _, ok := set[pos]; ok {
			return true
		}
		set[pos] = EST{}
	}
	return false
}
