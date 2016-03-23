package main

import(
	"fmt"
)

func sum_of_squares(Nth int) int{

	var sum_tot int = 0

	for i:=1;i<=Nth;i++{
		sum_tot+=i*i
	}
	return sum_tot
}

func square_of_sum(Nth int) int{
	var sum_tot int = 0

	for i:=1;i<=Nth;i++{
		sum_tot+=i
	}
	return sum_tot*sum_tot
}

func main(){
	var iSum_of_squares int = sum_of_squares(100)
	var iSquare_of_sum int = square_of_sum(100)

	fmt.Println("The difference between the square of the sum and the sum of squares is",iSquare_of_sum-iSum_of_squares)
}
