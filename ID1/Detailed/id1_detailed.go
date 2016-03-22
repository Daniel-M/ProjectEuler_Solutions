package main

import(
	"fmt"
)

func sum(values []int) int{
	var sum int
	sum = 0

	for i:=0; i < len(values); i++{
		sum+=values[i]
	}

	return sum
}

func main(){
	var multiples_of_3 []int
	var multiples_of_5 []int
	var multiples_of_3_and_5 []int
	const limit=1000

	for i:=1; i < limit; i++ {
		if( i%3 == 0){
			multiples_of_3 = append(multiples_of_3,i)
		}
		if(i%5 == 0){
			multiples_of_5 = append(multiples_of_5,i)
		}
		if( (i%3==0) || (i%5 == 0)){
			multiples_of_3_and_5 = append(multiples_of_3_and_5,i)
		}
	}
	
//	fmt.Println("List of multiples of 3\n",multiples_of_3,"\n")
//	fmt.Println("List of multiples of 5\n",multiples_of_5,"\n")
//	fmt.Println("List of multiples of 3 or 5\n",multiples_of_3_and_5,"\n")

	fmt.Println("The sum of multiples of 3 is",sum(multiples_of_3))
	fmt.Println("The sum of multiples of 5 is",sum(multiples_of_5))
	fmt.Println("The sum of multiples of 3 or 5 is",sum(multiples_of_3_and_5))
}
