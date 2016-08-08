package main

import "fmt"

func fib(n int) int {
	var f0 int = 0
	var f1 int = 1

	if n >= 0 {
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
		fmt.Printf("You have to input natural number.")
		return -1
	}
	// switch inputType := n.(type) {
	// case int:
	//
	// default:
	// 	fmt.Printf("Your input data is %d.", inputType)
	// 	fmt.Printf("You can use natural number only.")
	// }
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println("Fibonacci F(", i, "):", fib(i))
	}
}
