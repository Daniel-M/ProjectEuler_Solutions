package triangularNumbers

func TriangularN(Nth int) int{
	return Nth*(Nth+1)/2
}

func TriangularList(Nth int) []int{
	var triangular_list []int

	triangular_list = make([]int,Nth)

	for i:=1; i < Nth; i++{
		triangular_list[i] = TriangularN(i)
	}

	return triangular_list
}
