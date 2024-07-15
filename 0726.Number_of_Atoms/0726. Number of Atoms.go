package main

import (
	"sort"
	"strconv"
)

func countOfAtoms(s string) (ans string) {
	stack := []map[string]int{{}}
	i, n := 0, len(s)
	parseNum := func() (v int) {
		if i == n || !isDigit(s[i]) {
			return 1
		}
		for ; i < n && isDigit(s[i]); i++ {
			v = v*10 + int(s[i]-'0')
		}
		return
	}
	for i < n {
		switch s[i] {
		case '(':
			stack = append(stack, map[string]int{})
			i++
		case ')':
			mp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++
			v := parseNum()
			for s, c := range mp {
				stack[len(stack)-1][s] += c * v
			}
		default:
			st := i
			for i++; i < n && isLower(s[i]); i++ {
			}
			name := s[st:i]
			stack[len(stack)-1][name] += parseNum()
		}
	}
	mp := stack[0]
	atoms := make([]string, 0, len(mp))
	for k := range mp {
		atoms = append(atoms, k)
	}
	sort.Strings(atoms)
	for _, atom := range atoms {
		ans += atom
		if v := mp[atom]; v > 1 {
			ans += strconv.Itoa(v)
		}
	}
	return
}

func isDigit(b byte) bool { return '0' <= b && b <= '9' }
func isLower(b byte) bool { return 'a' <= b && b <= 'z' }
