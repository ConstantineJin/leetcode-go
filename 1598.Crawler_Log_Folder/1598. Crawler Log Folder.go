package main

func minOperations(logs []string) (ans int) {
	for _, log := range logs {
		switch log {
		case "./":
		case "../":
			ans = max(ans-1, 0)
		default:
			ans++
		}
	}
	return
}
