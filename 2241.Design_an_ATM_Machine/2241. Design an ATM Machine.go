package main

var denominations = [5]int{20, 50, 100, 200, 500}

type ATM [5]int

func Constructor() ATM { return ATM{} }

func (a *ATM) Deposit(banknotesCount []int) {
	for i, count := range banknotesCount {
		a[i] += count
	}
}

func (a *ATM) Withdraw(amount int) []int {
	if amount%10 > 0 {
		return []int{-1}
	}
	ans := make([]int, 5)
	for i := 4; i >= 0; i-- {
		ans[i] = min(amount/denominations[i], a[i])
		amount -= ans[i] * denominations[i]
	}
	if amount > 0 {
		return []int{-1}
	}
	for i, count := range ans {
		a[i] -= count
	}
	return ans
}
