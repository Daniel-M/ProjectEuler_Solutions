package id11
import(
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/sliceUtilities"
)

func Id11(){
	fmt.Println("Solution to the problem id")

	var block int = 4

	var overAllMaximums []int

	overAllMaximums = make([]int,0)

	var array_of_numbers [20][20]int
	array_of_numbers[0] =  [20]int{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8}
	array_of_numbers[1] =  [20]int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0}
	array_of_numbers[2] =  [20]int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65}
	array_of_numbers[3] =  [20]int{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91}
	array_of_numbers[4] =  [20]int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80}
	array_of_numbers[5] =  [20]int{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50}
	array_of_numbers[6] =  [20]int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70}
	array_of_numbers[7] =  [20]int{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21}
	array_of_numbers[8] =  [20]int{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72}
	array_of_numbers[9] =  [20]int{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95}
	array_of_numbers[10] = [20]int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92}
	array_of_numbers[11] = [20]int{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57}
	array_of_numbers[12] = [20]int{86, 56, 0, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58}
	array_of_numbers[13] = [20]int{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40}
	array_of_numbers[14] = [20]int{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66}
	array_of_numbers[15] = [20]int{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69}
	array_of_numbers[16] = [20]int{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36}
	array_of_numbers[17] = [20]int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16}
	array_of_numbers[18] = [20]int{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54}
	array_of_numbers[19] = [20]int{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48}

//	fmt.Println(array_of_numbers)

	fmt.Println("The array is")

	for i:=0; i < len(array_of_numbers[0]);i++{
		fmt.Println(array_of_numbers[i])
	}

	product := 1
	maximum := 0

	// First the forward backward products and store the maximum

	for i:=0; i < len(array_of_numbers[0]);i++{
		for j:=0; j <= len(array_of_numbers[0])-block;j++{

			for k:=0; k < block ; k++ {
				product *= array_of_numbers[i][j+k]
			//	fmt.Print(array_of_numbers[i][j+k]," ")
			}
			//fmt.Println("=",product)

			if( product > maximum ){
				maximum = product
			}

			product = 1

		}

	}

	overAllMaximums = append(overAllMaximums,maximum)
	maximum = 0


	// The up-down products and store the maximum too
	for i:=0; i <= len(array_of_numbers[0])-block;i++{
		for j:=0; j < len(array_of_numbers[0]);j++{

			for k:=0; k < block ; k++ {
				product *= array_of_numbers[i+k][j]
			//	fmt.Print(array_of_numbers[i+k][j]," ")
			}
//			fmt.Println("=",product)
			if( product > maximum ){
				maximum = product
			}

			product = 1

		}

	}

	overAllMaximums = append(overAllMaximums,maximum)
	maximum = 0

	// The right-down-diagonal products and store the maximum too
	for i:=0; i <= len(array_of_numbers[0])-block;i++{
		for j:=0; j <= len(array_of_numbers[0])-block;j++{

			for k:=0; k < block ; k++ {
				product *= array_of_numbers[i+k][j+k]
			//	fmt.Print(array_of_numbers[i+k][j+k]," ")
			}
			//fmt.Println("=",product)
			if( product > maximum ){
				maximum = product
			}
			product = 1

		}

	}

	overAllMaximums = append(overAllMaximums,maximum)
	maximum = 0

	// The left-down-diagonal products and store the maximum too
	for i:=0; i <= len(array_of_numbers[0])-block;i++{
		for j:=len(array_of_numbers[0])-1; j >= block-1;j--{

			for k:=0; k < block ; k++ {
				product *= array_of_numbers[i+k][j-k]
			//	fmt.Print(array_of_numbers[i+k][j-k]," ")
			}
			//fmt.Println("=",product)
			if( product > maximum ){
				maximum = product
			}
		product = 1

		}

	}

	overAllMaximums = append(overAllMaximums,maximum)
	maximum = 0

	fmt.Println("I've found the following maxima\n",overAllMaximums)

	fmt.Println("The maximum product is",sliceUtilities.Sl_max(overAllMaximums))
}
