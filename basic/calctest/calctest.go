package calctest

import (
	"fmt"
	"os"
)

func CalcSum(a int, b int) int {
	return a + b
}

func CalcSubtract(a int, b int) int {
	return a - b
}

func CalcMultiply(a int, b int) int {
	return a * b
}

func CalcDivide(a int, b int) int {
	return a / b
}

func opInfo(op string) {
	switch op {
	case "+":
		fmt.Println("Great! You gonna add two numbers.")
	case "-":
		fmt.Println("Great! You gonna subtract two numbers.")
	case "*":
		fmt.Println("Great! You gonna multiply two numbers.")
	case "/":
		fmt.Println("Great! You gonna divide two numbers.")
	default:
		fmt.Println("Please choose the correct operation")
		os.Exit(1)
	}
}

func opResult(o string, n1, n2 int) {
	ops := [4]string{"addition", "subtraction", "multiplying", "dividing"}

	switch o {
	case "+":
		fmt.Printf("The result of %v of %v and %v is %v\n", ops[0], n1, n2, CalcSum(n1, n2))
	case "-":
		fmt.Printf("The result of %v of %v and %v is %v\n", ops[1], n1, n2, CalcSubtract(n1, n2))
	case "*":
		fmt.Printf("The result of %v of %v and %v is %v\n", ops[2], n1, n2, CalcMultiply(n1, n2))
	case "/":
		fmt.Printf("The result of %v of %v and %v is %v\n", ops[3], n1, n2, CalcDivide(n1, n2))
	}
}

func CalctestExample() {
	var o string
	var n1 int
	var n2 int

	fmt.Println("Please choose the operation you want to perfom between two numbers by typing +, -, * or /: ")
	fmt.Scan(&o)

	opInfo(o)

	fmt.Println("Now please enter the first number: ")
	fmt.Scan(&n1)
	fmt.Println("Now please enter the second number: ")
	fmt.Scan(&n2)

	opResult(o, n1, n2)
}