package formatter

import "fmt"

func NameFormatter(name ...string) string {
	fmt.Println("Call from name formater")
	fmt.Println(name)
	return "Al faysal"
}
