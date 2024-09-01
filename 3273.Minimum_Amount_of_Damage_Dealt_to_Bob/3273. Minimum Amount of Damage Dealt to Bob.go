package main

import "sort"

type Enemy struct{ damage, time int }

func minDamage(power int, damage, health []int) int64 {
	n := len(damage)
	var damagePerSecond int
	for _, d := range damage {
		damagePerSecond += d
	}
	enemies := make([]Enemy, n)
	for i, d := range damage {
		enemies[i] = Enemy{d, (health[i] + power - 1) / power}
	}
	sort.Slice(enemies, func(i, j int) bool { return enemies[i].damage*enemies[j].time > enemies[j].damage*enemies[i].time })
	var ans int
	for _, enemy := range enemies {
		ans += damagePerSecond * enemy.time
		damagePerSecond -= enemy.damage
	}
	return int64(ans)
}
