// #  [PROBLEM ID6](https://projecteuler.net/problem=6)
//
// The sum of the squares of the first ten natural numbers is,
// 12 + 22 + ... + 102 = 385
//
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)2 = 552 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural
// numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred
// natural numbers and the square of the sum.
package id6

import (
	"fmt"
)

func sum_of_squares(Nth int) int {

	var sum_tot int = 0

	for i := 1; i <= Nth; i++ {
		sum_tot += i * i
	}
	return sum_tot
}

func square_of_sum(Nth int) int {
	var sum_tot int = 0

	for i := 1; i <= Nth; i++ {
		sum_tot += i
	}
	return sum_tot * sum_tot
}

//func main(){
func Solution() {
	fmt.Println("Solution to the problem id6")
	var iSum_of_squares int = sum_of_squares(100)
	var iSquare_of_sum int = square_of_sum(100)

	fmt.Println("The difference between the square of the sum and the sum of squares is", iSquare_of_sum-iSum_of_squares)
}
