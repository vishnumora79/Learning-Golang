package main

import (
	 "fmt"
	 "example.com/greetings"
)
func main() {
	message := greetings.Hello("Vinayaka")
	fmt.Println(message)
}