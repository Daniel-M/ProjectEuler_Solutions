// A concurrent prime sieve
// taken from https://play.golang.org/p/hgIlEQZB-2
package prime_sieve

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func PrimeList(iLimit int) []int{

	var iPrimes []int

	iPrimes = make([]int,iLimit)

	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < iLimit; i++ {
		prime := <-ch
		iPrimes[i] = prime
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	return iPrimes
}

// Checks if a given number is prime or not
func IsPrime(number int) bool{

	for i:=2; i < number; i++{
		//fmt.Println(i,number,number%i)
		if( (number%i == 0)  && (number != i) ){
			//break
			//fmt.Println(true)
			return false
		}
	}
	return true 
}
