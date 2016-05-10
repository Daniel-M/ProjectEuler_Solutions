package id12

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/exactlyDivisible"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/triangularNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/sliceUtilities"
)

//func main(){
func Solution(){
	fmt.Println("Solution to the problem id12")

	var factores []int

	factores = exactlyDivisible.ListOfFactors(28)

	fmt.Println("List of factors of 28",factores)

}
