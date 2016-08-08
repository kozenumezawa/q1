package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	if n >= 0 {
		var f0 int = 0
		var f1 int = 1
		switch n {
		case 0:
			return f0
		case 1:
			return f1
		default:
			var fib int = 0
			for i := 2; i <= n; i++ {
				fib = f0 + f1
				f0 = f1
				f1 = fib
			}
			return fib
		}
	} else {
		fmt.Println("You have to input positive natural number.")
		return -1
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "do nothing\n")
		os.Exit(1)
	}

	input := os.Args[1]

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fib := fib(number)
		if fib != -1 {
			fmt.Println("Fibonacci F(", number, "):", fib)
		}
	}
}
