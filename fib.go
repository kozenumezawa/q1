package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func fib(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("You have to input positive natural number.")
	}

	f0 := 0
	f1 := 1
	switch n {
	case 0:
		return f0, nil
	case 1:
		return f1, nil
	default:
		fib := 0
		for i := 2; i <= n; i++ {
			fib = f0 + f1
			f0 = f1
			f1 = fib
		}
		return fib, nil
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
		os.Exit(1)
	}

	fib, err := fib(number)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Fibonacci F(", number, "):", fib)
}
