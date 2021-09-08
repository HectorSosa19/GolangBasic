package main

import "fmt"

func main() {
	/*Declaracion de valor string,tambien se puede definir de la siguiente manera
	var name string
	name= "Hector"
	*/
	var name string = "Hector"
	/*Declaracion de valor integer,tambien se puede definir de la siguiente manera
	var age int
	age= 19
	*/
	var age, day int = 19, 20

	/* Valor flotante o valor decimal se declara de la siguiente manera */
	var Pi float32
	Pi = 3.14
	/*Otras maneras de definir variables*/
	var (
		Boolean  bool
		LastName = "Sosa"
		Phase    = "Hello World"
	)
	//Otra manera de declarar las variables pero se identifican automaticamente solo poniendole " :="
	school := "Maria inmaculada" //String
	number := 240                // Integer

	fmt.Println("My name is " + name)
	fmt.Println(Boolean, LastName, Phase)
	fmt.Println(Pi)
	fmt.Println(age, day)
	fmt.Println(school)
	fmt.Println(number)

}
