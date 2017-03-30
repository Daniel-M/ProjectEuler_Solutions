// sliceUtilities contain functions to process slices
package sliceUtilities

// Sl_produc calculates the product of the elements of a slice of int
func Sl_product(slice []int) int {
	var product int = 1
	for p := 0; p < len(slice); p++ {
		product *= slice[p]
	}
	return product
}

// Sl_max finds the maximum value of a slice of int
func Sl_max(slice []int) int {
	var max int = slice[0]
	for k := 1; k < len(slice); k++ {
		if slice[k] > max {
			max = slice[k]
		}
	}
	return max
}

// Map_max finds the maximum value of the values of a map[int]int
func Map_max(input_map map[int]int) (int, int) {
	var max int = 0
	pair := [2]int{0, 0}

	for k, v := range input_map {
		if v > max {
			pair[0] = k
			pair[1] = v
			max = v
		}
	}
	return pair[0], pair[1]
}
