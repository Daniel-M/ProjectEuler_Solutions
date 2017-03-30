// #  [PROBLEM ID5](https://projecteuler.net/problem=5)
//
// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
package id5

import (
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/exactlyDivisible"
)

//func main(){
func Solution() {
	fmt.Println("Solution to the problem id5")
	var div bool
	var condition int = 20

	for i := 2500; div != true; i++ {
		div = exactlyDivisible.IsEvenlyDivisible(i, condition)
		if div == true {
			fmt.Println("I've found", i)
		}
	}
}
