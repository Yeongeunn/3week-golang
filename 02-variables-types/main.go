package main

import "fmt"

func main() {
	var name string = "Alice"
	var age int = 25
	city := "Seoul"
	score := 98.5
	isStudent := true

	//`height`, `email`, `isAdmin` 변수를 추가하고 각각 출력해 보세요.
	height := 17
	email := "gmail.com"
	isAdmin := true

	fmt.Println("height:", height)
	fmt.Println("email:", email)
	fmt.Println("isAdmin", isAdmin)

	/********************************************************************/
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("city:", city)
	fmt.Println("score:", score)
	fmt.Println("isStudent:", isStudent)

	age = age + 1
	fmt.Printf("%s is %d years old next year.\n", name, age)

	fmt.Printf("type of score: %T\n", score)
	fmt.Printf("score with one decimal place: %.1f\n", score)
}
