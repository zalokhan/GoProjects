package main

import (
	"flag"
	"fmt"
)

func main() {
	var userName = flag.String("u", "Nobody", "Pass the name of the user")
	var age = flag.Int("a", 42, "Pass the age of the user")

	flag.Parse()
	fmt.Printf("Hello, my name is %s. I am %d years old.\n", *userName, *age)
}
