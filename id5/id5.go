package main

import(
	"fmt"
)

func is_divisible(number_to_divide int,dividing_by int) bool{

	if(number_to_divide%dividing_by == 0 && dividing_by > 1){
		return is_divisible(number_to_divide,dividing_by-1)
	} else if (dividing_by > 1){
		return false
	} else {
		return true
	}
}

func main(){
	var div bool
	var condition int = 20

	for i:=2500; div != true ;i++{
		div = is_divisible(i,condition)
		if(div == true){
			fmt.Println("I've found",i)
		}
	}
}
