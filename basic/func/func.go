package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	q, r := div(13, 3)
	println(q, r)

	if result, err := evalError(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	println(apply(pow, 3, 3))
	println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	println(sum(1, 2, 3))
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func evalError(a, b int, op string) (int, error) {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
	return result, nil
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args:"+
		"%d, %d", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
