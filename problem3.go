/**
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143?
**/

//if n%x[i] == 0 then x is not a prime

package main

import "fmt"

func main() {
	var factorMe int = 20
	var primes []int
	var n int = 0

	//Better for loop condition
	//for factorMe/n != 1 {
	for n = 2; n <= factorMe; {
		if factorMe%n == 0 {
			primes = append(primes, n)
			factorMe = factorMe / n
		} else {
			n++
		}
	}
	fmt.Println("All prime factors are ", primes)
	fmt.Println("Largest prime factor is ", n)
}
