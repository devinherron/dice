package main

import (
	"fmt"

	dice "github.com/devinherron/dice"
)

func main() {
	results, sum := dice.Roll(10, 1)
	fmt.Println(results)
	fmt.Println(sum)
}
