//package main
package id3

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/primeNumbers"
)

//func main(){
func Solution(){
	fmt.Println("Solution to the problem id3")

	//var desired_number_to_factor int = 13195
	var desired_number_to_factor int = 600851475143 
	var prime_factors_product int = 1
	var primes_pool []int
	var primes_factors []int


	// Create list of prime numbers
	// and store them into slice primes_pool
	primes_pool = primeNumbers.PrimeList(100000)

	// Show list of available prime numbers
	fmt.Println("A total of ", len(primes_pool),"prime numbers found")

	fmt.Println("Lets see the prime factors of",desired_number_to_factor,"under that set")

	// Lets iterate through each element of our pool of prime numbers
	// This is easy, check if the prime number divides exactly our desired number
	// In affirmative case, check what's the result of the division and repeat the process
	// with that number. This can be converted on a recursive function but meh.
	for j:=0;j < len(primes_pool);j++ {

		if(desired_number_to_factor%primes_pool[j] == 0){
			fmt.Println("Factor found",primes_pool[j])
			primes_factors = append(primes_factors,primes_pool[j])
			desired_number_to_factor = desired_number_to_factor/primes_pool[j]
			j=0
		}
	}

	for j:=0;j < len(primes_factors);j++ {
		prime_factors_product*= primes_factors[j]
	}

	fmt.Println("Set of factors found:")
	fmt.Println(primes_factors)
	fmt.Println("The product of them is",prime_factors_product)

}
