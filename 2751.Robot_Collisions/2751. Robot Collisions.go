package main

import "slices"

type Robot struct{ id, position, health, direction int }

func survivedRobotsHealths(positions, healths []int, directions string) []int {
	n := len(positions)
	robots := make([]Robot, n)
	for i, direction := range directions {
		if direction == 'L' {
			robots[i] = Robot{i, positions[i], healths[i], -1}
		} else {
			robots[i] = Robot{i, positions[i], healths[i], 1}
		}
	}
	slices.SortFunc(robots, func(a, b Robot) int { return a.position - b.position })
	var left, right []Robot // 用列表 left 维护向左的机器人，用栈 right 维护向右的机器人
next:
	for _, robot := range robots {
		if robot.direction == 1 { // 向右的机器人
			right = append(right, robot)
		} else {
			for len(right) > 0 {
				top := &right[len(right)-1]
				if top.health > robot.health {
					top.health--
					continue next
				} else if top.health == robot.health {
					right = right[:len(right)-1]
					continue next
				} else {
					robot.health--
					right = right[:len(right)-1]
				}
			}
			left = append(left, robot)
		}
	}
	left = append(left, right...)
	slices.SortFunc(left, func(a, b Robot) int { return a.id - b.id })
	ans := make([]int, len(left))
	for i, robot := range left {
		ans[i] = robot.health
	}
	return ans
}
