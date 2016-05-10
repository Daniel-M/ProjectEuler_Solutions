package id12

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/exactlyDivisible"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/triangularNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/sliceUtilities"
)

//func main(){
func Solution(){
	fmt.Println("Solution to the problem id12")

	var factores []int

	factores = exactlyDivisible.ListOfFactors(28)

	fmt.Println("List of factors of 28",factores)

	factores = make([]int,0)

	// since 28 is the 7th triangular number
	i := triangularNumbers.TriangularN(7)
	number_of_factors:= 1
	k:=0

	for number_of_factors < 5000{
		i = triangularNumbers.TriangularN(i+i-triangularNumbers.TriangularN(k))
		number_of_factors = len(exactlyDivisible.ListOfFactors(i))
		fmt.Println(i,number_of_factors)
		k+=10
	}

}
