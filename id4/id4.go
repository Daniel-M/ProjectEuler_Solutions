// #  [PROBLEM ID4](https://projecteuler.net/problem=4)
//
// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
package id4

import (
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/palindromicNumbers"
)

//func main(){
func Solution() {
	fmt.Println("Solution to the problem id4")

	var seed int = 999
	var iBuffer int
	var iMaxPalindrome int = 0
	var (
		pair_A int
		pair_B int
	)

	for i := seed; i > 100; i-- {
		for j := seed; j > 100; j-- {

			iBuffer = i * j

			if palindromicNumbers.IsPalindromic(iBuffer) && iMaxPalindrome < iBuffer {
				iMaxPalindrome = iBuffer
				pair_A = i
				pair_B = j
			}
		}
	}

	fmt.Println("The maximum palindromic number product of", pair_A, "*", pair_B, "is", iMaxPalindrome)
}
