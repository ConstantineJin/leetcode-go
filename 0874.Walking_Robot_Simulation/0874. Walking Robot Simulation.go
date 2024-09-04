package main

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func robotSim(commands []int, obstacles [][]int) (ans int) {
	mp := make(map[[2]int]struct{})
	for _, obstacle := range obstacles {
		mp[[2]int{obstacle[0], obstacle[1]}] = struct{}{}
	}
	var x, y, heading int // 朝向北/东/南/西对应的 heading 值分别为 0/1/2/3
	for _, command := range commands {
		switch command {
		case -2: // 左转 90°
			heading = (heading + 3) % 4
		case -1: // 右转 90°
			heading = (heading + 1) % 4
		default:
			for ; command > 0; command-- {
				x += dirs[heading][0]
				y += dirs[heading][1]
				if _, ok := mp[[2]int{x, y}]; ok { // 遇到障碍物，需要停在障碍物前那个格子
					x -= dirs[heading][0]
					y -= dirs[heading][1]
					break
				}
				ans = max(ans, x*x+y*y)
			}
		}
	}
	return
}
