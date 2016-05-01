package id10

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/primenumbers"
)

//func main(){
func Id10(){
	fmt.Println("Solution to the problem id10")
	var Max int = 2000000 
	var primes_pool []int = prime_sieve.PrimeList(Max)
	var sum int = 0

	fmt.Println("I have found",len(primes_pool),"prime numbers below",Max)

	for i:=0;i < len(primes_pool) ; i++{
		sum+=primes_pool[i]
	}
}
