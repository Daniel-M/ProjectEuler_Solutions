package id14

import(
	"fmt"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/fibonacciNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/primeNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/exactlyDivisible"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/triangularNumbers"
//	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/palindromicNumbers"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/sliceUtilities"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/collatzConjecture"
)

func Solution(){
	fmt.Println("Solution to the problem id14")

	const limit = 1000000

//	var collatz_list []int
//	collatz_list = make([]int,0)

	collatz_map := make(map[int]int)

	for j:=2; j <= limit; j++{
		long := len(collatzConjecture.CollatzList(j))
//		collatz_list =  append(collatz_list,long)
		collatz_map[j]= long
//		fmt.Println(j,long)
	}

//	fmt.Println(collatz_map)
	key,max := sliceUtilities.Map_max(collatz_map)
	fmt.Println("The maximum length is",max,"for the number",key)
}
