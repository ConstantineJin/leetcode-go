package main

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) (ans []string) {
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return ans
}
