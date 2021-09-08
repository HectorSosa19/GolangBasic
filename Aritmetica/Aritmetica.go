package main

import "fmt"

func main() {
	//Operadores aritmeticos
	var n1, n2 int
	fmt.Print("Enter the First Number: ")
	fmt.Scanln(&n1)
	fmt.Print("Enter the Second Number: ")
	fmt.Scanln(&n2)

	//Suma
	fmt.Println("The value total is :", n1+n2)

	/* Otra manera mas interesante de hacerlo es */
	Sum := n1 + n2
	Subtraction := n1 - n2
	Multiplication := n1 * n2
	Division := n1 / n2
	Mod := n1 % n2

	fmt.Println("The Total Sum is:", Sum)
	fmt.Println("The Total Subtraction is:", Subtraction)
	fmt.Println("The Total Multiplication is:", Multiplication)
	fmt.Println("The Total Division is:", Division)
	fmt.Println("The Total Module is:", Mod)
	fmt.Println("Are The same?:", n1 == n2)
	fmt.Println("Are different?:", n1 != n2)
}
