package main

import "fmt"

func main() {
	var name string
	var lastname string
	var age int
	//Programa de entrada por teclado usando Scanln
	fmt.Print("What's your name? ")
	fmt.Scanln(&name)
	fmt.Print("What's your Last Name? ")
	fmt.Scanln(&lastname)
	fmt.Print("How old you ? ")
	fmt.Scanln(&age)
	fmt.Println("Your name is", name, ",your Last name is", lastname, "and your Age is", age)
	fmt.Println("Nice to meet you!")
}
