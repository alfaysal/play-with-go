package main

import (
	"fmt"
	"github.com/alfaysal/play-with-go/creature"
)

// run the imported creature package first then inside the calle function.

func init() {
	fmt.Println("Initializing from main first")
}

func init() {
	fmt.Println("Initializing from main second")
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(creature.Creatures)
}
