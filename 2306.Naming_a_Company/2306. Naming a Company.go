package main

func distinctNames(ideas []string) (ans int64) {
	groups := [26]map[string]bool{}
	for i := range groups {
		groups[i] = make(map[string]bool)
	}
	for _, idea := range ideas {
		groups[idea[0]-'a'][idea[1:]] = true
	}
	for i, a := range groups {
		for _, b := range groups[:i] {
			var m int // 交集的大小
			for s := range a {
				if b[s] {
					m++
				}
			}
			ans += int64((len(a) - m) * (len(b) - m))
		}
	}
	return ans * 2
}
