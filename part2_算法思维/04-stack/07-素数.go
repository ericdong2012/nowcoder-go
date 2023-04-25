package main

import (
	"fmt"
	"math"
)

// 质数又称素数。一个大于1的自然数，如果除了1和它自身外，不能被其他自然数整除的数；
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(n int) []int {
	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	n := 100
	primes := findPrimes(n)
	fmt.Printf("The primes less than or equal to %d are: %v\n", n, primes)
}
