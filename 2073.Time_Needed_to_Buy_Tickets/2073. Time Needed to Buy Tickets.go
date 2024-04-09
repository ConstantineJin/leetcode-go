package main

func timeRequiredToBuy(tickets []int, k int) (ans int) {
	for i, ticket := range tickets {
		if i <= k {
			ans += min(ticket, tickets[k])
		} else {
			ans += min(ticket, tickets[k]-1)
		}
	}
	return
}
