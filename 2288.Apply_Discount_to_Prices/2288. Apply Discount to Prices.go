package main

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	d := 1 - float64(discount)/100
	for i, word := range words {
		if len(word) > 1 && word[0] == '$' {
			if price, err := strconv.Atoi(word[1:]); err == nil {
				words[i] = fmt.Sprintf("$%.2f", float64(price)*d)
			}
		}
	}
	return strings.Join(words, " ")
}
