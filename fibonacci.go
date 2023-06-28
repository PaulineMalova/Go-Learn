// Fibonacci sequence
//Fn = Fn-1 + Fn-2
package main

import "fmt"

// Iterative approach
func fib(number int) {

	first := 0
	second := 1
	next := 0
	for count:=0; count <= number; count ++ {
		fmt.Println(first)
		next = first + second
		first = second
		second = next
	}
	
}

// Recursive approach
func fibonacci() func(int) int {
	var fib func(term int) int
	fib = func(term int) int{
		if term <= 1 {
			return term
		}
		return fib(term-1) + fib(term-2)
	}
	return fib
}


func main() {
	fmt.Println("V1")
	fib(10)

	fmt.Println("\n\nV2")
	f := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(f(i))
	}
}
