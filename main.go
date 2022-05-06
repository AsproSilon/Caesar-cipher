package main

import (
	"fmt"
)

func main() {
	// enter message and scan it , i think it is simple
	var message string
	fmt.Println("Enter the message")
	fmt.Scan(&message)
	//loop to iterate
	for i := 0; i < len(message); i++ {
		c := message[i]
		//condition : if your letter <x , everything is simple, but else you have to minus from the alphabet
		if c >= 'a' && c <= 'z' {
			c = c + 3
			if c > 'z' {
				c = c - 26

			}
		}
		fmt.Printf("%c", c)
	}
}
