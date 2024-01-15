package main

func repeatLimitedString(s string, repeatLimit int) string {
	var cnt = [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++ // 遍历原字符串统计每个字母出现次数
	}
	var ans []byte
	for i := 25; i >= 0; i-- { // 贪心，倒序遍历字母表，尽可能优先添加更大的字母
		for cnt[i] > 0 {
			if cnt[i] <= repeatLimit { // 如果当前字母剩余可用次数不超过连续出现次数限制
				for ; cnt[i] > 0; cnt[i]-- { // 全部添加
					ans = append(ans, byte('a'+i))
				}
				break
			} else { // 否则按照限制次数添加
				for j := 0; j < repeatLimit; j++ {
					ans = append(ans, byte('a'+i))
					cnt[i]--
				}
				var j = i - 1
				for ; j >= 0 && cnt[j] == 0; j-- { // 继续贪心，寻找剩余可用的字母中小于当前字母的最大字母
				}
				if j >= 0 { // 如果存在这样的字母就添加
					ans = append(ans, byte('a'+j))
					cnt[j]--
				} else { // 否则无法再添加字母，直接返回
					return string(ans)
				}
			}
		}
	}
	return string(ans)
}
