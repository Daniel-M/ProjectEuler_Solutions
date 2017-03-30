// #  [PROBLEM ID1](https://projecteuler.net/problem=1)
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
package id1

import (
	"fmt"
)

func Solution() {

	fmt.Println("Solution to the problem id1")

	const limit = 1000

	var sum int

	sum = 0

	for i := 1; i < limit; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}

	fmt.Println("Total sum of multiples of 3 or 5 is", sum)
}
