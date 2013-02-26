package main

import (
	"fmt"
)

func Generate() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// START OMIT
func Filter(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	primes := Generate() // 2, 3, 4, 5, …
	for i := 0; i < 10; i++ {
		prime := <-primes
		fmt.Println(prime)
		primes = Filter(primes, prime)
	}
}

// END OMIT
