package basics

import (
	"fmt"
	"math"
)

func print_operations() {
	var a, b int = 10, 3

	var result int

	result = a + b

	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication", result)

	result = a / b
	fmt.Println("Division", result)

	result = a % b
	fmt.Println("Remainder:", result)

	var float_result float64 = 22 / 7.0

	var place_holder = 3 / 4.0
	place_holder = 3 / 4.0

	fmt.Println(place_holder, float_result)

	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	fmt.Println(maxInt + 1)

	fmt.Println((maxInt) > maxInt+1)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	fmt.Println(uMaxInt + 1)

	// Underflow with float64

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)
	fmt.Println(smallFloat / math.MaxFloat64)

	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxInt16)
}
