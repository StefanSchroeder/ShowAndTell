// Idea by Miek Gieben
package main

import "fmt"
import "time"

var c chan int

func cook(meal string, d int) {
	fmt.Printf("We are cooking %s today\n", meal)
	time.Sleep(time.Duration(d) * time.Second)
	fmt.Printf("%s is ready\n", meal)
	c <- 1
}

func main() {
	c = make(chan int)
	go cook("pizza", 10)
	go cook("spaghetti", 8)
	go cook("coffee", 1)
	go cook("tea", 3)

	<- c 
	<- c 
	<- c 
	<- c 
}
	
