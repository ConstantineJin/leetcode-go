package main

func calcPoints(points []int) (point int) {
	var idx = -1
	for i, p := range points {
		point += p
		if idx >= 0 && i-idx <= 2 {
			point += p
		}
		if p == 10 {
			idx = i
		}
	}
	return
}

func isWinner(player1 []int, player2 []int) int {
	var point1, point2 = calcPoints(player1), calcPoints(player2)
	if point1 > point2 {
		return 1
	} else if point1 < point2 {
		return 2
	} else {
		return 0
	}
}
