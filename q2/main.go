package main

import (
	"fmt"
)

func main() {

	// creating a channel
	msg := make(chan string)

	go chanl(msg)

	fmt.Println("Channel Info:", <-msg)

}

func chanl(msg chan string) {
	//send info into channel
	msg <- " What a Success, your code has ran!"
}
