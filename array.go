package main

import "fmt"

func main() {
	var helloWorld [10]interface{}
	helloWorld[0] = "Hello"
	helloWorld[1] = "World"
	fmt.Println(helloWorld[0], helloWorld[1])
	primes := [6]int{2, 3, 4, 5, 7, 11}
	fmt.Println(primes)

	var slc = primes[1:4]
	fmt.Println(slc)

	var slc1 []int
	fmt.Println(slc1)

	slc1 = append(slc1, 7)
	fmt.Println(slc1)
}
