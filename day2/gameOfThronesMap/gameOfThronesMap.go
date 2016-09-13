/*
Mission: Game of Thrones reloaded!

Starting with a slice of castles and a slice of kings, populate a game of Thrones
map whose key is the castle name and whose value is an adhoc slice of kings names.
Query the map and return the slice of minions if the castle is found or prints out a
failure if the castle can't be located.

Setup:
 Castles: "Red Keep", "Castle Back", "DragonStone"
 Kings:   "Melisandre", "Jon Snow", "Kal Drogo", "Gregor Clegane"

Expectations:
Looking for Red Keep yields: SUCCESS -- [Melisandre]
Looking for Red Keeper yields: FAIL! Castle Red Keeper Not Found!!
*/
package main

import "fmt"

func main() {
	got := make(map[string][]string) // Create a map with a string key and a slice value

	castles := []string{"Red Keep", "Castle Back", "DragonStone"}
	minions := []string{"Melisandre", "Jon Snow", "Kal Drogo", "Gregor Clegane"}

	for _, castle := range castles {
		// Add castle and minions to the map
		got[castle] = minions
	}

	queries := []string{"Red Keep", "Red Keeper"}
	for _, castle := range queries {
		fmt.Printf("Looking for castle `%s`...\n", castle)
		if v, ok := got[castle]; ok {
			fmt.Printf("\tSUCCESS -- %v\n", v)
		} else {
			fmt.Printf("\tFAIL! Castle `%s Not Found!!\n", castle)
		}
	}
}
