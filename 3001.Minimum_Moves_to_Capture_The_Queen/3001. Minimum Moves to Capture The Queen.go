package main

func isBetween(l, m, r int) bool { return min(l, r) < m && m < max(l, r) }

func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	if a == e && (c != e || !isBetween(b, d, f)) || b == f && (d != f || !isBetween(a, c, e)) || c+d == e+f && (a+b != e+f || !isBetween(c, a, e)) || c-d == e-f && (a-b != e-f || !isBetween(c, a, e)) {
		return 1
	}
	return 2
}
