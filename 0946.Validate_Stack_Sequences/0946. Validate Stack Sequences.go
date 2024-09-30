package main

func validateStackSequences(pushed, popped []int) bool {
	st := make([]int, 0, len(pushed))
	var i int
	for _, v := range pushed {
		st = append(st, v)
		for len(st) > 0 && st[len(st)-1] == popped[i] {
			st = st[:len(st)-1]
			i++
		}
	}
	return len(st) == 0
}
