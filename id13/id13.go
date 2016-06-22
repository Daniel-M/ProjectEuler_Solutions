package id13

import (
	"fmt"
	//"github.com/Daniel-M/ProjectEuler_Solutions/utilities/fibonacciNumbers"
	//"github.com/Daniel-M/ProjectEuler_Solutions/utilities/primeNumbers"
	//"github.com/Daniel-M/ProjectEuler_Solutions/utilities/exactlyDivisible"
	//"github.com/Daniel-M/ProjectEuler_Solutions/utilities/triangularNumbers"
	//"github.com/Daniel-M/ProjectEuler_Solutions/utilities/palindromicNumbers"
	//"github.com/Daniel-M/ProjectEuler_Solutions/utilities/sliceUtilities"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/numberStrings"
)

func Solution() {
	fmt.Println("Solution to the problem id13")

	fmt.Println(numberStrings.GetIntFigures(1000))
	fmt.Println(numberStrings.GetStringFigures(1000))
	fmt.Println(numberStrings.StrSliceToString(numberStrings.GetStringFigures(1050)))
	fmt.Println(numberStrings.StrSliceToInt(numberStrings.GetStringFigures(1050)))
	fmt.Println(numberStrings.StrDigits(5, 6))

}
