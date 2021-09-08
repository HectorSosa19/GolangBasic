package main

import "fmt"

func main() {
	var name string
	var lastname string
	var pi float32
	name = "Jose"
	lastname = "Castro"
	pi = 3.14
	Number := 2
	//Se concatena por comas (",")
	fmt.Println("Nombre:", name, "\nApellido:", lastname, "\nValor de Pi:", pi)
	/*Otra manera en la que se puede concatenar es la siguiente
	%s es para agregar un valor tipo string
	%d es para agregar un valor tipo integer
	%f es para agregar un valor tipo decimal
	*/
	fmt.Printf("Nombre:%s \nApellido:%s \nValor de Pi: %f \nNumber: %d", name, lastname, pi, Number)

}
