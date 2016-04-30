//package main
package id3

import(
	"fmt"
)

func is_prime(number int) bool{

	for i:=2; i < number; i++{
		//fmt.Println(i,number,number%i)
		if( (number%i == 0)  && (number != i) ){
			//break
			//fmt.Println(true)
			return false
		}
	}
	return true 
}

//func main(){
func Id3(){

	//var desired_number_to_factor int = 13195
	var desired_number_to_factor int = 600851475143 
	var primes_pool []int
	var primes_factors []int


	// Create list of prime numbers
	// and store them into slice primes_pool
	//for i:=1;i<=desired_number_to_factor;i++{
	for i:=1;i<=10000;i++{
		if(is_prime(i)){
			primes_pool = append(primes_pool,i)
		}
	}

	// Show list of available prime numbers
	fmt.Println(primes_pool)

	fmt.Println("Lets see the prime factors of",desired_number_to_factor)

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

	fmt.Println(primes_factors)
}
