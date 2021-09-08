package main

import "fmt"

func main() {
	var n1 int
	//Condicional
	fmt.Println("Ingresa un numero:")
	fmt.Scanln(&n1)
	//Todo numero dividido entre 2 si el modulo es igual a 0 entonces sera 2
	if n1 == 0 {
		fmt.Println("Es nulo")
	} else if n1%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("Es impar")
	}

}
