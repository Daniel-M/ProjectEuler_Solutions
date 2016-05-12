package collatzConjecture

func CollatzN(Nth int) int{
	if( Nth%2 == 0){
		return Nth/2
	} else {
		return 3*Nth+1
	}
}

func CollatzList(Nth int) []int{

	var collatz_list []int
	collatz_list = make([]int,0)

	var initial_number = Nth

	collatz_list = append(collatz_list,initial_number)

	for initial_number > 1{
		initial_number = CollatzN(initial_number)
		collatz_list = append(collatz_list,initial_number)
	}

	return collatz_list
}
