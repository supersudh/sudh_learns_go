// !!!WARNING!!!
// This is not feasible in golang. Do not try this

package main

import "fmt"

func main() {
	// BEGIN ERROR CODE, blunders!!!
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 289734789237423.23
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + intNum32 // invalid operation: floatNum32 + intNum32 (mismatched types float32 and int32)compilerMismatchedTypes
	// END ERROR CODE, blunders!!!
}
