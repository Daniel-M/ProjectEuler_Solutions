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
/*
	var factores []int

	for i:=2; i < 50; i++{

		factores = exactlyDivisible.ListOfFactors(i)

		fmt.Println("List of factors of",i,factores)

		factores = make([]int,0)
	}
*/
	number_of_factors:= 1
	counter:=0

	number := 0 

	const limit = 20000000

	for i:=7;(number_of_factors < 5000 && counter < limit); i++{
		number = triangularNumbers.TriangularN(i)
		number_of_factors = len(exactlyDivisible.ListOfFactors(number))
		fmt.Println(i,number,number_of_factors)
		counter++
	}

	if(counter >= limit)	{
		fmt.Println("I haven't found shit")
	} else {
		fmt.Println("I have found",number,"with",number_of_factors)
	}
}

