// primeNumbers contain a concurrent prime sieve.
// taken from https://play.golang.org/p/hgIlEQZB-2
package primeNumbers

// generate produces even numbers and sends them to a channel.
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// filter selects the prime numbers from the numbers incoming from a channel in
// and sends only the prime numbers to the out chanel.
// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// PrimeList returns a list of the prime numbers up to iLimit.
// The prime sieve: Daisy-chain Filter processes.
func PrimeList(iLimit int) []int {

	var iPrimes []int

	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Launch Generate goroutine.
	prime := 0
	for i := 1; prime < iLimit; i++ {
		iPrimes = append(iPrimes, prime)
		prime = <-ch
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

	return iPrimes
}

// IsPrime checks if a given number is a prime or not.
func IsPrime(number int) bool {

	for i := 2; i < number; i++ {
		if (number%i == 0) && (number != i) {
			return false
		}
	}
	return true
}
