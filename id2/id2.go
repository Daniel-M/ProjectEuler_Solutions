//package main
package id2

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/fibonacciNumbers"
)

//func main(){
func Solution(){
	fmt.Println("Solution to the problem id2")
	var sum_tot int = 0
	var counter int = 0
	var iBuffer int = 0
	const limit=4000000

	for i:=1; i < limit; i++ {

		iBuffer = fibonacciNumbers.FibonacciN(i)

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
