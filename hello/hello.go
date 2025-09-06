package main

import (
	 "fmt"
	 "log"
	 "example.com/greetings"
)
func main() {

	names := []string{"Vijay", "Vinay", "Vikas"}

	messages, error := greetings.Hellos(names)
	log.SetFlags(0)

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(messages)
}