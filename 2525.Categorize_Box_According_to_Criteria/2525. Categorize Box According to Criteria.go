package main

func categorizeBox(length int, width int, height int, mass int) string {
	var bulky, heavy = false, false
	if length >= 1e4 || width >= 1e4 || height >= 1e4 || length*width*height >= 1e9 {
		bulky = true
	}
	if mass >= 100 {
		heavy = true
	}
	if bulky && heavy {
		return "Both"
	} else if bulky {
		return "Bulky"
	} else if heavy {
		return "Heavy"
	} else {
		return "Neither"
	}
}
