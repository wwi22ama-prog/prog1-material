package main

import "fmt"

func main() {
	//Foo(10)
	//fmt.Println(Add(2, 3))
	//fmt.Println(Fib(100))
	fmt.Println(Ack(4, 4))
}

func Foo(n int) {
	if n == 0 {
		return
	}
	Foo(n - 1)
	fmt.Println(n)
}

func Add(m, n int) int {
	if n == 0 {
		return m
	}
	return 1 + Add(m, n-1)
}

// Fib(1) == 1
// Fib(2) == 1
// Fib(n) == Fib(n-1) + Fib(n-2)

func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)

	//     Fib(5)   +  Fib(4)
	//       ||          ||
	//        5           3
}

// Ack(0, n) == n+1
// Ack(m, 0) == Ack(m-1,1)
// Ack(m, n) == Ack(m-1,Ack(m,n-1))

func Ack(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if n == 0 {
		return Ack(m-1, 1)
	}
	return Ack(m-1, Ack(m, n-1))
}
