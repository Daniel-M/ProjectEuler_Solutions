package main

import(
	"fmt"
)

func is_prime(number int) bool{

	for i:=2; i < number; i++{
		if( (number%i == 0)  && (number != i) ){
			return false
			fmt.Println("You shoulda not see this bc of the return above")
			break
		}
	}
	return true
}

func main(){
	var desired_number_to_reach int = 10001
	var primes_pool []int

	// Create list of prime numbers
	// and store them into slice primes_pool
	for i:=2; len(primes_pool) < desired_number_to_reach; i++{
		if(is_prime(i)){
			primes_pool = append(primes_pool,i)
			//fmt.Println(len(primes_pool))
		}
	}

	fmt.Println("The",len(primes_pool),"-th prime is",primes_pool[len(primes_pool)-1])
}
