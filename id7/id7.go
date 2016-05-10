//package main
package id7

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/primeNumbers"
)

//func main(){
func Solution(){
	fmt.Println("Solution to the problem id7")
	var desired_number_to_reach int = 10001
	var primes_pool []int

	// Create list of prime numbers
	// and store them into slice primes_pool
	primes_pool = primeNumbers.PrimeList(desired_number_to_reach)

	fmt.Println("The",len(primes_pool),"-th prime is",primes_pool[len(primes_pool)-1])
}
