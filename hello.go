package main

import "fmt"
import "flag"

var name = flag.String("name", "Unknown", "a user name")

var greet = map[string] string {
	"Stefan": "Hello",
	"Otto": "Bonjour",
	"Heinz": "Guten Tag",
}

func main() {
	defer fmt.Println("Goodbye")
	defer fmt.Println("Almost done")
	
	flag.Parse()

	g, isthere := greet[*name]
	if ! isthere {
		g = "Good morning"
	}
	fmt.Printf("%s, %s\n", g, *name)
	µ := 3.141592653589793
	fmt.Printf("µ is equal %v\n", µ)

	var m string
	switch (*name) {
	case "Stefan": m = "Hi Boss"
	case "Karl": m = "What's up?"
	case "Fritz": m = "Weeell?"
	default: m = "Go away"
	}
	fmt.Println(m)
	var x int
	var y int32
	var z uint16
	fmt.Println(x, y, z, x + int(z))
}

