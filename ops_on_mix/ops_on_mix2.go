package main

import "fmt"

func ops_on_mix2() {
	// BEGIN ERROR CODE, blunders!!!
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 289734789237423.23
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println((result))
}
