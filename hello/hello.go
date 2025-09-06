package main

import (
	 "fmt"
	 "log"
	 "example.com/greetings"
)
func main() {
	message, error := greetings.Hello("")
	log.SetFlags(0)

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}