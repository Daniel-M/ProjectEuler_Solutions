package fibonacci

func FibonacciN(Nth int) int{
	var fibonacci_number  int = 1
	var fibonacci_number_b int = 1
	var fibonacci_seed int = 1

	for i:=1; i<Nth; i++{
		fibonacci_number_b = fibonacci_number + fibonacci_seed
		fibonacci_seed = fibonacci_number
		fibonacci_number = fibonacci_number_b
	}
	return fibonacci_number
}

func FibonacciList(Nth int) []int{
	var fibonacci_list []int

	fibonacci_list = make([]int,Nth)

	for i:=1; i < Nth; i++{
		fibonacci_list[i] = FibonacciN(i)
	}

	return fibonacci_list
}
