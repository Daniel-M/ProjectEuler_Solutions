package id14

import(
	"fmt"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/fibonacciNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/primeNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/exactlyDivisible"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/triangularNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/palindromicNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/sliceUtilities"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/collatzConjecture"
)

//func main(){
func Solution(){
	fmt.Println("Solution to the problem id14")
	lista := collatzConjecture.CollatzList(13)
	fmt.Println(lista)
}
