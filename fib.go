package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("Fibonacci Number calculator")
	fmt.Print("input number:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")

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
