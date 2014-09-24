package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("1 + 2 =", plus(1, 2))
}
