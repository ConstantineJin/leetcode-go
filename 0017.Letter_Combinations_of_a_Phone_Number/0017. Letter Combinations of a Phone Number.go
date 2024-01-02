package main

func letterCombinations(digits string) (ans []string) {
	var n = len(digits)
	if n == 0 {
		return
	}
	var arr = [8]string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"} // 数字与字符的对应关系
	var dfs func(int, string)
	dfs = func(i int, nums string) {
		if i == n-1 { // 到达最后一个字符
			for _, ch := range arr[digits[i]-'2'] {
				ans = append(ans, nums+string(ch)) // 保存到结果数组
			}
			return
		}
		for _, ch := range arr[digits[i]-'2'] { // 遍历当前数对应的每一个字母
			nums += string(ch)
			dfs(i+1, nums)            // 进行dfs
			nums = nums[:len(nums)-1] // 退出
		}
	}
	dfs(0, "")
	return
}
