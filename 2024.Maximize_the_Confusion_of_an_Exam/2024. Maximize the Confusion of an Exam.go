package main

func maxConsecutiveAnswers(answerKey string, k int) (ans int) {
	var cntT, cntF, left int
	for right, key := range answerKey { // 滑动窗口
		if key == 'T' {
			cntT++
		} else {
			cntF++
		}
		for min(cntT, cntF) > k {
			if answerKey[left] == 'T' {
				cntT--
			} else {
				cntF--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
