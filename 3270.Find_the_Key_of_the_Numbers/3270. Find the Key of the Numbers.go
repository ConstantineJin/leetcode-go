package main

import (
	"fmt"
	"strconv"
)

func generateKey(num1, num2, num3 int) int {
	s1, s2, s3 := fmt.Sprintf("%04d", num1), fmt.Sprintf("%04d", num2), fmt.Sprintf("%04d", num3)
	var ans string
	for i := 0; i < 4; i++ {
		ans += string(min(s1[i], s2[i], s3[i]))
	}
	key, _ := strconv.Atoi(ans)
	return key
}
