package main

// 库函数
//func dayOfTheWeek(day, month, year int) string {
//	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
//	return t.Weekday().String()
//}

// 蔡勒公式
func dayOfTheWeek(d int, m int, y int) string {
	if m < 3 {
		m += 12
		y -= 1
	}
	var c = y / 100
	y %= 100
	var w = (c/4 - 2*c + y + y/4 + 13*(m+1)/5 + d - 1) % 7
	var weeks = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return weeks[(w+7)%7]
}
