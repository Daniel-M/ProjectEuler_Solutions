// triangularNumbers contains functions to generate triangular numbers.
package triangularNumbers

// TriangularN returns the Nth triangular number. The formula is the same
// formula for the sum of the first Nth natural numbers, calculated by
// Gauss.
func TriangularN(Nth int) int {
	return Nth * (Nth + 1) / 2
}

// TriangularList returns a slice with the first Nth triangular numbers.
func TriangularList(Nth int) []int {
	var triangular_list []int

	triangular_list = make([]int, Nth)

	for i := 1; i < Nth; i++ {
		triangular_list[i] = TriangularN(i)
	}

	return triangular_list
}
