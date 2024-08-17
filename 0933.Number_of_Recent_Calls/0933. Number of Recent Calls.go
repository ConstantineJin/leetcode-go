package main

type RecentCounter []int

func Constructor() (_ RecentCounter) { return }

func (rc *RecentCounter) Ping(t int) int {
	for len(*rc) > 0 && (*rc)[0] < t-3000 {
		*rc = (*rc)[1:]
	}
	*rc = append(*rc, t)
	return len(*rc)
}
