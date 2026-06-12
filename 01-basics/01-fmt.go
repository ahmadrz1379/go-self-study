package main

import "fmt"

func main() { 
	name := "ali"
	age := 18
	score := 48.6850

	fmt.Println("--- Printing Variables ---")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(score)

	fmt.Println("\n--- Formatted Sentence ---")
	fmt.Printf("user: %s | Age is: %d | Score: %.2f\n", name, age, score)

	fmt.Println("\n--- Variable Types ---")
	fmt.Printf("type of score: %T\n", score)
	fmt.Printf("type of name: %T\n", name)
	fmt.Printf("type of age: %T\n", age)
}