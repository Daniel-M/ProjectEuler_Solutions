package palindromicnumbers


func IsPalindromic(number int) bool{

	var iBuffer int = number
	var iReversed int = 0

	for iBuffer != 0 {
		iReversed = iReversed *10
		iReversed = iReversed + iBuffer%10
		iBuffer = iBuffer/10
	}

	if(iReversed == number){
		return true
	} else {
		return false
	}
}
