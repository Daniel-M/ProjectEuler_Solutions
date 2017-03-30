// exactlyDivisible contains functions to check divisibility and multiplicative
// factors of integer numbers
package exactlyDivisible

import "sort" // To sort slices and arrays

// IsEvenlyDivisible checks if a number number_to_divide is evenly divisible
// by dividing_by. This is more than a simple modulo operation
// the function checks that all the dividing_by - 1 evenly divide
// number_to_divide with no reminder.
// for example, 12 is evenly divisible by 4 because
// 12/4 = 3
// 12/(4-1) = 12/3 = 4
// 12/(3-1) = 12/2 = 6
// Which is the same to say that
// 12%4 = 0
// 12%(4-1) = 12%3 = 0
// 12%(3-1) = 12%2 = 0
// where % is the modulo operation.
func IsEvenlyDivisible(number_to_divide int, dividing_by int) bool {

	if number_to_divide%dividing_by == 0 && dividing_by > 1 {
		return IsEvenlyDivisible(number_to_divide, dividing_by-1)
	} else if dividing_by > 1 {
		return false
	} else {
		return true
	}
}

// ListOfFactors returns a slice of the multiplicative factors of a number
// number_to_factor. For example the list of factors of 12 are {12,6,4,3,2,1}
func ListOfFactors(number_to_factor int) []int {

	var list_of_factors []int

	list_of_factors = make([]int, 0)

	up_limit := number_to_factor

	for i := 1; i < up_limit; i++ {

		if number_to_factor%i == 0 {
			list_of_factors = append(list_of_factors, i)
			up_limit = number_to_factor / i
			if up_limit > i {
				list_of_factors = append(list_of_factors, up_limit)
			}
		}
	}

	// Here the sorting takes place
	sort.Ints(list_of_factors)
	return list_of_factors
}
