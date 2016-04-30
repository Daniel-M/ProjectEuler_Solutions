package exactlydivisible


func IsDivisible(number_to_divide int,dividing_by int) bool{

	if(number_to_divide%dividing_by == 0 && dividing_by > 1){
		return IsDivisible(number_to_divide,dividing_by-1)
	} else if (dividing_by > 1){
		return false
	} else {
		return true
	}
}
