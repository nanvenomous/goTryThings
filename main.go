package main

import (
	"fmt"
	"strings"

	symdiff "github.com/mrgarelli/goTryThings/symmetricDifference"
)

func main() {
	one := "one, two"
	two := "one"
	subtractFrom := strings.Fields(one)
	toSubtract := strings.Fields(two)
	fmt.Println(symdiff.SymmetricDifference(subtractFrom, toSubtract))
}
