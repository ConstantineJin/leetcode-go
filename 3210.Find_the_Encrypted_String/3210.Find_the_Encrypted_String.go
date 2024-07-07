package main

func getEncryptedString(s string, k int) string {
	k %= len(s)
	return s[k:] + s[:k]
}
