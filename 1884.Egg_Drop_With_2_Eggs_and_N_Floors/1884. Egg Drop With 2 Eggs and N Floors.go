package main

import "math"

// O(1) 数学方法。设答案为 x，则转化为求满足 1+2+3+……+x = (1+x)*x/2 >= n 的最小 x
func twoEggDrop(n int) int { return int(math.Ceil((-1 + math.Sqrt(float64(1+8*n))) / 2)) }
