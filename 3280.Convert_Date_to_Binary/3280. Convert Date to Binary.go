package main

import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	parts := strings.Split(date, "-")
	for i := range parts {
		x, _ := strconv.Atoi(parts[i])
		parts[i] = strconv.FormatUint(uint64(x), 2)
	}
	return strings.Join(parts, "-")
}
