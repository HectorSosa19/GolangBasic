package main

import "fmt"

func main() {
	//saludar()
	/*nombre, edad := dato()
	fmt.Println("Nombre:", nombre, "& Edad:", edad)*/
	var a, b int
	fmt.Print("Ingresa un numero:")
	fmt.Scanln(&a)
	fmt.Print("Ingresa otro numero:")
	fmt.Scanln(&b)
	suma := sumar(a, b)
	fmt.Print("El resultado es :", suma)
}

//Como crear una funcion
func saludar() {
	var nombre string
	fmt.Println("Ingresa tu nombre:")
	fmt.Scanln(&nombre)
	fmt.Print("Hola ", nombre)

}

//Funcion retornar
func dato() (string, int) {
	var nombre string
	var edad int
	fmt.Println("Ingresa tu nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("Ingresa tu edad:")
	fmt.Scanln(&edad)
	return nombre, edad
}

//Otra manera de hacer una funcion
func sumar(a int, b int) int {
	suma := a + b
	return suma

}
