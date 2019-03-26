package main

import "fmt"

func main() {
	n := sum(1,2)
	fmt.Println(n)

	sum5 := curry(sum, 5)
	m := sum5(2)
	fmt.Println(m)

	sum6 := curry2(sum, 5)
	m2 := sum6(2)
	fmt.Println(m2)
}

func sum(x int, y int) int {
	return x + y
}

func curry(fn func(int, int) int, v int) func(int) int {
	return func(y int) int {
		return fn(v, y)
	}
}

type IntIntToInt func(int, int) int
type IntToInt func(int) int

func curry2(fn IntIntToInt, v int) IntToInt {
	return func(y int) int {
		return fn(v, y)
	}
}