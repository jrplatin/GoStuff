package main

import (
	"fmt"
)

func fib_rec(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else {
		return fib_rec(n-1) + fib_rec(n-2)
	}
}

func fib_mem(n int) int {
	first := 0
	second := 1
	third := 1
	if n <= 1 {
		return n
	} else {
		for i := 2; i <= n; i++  {
			third = second + first
			first = second
			second = third

		}
		return third
	}
}

func main() {
	fmt.Println(fib_rec(10))
	fmt.Println(fib_mem(10))

}
