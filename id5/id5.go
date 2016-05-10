//package main
package id5

import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/exactlyDivisible"
)



//func main(){
func Solution(){
	fmt.Println("Solution to the problem id5")
	var div bool
	var condition int = 20

	for i:=2500; div != true ;i++{
		div = exactlyDivisible.IsEvenlyDivisible(i,condition)
		if(div == true){
			fmt.Println("I've found",i)
		}
	}
}
