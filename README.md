# dice

Bare bones dice roller. Will expand functionality as needed.

# Usage
```go
func Roll(size int, number int) ([]int, int)
```
size determines the sides of the dice (e.g. number of sides) and number is the number of dice rolled.
Roll returns an array of the roll results, as well as the total sum of all of the dice rolls.
