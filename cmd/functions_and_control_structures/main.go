package main

import (
	"errors"
	"fmt"
)

func main() {
	printHelloWorld()

	printable := "apple"
	printArgument(printable)

	var sum int = integerAddition(5, 4)
	fmt.Println(sum)

	var result, remainder, err = integerDivision(5, 4)

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		// %v stands for value
		fmt.Printf("The result is %v\n", result)
	} else {
		fmt.Printf("The result is %v and remainder is %v\n", result, remainder)
	}
}

// function without parameters
func printHelloWorld() {
	fmt.Println("Hello World")
}

// adding parameters
func printArgument(param string) {
	fmt.Println(param)
}

// function which returns an int
func integerAddition(a int, b int) int {
	var result int = a + b
	return result
}

// function with multiple returns
func integerDivision(numerator int, denominator int) (int, int, error) {
	var err error // error data type var, default value nil

	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
