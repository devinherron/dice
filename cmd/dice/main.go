package main

import (
	"fmt"

	dice "github.com/devinherron/dice"
)

func main() {
	results, sum := dice.RollMulti(10, 2)
	fmt.Println(results)
	fmt.Println(sum)
	result := dice.Roll(6)
	fmt.Println(result)
}
