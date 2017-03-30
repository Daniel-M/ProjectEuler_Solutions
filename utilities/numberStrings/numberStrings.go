// numberStrings contains functions that parse numbers into strings, string
// slices, and int slices
package numberStrings

import "strconv"
import "math"

//import "fmt"

// GetStringFiguresInt recives an int and makes a slice of strings with the
// figures of the number given.
// The name stands for Get String slice of Figures of an Int
// It returns a slice of strings with the numbers given.
func GetStringFiguresInt(Number int) []string {
	number_str := strconv.Itoa(Number)

	var slice_figures []string

	slice_figures = make([]string, 0)

	for i := 0; i < len(number_str); i++ {
		slice_figures = append(slice_figures, number_str[i:i+1])
	}

	return slice_figures
}

// GetIntFiguresInt recives an int and makes a slice of ints with the figures
// of the input.
// The name stands for Get Integer slice of Figures of an Int
// Returns a slice of int with the figures of the input int.
func GetIntFiguresInt(Number int) []int {
	number_str := strconv.Itoa(Number)

	var slice_figures []int

	slice_figures = make([]int, 0)

	for i := 0; i < len(number_str); i++ {
		value, _ := strconv.Atoi(number_str[i : i+1])
		slice_figures = append(slice_figures, value)
	}

	return slice_figures
}

// GetIntFiguresString recives an int and makes a slice of ints with the figures
// of the input.
// The name stands for Get Integer slice of Figures of an Int
// Returns a slice of int with the figures of the input int.
func GetIntFiguresString(StrNumber string) []int {
	return StringSliceToIntSlice(GetStringFiguresString(StrNumber))
}

// StringSliceToInt converts a slice of strings representing a number into
// the actual int number.
// Returns an int corresponding to the slice of string given as input.
func StringSliceToInt(Number []string) int {
	var result_str = ""
	var parsed_number int

	for i := 0; i < len(Number); i++ {
		result_str += Number[i]
	}

	parsed_number, _ = strconv.Atoi(result_str)

	return parsed_number
}

// StringSliceToString converts a slice of strings into a string.
// returns a string made using the ordered elements of a slice of string.
func StringSliceToString(Number []string) string {
	var result_str = ""

	for i := 0; i < len(Number); i++ {
		result_str += Number[i]
	}

	return result_str
}

// StringSliceToIntSlice converts a slice of string into a slice of int.
func StringSliceToIntSlice(StrNumberSlice []string) []int {
	var IntSlice []int

	IntSlice = make([]int, 0)

	for i := 0; i < len(StrNumberSlice); i++ {
		value, _ := strconv.Atoi(StrNumberSlice[i])
		IntSlice = append(IntSlice, value)
	}

	return IntSlice

}

// IntSliceToStringSlice converts a slice of int into a slice of string.
func IntSliceToStringSlice(IntNumberSlice []int) []string {
	var StringSlice []string

	StringSlice = make([]string, 0)

	for i := 0; i < len(IntNumberSlice); i++ {
		value := strconv.Itoa(IntNumberSlice[i])
		StringSlice = append(StringSlice, value)
	}

	return StringSlice

}

// GetStringFiguresString recives a string representing a number and makes a
// slice with the digits.
// The name stands for Get String slice of Figures of an String
// It returns a slice with the figures of the input string number.
func GetStringFiguresString(Number string) []string {
	var slice_figures []string

	slice_figures = make([]string, 0)

	for i := 0; i < len(Number); i++ {
		slice_figures = append(slice_figures, Number[i:i+1])
	}

	return slice_figures
}

// LargeNumber allows handling really big numbers by wrapping them as a slice
// A number in base 10 can be seen as the polynomial
// A(s) = sum_n a_n * s^n
// evaluated at s = 10, each a_n of the sum corresponds to
// a_0 - unities
// a_1 - decenes
// a_2 - centenes
// and so on
// Take the number 509 for example, it can be written as the polynomial
// F(s) = a_0 + a_1*s + a_2*s^2
// with a_0 = 9, a_1 = 0, a_2 = 5
// F(s) = 9 + 0 + 5*s^2  evaluating at s = 10
// F(10) = 9 + 5*(100) = 509
// The convenience comes from the fact that the number is replaced by the set
// of coefficients a_n that represent it at the polynomial form evaluated at 10
type LargeNumber struct {
	strNumber string
	intSlice  []int
}

// InitString Initializes a LargeNumber using a string.
func (lN *LargeNumber) InitString(StrIn string) {
	lN.strNumber = StrIn
	lN.intSlice = GetIntFiguresString(StrIn)
}

// InitInt initializes a LargeNumber using an int.
func (lN *LargeNumber) InitInt(IntIn int) {
	lN.strNumber = strconv.Itoa(IntIn)
	lN.intSlice = GetIntFiguresInt(IntIn)
}

// GetString returns the equivalent string of the LargeNumber.
func (lN *LargeNumber) GetString() string {
	return lN.strNumber
}

// GetIntSlice returns the slice of the coefficients of the LargeNumber.
func (lN *LargeNumber) GetIntSlice() []int {
	return lN.intSlice
}

// GetStringSlice returns a slice of strings with the coefficients of the
// LargeNumber.
func (lN *LargeNumber) GetStringSlice() []string {
	return IntSliceToStringSlice(lN.intSlice)
}

// CountFigures returns the number of digits of the LargeNumber.
func (lN *LargeNumber) CountFigures() int {
	return len(lN.intSlice)
}

// sumLargeNumbers returns the sum of two LargeNumber objects.
func SumLargeNumbers(N1, N2 *LargeNumber) LargeNumber {

	var resultLN LargeNumber
	var buffer []int
	buffer = make([]int, 0)

	if len(N1.intSlice) >= len(N2.intSlice) {

		var SizeDiff int = len(N1.intSlice) - len(N2.intSlice)

		var e0, e1 []int
		e0 = make([]int, len(N1.intSlice))
		e1 = make([]int, len(N1.intSlice)+1)

		for i := len(N1.intSlice) - 1; i >= SizeDiff; i-- {
			if N1.intSlice[i]+N2.intSlice[i-SizeDiff] > 9 {
				//e0 = append(e0, int(math.Abs(float64(N1.intSlice[i]-N2.intSlice[i-SizeDiff]))))
				//e1 = append(e1, 1)
				e0[i] = int(math.Abs(float64(N1.intSlice[i] - N2.intSlice[i-SizeDiff])))
				e1[i] = 1
			} else {
				//e0 = append(e0, N1.intSlice[i]+N2.intSlice[i-SizeDiff])
				//e1 = append(e1, 0)
				e0[i] = N1.intSlice[i] + N2.intSlice[i-SizeDiff]
				e1[i] = 0
			}
			//buffer = append(buffer, N1.intSlice[i]+N2.intSlice[i-SizeDiff])
		}
		buffer = e0
		resultLN.InitString(StringSliceToString(IntSliceToStringSlice(buffer)))
	} else {
		resultLN = SumLargeNumbers(N2, N1)
	}

	return resultLN
}
