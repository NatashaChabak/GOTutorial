package main

import "fmt"

const pi float64 = 3.1415

func main() {
	a := 16
	b := 3
	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T\n", a/b, a%b, a/b)
}
