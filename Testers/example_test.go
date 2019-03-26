package main

import (
	"fmt"
	"strings"
)

func ExampleContainsRune() {
	// Descripcion Random
	fmt.Println(strings.ContainsRune("hola", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
	// Output:
	// true
	// false
}