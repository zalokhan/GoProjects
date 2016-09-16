/*
Mission: Game of Thrones is all about castles and their castles minions.
Write a function that takes an array of castles and an array of minions and
associates minions to castles (essentially zipping the array together) and returns
that array. You can make the assumption that both array are of the same size.
*/
package main

import "fmt"

type tuple struct {
	castle, minion string
}

func zip(castles [3]string, minions [3]string) [3]tuple {
	zipped := [3]tuple{}

	for i, castle := range castles {
		zipped[i] = tuple{castle, minions[i]}
	}
	return zipped
}

func main() {
	castles := [3]string{"Castle Black", "Red Keep", "DragonStone"}
	minions := [3]string{"melissandre", "Jon Snow", "Kal Drogo"}

	result := zip(castles, minions)
	fmt.Println(result)
}
