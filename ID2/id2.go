package main

import(
	"fmt"
)

func fibonacci_n(Nth int) int{
	var fibonacci_number  int = 1
	var fibonacci_number_b int = 1
	var fibonacci_seed int = 1

	for i:=1; i<Nth; i++{
		fibonacci_number_b = fibonacci_number + fibonacci_seed
		fibonacci_seed = fibonacci_number
		fibonacci_number = fibonacci_number_b
	}
	return fibonacci_number
}

func main(){
	var sum_tot int = 0
	var counter int = 0
	var iBuffer int = 0
	const limit=4000000

	for i:=1; i < limit; i++ {

		iBuffer = fibonacci_n(i)

		if( iBuffer%2 == 0 && iBuffer < limit){
			counter++;
			sum_tot+=iBuffer
		}
		if( iBuffer >= limit){
			break;
		}
	}

	fmt.Println("The sum of the first ",counter,"Fibonacci numbers is",sum_tot)
}
