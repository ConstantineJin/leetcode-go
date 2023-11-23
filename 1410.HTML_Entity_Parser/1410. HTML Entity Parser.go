package main

func entityParser(text string) string {
	start, ans := -1, make([]byte, 0)
	mp := map[string]byte{"quot": '"', "apos": '\'', "amp": '&', "gt": '>', "lt": '<', "frasl": '/'} // mp的key为HTML代码中的字符实体，value为字符本身
	for i, ch := range text {
		ans = append(ans, byte(ch))
		switch ch {
		case '&': // 如果当前字符是'&'，那么当前字符可能是一个字符实体的开始
			start = i // 记录下标
		case ';': // 如果当前字符是';'，那么当前字符可能是一个字符实体的结束
			if v, ok := mp[text[start+1:i]]; start >= 0 && ok { // 在mp中寻找从start到当前字符是否存在对应的字符实体
				ans = append(ans[:len(ans)-(i-start)-1], v) // 如果存在，将其替换为字符本身
			}
		}
	}
	return string(ans)
}
