//package main
package id4

import(
	"fmt"
)

func is_palindromic(number int) bool{

	var iBuffer int = number
	var iReversed int = 0

	for iBuffer != 0 {
		iReversed = iReversed *10
		iReversed = iReversed + iBuffer%10
		iBuffer = iBuffer/10
	}

	if(iReversed == number){
		return true
	} else {
		return false
	}
}

//func main(){
func Id4(){

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

			if( is_palindromic(iBuffer) && iMaxPalindrome < iBuffer){
				iMaxPalindrome = iBuffer
				pair_A = i
				pair_B = j
			}
		}
	}

	fmt.Println("The maximum palindromic number product of", pair_A,"*",pair_B,"is",iMaxPalindrome)
}
