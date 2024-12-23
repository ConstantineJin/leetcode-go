package main

import "slices"

type ExamRoom struct {
	n int
	s []int
}

func Constructor(n int) ExamRoom { return ExamRoom{n: n, s: []int{}} }

func (this *ExamRoom) Seat() int {
	var student, idx int
	if len(this.s) > 0 {
		dist := this.s[0]
		pre := -1
		for i, v := range this.s {
			if pre != -1 {
				d := (v - pre) / 2
				if d > dist {
					dist = d
					student = pre + d
					idx = i
				}
			}
			pre = v
		}
		if this.n-1-this.s[len(this.s)-1] > dist {
			student = this.n - 1
			idx = len(this.s)
		}
	}
	this.s = slices.Insert(this.s, idx, student)
	return student
}

func (this *ExamRoom) Leave(p int) {
	var idx int
	for i, v := range this.s {
		if v == p {
			idx = i
			break
		}
	}
	this.s = slices.Delete(this.s, idx, idx+1)
}
