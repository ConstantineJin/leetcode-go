package main

import "strconv"

func fractionAddition(expression string) string {
	denominator := 1                 // 通分后的分母，即所有分母的最小公倍数
	positive := expression[0] != '-' // 记录当前项是否为正数
	expression += "+"                // 结尾添加哨兵，无需在循环外额外处理最后一项
	var numerator int                // 通分后的分子
	var num []rune
	var n int // 当前项的分子
	for i, c := range expression {
		switch c {
		case '+', '-':
			if i == 0 {
				continue
			}
			d, _ := strconv.Atoi(string(num)) // 当前项的分母
			num = num[:0]
			newDenominator := lcm(denominator, d) // 通分
			numerator *= newDenominator / denominator
			denominator = newDenominator
			if positive {
				numerator += n * newDenominator / d
			} else {
				numerator -= n * newDenominator / d
			}
			positive = c == '+'
		case '/':
			n, _ = strconv.Atoi(string(num))
			num = num[:0]
		default:
			num = append(num, c)
		}
	}
	gcd := gcd(numerator, denominator)
	numerator /= gcd
	denominator /= gcd
	if denominator < 0 { // 如果分母为负，则需要将负号移到分子上
		numerator, denominator = -numerator, -denominator
	}
	return strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
