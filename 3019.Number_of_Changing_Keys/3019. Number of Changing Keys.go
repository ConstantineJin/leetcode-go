package main

func countKeyChanges(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		if (s[i]-s[i-1])%32 != 0 {
			ans++
		}
	}
	return
}
