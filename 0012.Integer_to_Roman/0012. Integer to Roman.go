package main

func intToRoman(num int) (ans string) {
	i := 0 // 记录位数
	for num > 0 {
		x := num % 10
		switch i {
		case 0:
			switch x {
			case 1:
				ans = "I"
			case 2:
				ans = "II"
			case 3:
				ans = "III"
			case 4:
				ans = "IV"
			case 5:
				ans = "V"
			case 6:
				ans = "VI"
			case 7:
				ans = "VII"
			case 8:
				ans = "VIII"
			case 9:
				ans = "IX"
			}
		case 1:
			switch x {
			case 1:
				ans = "X" + ans
			case 2:
				ans = "XX" + ans
			case 3:
				ans = "XXX" + ans
			case 4:
				ans = "XL" + ans
			case 5:
				ans = "L" + ans
			case 6:
				ans = "LX" + ans
			case 7:
				ans = "LXX" + ans
			case 8:
				ans = "LXXX" + ans
			case 9:
				ans = "XC" + ans
			}
		case 2:
			switch x {
			case 1:
				ans = "C" + ans
			case 2:
				ans = "CC" + ans
			case 3:
				ans = "CCC" + ans
			case 4:
				ans = "CD" + ans
			case 5:
				ans = "D" + ans
			case 6:
				ans = "DC" + ans
			case 7:
				ans = "DCC" + ans
			case 8:
				ans = "DCCC" + ans
			case 9:
				ans = "CM" + ans
			}
		case 3:
			switch x {
			case 1:
				ans = "M" + ans
			case 2:
				ans = "MM" + ans
			case 3:
				ans = "MMM" + ans
			}
		}
		num /= 10
		i++
	}
	return
}
