package main

type EST struct {
}

func vowelStrings(words []string, left int, right int) (cnt int) {
	set := map[uint8]EST{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}} // map加空结构体模拟set
	for i := left; i <= right; i++ {
		_, first := set[words[i][0]]
		_, last := set[words[i][len(words[i])-1]]
		if first && last {
			cnt++
		}
	}
	return
}
