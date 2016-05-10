package exactlyDivisible


func IsDivisible(number_to_divide int,dividing_by int) bool{

	if(number_to_divide%dividing_by == 0 && dividing_by > 1){
		return IsDivisible(number_to_divide,dividing_by-1)
	} else if (dividing_by > 1){
		return false
	} else {
		return true
	}
}

func ListOfFactors(number_to_factor int) []int{

	var list_of_factors []int

	list_of_factors = make([]int,0)

	for i:=2; i < number_to_factor; i++ {
		if(IsDivisible(number_to_factor,i) == true){
			list_of_factors = append(list_of_factors,i)
		}
	}

	return list_of_factors
}
