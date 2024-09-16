package main

func distanceBetweenBusStops(distance []int, start, destination int) int {
	n := len(distance)
	start, destination = min(start, destination), max(start, destination)
	var d1, d2 int
	for i := start; i < destination; i = (i + 1) % n {
		d1 += distance[i]
	}
	for i := start - 1; i >= 0; i-- {
		d2 += distance[i]
	}
	for i := n - 1; i >= destination; i-- {
		d2 += distance[i]
	}
	return min(d1, d2)
}
