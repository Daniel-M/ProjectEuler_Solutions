package numberStrings

import "strconv"

func StrDigits(number1 int, number2 int) []string{
	return GetStringFigures(number1 + number2)
}


func StrSliceToString(Number []string) string{
	var result_str = ""

	for i:=0;i<len(Number);i++ {
		result_str+=Number[i]
	}

	return result_str
}

func StrSliceToInt(Number []string) int{
	var result_str = ""
	var parsed_number int

	for i:=0;i<len(Number);i++ {
		result_str+=Number[i]
	}

	parsed_number,_ = strconv.Atoi(result_str)

	return parsed_number
}

func GetIntFigures(Number int) []int{
	number_str := strconv.Itoa(Number)

	var slice_figures []int

	slice_figures = make([]int,0)

	for i:=0;i<len(number_str);i++	{
		value,_ := strconv.Atoi(number_str[i:i+1])
		slice_figures = append(slice_figures,value)
	}

	return slice_figures
}

func GetStringFigures(Number int) []string{
	number_str := strconv.Itoa(Number)

	var slice_figures []string

	slice_figures = make([]string,0)

	for i:=0;i<len(number_str);i++	{
		slice_figures = append(slice_figures,number_str[i:i+1])
	}

	return slice_figures
}

func GetStringFiguresString(Number string) []string{
	var slice_figures []string

	slice_figures = make([]string,0)

	for i:=0;i<len(Number);i++	{
		slice_figures = append(slice_figures,Number[i:i+1])
	}

	return slice_figures
}

