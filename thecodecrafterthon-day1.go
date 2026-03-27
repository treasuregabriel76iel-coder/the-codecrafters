package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const (
		ColorRed    = "\033[0;31m"
		ColorGreen  = "\033[0;32m"
		ColorYellow = "\033[0;33m"
		ColorBlue   = "\033[0;34m"
		ColorReset  = "\033[0m" // Resets the color
		ColorPurple = "\033[35m"
	)

	scanner := bufio.NewReader(os.Stdin)
	fmt.Println(ColorBlue + "********* WELCOME TO MY CLI CALCULATOR *********" + ColorReset)
	fmt.Println()
	fmt.Println(ColorYellow + "✦✦✦ TYPE 'help' FOR COMMAND MENU : ✦✦✦" + ColorReset)
	for {
		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println(ColorRed + "Error reading input." + ColorReset)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		parts := strings.Fields(input)
		operation := parts[0]
		if operation == "quit" {
			fmt.Println(ColorPurple + "Goodbye!!" + ColorReset)
			break
		}
		if operation == "help" {
			fmt.Println(ColorBlue + "Supported commands:" + ColorReset)
			fmt.Println(ColorPurple + "  + (a) (b)   → addition" + ColorReset)
			fmt.Println(ColorGreen + "  - (a) (b)   → subtraction" + ColorReset)
			fmt.Println(ColorYellow + "  * (a) (b)   → multiplication" + ColorReset)
			fmt.Println(ColorPurple + "  / (a) (b)   → division" + ColorReset)
			fmt.Println(ColorRed + "  quit          → exit calculator" + ColorReset)
			continue
		}
		if len(parts) != 3 {
			fmt.Println(ColorRed + "Error: Expected format → operation (number) (number) or type 'help' for command menu " + ColorReset)
			continue
		}
		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			fmt.Println(ColorPurple + "Error: Invalid number input." + ColorReset)
			continue
		}
		fmt.Print(ColorPurple + "✦ Answer: " + ColorReset)
		switch operation {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			if b == 0 {
				fmt.Println(ColorPurple + "Error: Division by zero." + ColorReset)
				continue
			}
			fmt.Println(a / b)
		default:
			fmt.Println(ColorYellow + "Error: Unknown operation. Type 'help'." + ColorReset)
		}
	}
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	x := 10  -5
// 	fmt.Println(x)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for {
// 		fmt.Println("Input number:")
// 		var a float64
// 		var b float64
// 		fmt.Scanln(&a)
// 		fmt.Println("SELECT OPERATION")
// 		var operation string
// 		fmt.Scan(&operation)
// 		fmt.Println("Input number:")
// 		fmt.Scanln(&b)
// 		fmt.Println("Answer: ")

// 		// fmt.Println("Input number:")

// 		// var operation string
// 		// fmt.Scan(&operation)

// 		// fmt.Println("SELECT OPERATION")
// 		// fmt.Println("1.ADDITION")
// 		// fmt.Println("2.SUBTRACTION")
// 		// fmt.Println("3.MULTIPLICATION")
// 		// fmt.Println("4.DIVISION")
// 		//addition := +
// 		// subtraction := -
// 		// multiplication := *
// 		// division := /

// 		switch operation {
// 		case "+":
// 			fmt.Println(a + b)
// 		case "-":
// 			fmt.Println(a - b)
// 		case "*":
// 			fmt.Println(a * b)
// 		case "/":
// 			if b == 0 {
// 				fmt.Println("Error: Division by zero.")
// 			}
// 			fmt.Println(a / b)
// 		case "quit":
// 			fmt.Println("Goodbye")

// 		default:
// 			fmt.Println("Error: Unknown operation.")
// 		}

// 	}
// }
// a, err1 := strconv.ParseFloat(parts[1], 64)
// b, err2 := strconv.ParseFloat(parts[2], 64)

// if err1 != nil || err2 != nil {
// 	fmt.Println("Error: Invalid number input.")
// 	continue
// }
// c, err3 := strconv.ParseFloat(parts[3], 64)
