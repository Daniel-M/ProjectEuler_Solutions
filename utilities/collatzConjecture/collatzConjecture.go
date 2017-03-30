// collatzConjecture contains Collatz-sum related functions
package collatzConjecture

// CollatzN returns the result of the Nth term on the Collatz series.
// returns Nth/2 if Nth is even
// returns 3*Nth + 1 if Nth is odd
func CollatzN(Nth int) int {
	if Nth%2 == 0 {
		return Nth / 2
	} else {
		return 3*Nth + 1
	}
}

// CollatzList returns a slice of the Collatz-series beggining at Nth
func CollatzList(Nth int) []int {

	var collatz_list []int
	collatz_list = make([]int, 0)

	var initial_number = Nth

	collatz_list = append(collatz_list, initial_number)

	for initial_number > 1 {
		initial_number = CollatzN(initial_number)
		collatz_list = append(collatz_list, initial_number)
	}

	return collatz_list
}
