package sieve

import (
	"math"
)

type Sieve interface {
	NthPrime(n int64) int64
}

type sieveOfEratosthenes struct {
}

func NewSieve() Sieve {
	return &sieveOfEratosthenes{}
}

func (s *sieveOfEratosthenes) NthPrime(n int64) int64 {
	upper_bound := 0
	primes := []int64{}

	// Handle edge case for negative n and return -1 for invalid input
	if n < 0 {
		return -1
	}

	// Estimate the upper limit for the nth prime using the prime number theorem
	if n < 6 {
		upper_bound = 15
	} else {
		float_n := float64(n)
		upper_bound = int(float_n * (math.Log(float_n) + math.Log(math.Log(float_n))))
	}

	// Create boolean array and initialize as true
	sieve_slice := make([]bool, upper_bound+1)
	sieve_slice[0] = false
	sieve_slice[1] = false
	for i := 2; i <= upper_bound; i++ {
		sieve_slice[i] = true
	}

	// Sieve of Eratosthenes
	for num := 2; num <= int(math.Sqrt(float64(upper_bound))); num++ {
		if sieve_slice[num] {
			for composite_index := num * num; composite_index <= upper_bound; composite_index += num {
				sieve_slice[composite_index] = false
			}
		}
	}

	// Get list of primes
	for prime_index := 0; prime_index <= upper_bound; prime_index++ {
		if sieve_slice[prime_index] {
			primes = append(primes, int64(prime_index))
		}
	}

	return primes[n]
}
