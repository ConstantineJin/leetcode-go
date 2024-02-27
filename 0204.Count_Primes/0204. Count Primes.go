package main

// 埃氏筛 Sieve of Eratosthenes
//func countPrimes(n int) (ans int) {
//	var isPrime = make([]bool, n)
//	for i := range isPrime {
//		isPrime[i] = true
//	}
//	for i := 2; i < n; i++ {
//		if isPrime[i] {
//			ans++
//			for j := i * 2; j < n; j += i {
//				isPrime[j] = false
//			}
//		}
//	}
//	return
//}

// 线性筛
func countPrimes(n int) int {
	var primes []int
	var isPrime = make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, prime := range primes {
			if i*prime >= n {
				break
			}
			isPrime[i*prime] = false
			if i%prime == 0 {
				break
			}
		}
	}
	return len(primes)
}
