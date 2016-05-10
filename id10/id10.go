package id10

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/primeNumbers"
)

func Solution(){
	fmt.Println("Solution to the problem id10")
	var Max int = 2000000 // This is the max bound for prime numbers below 2000000 
	var primes_pool []int
	primes_pool = primeNumbers.PrimeList(Max)
	var sum int = 0

	fmt.Println("I have found",len(primes_pool),"prime numbers below 2000000\nI'm going to calculate their sum")

	for i:=0;i < len(primes_pool) ; i++{
		sum+=primes_pool[i]
	}
	fmt.Println("Their sum is equal to",sum)
}
