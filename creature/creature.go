package creature

import (
	"fmt"
	"math/rand"
)

var Creatures = []string{"shark", "jellyfish", "squid", "octopus", "dolphin"}

func init() {
	fmt.Println("Creature init")
}

func Random() string {
	i := rand.Intn(len(Creatures))
	return Creatures[i]
}
