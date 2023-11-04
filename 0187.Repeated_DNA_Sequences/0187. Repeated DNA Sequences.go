package main

// 暴力
//func findRepeatedDnaSequences(s string) (res []string) {
//	mp := make(map[string]int)
//	for i := 0; i < len(s)-9; i++ {
//		mp[s[i:i+10]]++
//	}
//	for i, v := range mp {
//		if v > 1 {
//			res = append(res, i)
//		}
//	}
//	return
//}

func findRepeatedDnaSequences(s string) (res []string) {
	if len(s) < 10 {
		return
	}
	var bin = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	code, cnt := 0, make(map[int]int)
	for _, char := range s[:9] {
		code = code<<2 | bin[byte(char)]
	}
	for i := 0; i <= len(s)-10; i++ {
		code = (code<<2 | bin[s[i+9]]) & 1048575 //1<<(20)-1 = 1048575
		cnt[code]++
		if cnt[code] == 2 {
			res = append(res, s[i:i+10])
		}
	}
	return
}
