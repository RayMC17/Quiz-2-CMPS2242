package main

import (
	"fmt"
	"time" //include in the package time
)

// create a function
func printAnimal(msg string) {
	fmt.Println(msg)
}

func main() {

	//two different goroutine
	//call the function with go routine
	go printAnimal("Bear")
	time.Sleep(6 * time.Second) //sleep the process for 6 sec
	go printAnimal("Tiger")
	time.Sleep(2 * time.Second)
	go printAnimal("Puppy")
	time.Sleep(6 * time.Second)
}
