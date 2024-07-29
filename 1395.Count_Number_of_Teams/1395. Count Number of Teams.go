package main

func numTeams(rating []int) (ans int) {
	n := len(rating)
	for j, rate := range rating {
		var leftLess, leftMore, rightLess, rightMore int
		for i := 0; i < j; i++ {
			if rating[i] < rate {
				leftLess++
			} else if rating[i] > rate {
				leftMore++
			}
		}
		for k := j + 1; k < n; k++ {
			if rating[k] < rate {
				rightLess++
			} else if rating[k] > rate {
				rightMore++
			}
		}
		ans += leftLess*rightMore + leftMore*rightLess
	}
	return
}
