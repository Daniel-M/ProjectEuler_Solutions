//package main
package id4

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/palindromicnumbers"
)



//func main(){
func Id4(){
	fmt.Println("Solution to the problem id4")

	var seed int = 999
	var iBuffer int
	var iMaxPalindrome int = 0
	var (
		pair_A int
		pair_B int
	)

	for i:=seed; i>100; i--{
		for j:=seed; j>100; j--{

			iBuffer = i*j

			if( palindromicnumbers.IsPalindromic(iBuffer) && iMaxPalindrome < iBuffer){
				iMaxPalindrome = iBuffer
				pair_A = i
				pair_B = j
			}
		}
	}

	fmt.Println("The maximum palindromic number product of", pair_A,"*",pair_B,"is",iMaxPalindrome)
}
