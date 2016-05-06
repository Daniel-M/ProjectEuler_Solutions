package id9

import(
	"fmt"
	"math"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/sliceUtilities"
)

func is_whole(number float64) bool{
	if ( math.Mod(number,1) == 0){
		return true
	} else {
		return false
	}
}

func is_pitagorean(a,b int) (bool,int) {
	var result float64 = math.Sqrt(math.Pow(float64(a),2) + math.Pow(float64(b),2))
	if( is_whole(result) ){
		return true, int(result)
	} else {
		return false, -1 
	}
}

func check_sum(triad []int,desired int) (bool,[]int,int){
	var sum int = 0
	for i:=0;i<len(triad);i++{
		sum+=triad[i]
	}

	if(sum == desired){
		return true, triad, sum
	} else {
		return false , triad, sum
	}
}


func Id9(){
	fmt.Println("Solution to the problem id9")

	var number, sum_result int
	//var status, progress bool
	var status bool
	var buffer, final_triad []int
	var triad_pool []([]int)

	status = false
	i:=1
	j:=1
	//for i:=1; i < 500; i++{
	for sum_result != 1000{
		j=i+1

		for ( j < 1000 && sum_result != 1000){
		//for j:=i+1; j < 500; j++{
			status, number = is_pitagorean(i,j)

			//fmt.Println(i,j)

			if(status == true){
				buffer = append(buffer,i,j,number)
				//fmt.Println("Triad found",buffer)
				triad_pool = append(triad_pool,buffer)
				//progress, final_triad , sum_result = check_sum(buffer,1000)
				_, final_triad , sum_result = check_sum(buffer,1000)
				//fmt.Println("Triad found",final_triad,sum_result)
				buffer = make([]int,0)
			}
			j++
		}

		i++
	}

	fmt.Println("Triad found",final_triad,"sums",sum_result," the product is",sliceUtilities.Sl_product(final_triad))

}
