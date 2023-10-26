package main

type EST struct {
}

func countDigits(num int) (cnt int) {
	set := make(map[int]EST)
	for i := 1; i < 10; i++ {
		if num%i == 0 {
			set[i] = EST{}
		}
	}
	for num > 0 {
		if _, ok := set[num%10]; ok {
			cnt++
		}
		num /= 10
	}
	return
}
