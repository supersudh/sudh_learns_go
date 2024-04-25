package main

import (
	"errors"
	"fmt"
)

func floatTest() {
	var printValue string = "Hello world"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
}

func printMe(str string) {
	fmt.Println(str)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainer int = numerator % denominator
	return result, remainer, err
}
