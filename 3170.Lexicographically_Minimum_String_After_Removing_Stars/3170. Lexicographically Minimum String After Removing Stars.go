package main

func clearStars(s string) string {
	var indices [26][]int
	str := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			str[i] = '#'
			for j, index := range indices {
				if len(index) > 0 {
					str[index[len(index)-1]] = '#'
					indices[j] = indices[j][:len(index)-1]
					break
				}
			}
		} else {
			str[i] = s[i]
			indices[s[i]-'a'] = append(indices[s[i]-'a'], i)
		}
	}
	var ans []byte
	for i, c := range str {
		if c != '#' {
			ans = append(ans, str[i])
		}
	}
	return string(ans)
}
