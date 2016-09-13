package main

import "fmt"

type tuple struct {
	castle string
	person string
}

func zip(people [3]string, castles [3]string) [3]tuple {

	res := [len(people)]tuple{}
	for i := 0; i < len(res); i++ {
		res[i].castle = castles[i]
		res[i].person = people[i]
	}
	return res
}

func main() {
	castles := [3]string{"red", "black", "winterfell"}
	people := [3]string{"someone", "jon snow", "stark"}

	fmt.Printf("%v\n", zip(people, castles))

}
