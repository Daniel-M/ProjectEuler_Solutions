package sliceUtilities

// Calculates the product of the elements of a slice of int
func Sl_product(slice []int) int{
	var product int = 1
	for p:=0;p<len(slice);p++{
		product*=slice[p]
	}
	return product
}

// Finds the maximum value of a slice of int
func Sl_max(slice []int) int{
	var max int = slice[0]
	for k:=1;k<len(slice);k++{
		if(slice[k]>max){
			max = slice[k]
		}
	}
	return max
}
