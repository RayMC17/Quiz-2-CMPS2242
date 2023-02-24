package main

import (
	"fmt"
	"sync"
)

func racecar1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\n first car")
}

func racecar2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\n second car")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2) //our do something function
	//here we are increasing the counter 2 because we have 2 goroutines

	go racecar1(wg)
	go racecar2(wg)

	wg.Wait() //blocking if execution
}

func main() {

	execute()
}
